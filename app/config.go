package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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
		Host:     "golang-web-demo-postgres",
		Port:     5432,
		User:     "user1",
		Password: "pA55w0rd1",
		Name:     "golang-web-demo",
		SSLmode:  "disable",
	}
}

type Config struct {
	Port     int            `json:"port"`
	Env      string         `json:"env"`
	Pepper   string         `json:"pepper"`
	HMACKey  string         `json:"hmac_key"`
	Database PostgresConfig `json:"database"`
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Port:     3000,
		Env:      "dev",
		Pepper:   "secret-random-string",
		HMACKey:  "secret-hmac-key",
		Database: DefaultPostgresConfig(),
	}
}

func LoadConfig(configReg bool) Config {
	f, err := os.Open("./config.json")

	if err != nil {
		if configReg {
			panic(err)
		}
		fmt.Println("Useing the default config...")
		return DefaultConfig()
	}

	var c Config
	dec := json.NewDecoder(f)
	fmt.Println(dec)

	err = dec.Decode(&c)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully loaded .config.json")
	return c
}
