package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type Default[T any] interface {
	Default() T
}

type Config struct {
	// path of the config file
	ConfigFile string

	// Prefix of environment variables
	EnvPrefix string

	// Delim of environment variables
	EnvDelim string
}

func (c Config) Default() Config {
	return Config{
		ConfigFile: "config.toml",
		EnvPrefix:  "",
		EnvDelim:   "__",
	}
}

func mergeDefault(a, b map[string]any) {
	for k, v := range a {
		bv, ok := b[k]

		if !ok {
			b[k] = v
			continue
		}

		switch bv := bv.(type) {
		case map[string]any:
			mergeDefault(v.(map[string]any), bv)
		default:
			if !reflect.ValueOf(v).IsZero() {
				b[k] = v
			}
		}
	}
}

// GetConfig reads configurations from default values, toml files, env vars.
//
// `c` should be passed as reference in order to get the configuration
//
// reading/overriding orders:
//  1. Environment variables
//  2. Config file
//  3. Default values
//
// example:
//
// given the following go code and setting EnvPrefix = "example", EnvDelim = "__"
//
//	     ```go
//
//		     type Example struct {
//		         Field FieldStruct `koanf:field`
//		     }
//
//		     type FeildStruct struct {
//		         Foo string `koanf:foo`
//		     }
//
//	       func (e Example) Default() ConfigInterface {...}
//
//	     ```
//
// config file will be like:
//
//	     ```toml
//
//	     [field]
//
//		    foo = bar
//
//	     ```
//
// the config can be overwritten by env vars like:
//
//	`example__field_foo=bar`, `EXAMPLE__field_foo=bar`, `Example__field_foo=bar`
func GetConfig[T any](conf Config, c Default[T]) (err error) {
	k := koanf.New(".")

	// Load default config.
	if err = k.Load(structs.Provider(c.Default(), "koanf"), nil); err != nil {
		return err
	}

	if !reflect.ValueOf(c).IsZero() {
		if err = k.Load(
			structs.Provider(c, "koanf"),
			nil,
			koanf.WithMergeFunc(func(src, dst map[string]any) error {
				mergeDefault(src, dst)
				return nil
			}),
		); err != nil {
			return err
		}
	}

	// Note: We chose Toml instead of Yaml because Yaml unmarshaller has some quirky behavior
	// on empty sections, which will override default values.
	// Load Toml config.
	if _, e := os.Stat(conf.ConfigFile); e == nil {
		if err = k.Load(file.Provider(conf.ConfigFile), toml.Parser()); err != nil {
			return err
		}
	}

	// Allows quick Env overwrite.
	if err = k.Load(env.Provider(conf.EnvPrefix, conf.EnvDelim, strings.ToLower), nil); err != nil {
		return err
	}

	if err = k.Unmarshal("", &c); err != nil {
		return err
	}

	return err
}
