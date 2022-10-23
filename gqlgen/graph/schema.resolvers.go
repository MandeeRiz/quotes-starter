package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gqlgen/graph/generated"
	"gqlgen/graph/model"
	"io"
	"net/http"
)

// CreateAQuote is the resolver for the createAQuote field.
func (r *mutationResolver) CreateAQuote(ctx context.Context, input model.QuoteInput) (*model.Quote, error) {
	quote := &model.Quote{
		ID:     "",
		Quote:  input.Quote,
		Author: input.Author,
	}
	quoteUnMarshalled, _ := json.Marshal(quote)
	bodyReturn := bytes.NewBuffer(quoteUnMarshalled)
	request, err := http.NewRequest("POST", "http://34.160.62.214:80/quotes", bodyReturn)
	ctxVal := ctx.Value("x-api-key")
	stringCtxVal := fmt.Sprintf("%v", ctxVal)
	request.Header.Set("x-api-key", stringCtxVal)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, _ := client.Do(request)
	if response.StatusCode == 401 || response.StatusCode == 400 {
		quote.ID = ""
		quote.Quote = ""
		quote.Author = ""
		err = errors.New(response.Status + " (insufficent characters)")
		return quote, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	json.Unmarshal(data, &quote)
	return quote, nil
}

// DeleteAquote is the resolver for the deleteAquote field.
func (r *mutationResolver) DeleteAquote(ctx context.Context, id string) (*string, error) {
	URL := "http://34.160.62.214:80/quotes/" + id
	request, err := http.NewRequest("DELETE", URL, nil)
	ctxVal := ctx.Value("x-api-key")
	stringCtxVal := fmt.Sprintf("%v", ctxVal)
	request.Header.Set("x-api-key", stringCtxVal)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, _ := client.Do(request)
	if response.StatusCode == 401 {
		status := "401: Unauthorized (incorrect headers)"
		return &status, err
	}
	if response.StatusCode == 404 {
		status := "404: Not Found (id not found)"
		return &status, err
	}
	if response.StatusCode == 204 {
		status := "204: No Content (delete successful)"
		return &status, err
	}
	dataArray := response.Status
	return &dataArray, err
}

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	request, err := http.NewRequest("GET", "http://34.160.62.214:80/quotes", nil)
	ctxVal := ctx.Value("x-api-key")
	stringCtxVal := fmt.Sprintf("%v", ctxVal)
	request.Header.Set("x-api-key", stringCtxVal)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, _ := client.Do(request)
	var quote model.Quote
	if response.StatusCode == 401 {
		err = errors.New(response.Status)
		return &quote, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	// var quote model.Quote
	json.Unmarshal(data, &quote)
	return &quote, err
}

// QuoteByID is the resolver for the quoteByID field.
func (r *queryResolver) QuoteByID(ctx context.Context, id string) (*model.Quote, error) {
	URL := "http://34.160.62.214:80/quotes/" + id
	request, err := http.NewRequest("GET", URL, nil)
	ctxVal := ctx.Value("x-api-key")
	stringCtxVal := fmt.Sprintf("%v", ctxVal)
	request.Header.Set("x-api-key", stringCtxVal)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, _ := client.Do(request)
	var quote model.Quote
	if response.StatusCode == 401 {
		err = errors.New(response.Status)
		return &quote, err
	}
	data, err := io.ReadAll(response.Body)
	json.Unmarshal(data, &quote)
	if response.StatusCode == 404 {
		err = errors.New(response.Status + " (id not found)")
		return &quote, err
	}
	return &quote, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
