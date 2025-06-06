// Code generated by go-swagger; DO NOT EDIT.

package tlv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/spensercai/nfc_apdu_runner/tools/nfc_analysis_platform/pkg/wapi/codegen/models"
)

// ExtractTlvValuesOKCode is the HTTP code returned for type ExtractTlvValuesOK
const ExtractTlvValuesOKCode int = 200

/*
ExtractTlvValuesOK Successful operation

swagger:response extractTlvValuesOK
*/
type ExtractTlvValuesOK struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewExtractTlvValuesOK creates ExtractTlvValuesOK with default headers values
func NewExtractTlvValuesOK() *ExtractTlvValuesOK {

	return &ExtractTlvValuesOK{}
}

// WithPayload adds the payload to the extract tlv values o k response
func (o *ExtractTlvValuesOK) WithPayload(payload *models.APIResponse) *ExtractTlvValuesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the extract tlv values o k response
func (o *ExtractTlvValuesOK) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExtractTlvValuesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
