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
)

// VirtualizationClusterTypesDeleteReader is a Reader for the VirtualizationClusterTypesDelete structure.
type VirtualizationClusterTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesDeleteNoContent creates a VirtualizationClusterTypesDeleteNoContent with default headers values
func NewVirtualizationClusterTypesDeleteNoContent() *VirtualizationClusterTypesDeleteNoContent {
	return &VirtualizationClusterTypesDeleteNoContent{}
}

/* VirtualizationClusterTypesDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterTypesDeleteNoContent virtualization cluster types delete no content
*/
type VirtualizationClusterTypesDeleteNoContent struct {
}

func (o *VirtualizationClusterTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesDeleteNoContent ", 204)
}

func (o *VirtualizationClusterTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationClusterTypesDeleteDefault creates a VirtualizationClusterTypesDeleteDefault with default headers values
func NewVirtualizationClusterTypesDeleteDefault(code int) *VirtualizationClusterTypesDeleteDefault {
	return &VirtualizationClusterTypesDeleteDefault{
		_statusCode: code,
	}
}

/* VirtualizationClusterTypesDeleteDefault describes a response with status code -1, with default header values.

VirtualizationClusterTypesDeleteDefault virtualization cluster types delete default
*/
type VirtualizationClusterTypesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types delete default response
func (o *VirtualizationClusterTypesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterTypesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/{id}/][%d] virtualization_cluster-types_delete default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationClusterTypesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
