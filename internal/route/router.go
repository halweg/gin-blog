package route

import (
    "github.com/gin-gonic/gin"
    _ "github.com/halweg/gin-blog/docs"
    v1 "github.com/halweg/gin-blog/internal/route/v1"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    article := v1.NewArticle()
    tag := v1.NewTag()

    apiv1 := r.Group("api/v1")
    {
        apiv1.POST("/tags", tag.Create)
        apiv1.DELETE("/tags/:id", tag.Delete)
        apiv1.PUT("/tags/:id", tag.Update)
        apiv1.PATCH("/tags/:id/state", tag.Update)
        apiv1.GET("/tags", tag.List)

        apiv1.POST("/articles", article.Create)
        apiv1.DELETE("/articles/:id", article.Delete)
        apiv1.PUT("/articles/:id", article.Update)
        apiv1.PATCH("/articles/:id/state", article.Get)
        apiv1.GET("/articles", article.List)

    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}