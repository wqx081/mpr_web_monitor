package monitor

import (
	"../common"
	"../config"
	"fmt"
)

type Monitor interface {
	Run(command <-chan common.MonitorCommand)
}

type SimpleMonitor struct {
	config config.Config
}

func NewSimpleMonitor(config config.Config) *SimpleMonitor {
	return &SimpleMonitor{
		config: config,
	}
}

func (sm *SimpleMonitor) Run(command <-chan common.MonitorCommand) {

	fmt.Printf("ping_interval: %d, ping_timeout: %d\n", sm.config.GetPingInterval(), sm.config.GetPingTimeout())

	for _, v := range sm.config.GetServers() {
		// How to Exist goroutine:
		// http://stackoverflow.com/questions/6807590/how-to-stop-a-goroutine
		//
		// go func() {
		//   si := &common.ServerInfo{ Name: v["name"], Url: v["url"]}
		//   hc := NewHealthChecker(si)
		//   common.AddServerStat(v["name"])
		//   hc.Check(si)
		// }
		fmt.Printf("checking Name: %s, Url: %s\n", v["name"], v["url"])
	}

	for {
		cm := <-command
		switch cm {
		case common.Stop:
			fmt.Printf("Stop monitor")
			return
		case common.Reconfig:
			fmt.Printf("Reconfig monitor")
			return
		default:
			fmt.Printf("Unknown command: %v", cm)
			break
		}
	}
}
