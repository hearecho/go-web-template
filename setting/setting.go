package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type App struct {
	Log             log
	File            file
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
type file struct {
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

var AppSetting = App{
	Log: log{
		SavePath:   "logs/",
		SaveName:   "log",
		FileExt:    "log",
		TimeFormat: "20060102",
	},
	File: file{
		PrefixUrl: "",
		SavePath:  "upload/images/",
		MaxSize:   0,
		AllowExts: nil,
	},
	RunMode:         "debug",
	Secret:          "123456",
	PageSize:        5,
	RuntimeRootPath: "runtime/",
}
var ServerSetting = Server{
	Port:         8080,
	ReadTimeOut:  60 * time.Second,
	WriteTimeOut: 60 * time.Second,
}
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
	// 需要判断是否存在该条属性 如果不存在则使用默认值
	if !viper.IsSet("database") {
		return
	}
	d.User = viper.GetString("database.user")
	d.Pwd = viper.GetString("database.pwd")
	d.Url = viper.GetString("database.url")
	d.Name = viper.GetString("database.name")
	d.TablePrefix = viper.GetString("database.tablePrefix")
}

func LoadServer(server *Server) {
	if !viper.IsSet("server") {
		return
	}
	server.Port = viper.GetInt("server.port")
	server.ReadTimeOut = viper.GetDuration("server.readTimeOut") * time.Second
	server.WriteTimeOut = viper.GetDuration("server.writeTimeOut") * time.Second
}

func LoadApp(app *App) {
	if !viper.IsSet("app") {
		return
	}
	app.RunMode = viper.GetString("app.runMode")
	app.Secret = viper.GetString("app.Secret")
	app.PageSize = viper.GetInt("app.pageSize")
	app.RuntimeRootPath = viper.GetString("app.runtimeRootPath")
	app.Log.SavePath = viper.GetString("app.log.savePath")
	app.Log.SaveName = viper.GetString("app.log.saveName")
	app.Log.FileExt = viper.GetString("app.log.fileExt")
	app.Log.TimeFormat = viper.GetString("app.log.timeFormat")
	app.File.PrefixUrl = viper.GetString("app.file.prefixUrl")
	app.File.SavePath = viper.GetString("app.file.savePath")
	app.File.MaxSize = viper.GetInt("app.file.maxSize") * 1024 * 1024 * 500
	app.File.AllowExts = viper.GetStringSlice("app.file.allowExts")
}
