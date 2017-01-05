package common

import (
	"sync"
)

// MonitorCommand
type MonitorCommand int

const (
	Stop MonitorCommand = iota
	Reconfig
)

// ServerStatus
type ServerStatus int

const (
	StatusRunning     ServerStatus = iota // 服务正常运行
	StatusStopped                         // 服务停止, 但是物理服务器还运行
	StatusUnreachable                     // 物理服务器停止
	StatusInit
)

func (ss ServerStatus) String() string {
	switch ss {
	case StatusRunning:
		return "Running"
	case StatusStopped:
		return "Stopped"
	case StatusUnreachable:
		return "Unreachable"
	case StatusInit:
		return "Init"
	default:
		return "Unknown"
	}
}

//
type ServerInfo struct {
	Name   string
	Url    string
	Status ServerStatus
}

type ServerStatistics map[string]*ServerInfo

var mu sync.Mutex
var server_statistics = make(ServerStatistics)

func AddServerStat(name string, url string) {
	mu.Lock()
	defer mu.Unlock()
	si := ServerInfo{Name: name, Url: url, Status: StatusInit}
	server_statistics[name] = &si
}

func UpdateServerStat(name string, status ServerStatus) {
	mu.Lock()
	defer mu.Unlock()
	server_statistics[name].Status = status
}

// 不需加锁, 返回的结果只用作读取
func GetServerStatistics() ServerStatistics {
	return server_statistics
}
