package amirazidni.parkinglot;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.core.io.ClassPathResource;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.util.StreamUtils;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

import java.io.InputStream;
import java.nio.charset.StandardCharsets;

@SpringBootTest
@AutoConfigureMockMvc
class ParkinglotApplicationTests {

	@Autowired
	private MockMvc mockMvc;

	@Test
	void contextLoads() {
	}

	@Test
	void parkTest() throws Exception {

		// 1) CREATE
		mockMvc.perform(post("/create_parking_lot/6"))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals("Created a parking lot with 6 slots\n", response);
				});

		// 2) PARK
		String[] tests = {
				"B-1234-RFS", "Black", "Allocated slot number: 1\n",
				"B-1999-RFD", "Green", "Allocated slot number: 2\n",
				"B-1000-RFS", "Black", "Allocated slot number: 3\n",
				"B-1777-RFU", "BlueSky", "Allocated slot number: 4\n",
				"B-1701-RFL", "Blue", "Allocated slot number: 5\n",
				"B-1141-RFS", "Black", "Allocated slot number: 6\n",
		}; // value, attribute, expected

		for (int c = 0; c < tests.length; c += 3) {
			final int i = c;
			String url = String.format("/park/%s/%s", tests[i], tests[i + 1]);
			mockMvc.perform(post(url))
					.andExpectAll(status().isOk())
					.andDo(result -> {
						String response = result.getResponse().getContentAsString();
						assertEquals(tests[i + 2], response);
					});
		}

		// 3) LEAVE PARK
		mockMvc.perform(post("/leave/4"))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals("Slot number 4 is free\n", response);
				});

		// 4) STATUS CHECK
		final var expected = "Slot No. Registration No Colour\n" +
				"1 B-1234-RFS Black\n" +
				"2 B-1999-RFD Green\n" +
				"3 B-1000-RFS Black\n" +
				"5 B-1701-RFL Blue\n" +
				"6 B-1141-RFS Black\n";

		mockMvc.perform(get("/status"))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals(expected, response);
				});

		// 5) PARK 2
		tests[0] = "B-1333-RFS";
		tests[1] = "Black";
		tests[2] = "Allocated slot number: 4\n";
		String url = String.format("/park/%s/%s", tests[0], tests[1]);
		mockMvc.perform(post(url))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals(tests[2], response);
				});

		tests[3] = "B-1989-RFU";
		tests[4] = "White";
		tests[5] = "Sorry, parking lot is full\n";
		url = String.format("/park/%s/%s", tests[3], tests[4]);
		mockMvc.perform(post(url))
				.andExpectAll(status().isBadRequest())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals(tests[5], response);
				});

		// 6) GET PLATE NUMBERS BY COLOR
		mockMvc.perform(get("/cars_registration_numbers/colour/Black"))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals("B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS\n", response);
				});

		// 7) GET SLOT NUMBERS BY COLOR
		mockMvc.perform(get("/cars_slot/colour/Black"))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals("1, 3, 4, 6\n", response);
				});

		// 8) GET A SLOT NUMBER BY PLATE NUMBER
		mockMvc.perform(get("/slot_number/car_registration_number/B-1701-RFL"))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals("5\n", response);
				});

		mockMvc.perform(get("/slot_number/car_registration_number/RI-1"))
				.andExpectAll(status().isNotFound())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals("Not found\n", response);
				});

	}

	@Test
	void bulkTest() throws Exception {
		ClassPathResource bulkClassPathResource = new ClassPathResource("bulk.txt");
		byte[] body = bulkClassPathResource.getContentAsByteArray();

		ClassPathResource expectedPathResource = new ClassPathResource("expected-bulk.txt");
		InputStream inputStream = expectedPathResource.getInputStream();
		String expected = StreamUtils.copyToString(inputStream, StandardCharsets.UTF_8) + "\n";

		mockMvc.perform(post("/bulk")
				.contentType(MediaType.TEXT_PLAIN_VALUE)
				.content(body))
				.andExpectAll(status().isOk())
				.andDo(result -> {
					String response = result.getResponse().getContentAsString();
					assertEquals(expected, response);
				});
	}

}
