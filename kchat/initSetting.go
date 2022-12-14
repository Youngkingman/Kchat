package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/service"
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
	"github.com/Youngkingman/Kchat/kchat/pkg/logger"
	"github.com/Youngkingman/Kchat/kchat/pkg/setting"
	"github.com/Youngkingman/Kchat/kchat/pkg/validator"
	"github.com/gin-gonic/gin/binding"
	"github.com/natefinch/lumberjack"
)

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("ChatRoom", &global.ChatRoomSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Qiniu", &global.QiniuSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.MySQL, err = dbutil.NewSQL(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.Redis, err = dbutil.NewRedisCli(global.RedisSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator

	return nil
}

func InitSetting() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
	}
	// err = setupTracer()
	// if err != nil {
	// 	log.Fatalf("init.setupTracer err: %v", err)
	// }
	// 服务器启动的时候加载所有房间
	err = service.LoadAllChatRoom()
	if err != nil {
		log.Fatalf("init.LoadChatRoom err: %v", err)
	}
}
