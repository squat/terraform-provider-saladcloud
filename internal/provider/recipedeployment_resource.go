// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/squat/terraform-provider-saladcloud/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RecipeDeploymentResource{}
var _ resource.ResourceWithImportState = &RecipeDeploymentResource{}

func NewRecipeDeploymentResource() resource.Resource {
	return &RecipeDeploymentResource{}
}

// RecipeDeploymentResource defines the resource implementation.
type RecipeDeploymentResource struct {
	client *sdk.SDK
}

// RecipeDeploymentResourceModel describes the resource data model.
type RecipeDeploymentResourceModel struct {
	CurrentState     ContainerGroupState       `tfsdk:"current_state"`
	DisplayName      types.String              `tfsdk:"display_name"`
	ID               types.String              `tfsdk:"id"`
	Name             types.String              `tfsdk:"name"`
	Networking       *ContainerGroupNetworking `tfsdk:"networking"`
	OrganizationName types.String              `tfsdk:"organization_name"`
	ProjectName      types.String              `tfsdk:"project_name"`
	Recipe           Recipe                    `tfsdk:"recipe"`
	RecipeName       types.String              `tfsdk:"recipe_name"`
	Replicas         types.Int64               `tfsdk:"replicas"`
}

func (r *RecipeDeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_recipe_deployment"
}

func (r *RecipeDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RecipeDeployment Resource",

		Attributes: map[string]schema.Attribute{
			"current_state": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Computed: true,
					},
					"finish_time": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"instance_status_count": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"allocating_count": schema.Int64Attribute{
								Computed: true,
							},
							"creating_count": schema.Int64Attribute{
								Computed: true,
							},
							"running_count": schema.Int64Attribute{
								Computed: true,
							},
						},
						Description: `Represents a container group instance status count`,
					},
					"start_time": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"status": schema.StringAttribute{
						Computed:    true,
						Description: `must be one of ["pending", "running", "stopped", "succeeded", "failed", "deploying"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"pending",
								"running",
								"stopped",
								"succeeded",
								"failed",
								"deploying",
							),
						},
					},
				},
				Description: `Represents a container group state`,
			},
			"display_name": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"networking": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"auth": schema.BoolAttribute{
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"dns": schema.StringAttribute{
						Computed: true,
					},
					"port": schema.Int64Attribute{
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"protocol": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required:    true,
						Description: `must be one of ["http"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"http",
							),
						},
					},
				},
				Description: `Represents recipe networking parameters`,
			},
			"organization_name": schema.StringAttribute{
				Required:    true,
				Description: `The unique organization name`,
			},
			"project_name": schema.StringAttribute{
				Required:    true,
				Description: `The unique project name`,
			},
			"recipe": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `The unique identifier`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `The recipe name`,
					},
					"networking": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"auth": schema.BoolAttribute{
								Computed: true,
							},
							"dns": schema.StringAttribute{
								Computed: true,
							},
							"port": schema.Int64Attribute{
								Computed: true,
							},
							"protocol": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["http"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"http",
									),
								},
							},
						},
						Description: `Represents recipe networking parameters`,
					},
					"readme": schema.StringAttribute{
						Computed:    true,
						Description: `A markdown file containing a brief summary of the recipe`,
					},
					"resources": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"cpu": schema.Int64Attribute{
								Computed: true,
							},
							"gpu_class": schema.StringAttribute{
								Computed: true,
							},
							"ram": schema.Int64Attribute{
								Computed: true,
							},
						},
						Description: `Represents a recipe resources`,
					},
				},
				Description: `Represents a recipe`,
			},
			"recipe_name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"replicas": schema.Int64Attribute{
				Required: true,
			},
		},
	}
}

func (r *RecipeDeploymentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *RecipeDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RecipeDeploymentResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	createRecipeDeployment := *data.ToCreateSDKType()
	organizationName := data.OrganizationName.ValueString()
	projectName := data.ProjectName.ValueString()
	request := operations.CreateRecipeDeploymentRequest{
		CreateRecipeDeployment: createRecipeDeployment,
		OrganizationName:       organizationName,
		ProjectName:            projectName,
	}
	res, err := r.client.RecipeDeployments.CreateRecipeDeployment(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RecipeDeployment == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.RecipeDeployment)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecipeDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RecipeDeploymentResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	organizationName := data.OrganizationName.ValueString()
	projectName := data.ProjectName.ValueString()
	recipeDeploymentName := data.Name.ValueString()
	request := operations.GetRecipeDeploymentRequest{
		OrganizationName:     organizationName,
		ProjectName:          projectName,
		RecipeDeploymentName: recipeDeploymentName,
	}
	res, err := r.client.RecipeDeployments.GetRecipeDeployment(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RecipeDeployment == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.RecipeDeployment)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecipeDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *RecipeDeploymentResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	updateRecipeDeployment := *data.ToUpdateSDKType()
	organizationName := data.OrganizationName.ValueString()
	projectName := data.ProjectName.ValueString()
	recipeDeploymentName := data.Name.ValueString()
	request := operations.UpdateRecipeDeploymentRequest{
		UpdateRecipeDeployment: updateRecipeDeployment,
		OrganizationName:       organizationName,
		ProjectName:            projectName,
		RecipeDeploymentName:   recipeDeploymentName,
	}
	res, err := r.client.RecipeDeployments.UpdateRecipeDeployment(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RecipeDeployment == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.RecipeDeployment)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecipeDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RecipeDeploymentResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	organizationName := data.OrganizationName.ValueString()
	projectName := data.ProjectName.ValueString()
	recipeDeploymentName := data.Name.ValueString()
	request := operations.DeleteRecipeDeploymentRequest{
		OrganizationName:     organizationName,
		ProjectName:          projectName,
		RecipeDeploymentName: recipeDeploymentName,
	}
	res, err := r.client.RecipeDeployments.DeleteRecipeDeployment(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 202 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *RecipeDeploymentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource recipe_deployment. Reason: composite imports strings not supported.")
}
