// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVpnconnectionsPutReader is a Reader for the PcloudVpnconnectionsPut structure.
type PcloudVpnconnectionsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVpnconnectionsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPutOK creates a PcloudVpnconnectionsPutOK with default headers values
func NewPcloudVpnconnectionsPutOK() *PcloudVpnconnectionsPutOK {
	return &PcloudVpnconnectionsPutOK{}
}

/* PcloudVpnconnectionsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVpnconnectionsPutOK struct {
	Payload *models.VPNConnection
}

func (o *PcloudVpnconnectionsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutOK  %+v", 200, o.Payload)
}
func (o *PcloudVpnconnectionsPutOK) GetPayload() *models.VPNConnection {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VPNConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutBadRequest creates a PcloudVpnconnectionsPutBadRequest with default headers values
func NewPcloudVpnconnectionsPutBadRequest() *PcloudVpnconnectionsPutBadRequest {
	return &PcloudVpnconnectionsPutBadRequest{}
}

/* PcloudVpnconnectionsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudVpnconnectionsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutUnauthorized creates a PcloudVpnconnectionsPutUnauthorized with default headers values
func NewPcloudVpnconnectionsPutUnauthorized() *PcloudVpnconnectionsPutUnauthorized {
	return &PcloudVpnconnectionsPutUnauthorized{}
}

/* PcloudVpnconnectionsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPutUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudVpnconnectionsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutForbidden creates a PcloudVpnconnectionsPutForbidden with default headers values
func NewPcloudVpnconnectionsPutForbidden() *PcloudVpnconnectionsPutForbidden {
	return &PcloudVpnconnectionsPutForbidden{}
}

/* PcloudVpnconnectionsPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsPutForbidden struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutForbidden  %+v", 403, o.Payload)
}
func (o *PcloudVpnconnectionsPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutNotFound creates a PcloudVpnconnectionsPutNotFound with default headers values
func NewPcloudVpnconnectionsPutNotFound() *PcloudVpnconnectionsPutNotFound {
	return &PcloudVpnconnectionsPutNotFound{}
}

/* PcloudVpnconnectionsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsPutNotFound struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutNotFound  %+v", 404, o.Payload)
}
func (o *PcloudVpnconnectionsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutUnprocessableEntity creates a PcloudVpnconnectionsPutUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsPutUnprocessableEntity() *PcloudVpnconnectionsPutUnprocessableEntity {
	return &PcloudVpnconnectionsPutUnprocessableEntity{}
}

/* PcloudVpnconnectionsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudVpnconnectionsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutInternalServerError creates a PcloudVpnconnectionsPutInternalServerError with default headers values
func NewPcloudVpnconnectionsPutInternalServerError() *PcloudVpnconnectionsPutInternalServerError {
	return &PcloudVpnconnectionsPutInternalServerError{}
}

/* PcloudVpnconnectionsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudVpnconnectionsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}