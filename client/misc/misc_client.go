// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new misc API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for misc API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	QueryBlockTimes(params *QueryBlockTimesParams) (*QueryBlockTimesOK, error)

	QueryDonation(params *QueryDonationParams) (*QueryDonationOK, error)

	QueryLastBlock(params *QueryLastBlockParams) (*QueryLastBlockOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  QueryBlockTimes queries block time

  Query the block time up to given height
*/
func (a *Client) QueryBlockTimes(params *QueryBlockTimesParams) (*QueryBlockTimesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryBlockTimesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryBlockTimes",
		Method:             "GET",
		PathPattern:        "/misc/block-times",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QueryBlockTimesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryBlockTimesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryBlockTimes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryDonation queries donations info

  Query donations by DonateToCommunityPool or CommentToken
*/
func (a *Client) QueryDonation(params *QueryDonationParams) (*QueryDonationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDonationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryDonation",
		Method:             "GET",
		PathPattern:        "/misc/donations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QueryDonationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDonationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryDonation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryLastBlock queries least block info

  Query least block info
*/
func (a *Client) QueryLastBlock(params *QueryLastBlockParams) (*QueryLastBlockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryLastBlockParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryLastBlock",
		Method:             "GET",
		PathPattern:        "/misc/height",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QueryLastBlockReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryLastBlockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryLastBlock: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}