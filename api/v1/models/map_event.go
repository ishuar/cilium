// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MapEvent Event on Map
//
// swagger:model MapEvent
type MapEvent struct {

	// Action type for event
	// Enum: ["update","delete"]
	Action string `json:"action,omitempty"`

	// Desired action to be performed after this event
	// Enum: ["ok","insert","delete"]
	DesiredAction string `json:"desired-action,omitempty"`

	// Map key on which the event occured
	Key string `json:"key,omitempty"`

	// Last error seen while performing desired action
	LastError string `json:"last-error,omitempty"`

	// Timestamp when the event occurred
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// Map value on which the event occured
	Value string `json:"value,omitempty"`
}

// Validate validates this map event
func (m *MapEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mapEventTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["update","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mapEventTypeActionPropEnum = append(mapEventTypeActionPropEnum, v)
	}
}

const (

	// MapEventActionUpdate captures enum value "update"
	MapEventActionUpdate string = "update"

	// MapEventActionDelete captures enum value "delete"
	MapEventActionDelete string = "delete"
)

// prop value enum
func (m *MapEvent) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mapEventTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MapEvent) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var mapEventTypeDesiredActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","insert","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mapEventTypeDesiredActionPropEnum = append(mapEventTypeDesiredActionPropEnum, v)
	}
}

const (

	// MapEventDesiredActionOk captures enum value "ok"
	MapEventDesiredActionOk string = "ok"

	// MapEventDesiredActionInsert captures enum value "insert"
	MapEventDesiredActionInsert string = "insert"

	// MapEventDesiredActionDelete captures enum value "delete"
	MapEventDesiredActionDelete string = "delete"
)

// prop value enum
func (m *MapEvent) validateDesiredActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mapEventTypeDesiredActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MapEvent) validateDesiredAction(formats strfmt.Registry) error {
	if swag.IsZero(m.DesiredAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDesiredActionEnum("desired-action", "body", m.DesiredAction); err != nil {
		return err
	}

	return nil
}

func (m *MapEvent) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this map event based on context it is used
func (m *MapEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MapEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MapEvent) UnmarshalBinary(b []byte) error {
	var res MapEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
