package network

import (
	"github.com/nadoo/glider/proxy"
	"github.com/nadoo/glider/proxy/http"
	"github.com/nadoo/glider/proxy/socks5"
	"github.com/nadoo/glider/proxy/ss"
	"github.com/nadoo/glider/proxy/ssh"
	"github.com/nadoo/glider/rule"
)

//type DirectRunner struct {
//	client *resty.Client
//}

func NewDirectDialer() (dialer *rule.Proxy, err error) {
	newProxy := rule.NewProxy(nil, &rule.Strategy{
		DialTimeout:  60, // sec
		RelayTimeout: 60, // sec
	}, nil)
	newProxy.Check()

	return newProxy, nil
}

//type ProxyRunner struct {
//	client *resty.Client
//}

func NewProxyDialers(forwards []string, strategy string) (dialer *rule.Proxy, err error) {
	newProxy := rule.NewProxy(forwards, &rule.Strategy{
		DialTimeout: 10, // sec
		Strategy:    strategy,
	}, nil)
	newProxy.Check()

	return newProxy, nil
}

func RegisterSosks5Proxy() {
	proxy.RegisterDialer("socks5", socks5.NewSocks5Dialer)
}

func RegisterHTTPProxy() {
	proxy.RegisterDialer("http", http.NewHTTPDialer)
}

func RegisterSSHProxy() {
	proxy.RegisterDialer("ssh", ssh.NewSSHDialer)
}

func RegisterSSProxy() {
	proxy.RegisterDialer("ss", ss.NewSSDialer)
}
