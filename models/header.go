package models

// Header represents the header information sent in every packet.
//
// Most important property is PacketID, which identifies the type of packet that was sent
// and will be used to decode all of the packet data and use it in the appropriate model.
//
// Packet IDs (information retrieved from the official F1 22 UDP Specification):
// 0 - Motion. Contains all motion data for the player's car - only sent while the player is in control.
// 1 - Session. Data about the session - track, time left.
// 2 - Lap Data. Data about all the lap times of the cars in the session.
// 3 - Event. Various notable events that happen during a session (see more information in Event model).
// 4 - Participants. List of participants in the session. Mostly relevant for multiplayer.
// 5 - Car Setups. Packet detailing car setups for all cars in a race.
// 6 - Car Telemetry. Telemetry data for all cars.
// 7 - Car Status. Status data for all cars.
// 8 - Final Classification. Final calssification confirmation at the end of a race.
// 9 - Lobby Info. Information about players in a multiplayer lobby.
// 10 - Car Damage. Damage status for all cars.
// 11 - Session History. Lap and tyre data for a session.
type Header struct {
	PacketFormat         uint16  // 2022
	GameMajorVersion     uint8   // Game major version (e.g., X.00)
	GameMinorVersion     uint8   // Game minor version (e.g., 1.XX)
	PacketVersion        uint8   // Version of the packet type (starts at 1)
	PacketID             uint8   // Identifier for packet type
	SessionUID           uint8   // Unique identifier for the session
	SessionTime          float64 // Session timestamp
	FrameIdentifer       uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex       uint8   // Index of player's car in the array
	SecondaryPlayerIndex uint8   // Index of secondary player's car in the array (split screen). 255 if no second player
}
