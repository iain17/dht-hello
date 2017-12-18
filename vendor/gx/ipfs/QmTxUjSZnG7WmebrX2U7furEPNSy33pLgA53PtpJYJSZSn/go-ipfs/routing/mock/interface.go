// Package mock provides a virtual routing server. To use it, create a virtual
// routing server and use the Client() method to get a routing client
// (IpfsRouting). The server quacks like a DHT but is really a local in-memory
// hash table.
package mockrouting

import (
	"context"

	delay "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/thirdparty/delay"
	"gx/ipfs/QmeDA8gNhvRTsbrjEieay5wezupJDiky8xvCzDABbsGzmp/go-testutil"

	routing "gx/ipfs/QmPCGUjMRuBcPybZFpjhzpifwPP9wPRoiy5geTQKU4vqWA/go-libp2p-routing"
	peer "gx/ipfs/QmWNY7dV54ZDYmTA1ykVdwNCqC11mpU4zSUp6XDpLTH9eG/go-libp2p-peer"
	ds "gx/ipfs/QmdHG8MAuARdGHxx4rPQASLcvhz24fzjSQq7AJRAQEorq5/go-datastore"
)

// Server provides mockrouting Clients
type Server interface {
	Client(p testutil.Identity) Client
	ClientWithDatastore(context.Context, testutil.Identity, ds.Datastore) Client
}

// Client implements IpfsRouting
type Client interface {
	routing.IpfsRouting
}

// NewServer returns a mockrouting Server
func NewServer() Server {
	return NewServerWithDelay(DelayConfig{
		ValueVisibility: delay.Fixed(0),
		Query:           delay.Fixed(0),
	})
}

// NewServerWithDelay returns a mockrouting Server with a delay!
func NewServerWithDelay(conf DelayConfig) Server {
	return &s{
		providers: make(map[string]map[peer.ID]providerRecord),
		delayConf: conf,
	}
}

type DelayConfig struct {
	// ValueVisibility is the time it takes for a value to be visible in the network
	// FIXME there _must_ be a better term for this
	ValueVisibility delay.D

	// Query is the time it takes to receive a response from a routing query
	Query delay.D
}
