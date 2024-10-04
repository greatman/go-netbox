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

// PatchedWritableRackRequestOuterUnit * `mm` - Millimeters * `in` - Inches
type PatchedWritableRackRequestOuterUnit string

// List of PatchedWritableRackRequest_outer_unit
const (
	PATCHEDWRITABLERACKREQUESTOUTERUNIT_MM PatchedWritableRackRequestOuterUnit = "mm"
	PATCHEDWRITABLERACKREQUESTOUTERUNIT_IN PatchedWritableRackRequestOuterUnit = "in"
	PATCHEDWRITABLERACKREQUESTOUTERUNIT_EMPTY PatchedWritableRackRequestOuterUnit = ""
)

// All allowed values of PatchedWritableRackRequestOuterUnit enum
var AllowedPatchedWritableRackRequestOuterUnitEnumValues = []PatchedWritableRackRequestOuterUnit{
	"mm",
	"in",
	"",
}

func (v *PatchedWritableRackRequestOuterUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableRackRequestOuterUnit(value)
	for _, existing := range AllowedPatchedWritableRackRequestOuterUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableRackRequestOuterUnit", value)
}

// NewPatchedWritableRackRequestOuterUnitFromValue returns a pointer to a valid PatchedWritableRackRequestOuterUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableRackRequestOuterUnitFromValue(v string) (*PatchedWritableRackRequestOuterUnit, error) {
	ev := PatchedWritableRackRequestOuterUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableRackRequestOuterUnit: valid values are %v", v, AllowedPatchedWritableRackRequestOuterUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableRackRequestOuterUnit) IsValid() bool {
	for _, existing := range AllowedPatchedWritableRackRequestOuterUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableRackRequest_outer_unit value
func (v PatchedWritableRackRequestOuterUnit) Ptr() *PatchedWritableRackRequestOuterUnit {
	return &v
}

type NullablePatchedWritableRackRequestOuterUnit struct {
	value *PatchedWritableRackRequestOuterUnit
	isSet bool
}

func (v NullablePatchedWritableRackRequestOuterUnit) Get() *PatchedWritableRackRequestOuterUnit {
	return v.value
}

func (v *NullablePatchedWritableRackRequestOuterUnit) Set(val *PatchedWritableRackRequestOuterUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequestOuterUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequestOuterUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequestOuterUnit(val *PatchedWritableRackRequestOuterUnit) *NullablePatchedWritableRackRequestOuterUnit {
	return &NullablePatchedWritableRackRequestOuterUnit{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequestOuterUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequestOuterUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

