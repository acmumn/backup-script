package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// LoadConfig loads the config file from the given path and returns the configuration information,
// panicking on error.
func LoadConfig(path string) (MysqlConfig, AwsConfig, map[Frequency]ScheduleConfig) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		log.Panicln(err)
	}

	schedule, err := config.Schedule()
	if err != nil {
		log.Panicln(err)
	}

	return config.Mysql, config.Aws, schedule
}

// Config is the global configuration for backups.
type Config struct {
	Mysql MysqlConfig `toml:"mysql"`
	Aws   AwsConfig   `toml:"aws"`

	ScheduleInner map[string]ScheduleConfig `toml:"schedule"`
}

func (c Config) Schedule() (map[Frequency]ScheduleConfig, error) {
	schedule := make(map[Frequency]ScheduleConfig, len(c.ScheduleInner))
	for k, v := range c.ScheduleInner {
		var k2 Frequency
		if err := k2.UnmarshalText([]byte(k)); err != nil {
			return nil, err
		}
		schedule[k2] = v
	}
	return schedule, nil
}

// MysqlConfig is the MySQL-specific configuration.
type MysqlConfig struct {
	Host string `toml:"host"`
	Port uint16 `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

// AwsConfig is the AWS-specific configuration.
type AwsConfig struct {
	SecretKey string `toml:"secret_key"`
	S3Bucket  string `toml:"s3_bucket"`
}

// ScheduleConfig is the configuration for a single item in the schedule.
type ScheduleConfig struct {
	Incremental bool `toml:"incremental"`
}

type Frequency struct {
	code int
}

func (f *Frequency) UnmarshalText(text []byte) error {
	textStr := string(text)
	if textStr == "daily" {
		f.code = 0
	} else if textStr == "weekly" {
		f.code = 1
	} else if textStr == "monthly" {
		f.code = 2
	} else if textStr == "yearly" {
		f.code = 3
	} else {
		return fmt.Errorf("Unknown frequency: %s", text)
	}
	return nil
}
