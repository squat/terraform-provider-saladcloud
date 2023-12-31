// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListRecipeDeploymentsRequest struct {
	// The unique organization name
	OrganizationName string `pathParam:"style=simple,explode=false,name=organization_name"`
	// The unique project name
	ProjectName string `pathParam:"style=simple,explode=false,name=project_name"`
}

func (o *ListRecipeDeploymentsRequest) GetOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.OrganizationName
}

func (o *ListRecipeDeploymentsRequest) GetProjectName() string {
	if o == nil {
		return ""
	}
	return o.ProjectName
}

type ListRecipeDeploymentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Not Found
	ProblemDetails *shared.ProblemDetails
	// OK
	RecipeDeploymentList *shared.RecipeDeploymentList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListRecipeDeploymentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRecipeDeploymentsResponse) GetProblemDetails() *shared.ProblemDetails {
	if o == nil {
		return nil
	}
	return o.ProblemDetails
}

func (o *ListRecipeDeploymentsResponse) GetRecipeDeploymentList() *shared.RecipeDeploymentList {
	if o == nil {
		return nil
	}
	return o.RecipeDeploymentList
}

func (o *ListRecipeDeploymentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRecipeDeploymentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
