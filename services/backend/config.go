package main

type ServiceConfig struct {
	DbUrl string `env:"DB_URL, required"`
}
