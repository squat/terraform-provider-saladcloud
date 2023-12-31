// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NewRelic struct {
	Host         string `json:"host"`
	IngestionKey string `json:"ingestion_key"`
}

func (o *NewRelic) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *NewRelic) GetIngestionKey() string {
	if o == nil {
		return ""
	}
	return o.IngestionKey
}

type Splunk struct {
	Host  string `json:"host"`
	Token string `json:"token"`
}

func (o *Splunk) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *Splunk) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type TCP struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

func (o *TCP) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *TCP) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

type Logging struct {
	NewRelic *NewRelic `json:"new_relic,omitempty"`
	Splunk   *Splunk   `json:"splunk,omitempty"`
	TCP      *TCP      `json:"tcp,omitempty"`
}

func (o *Logging) GetNewRelic() *NewRelic {
	if o == nil {
		return nil
	}
	return o.NewRelic
}

func (o *Logging) GetSplunk() *Splunk {
	if o == nil {
		return nil
	}
	return o.Splunk
}

func (o *Logging) GetTCP() *TCP {
	if o == nil {
		return nil
	}
	return o.TCP
}

// Container - Represents a container
type Container struct {
	Command              []string          `json:"command"`
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`
	Image                string            `json:"image"`
	Logging              *Logging          `json:"logging,omitempty"`
	// Represents a container resource requirements
	Resources ContainerResourceRequirements `json:"resources"`
}

func (o *Container) GetCommand() []string {
	if o == nil {
		return []string{}
	}
	return o.Command
}

func (o *Container) GetEnvironmentVariables() map[string]string {
	if o == nil {
		return nil
	}
	return o.EnvironmentVariables
}

func (o *Container) GetImage() string {
	if o == nil {
		return ""
	}
	return o.Image
}

func (o *Container) GetLogging() *Logging {
	if o == nil {
		return nil
	}
	return o.Logging
}

func (o *Container) GetResources() ContainerResourceRequirements {
	if o == nil {
		return ContainerResourceRequirements{}
	}
	return o.Resources
}
