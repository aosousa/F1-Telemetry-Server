package models

// LobbyInfoData represents the information of a player in a multiplayer lobby.
type LobbyInfoData struct {
	AIControlled uint8  // 0 = Human Controlled, 1 = AI Controlled
	TeamID       uint8  // Team ID (see UDP specification appendix for more information)
	Nationality  uint8  // Nationality of the driver (see UDP specification appendix for more information)
	Name         string // Name of the participant in UTF-8 format. Truncated with ... if too long
	CarNumber    uint8  // Car number of the player
	ReadyStatus  uint8  // 0 = Not Ready, 1 = Ready, 2 = Spectating
}

// PacketLobbyInfoData represents the information of all players in a multiplayer lobby.
type PacketLobbyInfoData struct {
	Header       Header            // Packet header
	NumPlayers   uint8             // Number of players in the lobby
	LobbyPlayers [22]LobbyInfoData // Information of all players in a multiplayer lobby
}
