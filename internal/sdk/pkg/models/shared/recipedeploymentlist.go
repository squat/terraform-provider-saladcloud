// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RecipeDeploymentList - Represents a list of recipe deployments
type RecipeDeploymentList struct {
	Items []RecipeDeployment `json:"items"`
}

func (o *RecipeDeploymentList) GetItems() []RecipeDeployment {
	if o == nil {
		return []RecipeDeployment{}
	}
	return o.Items
}
