// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"net/http"
)

type ContainerGroupInstanceRecreateRequest struct {
	// The unique container group name
	ContainerGroupName string `pathParam:"style=simple,explode=false,name=container_group_name"`
	// The unique machine identifier
	MachineID string `pathParam:"style=simple,explode=false,name=machine_id"`
	// The unique organization name
	OrganizationName string `pathParam:"style=simple,explode=false,name=organization_name"`
	// The unique project name
	ProjectName string `pathParam:"style=simple,explode=false,name=project_name"`
}

func (o *ContainerGroupInstanceRecreateRequest) GetContainerGroupName() string {
	if o == nil {
		return ""
	}
	return o.ContainerGroupName
}

func (o *ContainerGroupInstanceRecreateRequest) GetMachineID() string {
	if o == nil {
		return ""
	}
	return o.MachineID
}

func (o *ContainerGroupInstanceRecreateRequest) GetOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.OrganizationName
}

func (o *ContainerGroupInstanceRecreateRequest) GetProjectName() string {
	if o == nil {
		return ""
	}
	return o.ProjectName
}

type ContainerGroupInstanceRecreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Not Found
	ProblemDetails *shared.ProblemDetails
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ContainerGroupInstanceRecreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ContainerGroupInstanceRecreateResponse) GetProblemDetails() *shared.ProblemDetails {
	if o == nil {
		return nil
	}
	return o.ProblemDetails
}

func (o *ContainerGroupInstanceRecreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ContainerGroupInstanceRecreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
