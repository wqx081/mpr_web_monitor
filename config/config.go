package config

import (
	//	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	//	"log"
	//"time"
)

// YAML:
//
// ping_interval: 1000ms
// ping_timeout: 2s
// retry_count: 5
// servers:
//	- name: 'book server'
//	  url: 'http://book_server.com/api/'
//  - name: 'cp portal'
//    url: 'http://cp_portal.com/api/'
//
// JSON:
// {
//		"ping_interval": "1000ms",
//		"ping_timeout": "2s",
//		"retry_count": 5,
//		"servers": [
//         {
//             "name": "book_server",
//			   "url": "http://book_server/api"
//		    },
//		    {
//	           "name": "cp portal",
//			   "url": "http://cp_portal/api"
//			}
//	     ]
// }

type ServerDescriptorList []map[string]string

type Config interface {
	GetPingInterval() uint64
	GetPingTimeout() uint64
	GetRetryCount() uint64
	GetServers() ServerDescriptorList
}

type SimpleConfig struct {
	PingInterval uint64
	PingTimeout  uint64
	RetryCount   uint64
	Servers      ServerDescriptorList
}

// Interface
func (sc *SimpleConfig) GetPingInterval() uint64 {
	return sc.PingInterval
}

func (sc *SimpleConfig) GetPingTimeout() uint64 {
	return sc.PingTimeout
}

func (sc *SimpleConfig) GetRetryCount() uint64 {
	return sc.RetryCount
}

func (sc *SimpleConfig) GetServers() ServerDescriptorList {
	return sc.Servers
}

type yamlConfig struct {
	PingInterval uint64               `json:"ping_interval"`
	PingTimeout  uint64               `json:"ping_timeout"`
	RetryCount   uint64               `json:"retry_count"`
	Servers      ServerDescriptorList `json:"servers"` // server_name: url
}

func NewConfigFromFile(fpath string) (*SimpleConfig, error) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewConfigFromBytes(b)
}

func NewConfigFromBytes(b []byte) (*SimpleConfig, error) {

	yc := &yamlConfig{}

	err := yaml.Unmarshal(b, yc)
	if err != nil {
		return nil, err
	}

	cfg := &SimpleConfig{
		PingInterval: yc.PingInterval,
		PingTimeout:  yc.PingTimeout,
		RetryCount:   yc.RetryCount,
		Servers:      yc.Servers,
	}

	return cfg, nil
}

//func main() {
//	var FilePath string = "./config.yaml"
//
//	cfg, err := NewConfigFromFile(FilePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("ping_interval: %d, ping_timeout: %d, retry_count: %d\n", cfg.GetPingInterval(), cfg.GetPingTimeout(), cfg.GetRetryCount())
//	for _, v := range cfg.GetServers() {
//		fmt.Printf("name: %s, url: %s\n", v["name"], v["url"])
//	}
//}
