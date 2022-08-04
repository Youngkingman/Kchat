package config

import (
	"time"

	"github.com/spf13/viper"
)

type GlobalConfig struct {
	vp *viper.Viper
}

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	// DBType       string
	// UserName     string
	// Password     string
	// Host         string
	// DBName       string
	// TablePrefix  string
	// Charset      string
	// ParseTime    bool
	// MaxIdleConns int
	// MaxOpenConns int

}

type SecuritySettingS struct {
}

func (s *GlobalConfig) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func NewGlobalConfig() (*GlobalConfig, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &GlobalConfig{vp}, nil
}
