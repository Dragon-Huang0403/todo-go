package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type ExampleConfig struct {
	Nested NestedConfig `koanf:"nested"`
}

func (e ExampleConfig) Default() ExampleConfig {
	return ExampleConfig{
		Nested: NestedConfig{}.Default(),
	}
}

type NestedConfig struct {
	Foo string `koanf:"foo"`
}

func (f NestedConfig) Default() NestedConfig {
	return NestedConfig{
		Foo: "defaultFoo",
	}
}

func TestDefaultConfig(t *testing.T) {
	var c ExampleConfig
	conf := Config{}.Default()

	err := GetConfig(conf, &c)
	require.NoError(t, err)
	require.Equal(t, "defaultFoo", c.Nested.Foo)
}

func TestConfigFromFile(t *testing.T) {
	configContent := `
	[nested]
	foo = "fileFoo"
	`
	configFile := "test_config.toml"
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	require.NoError(t, err)
	defer os.Remove(configFile)

	var c ExampleConfig
	conf := Config{
		ConfigFile: configFile,
	}

	err = GetConfig(conf, &c)
	require.NoError(t, err)
	require.Equal(t, "fileFoo", c.Nested.Foo)
}

func TestConfigFromEnv(t *testing.T) {
	err := os.Setenv("NESTED__FOO", "envFoo")
	require.NoError(t, err)
	defer os.Unsetenv("NESTED__FOO")

	var c ExampleConfig
	conf := Config{}.Default()

	err = GetConfig(conf, &c)
	require.NoError(t, err)
	require.Equal(t, "envFoo", c.Nested.Foo)
}

func TestConfigMerge(t *testing.T) {
	configContent := `
	[nested]
	foo = "fileFoo"
	`
	configFile := "test_config.toml"
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	require.NoError(t, err)
	defer os.Remove(configFile)

	err = os.Setenv("NESTED__FOO", "envFoo")
	require.NoError(t, err)
	defer os.Unsetenv("NESTED__FOO")

	var c ExampleConfig
	conf := Config{
		ConfigFile: configFile,
		EnvDelim:   "__",
	}

	err = GetConfig(conf, &c)
	require.NoError(t, err)
	require.Equal(t, "envFoo", c.Nested.Foo)
}

func TestConfigNoOverrides(t *testing.T) {
	var c ExampleConfig
	conf := Config{
		ConfigFile: "nonexistent.toml",
		EnvDelim:   "__",
	}

	err := GetConfig(conf, &c)
	require.NoError(t, err)
	require.Equal(t, "defaultFoo", c.Nested.Foo)
}

func TestMergeDefault(t *testing.T) {
	a := map[string]any{
		"NESTED": map[string]any{
			"foo": "defaultFoo",
		},
	}
	b := map[string]any{
		"NESTED": map[string]any{
			"foo": "",
		},
	}

	mergeDefault(a, b)
	require.Equal(t, "defaultFoo", b["NESTED"].(map[string]any)["foo"])
}
