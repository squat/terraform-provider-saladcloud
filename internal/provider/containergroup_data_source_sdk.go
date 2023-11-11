// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"time"
)

func (r *ContainerGroupDataSourceModel) RefreshFromGetResponse(resp *shared.ContainerGroup) {
	r.AutostartPolicy = types.BoolValue(resp.AutostartPolicy)
	r.Container.Command = nil
	for _, v := range resp.Container.Command {
		r.Container.Command = append(r.Container.Command, types.StringValue(v))
	}
	if r.Container.EnvironmentVariables == nil && len(resp.Container.EnvironmentVariables) > 0 {
		r.Container.EnvironmentVariables = make(map[string]types.String)
		for key, value := range resp.Container.EnvironmentVariables {
			r.Container.EnvironmentVariables[key] = types.StringValue(value)
		}
	}
	r.Container.Image = types.StringValue(resp.Container.Image)
	if resp.Container.Logging == nil {
		r.Container.Logging = nil
	} else {
		r.Container.Logging = &Logging{}
		if resp.Container.Logging.NewRelic == nil {
			r.Container.Logging.NewRelic = nil
		} else {
			r.Container.Logging.NewRelic = &NewRelic{}
			r.Container.Logging.NewRelic.Host = types.StringValue(resp.Container.Logging.NewRelic.Host)
			r.Container.Logging.NewRelic.IngestionKey = types.StringValue(resp.Container.Logging.NewRelic.IngestionKey)
		}
		if resp.Container.Logging.Splunk == nil {
			r.Container.Logging.Splunk = nil
		} else {
			r.Container.Logging.Splunk = &Splunk{}
			r.Container.Logging.Splunk.Host = types.StringValue(resp.Container.Logging.Splunk.Host)
			r.Container.Logging.Splunk.Token = types.StringValue(resp.Container.Logging.Splunk.Token)
		}
		if resp.Container.Logging.TCP == nil {
			r.Container.Logging.TCP = nil
		} else {
			r.Container.Logging.TCP = &TCP{}
			r.Container.Logging.TCP.Host = types.StringValue(resp.Container.Logging.TCP.Host)
			r.Container.Logging.TCP.Port = types.Int64Value(resp.Container.Logging.TCP.Port)
		}
	}
	r.Container.Resources.CPU = types.Int64Value(resp.Container.Resources.CPU)
	if resp.Container.Resources.GpuClass != nil {
		r.Container.Resources.GpuClass = types.StringValue(*resp.Container.Resources.GpuClass)
	} else {
		r.Container.Resources.GpuClass = types.StringNull()
	}
	r.Container.Resources.GpuClasses = nil
	for _, v := range resp.Container.Resources.GpuClasses {
		r.Container.Resources.GpuClasses = append(r.Container.Resources.GpuClasses, types.StringValue(v))
	}
	r.Container.Resources.Memory = types.Int64Value(resp.Container.Resources.Memory)
	r.CountryCodes = nil
	for _, v := range resp.CountryCodes {
		r.CountryCodes = append(r.CountryCodes, types.StringValue(string(v)))
	}
	r.CreateTime = types.StringValue(resp.CreateTime.Format(time.RFC3339Nano))
	if resp.CurrentState.Description != nil {
		r.CurrentState.Description = types.StringValue(*resp.CurrentState.Description)
	} else {
		r.CurrentState.Description = types.StringNull()
	}
	r.CurrentState.FinishTime = types.StringValue(resp.CurrentState.FinishTime.Format(time.RFC3339Nano))
	r.CurrentState.InstanceStatusCount.AllocatingCount = types.Int64Value(resp.CurrentState.InstanceStatusCount.AllocatingCount)
	r.CurrentState.InstanceStatusCount.CreatingCount = types.Int64Value(resp.CurrentState.InstanceStatusCount.CreatingCount)
	r.CurrentState.InstanceStatusCount.RunningCount = types.Int64Value(resp.CurrentState.InstanceStatusCount.RunningCount)
	r.CurrentState.StartTime = types.StringValue(resp.CurrentState.StartTime.Format(time.RFC3339Nano))
	r.CurrentState.Status = types.StringValue(string(resp.CurrentState.Status))
	r.DisplayName = types.StringValue(resp.DisplayName)
	r.ID = types.StringValue(resp.ID)
	if resp.LivenessProbe == nil {
		r.LivenessProbe = nil
	} else {
		r.LivenessProbe = &ContainerGroupProbe{}
		if resp.LivenessProbe.Exec == nil {
			r.LivenessProbe.Exec = nil
		} else {
			r.LivenessProbe.Exec = &ContainerGroupProbeExec{}
			r.LivenessProbe.Exec.Command = nil
			for _, v := range resp.LivenessProbe.Exec.Command {
				r.LivenessProbe.Exec.Command = append(r.LivenessProbe.Exec.Command, types.StringValue(v))
			}
		}
		r.LivenessProbe.FailureThreshold = types.Int64Value(resp.LivenessProbe.FailureThreshold)
		if resp.LivenessProbe.Grpc == nil {
			r.LivenessProbe.Grpc = nil
		} else {
			r.LivenessProbe.Grpc = &ContainerGroupProbeGrpc{}
			r.LivenessProbe.Grpc.Port = types.Int64Value(resp.LivenessProbe.Grpc.Port)
			r.LivenessProbe.Grpc.Service = types.StringValue(resp.LivenessProbe.Grpc.Service)
		}
		if resp.LivenessProbe.HTTP == nil {
			r.LivenessProbe.HTTP = nil
		} else {
			r.LivenessProbe.HTTP = &ContainerGroupProbeHTTP{}
			if len(r.LivenessProbe.HTTP.Headers) > len(resp.LivenessProbe.HTTP.Headers) {
				r.LivenessProbe.HTTP.Headers = r.LivenessProbe.HTTP.Headers[:len(resp.LivenessProbe.HTTP.Headers)]
			}
			for headersCount, headersItem := range resp.LivenessProbe.HTTP.Headers {
				var headers1 HTTPHeaders
				headers1.Name = types.StringValue(headersItem.Name)
				headers1.Value = types.StringValue(headersItem.Value)
				if headersCount+1 > len(r.LivenessProbe.HTTP.Headers) {
					r.LivenessProbe.HTTP.Headers = append(r.LivenessProbe.HTTP.Headers, headers1)
				} else {
					r.LivenessProbe.HTTP.Headers[headersCount].Name = headers1.Name
					r.LivenessProbe.HTTP.Headers[headersCount].Value = headers1.Value
				}
			}
			r.LivenessProbe.HTTP.Path = types.StringValue(resp.LivenessProbe.HTTP.Path)
			r.LivenessProbe.HTTP.Port = types.Int64Value(resp.LivenessProbe.HTTP.Port)
			if resp.LivenessProbe.HTTP.Scheme != nil {
				r.LivenessProbe.HTTP.Scheme = types.StringValue(string(*resp.LivenessProbe.HTTP.Scheme))
			} else {
				r.LivenessProbe.HTTP.Scheme = types.StringNull()
			}
		}
		r.LivenessProbe.InitialDelaySeconds = types.Int64Value(resp.LivenessProbe.InitialDelaySeconds)
		r.LivenessProbe.PeriodSeconds = types.Int64Value(resp.LivenessProbe.PeriodSeconds)
		r.LivenessProbe.SuccessThreshold = types.Int64Value(resp.LivenessProbe.SuccessThreshold)
		if resp.LivenessProbe.TCP == nil {
			r.LivenessProbe.TCP = nil
		} else {
			r.LivenessProbe.TCP = &ContainerGroupProbeTCP{}
			r.LivenessProbe.TCP.Port = types.Int64Value(resp.LivenessProbe.TCP.Port)
		}
		r.LivenessProbe.TimeoutSeconds = types.Int64Value(resp.LivenessProbe.TimeoutSeconds)
	}
	r.Name = types.StringValue(resp.Name)
	if resp.Networking == nil {
		r.Networking = nil
	} else {
		r.Networking = &ContainerGroupNetworking{}
		r.Networking.Auth = types.BoolValue(resp.Networking.Auth)
		r.Networking.DNS = types.StringValue(resp.Networking.DNS)
		r.Networking.Port = types.Int64Value(resp.Networking.Port)
		r.Networking.Protocol = types.StringValue(string(resp.Networking.Protocol))
	}
	if resp.ReadinessProbe == nil {
		r.ReadinessProbe = nil
	} else {
		r.ReadinessProbe = &ContainerGroupProbe{}
		if resp.ReadinessProbe.Exec == nil {
			r.ReadinessProbe.Exec = nil
		} else {
			r.ReadinessProbe.Exec = &ContainerGroupProbeExec{}
			r.ReadinessProbe.Exec.Command = nil
			for _, v := range resp.ReadinessProbe.Exec.Command {
				r.ReadinessProbe.Exec.Command = append(r.ReadinessProbe.Exec.Command, types.StringValue(v))
			}
		}
		r.ReadinessProbe.FailureThreshold = types.Int64Value(resp.ReadinessProbe.FailureThreshold)
		if resp.ReadinessProbe.Grpc == nil {
			r.ReadinessProbe.Grpc = nil
		} else {
			r.ReadinessProbe.Grpc = &ContainerGroupProbeGrpc{}
			r.ReadinessProbe.Grpc.Port = types.Int64Value(resp.ReadinessProbe.Grpc.Port)
			r.ReadinessProbe.Grpc.Service = types.StringValue(resp.ReadinessProbe.Grpc.Service)
		}
		if resp.ReadinessProbe.HTTP == nil {
			r.ReadinessProbe.HTTP = nil
		} else {
			r.ReadinessProbe.HTTP = &ContainerGroupProbeHTTP{}
			if len(r.ReadinessProbe.HTTP.Headers) > len(resp.ReadinessProbe.HTTP.Headers) {
				r.ReadinessProbe.HTTP.Headers = r.ReadinessProbe.HTTP.Headers[:len(resp.ReadinessProbe.HTTP.Headers)]
			}
			for headersCount1, headersItem1 := range resp.ReadinessProbe.HTTP.Headers {
				var headers3 HTTPHeaders
				headers3.Name = types.StringValue(headersItem1.Name)
				headers3.Value = types.StringValue(headersItem1.Value)
				if headersCount1+1 > len(r.ReadinessProbe.HTTP.Headers) {
					r.ReadinessProbe.HTTP.Headers = append(r.ReadinessProbe.HTTP.Headers, headers3)
				} else {
					r.ReadinessProbe.HTTP.Headers[headersCount1].Name = headers3.Name
					r.ReadinessProbe.HTTP.Headers[headersCount1].Value = headers3.Value
				}
			}
			r.ReadinessProbe.HTTP.Path = types.StringValue(resp.ReadinessProbe.HTTP.Path)
			r.ReadinessProbe.HTTP.Port = types.Int64Value(resp.ReadinessProbe.HTTP.Port)
			if resp.ReadinessProbe.HTTP.Scheme != nil {
				r.ReadinessProbe.HTTP.Scheme = types.StringValue(string(*resp.ReadinessProbe.HTTP.Scheme))
			} else {
				r.ReadinessProbe.HTTP.Scheme = types.StringNull()
			}
		}
		r.ReadinessProbe.InitialDelaySeconds = types.Int64Value(resp.ReadinessProbe.InitialDelaySeconds)
		r.ReadinessProbe.PeriodSeconds = types.Int64Value(resp.ReadinessProbe.PeriodSeconds)
		r.ReadinessProbe.SuccessThreshold = types.Int64Value(resp.ReadinessProbe.SuccessThreshold)
		if resp.ReadinessProbe.TCP == nil {
			r.ReadinessProbe.TCP = nil
		} else {
			r.ReadinessProbe.TCP = &ContainerGroupProbeTCP{}
			r.ReadinessProbe.TCP.Port = types.Int64Value(resp.ReadinessProbe.TCP.Port)
		}
		r.ReadinessProbe.TimeoutSeconds = types.Int64Value(resp.ReadinessProbe.TimeoutSeconds)
	}
	r.Replicas = types.Int64Value(resp.Replicas)
	r.RestartPolicy = types.StringValue(string(resp.RestartPolicy))
	if resp.StartupProbe == nil {
		r.StartupProbe = nil
	} else {
		r.StartupProbe = &ContainerGroupProbe{}
		if resp.StartupProbe.Exec == nil {
			r.StartupProbe.Exec = nil
		} else {
			r.StartupProbe.Exec = &ContainerGroupProbeExec{}
			r.StartupProbe.Exec.Command = nil
			for _, v := range resp.StartupProbe.Exec.Command {
				r.StartupProbe.Exec.Command = append(r.StartupProbe.Exec.Command, types.StringValue(v))
			}
		}
		r.StartupProbe.FailureThreshold = types.Int64Value(resp.StartupProbe.FailureThreshold)
		if resp.StartupProbe.Grpc == nil {
			r.StartupProbe.Grpc = nil
		} else {
			r.StartupProbe.Grpc = &ContainerGroupProbeGrpc{}
			r.StartupProbe.Grpc.Port = types.Int64Value(resp.StartupProbe.Grpc.Port)
			r.StartupProbe.Grpc.Service = types.StringValue(resp.StartupProbe.Grpc.Service)
		}
		if resp.StartupProbe.HTTP == nil {
			r.StartupProbe.HTTP = nil
		} else {
			r.StartupProbe.HTTP = &ContainerGroupProbeHTTP{}
			if len(r.StartupProbe.HTTP.Headers) > len(resp.StartupProbe.HTTP.Headers) {
				r.StartupProbe.HTTP.Headers = r.StartupProbe.HTTP.Headers[:len(resp.StartupProbe.HTTP.Headers)]
			}
			for headersCount2, headersItem2 := range resp.StartupProbe.HTTP.Headers {
				var headers5 HTTPHeaders
				headers5.Name = types.StringValue(headersItem2.Name)
				headers5.Value = types.StringValue(headersItem2.Value)
				if headersCount2+1 > len(r.StartupProbe.HTTP.Headers) {
					r.StartupProbe.HTTP.Headers = append(r.StartupProbe.HTTP.Headers, headers5)
				} else {
					r.StartupProbe.HTTP.Headers[headersCount2].Name = headers5.Name
					r.StartupProbe.HTTP.Headers[headersCount2].Value = headers5.Value
				}
			}
			r.StartupProbe.HTTP.Path = types.StringValue(resp.StartupProbe.HTTP.Path)
			r.StartupProbe.HTTP.Port = types.Int64Value(resp.StartupProbe.HTTP.Port)
			if resp.StartupProbe.HTTP.Scheme != nil {
				r.StartupProbe.HTTP.Scheme = types.StringValue(string(*resp.StartupProbe.HTTP.Scheme))
			} else {
				r.StartupProbe.HTTP.Scheme = types.StringNull()
			}
		}
		r.StartupProbe.InitialDelaySeconds = types.Int64Value(resp.StartupProbe.InitialDelaySeconds)
		r.StartupProbe.PeriodSeconds = types.Int64Value(resp.StartupProbe.PeriodSeconds)
		r.StartupProbe.SuccessThreshold = types.Int64Value(resp.StartupProbe.SuccessThreshold)
		if resp.StartupProbe.TCP == nil {
			r.StartupProbe.TCP = nil
		} else {
			r.StartupProbe.TCP = &ContainerGroupProbeTCP{}
			r.StartupProbe.TCP.Port = types.Int64Value(resp.StartupProbe.TCP.Port)
		}
		r.StartupProbe.TimeoutSeconds = types.Int64Value(resp.StartupProbe.TimeoutSeconds)
	}
	r.UpdateTime = types.StringValue(resp.UpdateTime.Format(time.RFC3339Nano))
}
