package conf

import (
	"sync"
)

// ConnectionStats 结构体用于保存多个连接信息
type ConnectionStats struct {
	Id         int `gorm:"primaryKey;autoIncrement"`
	LocalPort  string
	RemoteAddr string
	RemotePort string
	Protocol   string
	TotalBytes uint64
}

// 全局转发协程等待组
var Wg sync.WaitGroup

// 全局协程通道 未初始化默认为nil
var Ch chan string