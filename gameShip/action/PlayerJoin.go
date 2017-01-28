package action

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/scriptnull/bingo/gameShip/game"
	"github.com/scriptnull/bingo/gameShip/gameShipRpc"
	"github.com/uber-go/zap"
)

// PlayerJoin allows player to join a game
func (s *GameShipRPCServer) PlayerJoin(ctx context.Context, in *gameShipRpc.PlayerJoinRequest) (*gameShipRpc.PlayerJoinResponse, error) {
	s.Logger.Info("Start PlayerJoin", zap.Object("PlayerJoinRequest", in))

	var newPlayer = game.NewPlayer(in.PlayerId)

	// game presence check
	reqGame, err := s.GameStore.GetByGameID(in.GameId)
	if err != nil {
		return nil, err
	}

	// player absence check
	var alreadyJoined bool
	for _, p := range reqGame.Players {
		if p.ID == in.PlayerId {
			alreadyJoined = true
			break
		}
	}
	if alreadyJoined {
		return nil, errors.New("PlayerAlreadyJoined")
	}

	reqGame.Players = append(reqGame.Players, newPlayer)

	var result = &gameShipRpc.PlayerJoinResponse{
		GameId: reqGame.ID,
	}

	s.Logger.Info("End PlayerJoin", zap.Object("PlayerJoinResponse", result))
	return result, nil
}
