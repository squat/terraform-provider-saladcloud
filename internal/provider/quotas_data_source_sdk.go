// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"time"
)

func (r *QuotasDataSourceModel) RefreshFromGetResponse(resp *shared.Quotas) {
	r.ContainerGroupsQuotas.ContainerInstanceQuota = types.Int64Value(resp.ContainerGroupsQuotas.ContainerInstanceQuota)
	if resp.ContainerGroupsQuotas.MaxContainerGroupReallocationsPerMinute != nil {
		r.ContainerGroupsQuotas.MaxContainerGroupReallocationsPerMinute = types.Int64Value(*resp.ContainerGroupsQuotas.MaxContainerGroupReallocationsPerMinute)
	} else {
		r.ContainerGroupsQuotas.MaxContainerGroupReallocationsPerMinute = types.Int64Null()
	}
	if resp.ContainerGroupsQuotas.MaxContainerGroupRecreatesPerMinute != nil {
		r.ContainerGroupsQuotas.MaxContainerGroupRecreatesPerMinute = types.Int64Value(*resp.ContainerGroupsQuotas.MaxContainerGroupRecreatesPerMinute)
	} else {
		r.ContainerGroupsQuotas.MaxContainerGroupRecreatesPerMinute = types.Int64Null()
	}
	if resp.ContainerGroupsQuotas.MaxContainerGroupRestartsPerMinute != nil {
		r.ContainerGroupsQuotas.MaxContainerGroupRestartsPerMinute = types.Int64Value(*resp.ContainerGroupsQuotas.MaxContainerGroupRestartsPerMinute)
	} else {
		r.ContainerGroupsQuotas.MaxContainerGroupRestartsPerMinute = types.Int64Null()
	}
	r.ContainerGroupsQuotas.MaxCreatedContainerGroups = types.Int64Value(resp.ContainerGroupsQuotas.MaxCreatedContainerGroups)
	if resp.CreateTime != nil {
		r.CreateTime = types.StringValue(resp.CreateTime.Format(time.RFC3339Nano))
	} else {
		r.CreateTime = types.StringNull()
	}
	r.RecipesQuotas.MaxCreatedRecipeDeployments = types.Int64Value(resp.RecipesQuotas.MaxCreatedRecipeDeployments)
	r.RecipesQuotas.RecipeInstanceQuota = types.Int64Value(resp.RecipesQuotas.RecipeInstanceQuota)
	if resp.UpdateTime != nil {
		r.UpdateTime = types.StringValue(resp.UpdateTime.Format(time.RFC3339Nano))
	} else {
		r.UpdateTime = types.StringNull()
	}
}
