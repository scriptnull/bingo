package action

import (
	"golang.org/x/net/context"

	"github.com/scriptnull/bingo/gameShip/gameShipRpc"
	"github.com/uber-go/zap"
)

// PlayerStrikeBox creates a Game object and saves in a global store
func (s *GameShipRPCServer) PlayerStrikeBox(ctx context.Context, in *gameShipRpc.PlayerStrikeBoxRequest) (*gameShipRpc.PlayerStrikeBoxResponse, error) {
	var funcName = zap.String("FuncName", "actions|PlayerStrikeBox|")
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "Start"))
	s.Logger.Debug("funcVariable", funcName, zap.Object("PlayerStrikeBoxRequest", in))

	reqGame, err := s.GameStore.GetByGameID(in.GameId)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	reqPlayer, err := reqGame.GetPlayerByID(in.PlayerId)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	err = reqGame.PlayerStrike(reqPlayer, in.Row, in.Column)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	var result = &gameShipRpc.PlayerStrikeBoxResponse{
		GameId:   reqGame.ID,
		PlayerId: reqPlayer.ID,
	}

	s.Logger.Debug("funcVariable", funcName, zap.Object("PlayerStrikeBoxResponse", result))
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "End"))
	return result, nil
}
