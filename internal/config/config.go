package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zggsong/message/internal/global"
	"github.com/zggsong/message/internal/model"
	"os"
)

// InitConfig
//
//	@Description: 初始化配置
//	@return model.Config
//	@return error
func InitConfig() (model.Config, error) {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("../../../config")
	viper.AddConfigPath(workDir + "/config")
	viper.OnConfigChange(func(e fsnotify.Event) {
		global.GLO_CONF_CH <- updateConfig()
	})
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return model.Config{}, err
	}
	return updateConfig(), nil
}

// updateConfig
//
//	@Description: 更新配置
//	@return model.Config
func updateConfig() model.Config {
	var config model.Config
	config.MsgEnabled = viper.GetBool("message.enabled")
	config.MsgType = viper.GetString("message.type")
	config.BarkUrl = viper.GetString("message.bark.url")
	config.BarkKey = viper.GetString("message.bark.key")
	config.MailHost = viper.GetString("message.mail.host")
	config.MailProtocol = viper.GetString("message.mail.protocol")
	config.MailPort = viper.GetInt("message.mail.port")
	config.MailUser = viper.GetString("message.mail.username")
	config.MailPwd = viper.GetString("message.mail.password")
	config.MailFromName = viper.GetString("message.mail.from_name")
	config.MailTo = viper.GetStringSlice("message.mail.to")
	return config
}
