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

// NeedRegisterEmailVerification 是否需要注册邮箱验证
var NeedRegisterEmailVerification = false

// GetConfig 获取配置
func GetConfig(key string) any {
	filter := map[string]string{"k": key}
	var result Config
	err := Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Errorf("find config error: %s", err)
		return nil
	}
	return result.Value
}

// LoadConfig 加载配置
func LoadConfig() {
	NeedRegisterEmailVerification = GetConfig("NeedRegisterEmailVerification").(bool)

	log.Debugf("NeedRegisterEmailVerification: %t", NeedRegisterEmailVerification)
}
