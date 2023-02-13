package config

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Config 配置项
type Config struct {
	Key   string `bson:"k"`
	Value any    `bson:"v"`
}

// Collection mongo集合
var Collection *mongo.Collection

var (
	NeedRegisterEmailVerification = false // 是否需要注册邮箱验证
)

// GetConfig 获取配置
func GetConfig(key string) any {
	filter := map[string]string{"k": key}
	var result Config
	err := Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Errorf("find config error: %s", err)
		return err
	}
	return result.Value
}

// SetConfig 设置配置
func SetConfig(key string, value any) {
	filter := map[string]string{"k": key}
	update := map[string]any{"$set": map[string]any{"v": value}}
	_, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Errorf("update config error: %s", err)
	}
}

// LoadConfig 加载配置
func LoadConfig() {
	NeedRegisterEmailVerification = GetConfig("NeedRegisterEmailVerification").(bool)

	log.Debugf("NeedRegisterEmailVerification: %t", NeedRegisterEmailVerification)
}

// SetDefaultConfigIfNotExist 设置默认配置
func SetDefaultConfigIfNotExist() {
	// 是否需要注册邮箱验证
	if GetConfig("NeedRegisterEmailVerification") == mongo.ErrNoDocuments {
		SetConfig("NeedRegisterEmailVerification", false)
	}
}
