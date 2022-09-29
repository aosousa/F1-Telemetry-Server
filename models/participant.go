package models

// ParticipantData represents the information of a participant in a race.
type ParticipantData struct {
	AIControlled  uint8    // 0 = Human Controlled, 1 = AI Controlled
	DriverID      uint8    // Driver ID. See UDP specification appendix for more information, 255 if network human
	NetworkID     uint8    // Network ID. Unique identifier for network players
	TeamID        uint8    // Team ID. See UDP specification appendix for more information
	MyTeam        uint8    // My Team flag. 1 = My Team, 0 otherwise
	RaceNumber    uint8    // Race number of the car
	Nationality   uint8    // Nationality of the driver. See UDP specification for more information
	Name          [48]rune // Name of participant in UTF-8. Truncated with ... if too long
	YourTelemetry uint8    // Player's UDP setting. 0 = Restricted, 1 = Public
}

// PacketParticipantData represents the information of all participants in a race.
type PacketParticipantData struct {
	Header        Header              // Packet header
	NumActiveCars uint8               // Number of active cars in the data - should match number of cars on HUD
	Participants  [22]ParticipantData // Data of all participants in a race
}
