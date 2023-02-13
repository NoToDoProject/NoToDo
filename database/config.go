package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

// ConfigInDatabase 配置项
type ConfigInDatabase struct {
	Key   string `bson:"k"`
	Value any    `bson:"v"`
}

// ConfigCollection mongo集合
var ConfigCollection *mongo.Collection

// Config 配置
var Config = struct {
	NeedRegisterEmailVerification bool          // 是否需要注册邮箱验证
	EmailVerificationTimeOut      time.Duration // 邮箱验证超时时间
	CanRegister                   bool          // 是否允许注册
}{
	false,
	time.Minute * 30,
	true,
}

// GetConfig 获取配置
func GetConfig(key string) any {
	filter := map[string]string{"k": key}
	var result ConfigInDatabase
	r := ConfigCollection.FindOne(context.Background(), filter)
	if r.Err() == mongo.ErrNoDocuments {
		// 未找到时，设置默认值
		log.Errorf("config %s not found, set default value", key)
		defaultValue := reflect.ValueOf(Config).FieldByName(key).Interface()
		_, err := ConfigCollection.InsertOne(context.Background(), ConfigInDatabase{key, defaultValue})
		if err != nil {
			log.Errorf("insert config error: %s", err)
		}
		return defaultValue
	}
	_ = r.Decode(&result)
	return result.Value
}

// LoadConfig 加载配置
func LoadConfig() {
	configValue := reflect.ValueOf(&Config)
	configElem := configValue.Elem()
	for i := 0; i < configElem.NumField(); i++ {
		field := configElem.Field(i)              // 获取字段
		name := configElem.Type().Field(i).Name   // 字段名
		typeOfField := configElem.Type().Field(i) // 字段Go类型
		kind := field.Kind()                      // 字段原始类型

		newValue := GetConfig(name)                       // 获取配置
		kindOfNewValue := reflect.TypeOf(newValue).Kind() // 获取配置原始类型
		if kind != kindOfNewValue {
			log.Errorf("config %s type error, need %s, got %s", name, kind, kindOfNewValue)
			continue
		}
		log.Debugf("config %s: %v", name, newValue)

		// 如果是时间类型，需要转换
		if typeOfField.Type == reflect.TypeOf(time.Minute) {
			newValue = time.Duration(newValue.(int64))
		}
		field.Set(reflect.ValueOf(newValue))
	}

	log.Debugf("Config in Database: %+v", Config)
}

// UpdateConfig 更新配置
func UpdateConfig(key string, value any) {
	filter := map[string]string{"k": key}
	update := map[string]any{"$set": map[string]any{"v": value}}
	_, err := ConfigCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Errorf("update config error: %s", err)
	}
	// use reflect to set config
	configValue := reflect.ValueOf(&Config)
	configElem := configValue.Elem()
	configElem.FieldByName(key).Set(reflect.ValueOf(value))
}
