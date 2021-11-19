package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"time"
)

var (
	// CoreConf conf
	CoreConf CoreConfig
)

// Init Config
func Init(conf string) {
	_, err := toml.DecodeFile(conf, &CoreConf)
	if err != nil {
		fmt.Printf("Err %v", err)
		os.Exit(1)
	}
	log.Println("read config file ==>", conf)
	log.Println("----- coreConf ----- \n", CoreConf)
}

// Init Config
func InitByContent(content string) {
	_, err := toml.Decode(content, &CoreConf)
	if err != nil {
		fmt.Printf("Err %v", err)
		os.Exit(1)
	}
	log.Println("-----remote coreConf ----- \n", CoreConf)
}

type CoreConfig struct {
	Log    Log
	Server Server
}

// Log Config
type Log struct {
	LogPath    string
	MaxSize    int
	Compress   bool
	MaxAge     int
	MaxBackups int
	LogLevel   string
	Format     string
}

// Server server config
type Server struct {
	Port            int
	MaxHTTPTime     int64
	MaxCoroutineNum int
	MachineCode     int64

	DB db
}

type db struct {
	DriveName    string
	Dsn          string
	MaxIdle      int
	MaxConn      int
	MaxQueryTime Duration
	LogMode      bool
}

type Duration struct {
	time.Duration
}

// UnmarshalText parse 10s to time.Time
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
