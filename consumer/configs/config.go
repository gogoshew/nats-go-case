package configs

import "github.com/spf13/viper"

type Config struct {
	Db struct {
		Host     string
		User     string
		Password string
		Dbname   string
		Port     string
		Sslmode  string
	}
}

func InitConfig() (string, error) {
	viper.SetConfigType("yml")
	viper.AddConfigPath("./consumer/configs/")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return "", err
	}

	dsn := "host=" + config.Db.Host + " user=" + config.Db.User + " password=" + config.Db.Password + " dbname=" + config.Db.Dbname +
		" port=" + config.Db.Port + " sslmode=" + config.Db.Sslmode

	return dsn, nil
}
