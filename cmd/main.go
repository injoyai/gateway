package main

import "github.com/injoyai/gateway/internal/boot"

func main() {
	boot.Init()
	select {}
}
