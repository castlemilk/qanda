package question

import (
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config provided to the Echo service for initialisation
type Config struct {
	Insecure    bool   `json:"insecure"        yaml:"insecure"     mapstructure:"insecure"`
	Port        int    `json:"port"            yaml:"port"         mapstructure:"port" validate:"gte=8000"`
	Credentials string `json:"credentials" yaml: "credentials" mapstructure:"credentials" validate:"required"`
}

const (
	defaultInsecure    = true
	defaultPort        = 8080
	defaultConcurrency = 1
)

func DefaultConfig() Config {
	return Config{
		Insecure: defaultInsecure,
		Port:     defaultPort,
		Credentials: "",
	}
}

// Env binds environment variables to main configuration struct populated by viper
func Env(v *viper.Viper, cfgPrefix, envPrefix string) {
	_ = v.BindEnv(fmt.Sprintf("%s.port", cfgPrefix), fmt.Sprintf("%s_PORT", envPrefix))
	_ = v.BindEnv(fmt.Sprintf("%s.insecure", cfgPrefix), fmt.Sprintf("%s_INSECURE", envPrefix))
	_ = v.BindEnv(fmt.Sprintf("%s.credentials", cfgPrefix), "GOOGLE_APPLICATION_CREDENTIALS")
}

// Flags binds flags to main configuration struct populated by viper
func Flags(f *flag.FlagSet, prefix string) {
	f.IntP(fmt.Sprintf("%s.port", prefix), "p", defaultPort, "default port agent will listen on")
}
