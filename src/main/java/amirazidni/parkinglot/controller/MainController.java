package amirazidni.parkinglot.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import amirazidni.parkinglot.service.MainService;

@RestController
public class MainController {

    @Autowired
    private MainService mainService;

    @PostMapping("/create_parking_lot/{value}")
    public String createHandler(@PathVariable("value") String value) {
        return mainService.createParkingLot(value);
    }

    @PostMapping("/park/{value}/{attribute}")
    public String parkHandler(
            @PathVariable("value") String value,
            @PathVariable("attribute") String attribute) {
        return mainService.parkCar(value, attribute);
    }

    @PostMapping("/leave/{value}")
    public String leaveHandler(@PathVariable("value") String value) {
        return mainService.leavePark(value);
    }

    @GetMapping("/status")
    public String getStatusHandler() {
        return mainService.getParkingLotStatus();
    }

    @GetMapping("/cars_registration_numbers/colour/{value}")
    public String getCarsPlateHandler(@PathVariable("value") String value) {
        return mainService.getCarsPlate(value);
    }

    @GetMapping("/cars_slot/colour/{value}")
    public String getCarsSlotHandler(@PathVariable("value") String value) {
        return mainService.getCarsSlot(value);
    }
}
