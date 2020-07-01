package Handle

import (
	"net"
	"sync"
)

//协程调用

type limitListener struct {
	net.Listener
	sem       chan struct{}
	closeOnce sync.Once //确保已完成的chan只关闭一次
	done      chan struct{}
}

// n:最多同时接受n个监听器
// l:来自侦听器的连接。
func LimitListener(l net.Listener, n int) net.Listener {
	return &limitListener{
		Listener: l,
		sem:      make(chan struct{}, n),
		done:     make(chan struct{}),
	}
}
// 读取限制的信号量成功返回true，关闭且未返回则为false
func (l *limitListener) acquire() bool {
	select {
	case <-l.done:
		return false
	case l.sem <- struct{}{}:
		return true
	}
}
func (l *limitListener) release() { <-l.sem }

func (l *limitListener) Accept() (net.Conn, error) {
	acquired := l.acquire()
	//如果由于侦听器关闭而没有获得信号量，调用不会阻塞，但会立即返回一个错误。
	c, err := l.Listener.Accept()
	if err != nil {
		if acquired {
			l.release()
		}
		return nil, err
	}
	return &limitListenerConn{Conn: c, release: l.release}, nil
}

func (l *limitListener) Close() error {
	err := l.Listener.Close()
	l.closeOnce.Do(func() { close(l.done) })
	return err
}

type limitListenerConn struct {
	net.Conn
	releaseOnce sync.Once
	release     func()
}

func (l *limitListenerConn) Close() error {
	err := l.Conn.Close()
	l.releaseOnce.Do(l.release)
	return err
}
