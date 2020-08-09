package settings

import "github.com/spf13/viper"

type Settings struct {
    vp *viper.Viper
}

func NewSetting() (*Settings, error) {
    vp := viper.New()
    vp.SetConfigName("config")
    vp.AddConfigPath("configs/")
    vp.SetConfigType("yaml")

    err := vp.ReadInConfig()

    if err != nil {
        return nil, err
    }

    return &Settings{vp}, nil

}
