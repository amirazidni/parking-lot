package amirazidni.parkinglot.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class MainController {

    @GetMapping("/status")
    public String status() {
        return "Hello, this is status endpoint";
    }
}
