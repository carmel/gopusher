package comet

import (
	"gopusher/api"
	"gopusher/configuration"
	"gopusher/connection/websocket"
	"gopusher/contracts"
)

func Run() {
	config := configuration.GetCometConfig()

	server := getCometServer(config)

	go server.Run()

	go server.JoinCluster()

	api.InitRpcServer(server, config)
}

func getCometServer(config *configuration.CometConfig) contracts.Server {
	switch config.SocketProtocol {
	case "ws":
		fallthrough
	case "wss":
		return websocket.NewWebSocketServer(config)
	case "tcp": //暂时不处理
		panic("Unsupported protocol: " + config.SocketProtocol)
	default:
		panic("Unsupported protocol: " + config.SocketProtocol)
	}
}
