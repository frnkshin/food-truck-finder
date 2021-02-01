package foodtruck

import (
	"encoding/json"
	"fmt"
	soda "github.com/SebastiaanKlippert/go-soda"
	"io/ioutil"
	"log"
	"time"
)

const (
	DefaultBaseURL   = "https://data.sfgov.org/resource/jjew-r69b"
	DefaultPageLimit = 10
	DefaultSorted    = true
	DefaultFormat    = "json"
	DefaultToken     = ""
	OpenNowStatement = `dayofweekstr='%s' AND start24<='%s' AND end24>'%s'`
)

type Bytes []byte
type callback func(trucks *FoodTrucks)
type client struct {
	baseURL string
	limit   uint
	sorted  bool
	token   string
	format  string
}

func NewClient(baseURL string, limit uint, sorted bool, format string, token string) *client {
	return &client{
		baseURL: baseURL,
		limit:   limit,
		sorted:  sorted,
		format:  format,
		token:   token,
	}
}

func (c *client) GetFoodtrucks(cb callback) {
	hhmm := getCurrentHHMM("%d:%d")
	requestBuilder := NewRequestBuilder(c.baseURL, c.token).setFormat(c.format).setOrder("applicant", soda.DirAsc).setWhere(fmt.Sprintf(OpenNowStatement, time.Now().Weekday(), hhmm, hhmm))
	resp, err := requestBuilder.getRequest.Get()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var trucks FoodTrucks
	if err := json.Unmarshal(body, &trucks); err != nil {
		log.Fatal(err)
	}
	cb(&trucks)
}

func (c *client) GetFoodtrucksPaginated(cb callback) {
	hhmm := getCurrentHHMM("%d:%d")
	requestBuilder := NewRequestBuilder(c.baseURL, c.token).setFormat(c.format).setOrder("applicant", soda.DirAsc).setWhere(fmt.Sprintf(OpenNowStatement, time.Now().Weekday(), hhmm, hhmm)).wrapWithOffsetRequest()
	offsetGetRequest := requestBuilder.offsetGetRequest
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
		cb(&trucks)

		offsetGetRequest.Done()
		resp.Body.Close()
	}
}
