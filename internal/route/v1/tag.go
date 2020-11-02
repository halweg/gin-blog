package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/halweg/gin-blog/global"
    _ "github.com/halweg/gin-blog/global"
    _ "github.com/halweg/gin-blog/internal/model"
    "github.com/halweg/gin-blog/internal/service"
    _ "github.com/halweg/gin-blog/internal/service"
    "github.com/halweg/gin-blog/pkg/app"
    _ "github.com/halweg/gin-blog/pkg/app"
    "github.com/halweg/gin-blog/pkg/errcode"
    _ "github.com/halweg/gin-blog/pkg/errcode"
    _ "github.com/halweg/gin-blog/pkg/errcode"
    "net/http"
)


type Tag struct {

}

func NewTag() Tag {
    return Tag{}
}

// @Summary 获取单个标签
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {object} model.Tag "成功
// @Failure 400 {object} errcode.ErrorSwagger "请求错误"
// @Failure 500 {object} errcode.ErrorSwagger "内部错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context)  {
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxLength(100)
// @Param state query int false "标签状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功
// @Failure 400 {object} errcode.ErrorSwagger "请求错误"
// @Failure 500 {object} errcode.ErrorSwagger "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {

    var data service.TagListRequest
    ret := app.NewResponse(c)
    if valida, e := app.BindAndValid(c, &data); valida {
        global.Logger.Debugf("app.BindAndValid errs:%v", e)
        ret.ToResponseError(errcode.InvalidParams.WithDetails(e.Errors()...))
        return
    } else {
        c.JSON(http.StatusOK, gin.H{
            "data": data,
        })
        return
    }
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" maxLength(100) minlength(3)
// @Param state body int false "标签状态" Enums(0, 1) default(1)
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {object} errcode.ErrorSwagger "请求错误"
// @Failure 500 {object} errcode.ErrorSwagger "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context)  {

}

// @Summary 修改标签
// @Produce json
// @Param id path int true "标签id"
// @Param name body string false "标签名称" maxlength(100) minlength(3)
// @Param state body int false "标签状态" Enums(0, 1) default(1)
// @Param modify_by body string true "修改人/修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功
// @Failure 400 {object} errcode.ErrorSwagger "请求错误"
// @Failure 500 {object} errcode.ErrorSwagger "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update (c *gin.Context) {

}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签id"
// @Param deleted_by body string true "修改人/修改者" minlength(3) maxlength(100)
// @Success 200 {object} string "成功
// @Failure 400 {object} errcode.ErrorSwagger "请求错误"
// @Failure 500 {object} errcode.ErrorSwagger "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete (c *gin.Context) {

}
