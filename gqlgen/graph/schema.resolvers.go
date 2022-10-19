package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
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
		Author: input.Quote,
	}
	//json marshall (byte array)
	quoteUnMarshalled, _ := json.Marshal(quote)
	//bytes.newbuffer
	bodyReturn := bytes.NewBuffer(quoteUnMarshalled)
	request, err := http.NewRequest("POST", "http://34.160.62.214:80/quotes", bodyReturn)
	//send the headers for verification
	ctxVal := ctx.Value("x-api-key")
	stringCtxVal := fmt.Sprintf("%v", ctxVal)
	request.Header.Set("x-api-key", stringCtxVal)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, _ := client.Do(request)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	json.Unmarshal(data, &quote)
	return quote, nil
}

// DeleteAquote is the resolver for the deleteAquote field.
func (r *mutationResolver) DeleteAquote(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented: DeleteAquote - deleteAquote"))
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
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var quote model.Quote
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
	data, err := io.ReadAll(response.Body)
	var quote model.Quote
	json.Unmarshal(data, &quote)
	return &quote, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
