package models

// In multiplayer games, other player cars will appear as blank.
// It is only possible to see your own car setup, as well as the setup of AI-controlled cars.

// CarSetupData represents the car setup details of a vehicle in a session.
type CarSetupData struct {
	FrontWing              uint8   // Front wing aero
	RearWing               uint8   // Rear wing aero
	DiffOnThrottle         uint8   // Differential adjustment on throttle (percentage)
	DiffOffThrottle        uint8   // Differential adjustment off throttle (percentage)
	FrontCamber            float64 // Front camber angle (suspension geometry)
	RearCamber             float64 // Rear camber angle (suspension geometry)
	FrontToe               float64 // Front toe angle (suspension geometry)
	RearToe                float64 // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   // Front suspension
	RearSuspension         uint8   // Rear suspension
	FrontAntiRollBar       uint8   // Front anti-roll bar
	RearAntiRollBar        uint8   // Rear anti-roll bar
	FrontSuspensionHeight  uint8   // Front ride height
	RearSuspensionHeight   uint8   // Rear ride height
	BrakePressure          uint8   // Brake pressure (percentage)
	BrakeBias              uint8   // Brake bias (percentage)
	RearLeftTyrePressure   float64 // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float64 // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float64 // Front left tyre pressure (PSI)
	FrontRightTyrePressure float64 // Front right tyre pressure (PSI)
	Ballast                uint8   // Ballast
	FuelLoad               float64 // Fuel load
}

// PacketCarSetupData represents the car setup details of all vehicles in a session
type PacketCarSetupData struct {
	Header    Header           // Packet header
	CarSetups [22]CarSetupData // Setup data of all vehicles in a session
}
