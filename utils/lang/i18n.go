package lang

import (
	"fmt"
	"gin/config"
	"gin/utils/ctx"
	"github.com/goccy/go-json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

var (
	Bundle     *i18n.Bundle
	Localizers = map[string]*i18n.Localizer{}
)

// LoadLang 初始化翻译
func LoadLang() {
	Bundle = i18n.NewBundle(language.Chinese)
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	Bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	baseDir := config.Conf.I18n.Dir
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		config.ZapLogger.Info(nil, fmt.Sprintf("i18n baseDir not found: %s", baseDir))
		return
	}

	langs := strings.Split(config.Conf.I18n.Lang, ",")

	// 遍历语言目录
	for _, lang := range langs {
		langDir := filepath.Join(baseDir, lang)
		loadLangDir(lang, langDir)
	}

	// 初始化Localizer
	for _, lang := range langs {
		Localizers[lang] = i18n.NewLocalizer(Bundle, lang)
	}
}

// loadLangDir 递归加载指定语言目录下的所有翻译文件
func loadLangDir(lang, dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			config.ZapLogger.Info(nil, err.Error())
			return nil
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".json" && ext != ".yaml" && ext != ".yml" {
			config.ZapLogger.Info(nil, "Unsupported lang file type: "+ext)
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			config.ZapLogger.Info(nil, err.Error())
		}

		// 模拟路径格式如zh.json/en.yaml,让go-i18n能识别语言
		virtualFileName := fmt.Sprintf("%s%s", lang, ext)
		_, err = Bundle.ParseMessageFileBytes(data, virtualFileName)
		if err != nil {
			config.ZapLogger.Info(nil, err.Error())
		}

		return nil
	})
	if err != nil {
		config.ZapLogger.Info(nil, err.Error())
	}
}

// T 翻译
func T(messageID string, data map[string]interface{}) string {
	context := ctx.GetContext(ctx.KeyLang)
	langCode := context.GetString(ctx.KeyLang)
	if langCode == "" {
		langCode = "zh"
	}

	localizer, ok := Localizers[langCode]
	if !ok {
		localizer = Localizers["zh"]
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
	if err != nil {
		config.ZapLogger.Info(nil, fmt.Sprintf("缺少翻译: %s (%s)\n", messageID, langCode))
		return messageID
	}
	return msg
}
