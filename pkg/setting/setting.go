package setting

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type App struct{
	RunMode string
	PageSize int
	JwtSecret string
}

var AppS *App

type Server struct {
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerS  *Server

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}
var DatabaseS *Database

type Setting struct {
	viper *viper.Viper
}

func NewSetting() (*Setting,error) {
	vp:=viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("conf/")
	err:=vp.ReadInConfig()
	if err!=nil {
		return nil ,err
	}
	return &Setting{viper: vp}, err
}

func (s *Setting) Section(k string ,v interface{}) error {
	err:=s.viper.UnmarshalKey(k,&v)
	if err != nil {
		return err
	}
	return nil
}

func SetUpSetting()  {
	s,err:=NewSetting()
	if err!=nil{
		log.Fatalf("read config failed: %v", err)
	}
	err=s.Section("Server",&ServerS)
	if err!=nil{
		log.Fatalf("read Server failed: %v", err)
	}
	ServerS.ReadTimeout*=time.Second
	ServerS.WriteTimeout*=time.Second
	err=s.Section("App",&AppS)
	if err!=nil{
		log.Fatalf("read Server failed: %v", err)
	}
	err=s.Section("Database",&DatabaseS)
	if err!=nil{
		log.Fatalf("read Server failed: %v", err)
	}

}