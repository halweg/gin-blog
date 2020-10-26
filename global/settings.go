package global

import (
    "github.com/halweg/gin-blog/pkg/logger"
    "github.com/halweg/gin-blog/pkg/settings"
)

var (
    ServerSettings *settings.ServerSettings
    AppSettings *settings.AppSettings
    DatabaseSettings *settings.DataBaseSettings
    Logger *logger.Logger
)
