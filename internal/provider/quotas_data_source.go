// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/squat/terraform-provider-saladcloud/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &QuotasDataSource{}
var _ datasource.DataSourceWithConfigure = &QuotasDataSource{}

func NewQuotasDataSource() datasource.DataSource {
	return &QuotasDataSource{}
}

// QuotasDataSource is the data source implementation.
type QuotasDataSource struct {
	client *sdk.SDK
}

// QuotasDataSourceModel describes the data model.
type QuotasDataSourceModel struct {
	ContainerGroupsQuotas ContainerGroupsQuotas `tfsdk:"container_groups_quotas"`
	CreateTime            types.String          `tfsdk:"create_time"`
	OrganizationName      types.String          `tfsdk:"organization_name"`
	RecipesQuotas         RecipesQuotas         `tfsdk:"recipes_quotas"`
	UpdateTime            types.String          `tfsdk:"update_time"`
}

// Metadata returns the data source type name.
func (r *QuotasDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quotas"
}

// Schema defines the schema for the data source.
func (r *QuotasDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Quotas DataSource",

		Attributes: map[string]schema.Attribute{
			"container_groups_quotas": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"container_instance_quota": schema.Int64Attribute{
						Computed: true,
					},
					"max_container_group_reallocations_per_minute": schema.Int64Attribute{
						Computed:    true,
						Description: `Default: 10`,
					},
					"max_container_group_recreates_per_minute": schema.Int64Attribute{
						Computed:    true,
						Description: `Default: 10`,
					},
					"max_container_group_restarts_per_minute": schema.Int64Attribute{
						Computed:    true,
						Description: `Default: 10`,
					},
					"max_created_container_groups": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"create_time": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The time the resource was created`,
			},
			"organization_name": schema.StringAttribute{
				Required:    true,
				Description: `The unique organization name`,
			},
			"recipes_quotas": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"max_created_recipe_deployments": schema.Int64Attribute{
						Computed: true,
					},
					"recipe_instance_quota": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"update_time": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The time the resource was last updated`,
			},
		},
	}
}

func (r *QuotasDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *QuotasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *QuotasDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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
	request := operations.GetQuotasRequest{
		OrganizationName: organizationName,
	}
	res, err := r.client.Quotas.GetQuotas(ctx, request)
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
	if res.Quotas == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Quotas)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
