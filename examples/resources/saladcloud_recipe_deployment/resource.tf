resource "saladcloud_recipe_deployment" "my_recipedeployment" {
  display_name = "...my_display_name..."
  name         = "Mrs. Elbert Buckridge"
  networking = {
    auth     = false
    dns      = "...my_dns..."
    port     = 9
    protocol = "http"
  }
  organization_name = "...my_organization_name..."
  project_name      = "...my_project_name..."
  recipe_name       = "...my_recipe_name..."
  replicas          = 7
}