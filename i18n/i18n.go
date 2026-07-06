package i18n

import (
	"embed"
	"encoding/json"
)

//go:embed zh.json en.json
var locales embed.FS

var localesCache map[string]map[string]string

func init() {
	localesCache = make(map[string]map[string]string)
	for _, lang := range []string{"zh", "en"} {
		data, err := locales.ReadFile(lang + ".json")
		if err != nil {
			panic("Cannot read locale file: " + lang + ".json")
		}
		msgs := make(map[string]string)
		if err := json.Unmarshal(data, &msgs); err != nil {
			panic("Cannot parse locale file: " + lang + ".json")
		}
		localesCache[lang] = msgs
	}
}

type Locale struct {
	Messages map[string]string
}

func GetLocale(lang string) *Locale {
	if msgs, ok := localesCache[lang]; ok {
		return &Locale{Messages: msgs}
	}
	return &Locale{Messages: localesCache["zh"]}
}

func T(key, lang string) string {
	if msgs, ok := localesCache[lang]; ok {
		if msg, ok2 := msgs[key]; ok2 {
			return msg
		}
	}
	if msgs, ok := localesCache["zh"]; ok {
		if msg, ok2 := msgs[key]; ok2 {
			return msg
		}
	}
	return key
}
