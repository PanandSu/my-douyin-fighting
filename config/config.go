package config

type System struct {
	MysqlConfig *MysqlConfig `mapstructure:"mysql"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
	GinConfig   *GinConfig   `mapstructure:"gin"`
	JWTConfig   *JWTConfig   `mapstructure:"jwt"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type GinConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"signing_key"`
}
