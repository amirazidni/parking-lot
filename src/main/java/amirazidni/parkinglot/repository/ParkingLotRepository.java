package amirazidni.parkinglot.repository;

import org.springframework.stereotype.Repository;

import amirazidni.parkinglot.model.CarSlot;

@Repository
public interface ParkingLotRepository {
    int createNew(int val);

    int allocateParkingLot(String regisNum, String color);

    CarSlot[] getParkingLot();

    void leaveParkingLot(int id);
}
