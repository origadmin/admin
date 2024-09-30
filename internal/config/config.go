package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/godcong/dl"
	"github.com/origadmin/toolkits/codec"
)

type Config struct {
	Dir         string // From command arguments
	Debug       bool
	PrintConfig bool
	Settings    Settings
}

func (c *Config) Print() {
	if !c.PrintConfig {
		return
	}
	fmt.Print(UI)
	fmt.Println("// ----------------------- Load configurations start ------------------------")
	_, _ = fmt.Println(c)
	fmt.Println("// ----------------------- Load configurations end --------------------------")
}

type configErr struct {
	name    string
	message string
	error
}

func (c configErr) Error() string {
	return fmt.Sprintf("file %s: %s,error:%v", c.name, c.message, c.error)
}

func LoadError(err error, name string, message string) error {
	return configErr{name: name, message: message, error: err}
}

func Load(dir string, names ...string) (*Config, error) {
	var c Config
	err := dl.Load(&c)
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		fullname := filepath.Join(dir, name)
		info, err := os.Stat(fullname)
		if err != nil {
			return nil, LoadError(err, name, "failed to state file")
		}

		if info.IsDir() {
			err := filepath.WalkDir(fullname, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return LoadError(err, name, "failed to get config file")
				} else if d.IsDir() {
					return nil
				}
				return parseConfigFile(&c, path)
			})
			if err != nil {
				return nil, LoadError(err, name, "failed to walk config dir")
			}
			continue
		}
		if err := parseConfigFile(&c, fullname); err != nil {
			return nil, err
		}
	}
	c.Dir = dir
	return &c, nil
}

const (
	extNames = `.json,.toml,.yaml",.yml`
)

func parseConfigFile(c *Config, path string) error {
	ext := filepath.Ext(path)
	if ext == "" || !strings.Contains(extNames, ext) {
		return nil
	}
	err := codec.DecodeFile(path, c)
	if err != nil {
		return err
	}

	return nil
}