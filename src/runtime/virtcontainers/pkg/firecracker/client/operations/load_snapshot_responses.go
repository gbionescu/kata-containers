// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// LoadSnapshotReader is a Reader for the LoadSnapshot structure.
type LoadSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoadSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewLoadSnapshotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLoadSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewLoadSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoadSnapshotNoContent creates a LoadSnapshotNoContent with default headers values
func NewLoadSnapshotNoContent() *LoadSnapshotNoContent {
	return &LoadSnapshotNoContent{}
}

/*LoadSnapshotNoContent handles this case with default header values.

Snapshot loaded
*/
type LoadSnapshotNoContent struct {
}

func (o *LoadSnapshotNoContent) Error() string {
	return fmt.Sprintf("[PUT /snapshot/load][%d] loadSnapshotNoContent ", 204)
}

func (o *LoadSnapshotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoadSnapshotBadRequest creates a LoadSnapshotBadRequest with default headers values
func NewLoadSnapshotBadRequest() *LoadSnapshotBadRequest {
	return &LoadSnapshotBadRequest{}
}

/*LoadSnapshotBadRequest handles this case with default header values.

Snapshot cannot be loaded due to bad input
*/
type LoadSnapshotBadRequest struct {
	Payload *models.Error
}

func (o *LoadSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /snapshot/load][%d] loadSnapshotBadRequest  %+v", 400, o.Payload)
}

func (o *LoadSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoadSnapshotDefault creates a LoadSnapshotDefault with default headers values
func NewLoadSnapshotDefault(code int) *LoadSnapshotDefault {
	return &LoadSnapshotDefault{
		_statusCode: code,
	}
}

/*LoadSnapshotDefault handles this case with default header values.

Internal server error
*/
type LoadSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the load snapshot default response
func (o *LoadSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *LoadSnapshotDefault) Error() string {
	return fmt.Sprintf("[PUT /snapshot/load][%d] loadSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *LoadSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}