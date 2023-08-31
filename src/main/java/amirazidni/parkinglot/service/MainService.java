package amirazidni.parkinglot.service;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import amirazidni.parkinglot.model.CarSlot;
import amirazidni.parkinglot.repository.ParkingLotRepository;

enum ReturnCase {
    ID,
    PlateNumber,
}

@Service
public class MainService {

    @Autowired
    private ParkingLotRepository parkingLotRepository;

    private final ResponseStatusException EmptyParkingLot = new ResponseStatusException(
            HttpStatus.BAD_REQUEST,
            "Empty parking lot");

    public String createParkingLot(String value) {
        int qty = Integer.parseInt(value);

        if (qty < 1) {
            throw new ResponseStatusException(
                    HttpStatus.BAD_REQUEST,
                    "invalid slot value");
        }

        int total = parkingLotRepository.createNew(qty);

        return String.format("Created a parking lot with %s slots\n", total);
    }

    public String parkCar(String regisNum, String color) {
        if (regisNum.trim().isEmpty() || color.trim().isEmpty()) {
            throw new ResponseStatusException(
                    HttpStatus.BAD_REQUEST,
                    "plate number or color should not empty");
        }
        int no = parkingLotRepository.allocateParkingLot(regisNum, color);

        switch (no) {
            case 0:
                throw EmptyParkingLot;
            case -1:
                throw new ResponseStatusException(
                        HttpStatus.BAD_REQUEST,
                        "Sorry, parking lot is full");
            default:
                break;
        }

        return String.format("Allocated slot number: %d\n", no);
    }

    public String leavePark(String slotNumber) {
        if (slotNumber.trim().isEmpty()) {
            throw new ResponseStatusException(
                    HttpStatus.BAD_REQUEST,
                    "invalid slot number");
        }

        int no = Integer.parseInt(slotNumber);

        int result = parkingLotRepository.leaveParkingLot(no);

        switch (result) {
            case 0:
                throw EmptyParkingLot;
            case -1:
                throw new ResponseStatusException(
                        HttpStatus.BAD_REQUEST,
                        "selected slot is unavailable");
            default:
                break;
        }

        if (result != no) {
            throw new ResponseStatusException(
                    HttpStatus.INTERNAL_SERVER_ERROR,
                    "failed to process leave request");
        }

        return String.format("Slot number %d is free\n", result);
    }

    public String getParkingLotStatus() {
        List<String> responses = new ArrayList<String>();
        responses.add("Slot No. Registration No Colour");

        CarSlot[] parkingLot = parkingLotRepository.getParkingLot();

        for (CarSlot carSlot : parkingLot) {
            if (carSlot != null) {
                responses.add(
                        String.format("%d %s %s",
                                carSlot.getId(),
                                carSlot.getPlateNumber(),
                                carSlot.getColor()));
            }
        }

        return String.join("\n", responses) + "\n";
    }

    public String getCarsPlate(String colour) {
        return getCarsByColor(colour, ReturnCase.PlateNumber);
    }

    // util function
    private String getCarsByColor(String colour, ReturnCase returnCase) {
        if (colour.trim().isEmpty()) {
            throw new ResponseStatusException(
                    HttpStatus.BAD_REQUEST,
                    "colour value should not be empty");
        }

        List<String> cars = new ArrayList<String>();
        CarSlot[] parkingLot = parkingLotRepository.getParkingLot();
        for (CarSlot carSlot : parkingLot) {
            if (carSlot.getColor().equals(colour)) {
                switch (returnCase) {
                    case ID:
                        cars.add(String.format("%d", carSlot.getId()));
                        break;
                    case PlateNumber:
                        cars.add(carSlot.getPlateNumber());
                        break;
                    default:
                        throw new ResponseStatusException(
                                HttpStatus.BAD_REQUEST,
                                "not recognized");
                }
            }
        }

        String responses = String.join(", ", cars) + "\n";
        return responses;
    }

}
