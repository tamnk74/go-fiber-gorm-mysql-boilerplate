package config

var MAIL = struct {
	HOST     string
	PORT     string
	USERNAME string
	PASSWORD string
	FROM     string
}{
	HOST:     getEnv("MAIL_HOST", "smtp.gmail.com"),
	PORT:     getEnv("MAIL_PORT", "587"),
	USERNAME: getEnv("MAIL_USERNAME", ""),
	PASSWORD: getEnv("MAIL_PASSWORD", ""),
	FROM:     getEnv("MAIL_FROM", "super_admin@gmail.com"),
}
