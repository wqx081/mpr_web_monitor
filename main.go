package main

import (
	"./common"
	"./config"
	"./monitor"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Chain use to control Monitor
var command = make(chan common.MonitorCommand)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Bad request")
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1>", "Hello, World")
}

func main() {
	configFilePath := os.Args[1]
	cfg, err := config.NewConfigFromFile(configFilePath)

	if err != nil {
		log.Fatalf("Parse config file:%s error:%v", configFilePath, err)
	}

	log.Println("Start Monitor...")
	go func() {
		monitor := monitor.NewSimpleMonitor(cfg)
		monitor.Run(command)
	}()

	log.Println("Start Http API...")
	http.HandleFunc("/", HealthCheckHandler)
	http.ListenAndServe(":8080", nil)
}

/*
Start(config *Config) {
	pingInterval := config.GetPingInterval()
	pingTimeout := config.GetPingTimeout()
	retryCount := config.GetRetryCount()
	servers := config.GetServers()
}

enum ServerStatus {
	RUNNING,
	DEAD,
}

struct ServerInfo {
	Name string
	Url string

}

struct ServerMonitor {
	sync.Mutex

	Servers []ServerInfo
}


func main() {
	go func() {
		cfg, err := NewConfigFromFile("config.yaml")
        monitor, err := NewMonitor(cfg)
		monitor.Run()
	}

	http.HandleFunc("/", HealthCheckHandler)
	http.ListenAndServe(":8080", nil)
}
*/
