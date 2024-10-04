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

// InterfaceTemplateRequestRfRole * `ap` - Access point * `station` - Station
type InterfaceTemplateRequestRfRole string

// List of InterfaceTemplateRequest_rf_role
const (
	INTERFACETEMPLATEREQUESTRFROLE_AP InterfaceTemplateRequestRfRole = "ap"
	INTERFACETEMPLATEREQUESTRFROLE_STATION InterfaceTemplateRequestRfRole = "station"
	INTERFACETEMPLATEREQUESTRFROLE_EMPTY InterfaceTemplateRequestRfRole = ""
)

// All allowed values of InterfaceTemplateRequestRfRole enum
var AllowedInterfaceTemplateRequestRfRoleEnumValues = []InterfaceTemplateRequestRfRole{
	"ap",
	"station",
	"",
}

func (v *InterfaceTemplateRequestRfRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceTemplateRequestRfRole(value)
	for _, existing := range AllowedInterfaceTemplateRequestRfRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceTemplateRequestRfRole", value)
}

// NewInterfaceTemplateRequestRfRoleFromValue returns a pointer to a valid InterfaceTemplateRequestRfRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceTemplateRequestRfRoleFromValue(v string) (*InterfaceTemplateRequestRfRole, error) {
	ev := InterfaceTemplateRequestRfRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceTemplateRequestRfRole: valid values are %v", v, AllowedInterfaceTemplateRequestRfRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceTemplateRequestRfRole) IsValid() bool {
	for _, existing := range AllowedInterfaceTemplateRequestRfRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceTemplateRequest_rf_role value
func (v InterfaceTemplateRequestRfRole) Ptr() *InterfaceTemplateRequestRfRole {
	return &v
}

type NullableInterfaceTemplateRequestRfRole struct {
	value *InterfaceTemplateRequestRfRole
	isSet bool
}

func (v NullableInterfaceTemplateRequestRfRole) Get() *InterfaceTemplateRequestRfRole {
	return v.value
}

func (v *NullableInterfaceTemplateRequestRfRole) Set(val *InterfaceTemplateRequestRfRole) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplateRequestRfRole) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplateRequestRfRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplateRequestRfRole(val *InterfaceTemplateRequestRfRole) *NullableInterfaceTemplateRequestRfRole {
	return &NullableInterfaceTemplateRequestRfRole{value: val, isSet: true}
}

func (v NullableInterfaceTemplateRequestRfRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplateRequestRfRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

