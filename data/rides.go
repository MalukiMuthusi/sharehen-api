package data

// for a detailed explanation on the use of pointers for specifying the API see
// https://willnorris.com/2014/05/go-rest-apis-and-pointers/

// Rider requests for rides
type Rider struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Ride owner volunteers to share ride
type RideOwner struct {
	Id    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	CarId *string `json:"car_id,omitempty"`
}

// RideOwner volunteers and asks for riders
type RideOwnerOffer struct {
	Id          *string  `json:"id,omitempty"`
	OfferAmount *FeeType `json:"offer_amount,omitempty"`
	RideOwnerId *string  `json:"ride_owner_id,omitempty"`
}

// A ride is shared by a rideOwner and a rider
type Ride struct {
	Id        *string    `json:"id,omitempty"`
	Rider     *Rider     `json:"rider,omitempty"`
	RideOwner *RideOwner `json:"ride_owner"`
}

// The request made by a rider for a ride
type RiderRequest struct {
	Id          *string  `json:"id,omitempty"`
	RiderId     *string  `json:"rider_ide,omitempty"`
	OfferAmount *FeeType `json:"offer_amount,omitempty"`
}

// the currency and the amount
type FeeType struct {
	Currency *string  `json:"currency,omitempty"`
	Amount   *float32 `json:"amount,omitempty"`
}
