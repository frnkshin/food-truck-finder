package foodtruck

import (
	soda "github.com/SebastiaanKlippert/go-soda"
	"log"
)

type requestBuilder struct {
	getRequest       *soda.GetRequest
	offsetGetRequest *soda.OffsetGetRequest
}

func NewRequestBuilder(baseURL string, token string) *requestBuilder {
	return &requestBuilder{getRequest: soda.NewGetRequest(baseURL, token)}
}

func (b *requestBuilder) setFormat(format string) *requestBuilder {
	b.getRequest.Format = format
	return b
}

func (b *requestBuilder) setWhere(where string) *requestBuilder {
	b.getRequest.Query.Where = where
	return b
}

func (b *requestBuilder) setOrder(column string, order soda.Direction) *requestBuilder {
	b.getRequest.Query.AddOrder(column, order)
	return b
}

func (b *requestBuilder) wrapWithOffsetRequest() *requestBuilder {
	offsetRequest, err := soda.NewOffsetGetRequest(b.getRequest)
	if err != nil {
		log.Fatal(err)
	}
	b.offsetGetRequest = offsetRequest
	return b
}
