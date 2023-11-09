// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContainerGroupProbeGrpc struct {
	Port    int64  `json:"port"`
	Service string `json:"service"`
}

func (o *ContainerGroupProbeGrpc) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ContainerGroupProbeGrpc) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}
