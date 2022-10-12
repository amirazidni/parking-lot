package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Case struct {
	Method          string
	Path            string
	ExpectedPayload string
}

func main() {

	url := flag.String("url", "http://localhost:8080", "targeted server url")
	testCase := flag.Int("case", 1, "which test case")

	flag.Parse()

	testSteps := testSteps1()
	if *testCase != 1 {
		testSteps = testSteps2()
	}

	for _, c := range testSteps {
		path := *url + c.Path

		if c.Method == http.MethodPost {
			post(path, c.ExpectedPayload)
			continue
		}

		get(path, c.ExpectedPayload)
		continue
	}
	log.Println("tests passed!")
}

func post(path, expectedPayload string) {
	req := []byte(bulkReq())
	resp, err := http.Post(path, "text/plain", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("http post: %s\n", err)
	}
	defer resp.Body.Close()

	handleResp(resp, expectedPayload)
}

func get(path, expectedPayload string) {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatalf("http post: %s\n", err)
	}
	defer resp.Body.Close()

	handleResp(resp, expectedPayload)
}

func handleResp(resp *http.Response, expectedPayload string) {
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code not OK, status code: %d\n", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read body: %s\n", err)
	}

	if string(body) != expectedPayload {
		fmt.Println("response: ", string(body))
		fmt.Println("expected: ", expectedPayload)
		log.Fatalf("test failed!\n")
	}
}

func testSteps1() []*Case {
	return []*Case{
		{
			Method:          http.MethodPost,
			Path:            "/create_parking_lot/6",
			ExpectedPayload: "Created a parking lot with 6 slot(s)\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1234-RFS/Black",
			ExpectedPayload: "Allocated slot number: 1\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1999-RFD/Green",
			ExpectedPayload: "Allocated slot number: 2\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1000-RFS/Black",
			ExpectedPayload: "Allocated slot number: 3\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1777-RFU/BlueSky",
			ExpectedPayload: "Allocated slot number: 4\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1701-RFL/Blue",
			ExpectedPayload: "Allocated slot number: 5\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1141-RFS/Black",
			ExpectedPayload: "Allocated slot number: 6\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/leave/4",
			ExpectedPayload: "Slot number 4 is free\n",
		},
		{
			Method:          http.MethodGet,
			Path:            "/status",
			ExpectedPayload: status(),
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1333-RFS/Black",
			ExpectedPayload: "Allocated slot number: 4\n",
		},
		{
			Method:          http.MethodPost,
			Path:            "/park/B-1989-RFU/White",
			ExpectedPayload: "Sorry, parking lot is full\n",
		},
		{
			Method:          http.MethodGet,
			Path:            "/cars_registration_numbers/colour/Black",
			ExpectedPayload: "B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS\n",
		},
		{
			Method:          http.MethodGet,
			Path:            "/cars_slot/colour/Black",
			ExpectedPayload: "1, 3, 4, 6\n",
		},
		{
			Method:          http.MethodGet,
			Path:            "/slot_number/car_registration_number/B-1701-RFL",
			ExpectedPayload: "5\n",
		},
		{
			Method:          http.MethodGet,
			Path:            "/slot_number/car_registration_number/RI-1",
			ExpectedPayload: "Not found\n",
		},
	}
}

func testSteps2() []*Case {
	return []*Case{
		{
			Method:          http.MethodPost,
			Path:            "/bulk",
			ExpectedPayload: bulkResp(),
		},
	}
}

func status() string {
	return `Slot No. Registration No Colour
1 B-1234-RFS Black
2 B-1999-RFD Green
3 B-1000-RFS Black
5 B-1701-RFL Blue
6 B-1141-RFS Black
`
}

func bulkReq() string {
	return `create_parking_lot 6
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
slot_number_for_registration_number RI-1`
}

func bulkResp() string {
	return `Created a parking lot with 6 slot(s)
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
`
}
