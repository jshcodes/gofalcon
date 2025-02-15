// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new device control policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for device control policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeviceControlPolicies(params *CreateDeviceControlPoliciesParams, opts ...ClientOption) (*CreateDeviceControlPoliciesCreated, error)

	DeleteDeviceControlPolicies(params *DeleteDeviceControlPoliciesParams, opts ...ClientOption) (*DeleteDeviceControlPoliciesOK, error)

	GetDefaultDeviceControlPolicies(params *GetDefaultDeviceControlPoliciesParams, opts ...ClientOption) (*GetDefaultDeviceControlPoliciesOK, error)

	GetDeviceControlPolicies(params *GetDeviceControlPoliciesParams, opts ...ClientOption) (*GetDeviceControlPoliciesOK, error)

	PerformDeviceControlPoliciesAction(params *PerformDeviceControlPoliciesActionParams, opts ...ClientOption) (*PerformDeviceControlPoliciesActionOK, error)

	QueryCombinedDeviceControlPolicies(params *QueryCombinedDeviceControlPoliciesParams, opts ...ClientOption) (*QueryCombinedDeviceControlPoliciesOK, error)

	QueryCombinedDeviceControlPolicyMembers(params *QueryCombinedDeviceControlPolicyMembersParams, opts ...ClientOption) (*QueryCombinedDeviceControlPolicyMembersOK, error)

	QueryDeviceControlPolicies(params *QueryDeviceControlPoliciesParams, opts ...ClientOption) (*QueryDeviceControlPoliciesOK, error)

	QueryDeviceControlPolicyMembers(params *QueryDeviceControlPolicyMembersParams, opts ...ClientOption) (*QueryDeviceControlPolicyMembersOK, error)

	SetDeviceControlPoliciesPrecedence(params *SetDeviceControlPoliciesPrecedenceParams, opts ...ClientOption) (*SetDeviceControlPoliciesPrecedenceOK, error)

	UpdateDefaultDeviceControlPolicies(params *UpdateDefaultDeviceControlPoliciesParams, opts ...ClientOption) (*UpdateDefaultDeviceControlPoliciesOK, error)

	UpdateDeviceControlPolicies(params *UpdateDeviceControlPoliciesParams, opts ...ClientOption) (*UpdateDeviceControlPoliciesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDeviceControlPolicies creates device control policies by specifying details about the policy to create
*/
func (a *Client) CreateDeviceControlPolicies(params *CreateDeviceControlPoliciesParams, opts ...ClientOption) (*CreateDeviceControlPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeviceControlPolicies",
		Method:             "POST",
		PathPattern:        "/policy/entities/device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeviceControlPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeviceControlPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDeviceControlPolicies deletes a set of device control policies by specifying their i ds
*/
func (a *Client) DeleteDeviceControlPolicies(params *DeleteDeviceControlPoliciesParams, opts ...ClientOption) (*DeleteDeviceControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDeviceControlPolicies",
		Method:             "DELETE",
		PathPattern:        "/policy/entities/device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeviceControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDeviceControlPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDefaultDeviceControlPolicies retrieves the configuration for a default device control policy
*/
func (a *Client) GetDefaultDeviceControlPolicies(params *GetDefaultDeviceControlPoliciesParams, opts ...ClientOption) (*GetDefaultDeviceControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDefaultDeviceControlPolicies",
		Method:             "GET",
		PathPattern:        "/policy/entities/default-device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefaultDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDefaultDeviceControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDefaultDeviceControlPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceControlPolicies retrieves a set of device control policies by specifying their i ds
*/
func (a *Client) GetDeviceControlPolicies(params *GetDeviceControlPoliciesParams, opts ...ClientOption) (*GetDeviceControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceControlPolicies",
		Method:             "GET",
		PathPattern:        "/policy/entities/device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceControlPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PerformDeviceControlPoliciesAction performs the specified action on the device control policies specified in the request
*/
func (a *Client) PerformDeviceControlPoliciesAction(params *PerformDeviceControlPoliciesActionParams, opts ...ClientOption) (*PerformDeviceControlPoliciesActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformDeviceControlPoliciesActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "performDeviceControlPoliciesAction",
		Method:             "POST",
		PathPattern:        "/policy/entities/device-control-actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformDeviceControlPoliciesActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PerformDeviceControlPoliciesActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PerformDeviceControlPoliciesActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedDeviceControlPolicies searches for device control policies in your environment by providing an f q l filter and paging details returns a set of device control policies which match the filter criteria
*/
func (a *Client) QueryCombinedDeviceControlPolicies(params *QueryCombinedDeviceControlPoliciesParams, opts ...ClientOption) (*QueryCombinedDeviceControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedDeviceControlPolicies",
		Method:             "GET",
		PathPattern:        "/policy/combined/device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryCombinedDeviceControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedDeviceControlPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedDeviceControlPolicyMembers searches for members of a device control policy in your environment by providing an f q l filter and paging details returns a set of host details which match the filter criteria
*/
func (a *Client) QueryCombinedDeviceControlPolicyMembers(params *QueryCombinedDeviceControlPolicyMembersParams, opts ...ClientOption) (*QueryCombinedDeviceControlPolicyMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedDeviceControlPolicyMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedDeviceControlPolicyMembers",
		Method:             "GET",
		PathPattern:        "/policy/combined/device-control-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedDeviceControlPolicyMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryCombinedDeviceControlPolicyMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedDeviceControlPolicyMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDeviceControlPolicies searches for device control policies in your environment by providing an f q l filter and paging details returns a set of device control policy i ds which match the filter criteria
*/
func (a *Client) QueryDeviceControlPolicies(params *QueryDeviceControlPoliciesParams, opts ...ClientOption) (*QueryDeviceControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryDeviceControlPolicies",
		Method:             "GET",
		PathPattern:        "/policy/queries/device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDeviceControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDeviceControlPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDeviceControlPolicyMembers searches for members of a device control policy in your environment by providing an f q l filter and paging details returns a set of agent i ds which match the filter criteria
*/
func (a *Client) QueryDeviceControlPolicyMembers(params *QueryDeviceControlPolicyMembersParams, opts ...ClientOption) (*QueryDeviceControlPolicyMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDeviceControlPolicyMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryDeviceControlPolicyMembers",
		Method:             "GET",
		PathPattern:        "/policy/queries/device-control-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDeviceControlPolicyMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDeviceControlPolicyMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDeviceControlPolicyMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SetDeviceControlPoliciesPrecedence sets the precedence of device control policies based on the order of i ds specified in the request the first ID specified will have the highest precedence and the last ID specified will have the lowest you must specify all non default policies for a platform when updating precedence
*/
func (a *Client) SetDeviceControlPoliciesPrecedence(params *SetDeviceControlPoliciesPrecedenceParams, opts ...ClientOption) (*SetDeviceControlPoliciesPrecedenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDeviceControlPoliciesPrecedenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setDeviceControlPoliciesPrecedence",
		Method:             "POST",
		PathPattern:        "/policy/entities/device-control-precedence/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetDeviceControlPoliciesPrecedenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetDeviceControlPoliciesPrecedenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetDeviceControlPoliciesPrecedenceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateDefaultDeviceControlPolicies updates the configuration for a default device control policy
*/
func (a *Client) UpdateDefaultDeviceControlPolicies(params *UpdateDefaultDeviceControlPoliciesParams, opts ...ClientOption) (*UpdateDefaultDeviceControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDefaultDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDefaultDeviceControlPolicies",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/default-device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDefaultDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDefaultDeviceControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDefaultDeviceControlPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDeviceControlPolicies updates device control policies by specifying the ID of the policy and details to update
*/
func (a *Client) UpdateDeviceControlPolicies(params *UpdateDeviceControlPoliciesParams, opts ...ClientOption) (*UpdateDeviceControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceControlPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDeviceControlPolicies",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/device-control/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeviceControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateDeviceControlPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
