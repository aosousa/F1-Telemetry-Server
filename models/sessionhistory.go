package models

// Session History packet contains lap times and tyre usage for the session. This packet works
// slightly differently to the remaining packets. In order to reduce CPU and bandwidth, each packet
// is related to a specific vehicle and is sent every 1/20s, adn the vehicle being sent is cycled through.
//
// In a 20 car race, there should be an update for each vehicle at leat once per second.

// LapHistoryData represents the information of one lap of a car in a race.
//
// LapValidBitFlags constants:
// 0x01 - Lap Valid
// 0x02 - Sector 1 Valid
// 0x04 - Sector 2 Valid
// 0x08 - Sector 3 Valid
type LapHistoryData struct {
	LapTime          uint32 // Lap time in milliseconds
	SectorOneTime    uint16 // Sector 1 time in milliseconds
	SectorTwoTime    uint16 // Sector 2 time in milliseconds
	SectorThreeTime  uint16 // Sector 3 time in milliseconds
	LapValidBitFlags uint8  // See LapValidBitFlags constants above for more information
}

// TyreStintHistoryData represents the information of the tyre stint of a car in a race.
type TyreStintHistoryData struct {
	EndLap             uint8 // Lap the tyre usage ends on
	TyreActualCompound uint8 // Actual tyres used by the driver
	TyreVisualCompound uint8 // Visual tyres used by the driver
}

// PacketSessionHistoryData represents the lap and tyre stint data of a car in a race.
type PacketSessionHistoryData struct {
	Header                Header                  // Packet header
	CarIdx                uint8                   // Index of the car this lap data relates to
	Numlaps               uint8                   // Number of laps in the data (including current partial lap)
	NumTyreStints         uint8                   // Number of tyre stints in the data
	BestLapTimeLapNum     uint8                   // Lap the best lap time was achieved on
	BestSectorOneLapNum   uint8                   // Lap the best Sector 1 time was achieved on
	BestSectorTwoLapNum   uint8                   // Lap the best Sector 2 time was achieved on
	BestSectorThreeLapNum uint8                   // Lap the best Sector 3 time was achieved on
	LapHistoryData        [100]LapHistoryData     // Maximum of 100 laps history data
	TyreStintHistoryData  [8]TyreStintHistoryData // Maximum of 8 tyre stints history data
}
