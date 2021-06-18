package data

import (
	"encoding/json"
	"net/http"

	"github.com/MalukiMuthusi/sharehen-api/utils"
)

// returns all the rides that are active in the system
// that is the rides the rider and rideOwner have agreed on
func GetRidesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// create a dummy list of rides
	// return the json value of the rides

	b, err := json.Marshal(ridesList())
	utils.PanicWithError(err, "failed to marshall")

	_, err = w.Write(b)
	utils.PanicWithError(err, "failed to write the response")
}

func ridesList() []Ride {
	ride1Id := string("0gytba")
	ride2Id := string("anuh1q")

	rider1Name := string("Joseph Rider")
	riderId2 := string("0gytba")
	riderName2 := string("Galileo Galilee")

	rideOwnerId1 := string("0nba_bsn")
	riderOwnerName1 := string("Anthony Carlos")

	rideOwnerId2 := string("qwn_y")
	rideOwnerName2 :=
		string("Silas Othonios")

	carId1 := string("kbk100N")
	carId2 := string("zedonkn")

	rides := []Ride{
		{Id: &ride1Id, Rider: &Rider{Id: &ride1Id, Name: &rider1Name}, RideOwner: &RideOwner{Id: &rideOwnerId1, Name: &riderOwnerName1, CarId: &carId1}},
		{Id: &ride2Id, Rider: &Rider{Id: &riderId2, Name: &riderName2}, RideOwner: &RideOwner{Id: &rideOwnerId2, Name: &rideOwnerName2, CarId: &carId2}},
	}

	return rides
}
