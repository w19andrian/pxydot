package app

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"

	"github.com/w19andrian/pxydot/config"
)

type Upstream struct {
	Addr string
}

type PoolUpstream struct {
	Pool []*Upstream
	Last uint64
}

// InitUpstreamServers unloads the Upstream_Servers from
// the configuration into the Ustream structs
func InitUpstreamServers(c *config.Config) *PoolUpstream {
	p := &PoolUpstream{}

	for i := range c.Upstream_Servers {
		us := &Upstream{
			Addr: c.Upstream_Servers[i],
		}
		p.Pool = append(p.Pool, us)
	}
	return p
}

// NextIndex will increment the index of []*Upstream
func (p *PoolUpstream) NextIndex() int {
	return int(atomic.AddUint64(&p.Last, uint64(1)) % uint64(len(p.Pool)))
}

// GetNextUpstream is an implementation of Roundrobin load balancing
// will return the []*Upstream of the index n of the *PoolStream and
// will reset the index back to the beginning after it reached the final index.
func (p *PoolUpstream) GetNextUpstream() *Upstream {

	n := p.NextIndex()
	l := len(p.Pool) + n

	for i := n; i < l; i++ {
		x := i % len(p.Pool)
		if p.Pool[x].IsAlive() {
			if i != n {
				atomic.StoreUint64(&p.Last, uint64(x))
			}
			return p.Pool[i]
		}
	}
	return nil
}

// IsAlive checks if the upstream is available for connection on tcp/853
func (u Upstream) IsAlive() bool {
	if _, err := net.Dial("tcp", fmt.Sprintf("%s:853", u.Addr)); err != nil {
		log.Printf("Upstream server %v is not available\n", u.Addr)
		return false
	}
	return true
}
