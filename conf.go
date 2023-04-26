package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
)

// Config represents application config
type Config struct {
	NutsUrl     string `env:"NUTS_URL" validate:"required"`
	NutsSubject string `env:"NUTS_SUBJECT" validate:"required"`
	PoolSize    int    `env:"POOL_SIZE" validate:"required,min=1"`
}

func NewConfig() *Config {
	return &Config{}
}

func (r *Config) Read() error {
	if err := env.ParseWithFuncs(r, nil); err != nil {
		return err
	}

	// Validating the config values. For example, empty\default values are NOT
	// allowed
	if err := validator.New().Struct(r); err != nil {
		return err
	}

	return nil
}
