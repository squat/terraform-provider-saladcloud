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
var _ datasource.DataSource = &ContainerGroupsDataSource{}
var _ datasource.DataSourceWithConfigure = &ContainerGroupsDataSource{}

func NewContainerGroupsDataSource() datasource.DataSource {
	return &ContainerGroupsDataSource{}
}

// ContainerGroupsDataSource is the data source implementation.
type ContainerGroupsDataSource struct {
	client *sdk.SDK
}

// ContainerGroupsDataSourceModel describes the data model.
type ContainerGroupsDataSourceModel struct {
	Items            []ContainerGroup `tfsdk:"items"`
	OrganizationName types.String     `tfsdk:"organization_name"`
	ProjectName      types.String     `tfsdk:"project_name"`
}

// Metadata returns the data source type name.
func (r *ContainerGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_container_groups"
}

// Schema defines the schema for the data source.
func (r *ContainerGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ContainerGroups DataSource",

		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"autostart_policy": schema.BoolAttribute{
							Computed: true,
						},
						"container": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"command": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"environment_variables": schema.MapAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"image": schema.StringAttribute{
									Computed: true,
								},
								"logging": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"new_relic": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Computed: true,
												},
												"ingestion_key": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"splunk": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Computed: true,
												},
												"token": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"tcp": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Computed: true,
												},
												"port": schema.Int64Attribute{
													Computed: true,
												},
											},
										},
									},
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
										"gpu_classes": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"memory": schema.Int64Attribute{
											Computed: true,
										},
									},
									Description: `Represents a container resource requirements`,
								},
							},
							Description: `Represents a container`,
						},
						"country_codes": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"create_time": schema.StringAttribute{
							Computed: true,
						},
						"current_state": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Computed: true,
								},
								"finish_time": schema.StringAttribute{
									Computed: true,
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
								},
								"status": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["pending", "running", "stopped", "succeeded", "failed", "deploying"]`,
								},
							},
							Description: `Represents a container group state`,
						},
						"display_name": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"liveness_probe": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"exec": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"command": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
									},
								},
								"failure_threshold": schema.Int64Attribute{
									Computed: true,
								},
								"grpc": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Computed: true,
										},
										"service": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"http": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"headers": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Computed: true,
													},
													"value": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
										"path": schema.StringAttribute{
											Computed: true,
										},
										"port": schema.Int64Attribute{
											Computed: true,
										},
										"scheme": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["http"]`,
										},
									},
								},
								"initial_delay_seconds": schema.Int64Attribute{
									Computed: true,
								},
								"period_seconds": schema.Int64Attribute{
									Computed: true,
								},
								"success_threshold": schema.Int64Attribute{
									Computed: true,
								},
								"tcp": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Computed: true,
										},
									},
								},
								"timeout_seconds": schema.Int64Attribute{
									Computed: true,
								},
							},
							Description: `Represents container group probe`,
						},
						"name": schema.StringAttribute{
							Computed: true,
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
								},
							},
							Description: `Represents container group networking parameters`,
						},
						"readiness_probe": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"exec": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"command": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
									},
								},
								"failure_threshold": schema.Int64Attribute{
									Computed: true,
								},
								"grpc": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Computed: true,
										},
										"service": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"http": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"headers": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Computed: true,
													},
													"value": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
										"path": schema.StringAttribute{
											Computed: true,
										},
										"port": schema.Int64Attribute{
											Computed: true,
										},
										"scheme": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["http"]`,
										},
									},
								},
								"initial_delay_seconds": schema.Int64Attribute{
									Computed: true,
								},
								"period_seconds": schema.Int64Attribute{
									Computed: true,
								},
								"success_threshold": schema.Int64Attribute{
									Computed: true,
								},
								"tcp": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Computed: true,
										},
									},
								},
								"timeout_seconds": schema.Int64Attribute{
									Computed: true,
								},
							},
							Description: `Represents container group probe`,
						},
						"replicas": schema.Int64Attribute{
							Computed: true,
						},
						"restart_policy": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["always", "on_failure", "never"]`,
						},
						"startup_probe": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"exec": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"command": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
									},
								},
								"failure_threshold": schema.Int64Attribute{
									Computed: true,
								},
								"grpc": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Computed: true,
										},
										"service": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"http": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"headers": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Computed: true,
													},
													"value": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
										"path": schema.StringAttribute{
											Computed: true,
										},
										"port": schema.Int64Attribute{
											Computed: true,
										},
										"scheme": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["http"]`,
										},
									},
								},
								"initial_delay_seconds": schema.Int64Attribute{
									Computed: true,
								},
								"period_seconds": schema.Int64Attribute{
									Computed: true,
								},
								"success_threshold": schema.Int64Attribute{
									Computed: true,
								},
								"tcp": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Computed: true,
										},
									},
								},
								"timeout_seconds": schema.Int64Attribute{
									Computed: true,
								},
							},
							Description: `Represents container group probe`,
						},
						"update_time": schema.StringAttribute{
							Computed: true,
						},
					},
				},
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

func (r *ContainerGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ContainerGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ContainerGroupsDataSourceModel
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
	projectName := data.ProjectName.ValueString()
	request := operations.ListContainerGroupsRequest{
		OrganizationName: organizationName,
		ProjectName:      projectName,
	}
	res, err := r.client.ContainerGroups.ListContainerGroups(ctx, request)
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
	if res.ContainerGroupList == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.ContainerGroupList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
