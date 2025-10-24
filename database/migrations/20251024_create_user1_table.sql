CREATE TABLE `user1` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
                        `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
                        `username` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
                        `full_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓名',
                        `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
                        `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
                        `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
                        `gender` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '性别 1=男 2=女',
                        `age` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
                        `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态 1=启用 2=停用',
                        `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                        `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                        `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                        PRIMARY KEY (`id`) USING BTREE,
                        UNIQUE KEY `uniq_idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `uniq_idx_email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';