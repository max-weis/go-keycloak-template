package config

import "os"

type Config struct {
	Port         string
	KeycloakHost string
	ClientId     string
	ClientSecret string
	Realm        string
	AdminUser    string
	AdminPass    string
}

func NewConfig() Config {
	return Config{
		Port:         getEnvVar("PORT", ":8000"),
		KeycloakHost: getEnvVar("KEYCLOAK_HOST", "http://localhost:8080"),
		ClientId:     getEnvVar("KEYCLOAK_CLIENT_ID", "go-demo"),
		ClientSecret: getEnvVar("KEYCLOAK_CLIENT_SECRET", ""),
		Realm:        getEnvVar("KEYCLOAK_REALM", "master"),
		AdminUser:    getEnvVar("KEYCLOAK_ADMIN_USER", "admin"),
		AdminPass:    getEnvVar("KEYCLOAK_ADMIN_PASS", "admin"),
	}
}

func getEnvVar(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}
