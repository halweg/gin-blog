package app

import (
    "github.com/gin-gonic/gin"
    "github.com/halweg/gin-blog/global"
    "github.com/halweg/gin-blog/pkg/convert"
)

func GetPage(c *gin.Context) int {
    page := convert.StrTo(c.Query("page")).MustInt()
    if page > 0 {
        return page
    }
    return 1
}

func GetPageSize(c *gin.Context) int {

    pageSize := convert.StrTo(c.Query("pageSize")).MustInt()

    if pageSize > 0 {
        return global.AppSettings.DefaultPageSize
    }

    if pageSize > global.AppSettings.DefaultPageSize {
        return global.AppSettings.MaxPageSize
    }

    return pageSize
}

func GetPageOffset(page, pageSize int) int {
    res := 0
    if page > 0 {
        res = (page-1) * pageSize
    }

    return res
}