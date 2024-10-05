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

// AuthenticationType3 * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise
type AuthenticationType3 string

// List of Authentication_type_3
const (
	AUTHENTICATIONTYPE3_OPEN AuthenticationType3 = "open"
	AUTHENTICATIONTYPE3_WEP AuthenticationType3 = "wep"
	AUTHENTICATIONTYPE3_WPA_PERSONAL AuthenticationType3 = "wpa-personal"
	AUTHENTICATIONTYPE3_WPA_ENTERPRISE AuthenticationType3 = "wpa-enterprise"
	AUTHENTICATIONTYPE3_EMPTY AuthenticationType3 = ""
)

// All allowed values of AuthenticationType3 enum
var AllowedAuthenticationType3EnumValues = []AuthenticationType3{
	"open",
	"wep",
	"wpa-personal",
	"wpa-enterprise",
	"",
}

func (v *AuthenticationType3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticationType3(value)
	for _, existing := range AllowedAuthenticationType3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticationType3", value)
}

// NewAuthenticationType3FromValue returns a pointer to a valid AuthenticationType3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationType3FromValue(v string) (*AuthenticationType3, error) {
	ev := AuthenticationType3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticationType3: valid values are %v", v, AllowedAuthenticationType3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticationType3) IsValid() bool {
	for _, existing := range AllowedAuthenticationType3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Authentication_type_3 value
func (v AuthenticationType3) Ptr() *AuthenticationType3 {
	return &v
}

type NullableAuthenticationType3 struct {
	value *AuthenticationType3
	isSet bool
}

func (v NullableAuthenticationType3) Get() *AuthenticationType3 {
	return v.value
}

func (v *NullableAuthenticationType3) Set(val *AuthenticationType3) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationType3) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationType3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationType3(val *AuthenticationType3) *NullableAuthenticationType3 {
	return &NullableAuthenticationType3{value: val, isSet: true}
}

func (v NullableAuthenticationType3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationType3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

