// Package foodtruck is responsible creating client object and using soda library to talk to
// the REST API and process response.
package foodtruck

import (
	soda "github.com/SebastiaanKlippert/go-soda"
	"log"
)

// Struct requestBuilder is used to create a soda.GetRequest and also creates
// soda.OffsetGetRequest to be supplied to soda library to use the request
// to send to the REST API endpoint.
type requestBuilder struct {
	getRequest       *soda.GetRequest
	offsetGetRequest *soda.OffsetGetRequest
}

func NewRequestBuilder(baseURL string, token string) *requestBuilder {
	return &requestBuilder{getRequest: soda.NewGetRequest(baseURL, token)}
}

// SetFormat sets format and returns the current requestBuilder.
func (b *requestBuilder) SetFormat(format string) *requestBuilder {
	b.getRequest.Format = format
	return b
}

// SetWhere sets where and returns the current requestBuilder.
func (b *requestBuilder) SetWhere(where string) *requestBuilder {
	b.getRequest.Query.Where = where
	return b
}

// SetOrder sets order and returns the current requestBuilder.
func (b *requestBuilder) SetOrder(column string, order soda.Direction) *requestBuilder {
	b.getRequest.Query.AddOrder(column, order)
	return b
}

// wrapWithOffsetGetRequest is a private method intended to be used by client to create and get
// soda.OffsetGetRequest from requestBuilder.getRequest.
func (b *requestBuilder) wrapWithOffsetGetRequest() *soda.OffsetGetRequest {
	offsetRequest, err := soda.NewOffsetGetRequest(b.getRequest)
	if err != nil {
		log.Fatal(err)
	}
	b.offsetGetRequest = offsetRequest
	return b.offsetGetRequest
}
