package common

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal(fn func(sig os.Signal) bool) {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGABRT, syscall.SIGPIPE, syscall.SIGTERM) // 注册信号


	for sig := range ch {
		if !fn(sig) {
			close(ch)
			break
		}
	}
}