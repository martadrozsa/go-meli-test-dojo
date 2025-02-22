package api

import (
	"github.com/klasrak/go-meli-test-dojo/clients/swapi"
	"github.com/klasrak/go-meli-test-dojo/errors"
	"github.com/klasrak/go-meli-test-dojo/mockeable"
	"github.com/klasrak/go-meli-test-dojo/models"
	"net/http"
	"testing"
)

func TestGetStarshipsHandlerBadRequest(t *testing.T) {

	url := "/api/v1/starships/invalid_id"
	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 400

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}
}

func TestGetStarshipHandlerSuccess(t *testing.T) {

	url := "/api/v1/starships/9"

	mock := swapi.MockClient{
		GetStarshipFunc: func(id int) (models.Starship, error) {
			if id != 9 {
				t.Errorf("Assertion error. Expected: %d, Got: %d", 9, id)
			}

			return models.Starship{
				Name:                 "Death Star",
				Model:                "DS-1 Orbital Battle Station",
				Manufacturer:         "Imperial Department of Military Research, Sienar Fleet Systems",
				CostInCredits:        "1000000000000",
				Length:               "120000",
				MaxAtmospheringSpeed: "n/a",
				Crew:                 "342953",
				Passengers:           "843342",
				CargoCapacity:        "1000000000000",
				Consumables:          "3 years",
				HyperdriveRating:     "4.0",
				MGLT:                 "10",
				Class:                "Deep Space Mobile Battlestation",
				Films: []string{
					"https://swapi.dev/api/films/1/",
				},
			}, nil
		},
		GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 200
	expectedBody := `{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null}`

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}

	if response.StringBody() != expectedBody {
		t.Errorf("Assertion error. Expected: %s, Got: %s", expectedBody, response.StringBody())
	}
}

func TestGetStarshipHandlerNotFound(t *testing.T) {
	url := "/api/v1/starships/9"
	expectedError := 404

	mock := swapi.MockClient{
		GetStarshipFunc: func(id int) (models.Starship, error) {
			if id != 9 {
				t.Errorf("Assertion error. Expected: %d, Got: %d", 9, id)
			}

			return models.Starship{}, errors.NewNotFound("Not found", "starships not found")
		},
		GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetStarshipHandlerInternalServerError(t *testing.T) {
	url := "/api/v1/starships/9"
	expectedError := 500

	mock := swapi.MockClient{
		GetStarshipFunc: func(id int) (models.Starship, error) {
			if id != 9 {
				t.Errorf("Assertion error. Expected: %d, Got: %d", 9, id)
			}
			return models.Starship{}, errors.NewInternal()
		},
		GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetStarshipsHandlerInternalServerError(t *testing.T) {
	url := "/api/v1/starships"
	expectedError := 500

	mock := swapi.MockClient{
		GetStarshipsFunc: func() (models.Starships, error) {
			return models.Starships{}, errors.NewInternal()
		},
		GetStarshipsFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetStarshipsHandlerNotFound(t *testing.T) {
	url := "/api/v1/starships"
	expectedError := 404

	mock := swapi.MockClient{
		GetStarshipsFunc: func() (models.Starships, error) {
			return models.Starships{}, errors.NewNotFound("Not found", "starships not found")
		},
		GetStarshipsFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetStarshipsHandlerSuccess(t *testing.T) {
	url := "/api/v1/starships"

	mock := swapi.MockClient{
		GetStarshipsFunc: func() (models.Starships, error) {
			return models.Starships{
				Count: 2,
				Results: []models.Starship{
					{
						Name:                 "CR90 corvette",
						Model:                "CR90 corvette",
						Class:                "1 year",
						Manufacturer:         "corvette",
						CostInCredits:        "Corellian Engineering Corporation",
						Length:               "3500000",
						Crew:                 "30-165",
						Passengers:           "600",
						MaxAtmospheringSpeed: "150",
						HyperdriveRating:     "60",
						MGLT:                 "3000000",
						CargoCapacity:        "950",
						Consumables:          "2.0",
						Films: []string{
							"https://swapi.dev/api/films/1/",
							"https://swapi.dev/api/films/3/",
							"https://swapi.dev/api/films/6/"},
						Pilots: []string{},
					},
					{
						Name:                 "Star Destroyer",
						Model:                "Imperial I-class Star Destroyer",
						Class:                "Star Destroyer",
						Manufacturer:         "Kuat Drive Yards",
						CostInCredits:        "150000000",
						Length:               "1,600",
						Crew:                 "47,060",
						Passengers:           "n/a",
						MaxAtmospheringSpeed: "975",
						HyperdriveRating:     "2.0",
						MGLT:                 "60",
						CargoCapacity:        "36000000",
						Consumables:          "2 years",
						Films: []string{
							"https://swapi.dev/api/films/1/",
							"https://swapi.dev/api/films/2/",
							"https://swapi.dev/api/films/3/"},
						Pilots: []string{},
					},
				},
			}, nil
		},
		GetStarshipsFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 200
	expectedBody := `{"count":2,"results":[{"name":"CR90 corvette","model":"CR90 corvette","starship_class":"1 year","manufacturer":"corvette","cost_in_credits":"Corellian Engineering Corporation","length":"3500000","crew":"30-165","passengers":"600","max_atmosphering_speed":"150","hyperdrive_rating":"60","MGLT":"3000000","cargo_capacity":"950","consumables":"2.0","films":["https://swapi.dev/api/films/1/","https://swapi.dev/api/films/3/","https://swapi.dev/api/films/6/"],"pilots":[]},{"name":"Star Destroyer","model":"Imperial I-class Star Destroyer","starship_class":"Star Destroyer","manufacturer":"Kuat Drive Yards","cost_in_credits":"150000000","length":"1,600","crew":"47,060","passengers":"n/a","max_atmosphering_speed":"975","hyperdrive_rating":"2.0","MGLT":"60","cargo_capacity":"36000000","consumables":"2 years","films":["https://swapi.dev/api/films/1/","https://swapi.dev/api/films/2/","https://swapi.dev/api/films/3/"],"pilots":[]}]}`

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}

	if response.StringBody() != expectedBody {
		t.Errorf("Assertion error. Expected: %s, Got: %s", expectedBody, response.StringBody())
	}
}

func TestGetPeopleHandlerBadRequest(t *testing.T) {

	url := "/api/v1/people/XX"
	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 400

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}
}

func TestGetPeopleHandlerSuccess(t *testing.T) {

	url := "/api/v1/people/1"

	mock := swapi.MockClient{GetPeopleFunc: func(id int) (models.People, error) {
		if id != 1 {
			t.Errorf("Assertion error. Expected: %d, Got: %d", 1, id)
		}
		return models.People{
			Name:      "Luke Skywalker",
			BirthYear: "19BBY",
			EyeColor:  "blue",
			Gender:    "male",
			HairColor: "blond",
			Height:    "172",
			Mass:      "77",
			SkinColor: "fair",
			Homeworld: "https://swapi.dev/api/planets/1/",
			Films: []string{
				"https://swapi.dev/api/films/1/",
				"https://swapi.dev/api/films/2/",
				"https://swapi.dev/api/films/3/",
				"https://swapi.dev/api/films/6/"},
			Species: []string{},
			Starships: []string{
				"https://swapi.dev/api/starships/12/",
				"https://swapi.dev/api/starships/22/",
			},
		}, nil
	},
		GetPeopleFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 200
	expectedBody := `{"name":"Luke Skywalker","birth_year":"19BBY","eye_color":"blue","gender":"male","hair_color":"blond","height":"172","mass":"77","skin_color":"fair","homeworld":"https://swapi.dev/api/planets/1/","films":["https://swapi.dev/api/films/1/","https://swapi.dev/api/films/2/","https://swapi.dev/api/films/3/","https://swapi.dev/api/films/6/"],"species":[],"starships":["https://swapi.dev/api/starships/12/","https://swapi.dev/api/starships/22/"]}`

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}

	if response.StringBody() != expectedBody {
		t.Errorf("Assertion error. Expected: %s, Got: %s", expectedBody, response.StringBody())
	}
}

func TestGetPeopleHandlerNotFound(t *testing.T) {

	url := "/api/v1/people/1"
	expectedError := 404

	mock := swapi.MockClient{GetPeopleFunc: func(id int) (models.People, error) {
		if id != 1 {
			t.Errorf("Assertion error. Expected: %d, Got: %d", 1, id)
		}
		return models.People{}, errors.NewNotFound("Not found", "people not found")
	},
		GetPeopleFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetPeopleHandlerInternalServerError(t *testing.T) {

	url := "/api/v1/people/1"
	expectedError := 500

	mock := swapi.MockClient{GetPeopleFunc: func(id int) (models.People, error) {
		if id != 1 {
			t.Errorf("Assertion error. Expected: %d, Got: %d", 1, id)
		}
		return models.People{}, errors.NewInternal()
	},
		GetPeopleFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetPeopleListHandlerNotFound(t *testing.T) {
	url := "/api/v1/people"
	expectedError := 404

	mock := swapi.MockClient{
		GetPeopleListFunc: func() (models.PeopleList, error) {
			return models.PeopleList{}, errors.NewNotFound("Not found", "peoples not found")
		},
		GetPeopleListFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetPeopleListHandlerInternalServerError(t *testing.T) {
	url := "/api/v1/people"
	expectedError := 500

	mock := swapi.MockClient{
		GetPeopleListFunc: func() (models.PeopleList, error) {
			return models.PeopleList{}, errors.NewInternal()
		},
		GetPeopleListFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")

	if response.StatusCode != expectedError {
		t.Errorf("Assertion error. Expected: %d, Got: %d", expectedError, response.StatusCode)
	}
}

func TestGetPeopleListHandlerSuccess(t *testing.T) {
	url := "/api/v1/people"

	mock := swapi.MockClient{
		GetPeopleListFunc: func() (models.PeopleList, error) {
			return models.PeopleList{
				Count: 2,
				Results: []models.People{
					{
						Name:      "Luke Skywalker",
						BirthYear: "19BBY",
						EyeColor:  "blue",
						Gender:    "male",
						HairColor: "blond",
						Height:    "172",
						Mass:      "77",
						SkinColor: "fair",
						Homeworld: "https://swapi.dev/api/planets/1/",
						Films: []string{
							"https://swapi.dev/api/films/1/",
							"https://swapi.dev/api/films/2/",
							"https://swapi.dev/api/films/3/",
							"https://swapi.dev/api/films/6/",
						},
						Species: []string{},
						Starships: []string{
							"https://swapi.dev/api/starships/12/",
							"https://swapi.dev/api/starships/22/",
						},
					},
					{
						Name:      "C-3PO",
						BirthYear: "112BBY",
						EyeColor:  "yellow",
						Gender:    "n/a",
						HairColor: "n/a",
						Height:    "167",
						Mass:      "75",
						SkinColor: "gold",
						Homeworld: "https://swapi.dev/api/planets/1/",
						Films: []string{
							"https://swapi.dev/api/films/1/",
							"https://swapi.dev/api/films/2/",
							"https://swapi.dev/api/films/3/",
							"https://swapi.dev/api/films/4/",
							"https://swapi.dev/api/films/5/",
							"https://swapi.dev/api/films/6/",
						},
						Species: []string{
							"https://swapi.dev/api/species/2/",
						},
						Starships: []string{},
					},
				},
			}, nil
		},
		GetPeopleListFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 200
	expectedBody := `{"count":2,"results":[{"name":"Luke Skywalker","birth_year":"19BBY","eye_color":"blue","gender":"male","hair_color":"blond","height":"172","mass":"77","skin_color":"fair","homeworld":"https://swapi.dev/api/planets/1/","films":["https://swapi.dev/api/films/1/","https://swapi.dev/api/films/2/","https://swapi.dev/api/films/3/","https://swapi.dev/api/films/6/"],"species":[],"starships":["https://swapi.dev/api/starships/12/","https://swapi.dev/api/starships/22/"]},{"name":"C-3PO","birth_year":"112BBY","eye_color":"yellow","gender":"n/a","hair_color":"n/a","height":"167","mass":"75","skin_color":"gold","homeworld":"https://swapi.dev/api/planets/1/","films":["https://swapi.dev/api/films/1/","https://swapi.dev/api/films/2/","https://swapi.dev/api/films/3/","https://swapi.dev/api/films/4/","https://swapi.dev/api/films/5/","https://swapi.dev/api/films/6/"],"species":["https://swapi.dev/api/species/2/"],"starships":[]}]}`

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}

	if response.StringBody() != expectedBody {
		t.Errorf("Assertion error. Expected: %s, Got: %s", expectedBody, response.StringBody())
	}
}
