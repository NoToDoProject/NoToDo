package common

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"os"
	"reflect"
)

// IsFileExists 判断文件是否存在
func IsFileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if fileInfo != nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// EncryptPassword 加密密码
func EncryptPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panicf("encrypt password error: %s", err)
	}
	return hashedPassword
}

// ComparePassword 比较密码
func ComparePassword(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}

// CopyStruct 复制结构体
func CopyStruct(src, dst interface{}) {
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcValue.Kind() != reflect.Ptr || dstValue.Kind() != reflect.Ptr {
		return
	}
	if srcValue.IsNil() || dstValue.IsNil() {
		return
	}

	srcElem := srcValue.Elem()
	dstElem := dstValue.Elem()
	for i := 0; i < dstElem.NumField(); i++ {
		name := dstElem.Type().Field(i).Name
		if srcElem.FieldByName(name).IsValid() {
			if srcElem.FieldByName(name).Type() == dstElem.Field(i).Type() {
				dstElem.Field(i).Set(srcElem.FieldByName(name))
			}
		}
	}
}
