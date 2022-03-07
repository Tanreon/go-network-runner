package network_runner

type IDialOptions interface {
	IsDialTimeoutSet() bool
	SetDialTimeout(timeout int)
	DialTimeout() int

	IsRelayTimeoutSet() bool
	SetRelayTimeout(timeout int)
	RelayTimeout() int
}

type IDirectDialOptions struct {
	IDialOptions
}

type DirectDialOptions struct {
	dialTimeout  *int
	relayTimeout *int
}

func (d *DirectDialOptions) IsDialTimeoutSet() bool {
	return d.dialTimeout != nil
}
func (d *DirectDialOptions) SetDialTimeout(timeout int) {
	d.dialTimeout = &timeout
}
func (d *DirectDialOptions) DialTimeout() int {
	return *d.dialTimeout
}

func (d *DirectDialOptions) IsRelayTimeoutSet() bool {
	return d.relayTimeout != nil
}
func (d *DirectDialOptions) SetRelayTimeout(timeout int) {
	d.relayTimeout = &timeout
}
func (d *DirectDialOptions) RelayTimeout() int {
	return *d.relayTimeout
}

type IProxyDialOptions interface {
	IDialOptions

	IsStrategySet() bool
	SetStrategy(strategy string)
	Strategy() string // rr, ha, lha, dh

	IsCheckUrlSet() bool
	SetCheckUrl(url string)
	CheckUrl() string

	IsCheckIntervalSet() bool
	SetCheckInterval(interval int)
	CheckInterval() int

	IsCheckTimeoutSet() bool
	SetCheckTimeout(timeout int)
	CheckTimeout() int

	IsCheckToleranceSet() bool
	SetCheckTolerance(tolerance int)
	CheckTolerance() int

	IsCheckDisabledOnlySet() bool
	SetCheckDisabledOnly(value bool)
	CheckDisabledOnly() bool

	IsCheckMaxFailuresSet() bool
	SetCheckMaxFailures(value int)
	MaxFailures() int
}

type ProxyDialOptions struct {
	dialTimeout       *int
	relayTimeout      *int
	strategy          *string
	checkUrl          *string
	checkInterval     *int
	checkTimeout      *int
	checkTolerance    *int
	checkDisabledOnly *bool
	maxFailures       *int
}

func (p *ProxyDialOptions) IsDialTimeoutSet() bool {
	return p.dialTimeout != nil
}
func (p *ProxyDialOptions) SetDialTimeout(timeout int) {
	p.dialTimeout = &timeout
}
func (p *ProxyDialOptions) DialTimeout() int {
	return *p.dialTimeout
}

func (p *ProxyDialOptions) IsRelayTimeoutSet() bool {
	return p.relayTimeout != nil
}
func (p *ProxyDialOptions) SetRelayTimeout(timeout int) {
	p.relayTimeout = &timeout
}
func (p *ProxyDialOptions) RelayTimeout() int {
	return *p.relayTimeout
}

func (p *ProxyDialOptions) IsStrategySet() bool {
	return p.strategy != nil
}
func (p *ProxyDialOptions) SetStrategy(strategy string) error {
	switch strategy {
	case "rr":
		fallthrough
	case "ha":
		fallthrough
	case "lha":
		fallthrough
	case "dh":
		p.strategy = &strategy
	default:
		return ErrIncorrectStrategyMethod
	}

	return nil
}
func (p *ProxyDialOptions) Strategy() string {
	return *p.strategy
}

func (p *ProxyDialOptions) IsCheckUrlSet() bool {
	return p.checkUrl != nil
}

// SetCheckUrl http://www.msftconnecttest.com/connecttest.txt#expect=200
func (p *ProxyDialOptions) SetCheckUrl(url string) {
	p.checkUrl = &url
}
func (p *ProxyDialOptions) CheckUrl() string {
	return *p.checkUrl
}

func (p *ProxyDialOptions) IsCheckIntervalSet() bool {
	return p.checkInterval != nil
}
func (p *ProxyDialOptions) SetCheckInterval(interval int) {
	p.checkInterval = &interval
}
func (p *ProxyDialOptions) CheckInterval() int {
	return *p.checkInterval
}

func (p *ProxyDialOptions) IsCheckTimeoutSet() bool {
	return p.checkTimeout != nil
}
func (p *ProxyDialOptions) SetCheckTimeout(timeout int) {
	p.checkTimeout = &timeout
}
func (p *ProxyDialOptions) CheckTimeout() int {
	return *p.checkTimeout
}

func (p *ProxyDialOptions) IsCheckToleranceSet() bool {
	return p.checkTolerance != nil
}
func (p *ProxyDialOptions) SetCheckTolerance(tolerance int) {
	p.checkTolerance = &tolerance
}
func (p *ProxyDialOptions) CheckTolerance() int {
	return *p.checkTolerance
}

func (p *ProxyDialOptions) IsCheckDisabledOnlySet() bool {
	return p.checkDisabledOnly != nil
}
func (p *ProxyDialOptions) SetCheckDisabledOnly(value bool) {
	p.checkDisabledOnly = &value
}
func (p *ProxyDialOptions) CheckDisabledOnly() bool {
	return *p.checkDisabledOnly
}

func (p *ProxyDialOptions) IsCheckMaxFailuresSet() bool {
	return p.maxFailures != nil
}
func (p *ProxyDialOptions) SetCheckMaxFailures(value int) {
	p.maxFailures = &value
}
func (p *ProxyDialOptions) MaxFailures() int {
	return *p.maxFailures
}
