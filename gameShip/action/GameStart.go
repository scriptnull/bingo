package action

import (
	"golang.org/x/net/context"

	"github.com/scriptnull/bingo/gameShip/gameShipRpc"
	"github.com/uber-go/zap"
)

// GameStart starts the game
func (s *GameShipRPCServer) GameStart(ctx context.Context, in *gameShipRpc.GameStartRequest) (*gameShipRpc.GameStartResponse, error) {
	s.Logger.Info("Start GameStart", zap.Object("GameStartRequest", in))

	// game presence check
	reqGame, err := s.GameStore.GetByGameID(in.GameId)
	if err != nil {
		return nil, err
	}

	reqGame.Start()

	s.Logger.Debug("CurrentPlayer", zap.Object("reqGame.CurrentPlayer", reqGame.CurrentPlayer))

	var result = &gameShipRpc.GameStartResponse{
		GameId:   reqGame.ID,
		PlayerId: reqGame.CurrentPlayer.ID,
	}
	s.Logger.Info("End GameStart", zap.Object("GameStartResponse", result))
	return result, nil
}
