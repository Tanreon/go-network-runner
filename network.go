package network_runner

import (
	"errors"

	"github.com/nadoo/glider/proxy"
	"github.com/nadoo/glider/proxy/http"
	"github.com/nadoo/glider/proxy/socks5"
	"github.com/nadoo/glider/proxy/ss"
	"github.com/nadoo/glider/proxy/ssh"
	"github.com/nadoo/glider/rule"
)

var ErrIncorrectStrategyMethod = errors.New("INCORRECT_STRATEGY_TYPE")

type DialOptions struct {
	DialTimeout  int
	RelayTimeout int
	Strategy     string
}

func NewDirectDialer(options ...IDirectDialOptions) (dialer *rule.Proxy, err error) {
	ruleStrategy := rule.Strategy{}

	if len(options) > 0 {
		directDialOptions := options[0]

		if directDialOptions.IsDialTimeoutSet() {
			ruleStrategy.DialTimeout = directDialOptions.DialTimeout()
		} else {
			ruleStrategy.DialTimeout = 30
		}
		if directDialOptions.IsRelayTimeoutSet() {
			ruleStrategy.RelayTimeout = directDialOptions.RelayTimeout()
		} else {
			ruleStrategy.RelayTimeout = 60
		}
	}

	newProxy := rule.NewProxy(nil, &ruleStrategy, nil)
	newProxy.Check()

	return newProxy, nil
}

func NewProxyDialers(forwards []string, options ...IProxyDialOptions) (dialer *rule.Proxy, err error) {
	ruleStrategy := rule.Strategy{}

	if len(options) > 0 {
		proxyDialOptions := options[0]

		if proxyDialOptions.IsStrategySet() {
			ruleStrategy.Strategy = proxyDialOptions.Strategy()
		} else {
			ruleStrategy.Strategy = "rr"
		}
		if proxyDialOptions.IsCheckUrlSet() {
			ruleStrategy.Check = proxyDialOptions.CheckUrl()
		}
		if proxyDialOptions.IsCheckIntervalSet() {
			ruleStrategy.CheckInterval = proxyDialOptions.CheckInterval()
		}
		if proxyDialOptions.IsCheckTimeoutSet() {
			ruleStrategy.CheckTimeout = proxyDialOptions.CheckTimeout()
		}
		if proxyDialOptions.IsCheckToleranceSet() {
			ruleStrategy.CheckTolerance = proxyDialOptions.CheckTolerance()
		}
		if proxyDialOptions.IsCheckDisabledOnlySet() {
			ruleStrategy.CheckDisabledOnly = proxyDialOptions.CheckDisabledOnly()
		}
		if proxyDialOptions.IsCheckMaxFailuresSet() {
			ruleStrategy.MaxFailures = proxyDialOptions.MaxFailures()
		}
		if proxyDialOptions.IsDialTimeoutSet() {
			ruleStrategy.DialTimeout = proxyDialOptions.DialTimeout()
		} else {
			ruleStrategy.DialTimeout = 10
		}
		if proxyDialOptions.IsRelayTimeoutSet() {
			ruleStrategy.RelayTimeout = proxyDialOptions.RelayTimeout()
		}
	}

	newProxy := rule.NewProxy(forwards, &ruleStrategy, nil)
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
