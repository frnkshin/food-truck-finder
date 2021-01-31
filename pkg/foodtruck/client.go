package foodtruck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	soda "github.com/SebastiaanKlippert/go-soda"
)

const (
	DefaultBaseURL = "https://data.sfgov.org/resource/jjew-r69b.json"
	DefaultPageLimit = 10
	DefaultSorted = true
	DefaultTimeoutUnit = time.Second
)

type Bytes []byte

type client struct {
	client *http.Client
	config *config
}

type config struct {
	baseURL string
	limit uint
	sorted bool
	timeout time.Duration
}

func NewClient(config *config) *client {
	return &client{
		client:  &http.Client{
			Timeout: config.timeout,
		},
		config: config,
	}
}

func NewConfig(baseURL string, limit uint, sorted bool, timeoutDuration time.Duration) *config {
	return &config{
		baseURL: baseURL,
		limit:   limit,
		sorted:  sorted,
		timeout: timeoutDuration,
	}
}

func (c *client) ToString() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("baseURL: %s\n", c.config.baseURL))
	builder.WriteString(fmt.Sprintf("limit: %d\n", c.config.limit))
	builder.WriteString(fmt.Sprintf("sorted: %t\n", c.config.sorted))
	builder.WriteString(fmt.Sprintf("timeout %s", c.config.timeout))
	return builder.String()
}

func (c *client) fetch(url string) (Bytes, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *client) Get(query map[string]string) (FoodTrucks, error) {
	req, err := http.NewRequest("GET", c.config.baseURL, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	for key, element := range query {
		q.Add(key, element)
	}
	req.URL.RawQuery = q.Encode()
	if body, err := c.fetch(req.URL.String()); err != nil {
		return nil, err
	} else {
		return body.toJson(&FoodTrucks{})
	}
}

func (c *client) GetOpenFoodTrucks(query map[string]string) (FoodTrucks, error) {
	req, err := http.NewRequest("GET", c.getQuerySupportedURL(), nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("dayofweekstr", time.Now().Weekday().String())
	q.Add("start24", string(time.Now().Hour()))
	q.Add("end24", string(time.Now().Hour()))
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.RawQuery)

	if body, err := c.fetch(req.URL.String()); err != nil {
		return nil, err
	} else {
		return body.toJson(&FoodTrucks{})
	}
}

func (c *client) getQuerySupportedURL() string {
	return c.config.baseURL + "$where="
}

func (b Bytes) toJson(trucks *FoodTrucks) (FoodTrucks, error) {
	err := json.Unmarshal(b, &trucks)
	return *trucks, err
}