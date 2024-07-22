package types

import "encoding/json"

type ServiceBindings struct {
	Items []ServiceBinding `json:"items" yaml:"items"`
}

type ServiceBinding struct {
	Common
	PagingSequence int64 `json:"-" yaml:"-"`

	Credentials json.RawMessage `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	ServiceInstanceID   string `json:"service_instance_id" yaml:"service_instance_id,omitempty"`
	ServiceInstanceName string `json:"service_instance_name,omitempty" yaml:"service_instance_name,omitempty"`

	SyslogDrainURL  string          `json:"syslog_drain_url,omitempty" yaml:"syslog_drain_url,omitempty"`
	RouteServiceURL string          `json:"route_service_url,omitempty"`
	VolumeMounts    json.RawMessage `json:"-" yaml:"-"`
	Endpoints       json.RawMessage `json:"-" yaml:"-"`
	Context         json.RawMessage `json:"context,omitempty" yaml:"context,omitempty"`
	Parameters      json.RawMessage `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	BindResource    json.RawMessage `json:"-" yaml:"-"`

	LastOperation *Operation `json:"last_operation,omitempty" yaml:"last_operation,omitempty"`
}