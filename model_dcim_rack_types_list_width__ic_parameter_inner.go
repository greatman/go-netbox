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

// DcimRackTypesListWidthIcParameterInner the model 'DcimRackTypesListWidthIcParameterInner'
type DcimRackTypesListWidthIcParameterInner int64

// List of dcim_rack_types_list_width__ic_parameter_inner
const (
	DCIMRACKTYPESLISTWIDTHICPARAMETERINNER__10 DcimRackTypesListWidthIcParameterInner = 10
	DCIMRACKTYPESLISTWIDTHICPARAMETERINNER__19 DcimRackTypesListWidthIcParameterInner = 19
	DCIMRACKTYPESLISTWIDTHICPARAMETERINNER__21 DcimRackTypesListWidthIcParameterInner = 21
	DCIMRACKTYPESLISTWIDTHICPARAMETERINNER__23 DcimRackTypesListWidthIcParameterInner = 23
)

// All allowed values of DcimRackTypesListWidthIcParameterInner enum
var AllowedDcimRackTypesListWidthIcParameterInnerEnumValues = []DcimRackTypesListWidthIcParameterInner{
	10,
	19,
	21,
	23,
}

func (v *DcimRackTypesListWidthIcParameterInner) UnmarshalJSON(src []byte) error {
	var value int64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimRackTypesListWidthIcParameterInner(value)
	for _, existing := range AllowedDcimRackTypesListWidthIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimRackTypesListWidthIcParameterInner", value)
}

// NewDcimRackTypesListWidthIcParameterInnerFromValue returns a pointer to a valid DcimRackTypesListWidthIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimRackTypesListWidthIcParameterInnerFromValue(v int64) (*DcimRackTypesListWidthIcParameterInner, error) {
	ev := DcimRackTypesListWidthIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimRackTypesListWidthIcParameterInner: valid values are %v", v, AllowedDcimRackTypesListWidthIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimRackTypesListWidthIcParameterInner) IsValid() bool {
	for _, existing := range AllowedDcimRackTypesListWidthIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_rack_types_list_width__ic_parameter_inner value
func (v DcimRackTypesListWidthIcParameterInner) Ptr() *DcimRackTypesListWidthIcParameterInner {
	return &v
}

type NullableDcimRackTypesListWidthIcParameterInner struct {
	value *DcimRackTypesListWidthIcParameterInner
	isSet bool
}

func (v NullableDcimRackTypesListWidthIcParameterInner) Get() *DcimRackTypesListWidthIcParameterInner {
	return v.value
}

func (v *NullableDcimRackTypesListWidthIcParameterInner) Set(val *DcimRackTypesListWidthIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimRackTypesListWidthIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimRackTypesListWidthIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimRackTypesListWidthIcParameterInner(val *DcimRackTypesListWidthIcParameterInner) *NullableDcimRackTypesListWidthIcParameterInner {
	return &NullableDcimRackTypesListWidthIcParameterInner{value: val, isSet: true}
}

func (v NullableDcimRackTypesListWidthIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimRackTypesListWidthIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

