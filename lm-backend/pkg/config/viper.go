package config

import "github.com/spf13/viper"

type ApplicationCfg struct {
	DBDriver 		string `mapstructure:"LM_DB_DRIVER"`
	DBURL 			string `mapstructure:"LM_DB_URL"`
	JWTSecret 		string `mapstructure:"LM_JWT_SECRET"`
	ServerAddress 	string `mapstructure:"LM_SERVER_ADDRESS"`
	AdminMail		string `mapstructure:"LM_ADMIN_MAIL"`
	AdminPassword	string `mapstructure:"LM_ADMIN_PASSWORD"`
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