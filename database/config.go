package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

// ConfigInDatabase config item in database
type ConfigInDatabase struct {
	Key   string `bson:"k"`
	Value any    `bson:"v"`
}

// ConfigCollection mongo collection
var ConfigCollection *mongo.Collection

// Config config struct
var Config = struct {
	NeedRegisterEmailVerification bool          // need register email verification when register
	EmailVerificationTimeOut      time.Duration // time out of email verification
	CanRegister                   bool          // enable register
}{
	// default values
	false,
	time.Minute * 30,
	true,
}

// GetConfig get config from database
func GetConfig(key string) any {
	filter := map[string]string{"k": key}
	var result ConfigInDatabase
	r := ConfigCollection.FindOne(context.Background(), filter)
	if r.Err() == mongo.ErrNoDocuments {
		// set default value when config not found, use reflect to get default value
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

// LoadConfig use reflect to load config
func LoadConfig() {
	configValue := reflect.ValueOf(&Config)
	configElem := configValue.Elem()
	for i := 0; i < configElem.NumField(); i++ {
		field := configElem.Field(i)              // field of config
		name := configElem.Type().Field(i).Name   // name of config
		typeOfField := configElem.Type().Field(i) // go type of field
		kind := field.Kind()                      // raw type of field

		newValue := GetConfig(name)                       // config value
		kindOfNewValue := reflect.TypeOf(newValue).Kind() // raw type of config
		if kind != kindOfNewValue {
			log.Errorf("config %s type error, need %s, got %s", name, kind, kindOfNewValue)
			continue
		}
		log.Debugf("config %s: %v", name, newValue)

		// transform type
		if typeOfField.Type == reflect.TypeOf(time.Minute) {
			newValue = time.Duration(newValue.(int64))
		}
		field.Set(reflect.ValueOf(newValue))
	}
}

// UpdateConfig set new config
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
