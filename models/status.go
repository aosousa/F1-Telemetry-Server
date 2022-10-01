package models

// CarStatusData represents the status of a car in a race.
// This includes things like assist settings, tyre compound in use, fuel available, etc.
//
// ActualTyreCompound constants:
// 7 - Inter (F1 Modern)
// 8 - Wet (F1 Modern)
// 9 - Dry (F1 Classic)
// 10 - Wet (F1 Classic)
// 11 - Super Soft (F2)
// 12 - Soft (F2)
// 13 - Medium (F2)
// 14 - Hard (F2)
// 15 - Wet (F2)
// 16 - C5 (F1 Modern)
// 17 - C4 (F1 Modern)
// 18 - C3 (F1 Modern)
// 19 - C2 (F1 Modern)
// 20 - C1 (F1 Modern)
//
// VisualTyreCompound constants:
// 7 - Inter (F1 Modern and Classic)
// 8 - Wet (F1 Modern and Classic)
// 15 - Wet (F2 '19)
// 16 - Soft (F1 Modern and Classic)
// 17 - Medium (F1 Modern and Classic)
// 18 - Hard (F1 Modern and Classic)
// 19 - Super Soft (F2 '19)
// 20 - Soft (F2 '19)
// 21 - Medium (F2 '19)
// 22 - Hard (F2 '19)
//
// VehicleFIAFlags constants:
// -1 - Invalid/Unknown
// 0 - None
// 1 - Green
// 2 - Blue
// 3 - Yellow
// 4 - Red
//
// ERSDeployMode constants:
// 0 - None
// 1 - Medium
// 2 - Hotlap
// 3 - Overtake
type CarStatusData struct {
	TracionControl            uint8   // Traction control (0 = Off, 1 = Medium, 2 = Full)
	AntiLockBrakes            uint8   // ABS status (0 = Off, 1 = On)
	FuelMix                   uint8   // Fuel mix (0 = Lean, 1 = Standard, 2 = Rich, 3 = Max)
	FrontBrakeBias            uint8   // Front brake bias (percentage)
	PitLimiterStatus          uint8   // Pit limiter status (0 = Off, 1 = On)
	FuelInTank                float64 // Current fuel mass
	FuelCapacity              float64 // Fuel capacity
	FuelRemainingLaps         float64 // Fuel remaining in terms of laps (value shown on MFD)
	MaxRPM                    uint16  // Max RPM of the car, point of rev limiter
	IdleRPM                   uint16  // Idle RPM of the car
	MaxGears                  uint8   // Maximum number of gears
	DRSAllowed                uint8   // 0 = Not Allowed, 1 = Allowed
	DRSActivationDistance     uint16  // 0 = DRS not available, otherwise DRS available
	ActualTyreCompound        uint8   // Tyre compound in the car. See ActualTyreCompound constants above for more information
	VisualTyreCompound        uint8   // F1 visual (can be different from actual compound). See VisualTyreCompound constants above for more information
	TyresAgeLaps              uint8   // Age of the current set of tyres (in laps)
	VehicleFIAFlags           int8    // See VehicleFIAFlags constants above for more information
	ERSStoreEnergy            float64 // ERS energy store in Joules
	ERSDeployMode             uint8   // ERS deployment mode. See ERSDeployMode constants above for more information
	ERSHarvestedInThisLapMGUK float64 // ERS energy harvested this lap by MGU-K
	ERSHarvestedInThisLapMGUH float64 // ERS energy harvested this lap by MGU-H
	ERSDeployedThisLap        float64 // ERS energy deployed this lap
	NetworkPaused             uint8   // Whether the car is paused in a network game
}

// PacketCarStatusData represents the status of all cars in a race.
type PacketCarStatusData struct {
	Header        Header            // Packet header
	CarStatusData [22]CarStatusData // Status of all cars in a race
}
