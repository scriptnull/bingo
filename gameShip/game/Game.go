package game

import (
	"errors"
	"time"
)

// Game instance
type Game struct {
	ID                 string
	Creator            *Player
	Players            []*Player
	GameLoop           *time.Timer
	currentPlayerIndex int
	CurrentPlayer      *Player
	Results            []*Result
	Mode               Mode
}

// Start the game
func (g *Game) Start() {
	g.currentPlayerIndex = -1
	g.startNextTurn()
}

// startNextTurn will give the turn to next user on line
func (g *Game) startNextTurn() {
	var playerIndex = g.currentPlayerIndex + 1
	if playerIndex == len(g.Players) {
		playerIndex = 0
	}

	g.currentPlayerIndex = playerIndex
	g.CurrentPlayer = g.Players[playerIndex]

	if g.Mode == TimerMode {
		g.GameLoop = time.NewTimer(1 * time.Second)
		<-g.GameLoop.C
		// write default behavior when the time out happens
		// may be strike a random box in name of the currentPlayer
		// or just give the chance to next user
		g.startNextTurn()
	}
}

// PlayerStrike used to strike a box in the player's board
func (g *Game) PlayerStrike(p *Player, rowIndex, columnIndex int32) error {
	if p != g.CurrentPlayer {
		return errors.New("Could not strike, as its not the turn of the player")
	}

	// stop the current game loop timer
	if g.Mode == TimerMode {
		var isTurnTimerStopped = g.GameLoop.Stop()
		if isTurnTimerStopped {
			// the turn timeout expired and so, throw an error saying that turn timedout
			return errors.New("Could not strike, as the turn timed out already")
		}
	}

	// strike the number in player's board
	err := p.Strike(rowIndex, columnIndex)

	if err != nil {
		return err
	}

	// start the next turn
	g.startNextTurn()
	return nil
}

// PlayerBingo executes for a player, when he tries to press bingo in frontend
func (g *Game) PlayerBingo(p *Player) (*Result, error) {
	err := p.Bingo()

	if err != nil {
		return nil, err
	}

	var result = &Result{
		Position: int32(len(g.Results)),
		Player:   p,
	}

	g.Results = append(g.Results, result)

	return result, nil
}

// GetPlayerByID gets player in a game by ID
func (g *Game) GetPlayerByID(playerID string) (*Player, error) {
	var reqPlayer *Player

	for _, p := range g.Players {
		if p.ID == playerID {
			reqPlayer = p
			break
		}
	}

	if reqPlayer == nil {
		return nil, errors.New("PlayerNotFound")
	}
	return reqPlayer, nil
}
