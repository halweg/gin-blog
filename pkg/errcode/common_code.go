package errcode

var (
    Success             = NewError(0, "成功")
    ServerError         = NewError(1000000, "服务器内部错误")
    InvalidParams       = NewError(10000001, "入参错误")
    NotFound            = NewError(10000002, "入参错误")
    UnauthorizedAuthNotExist  = NewError(1000003, "鉴权失败，找不到对应的AppID 或 AppSecret")
    UnauthorizedTokenError    = NewError(1000004, "鉴权失败，Token无效")
    UnauthorizedTokenTimeout  = NewError(1000005, "鉴权失败，Token超时")
    UnauthorizedTokenGenerate = NewError(1000006, "Token生成失败")
    TooManyRequest            = NewError(1000007, "请求太多")
)
