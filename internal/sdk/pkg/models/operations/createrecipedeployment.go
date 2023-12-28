// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateRecipeDeploymentRequest struct {
	CreateRecipeDeployment shared.CreateRecipeDeployment `request:"mediaType=application/json"`
	// The unique organization name
	OrganizationName string `pathParam:"style=simple,explode=false,name=organization_name"`
	// The unique project name
	ProjectName string `pathParam:"style=simple,explode=false,name=project_name"`
}

func (o *CreateRecipeDeploymentRequest) GetCreateRecipeDeployment() shared.CreateRecipeDeployment {
	if o == nil {
		return shared.CreateRecipeDeployment{}
	}
	return o.CreateRecipeDeployment
}

func (o *CreateRecipeDeploymentRequest) GetOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.OrganizationName
}

func (o *CreateRecipeDeploymentRequest) GetProjectName() string {
	if o == nil {
		return ""
	}
	return o.ProjectName
}

type CreateRecipeDeploymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// Bad Request
	ProblemDetails *shared.ProblemDetails
	// Created
	RecipeDeployment *shared.RecipeDeployment
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateRecipeDeploymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRecipeDeploymentResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateRecipeDeploymentResponse) GetProblemDetails() *shared.ProblemDetails {
	if o == nil {
		return nil
	}
	return o.ProblemDetails
}

func (o *CreateRecipeDeploymentResponse) GetRecipeDeployment() *shared.RecipeDeployment {
	if o == nil {
		return nil
	}
	return o.RecipeDeployment
}

func (o *CreateRecipeDeploymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRecipeDeploymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
