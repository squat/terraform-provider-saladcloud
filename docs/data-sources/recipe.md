---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "saladcloud_recipe Data Source - terraform-provider-saladcloud"
subcategory: ""
description: |-
  Recipe DataSource
---

# saladcloud_recipe (Data Source)

Recipe DataSource

## Example Usage

```terraform
data "saladcloud_recipe" "my_recipe" {
  name              = "Beth Quigley"
  organization_name = "...my_organization_name..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The unique recipe name
- `organization_name` (String) The unique organization name

### Read-Only

- `id` (String) The unique identifier
- `networking` (Attributes) Represents recipe networking parameters (see [below for nested schema](#nestedatt--networking))
- `readme` (String) A markdown file containing a brief summary of the recipe
- `resources` (Attributes) Represents a recipe resources (see [below for nested schema](#nestedatt--resources))

<a id="nestedatt--networking"></a>
### Nested Schema for `networking`

Read-Only:

- `auth` (Boolean)
- `dns` (String)
- `port` (Number)
- `protocol` (String) must be one of ["http"]


<a id="nestedatt--resources"></a>
### Nested Schema for `resources`

Read-Only:

- `cpu` (Number)
- `gpu_class` (String)
- `ram` (Number)


