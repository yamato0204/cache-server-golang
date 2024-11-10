package provider

import (
	"github.com/yamato0204/cache-server-golang/internal/handler"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
)

type GameServer struct {
	Router *handler.Router
	Db     *mysql.ApplicationDB
}

func NewGameServer(router *handler.Router, db *mysql.ApplicationDB) *GameServer {
	return &GameServer{Router: router, Db: db}
}
