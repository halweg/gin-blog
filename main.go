package main

import (
    "github.com/gin-gonic/gin"
    "github.com/halweg/gin-blog/global"
    "github.com/halweg/gin-blog/internal/model"
    "github.com/halweg/gin-blog/internal/route"
    "github.com/halweg/gin-blog/pkg/logger"
    "github.com/halweg/gin-blog/pkg/settings"
    "gopkg.in/natefinch/lumberjack.v2"
    "log"
    "net/http"
    "time"
)

func init() {
    err := setupSettings()
    if err != nil {
        log.Fatalf("init.setupSetting err : %v", err)
    }

    err = setupDBEngine()

    if err != nil {
        log.Fatalf("init.setupDBEngine err : %v", err)
    }

    err = setupLogger()

    if err != nil {
        log.Fatalf("init.setupLogger err %v", err)
    }

}

func main() {
    gin.SetMode(global.ServerSettings.RunMode)
    router := route.NewRouter()
    global.Logger.Infof("%s:gin-blog is start! /%s", "halweg", "blog-service")

    s := &http.Server{
        Addr: ":" + global.ServerSettings.HttpPort,
        Handler: router,
        ReadHeaderTimeout: global.ServerSettings.ReadTimeout,
        WriteTimeout: global.ServerSettings.ReadTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()


}

func setupSettings() error{
    settings, err := settings.NewSetting()
    if err != nil {
        return err
    }

    err = settings.ReadSection("Server", &global.ServerSettings)
    if err != nil {
        return err
    }

    err = settings.ReadSection("App", &global.AppSettings)
    if err != nil {
        return err
    }

    err = settings.ReadSection("Database", &global.DatabaseSettings)
    if err != nil {
        return err
    }

    global.ServerSettings.ReadTimeout *= time.Second
    global.ServerSettings.WriteTimeout *= time.Second

    return nil

}

func setupDBEngine()  error {
    var err error

    global.DBEngine, err = model.NewDBEngine(global.DatabaseSettings)

    if err != nil {
        return err
    }

    return nil
}

func setupLogger() error {
    global.Logger = logger.NewLogger(&lumberjack.Logger{
        Filename:   global.AppSettings.LogSavePath  + "/" +global.AppSettings.LogFileName + "/" + global.AppSettings.LogFileExt ,
        MaxSize:    600,
        MaxAge:     10,
        LocalTime:  true,
    }, "", log.LstdFlags).WithCaller(2)
    return nil
}
