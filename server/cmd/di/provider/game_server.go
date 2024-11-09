package provider

import "github.com/yamato0204/cache-server-golang/internal/handler"

type GameServer struct {
	Router *handler.Router
}

func NewGameServer(router *handler.Router) *GameServer {
	return &GameServer{Router: router}
}
