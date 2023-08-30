package amirazidni.parkinglot;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.web.servlet.MockMvc;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

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

	}
}
