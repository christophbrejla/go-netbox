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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimModuleBaysBulkUpdateReader is a Reader for the DcimModuleBaysBulkUpdate structure.
type DcimModuleBaysBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBaysBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysBulkUpdateOK creates a DcimModuleBaysBulkUpdateOK with default headers values
func NewDcimModuleBaysBulkUpdateOK() *DcimModuleBaysBulkUpdateOK {
	return &DcimModuleBaysBulkUpdateOK{}
}

/* DcimModuleBaysBulkUpdateOK describes a response with status code 200, with default header values.

DcimModuleBaysBulkUpdateOK dcim module bays bulk update o k
*/
type DcimModuleBaysBulkUpdateOK struct {
	Payload *models.ModuleBay
}

func (o *DcimModuleBaysBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcimModuleBaysBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimModuleBaysBulkUpdateOK) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysBulkUpdateDefault creates a DcimModuleBaysBulkUpdateDefault with default headers values
func NewDcimModuleBaysBulkUpdateDefault(code int) *DcimModuleBaysBulkUpdateDefault {
	return &DcimModuleBaysBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimModuleBaysBulkUpdateDefault describes a response with status code -1, with default header values.

DcimModuleBaysBulkUpdateDefault dcim module bays bulk update default
*/
type DcimModuleBaysBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bays bulk update default response
func (o *DcimModuleBaysBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleBaysBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcim_module-bays_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimModuleBaysBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}