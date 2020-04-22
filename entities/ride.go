package entities

//Ride is ride object
type Ride struct {
	Aid int `json:"aid"`
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

//RideDetail is struct for viewing details of ride
type RideDetail struct {
	Rid int `json:"rid"`
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
	Driver string `json:"driver"`
	DriverPicture string `json:"driver_picture"`
	Rides int `json:"rides"`
	Ratings int `json:"ratings"`
	AverageRating float32 `json:"average_rating"`
	CommentsAboutDriver []struct {
		Rid int `json:"rid"`
		Date string `json:"date"`
		Rating int `json:"rid"`
		Comment string `json:"rid"`	
	}

}

//CreateRideDetail inits object
func CreateRideDetail(ride Ride, account interface{}) RideDetail{
	
}
//Cra
