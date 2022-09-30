package models

// Event details packet is different for each type of event so I am using the packet header
// and event string code properties in each struct for this one and doing a third buffer
// read according to the event string code received.
// TODO: Research a better implementation of struct / union combination.
//
// Event string codes (Code - Event - Description):
// SSTA - Session Started - Sent when the session starts
// SEND - Session Ended - Sent when the session ends
// FTLP - Fastest Lap - When a driver achieves the fastest lap
// RTMT - Retirement - When a driver retires
// DRSE - DRS Enabled - Race control have enabled DRS
// DRSD - DRS Disabled - Race control have disabled DRS
// TMPT - Team Mate in Pits - Team mate has entered the pits
// CHQF - Chequered Flag - Chequered flag has been waved
// RCWN - Race Winner - Race winner is announced
// PENA - Penalty Issued - A penalty has been issued (details in event)
// SPTP - Speed Trap Triggered - Speed trap has been triggered by fastest speed
// STLG - Start Lights - Start lights (number shown)
// LGOT - Lights Out - Lights out
// DTSV - Drive Through Served - Drive through penalty served
// SGSV - Stop Go Served - Stop go penalty served
// FLBK - Flashback - Flashback activated
// BUTN - Button Status - Button status changed

// FastestLap represents the information sent in a FTLP event
type FastestLap struct {
	Header          Header  // Packet header
	EventStringCode string  // Event string code
	VehicleIdx      uint8   // Vehicle index of car achieving fastest lap
	LapTime         float64 // Lap time in seconds
}

// Retirement represents the information sent in a RTMT event
type Retirement struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
	VehicleIdx      uint8  // Vehicle index of car retiring
}

// TeamMateInPits represents the information sent in a TMPT event
type TeamMateInPits struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
	VehicleIdx      uint8  // Vehicle index of team mate
}

// RaceWinner represents the information sent in a RCWN event
type RaceWinner struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
	VehicleIdx      uint8  // Vehicle index of the race winner
}

// Penalty represents the information sent in a PENA event
type Penalty struct {
	Header            Header // Packet header
	EventStringCode   string // Event string code
	PenaltyType       uint8  // Penalty type - see UDP specification appendix for more information
	InfringementType  uint8  // Infringement type - see UDP specification appendix for more information
	VehicleIdx        uint8  // Vehicle index of the car the penalty is applied to
	OtherVehicleIndex uint8  // Vehicle index of the car involved
	Time              uint8  // Time gained, or time spent doing action in seconds
	LapNumber         uint8  // Lap the penalty occurred on
	PlacesGained      uint8  // Number of places gained by this
}

// SpeedTrap represents the information sent in a SPTP event
type SpeedTrap struct {
	Header                     Header  // Packet header
	EventStringCode            string  // Event string code
	VehicleIdx                 uint8   // Vehicle index of the vehicle triggering the speed trap
	Speed                      float64 // Top speed achieved in kp/h
	IsOverallFastestInSession  uint8   // 1 if overall fastest speed in session, 0 otherwise
	IsDriverFastestInSession   uint8   // 1 if fastest speed for driver in session, 0 otherwise
	FastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest in the session
	FastestSpeedInSession      float64 // Speed of the vehicle that's the fastest in this session
}

// StartLights represents the information sent in a STLG event
type StartLights struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
	NumLights       uint8  // Number of lights showing
}

// DriveThroughPenaltyServed represents the information sent in a DTSV event
type DriveThroughPenaltyServed struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
	VehicleIdx      uint8  // Vehicle index of the vehicle serving a drive through penalty
}

// StopGoPenaltyServed represents the information sent in a SGSV event
type StopGoPenaltyServed struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
	VehicleIdx      uint8  // Vehicle index of the vehicle serving a stop and go penalty
}

// Flashback represents the information sent in a FLBK event
type Flashback struct {
	Header                   Header  // Packet header
	EventStringCode          string  // Event string code
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float64 // Session time flashed back to
}

// Buttons represents the information sent in a BUTN event
type Buttons struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
	ButtonStatus    uint32 // Bit flags specifying which buttons are being pressed currently - see UDP specification appendix for more information
}

// PacketEventData represents the information sent in an event packet, and used to
// identify which specific event occurred
type PacketEventData struct {
	Header          Header // Packet header
	EventStringCode string // Event string code
}
