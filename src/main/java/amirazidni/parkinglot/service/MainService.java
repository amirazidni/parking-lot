package amirazidni.parkinglot.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import amirazidni.parkinglot.repository.ParkingLotRepository;

@Service
public class MainService {

    @Autowired
    private ParkingLotRepository parkingLotRepository;

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

}
