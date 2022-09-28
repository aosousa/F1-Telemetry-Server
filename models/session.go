package models

// MarshalZone represents the information of a marshal post
//
// ZoneFlag constants:
// -1 - Invalid / Unkown
// 0 - None
// 1 - Green
// 2 - Blue
// 3 - Yellow
// 4 - Red
type MarshalZone struct {
	ZoneStart float64 // Fraction (0..1) of way through the lap where the marshal zone starts
	ZoneFlag  int8    // Flag information
}

// WeatherForecastSample represents the information of a weather forecast sample of a session
//
// SessionType constants:
// 0 - Unknown
// 1 - FP1
// 2 - FP2
// 3 - FP3
// 4 - Short Practice
// 5 - Q1
// 6 - Q2
// 7 - Q3
// 8 - Short Q
// 9 - One Shot Qualifying
// 10 - Race
// 11 - Race 2
// 12 - Race 3
// 13 - Time Trial
//
// Weather constants:
// 0 - Clear
// 1 - Light Cloud
// 2 - Overcast
// 3 - Light Rain
// 4 - Heavy Rain
// 5 - Storm
type WeatherForecastSample struct {
	SessionType            uint8 // See SessionType constants above for more information
	TimeOffset             uint8 // Time in minutes the forecast is for
	Weather                uint8 // Weather information. See Weather constants above for more infromation
	TrackTemperature       int8  // Track temperature in Celsius
	TrackTemperatureChange int8  // Track temperature change. 0 = Up, 1 = Down, 2 = No Change
	AirTemperature         int8  // Air temperature in Celsius
	AirTemperatureChange   int8  // Air temperature change. 0 = Up, 1 = Down, 2 = No Change
	RainPercentage         uint8 // Rain percentage (0-100)
}

// PacketSessionData represents the full session data.
//
// Formula constants:
// 0 = F1 Modern
// 1 = F1 Classic
// 2 = F2
// 3 = F1 Generic
// 4 = Beta
// 5 = Supercars
// 6 = Esports
// 7 = F2 2021
//
// SafetyCarStatus constants:
// 0 - No safety car status
// 1 - Full safety car
// 2 - Virtual safety car
// 3 - Formation lap
//
// SessionLength constants:
// 0 - None
// 2 - Very Short
// 3 - Short
// 4 - Medium
// 5 - Medium Long
// 6 - Long
// 7 - Full
type PacketSessionData struct {
	Header                    Header                    // Header
	Weather                   uint8                     // Weather information. See Weather constants above for more information
	TrackTemperature          int8                      // Track temperature in Celsius
	AirTemperature            int8                      // Air temperature in Celsius
	TotalLaps                 uint8                     // Total number of laps in this race
	TrackLength               uint16                    // Track length in metres
	SessionType               uint8                     // Session type. See Session Type constants above for more information
	TrackID                   int8                      // ID of the track. -1 for unknown, see UDP specification manual for IDs of the tracks
	Formula                   uint8                     // Formula type. See Formula constants above for more information
	SessionTimeLeft           uint16                    // Time left in session in seconds
	SessionDuration           uint16                    // Session duration in seconds
	PitSpeedLimit             uint8                     // Pit speed limit (kp/h)
	GamePaused                uint8                     // Whether the game is paused - network game only
	IsSpectating              uint8                     // Whether the game is spectating
	SpectatorCarIndex         uint8                     // Index of the car being select
	SLIProNativeSupport       uint8                     // SLI Pro support, 0 = Inactive, 1 = Active
	NumMarshalZones           uint8                     // Number of marshal zones to follow
	MarshalZones              [21]MarshalZone           // List of marshal zones (maximum 21)
	SafetyCarStatus           uint8                     // Safety car information. See SafetyCarStatus constants above for more information
	NetworkGame               uint8                     // 0 - Offline, 1 - Online
	NumWeatherForecastSamples uint8                     // Number of weather samples to follow
	WeatherForecastSamples    [56]WeatherForecastSample // Array of weather forecast samples
	ForecastAccuracy          uint8                     // 0 = Perfect, 1 = Approximate
	AIDifficulty              uint8                     // AI Difficulty Rating (0-110)
	SeasonLinkIdentifier      uint32                    // Identifier for season (persists across saves)
	WeekendLinkIdentifier     uint32                    // Identifier for weekend (persists across saves)
	SessionLinkIdentifier     uint32                    // Identifier for session (persists across saves)
	PitStopWindowIdealLap     uint8                     // Ideal lap to pit on for current strategy (player)
	PitStopWindowLatestLap    uint8                     // Latest lap to pit on for current strategy (player)
	PitStopRejoinPosition     uint8                     // Predicted position to rejoin on track (player)
	SteeringAssist            uint8                     // 0 = Off, 1 = On
	BrakingAssist             uint8                     // 0 = Off, 1 = Low, 2 = Medium, 3 = High
	GearboxAssist             uint8                     // 1 = Manual, 2 = Manual & Suggested Gear, 3 = Auto
	PitAssist                 uint8                     // 0 = Off, 1 = On
	PitReleaseAssist          uint8                     // 0 = Off, 1 = On
	ERSAssist                 uint8                     // 0 = Off, 1 = On
	DRSAssist                 uint8                     // 0 = Off, 1 = On
	DynamicRacingLine         uint8                     // 0 = Off, 1 = Corners Only, 2 = Full
	DynamicRacingLineType     uint8                     // 0 = 2D, 1 = 3D
	GameMode                  uint8                     // Game mode ID - see UDP specification information
	RuleSet                   uint8                     // Rule set - see UDP specification information
	TimeOfDay                 uint32                    // Local time of day - minutes since midnight
	SessionLength             uint                      // Session length. See SessionLength constants above for more information
}

// TODO: each model should have its own save method (to save data in database)
