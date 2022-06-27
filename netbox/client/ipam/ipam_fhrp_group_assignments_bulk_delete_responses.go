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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamFhrpGroupAssignmentsBulkDeleteReader is a Reader for the IpamFhrpGroupAssignmentsBulkDelete structure.
type IpamFhrpGroupAssignmentsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamFhrpGroupAssignmentsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupAssignmentsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupAssignmentsBulkDeleteNoContent creates a IpamFhrpGroupAssignmentsBulkDeleteNoContent with default headers values
func NewIpamFhrpGroupAssignmentsBulkDeleteNoContent() *IpamFhrpGroupAssignmentsBulkDeleteNoContent {
	return &IpamFhrpGroupAssignmentsBulkDeleteNoContent{}
}

/* IpamFhrpGroupAssignmentsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamFhrpGroupAssignmentsBulkDeleteNoContent ipam fhrp group assignments bulk delete no content
*/
type IpamFhrpGroupAssignmentsBulkDeleteNoContent struct {
}

func (o *IpamFhrpGroupAssignmentsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsBulkDeleteNoContent ", 204)
}

func (o *IpamFhrpGroupAssignmentsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamFhrpGroupAssignmentsBulkDeleteDefault creates a IpamFhrpGroupAssignmentsBulkDeleteDefault with default headers values
func NewIpamFhrpGroupAssignmentsBulkDeleteDefault(code int) *IpamFhrpGroupAssignmentsBulkDeleteDefault {
	return &IpamFhrpGroupAssignmentsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* IpamFhrpGroupAssignmentsBulkDeleteDefault describes a response with status code -1, with default header values.

IpamFhrpGroupAssignmentsBulkDeleteDefault ipam fhrp group assignments bulk delete default
*/
type IpamFhrpGroupAssignmentsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp group assignments bulk delete default response
func (o *IpamFhrpGroupAssignmentsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamFhrpGroupAssignmentsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/][%d] ipam_fhrp-group-assignments_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamFhrpGroupAssignmentsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}