package manage

import (
	parkinglot "parking-lot/pkg/parking-lot"
)

func NewManager() ManagerService {
	return &Manager{}
}

type Manager struct {
	ParkingLot *[]parkinglot.CarSlot
}
