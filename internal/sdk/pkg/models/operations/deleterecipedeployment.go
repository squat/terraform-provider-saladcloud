// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteRecipeDeploymentRequest struct {
	// The unique organization name
	OrganizationName string `pathParam:"style=simple,explode=false,name=organization_name"`
	// The unique project name
	ProjectName string `pathParam:"style=simple,explode=false,name=project_name"`
	// The unique recipe deployment name
	RecipeDeploymentName string `pathParam:"style=simple,explode=false,name=recipe_deployment_name"`
}

func (o *DeleteRecipeDeploymentRequest) GetOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.OrganizationName
}

func (o *DeleteRecipeDeploymentRequest) GetProjectName() string {
	if o == nil {
		return ""
	}
	return o.ProjectName
}

func (o *DeleteRecipeDeploymentRequest) GetRecipeDeploymentName() string {
	if o == nil {
		return ""
	}
	return o.RecipeDeploymentName
}

type DeleteRecipeDeploymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Forbidden
	ProblemDetails *shared.ProblemDetails
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteRecipeDeploymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteRecipeDeploymentResponse) GetProblemDetails() *shared.ProblemDetails {
	if o == nil {
		return nil
	}
	return o.ProblemDetails
}

func (o *DeleteRecipeDeploymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteRecipeDeploymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
