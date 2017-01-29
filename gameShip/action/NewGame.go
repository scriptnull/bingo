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
	var funcName = zap.String("FuncName", "actions|NewGame|")
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "Start"))
	s.Logger.Debug("funcVariable", funcName, zap.Object("NewGameRequest", in))

	var creator = game.NewPlayer(in.CreatorId)

	var newGame = &game.Game{
		ID:      uuid.NewV4().String(),
		Players: []*game.Player{creator},
		Creator: creator,
	}

	if err := s.GameStore.Add(newGame); err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	var result = &gameShipRpc.NewGameResponse{
		GameId: newGame.ID,
	}

	s.Logger.Debug("funcVariable", funcName, zap.Object("NewGameResponse", result))
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "End"))
	return result, nil
}
