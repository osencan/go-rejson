package rejson

import (
	redigo "github.com/gomodule/redigo/redis"
	"github.com/osencan/go-rejson/clients"
	"github.com/osencan/go-rejson/rjs"
	goredis "github.com/osencan/redis"
)

// RedisClient provides interface for Client handling in the ReJSON Handler
type RedisClient interface {
	SetClientInactive()
	SetRedigoClient(redigo.Conn)
	SetGoRedisClient(conn *goredis.Client)
}

// SetClientInactive resets the handler and unset any client, set to the handler
func (r *Handler) SetClientInactive() {
	_t := &Handler{clientName: rjs.ClientInactive}
	r.clientName = _t.clientName
	r.implementation = _t.implementation
}

// SetRedigoClient sets Redigo (https://github.com/gomodule/redigo/redis) client
// to the handler
func (r *Handler) SetRedigoClient(conn redigo.Conn) {
	r.clientName = "redigo"
	r.implementation = &clients.Redigo{Conn: conn}
}

// SetGoRedisClient sets Go-Redis (https://github.com/osencan/redis) client to
// the handler
func (r *Handler) SetGoRedisClient(conn *goredis.Client) {
	r.clientName = "goredis"
	r.implementation = &clients.GoRedis{Conn: conn}
}
