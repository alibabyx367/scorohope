package ccv2

import (
	"bytes"
	"encoding/json"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/internal"
)

// ServiceBinding represents a Cloud Controller Service Binding.
type ServiceBinding struct {
	GUID string
}

// UnmarshalJSON helps unmarshal a Cloud Controller Service Binding response.
func (serviceBinding *ServiceBinding) UnmarshalJSON(data []byte) error {
	var ccServiceBinding struct {
		Metadata internal.Metadata
	}
	err := json.Unmarshal(data, &ccServiceBinding)
	if err != nil {
		return err
	}

	serviceBinding.GUID = ccServiceBinding.Metadata.GUID
	return nil
}

// serviceBindingRequestBody represents the body of the service binding create
// request.
type serviceBindingRequestBody struct {
	ServiceInstanceGUID string                 `json:"service_instance_guid"`
	AppGUID             string                 `json:"app_guid"`
	Parameters          map[string]interface{} `json:"parameters"`
}

// CreateServiceBinding creates a service binding
func (client *Client) CreateServiceBinding(appGUID string, serviceInstanceGUID string, parameters map[string]interface{}) (ServiceBinding, Warnings, error) {
	requestBody := serviceBindingRequestBody{
		ServiceInstanceGUID: serviceInstanceGUID,
		AppGUID:             appGUID,
		Parameters:          parameters,
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return ServiceBinding{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostServiceBindingRequest,
		Body:        bytes.NewReader(bodyBytes),
	})
	if err != nil {
		return ServiceBinding{}, nil, err
	}

	var serviceBinding ServiceBinding
	response := cloudcontroller.Response{
		Result: &serviceBinding,
	}

	err = client.connection.Make(request, &response)
	if err != nil {
		return ServiceBinding{}, response.Warnings, err
	}

	return serviceBinding, response.Warnings, nil
}

// GetServiceBindings returns back a list of Service Bindings based off of the
// provided queries.
func (client *Client) GetServiceBindings(queries []Query) ([]ServiceBinding, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceBindingsRequest,
		Query:       FormatQueryParameters(queries),
	})
	if err != nil {
		return nil, nil, err
	}

	var fullBindingsList []ServiceBinding
	warnings, err := client.paginate(request, ServiceBinding{}, func(item interface{}) error {
		if binding, ok := item.(ServiceBinding); ok {
			fullBindingsList = append(fullBindingsList, binding)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   ServiceBinding{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullBindingsList, warnings, err
}

// DeleteServiceBinding will destroy the requested Service Binding.
func (client *Client) DeleteServiceBinding(serviceBindingGUID string) (Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.DeleteServiceBindingRequest,
		URIParams:   map[string]string{"service_binding_guid": serviceBindingGUID},
	})
	if err != nil {
		return nil, err
	}

	var response cloudcontroller.Response
	err = client.connection.Make(request, &response)
	return response.Warnings, err
}
