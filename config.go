package main

import "fmt"

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"post"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SSLmode  string `json:"sslmode"`
}

func (c PostgresConfig) ConnectionInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Password, c.Name, c.SSLmode)
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "ec2-23-23-184-76.compute-1.amazonaws.com",
		Port:     5432,
		User:     "hrgcmhzjkgllyf",
		Password: "f867d132e78e27e50a27d0b7522dbf3f44dc835c903eb3040d74ecd5daf5c633",
		Name:     "d61hvpjfrp6em7",
		SSLmode:  "require",
	}
}

type Config struct {
	Port    int    `json:"port"`
	Env     string `json:"env"`
	Pepper  string `json:"pepper"`
	HMACKey string `json:"hmac_key"`
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Port:    3000,
		Env:     "dev",
		Pepper:  "secret-random-string",
		HMACKey: "secret-hmac-key",
	}
}
