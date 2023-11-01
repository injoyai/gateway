package common

import (
	"fmt"
	"github.com/injoyai/gateway/internal/data/listen"
	"github.com/injoyai/goutil/frame/gin"
	"github.com/injoyai/io"
	"github.com/injoyai/io/dial"
	"github.com/injoyai/logs"
)

var (
	HTTPServer   *gin.Server
	MQTTServer   *io.Server
	SerialClient *io.Client
	TCPServer    *io.Server
	UDPServer    *io.Server
)

func initListen() {
	initListenHTTP()
	initListenMQTT()
	initListenSerial()
	initListenTCP()
	initListenUDP()
}

func initListenHTTP() {
	var err error
	httpPort := Cfg.GetInt("listen.http.port")
	HTTPServer, err = listen.NewHTTP(httpPort)
	logs.PrintErr(err)
	if HTTPServer != nil {
		go HTTPServer.Run()
	}
}

func initListenMQTT() {
	var err error
	mqttPort := Cfg.GetInt("listen.mqtt.port")
	MQTTServer, err = listen.NewMQTT(mqttPort)
	logs.PrintErr(err)
	if MQTTServer != nil {
		go MQTTServer.Run()
	}
}

func initListenSerial() {
	var err error
	SerialClient, err = listen.NewSerial(&dial.SerialConfig{
		Address:  Cfg.GetString("listen.serial.address"),
		BaudRate: Cfg.GetInt("listen.serial.baudRate"),
		DataBits: Cfg.GetInt("listen.serial.dataBits"),
		StopBits: Cfg.GetInt("listen.serial.stopBits"),
		Parity:   Cfg.GetString("listen.serial.parity"),
		Timeout:  Cfg.GetMicrosecond("listen.serial.timeout"),
		RS485: dial.SerialRS485Config{
			Enabled:            Cfg.GetBool("listen.serial.rs485.enabled"),
			DelayRtsBeforeSend: Cfg.GetMillisecond("listen.serial.rs485.delayRtsBeforeSend"),
			DelayRtsAfterSend:  Cfg.GetMillisecond("listen.serial.rs485.delayRtsAfterSend"),
			RtsHighDuringSend:  Cfg.GetBool("listen.serial.rs485.rtsHighDuringSend"),
			RtsHighAfterSend:   Cfg.GetBool("listen.serial.rs485.rtsHighAfterSend"),
			RxDuringTx:         Cfg.GetBool("listen.serial.rs485.rxDuringTx"),
		},
	})
	logs.PrintErr(err)
	if SerialClient != nil {
		go SerialClient.Run()
	}
}

func initListenTCP() {
	var err error
	tcpPort := Cfg.GetInt("listen.tcp.port")
	TCPServer, err = listen.NewTCP(tcpPort, func(s *io.Server) {
		s.SetTimeout(Cfg.GetMillisecond("listen.tcp.timeout"))
	})
	logs.PrintErr(err)
	if TCPServer != nil {
		go TCPServer.Run()
	}
}

func initListenUDP() {
	var err error
	udpPort := Cfg.GetInt("listen.udp.port")
	UDPServer, err = listen.NewUDP(udpPort, func(s *io.Server) {
		s.SetKey(fmt.Sprintf(":%d", udpPort))
	})
	logs.PrintErr(err)
	if UDPServer != nil {
		go UDPServer.Run()
	}
}
