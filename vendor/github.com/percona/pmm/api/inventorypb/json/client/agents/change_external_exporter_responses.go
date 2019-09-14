// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeExternalExporterReader is a Reader for the ChangeExternalExporter structure.
type ChangeExternalExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeExternalExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeExternalExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeExternalExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeExternalExporterOK creates a ChangeExternalExporterOK with default headers values
func NewChangeExternalExporterOK() *ChangeExternalExporterOK {
	return &ChangeExternalExporterOK{}
}

/*ChangeExternalExporterOK handles this case with default header values.

A successful response.
*/
type ChangeExternalExporterOK struct {
	Payload *ChangeExternalExporterOKBody
}

func (o *ChangeExternalExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeExternalExporter][%d] changeExternalExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeExternalExporterOK) GetPayload() *ChangeExternalExporterOKBody {
	return o.Payload
}

func (o *ChangeExternalExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeExternalExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeExternalExporterDefault creates a ChangeExternalExporterDefault with default headers values
func NewChangeExternalExporterDefault(code int) *ChangeExternalExporterDefault {
	return &ChangeExternalExporterDefault{
		_statusCode: code,
	}
}

/*ChangeExternalExporterDefault handles this case with default header values.

An error response.
*/
type ChangeExternalExporterDefault struct {
	_statusCode int

	Payload *ChangeExternalExporterDefaultBody
}

// Code gets the status code for the change external exporter default response
func (o *ChangeExternalExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeExternalExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeExternalExporter][%d] ChangeExternalExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeExternalExporterDefault) GetPayload() *ChangeExternalExporterDefaultBody {
	return o.Payload
}

func (o *ChangeExternalExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeExternalExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeExternalExporterDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeExternalExporterDefaultBody
*/
type ChangeExternalExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change external exporter default body
func (o *ChangeExternalExporterDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeExternalExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeExternalExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeExternalExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeExternalExporterOKBody change external exporter OK body
swagger:model ChangeExternalExporterOKBody
*/
type ChangeExternalExporterOKBody struct {

	// external exporter
	ExternalExporter *ChangeExternalExporterOKBodyExternalExporter `json:"external_exporter,omitempty"`
}

// Validate validates this change external exporter OK body
func (o *ChangeExternalExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternalExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeExternalExporterOKBody) validateExternalExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalExporter) { // not required
		return nil
	}

	if o.ExternalExporter != nil {
		if err := o.ExternalExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeExternalExporterOk" + "." + "external_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeExternalExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeExternalExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeExternalExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeExternalExporterOKBodyExternalExporter ExternalExporter does not run on any Inventory Node.
swagger:model ChangeExternalExporterOKBodyExternalExporter
*/
type ChangeExternalExporterOKBodyExternalExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// URL for scraping metrics.
	MetricsURL string `json:"metrics_url,omitempty"`
}

// Validate validates this change external exporter OK body external exporter
func (o *ChangeExternalExporterOKBodyExternalExporter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeExternalExporterOKBodyExternalExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeExternalExporterOKBodyExternalExporter) UnmarshalBinary(b []byte) error {
	var res ChangeExternalExporterOKBodyExternalExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
