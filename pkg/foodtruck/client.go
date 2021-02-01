// Package foodtruck is responsible creating client object and using soda library to talk to
// the REST API and process response.
package foodtruck

import (
	"encoding/json"
	soda "github.com/SebastiaanKlippert/go-soda"
	"io/ioutil"
	"log"
)

const (
	DefaultBaseURL   = "https://data.sfgov.org/resource/jjew-r69b"
	DefaultPageLimit = 10
	DefaultSorted    = true
	DefaultFormat    = "json"
	DefaultToken     = ""
)

type Bytes []byte
type handlerFunc func(trucks *FoodTrucks)

// Struct client defines the client used to process request by using soda library.
type client struct {
	baseURL string
	limit   uint
	sorted  bool
	token   string
	format  string
}

// NewClient is the constructor for struct client.
func NewClient(baseURL string, limit uint, sorted bool, format string, token string) *client {
	return &client{
		baseURL: baseURL,
		limit:   limit,
		sorted:  sorted,
		format:  format,
		token:   token,
	}
}

// GetFoodtrucksPaginated creates soda.OffsetGetRequest from requestBuilder and calls
// REST API with offset.
// GetFoodtrucksPaginated also takes a handlerFunc function to handle returned
// trucks and how to use trucks (For this instance, by printing it to console).
func (c *client) GetFoodtrucksPaginated(reqBuilder requestBuilder, handle handlerFunc) {
	offsetGetRequest := reqBuilder.wrapWithOffsetGetRequest()
	for !offsetGetRequest.IsDone() {
		offsetGetRequest.Add(1)
		resp, err := offsetGetRequest.Next(c.limit)
		if err == soda.ErrDone {
			log.Println("All trucks are printed")
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var trucks FoodTrucks
		if err := json.Unmarshal(body, &trucks); err != nil {
			log.Fatal(err)
		}

		handle(&trucks)
		offsetGetRequest.Done()
		resp.Body.Close()
	}
}

// Getter functions for client
func (c *client) BaseURL() string {
	return c.baseURL
}

func (c *client) Token() string {
	return c.token
}

func (c *client) Format() string {
	return c.format
}

func (c *client) Sorted() bool {
	return c.sorted
}