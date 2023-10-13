package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type Hook interface {
	WithSignals(signals ...syscall.Signal) Hook //注册信号

	Close(funcs ...func()) //信号触发的动作
}

type hook struct {
	ctx chan os.Signal //接收信号的通道
}

// 初始化注册信号，以便在接收到信号时出发相应的处理
func NewHook() Hook {
	hook := &hook{
		ctx: make(chan os.Signal, 1),
	}
	//中断信号 和 终止信号 注册到 h.ctx中
	return hook.WithSignals(syscall.SIGINT, syscall.SIGTERM)
}

func (h *hook) WithSignals(signals ...syscall.Signal) Hook {
	for _, s := range signals {
		signal.Notify(h.ctx, s)
	}
	return h
}

// 关闭钩子，停止接收信号
func (h *hook) Close(funcs ...func()) {

	//阻塞在此，当接收到上述两种信号时才会往下继续执行
	select {
	case <-h.ctx:
	}

	signal.Stop(h.ctx)
	for _, f := range funcs {
		f()
	}
}
