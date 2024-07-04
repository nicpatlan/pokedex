package pokedexAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
	Next     string
	Previous string
}

type PokeAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetAreas(url *string) Config {
	var requestURL string
	if url == nil {
		requestURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		requestURL = *url
	}
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	data := []byte(body)
	areas := PokeAreaResponse{}
	e := json.Unmarshal(data, &areas)
	if e != nil {
		log.Fatal(e)
	}

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	fmt.Println()

	return Config{
		Next:     areas.Next,
		Previous: areas.Previous,
	}
}
