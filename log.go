package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为Text格式
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 时间格式
		ForceColors:     true,                  // 强制颜色
		//FullTimestamp:   true,                  // 显示完整时间
	})

	// 设置将日志输出到标准输出（默认的输出为stderr）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
}
