// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/treeverse/lakefs/api/gen/models"
)

// ListBranchesReader is a Reader for the ListBranches structure.
type ListBranchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBranchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBranchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListBranchesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListBranchesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBranchesOK creates a ListBranchesOK with default headers values
func NewListBranchesOK() *ListBranchesOK {
	return &ListBranchesOK{}
}

/*ListBranchesOK handles this case with default header values.

branch list
*/
type ListBranchesOK struct {
	Payload []*models.Refspec
}

func (o *ListBranchesOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/branches][%d] listBranchesOK  %+v", 200, o.Payload)
}

func (o *ListBranchesOK) GetPayload() []*models.Refspec {
	return o.Payload
}

func (o *ListBranchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBranchesUnauthorized creates a ListBranchesUnauthorized with default headers values
func NewListBranchesUnauthorized() *ListBranchesUnauthorized {
	return &ListBranchesUnauthorized{}
}

/*ListBranchesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListBranchesUnauthorized struct {
	Payload *models.Error
}

func (o *ListBranchesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/branches][%d] listBranchesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListBranchesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListBranchesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBranchesDefault creates a ListBranchesDefault with default headers values
func NewListBranchesDefault(code int) *ListBranchesDefault {
	return &ListBranchesDefault{
		_statusCode: code,
	}
}

/*ListBranchesDefault handles this case with default header values.

generic error response
*/
type ListBranchesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list branches default response
func (o *ListBranchesDefault) Code() int {
	return o._statusCode
}

func (o *ListBranchesDefault) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/branches][%d] listBranches default  %+v", o._statusCode, o.Payload)
}

func (o *ListBranchesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListBranchesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}