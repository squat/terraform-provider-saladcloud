// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ContainerGroupInstanceStatusCount struct {
	AllocatingCount types.Int64 `tfsdk:"allocating_count"`
	CreatingCount   types.Int64 `tfsdk:"creating_count"`
	RunningCount    types.Int64 `tfsdk:"running_count"`
}
