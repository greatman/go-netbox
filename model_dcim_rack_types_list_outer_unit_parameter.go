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

// DcimRackTypesListOuterUnitParameter the model 'DcimRackTypesListOuterUnitParameter'
type DcimRackTypesListOuterUnitParameter string

// List of dcim_rack_types_list_outer_unit_parameter
const (
	DCIMRACKTYPESLISTOUTERUNITPARAMETER_IN DcimRackTypesListOuterUnitParameter = "in"
	DCIMRACKTYPESLISTOUTERUNITPARAMETER_MM DcimRackTypesListOuterUnitParameter = "mm"
)

// All allowed values of DcimRackTypesListOuterUnitParameter enum
var AllowedDcimRackTypesListOuterUnitParameterEnumValues = []DcimRackTypesListOuterUnitParameter{
	"in",
	"mm",
}

func (v *DcimRackTypesListOuterUnitParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimRackTypesListOuterUnitParameter(value)
	for _, existing := range AllowedDcimRackTypesListOuterUnitParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimRackTypesListOuterUnitParameter", value)
}

// NewDcimRackTypesListOuterUnitParameterFromValue returns a pointer to a valid DcimRackTypesListOuterUnitParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimRackTypesListOuterUnitParameterFromValue(v string) (*DcimRackTypesListOuterUnitParameter, error) {
	ev := DcimRackTypesListOuterUnitParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimRackTypesListOuterUnitParameter: valid values are %v", v, AllowedDcimRackTypesListOuterUnitParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimRackTypesListOuterUnitParameter) IsValid() bool {
	for _, existing := range AllowedDcimRackTypesListOuterUnitParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_rack_types_list_outer_unit_parameter value
func (v DcimRackTypesListOuterUnitParameter) Ptr() *DcimRackTypesListOuterUnitParameter {
	return &v
}

type NullableDcimRackTypesListOuterUnitParameter struct {
	value *DcimRackTypesListOuterUnitParameter
	isSet bool
}

func (v NullableDcimRackTypesListOuterUnitParameter) Get() *DcimRackTypesListOuterUnitParameter {
	return v.value
}

func (v *NullableDcimRackTypesListOuterUnitParameter) Set(val *DcimRackTypesListOuterUnitParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimRackTypesListOuterUnitParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimRackTypesListOuterUnitParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimRackTypesListOuterUnitParameter(val *DcimRackTypesListOuterUnitParameter) *NullableDcimRackTypesListOuterUnitParameter {
	return &NullableDcimRackTypesListOuterUnitParameter{value: val, isSet: true}
}

func (v NullableDcimRackTypesListOuterUnitParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimRackTypesListOuterUnitParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

