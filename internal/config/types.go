package config

type (
	Config struct {
		Service Service `mapstructure:"service"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}
)
