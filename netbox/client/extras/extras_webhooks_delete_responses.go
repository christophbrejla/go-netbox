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

// ExtrasWebhooksDeleteReader is a Reader for the ExtrasWebhooksDelete structure.
type ExtrasWebhooksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasWebhooksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasWebhooksDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasWebhooksDeleteNoContent creates a ExtrasWebhooksDeleteNoContent with default headers values
func NewExtrasWebhooksDeleteNoContent() *ExtrasWebhooksDeleteNoContent {
	return &ExtrasWebhooksDeleteNoContent{}
}

/* ExtrasWebhooksDeleteNoContent describes a response with status code 204, with default header values.

ExtrasWebhooksDeleteNoContent extras webhooks delete no content
*/
type ExtrasWebhooksDeleteNoContent struct {
}

func (o *ExtrasWebhooksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extrasWebhooksDeleteNoContent ", 204)
}

func (o *ExtrasWebhooksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasWebhooksDeleteDefault creates a ExtrasWebhooksDeleteDefault with default headers values
func NewExtrasWebhooksDeleteDefault(code int) *ExtrasWebhooksDeleteDefault {
	return &ExtrasWebhooksDeleteDefault{
		_statusCode: code,
	}
}

/* ExtrasWebhooksDeleteDefault describes a response with status code -1, with default header values.

ExtrasWebhooksDeleteDefault extras webhooks delete default
*/
type ExtrasWebhooksDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras webhooks delete default response
func (o *ExtrasWebhooksDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasWebhooksDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extras_webhooks_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasWebhooksDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
