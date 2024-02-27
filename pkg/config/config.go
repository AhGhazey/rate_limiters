package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

// go get github.com/spf13/viper
const (
	DefaultAssetsPath = "/etc/config/rate_limiter"
)

type Config struct {
	server *serverConfig
	redis  *redisConfig
	topics *redisTopics
}

type serverConfig struct {
	host         string
	port         int
	readTimeOut  time.Duration
	writeTimeOut time.Duration
	idleTimeOut  time.Duration
	waitTimeOut  time.Duration
}

type redisConfig struct {
	host     string
	port     int
	password string
	db       int
}

type redisTopics struct {
	rateLimiterTopic string
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.server.host, c.server.port)
}

func (c *Config) ReadTimeOut() time.Duration {
	return c.server.readTimeOut
}

func (c *Config) WriteTimeOut() time.Duration {
	return c.server.writeTimeOut
}

func (c *Config) IdleTimeOut() time.Duration {
	return c.server.idleTimeOut
}

func (c *Config) WaitTimeOut() time.Duration {
	return c.server.waitTimeOut
}

func (c *Config) RedisHost() string {
	return c.redis.host
}

func (c *Config) RedisPort() int {
	return c.redis.port
}

func (c *Config) RedisPassword() string {
	return c.redis.password
}

func (c *Config) RedisDb() int {
	return c.redis.db
}

func (c *Config) RateLimiterTopic() string {
	return c.topics.rateLimiterTopic
}

func LoadConfig() (config *Config, err error) {
	path := DefaultAssetsPath
	if assetsPath := os.Getenv("ASSETS_PATH"); assetsPath != "" {
		path = assetsPath
	}
	v := newDefaultConfig(path)
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config = newConfig(v)
	return config, nil
}

func newDefaultConfig(path string) (v *viper.Viper) {
	v = viper.New()
	v.SetConfigName("env")
	v.SetConfigType("yml")
	v.AddConfigPath(path)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(
		strings.NewReplacer(".", "_", "-", "_"), // replace the (.) and (-) with (_)
	)
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.readTimeOut", "10s")
	v.SetDefault("server.writeTimeOut", "10s")
	v.SetDefault("server.idleTimeOut", "10s")
	v.SetDefault("server.waitTimeOut", "10s")
	v.SetDefault("redis.host", "localhost")
	v.SetDefault("redis.port", 6379)
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)
	v.SetDefault("topics.rateLimiterTopic", "requests")

	return v
}

func newConfig(v *viper.Viper) (config *Config) {
	config = &Config{
		server: &serverConfig{
			host:         v.GetString("server.host"),
			port:         v.GetInt("server.port"),
			readTimeOut:  v.GetDuration("server.readTimeOut"),
			writeTimeOut: v.GetDuration("server.writeTimeOut"),
			idleTimeOut:  v.GetDuration("server.idleTimeOut"),
			waitTimeOut:  v.GetDuration("server.waitTimeOut"),
		},
		redis: &redisConfig{
			host:     v.GetString("redis.host"),
			port:     v.GetInt("redis.port"),
			password: v.GetString("redis.password"),
			db:       v.GetInt("redis.db"),
		},
		topics: &redisTopics{
			rateLimiterTopic: v.GetString("topics.rateLimiterTopic"),
		},
	}
	return config
}
