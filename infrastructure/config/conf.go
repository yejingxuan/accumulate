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
	General string
	Log     Log
	Server  Server
	File    File
	Nacos   Nacos
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
	Port                   int
	GrpcPort               int
	TaskRunnerServer       string
	MaxHTTPTime            int64
	DB                     db
	MachineCode            int64
	MaxCoroutineNum        int
	UploadBasePath         string
	FileStorePath          string
	HostOpFilePath         string
	ClearPackageCron       string
	ClearPackageCronEnable bool
	Env                    map[string]EnvConfig
}

type EnvConfig struct {
	Value string
}

type db struct {
	DriveName    string
	Dsn          string
	MaxIdle      int
	MaxConn      int
	MaxQueryTime Duration
	LogMode      bool
}

type File struct {
	OpFilePath         string
	OpAdjunctPath      string
	SvrFilePath        string
	ImageFilePath      string
	DockerFileTempPath string
	ImageUrl           string
	OpImageProject     string
	ImageUserName      string
	ImagePassWord      string
	ImagePrefix        string
	SvrPrefix          string
}

type Nacos struct {
	IpAddr      string
	Port        uint64
	ContextPath string
	NamespaceId string
	DataId      string
	Group       string
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
