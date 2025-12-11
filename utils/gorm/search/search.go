package search

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"unicode"
)

/*
BuildCondition 构建动态sql条件(and/or、exist、普通字段、json 字段)
query/body 传递json字符串 参数__search(接收参数数据类型必须是map[string]interface{})
示例1:

	__search = {
	  "or": [
		{"username": ["like", "John"]},
		{"email": ["like", "@example.com"]},
		{"exist": {"userRoles.name": "admin"}}
	  ],
	  "and": [
		{"createdAt": [">", "2025-01-01"]},
		{"createdAt": ["<", "2026-01-01"]},
		{"not exist": {"userRoles.name": ""}}
	  ]
	}

示例2:

	__search = {
	  "or": [
		{
		  "and": [
			{ "createdAt": [">", "2025-01-01"] },
			{ "createdAt": ["<", "2026-01-01"] },
			{ "not exist": { "UserRoles.name": "" } }
		  ]
		},
		{
		  "username": "admin"
		}
	  ]
	}

db := global.DB.Model(&model.User{}).Preload("UserRoles")

sql, args := BuildCondition(__search, db, model.User{})

	if sql != "" {
		db = db.Where(sql, args...)
	}
*/
func BuildCondition(filters map[string]interface{}, db *gorm.DB, model interface{}) (string, []interface{}) {
	// 必须预解析Schema否则Relationships = nil会导致preload关系解析崩溃
	// 必须传递非切片类型模型
	// 例如resolveRelation中访问curSchema.Relationships.Relations时会panic
	_ = db.Statement.Parse(&model)

	sql, args := parseLogic(filters, db, "and")
	if sql != "" {
		sql = "(" + sql + ")"
	}
	return sql, args
}

/*
parseLogic 解析and/or/exist/not exist逻辑(递归处理)

filters:
  - 若 key = "and"/"or",则遍历数组递归构建
  - 若 key = "exist"/"not exist",则走子查询逻辑
  - 其他情况为普通字段查询
*/
func parseLogic(filters map[string]interface{}, db *gorm.DB, logic string) (string, []interface{}) {
	var parts []string
	var args []interface{}

	for key, v := range filters {
		conditionKey := strings.ToLower(key)

		// 处理and/or递归逻辑
		if conditionKey == "and" || conditionKey == "or" {
			items, _ := v.([]interface{}) // 子项数组
			var subParts []string
			var subArgs []interface{}

			for _, item := range items {
				// map[string]interface{}
				// []interface{} (如["username","admin"]或["username","like","John%"])
				// 其它忽略
				if m, _ok := item.(map[string]interface{}); _ok {
					sql, a := parseLogic(m, db, conditionKey)
					if sql != "" {
						subParts = append(subParts, "("+sql+")")
						subArgs = append(subArgs, a...)
					}
					continue
				}

				if arr, _ok := item.([]interface{}); _ok && len(arr) > 0 {
					// 转换成map[string]interface{},然后递归处理
					if normalized := normalizeArrayItemToMap(arr); normalized != nil {
						sql, a := parseLogic(normalized, db, conditionKey)
						if sql != "" {
							subParts = append(subParts, "("+sql+")")
							subArgs = append(subArgs, a...)
						}
					}
					continue
				}
				// 其它类型忽略
			}

			if len(subParts) > 0 {
				// 使用and/or拼接多个(xxx)
				parts = append(parts, strings.Join(subParts, " "+strings.ToUpper(conditionKey)+" "))
				args = append(args, subArgs...)
			}
			continue
		}

		// exist/not exist子查询
		if conditionKey == "exist" || conditionKey == "not exist" {
			field, value := getExistPair(v)
			if field == "" {
				continue
			}

			sql, a := buildExistSql(db, field, conditionKey == "exist", value)
			parts = append(parts, sql)
			args = append(args, a...)
			continue
		}

		// 普通字段处理
		field := key
		op := "="
		value := v

		// 若value是[op, val]格式则提取
		if arr, ok := v.([]interface{}); ok && len(arr) > 0 {
			op, _ = arr[0].(string)
			if len(arr) > 1 {
				value = arr[1]
			}
		}

		sql, a := buildBasicSQL(db, field, op, value)
		if sql != "" {
			parts = append(parts, sql)
			args = append(args, a...)
		}
	}

	return strings.Join(parts, " "+strings.ToUpper(logic)+" "), args
}

// normalizeArrayItemToMap 将数组形式的子项转换成map[string]interface{}
// 支持：["field", value]或["field", "op", value]
func normalizeArrayItemToMap(arr []interface{}) map[string]interface{} {
	if len(arr) == 0 {
		return nil
	}
	field := fmt.Sprint(arr[0])
	// len==1 -> treat as field exists with empty value
	if len(arr) == 1 {
		return map[string]interface{}{field: []interface{}{"=", ""}}
	}
	// len==2 -> [field, value] => "="
	if len(arr) == 2 {
		return map[string]interface{}{field: []interface{}{"=", arr[1]}}
	}
	// len>=3 -> [field, op, value]
	op := fmt.Sprint(arr[1])
	val := arr[2]
	return map[string]interface{}{field: []interface{}{op, val}}
}

/*
getExistPair 提取exist/not exist的key-value

例：
"exist": {"UserRoles.name": "admin"}
返回：

	field = "UserRoles.name"
	value = "admin"
*/
func getExistPair(v interface{}) (field string, value interface{}) {
	m, ok := v.(map[string]interface{})
	if !ok {
		return "", nil
	}
	for k, val := range m {
		return k, val
	}
	return "", nil
}

/*
buildBasicSQL 构建普通字段条件(包括like/in/between等)
支持JSON_EXTRACT语法
*/
func buildBasicSQL(db *gorm.DB, field, op string, val interface{}) (string, []interface{}) {
	// JSON_EXTRACT JSON字段($.xxx 或 json->xxx)
	if strings.HasPrefix(field, "$.") || strings.HasPrefix(field, "json->") {
		return buildJsonSql(field, op, val)
	}

	// 判断是否为关联字段 a.b
	if strings.Contains(field, ".") {
		// 自动改写为关联 exist 子查询
		return buildRelationFieldSql(db, field, op, val)
	}

	// 字段名转 snake_case
	snakeField := camelToSnake(field)
	// 普通操作符
	col := columnName(db, snakeField)
	op = strings.ToLower(op)

	// 构建常规sql
	switch op {
	case "=", "eq":
		return fmt.Sprintf("%s = ?", col), []interface{}{val}

	case "!=", "<>", "ne":
		return fmt.Sprintf("%s <> ?", col), []interface{}{val}

	case ">", "gt":
		return fmt.Sprintf("%s > ?", col), []interface{}{val}

	case ">=", "gte":
		return fmt.Sprintf("%s >= ?", col), []interface{}{val}

	case "<", "lt":
		return fmt.Sprintf("%s < ?", col), []interface{}{val}

	case "<=", "lte":
		return fmt.Sprintf("%s <= ?", col), []interface{}{val}

	case "in":
		return fmt.Sprintf("%s IN (?)", col), []interface{}{val}

	case "not in":
		return fmt.Sprintf("%s NOT IN (?)", col), []interface{}{val}

	case "is null":
		return fmt.Sprintf("%s IS NULL", col), nil

	case "is not null":
		return fmt.Sprintf("%s IS NOT NULL", col), nil

	case "like":
		return fmt.Sprintf("%s LIKE ?", col), []interface{}{"%" + fmt.Sprint(val) + "%"}

	case "left like":
		return fmt.Sprintf("%s LIKE ?", col), []interface{}{fmt.Sprint(val) + "%"}

	case "right like":
		return fmt.Sprintf("%s LIKE ?", col), []interface{}{"%" + fmt.Sprint(val)}

	case "between":
		arr, ok := val.([]interface{})
		if ok && len(arr) == 2 {
			return fmt.Sprintf("%s BETWEEN ? AND ?", col), []interface{}{arr[0], arr[1]}
		}

	case "not between":
		arr, ok := val.([]interface{})
		if ok && len(arr) == 2 {
			return fmt.Sprintf("%s NOT BETWEEN ? AND ?", col), []interface{}{arr[0], arr[1]}
		}
	}

	// 默认"="
	return fmt.Sprintf("%s = ?", col), []interface{}{val}
}

// buildRelationFieldSql 自动将a.b=x转成EXISTS子查询
func buildRelationFieldSql(db *gorm.DB, field string, op string, val interface{}) (string, []interface{}) {
	table, column := resolveRelation(db, field)

	sub := fmt.Sprintf("SELECT 1 FROM %s WHERE %s = ?", table, column)

	op = strings.ToLower(op)
	switch op {
	case "=", "eq":
		return fmt.Sprintf("EXISTS (%s)", sub), []interface{}{val}
	case "!=", "<>", "ne":
		return fmt.Sprintf("NOT EXISTS (%s)", sub), []interface{}{val}
	case "in":
		return fmt.Sprintf("EXISTS (%s IN (?))", sub), []interface{}{val}
	}

	return fmt.Sprintf("EXISTS (%s)", sub), []interface{}{val}
}

/*
buildJsonSql 构建json查询语句(JSON_EXTRACT,JSON_CONTAINS等)
支持：
- $.meta.flag
- json->meta.flag
*/
func buildJsonSql(field, op string, val interface{}) (string, []interface{}) {
	var col, path string

	// $.meta.flag
	if strings.HasPrefix(field, "$.") {
		parts := strings.Split(field, ".")
		col = camelToSnake(parts[1])
		path = "$." + strings.Join(parts[2:], ".")
	}

	// json->meta.flag
	if strings.HasPrefix(field, "json->") {
		f := strings.TrimPrefix(field, "json->")
		parts := strings.Split(f, ".")
		col = camelToSnake(parts[0])
		path = "$." + strings.Join(parts[1:], ".")
	}

	op = strings.ToLower(op)

	switch op {
	case "json_contains":
		return fmt.Sprintf("JSON_CONTAINS(%s, ?)", col), []interface{}{val}
	case "in":
		return fmt.Sprintf("JSON_EXTRACT(%s, '%s') IN (?)", col, path), []interface{}{val}
	default:
		return fmt.Sprintf("JSON_EXTRACT(%s, '%s') = ?", col, path), []interface{}{val}
	}
}

/*
buildExistSql 构建EXISTS/NOT EXISTS子查询(支持preload关系)

例：
exist: {"UserRoles.name": "admin"}

生成 SQL：
EXISTS (

	SELECT 1 FROM user_roles WHERE user_roles.name = ?

)
*/
func buildExistSql(db *gorm.DB, field string, positive bool, value interface{}) (string, []interface{}) {
	// 从preload关系中推导子表名和字段名
	table, column := resolveRelation(db, field)

	sub := fmt.Sprintf("SELECT 1 FROM %s WHERE %s = ?", table, column)

	if positive {
		return fmt.Sprintf("EXISTS (%s)", sub), []interface{}{value}
	}
	return fmt.Sprintf("NOT EXISTS (%s)", sub), []interface{}{value}
}

/*
resolveRelation 解析preload关系链,自动找到最终的表名与字段名

示例：
UserRoles.name → user_roles.name
Profile.Avatar.url → profile.avatar.url(通过gorm schema推导)
*/
//func resolveRelation(db *gorm.DB, field string) (subTable string, column string) {
//	parts := strings.Split(field, ".")
//
//	curSchema := db.Statement.Schema
//	// 逐级关系解析：UserRoles → user_roles 表
//	for i := 0; i < len(parts)-1; i++ {
//		name := parts[i]
//		if rel, ok := curSchema.Relationships.Relations[name]; ok {
//			curSchema = rel.FieldSchema // 下一级schema
//		}
//	}
//
//	// 当前schema最终表
//	subTable = curSchema.Table
//	// 最后一个字段,如name → name
//	column = camelToSnake(parts[len(parts)-1])
//	return
//}
func resolveRelation(db *gorm.DB, field string) (subTable string, column string) {
	parts := strings.Split(field, ".")
	if len(parts) == 0 {
		return "", ""
	}

	// 统一格式化关联名user_roles → UserRoles
	parts[0] = toPascalCase(parts[0])

	curSchema := db.Statement.Schema

	// 按顺序解析关系链,例如UserRoles → Profile → Avatar
	for i := 0; i < len(parts)-1; i++ {
		name := parts[i]

		if rel, ok := curSchema.Relationships.Relations[name]; ok {
			curSchema = rel.FieldSchema
			continue
		}

		// 若存在小写(容错)
		lower := strings.ToLower(name)
		if rel, ok := curSchema.Relationships.Relations[lower]; ok {
			curSchema = rel.FieldSchema
			continue
		}

		// 若存在下划线(容错)
		snake := camelToSnake(name)
		if rel, ok := curSchema.Relationships.Relations[snake]; ok {
			curSchema = rel.FieldSchema
			continue
		}

		// 未找到关系
		break
	}

	return curSchema.Table, camelToSnake(parts[len(parts)-1])
}

// toPascalCase 转大驼峰命名
func toPascalCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if strings.Contains(s, "_") {
		s = snakeToCamel(s)
	}

	// 处理首字母小写
	if unicode.IsLower(rune(s[0])) {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

/*
columnName 返回完整的表名.字段名
如：users.username
*/
func columnName(db *gorm.DB, field string) string {
	return db.Statement.Table + "." + field
}

// 将驼峰命名转为下划线命名
func camelToSnake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// 将下划线命名转为驼峰命名
func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + strings.ToLower(parts[i][1:])
		}
	}
	return strings.Join(parts, "")
}
