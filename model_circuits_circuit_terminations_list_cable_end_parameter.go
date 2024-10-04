/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// CircuitsCircuitTerminationsListCableEndParameter the model 'CircuitsCircuitTerminationsListCableEndParameter'
type CircuitsCircuitTerminationsListCableEndParameter string

// List of circuits_circuit_terminations_list_cable_end_parameter
const (
	CIRCUITSCIRCUITTERMINATIONSLISTCABLEENDPARAMETER_A CircuitsCircuitTerminationsListCableEndParameter = "A"
	CIRCUITSCIRCUITTERMINATIONSLISTCABLEENDPARAMETER_B CircuitsCircuitTerminationsListCableEndParameter = "B"
)

// All allowed values of CircuitsCircuitTerminationsListCableEndParameter enum
var AllowedCircuitsCircuitTerminationsListCableEndParameterEnumValues = []CircuitsCircuitTerminationsListCableEndParameter{
	"A",
	"B",
}

func (v *CircuitsCircuitTerminationsListCableEndParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CircuitsCircuitTerminationsListCableEndParameter(value)
	for _, existing := range AllowedCircuitsCircuitTerminationsListCableEndParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CircuitsCircuitTerminationsListCableEndParameter", value)
}

// NewCircuitsCircuitTerminationsListCableEndParameterFromValue returns a pointer to a valid CircuitsCircuitTerminationsListCableEndParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCircuitsCircuitTerminationsListCableEndParameterFromValue(v string) (*CircuitsCircuitTerminationsListCableEndParameter, error) {
	ev := CircuitsCircuitTerminationsListCableEndParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CircuitsCircuitTerminationsListCableEndParameter: valid values are %v", v, AllowedCircuitsCircuitTerminationsListCableEndParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CircuitsCircuitTerminationsListCableEndParameter) IsValid() bool {
	for _, existing := range AllowedCircuitsCircuitTerminationsListCableEndParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to circuits_circuit_terminations_list_cable_end_parameter value
func (v CircuitsCircuitTerminationsListCableEndParameter) Ptr() *CircuitsCircuitTerminationsListCableEndParameter {
	return &v
}

type NullableCircuitsCircuitTerminationsListCableEndParameter struct {
	value *CircuitsCircuitTerminationsListCableEndParameter
	isSet bool
}

func (v NullableCircuitsCircuitTerminationsListCableEndParameter) Get() *CircuitsCircuitTerminationsListCableEndParameter {
	return v.value
}

func (v *NullableCircuitsCircuitTerminationsListCableEndParameter) Set(val *CircuitsCircuitTerminationsListCableEndParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitsCircuitTerminationsListCableEndParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitsCircuitTerminationsListCableEndParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitsCircuitTerminationsListCableEndParameter(val *CircuitsCircuitTerminationsListCableEndParameter) *NullableCircuitsCircuitTerminationsListCableEndParameter {
	return &NullableCircuitsCircuitTerminationsListCableEndParameter{value: val, isSet: true}
}

func (v NullableCircuitsCircuitTerminationsListCableEndParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitsCircuitTerminationsListCableEndParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

