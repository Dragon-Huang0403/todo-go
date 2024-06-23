package httpserver

import "time"

type Config struct {
	AddrPort        string        `koanf:"addr_port" validate:"required,tcp4_addr"`
	ShutdownTimeout time.Duration `koanf:"shutdown_timeout" validate:"required"`
}

func (Config) Default() Config {
	return Config{
		AddrPort:        "127.0.0.1:8080",
		ShutdownTimeout: 5 * time.Second,
	}
}
