package conf

import ( 
	"fmt"
	"sync"
	"time"

	"github.com/caarlos0/env"
)


var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		var err error
		if instance == nil {
			instance, err = getConfig() // not thread safe
			if err != nil {
				panic(err)
			}
		}
	})
	return instance
}

type Config struct {
	// server config
	ServerHost         string `env:"WEBOOK_SERVER_HOST" envDefault:"0.0.0.0"`
	ServerPort         string `env:"WEBOOK_SERVER_PORT" envDefault:"8001"`
	ServerMode         string `env:"WEBOOK_SERVER_MODE" envDefault:"debug"`
	SecretKey          string `env:"WEBOOK_SECRET_KEY" envDefault:"8xEMrWkBARcDDYQ"`
   
	// storage config
	MySQLAddr        string `env:"WEBOOK_MYSQL_ADDR" envDefault:"localhost"`
	MySQLPort        string `env:"WEBOOK_MYSQL_PORT" envDefault:"3306"`
	MySQLUser        string `env:"WEBOOK_MYSQL_USER" envDefault:"root"`
	MySQLPassword    string `env:"WEBOOK_MYSQL_PASSWORD" envDefault:"123456"`
	MySQLDatabase    string `env:"WEBOOK_MYSQL_DATABASE" envDefault:"webook_backend"`
	MySQLParameters  string `env:"WEBOOK_MYSQL_PARAMETERS" envDefault:"charset=utf8mb4&parseTime=true&loc=Local"`

	// cache config
	RedisAddr     string `env:"WEBOOK_REDIS_ADDR" envDefault:"localhost"`
	RedisPort     string `env:"WEBOOK_REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"WEBOOK_REDIS_PASSWORD" envDefault:"illa2022"`
	RedisDatabase int    `env:"WEBOOK_REDIS_DATABASE" envDefault:"0"`

	// oss config
	DriveType             string `env:"WEBOOK_DRIVE_TYPE" envDefault:""`
	DriveAccessKeyID      string `env:"ILLA_DRIVE_ACCESS_KEY_ID" envDefault:""`
	DriveAccessKeySecret  string `env:"ILLA_DRIVE_ACCESS_KEY_SECRET" envDefault:""`
	DriveRegion           string `env:"ILLA_DRIVE_REGION" envDefault:""`
	DriveEndpoint         string `env:"ILLA_DRIVE_ENDPOINT" envDefault:""`
	DriveSystemBucketName string `env:"ILLA_DRIVE_SYSTEM_BUCKET_NAME" envDefault:"illa-cloud"`
	DriveTeamBucketName   string `env:"ILLA_DRIVE_TEAM_BUCKET_NAME" envDefault:"illa-cloud-team"`
	DriveUploadTimeoutRaw string `env:"ILLA_DRIVE_UPLOAD_TIMEOUT" envDefault:"30s"`
	DriveUploadTimeout    time.Duration

	// google config
	IllaGoogleSheetsClientID     string `env:"ILLA_GS_CLIENT_ID" envDefault:""`
	IllaGoogleSheetsClientSecret string `env:"ILLA_GS_CLIENT_SECRET" envDefault:""`
	IllaGoogleSheetsRedirectURI  string `env:"ILLA_GS_REDIRECT_URI" envDefault:""`
}

func getConfig() (*Config, error) {
	// fetch
	cfg := &Config{}
	err := env.Parse(cfg)

	// process data
	var errInParseDuration error
	cfg.DriveUploadTimeout, errInParseDuration = time.ParseDuration(cfg.DriveUploadTimeoutRaw)
	if errInParseDuration != nil {
		return nil, errInParseDuration
	}
	// ok
	fmt.Printf("----------------\n")
	fmt.Printf("run by following config: %+v\n", cfg)
	fmt.Printf("parse config error info: %+v\n", err)

	return cfg, err
}


func (c *Config) GetSecretKey() string {
	return c.SecretKey
}

func (c *Config) GetMySQLAddr() string {
	return c.MySQLAddr
}

func (c *Config) GetMySQLPort() string {
	return c.MySQLPort
}

func (c *Config) GetMySQLUser() string {
	return c.MySQLUser
}

func (c *Config) GetMySQLPassword() string {
	return c.MySQLPassword
}

func (c *Config) GetMySQLDatabase() string {
	return c.MySQLDatabase
}

func (c *Config) GetMySQLParameters() string {
	return c.MySQLParameters
}



func (c *Config) GetRedisAddr() string {
	return c.RedisAddr
}

func (c *Config) GetRedisPort() string {
	return c.RedisPort
}

func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}

func (c *Config) GetRedisDatabase() int {
	return c.RedisDatabase
}

func (c *Config) GetDriveType() string {
	return c.DriveType
}


func (c *Config) GetAWSS3Endpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetAWSS3AccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetAWSS3AccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetAWSS3Region() string {
	return c.DriveRegion
}

func (c *Config) GetAWSS3SystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetAWSS3TeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetAWSS3Timeout() time.Duration {
	return c.DriveUploadTimeout
}

func (c *Config) GetMINIOAccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetMINIOAccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetMINIOEndpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetMINIOSystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetMINIOTeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetMINIOTimeout() time.Duration {
	return c.DriveUploadTimeout
}


func (c *Config) GetIllaGoogleSheetsClientID() string {
	return c.IllaGoogleSheetsClientID
}

func (c *Config) GetIllaGoogleSheetsClientSecret() string {
	return c.IllaGoogleSheetsClientSecret
}

func (c *Config) GetIllaGoogleSheetsRedirectURI() string {
	return c.IllaGoogleSheetsRedirectURI
}
