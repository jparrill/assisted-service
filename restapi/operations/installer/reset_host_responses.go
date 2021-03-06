// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// ResetHostOKCode is the HTTP code returned for type ResetHostOK
const ResetHostOKCode int = 200

/*ResetHostOK Success.

swagger:response resetHostOK
*/
type ResetHostOK struct {

	/*
	  In: Body
	*/
	Payload *models.Host `json:"body,omitempty"`
}

// NewResetHostOK creates ResetHostOK with default headers values
func NewResetHostOK() *ResetHostOK {

	return &ResetHostOK{}
}

// WithPayload adds the payload to the reset host o k response
func (o *ResetHostOK) WithPayload(payload *models.Host) *ResetHostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reset host o k response
func (o *ResetHostOK) SetPayload(payload *models.Host) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResetHostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ResetHostUnauthorizedCode is the HTTP code returned for type ResetHostUnauthorized
const ResetHostUnauthorizedCode int = 401

/*ResetHostUnauthorized Unauthorized.

swagger:response resetHostUnauthorized
*/
type ResetHostUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewResetHostUnauthorized creates ResetHostUnauthorized with default headers values
func NewResetHostUnauthorized() *ResetHostUnauthorized {

	return &ResetHostUnauthorized{}
}

// WithPayload adds the payload to the reset host unauthorized response
func (o *ResetHostUnauthorized) WithPayload(payload *models.InfraError) *ResetHostUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reset host unauthorized response
func (o *ResetHostUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResetHostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ResetHostForbiddenCode is the HTTP code returned for type ResetHostForbidden
const ResetHostForbiddenCode int = 403

/*ResetHostForbidden Forbidden.

swagger:response resetHostForbidden
*/
type ResetHostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewResetHostForbidden creates ResetHostForbidden with default headers values
func NewResetHostForbidden() *ResetHostForbidden {

	return &ResetHostForbidden{}
}

// WithPayload adds the payload to the reset host forbidden response
func (o *ResetHostForbidden) WithPayload(payload *models.InfraError) *ResetHostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reset host forbidden response
func (o *ResetHostForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResetHostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ResetHostNotFoundCode is the HTTP code returned for type ResetHostNotFound
const ResetHostNotFoundCode int = 404

/*ResetHostNotFound Error.

swagger:response resetHostNotFound
*/
type ResetHostNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewResetHostNotFound creates ResetHostNotFound with default headers values
func NewResetHostNotFound() *ResetHostNotFound {

	return &ResetHostNotFound{}
}

// WithPayload adds the payload to the reset host not found response
func (o *ResetHostNotFound) WithPayload(payload *models.Error) *ResetHostNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reset host not found response
func (o *ResetHostNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResetHostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ResetHostConflictCode is the HTTP code returned for type ResetHostConflict
const ResetHostConflictCode int = 409

/*ResetHostConflict Error.

swagger:response resetHostConflict
*/
type ResetHostConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewResetHostConflict creates ResetHostConflict with default headers values
func NewResetHostConflict() *ResetHostConflict {

	return &ResetHostConflict{}
}

// WithPayload adds the payload to the reset host conflict response
func (o *ResetHostConflict) WithPayload(payload *models.Error) *ResetHostConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reset host conflict response
func (o *ResetHostConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResetHostConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ResetHostInternalServerErrorCode is the HTTP code returned for type ResetHostInternalServerError
const ResetHostInternalServerErrorCode int = 500

/*ResetHostInternalServerError Error.

swagger:response resetHostInternalServerError
*/
type ResetHostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewResetHostInternalServerError creates ResetHostInternalServerError with default headers values
func NewResetHostInternalServerError() *ResetHostInternalServerError {

	return &ResetHostInternalServerError{}
}

// WithPayload adds the payload to the reset host internal server error response
func (o *ResetHostInternalServerError) WithPayload(payload *models.Error) *ResetHostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reset host internal server error response
func (o *ResetHostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResetHostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
