package global

import (
	"time"
)

var (
	Conf = new(config)
)

type config struct {
	Server   *serverConfig
	App      *appConfig
	Database *dataBaseConfig
	Logs     *logsConfig
	Redis    *redisConf
}

type serverConfig struct {
	RunModel      string
	HttpPort      string
	ReadTimeOut   time.Duration
	WriteTimeOut  time.Duration
	UrlPathPrefix string
	RsaPublicKey  string
	RsaPrivateKey string
}

type appConfig struct {
	DefaultPageSize int
	MaxPageSize     int
}

type dataBaseConfig struct {
	DBType       string
	UserName     string
	PassWord     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type logsConfig struct {
	Path string
}

type redisConf struct {
	Host         string
	PassWord     string
	DB           int
	PoolSize     int
	MaxRetries   int
	MinIdleConns int
}
