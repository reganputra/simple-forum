package configs

import "github.com/spf13/viper"

var config *Config

type option struct {
	configFolder []string
	configFile   string
	configType   string
}

func Init(opts ...Option) error {
	opt := &option{
		configFolder: getDefaultConfigFolders(),
		configFile:   getDefaultConfigFile(),
		configType:   getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	for _, configFolder := range opt.configFolder {
		viper.AddConfigPath(configFolder)
	}
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	config = new(Config)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}

	return nil
}

type Option func(*option)

func getDefaultConfigFolders() []string {
	return []string{"./configs"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(folder []string) Option {
	return func(opt *option) {
		opt.configFolder = folder
	}
}

func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
