package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type App struct {
	Log             log
	Image           image
	RunMode         string
	Secret          string
	PageSize        int
	RuntimeRootPath string
}
type log struct {
	SavePath   string
	SaveName   string
	FileExt    string
	TimeFormat string
}
type image struct {
	PrefixUrl string
	SavePath  string
	MaxSize   int
	AllowExts []string
}
type Server struct {
	Port         int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}
type DataBase struct {
	Type        string
	User        string
	Pwd         string
	Url         string
	Name        string
	TablePrefix string
}

var AppSetting App
var ServerSetting Server
var DBSetting DataBase

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	LoadApp(&AppSetting)
	LoadDB(&DBSetting)
	LoadServer(&ServerSetting)
}

func LoadDB(d *DataBase) {
	d.Type = viper.GetString("database.type")
	d.User = viper.GetString("database.user")
	d.Pwd = viper.GetString("database.pwd")
	d.Url = viper.GetString("database.url")
	d.Name = viper.GetString("database.name")
	d.TablePrefix = viper.GetString("database.tablePrefix")
}

func LoadServer(server *Server) {
	server.Port = viper.GetInt("server.port")
	server.ReadTimeOut = viper.GetDuration("server.readTimeOut") * time.Second
	server.WriteTimeOut = viper.GetDuration("server.writeTimeOut") * time.Second
}

func LoadApp(app *App) {
	app.RunMode = viper.GetString("app.runMode")
	app.Secret = viper.GetString("app.Secret")
	app.PageSize = viper.GetInt("app.pageSize")
	app.RuntimeRootPath = viper.GetString("app.runtimeRootPath")
	app.Log.SavePath = viper.GetString("app.log.savePath")
	app.Log.SaveName = viper.GetString("app.log.saveName")
	app.Log.FileExt = viper.GetString("app.log.fileExt")
	app.Log.TimeFormat = viper.GetString("app.log.timeFormat")
	app.Image.PrefixUrl = viper.GetString("app.image.prefixUrl")
	app.Image.SavePath = viper.GetString("app.image.savePath")
	app.Image.MaxSize = viper.GetInt("app.image.maxSize")* 1024 * 1024
	app.Image.AllowExts = viper.GetStringSlice("app.image.allowExts")
}
