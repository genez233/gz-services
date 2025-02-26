package sett

import (
	"github.com/spf13/viper"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("./configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

type ServerSetting struct {
	ServerName   string
	Version      string
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize   int
	LogSavePath       string
	LogFileName       string
	LogFileExt        string
	UploadSavePath    string
	UploadFileMaxSize int
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	Port         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
