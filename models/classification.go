package models

// FinalClassificationData represents the final classification of a car at the end of a race.
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
type FinalClassificationData struct {
	Position          uint8    // Finishing position
	NumLaps           uint8    // Number of laps completed
	GridPosition      uint8    // Grid position of the car
	Points            uint8    // Number of points scored
	NumPitStops       uint8    // Number of pit stops made
	ResultStatus      uint8    // Result status. See ResultStatus constants above for more information
	BestLapTime       uint32   // Best lap time of the session in milliseconds
	TotalRaceTime     float64  // Total race time in seconds without penalties
	PenaltiesTime     uint8    // Total penalties accumulated in seconds
	NumPenalties      uint8    // Number of penalties applied to this driver
	NumTyreStints     uint8    // Number of tyre stints up to maximum
	TyreStintsActual  [8]uint8 // Actual tyres used by this driver
	TyreStintsVisual  [8]uint8 // Visual tyres used by this driver
	TyreStintsEndLaps [8]uint8 // The lap number that the stints end on
}

// PacketFinalClassificationData represents the final classification of all cars at the end of a race.
//
// Sent once at the end of a race, and the data will match with the post race results screen.
//
// Especially useful for multiplayer games where it is not always possible to send lap times
// on the final frame because of network delay.
type PacketFinalClassificationData struct {
	Header             Header                      // Packet header
	NumCars            uint8                       // Number of cars in the final classification
	ClassificationData [22]FinalClassificationData // Final classification data of all cars in a race
}
