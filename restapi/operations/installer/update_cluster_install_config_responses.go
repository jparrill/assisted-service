// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// UpdateClusterInstallConfigCreatedCode is the HTTP code returned for type UpdateClusterInstallConfigCreated
const UpdateClusterInstallConfigCreatedCode int = 201

/*UpdateClusterInstallConfigCreated Success.

swagger:response updateClusterInstallConfigCreated
*/
type UpdateClusterInstallConfigCreated struct {
}

// NewUpdateClusterInstallConfigCreated creates UpdateClusterInstallConfigCreated with default headers values
func NewUpdateClusterInstallConfigCreated() *UpdateClusterInstallConfigCreated {

	return &UpdateClusterInstallConfigCreated{}
}

// WriteResponse to the client
func (o *UpdateClusterInstallConfigCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// UpdateClusterInstallConfigBadRequestCode is the HTTP code returned for type UpdateClusterInstallConfigBadRequest
const UpdateClusterInstallConfigBadRequestCode int = 400

/*UpdateClusterInstallConfigBadRequest Error.

swagger:response updateClusterInstallConfigBadRequest
*/
type UpdateClusterInstallConfigBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallConfigBadRequest creates UpdateClusterInstallConfigBadRequest with default headers values
func NewUpdateClusterInstallConfigBadRequest() *UpdateClusterInstallConfigBadRequest {

	return &UpdateClusterInstallConfigBadRequest{}
}

// WithPayload adds the payload to the update cluster install config bad request response
func (o *UpdateClusterInstallConfigBadRequest) WithPayload(payload *models.Error) *UpdateClusterInstallConfigBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install config bad request response
func (o *UpdateClusterInstallConfigBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallConfigBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallConfigUnauthorizedCode is the HTTP code returned for type UpdateClusterInstallConfigUnauthorized
const UpdateClusterInstallConfigUnauthorizedCode int = 401

/*UpdateClusterInstallConfigUnauthorized Unauthorized.

swagger:response updateClusterInstallConfigUnauthorized
*/
type UpdateClusterInstallConfigUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUpdateClusterInstallConfigUnauthorized creates UpdateClusterInstallConfigUnauthorized with default headers values
func NewUpdateClusterInstallConfigUnauthorized() *UpdateClusterInstallConfigUnauthorized {

	return &UpdateClusterInstallConfigUnauthorized{}
}

// WithPayload adds the payload to the update cluster install config unauthorized response
func (o *UpdateClusterInstallConfigUnauthorized) WithPayload(payload *models.InfraError) *UpdateClusterInstallConfigUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install config unauthorized response
func (o *UpdateClusterInstallConfigUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallConfigUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallConfigForbiddenCode is the HTTP code returned for type UpdateClusterInstallConfigForbidden
const UpdateClusterInstallConfigForbiddenCode int = 403

/*UpdateClusterInstallConfigForbidden Forbidden.

swagger:response updateClusterInstallConfigForbidden
*/
type UpdateClusterInstallConfigForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUpdateClusterInstallConfigForbidden creates UpdateClusterInstallConfigForbidden with default headers values
func NewUpdateClusterInstallConfigForbidden() *UpdateClusterInstallConfigForbidden {

	return &UpdateClusterInstallConfigForbidden{}
}

// WithPayload adds the payload to the update cluster install config forbidden response
func (o *UpdateClusterInstallConfigForbidden) WithPayload(payload *models.InfraError) *UpdateClusterInstallConfigForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install config forbidden response
func (o *UpdateClusterInstallConfigForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallConfigForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallConfigNotFoundCode is the HTTP code returned for type UpdateClusterInstallConfigNotFound
const UpdateClusterInstallConfigNotFoundCode int = 404

/*UpdateClusterInstallConfigNotFound Error.

swagger:response updateClusterInstallConfigNotFound
*/
type UpdateClusterInstallConfigNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallConfigNotFound creates UpdateClusterInstallConfigNotFound with default headers values
func NewUpdateClusterInstallConfigNotFound() *UpdateClusterInstallConfigNotFound {

	return &UpdateClusterInstallConfigNotFound{}
}

// WithPayload adds the payload to the update cluster install config not found response
func (o *UpdateClusterInstallConfigNotFound) WithPayload(payload *models.Error) *UpdateClusterInstallConfigNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install config not found response
func (o *UpdateClusterInstallConfigNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallConfigMethodNotAllowedCode is the HTTP code returned for type UpdateClusterInstallConfigMethodNotAllowed
const UpdateClusterInstallConfigMethodNotAllowedCode int = 405

/*UpdateClusterInstallConfigMethodNotAllowed Method Not Allowed.

swagger:response updateClusterInstallConfigMethodNotAllowed
*/
type UpdateClusterInstallConfigMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallConfigMethodNotAllowed creates UpdateClusterInstallConfigMethodNotAllowed with default headers values
func NewUpdateClusterInstallConfigMethodNotAllowed() *UpdateClusterInstallConfigMethodNotAllowed {

	return &UpdateClusterInstallConfigMethodNotAllowed{}
}

// WithPayload adds the payload to the update cluster install config method not allowed response
func (o *UpdateClusterInstallConfigMethodNotAllowed) WithPayload(payload *models.Error) *UpdateClusterInstallConfigMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install config method not allowed response
func (o *UpdateClusterInstallConfigMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallConfigMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallConfigInternalServerErrorCode is the HTTP code returned for type UpdateClusterInstallConfigInternalServerError
const UpdateClusterInstallConfigInternalServerErrorCode int = 500

/*UpdateClusterInstallConfigInternalServerError Error.

swagger:response updateClusterInstallConfigInternalServerError
*/
type UpdateClusterInstallConfigInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallConfigInternalServerError creates UpdateClusterInstallConfigInternalServerError with default headers values
func NewUpdateClusterInstallConfigInternalServerError() *UpdateClusterInstallConfigInternalServerError {

	return &UpdateClusterInstallConfigInternalServerError{}
}

// WithPayload adds the payload to the update cluster install config internal server error response
func (o *UpdateClusterInstallConfigInternalServerError) WithPayload(payload *models.Error) *UpdateClusterInstallConfigInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install config internal server error response
func (o *UpdateClusterInstallConfigInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallConfigInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
