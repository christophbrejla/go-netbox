// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceProtocol Protocol
//
// swagger:model serviceProtocol
type ServiceProtocol struct {

	// label
	// Required: true
	// Enum: [TCP UDP]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [tcp udp]
	Value *string `json:"value"`
}

// Validate validates this service protocol
func (m *ServiceProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceProtocolTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TCP","UDP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceProtocolTypeLabelPropEnum = append(serviceProtocolTypeLabelPropEnum, v)
	}
}

const (

	// ServiceProtocolLabelTCP captures enum value "TCP"
	ServiceProtocolLabelTCP string = "TCP"

	// ServiceProtocolLabelUDP captures enum value "UDP"
	ServiceProtocolLabelUDP string = "UDP"
)

// prop value enum
func (m *ServiceProtocol) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceProtocolTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceProtocol) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var serviceProtocolTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceProtocolTypeValuePropEnum = append(serviceProtocolTypeValuePropEnum, v)
	}
}

const (

	// ServiceProtocolValueTCP captures enum value "tcp"
	ServiceProtocolValueTCP string = "tcp"

	// ServiceProtocolValueUDP captures enum value "udp"
	ServiceProtocolValueUDP string = "udp"
)

// prop value enum
func (m *ServiceProtocol) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceProtocolTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceProtocol) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceProtocol) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceProtocol) UnmarshalBinary(b []byte) error {
	var res ServiceProtocol
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}