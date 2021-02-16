// Exercise 2.10: Looping Over a Map

package main

import (
	"fmt"
)

func main() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "info",
		"version":  "1.0.1",
	}
	for key, value := range config {
		fmt.Println(key, "=", value)
	}

	zabbix_agent_config := map[string]string{
		"LogFile":     "/var/log/zabbix-agent/zabbix_agentd.log",
		"LogFileSize": "1GB",
		"Server":      "127.0.0.1",
	}
	for k, v := range zabbix_agent_config {
		fmt.Println(k, "=", v)
	}
}
