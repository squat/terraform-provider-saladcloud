#! /usr/bin/env sh
# Setup
F=$(mktemp)
trap 'cat "$F"; exit' EXIT
export YQ="go run github.com/mikefarah/yq/v4@v4.35.2"
cat - > "$F"
# General Patches
$YQ -i '.info.title = "SaladCloud Provider"' "$F"
$YQ -i '.info.description = "The SaladCloud provider enables declaratively managing resources provided by SaladCloud. The provider needs to be configured with the proper credentials before it can be used.\n\nFor information on obtaining an API key for SaladCloud, refer to [Authentication](https://docs.salad.com/reference/api-reference#authentication) from the [SaladCloud Documentation](https://docs.salad.com/)."' "$F"
$YQ -i '.components.securitySchemes.apiKey = .components.securitySchemes.ApiKeyAuth' "$F"
$YQ -i 'del(.components.securitySchemes.ApiKeyAuth)' "$F"
$YQ -i 'del(.security[0])' "$F"
$YQ -i '.security[0] = {"apiKey": []}' "$F"
$YQ -i '.components.schemas.ContainerResourceRequirements.properties.memory.maximum = 30720' "$F"
$YQ -i '.components.schemas.ContainerResourceRequirements.properties.gpu_class.deprecated = true' "$F"
$YQ -i '.components.schemas.ContainerResourceRequirements.properties.gpu_classes = {"nullable": true, "type": "array", "items": {"type": "string"}}' "$F"
# Container Group
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers"].post.x-speakeasy-entity-operation = "ContainerGroup#create"' "$F"
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].get.x-speakeasy-entity-operation = "ContainerGroup#read"' "$F"
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].patch.x-speakeasy-entity-operation = "ContainerGroup#update"' "$F"
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].delete.x-speakeasy-entity-operation = "ContainerGroup#delete"' "$F"
$YQ -i '.components.schemas.ContainerGroup.x-speakeasy-entity = "ContainerGroup"' "$F"
$YQ -i '.components.schemas.CreateContainerGroup.x-speakeasy-entity = "ContainerGroup"' "$F"
$YQ -i '.components.schemas.ContainerGroup.properties.autostart_policy = {"type": "boolean"}' "$F"
$YQ -i '.components.schemas.ContainerGroup.required += ["autostart_policy"]' "$F"
$YQ -i '.components.schemas.CreateContainerGroup.properties.autostart_policy = {"type": "boolean"}' "$F"
$YQ -i '.components.schemas.CreateContainerGroup.required += ["autostart_policy"]' "$F"
$YQ -i '.components.parameters.container_group_name.x-speakeasy-match = "name"' "$F"
# Container Group Instances
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances"].get.x-speakeasy-entity-operation = "ContainerGroupInstances#read"' "$F"
$YQ -i '.components.schemas.ContainerGroupInstances.x-speakeasy-entity = "ContainerGroupInstances"' "$F"
# Container Groups
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers"].get.x-speakeasy-entity-operation = "ContainerGroups#read"' "$F"
$YQ -i '.components.schemas.ContainerGroupList.x-speakeasy-entity = "ContainerGroups"' "$F"
# GPU Classes
$YQ -i '.paths.["/organizations/{organization_name}/gpu-classes"].get.x-speakeasy-entity-operation = "GPUClasses#read"' "$F"
$YQ -i '.components.schemas.GpuClassesList.x-speakeasy-entity = "GPUClasses"' "$F"
# Quotas
$YQ -i '.paths.["/organizations/{organization_name}/quotas"].get.x-speakeasy-entity-operation = "Quotas#read"' "$F"
$YQ -i '.components.schemas.Quotas.x-speakeasy-entity = "Quotas"' "$F"
# Recipe
$YQ -i '.paths.["/organizations/{organization_name}/recipes/{recipe_name}"].get.x-speakeasy-entity-operation = "Recipe#read"' "$F"
$YQ -i '.components.schemas.Recipe.x-speakeasy-entity = "Recipe"' "$F"
$YQ -i '.components.parameters.recipe_name.x-speakeasy-match = "name"' "$F"
# Recipes
$YQ -i '.paths.["/organizations/{organization_name}/recipes"].get.x-speakeasy-entity-operation = "Recipes#read"' "$F"
$YQ -i '.components.schemas.RecipeList.x-speakeasy-entity = "Recipes"' "$F"
# Recipe Deployment
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments"].post.x-speakeasy-entity-operation = "RecipeDeployment#create"' "$F"
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].get.x-speakeasy-entity-operation = "RecipeDeployment#read"' "$F"
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].patch.x-speakeasy-entity-operation = "RecipeDeployment#update"' "$F"
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].delete.x-speakeasy-entity-operation = "RecipeDeployment#delete"' "$F"
$YQ -i '.components.schemas.RecipeDeployment.x-speakeasy-entity = "RecipeDeployment"' "$F"
$YQ -i '.components.parameters.recipe_deployment_name.x-speakeasy-match = "name"' "$F"
# Recipe Deployment Instances
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/instances"].get.x-speakeasy-entity-operation = "RecipeDeploymentInstances#read"' "$F"
$YQ -i '.components.schemas.RecipeDeploymentInstances.x-speakeasy-entity = "RecipeDeploymentInstances"' "$F"
# Recipe Deployments
$YQ -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments"].get.x-speakeasy-entity-operation = "RecipeDeployments#read"' "$F"
$YQ -i '.components.schemas.RecipeDeploymentList.x-speakeasy-entity = "RecipeDeployments"' "$F"
