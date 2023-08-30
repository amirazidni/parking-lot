package amirazidni.parkinglot.repository;

import org.springframework.stereotype.Repository;

import amirazidni.parkinglot.model.CarSlot;

@Repository
public class ParkingLotImpl implements ParkingLotRepository {
    private CarSlot[] parkingLot;

    public int createNew(int val) {
        this.parkingLot = new CarSlot[val];
        return this.parkingLot.length;
    }

    public CarSlot[] getParkingLot() {
        return this.parkingLot;
    }

    public int allocateParkingLot(String regisNum, String color) {
        if (this.parkingLot == null) {
            return 0;
        }

        int c = 0; // index iteration
        for (CarSlot carSlot : this.parkingLot) {
            if (carSlot == null) {
                CarSlot newCarSlot = new CarSlot();
                newCarSlot.setId(c + 1);
                newCarSlot.setPlateNumber(regisNum);
                newCarSlot.setColor(color);
                this.parkingLot[c] = newCarSlot;
                return c + 1;
            }
            c++;
        }

        return -1;
    }

    public int leaveParkingLot(int slotID) {
        if (this.parkingLot == null) {
            return 0;
        }

        if (slotID > this.parkingLot.length) {
            return -1;
        }

        this.parkingLot[slotID - 1] = null;

        return slotID;
    }
}
