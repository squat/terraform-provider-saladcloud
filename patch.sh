#! /usr/bin/env sh
# Setup
F=$(mktemp)
trap 'cat "$F"; exit' EXIT
cat - > "$F"
# General Patches
bin/yq -i '.info.title = "SaladCloud Provider"' "$F"
bin/yq -i '.info.description = "The SaladCloud provider enables declaratively managing resources provided by SaladCloud. The provider needs to be configured with the proper credentials before it can be used.\n\nFor information on obtaining an API key for SaladCloud, refer to [Authentication](https://docs.salad.com/reference/api-reference#authentication) from the [SaladCloud Documentation](https://docs.salad.com/)."' "$F"
bin/yq -i '.components.schemas.ContainerResourceRequirements.properties.memory.maximum = 30720' "$F"
bin/yq -i '.components.schemas.ContainerResourceRequirements.properties.gpu_class.deprecated = true' "$F"
bin/yq -i '.components.schemas.ContainerResourceRequirements.properties.gpu_classes = {"nullable": true, "type": "array", "items": {"type": "string"}}' "$F"
# Container Group
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers"].post.x-speakeasy-entity-operation = "ContainerGroup#create"' "$F"
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].get.x-speakeasy-entity-operation = "ContainerGroup#read"' "$F"
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].patch.x-speakeasy-entity-operation = "ContainerGroup#update"' "$F"
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].delete.x-speakeasy-entity-operation = "ContainerGroup#delete"' "$F"
bin/yq -i '.components.schemas.ContainerGroup.x-speakeasy-entity = "ContainerGroup"' "$F"
bin/yq -i '.components.schemas.CreateContainerGroup.x-speakeasy-entity = "ContainerGroup"' "$F"
bin/yq -i '.components.schemas.ContainerGroup.properties.autostart_policy = {"type": "boolean"}' "$F"
bin/yq -i '.components.schemas.ContainerGroup.required += ["autostart_policy"]' "$F"
bin/yq -i '.components.schemas.CreateContainerGroup.properties.autostart_policy = {"type": "boolean"}' "$F"
bin/yq -i '.components.schemas.CreateContainerGroup.required += ["autostart_policy"]' "$F"
bin/yq -i '.components.parameters.container_group_name.x-speakeasy-match = "name"' "$F"
# Container Group Instances
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances"].get.x-speakeasy-entity-operation = "ContainerGroupInstances#read"' "$F"
bin/yq -i '.components.schemas.ContainerGroupInstances.x-speakeasy-entity = "ContainerGroupInstances"' "$F"
# Container Groups
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers"].get.x-speakeasy-entity-operation = "ContainerGroups#read"' "$F"
bin/yq -i '.components.schemas.ContainerGroupList.x-speakeasy-entity = "ContainerGroups"' "$F"
# GPU Classes
bin/yq -i '.paths.["/organizations/{organization_name}/gpu-classes"].get.x-speakeasy-entity-operation = "GPUClasses#read"' "$F"
bin/yq -i '.components.schemas.GpuClassesList.x-speakeasy-entity = "GPUClasses"' "$F"
# Quotas
bin/yq -i '.paths.["/organizations/{organization_name}/quotas"].get.x-speakeasy-entity-operation = "Quotas#read"' "$F"
bin/yq -i '.components.schemas.Quotas.x-speakeasy-entity = "Quotas"' "$F"
# Recipe
bin/yq -i '.paths.["/organizations/{organization_name}/recipes/{recipe_name}"].get.x-speakeasy-entity-operation = "Recipe#read"' "$F"
bin/yq -i '.components.schemas.Recipe.x-speakeasy-entity = "Recipe"' "$F"
bin/yq -i '.components.parameters.recipe_name.x-speakeasy-match = "name"' "$F"
# Recipes
bin/yq -i '.paths.["/organizations/{organization_name}/recipes"].get.x-speakeasy-entity-operation = "Recipes#read"' "$F"
bin/yq -i '.components.schemas.RecipeList.x-speakeasy-entity = "Recipes"' "$F"
# Recipe Deployment
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments"].post.x-speakeasy-entity-operation = "RecipeDeployment#create"' "$F"
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].get.x-speakeasy-entity-operation = "RecipeDeployment#read"' "$F"
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].patch.x-speakeasy-entity-operation = "RecipeDeployment#update"' "$F"
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].delete.x-speakeasy-entity-operation = "RecipeDeployment#delete"' "$F"
bin/yq -i '.components.schemas.RecipeDeployment.x-speakeasy-entity = "RecipeDeployment"' "$F"
bin/yq -i '.components.parameters.recipe_deployment_name.x-speakeasy-match = "name"' "$F"
# Recipe Deployment Instances
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/instances"].get.x-speakeasy-entity-operation = "RecipeDeploymentInstances#read"' "$F"
bin/yq -i '.components.schemas.RecipeDeploymentInstances.x-speakeasy-entity = "RecipeDeploymentInstances"' "$F"
# Recipe Deployments
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments"].get.x-speakeasy-entity-operation = "RecipeDeployments#read"' "$F"
bin/yq -i '.components.schemas.RecipeDeploymentList.x-speakeasy-entity = "RecipeDeployments"' "$F"
