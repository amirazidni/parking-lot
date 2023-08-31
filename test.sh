#!/usr/bin/env bash

curl -X POST localhost:8080/create_parking_lot/6
curl -X POST http://localhost:8080/park/B-1234-RFS/Black
curl -X POST http://localhost:8080/park/B-1999-RFD/Green
curl -X POST http://localhost:8080/park/B-1000-RFS/Black
curl -X POST http://localhost:8080/park/B-1777-RFU/BlueSky
curl -X POST http://localhost:8080/park/B-1701-RFL/Blue
curl -X POST http://localhost:8080/park/B-1141-RFS/Black
curl -X POST http://localhost:8080/leave/4
curl -X GET http://localhost:8080/status
curl -X POST http://localhost:8080/park/B-1333-RFS/Black
curl -X POST http://localhost:8080/park/B-1989-RFU/White
curl -X GET http://localhost:8080/cars_registration_numbers/colour/Black
curl -X GET http://localhost:8080/cars_slot/colour/Black
curl -X GET http://localhost:8080/slot_number/car_registration_number/B-1701-RFL
curl -X GET http://localhost:8080/slot_number/car_registration_number/RI-1
