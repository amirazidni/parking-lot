package manage

import (
	"context"
	"fmt"
	parkinglot "parking-lot/pkg/parking-lot"
)

func (m *Manager) CreateParkingLot(ctx context.Context, slot uint8) error {
	newParkingLot := make([]parkinglot.CarSlot, slot)
	m.ParkingLot = &newParkingLot
	return nil
}

func (m *Manager) GetParkingLotStatus(ctx context.Context) (int, error) {
	if m.ParkingLot == nil {
		return 0, fmt.Errorf("Empty parking lot")
	}
	slot := len(*m.ParkingLot)
	return slot, nil
}
