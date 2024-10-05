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

// ModuleStatusLabel the model 'ModuleStatusLabel'
type ModuleStatusLabel string

// List of Module_status_label
const (
	MODULESTATUSLABEL_OFFLINE ModuleStatusLabel = "Offline"
	MODULESTATUSLABEL_ACTIVE ModuleStatusLabel = "Active"
	MODULESTATUSLABEL_PLANNED ModuleStatusLabel = "Planned"
	MODULESTATUSLABEL_STAGED ModuleStatusLabel = "Staged"
	MODULESTATUSLABEL_FAILED ModuleStatusLabel = "Failed"
	MODULESTATUSLABEL_DECOMMISSIONING ModuleStatusLabel = "Decommissioning"
)

// All allowed values of ModuleStatusLabel enum
var AllowedModuleStatusLabelEnumValues = []ModuleStatusLabel{
	"Offline",
	"Active",
	"Planned",
	"Staged",
	"Failed",
	"Decommissioning",
}

func (v *ModuleStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModuleStatusLabel(value)
	for _, existing := range AllowedModuleStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModuleStatusLabel", value)
}

// NewModuleStatusLabelFromValue returns a pointer to a valid ModuleStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModuleStatusLabelFromValue(v string) (*ModuleStatusLabel, error) {
	ev := ModuleStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModuleStatusLabel: valid values are %v", v, AllowedModuleStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModuleStatusLabel) IsValid() bool {
	for _, existing := range AllowedModuleStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Module_status_label value
func (v ModuleStatusLabel) Ptr() *ModuleStatusLabel {
	return &v
}

type NullableModuleStatusLabel struct {
	value *ModuleStatusLabel
	isSet bool
}

func (v NullableModuleStatusLabel) Get() *ModuleStatusLabel {
	return v.value
}

func (v *NullableModuleStatusLabel) Set(val *ModuleStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatusLabel(val *ModuleStatusLabel) *NullableModuleStatusLabel {
	return &NullableModuleStatusLabel{value: val, isSet: true}
}

func (v NullableModuleStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

