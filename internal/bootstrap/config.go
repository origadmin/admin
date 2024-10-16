package bootstrap

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/godcong/dl"
	ggb "github.com/goexts/ggb"
	"github.com/origadmin/toolkits/codec"

	"origadmin/backend/internal/errors"
)

type Type string

const (
	TypeFile       Type = "file"
	TypeConsul     Type = "consul"
	TypeKubernetes Type = "kubernetes"
)

type Consul struct {
	Addr string `json:"addr" toml:"addr" yaml:"addr" default:"127.0.0.1:8500"`
}

type Config struct {
	Type   Type   `json:"type" toml:"type" yaml:"type" default:"file"`
	Name   string `json:"name" toml:"name" yaml:"name" default:"boostrap.toml"`
	Consul Consul `json:"consul" toml:"consul" yaml:"consul"`
}

func (c *Config) String() string {
	typo := codec.TypeFromExt(filepath.Ext(c.Name))
	return string(ggb.Must(typo.Marshal(c)))
}
func NewConfig() *Config {
	return &Config{
		Type: "file",
	}
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
			return nil, errors.LoadError(err, name, "failed to state file")
		}

		if info.IsDir() {
			err := filepath.WalkDir(fullname, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return errors.LoadError(err, name, "failed to get config file")
				} else if d.IsDir() {
					return nil
				}
				return parseConfigFile(&c, path)
			})
			if err != nil {
				return nil, errors.LoadError(err, name, "failed to walk config dir")
			}
			continue
		}
		if err := parseConfigFile(&c, fullname); err != nil {
			return nil, err
		}
	}
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
	err := codec.DecodeFromFile(path, c)
	if err != nil {
		return err
	}

	return nil
}
