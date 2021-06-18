package data

// for a detailed explanation on the use of pointers for specifying the API see
// https://willnorris.com/2014/05/go-rest-apis-and-pointers/

// Rider requests for rides
type Rider struct {
	Id   *string
	Name *string
}

// Ride owner volunteers to share ride
type RideOwner struct {
	Id    *string
	Name  *string
	CarId *string
}

// RideOwner volunteers and asks for riders
type RideOwnerOffer struct {
	Id          *string
	OfferAmount *FeeType
	RideOwnerId *string
}

// A ride is shared by a rideOwner and a rider
type Ride struct {
	Id        *string    `json:"ID"`
	Rider     *Rider     `json:"rider"`
	RideOwner *RideOwner `json:"ride_owner"`
}

// The request made by a rider for a ride
type RiderRequest struct {
	Id          *string
	RiderId     *string
	OfferAmount *FeeType
}

// the currency and the amount
type FeeType struct {
	Currency *string
	Amount   *float32
}
