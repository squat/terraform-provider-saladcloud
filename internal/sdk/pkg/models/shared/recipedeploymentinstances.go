// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/utils"
	"time"
)

// RecipeDeploymentInstancesState - The state of the recipe deployment instance
type RecipeDeploymentInstancesState string

const (
	RecipeDeploymentInstancesStateCreating RecipeDeploymentInstancesState = "creating"
	RecipeDeploymentInstancesStateRunning  RecipeDeploymentInstancesState = "running"
)

func (e RecipeDeploymentInstancesState) ToPointer() *RecipeDeploymentInstancesState {
	return &e
}

func (e *RecipeDeploymentInstancesState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "creating":
		fallthrough
	case "running":
		*e = RecipeDeploymentInstancesState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RecipeDeploymentInstancesState: %v", v)
	}
}

type RecipeDeploymentInstancesInstances struct {
	// The organization-specific machine ID
	MachineID string `json:"machine_id"`
	// The state of the recipe deployment instance
	State RecipeDeploymentInstancesState `json:"state"`
	// The UTC date & time when the workload on this machine transitioned to the current state
	UpdateTime time.Time `json:"update_time"`
}

func (r RecipeDeploymentInstancesInstances) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RecipeDeploymentInstancesInstances) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RecipeDeploymentInstancesInstances) GetMachineID() string {
	if o == nil {
		return ""
	}
	return o.MachineID
}

func (o *RecipeDeploymentInstancesInstances) GetState() RecipeDeploymentInstancesState {
	if o == nil {
		return RecipeDeploymentInstancesState("")
	}
	return o.State
}

func (o *RecipeDeploymentInstancesInstances) GetUpdateTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdateTime
}

// RecipeDeploymentInstances - Represents a list of recipe deployment instances
type RecipeDeploymentInstances struct {
	Instances []RecipeDeploymentInstancesInstances `json:"instances"`
}

func (o *RecipeDeploymentInstances) GetInstances() []RecipeDeploymentInstancesInstances {
	if o == nil {
		return []RecipeDeploymentInstancesInstances{}
	}
	return o.Instances
}
