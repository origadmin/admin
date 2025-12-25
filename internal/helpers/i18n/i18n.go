package i18n

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path"
	"strings"
	"sync"
)

//go:embed locales/*.json
var localesFS embed.FS

const (
	// DefaultLang defines the language used for build-time tasks and as a fallback.
	DefaultLang = "en"
)

var (
	globalManager *Manager
	once          sync.Once
)

// Manager handles loading and retrieving translations.
type Manager struct {
	translations map[string]map[string]string // map[lang]map[key]value
}

// initManager loads all locale files from the embedded filesystem.
func initManager() {
	once.Do(func() {
		m := &Manager{
			translations: make(map[string]map[string]string),
		}

		files, err := fs.ReadDir(localesFS, "locales")
		if err != nil {
			// A broken build should stop the process.
			panic("i18n: failed to read embedded locales directory: " + err.Error())
		}

		for _, file := range files {
			if file.IsDir() || path.Ext(file.Name()) != ".json" {
				continue
			}

			lang := strings.TrimSuffix(file.Name(), ".json")
			content, err := fs.ReadFile(localesFS, path.Join("locales", file.Name()))
			if err != nil {
				panic("i18n: failed to read locale file " + file.Name() + ": " + err.Error())
			}

			var translationsForLang map[string]string
			if err := json.Unmarshal(content, &translationsForLang); err != nil {
				panic("i18n: failed to parse locale file " + file.Name() + ": " + err.Error())
			}
			m.translations[lang] = translationsForLang
		}
		globalManager = m
	})
}

// Text returns the translation for the given key in the **default language**.
// This function is 100% backward compatible and is used for build-time tasks
// like `go generate` and as the default for runtime usage.
// IT MUST NOT BE RENAMED OR HAVE ITS SIGNATURE CHANGED.
func Text(key string) string {
	initManager()
	if defaultTranslations, ok := globalManager.translations[DefaultLang]; ok {
		if val, ok := defaultTranslations[key]; ok {
			return val
		}
	}
	return key // Fallback to the key itself.
}

// TextFor returns the translation for a given key in a specific language.
// This is the new function to be used at RUNTIME for multi-language support.
func TextFor(lang, key string) string {
	initManager()
	// Try to get the translation for the requested language.
	if langTranslations, ok := globalManager.translations[lang]; ok {
		if val, ok := langTranslations[key]; ok {
			return val
		}
	}
	// Fallback to the default language.
	return Text(key)
}
