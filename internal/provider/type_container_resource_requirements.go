// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ContainerResourceRequirements struct {
	CPU      types.Int64  `tfsdk:"cpu"`
	GpuClass types.String `tfsdk:"gpu_class"`
	Memory   types.Int64  `tfsdk:"memory"`
}
