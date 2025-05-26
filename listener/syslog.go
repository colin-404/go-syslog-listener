package listener

import (
	"log"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func ListenSyslog(protocol string, listenPort string) {
	// 初始化Elasticsearch客户端
	// 设置syslog服务器
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)

	if protocol == "udp" {
		server.ListenUDP("0.0.0.0:" + listenPort)
	} else if protocol == "tcp" {
		server.ListenTCP("0.0.0.0:" + listenPort)
	}

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			// log.Println(logParts)

			// 解析syslog消息
			if logParts["message"] != nil {
				message := logParts["message"].(string)

				if xKey, exists := logParts["x_key"]; exists && xKey == "dlp_block" {
					continue
				}
				log.Println(message)

			}

		}
	}(channel)

	server.Wait()
}
