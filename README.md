# Parking Lot Quiz with Go
I own a parking lot that can hold up to 'n' cars at any given point in time.
Each slot is given a number starting at 1 increasing with the increasing distance from the entry point in steps of one.
I want to create an automated ticketing system in the cloud that allows my customers to use my parking lot without human intervention.
When a car enters my parking lot, I want to have a ticket issued to the driver.
The ticket issuing process includes us documenting the registration number (number plate) and the colour of the car and allocating an available parking slot to the car before actually handing over a ticket to the driver (we assume that our customers are nice enough to always park in the slots allocated to them).
The customer should be allocated to a parking slot which is nearest to the entry.
At the exit the customer returns the ticket which then marks the slot they were using as being available.

Due to government regulation, the system should provide me with the ability to find out:
- Registration numbers of all cars of a particular colour.
- Slot number of a car with a given registration number is parked.
- Slot numbers of all slots where cars of a particular colour are parked.

We interact with the system via a set of endpoints which produce a specific output.
Please take a look at the example below, which includes all the endpoints you need to support - they're selfexplanatory.
The system should allow input in two ways. Just to clarify, the same codebase should support both modes of input - we don't want two distinct submissions.
1. It should provide us with several HTTP APIs as the commands.
2. It should accept a POST HTTP Request that accepts plain text payload which contains all the commands and reads the commands from the payload.

## Example:

### HTTP API Request

Assuming a parking lot with 6 slots, the following commands should be run in sequence by
typing them in a tool like cURL and should produce output as described below the command.

```bash
$ curl -X POST localhost:8080/create_parking_lot/6
Created a parking lot with 6 slots
```

```bash
$ curl -X POST http://localhost:8080/park/B-1234-RFS/Black
Allocated slot number: 1
```

```bash
$ curl -X POST http://localhost:8080/park/B-1999-RFD/Green
Allocated slot number: 2
```

```bash
$ curl -X POST http://localhost:8080/park/B-1000-RFS/Black
Allocated slot number: 3
```

```bash
$ curl -X POST http://localhost:8080/park/B-1777-RFU/BlueSky
Allocated slot number: 4
```

```bash
$ curl -X POST http://localhost:8080/park/B-1701-RFL/Blue
Allocated slot number: 5
```

```bash
$ curl -X POST http://localhost:8080/park/B-1141-RFS/Black
Allocated slot number: 6
```

```bash
$ curl -X POST http://localhost:8080/leave/4
Slot number 4 is free
```

```bash
$ curl -X GET http://localhost:8080/status
Slot No. Registration No Colour
1 B-1234-RFS Black
2 B-1999-RFD Green
3 B-1000-RFS Black
5 B-1701-RFL Blue
6 B-1141-RFS Black
```

```bash
$ curl -X POST http://localhost:8080/park/B-1333-RFS/Black
Allocated slot number: 4
```

```bash
$ curl -X POST http://localhost:8080/park/B-1989-RFU/White
Sorry, parking lot is full
```

```bash
$ curl -X GET http://localhost:8080/cars_registration_numbers/colour/Black
B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS
```

```bash
$ curl -X GET http://localhost:8080/cars_slot/colour/Black
1, 3, 4, 6
```

```bash
$ curl -X GET http://localhost:8080/slot_number/car_registration_number/B-1701-RFL
5
```

```bash
$ curl -X GET http://localhost:8080/slot_number/car_registration_number/RI-1
Not found
```

### HTTP API Request with Plain Text Payload

Path:
```bash
$ curl -X POST http://localhost:8080/bulk
```

Request body:
```bash
create_parking_lot 6
park B-1234-RFS Black
park B-1999-RFD Green
park B-1000-RFS Black
park B-1777-RFU BlueSky
park B-1701-RFL Blue
park B-1141-RFS Black
leave 4
status
park B-1333-RFS Black
park B-1989-RFU BlueSky
registration_numbers_for_cars_with_colour Black
slot_numbers_for_cars_with_colour Black
slot_number_for_registration_number B-1701-RFL
slot_number_for_registration_number RI-1
```

Expected response:
```bash
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No. Registration No Colour
1 B-1234-RFS Black
2 B-1999-RFD Green
3 B-1000-RFS Black
5 B-1701-RFL Blue
6 B-1141-RFS Black
Allocated slot number: 4
Sorry, parking lot is full
B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS
1, 3, 4, 6
5
Not found
```

## How to Run

Make sure Go >= 1.15.3 already installed and clone this repo at GOPATH

### Functional testing

1. Run the main service:
```bash
go run .
```

2. Run the functional testing:
```bash
go run test/functional/functional_testing.go -url=http://localhost:8080 -case=1
```

```bash
go run test/functional/functional_testing.go -url=http://localhost:8080 -case=2
```

### API testing via shell script
1. Set script to executed allowed
```bash
sudo chmod +x test.sh
```

2. Run script test
```bash
./test.sh
```

3. Output
```bash
Created a parking lot with 6 slot(s)
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No. Registration No Colour
1 B-1234-RFS Black
2 B-1999-RFD Green
3 B-1000-RFS Black
5 B-1701-RFL Blue
6 B-1141-RFS Black
Allocated slot number: 4
Sorry, parking lot is full
B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS
1, 3, 4, 6
5
Not found
```
