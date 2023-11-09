// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContainerGroupInstanceStatusCount - Represents a container group instance status count
type ContainerGroupInstanceStatusCount struct {
	AllocatingCount int64 `json:"allocating_count"`
	CreatingCount   int64 `json:"creating_count"`
	RunningCount    int64 `json:"running_count"`
}

func (o *ContainerGroupInstanceStatusCount) GetAllocatingCount() int64 {
	if o == nil {
		return 0
	}
	return o.AllocatingCount
}

func (o *ContainerGroupInstanceStatusCount) GetCreatingCount() int64 {
	if o == nil {
		return 0
	}
	return o.CreatingCount
}

func (o *ContainerGroupInstanceStatusCount) GetRunningCount() int64 {
	if o == nil {
		return 0
	}
	return o.RunningCount
}
