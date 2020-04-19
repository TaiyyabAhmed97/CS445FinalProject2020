package entities

//Ride is ride object
type Ride struct {
	LocationInfo struct {
		FromCity string `json:"from_city"`
		FromZip  string `json:"from_zip"`
		ToCity   string `json:"to_city"`
		ToZip    string `json:"to_zip"`
	} `json:"location_info"`
	DateTime struct {
		Date string `json:"date"`
		Time string `json:"time"`
	} `json:"date_time"`
	CarInfo struct {
		Make        string `json:"make"`
		Model       string `json:"model"`
		Color       string `json:"color"`
		PlateState  string `json:"plate_state"`
		PlateSerial string `json:"plate_serial"`
	} `json:"car_info"`
	MaxPassengers      int     `json:"max_passengers"`
	AmountPerPassenger float32 `json:"amount_per_passenger"`
	Conditions         string  `json:"conditions"`
}

//Cra
