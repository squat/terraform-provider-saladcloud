// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ContainerGroupInstancesDataSource{}
var _ datasource.DataSourceWithConfigure = &ContainerGroupInstancesDataSource{}

func NewContainerGroupInstancesDataSource() datasource.DataSource {
	return &ContainerGroupInstancesDataSource{}
}

// ContainerGroupInstancesDataSource is the data source implementation.
type ContainerGroupInstancesDataSource struct {
	client *sdk.SDK
}

// ContainerGroupInstancesDataSourceModel describes the data model.
type ContainerGroupInstancesDataSourceModel struct {
	Instances        []Instances  `tfsdk:"instances"`
	Name             types.String `tfsdk:"name"`
	OrganizationName types.String `tfsdk:"organization_name"`
	ProjectName      types.String `tfsdk:"project_name"`
}

// Metadata returns the data source type name.
func (r *ContainerGroupInstancesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_container_group_instances"
}

// Schema defines the schema for the data source.
func (r *ContainerGroupInstancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ContainerGroupInstances DataSource",

		Attributes: map[string]schema.Attribute{
			"instances": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"machine_id": schema.StringAttribute{
							Computed:    true,
							Description: `The organization-specific machine ID`,
						},
						"state": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: `must be one of ["creating", "running"]` + "\n" +
								`The state of the container group instance`,
						},
						"update_time": schema.StringAttribute{
							Computed:    true,
							Description: `The UTC date & time when the workload on this machine transitioned to the current state`,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The unique container group name`,
			},
			"organization_name": schema.StringAttribute{
				Required:    true,
				Description: `The unique organization name`,
			},
			"project_name": schema.StringAttribute{
				Required:    true,
				Description: `The unique project name`,
			},
		},
	}
}

func (r *ContainerGroupInstancesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ContainerGroupInstancesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ContainerGroupInstancesDataSourceModel
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

	containerGroupName := data.Name.ValueString()
	organizationName := data.OrganizationName.ValueString()
	projectName := data.ProjectName.ValueString()
	request := operations.ListContainerGroupInstancesRequest{
		ContainerGroupName: containerGroupName,
		OrganizationName:   organizationName,
		ProjectName:        projectName,
	}
	res, err := r.client.ContainerGroups.ListContainerGroupInstances(ctx, request)
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
	if res.ContainerGroupInstances == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.ContainerGroupInstances)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
