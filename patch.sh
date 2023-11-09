#! /usr/bin/env sh
# General Patches
bin/yq -i '.components.schemas.ContainerResourceRequirements.properties.memory.maximum = 30720' saladcloud.yaml
# Container Group
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers"].post.x-speakeasy-entity-operation = "ContainerGroup#create"' saladcloud.yaml
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].get.x-speakeasy-entity-operation = "ContainerGroup#read"' saladcloud.yaml
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].patch.x-speakeasy-entity-operation = "ContainerGroup#update"' saladcloud.yaml
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}"].delete.x-speakeasy-entity-operation = "ContainerGroup#delete"' saladcloud.yaml
bin/yq -i '.components.schemas.ContainerGroup.x-speakeasy-entity = "ContainerGroup"' saladcloud.yaml
bin/yq -i '.components.parameters.container_group_name.x-speakeasy-match = "name"' saladcloud.yaml
# Container Group Instances
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances"].get.x-speakeasy-entity-operation = "ContainerGroupInstances#read"' saladcloud.yaml
bin/yq -i '.components.schemas.ContainerGroupInstances.x-speakeasy-entity = "ContainerGroupInstances"' saladcloud.yaml
# Container Groups
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/containers"].get.x-speakeasy-entity-operation = "ContainerGroups#read"' saladcloud.yaml
bin/yq -i '.components.schemas.ContainerGroupList.x-speakeasy-entity = "ContainerGroups"' saladcloud.yaml
# GPU Classes
bin/yq -i '.paths.["/organizations/{organization_name}/gpu-classes"].get.x-speakeasy-entity-operation = "GPUClasses#read"' saladcloud.yaml
bin/yq -i '.components.schemas.GpuClassesList.x-speakeasy-entity = "GPUClasses"' saladcloud.yaml
# Quotas
bin/yq -i '.paths.["/organizations/{organization_name}/quotas"].get.x-speakeasy-entity-operation = "Quotas#read"' saladcloud.yaml
bin/yq -i '.components.schemas.Quotas.x-speakeasy-entity = "Quotas"' saladcloud.yaml
# Recipe
bin/yq -i '.paths.["/organizations/{organization_name}/recipes/{recipe_name}"].get.x-speakeasy-entity-operation = "Recipe#read"' saladcloud.yaml
bin/yq -i '.components.schemas.Recipe.x-speakeasy-entity = "Recipe"' saladcloud.yaml
bin/yq -i '.components.parameters.recipe_name.x-speakeasy-match = "name"' saladcloud.yaml
# Recipes
bin/yq -i '.paths.["/organizations/{organization_name}/recipes"].get.x-speakeasy-entity-operation = "Recipes#read"' saladcloud.yaml
bin/yq -i '.components.schemas.RecipeList.x-speakeasy-entity = "Recipes"' saladcloud.yaml
# Recipe Deployment
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments"].post.x-speakeasy-entity-operation = "RecipeDeployment#create"' saladcloud.yaml
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].get.x-speakeasy-entity-operation = "RecipeDeployment#read"' saladcloud.yaml
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].patch.x-speakeasy-entity-operation = "RecipeDeployment#update"' saladcloud.yaml
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}"].delete.x-speakeasy-entity-operation = "RecipeDeployment#delete"' saladcloud.yaml
bin/yq -i '.components.schemas.RecipeDeployment.x-speakeasy-entity = "RecipeDeployment"' saladcloud.yaml
bin/yq -i '.components.parameters.recipe_deployment_name.x-speakeasy-match = "name"' saladcloud.yaml
# Recipe Deployment Instances
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/instances"].get.x-speakeasy-entity-operation = "RecipeDeploymentInstances#read"' saladcloud.yaml
bin/yq -i '.components.schemas.RecipeDeploymentInstances.x-speakeasy-entity = "RecipeDeploymentInstances"' saladcloud.yaml
# Recipe Deployments
bin/yq -i '.paths.["/organizations/{organization_name}/projects/{project_name}/recipe-deployments"].get.x-speakeasy-entity-operation = "RecipeDeployments#read"' saladcloud.yaml
bin/yq -i '.components.schemas.RecipeDeploymentList.x-speakeasy-entity = "RecipeDeployments"' saladcloud.yaml
