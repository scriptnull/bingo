package action

import (
	"github.com/scriptnull/bingo/gameShip/store"
	"github.com/uber-go/zap"
)

// GameShipRPCServer implments the gameShipRpc interface
type GameShipRPCServer struct {
	Logger    zap.Logger
	GameStore store.GameStore
}
