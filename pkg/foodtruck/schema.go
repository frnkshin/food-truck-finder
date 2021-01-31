package foodtruck

type FoodTruck struct {
	Dayorder         string `json:"dayorder"`
	Dayofweekstr     string `json:"dayofweekstr"`
	Starttime        string `json:"starttime"`
	Endtime          string `json:"endtime"`
	Permit           string `json:"permit"`
	Location         string `json:"location"`
	Locationdesc     string `json:"locationdesc"`
	Optionaltext     string `json:"optionaltext"`
	Locationid       string `json:"locationid"`
	Start24          string `json:"start24"`
	End24            string `json:"end24"`
	Cnn              string `json:"cnn"`
	AddrDateCreate   string `json:"addr_date_create"`
	AddrDateModified string `json:"addr_date_modified"`
	Block            string `json:"block"`
	Lot              string `json:"lot"`
	Coldtruck        string `json:"coldtruck"`
	Applicant        string `json:"applicant"`
	X                string `json:"x"`
	Y                string `json:"y"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	Location2        struct {
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		HumanAddress string `json:"human_address"`
	} `json:"location_2"`
}

type FoodTrucks []FoodTruck
