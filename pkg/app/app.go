package app

import "IMfourm-go/pkg/config"

func IsLocal() bool{
	return config.Get("app.env") == "local"
}
func IsProduction() bool{
	return config.Get("app.env") == "production"
}
func IsTesting() bool{
	return config.Get("app.env") == "testing"
}
