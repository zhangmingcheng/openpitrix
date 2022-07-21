// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new runtime manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for runtime manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDebugRuntime creates debug runtime
*/
func (a *Client) CreateDebugRuntime(params *CreateDebugRuntimeParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDebugRuntimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDebugRuntimeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDebugRuntime",
		Method:             "POST",
		PathPattern:        "/v1/debug_runtimes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateDebugRuntimeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDebugRuntimeOK), nil

}

/*
CreateDebugRuntimeCredential creates debug runtime credential
*/
func (a *Client) CreateDebugRuntimeCredential(params *CreateDebugRuntimeCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDebugRuntimeCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDebugRuntimeCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDebugRuntimeCredential",
		Method:             "POST",
		PathPattern:        "/v1/debug_runtimes/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateDebugRuntimeCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDebugRuntimeCredentialOK), nil

}

/*
CreateRuntime creates runtime
*/
func (a *Client) CreateRuntime(params *CreateRuntimeParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRuntimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRuntimeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRuntime",
		Method:             "POST",
		PathPattern:        "/v1/runtimes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRuntimeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRuntimeOK), nil

}

/*
CreateRuntimeCredential creates runtime credential
*/
func (a *Client) CreateRuntimeCredential(params *CreateRuntimeCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRuntimeCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRuntimeCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRuntimeCredential",
		Method:             "POST",
		PathPattern:        "/v1/runtimes/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRuntimeCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRuntimeCredentialOK), nil

}

/*
DeleteRuntimeCredentials batches delete runtime credentials
*/
func (a *Client) DeleteRuntimeCredentials(params *DeleteRuntimeCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRuntimeCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuntimeCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRuntimeCredentials",
		Method:             "DELETE",
		PathPattern:        "/v1/runtimes/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRuntimeCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRuntimeCredentialsOK), nil

}

/*
DeleteRuntimes deletes runtimes
*/
func (a *Client) DeleteRuntimes(params *DeleteRuntimesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRuntimesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuntimesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRuntimes",
		Method:             "DELETE",
		PathPattern:        "/v1/runtimes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRuntimesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRuntimesOK), nil

}

/*
DescribeDebugRuntimeCredentials gets debug runtime credentials filter with these fields runtime credential id status provider owner default return all debug runtime credentials
*/
func (a *Client) DescribeDebugRuntimeCredentials(params *DescribeDebugRuntimeCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeDebugRuntimeCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDebugRuntimeCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeDebugRuntimeCredentials",
		Method:             "GET",
		PathPattern:        "/v1/debug_runtimes/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeDebugRuntimeCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeDebugRuntimeCredentialsOK), nil

}

/*
DescribeDebugRuntimes gets debug runtimes can filter with these fields runtime id provider zone status owner default return all debug runtimes
*/
func (a *Client) DescribeDebugRuntimes(params *DescribeDebugRuntimesParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeDebugRuntimesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDebugRuntimesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeDebugRuntimes",
		Method:             "GET",
		PathPattern:        "/v1/debug_runtimes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeDebugRuntimesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeDebugRuntimesOK), nil

}

/*
DescribeRuntimeCredentials gets runtime credentials filter with these fields runtime credential id status provider owner default return all runtime credentials
*/
func (a *Client) DescribeRuntimeCredentials(params *DescribeRuntimeCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeRuntimeCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeRuntimeCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeRuntimeCredentials",
		Method:             "GET",
		PathPattern:        "/v1/runtimes/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeRuntimeCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeRuntimeCredentialsOK), nil

}

/*
DescribeRuntimeProviderZones gets runtime provider zones
*/
func (a *Client) DescribeRuntimeProviderZones(params *DescribeRuntimeProviderZonesParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeRuntimeProviderZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeRuntimeProviderZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeRuntimeProviderZones",
		Method:             "GET",
		PathPattern:        "/v1/runtimes/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeRuntimeProviderZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeRuntimeProviderZonesOK), nil

}

/*
DescribeRuntimes gets runtimes can filter with these fields runtime id provider zone status owner default return all runtimes
*/
func (a *Client) DescribeRuntimes(params *DescribeRuntimesParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeRuntimesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeRuntimesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeRuntimes",
		Method:             "GET",
		PathPattern:        "/v1/runtimes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeRuntimesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeRuntimesOK), nil

}

/*
GetRuntimeStatistics gets statistics of runtime
*/
func (a *Client) GetRuntimeStatistics(params *GetRuntimeStatisticsParams, authInfo runtime.ClientAuthInfoWriter) (*GetRuntimeStatisticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuntimeStatisticsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRuntimeStatistics",
		Method:             "GET",
		PathPattern:        "/v1/runtimes/statistics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRuntimeStatisticsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRuntimeStatisticsOK), nil

}

/*
ModifyRuntime modifies runtime
*/
func (a *Client) ModifyRuntime(params *ModifyRuntimeParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyRuntimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRuntimeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyRuntime",
		Method:             "PATCH",
		PathPattern:        "/v1/runtimes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyRuntimeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyRuntimeOK), nil

}

/*
ModifyRuntimeCredential modifies runtime credential
*/
func (a *Client) ModifyRuntimeCredential(params *ModifyRuntimeCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyRuntimeCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRuntimeCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyRuntimeCredential",
		Method:             "PATCH",
		PathPattern:        "/v1/runtimes/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyRuntimeCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyRuntimeCredentialOK), nil

}

/*
ValidateRuntimeCredential validates runtime credential
*/
func (a *Client) ValidateRuntimeCredential(params *ValidateRuntimeCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateRuntimeCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateRuntimeCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ValidateRuntimeCredential",
		Method:             "POST",
		PathPattern:        "/v1/runtimes/credentials:validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateRuntimeCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateRuntimeCredentialOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}