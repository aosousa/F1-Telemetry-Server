package models

// CarTelemetryData represents telemetry details for a car in a race.
// It details various values that would be recorded on the car such as speed, throttle
// application, DRS status, etc.
type CarTelemetryData struct {
	Speed                   uint16     // Speed of the car in kp/h
	Throttle                float64    // Amount of throttle applied (0.0 to 1.0, can be converted to a percentage)
	Steer                   float64    // Amount of steering applied (-1.0 full lock left, 1.0 full lock right)
	Brake                   float64    // Amount of brake applied (0.0 to 1.0, can be converted to a percentage)
	Clutch                  uint8      // Amount of clutch applied (0 to 100)
	Gear                    int8       // Gear selected (1-8, 0 = Neutral, -1 = Reverse)
	EngineRPM               uint16     // Engine RPM
	DRS                     uint8      // Engine RPM
	RevLightsPercent        uint8      // Rev lights indicator (percentage)
	RevLightsBitValue       uint16     // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	BrakesTemperature       [4]uint16  // Brakes temperature in Celsius. Order: RL, RR, FL, FR
	TyresSurfaceTemperature [4]uint8   // Tyres surface temperature in Celsius. Order: RL, RR, FL, FR
	TyresInnerTemperature   [4]uint8   // Tyres inner temperature in Celsius. Order: RL, RR, FL, FR
	EngineTemperature       uint16     // Engine temperature in Celsius
	TyresPressure           [4]float64 // Tyres pressure (PSI). Order: RL, RR, FL, FR
	SurfaceType             [4]uint8   // Driving surface. See UDP specification appendix for more information
}

// PacketCarTelemetryData represents telemetry details for all cars in a race.
type PacketCarTelemetryData struct {
	Header           Header               // Packet header
	CarTelemetryData [22]CarTelemetryData // Telemetry details for all cars in a race
}
