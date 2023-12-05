package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	//if err := godotenv.Load("../../.env"); err != nil {
	//	panic(err)
	//}
	// Define a structure for specifying input and output data of multi test cases.
	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		body          map[string]interface{}
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "logout test",
			route:         "/api/v1/logout",
			method:        "POST",
			body:          map[string]interface{}{},
			expectedError: false,
			expectedCode:  204,
		},
		{
			description:   "aws resource test",
			route:         fmt.Sprintf("/api/v1/aws-resource?region=%s&day=7&email=%s", os.Getenv("AWS_REGION"), os.Getenv("TEST_EMAIL")),
			method:        "GET",
			body:          map[string]interface{}{},
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "alert messages test",
			route:         fmt.Sprintf("/api/v1/alert-messages?email=%s", os.Getenv("TEST_EMAIL")),
			method:        "GET",
			body:          map[string]interface{}{},
			expectedError: false,
			expectedCode:  200,
		},
		{
			description: "user key save test",
			route:       "/api/v1/user-key",
			method:      "POST",
			body: map[string]interface{}{
				"email":     os.Getenv("TEST_EMAIL"),
				"accessKey": os.Getenv("TEST_AWS_ACCESS_KEY"),
				"secretKey": os.Getenv("TEST_AWS_SECRET_KEY"),
			},
			expectedError: false,
			expectedCode:  200,
		},
		{
			description: "alert setting save test",
			route:       "/api/v1/alert-setting",
			method:      "POST",
			body: map[string]interface{}{
				"email":      os.Getenv("TEST_EMAIL"),
				"timeEnd":    "2023-01-01 00:00:00",
				"targetCost": 1,
			},
			expectedError: false,
			expectedCode:  200,
		},
	}

	// Define a new Fiber app.
	app := fiber.New()

	// Define routes.
	PrivateRoutes(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		b, err := json.Marshal(test.body)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(test.method, test.route, bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency
		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
