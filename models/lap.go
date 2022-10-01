package models

// LapData represents the details of a car in a session.
//
// DriverStatus constants:
// 0 - In Garage
// 1 - Flying Lap
// 2 - In Lap
// 3 - Out Lap
// 4 - On Track
//
// ResultStatus constants:
// 0 - Invalid
// 1 - Inactive
// 2 - Active
// 3 - Finished
// 4 - DNF
// 5 - DSQ
// 6 - Not Classified
// 7 - Retired
type LapData struct {
	LastLapTime                 uint32  // Last lap time in milliseconds
	CurrentLapTime              uint32  // Current time around the lap in milliseconds
	SectorOneTime               uint16  // Sector 1 time in milliseconds
	SectorTwoTime               uint16  // Sector 2 time in milliseconds
	LapDistance                 float64 // Distance vehicle is around current lap in metres (can be negative if the line hasn't been crossed yet)
	TotalDistance               float64 // Total distance travelled in session in metres (can be negative if the line hasn't been crossed yet)
	SafetyCarDelta              float64 // Delta in seconds for safety car
	CarPosition                 uint8   // Car race position
	CurrentLapNum               uint8   // Current lap number
	PitStatus                   uint8   // 0 = None, 1 = Pitting, 2 = In Pit
	NumPitStops                 uint8   // Number of pit stops taken in this race
	Sector                      uint8   // 0 = S1, 1 = S2, 2 = S3
	CurrentLapInvalid           uint8   // Current lap valid status (0 = Valid, 1 = Invalid)
	Penalties                   uint8   // Accumulated time penalties in seconds to be added
	Warnings                    uint8   // Accumulated number of warnings issued
	NumUnservedDriveThroughPens uint8   // Number of drive through penalties left to serve
	NumUnservedStopGoPens       uint8   // Number of stop and go penalties left to serve
	GridPosition                uint8   // Grid position the vehicle started the race in
	DriverStatus                uint8   // Status of the driver. See DriverStatus constants above for more information
	ResultStatus                uint8   // Result status. See ResultStatus constants above for more information
	PitLaneTimerActive          uint8   // Pit lane timing. 0 = Inactive, 1 = Active
	PitLaneTimeInLane           uint16  // If PitLaneTimerActive = 1, the current time spent in the pit lane in milliseconds
	PitStopTimer                uint16  // Time of the actual pit stop in milliseconds
	PitStopShouldServePen       uint8   // Whether the car should serve a penalty at the current pit stop
}

// PacketLapData represents the details of all cars in a session.
type PacketLapData struct {
	Header                  Header      // Packet header
	LapData                 [22]LapData // Lap data for all cars on track
	TimeTrialPBCarIndex     uint8       // Index of Personal Best car in time trial (255 if invalid)
	TimeTrialRivalCardIndex uint8       // Index of Rival car in time trial (255 if invalid)
}
