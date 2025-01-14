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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyContactsBulkDeleteReader is a Reader for the TenancyContactsBulkDelete structure.
type TenancyContactsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyContactsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsBulkDeleteNoContent creates a TenancyContactsBulkDeleteNoContent with default headers values
func NewTenancyContactsBulkDeleteNoContent() *TenancyContactsBulkDeleteNoContent {
	return &TenancyContactsBulkDeleteNoContent{}
}

/* TenancyContactsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyContactsBulkDeleteNoContent tenancy contacts bulk delete no content
*/
type TenancyContactsBulkDeleteNoContent struct {
}

func (o *TenancyContactsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contacts/][%d] tenancyContactsBulkDeleteNoContent ", 204)
}

func (o *TenancyContactsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyContactsBulkDeleteDefault creates a TenancyContactsBulkDeleteDefault with default headers values
func NewTenancyContactsBulkDeleteDefault(code int) *TenancyContactsBulkDeleteDefault {
	return &TenancyContactsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* TenancyContactsBulkDeleteDefault describes a response with status code -1, with default header values.

TenancyContactsBulkDeleteDefault tenancy contacts bulk delete default
*/
type TenancyContactsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contacts bulk delete default response
func (o *TenancyContactsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contacts/][%d] tenancy_contacts_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyContactsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
