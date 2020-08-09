package settings

import "time"

type ServerSettings struct {
    RunMode string
    HttpPort string
    ReadTimeout time.Duration
    WriteTimeout time.Duration
}

type AppSettings struct {
    DefaultPageSize int
    MaxPageSize int
    LogSavePath string
    LogFileName string
    LogFileExt string
}

type DataBaseSettings struct {
    DBType string
    UserName string
    PassWord string
    Host string
    DBName string
    TablePreFix string
    Charset string
    ParseTime bool
    MaxIdleConns int
    MaxOpenConns int
}

func (s *Settings) ReadSection(k string, v interface{}) error {
    err := s.vp.UnmarshalKey(k, v)
    if err != nil {
        return err
    }

    return nil
}
