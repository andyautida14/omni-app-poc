package main

type ServiceConfig struct {
	DbUrl         string `env:"DB_URL, required"`
	StaticPath    string `env:"STATIC_PATH"`
	TemplatePath  string `env:"TEMPLATE_PATH"`
	TemplateCache string `env:"TEMPLATE_CACHE, default=lazy"`
}
