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

    @GetMapping("/status")
    public String getStatusHandler() {
        return "Hello, this is status endpoint";
    }
}
