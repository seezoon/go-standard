package base

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Grpc *GrpcConf
	Log  *LogConf
	Db   *DBConf
}

type GrpcConf struct {
	Network string
	Listen  string
}

type LogConf struct {
	Filename string
	// M
	MaxSize int
	// Day
	MaxAge          int
	Level           string
	ReportCaller    bool
	OutputToConsole bool
}

type DBConf struct {
	Dsn          string
	MaxIdleConns int
	// S
	ConnMaxLifetime time.Duration
	// S
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
}

var (
	ConfInstance Config
	once         sync.Once
)

const (
	configPath = "conf"
	configName = "seezoon"
	configType = "yml"
	// 指定的配置路径优先
	ENV_CONF_PATH_KEY = "CONF_PATH"
)

func init() {
	once.Do(func() {
		viper := viper.New()
		confPath := os.Getenv(ENV_CONF_PATH_KEY)
		if confPath != "" {
			viper.AddConfigPath(confPath)
		} else {
			workDir, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			viper.AddConfigPath(workDir + string(os.PathSeparator) + configPath)
		}
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		// 环境变量大写加_ ,可以自动取环境变量
		viper.AutomaticEnv()
		replacer := strings.NewReplacer(".", "_")
		viper.SetEnvKeyReplacer(replacer)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&ConfInstance); err != nil {
			panic(err)
		}
		if marshal, err := json.Marshal(&ConfInstance); err != nil {
			panic(err)
		} else {
			log.Debugf("config is : %s", marshal)
		}
	})
}
