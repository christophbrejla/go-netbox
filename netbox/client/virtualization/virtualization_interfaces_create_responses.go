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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// VirtualizationInterfacesCreateReader is a Reader for the VirtualizationInterfacesCreate structure.
type VirtualizationInterfacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationInterfacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesCreateCreated creates a VirtualizationInterfacesCreateCreated with default headers values
func NewVirtualizationInterfacesCreateCreated() *VirtualizationInterfacesCreateCreated {
	return &VirtualizationInterfacesCreateCreated{}
}

/* VirtualizationInterfacesCreateCreated describes a response with status code 201, with default header values.

VirtualizationInterfacesCreateCreated virtualization interfaces create created
*/
type VirtualizationInterfacesCreateCreated struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/interfaces/][%d] virtualizationInterfacesCreateCreated  %+v", 201, o.Payload)
}
func (o *VirtualizationInterfacesCreateCreated) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationInterfacesCreateDefault creates a VirtualizationInterfacesCreateDefault with default headers values
func NewVirtualizationInterfacesCreateDefault(code int) *VirtualizationInterfacesCreateDefault {
	return &VirtualizationInterfacesCreateDefault{
		_statusCode: code,
	}
}

/* VirtualizationInterfacesCreateDefault describes a response with status code -1, with default header values.

VirtualizationInterfacesCreateDefault virtualization interfaces create default
*/
type VirtualizationInterfacesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization interfaces create default response
func (o *VirtualizationInterfacesCreateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /virtualization/interfaces/][%d] virtualization_interfaces_create default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationInterfacesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
