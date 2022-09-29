package models

// CarMotionData represents the physics data for a car.
type CarMotionData struct {
	WorldPositionX     float64 // World space X position
	WorldPositionY     float64 // World space Y position
	WorldPositionZ     float64 // World space Z position
	WorldVelocityX     float64 // Velocity in world space X
	WorldVelocityY     float64 // Velocity in world space Y
	WorldVelocityZ     float64 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float64 // Lateral G-Force component
	GForceLongitudinal float64 // Longitudinal G-Force component
	GForceVertical     float64 // Vertical G-Force component
	Yaw                float64 // Yaw angle in radians
	Pitch              float64 // Pitch angle in radians
	Roll               float64 // Roll angle in radians
}

// PacketMotionData represents the physics data for all cars on track in a session.
type PacketMotionData struct {
	Header                 Header            // Packet header
	CarMotionData          [22]CarMotionData // Car motion data for all cars on track
	SuspensionPosition     [4]float64        // Order: Rear Left (RL), Rear Right (RR), Front Left (FL), Front Right (FR)
	SuspensionVelocity     [4]float64        // Order: RL, RR, FL, FR
	SuspensionAcceleration [4]float64        // Order: RL, RR, FL, FR
	WheelSpeed             [4]float64        // Speed of each wheel. Order: RL, RR, FL, FR
	WheelSlip              [4]float64        // Slip ratio for each wheel. Order: RL, RR, FL, FR
	LocalVelocityX         float64           // Velocity in local space X
	LocalVelocityY         float64           // Velocity in local space Y
	LocalVelocityZ         float64           // Velocity in local space Z
	AngularVelocityX       float64           // Angular velocity X-component
	AngularVelocityY       float64           // Angular velocity Y-component
	AngularVelocityZ       float64           // Angular velocity Z-component
	AngularAccelerationX   float64           // Angular acceleration X-component
	AngularAccelerationY   float64           // Angular acceleration Y-component
	AngularAccelerationZ   float64           // Angular acceleration Z-component
	FrontWheelsAngle       float64           // Current front wheels angle in radians
}
