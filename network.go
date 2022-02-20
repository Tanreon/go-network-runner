package network_runner

import (
	"github.com/nadoo/glider/proxy"
	"github.com/nadoo/glider/proxy/http"
	"github.com/nadoo/glider/proxy/socks5"
	"github.com/nadoo/glider/proxy/ss"
	"github.com/nadoo/glider/proxy/ssh"
	"github.com/nadoo/glider/rule"
)

type DialOptions struct {
	DialTimeout  int
	RelayTimeout int
	Strategy     string
}

func NewDirectDialer(options ...DialOptions) (dialer *rule.Proxy, err error) {
	dialTimeout := 60
	relayTimeout := 60

	if len(options) > 0 {
		dialTimeout = options[0].DialTimeout
		relayTimeout = options[0].RelayTimeout
	}

	newProxy := rule.NewProxy(nil, &rule.Strategy{
		DialTimeout:  dialTimeout,  // sec
		RelayTimeout: relayTimeout, // sec
	}, nil)
	newProxy.Check()

	return newProxy, nil
}

func NewProxyDialers(forwards []string, options ...DialOptions) (dialer *rule.Proxy, err error) {
	dialTimeout := 10
	strategy := "ha"

	if len(options) > 0 {
		dialTimeout = options[0].DialTimeout
		strategy = options[0].Strategy
	}

	newProxy := rule.NewProxy(forwards, &rule.Strategy{
		DialTimeout: dialTimeout, // sec
		Strategy:    strategy,    // rr, ha, lha, dh
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
