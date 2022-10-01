package models

// CarDamageData represents the damage and wear information for a car in a race.
type CarDamageData struct {
	TyresWear            [4]float64 // Tyre wear (percentage). Order: RL, RR, FL, FR
	TyresDamage          [4]uint8   // Tyre damage (percentage). Order: RL, RR, FL, FR
	BrakesDamage         [4]uint8   // Brakes damage (percentage). Order: RL, RR, FL, FR
	FrontLeftWingDamage  uint8      // Front left wing damage (percentage)
	FrontRightWingDamage uint8      // Front right wind damage (percentage)
	RearWingDamage       uint8      // Rear wing damage (percentage)
	FloorDamage          uint8      // Floor damage (percentage)
	DiffuserDamage       uint8      // Diffuser damage (percentage)
	SidepodDamage        uint8      // Sidepod damage (percentage)
	DRSFault             uint8      // Indicator for DRS fault (0 = OK, 1 = Fault)
	ERSFault             uint8      // Indicator for ERS fault (0 = OK, 1 = Fault)
	GearboxDamage        uint8      // Gear box damage (percentage)
	EngineDamage         uint8      // Engine damage (percentage)
	EngineMGUHWear       uint8      // MGU-H wear (percentage)
	EngineESWear         uint8      // Energy Store wear (percentage)
	EngineCEWear         uint8      // Control Electronics wear
	EngineICEWear        uint8      // ICE wear (percentage)
	EngineMGUKWear       uint8      // MGU-K wear (percentage)
	EngineTCWear         uint8      // Turbo Charger wear (percentage)
	EngineBlown          uint8      // Indicator for engine blown (0 = OK, 1 = Fault)
	EngineSeized         uint8      // Indicator for engine seized (0 = OK, 1 = Fault)
}

// PacketCarDamageData represents the damage and wear information for all cars in a race.
type PacketCarDamageData struct {
	Header        Header            // Packet header
	CarDamageData [22]CarDamageData // Damage and wear information for all cars in a race
}
