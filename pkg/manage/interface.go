package manage

import (
	"context"
	parkinglot "parking-lot/pkg/parking-lot"
)

type ManagerService interface {
	CreateParkingLot(ctx context.Context, slot int) error
	AllocateParkingLot(ctx context.Context, regisNum, color string) (slotNum int, err error)
	GetParkingLot(ctx context.Context) (*[]parkinglot.CarSlot, error)
	LeaveParkingLot(ctx context.Context, slotID int) error
}
