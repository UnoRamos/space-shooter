package messages

import (
	"space-shooter/game/component"
	"space-shooter/game/types"
)

type PlayerData struct {
	PlayerId   types.PlayerId
	PlayerName string
	Position   component.PositionData
}

type ConnectionHandshake struct {
	PlayerName string
}

type ConnectionHandshakeResponse struct {
	IsRoomFull bool

	PlayerId   types.PlayerId
	PlayerData []PlayerData
}

type UpdatePosition struct {
	PlayerId types.PlayerId
	Position component.PositionData
}

// Message sent from the client to the server to tell the
// server that the client detected a move by the player.
type RegisterPlayerMove struct {
	Move types.PlayerMove

	// We send the position to see if it matches how the server moved
	// the player.
	Position component.PositionData
}

// Message sent from the server to the clients to render the
// player move.
type EventPlayerMove struct {
	Move     types.PlayerMove
	PlayerId types.PlayerId
}

// Message sent from the server to the clients to render the
// following position of the new player.
type EventPlayerConnected struct {
	PlayerId   types.PlayerId
	PlayerName string
	Position   component.PositionData
}

// Message sent from the server to the clients to tell the clients that the
// corresponding PlayerId has disconnected
type EventPlayerDisconnected struct {
	PlayerId types.PlayerId
}

type EventPlayerFireBullet struct {
	PlayerId types.PlayerId
}

type EventUpdateHealth struct {
	PlayerId types.PlayerId // The player whose health is being updated
	Health   float64        // The updated health value of the player
}

type EventPlayerDied struct {
	PlayerId types.PlayerId // The player whose health is being updated
	KilledBy types.PlayerId
}

type EventPlayerRespawned struct {
	PlayerId types.PlayerId
	Position component.PositionData
}
