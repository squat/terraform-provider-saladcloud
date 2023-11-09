// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateRecipeDeploymentRequest struct {
	UpdateRecipeDeployment shared.UpdateRecipeDeployment `request:"mediaType=application/json"`
	// The unique organization name
	OrganizationName string `pathParam:"style=simple,explode=false,name=organization_name"`
	// The unique project name
	ProjectName string `pathParam:"style=simple,explode=false,name=project_name"`
	// The unique recipe deployment name
	RecipeDeploymentName string `pathParam:"style=simple,explode=false,name=recipe_deployment_name"`
}

func (o *UpdateRecipeDeploymentRequest) GetUpdateRecipeDeployment() shared.UpdateRecipeDeployment {
	if o == nil {
		return shared.UpdateRecipeDeployment{}
	}
	return o.UpdateRecipeDeployment
}

func (o *UpdateRecipeDeploymentRequest) GetOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.OrganizationName
}

func (o *UpdateRecipeDeploymentRequest) GetProjectName() string {
	if o == nil {
		return ""
	}
	return o.ProjectName
}

func (o *UpdateRecipeDeploymentRequest) GetRecipeDeploymentName() string {
	if o == nil {
		return ""
	}
	return o.RecipeDeploymentName
}

type UpdateRecipeDeploymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Bad Request
	ProblemDetails *shared.ProblemDetails
	// OK
	RecipeDeployment *shared.RecipeDeployment
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateRecipeDeploymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRecipeDeploymentResponse) GetProblemDetails() *shared.ProblemDetails {
	if o == nil {
		return nil
	}
	return o.ProblemDetails
}

func (o *UpdateRecipeDeploymentResponse) GetRecipeDeployment() *shared.RecipeDeployment {
	if o == nil {
		return nil
	}
	return o.RecipeDeployment
}

func (o *UpdateRecipeDeploymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRecipeDeploymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
