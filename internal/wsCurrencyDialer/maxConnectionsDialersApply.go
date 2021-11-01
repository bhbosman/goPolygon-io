package wsCurrencyDialer

type maxConnectionsDialersApply struct {
	maxConnections int
}

func MaxConnections(maxConnections int) *maxConnectionsDialersApply {
	return &maxConnectionsDialersApply{maxConnections: maxConnections}
}

func (self maxConnectionsDialersApply) apply(settings *DialerSettings) {
	settings.maxConnections = self.maxConnections
}
