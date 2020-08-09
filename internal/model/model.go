package model

import (
    "fmt"
    "github.com/halweg/gin-blog/global"
    "github.com/halweg/gin-blog/pkg/settings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

)

type Model struct {
    ID uint32 `gorm:"primary_key" json:"id" `
    CreatedBy string `json:"created_by"`
    CreatedOn uint32 `json:"created_on"`
    ModifiedBy string `json:"modified_by"`
    ModifiedOn string `json:"modified_on"`
    DeletedOn uint32 `json:"deleted_on"`
    IsDel uint8 `json:"is_del"`
}

func NewDBEngine (dataBaseSettings *settings.DataBaseSettings)  (*gorm.DB, error) {

    db, err := gorm.Open(dataBaseSettings.DBType,
        fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
            dataBaseSettings.UserName,
            dataBaseSettings.PassWord,
            dataBaseSettings.Host,
            dataBaseSettings.DBName,
            dataBaseSettings.Charset,
            dataBaseSettings.ParseTime,
        ))

    if err != nil {
        return nil, err
    }

    if global.ServerSettings.RunMode == "debug" {
        db.LogMode(true)
    }

    db.SingularTable(true)
    db.DB().SetMaxIdleConns(dataBaseSettings.MaxIdleConns)
    db.DB().SetMaxOpenConns(dataBaseSettings.MaxOpenConns)

    return db, nil
}