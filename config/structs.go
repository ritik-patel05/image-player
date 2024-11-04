package config

type Store struct {
	ActiveEnv string `mapstructure:"activeEnv" validate:"required"`
	Redis     Redis  `mapstructure:"redis"     validate:"required"`
}

type Redis struct {
	ImageRedis RedisConfig `mapstructure:"imageRedis" validate:"required"`
}

type RedisConfig struct {
	Host        []string `mapstructure:"host"        validate:"required"`
	IsClustered bool     `mapstructure:"isClustered" validate:"required"`
}
