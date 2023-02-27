package config

import "github.com/spf13/viper"

type ApplicationCfg struct {
	DBDriver 		string `mapstructure:"LM_DB_DRIVER"`
	DBURL 			string `mapstructure:"LM_DB_URL"`
	JWTSecret 		string `mapstructure:"LM_JWT_SECRET"`
	ServerAddress 	string `mapstructure:"LM_SERVER_ADDRESS"`
}

func LoadCfg(path string) ApplicationCfg {

	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg ApplicationCfg

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg
}