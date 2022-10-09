package manage

import (
	"context"
)

type ManagerService interface {
	CreateParkingLot(ctx context.Context, slot uint8) error
	GetParkingLotStatus(ctx context.Context) (int, error)
}
