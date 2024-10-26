package main

type ServiceConfig struct {
	DbUrl          string `env:"DB_URL, required"`
	StaticPath     string `env:"STATIC_PATH"`
	TemplatePath   string `env:"TEMPLATE_PATH"`
	CacheTemplates bool   `env:"CACHE_TEMPLATES, default=true"`
}
