package action

import (
	"golang.org/x/net/context"

	"github.com/scriptnull/bingo/gameShip/gameShipRpc"
	"github.com/uber-go/zap"
)

// PlayerStrikeBox creates a Game object and saves in a global store
func (s *GameShipRPCServer) PlayerStrikeBox(ctx context.Context, in *gameShipRpc.PlayerStrikeBoxRequest) (*gameShipRpc.PlayerStrikeBoxResponse, error) {
	s.Logger.Info("Start PlayerStrikeBox", zap.Object("PlayerStrikeBoxRequest", in))

	reqGame, err := s.GameStore.GetByGameID(in.GameId)
	if err != nil {
		return nil, err
	}

	reqPlayer, err := reqGame.GetPlayerByID(in.PlayerId)
	if err != nil {
		return nil, err
	}

	s.Logger.Debug("mylog before", zap.Object("reqGame.CurrentPlayer.ID", reqGame.CurrentPlayer.ID))

	err = reqGame.PlayerStrike(reqPlayer, in.Row, in.Column)
	if err != nil {
		return nil, err
	}

	var result = &gameShipRpc.PlayerStrikeBoxResponse{
		GameId:   reqGame.ID,
		PlayerId: reqPlayer.ID,
	}
	// s.Logger.Debug("mylog", zap.Object("player", reqPlayer))
	s.Logger.Debug("mylog after", zap.Object("reqGame.CurrentPlayer.ID", reqGame.CurrentPlayer.ID))
	s.Logger.Info("End NewGame", zap.Object("PlayerStrikeBoxResponse", result))
	return result, nil
}
