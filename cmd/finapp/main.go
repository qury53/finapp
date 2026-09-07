package main

import (
	"context"
	"os/signal"
	"syscall"
)

func main() {
	signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	
}