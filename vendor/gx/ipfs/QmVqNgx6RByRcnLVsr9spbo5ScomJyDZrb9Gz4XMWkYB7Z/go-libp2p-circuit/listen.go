package relay

import (
	"fmt"
	"net"

	pb "gx/ipfs/QmVqNgx6RByRcnLVsr9spbo5ScomJyDZrb9Gz4XMWkYB7Z/go-libp2p-circuit/pb"

	tpt "gx/ipfs/QmQGRkVSA5vTXt9WpJ6nBGnrvq9zRNsfjoNPq8uQrhnBoq/go-libp2p-transport"
	ma "gx/ipfs/QmW8s4zTsUoX1Q6CeYxVKPyqSKbF7H1YDUyTostBtZ8DaG/go-multiaddr"
	peer "gx/ipfs/QmWNY7dV54ZDYmTA1ykVdwNCqC11mpU4zSUp6XDpLTH9eG/go-libp2p-peer"
	filter "gx/ipfs/QmcrNvQBZkpbT7zLcGyjdPCiCHLVtaYeg4UVUH7XLfWJWy/go-maddr-filter"
)

var _ tpt.Listener = (*RelayListener)(nil)

type RelayListener Relay

func (l *RelayListener) Relay() *Relay {
	return (*Relay)(l)
}

func (r *Relay) Listener() *RelayListener {
	return (*RelayListener)(r)
}

func (l *RelayListener) Accept() (tpt.Conn, error) {
	select {
	case c := <-l.incoming:
		err := l.Relay().writeResponse(c.Stream, pb.CircuitRelay_SUCCESS)
		if err != nil {
			log.Debugf("error writing relay response: %s", err.Error())
			c.Stream.Reset()
			return nil, err
		}

		log.Infof("accepted relay connection: %s", c.ID())

		return c, nil
	case <-l.ctx.Done():
		return nil, l.ctx.Err()
	}
}

func (l *RelayListener) Addr() net.Addr {
	return &NetAddr{
		Relay:  "any",
		Remote: "any",
	}
}

func (l *RelayListener) Multiaddr() ma.Multiaddr {
	a, err := ma.NewMultiaddr(fmt.Sprintf("/p2p-circuit/ipfs/%s", l.self.Pretty()))
	if err != nil {
		panic(err)
	}
	return a
}

func (l *RelayListener) LocalPeer() peer.ID {
	return l.self
}

func (l *RelayListener) SetAddrFilters(f *filter.Filters) {
	// noop ?
}

func (l *RelayListener) Close() error {
	// TODO: noop?
	return nil
}
