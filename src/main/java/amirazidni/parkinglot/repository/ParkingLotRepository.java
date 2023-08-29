package amirazidni.parkinglot.repository;

import org.springframework.stereotype.Repository;

import amirazidni.parkinglot.model.CarSlot;

@Repository
public class ParkingLotRepository {

    private CarSlot[] parkingLot;

    public int createNew(int val) {
        this.parkingLot = new CarSlot[val];
        return this.parkingLot.length;
    }
}
