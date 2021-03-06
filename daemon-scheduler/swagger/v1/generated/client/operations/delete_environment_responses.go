// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteEnvironmentReader is a Reader for the DeleteEnvironment structure.
type DeleteEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEnvironmentOK creates a DeleteEnvironmentOK with default headers values
func NewDeleteEnvironmentOK() *DeleteEnvironmentOK {
	return &DeleteEnvironmentOK{}
}

/*DeleteEnvironmentOK handles this case with default header values.

OK
*/
type DeleteEnvironmentOK struct {
}

func (o *DeleteEnvironmentOK) Error() string {
	return fmt.Sprintf("[DELETE /environments/{name}][%d] deleteEnvironmentOK ", 200)
}

func (o *DeleteEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvironmentBadRequest creates a DeleteEnvironmentBadRequest with default headers values
func NewDeleteEnvironmentBadRequest() *DeleteEnvironmentBadRequest {
	return &DeleteEnvironmentBadRequest{}
}

/*DeleteEnvironmentBadRequest handles this case with default header values.

Bad Request
*/
type DeleteEnvironmentBadRequest struct {
	Payload string
}

func (o *DeleteEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /environments/{name}][%d] deleteEnvironmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
