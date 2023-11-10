// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContainerResourceRequirements - Represents a container resource requirements
type ContainerResourceRequirements struct {
	CPU int64 `json:"cpu"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	GpuClass   *string  `json:"gpu_class,omitempty"`
	GpuClasses []string `json:"gpu_classes,omitempty"`
	Memory     int64    `json:"memory"`
}

func (o *ContainerResourceRequirements) GetCPU() int64 {
	if o == nil {
		return 0
	}
	return o.CPU
}

func (o *ContainerResourceRequirements) GetGpuClass() *string {
	if o == nil {
		return nil
	}
	return o.GpuClass
}

func (o *ContainerResourceRequirements) GetGpuClasses() []string {
	if o == nil {
		return nil
	}
	return o.GpuClasses
}

func (o *ContainerResourceRequirements) GetMemory() int64 {
	if o == nil {
		return 0
	}
	return o.Memory
}
