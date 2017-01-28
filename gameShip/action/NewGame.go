package action

import (
	"golang.org/x/net/context"

	uuid "github.com/satori/go.uuid"
	"github.com/scriptnull/bingo/gameShip/game"
	"github.com/scriptnull/bingo/gameShip/gameShipRpc"
	"github.com/uber-go/zap"
)

// NewGame creates a Game object and saves in a global store
func (s *GameShipRPCServer) NewGame(ctx context.Context, in *gameShipRpc.NewGameRequest) (*gameShipRpc.NewGameResponse, error) {
	s.Logger.Info("Start NewGame", zap.Object("NewGameRequest", in))

	var creator = game.NewPlayer(in.CreatorId)

	var newGame = &game.Game{
		ID:      uuid.NewV4().String(),
		Players: []*game.Player{creator},
		Creator: creator,
	}

	if err := s.GameStore.Add(newGame); err != nil {
		return nil, err
	}

	var result = &gameShipRpc.NewGameResponse{
		GameId: newGame.ID,
	}

	s.Logger.Info("End NewGame", zap.Object("NewGameResponse", result))
	return result, nil
}
