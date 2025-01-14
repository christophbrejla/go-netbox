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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomFieldsDeleteReader is a Reader for the ExtrasCustomFieldsDelete structure.
type ExtrasCustomFieldsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomFieldsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsDeleteNoContent creates a ExtrasCustomFieldsDeleteNoContent with default headers values
func NewExtrasCustomFieldsDeleteNoContent() *ExtrasCustomFieldsDeleteNoContent {
	return &ExtrasCustomFieldsDeleteNoContent{}
}

/* ExtrasCustomFieldsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomFieldsDeleteNoContent extras custom fields delete no content
*/
type ExtrasCustomFieldsDeleteNoContent struct {
}

func (o *ExtrasCustomFieldsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/{id}/][%d] extrasCustomFieldsDeleteNoContent ", 204)
}

func (o *ExtrasCustomFieldsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasCustomFieldsDeleteDefault creates a ExtrasCustomFieldsDeleteDefault with default headers values
func NewExtrasCustomFieldsDeleteDefault(code int) *ExtrasCustomFieldsDeleteDefault {
	return &ExtrasCustomFieldsDeleteDefault{
		_statusCode: code,
	}
}

/* ExtrasCustomFieldsDeleteDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldsDeleteDefault extras custom fields delete default
*/
type ExtrasCustomFieldsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom fields delete default response
func (o *ExtrasCustomFieldsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomFieldsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/{id}/][%d] extras_custom-fields_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasCustomFieldsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
