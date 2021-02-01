package foodtruck

import (
	"fmt"
	"strings"
)

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

func (truck *FoodTruck) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Name                 : %s\n", truck.Applicant)
	fmt.Fprintf(&sb, "Description          : %s\n", truck.Optionaltext)
	fmt.Fprintf(&sb, "Business Hour        : %s ~ %s on %s\n", truck.Starttime, truck.Endtime, truck.Dayofweekstr)
	fmt.Fprintf(&sb, "Location Description : %s\n", truck.Locationdesc)
	fmt.Fprintf(&sb, "Lat/Long             : %s, %s\n", truck.Latitude, truck.Longitude)
	return sb.String()
}
