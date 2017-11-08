// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new routes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for routes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAppsAppRoutesRoute deletes the route

Deletes the route.
*/
func (a *Client) DeleteAppsAppRoutesRoute(params *DeleteAppsAppRoutesRouteParams) (*DeleteAppsAppRoutesRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppsAppRoutesRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAppsAppRoutesRoute",
		Method:             "DELETE",
		PathPattern:        "/apps/{app}/routes/{route}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAppsAppRoutesRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAppsAppRoutesRouteOK), nil

}

/*
GetAppsAppRoutes gets route list by app name

This will list routes for a particular app.
*/
func (a *Client) GetAppsAppRoutes(params *GetAppsAppRoutesParams) (*GetAppsAppRoutesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsAppRoutesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAppsAppRoutes",
		Method:             "GET",
		PathPattern:        "/apps/{app}/routes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppsAppRoutesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppsAppRoutesOK), nil

}

/*
GetAppsAppRoutesRoute gets route by name

Gets a route by name.
*/
func (a *Client) GetAppsAppRoutesRoute(params *GetAppsAppRoutesRouteParams) (*GetAppsAppRoutesRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsAppRoutesRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAppsAppRoutesRoute",
		Method:             "GET",
		PathPattern:        "/apps/{app}/routes/{route}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppsAppRoutesRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppsAppRoutesRouteOK), nil

}

/*
PatchAppsAppRoutesRoute updates a route

Update a route
*/
func (a *Client) PatchAppsAppRoutesRoute(params *PatchAppsAppRoutesRouteParams) (*PatchAppsAppRoutesRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAppsAppRoutesRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAppsAppRoutesRoute",
		Method:             "PATCH",
		PathPattern:        "/apps/{app}/routes/{route}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchAppsAppRoutesRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchAppsAppRoutesRouteOK), nil

}

/*
PostAppsAppRoutes creates new route

Create a new route in an app, if app doesn't exists, it creates the app
*/
func (a *Client) PostAppsAppRoutes(params *PostAppsAppRoutesParams) (*PostAppsAppRoutesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAppsAppRoutesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAppsAppRoutes",
		Method:             "POST",
		PathPattern:        "/apps/{app}/routes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostAppsAppRoutesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAppsAppRoutesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
