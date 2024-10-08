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

// VpnTunnelsListStatusIcParameterInner the model 'VpnTunnelsListStatusIcParameterInner'
type VpnTunnelsListStatusIcParameterInner string

// List of vpn_tunnels_list_status__ic_parameter_inner
const (
	VPNTUNNELSLISTSTATUSICPARAMETERINNER_ACTIVE VpnTunnelsListStatusIcParameterInner = "active"
	VPNTUNNELSLISTSTATUSICPARAMETERINNER_DISABLED VpnTunnelsListStatusIcParameterInner = "disabled"
	VPNTUNNELSLISTSTATUSICPARAMETERINNER_PLANNED VpnTunnelsListStatusIcParameterInner = "planned"
)

// All allowed values of VpnTunnelsListStatusIcParameterInner enum
var AllowedVpnTunnelsListStatusIcParameterInnerEnumValues = []VpnTunnelsListStatusIcParameterInner{
	"active",
	"disabled",
	"planned",
}

func (v *VpnTunnelsListStatusIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VpnTunnelsListStatusIcParameterInner(value)
	for _, existing := range AllowedVpnTunnelsListStatusIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VpnTunnelsListStatusIcParameterInner", value)
}

// NewVpnTunnelsListStatusIcParameterInnerFromValue returns a pointer to a valid VpnTunnelsListStatusIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVpnTunnelsListStatusIcParameterInnerFromValue(v string) (*VpnTunnelsListStatusIcParameterInner, error) {
	ev := VpnTunnelsListStatusIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VpnTunnelsListStatusIcParameterInner: valid values are %v", v, AllowedVpnTunnelsListStatusIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VpnTunnelsListStatusIcParameterInner) IsValid() bool {
	for _, existing := range AllowedVpnTunnelsListStatusIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vpn_tunnels_list_status__ic_parameter_inner value
func (v VpnTunnelsListStatusIcParameterInner) Ptr() *VpnTunnelsListStatusIcParameterInner {
	return &v
}

type NullableVpnTunnelsListStatusIcParameterInner struct {
	value *VpnTunnelsListStatusIcParameterInner
	isSet bool
}

func (v NullableVpnTunnelsListStatusIcParameterInner) Get() *VpnTunnelsListStatusIcParameterInner {
	return v.value
}

func (v *NullableVpnTunnelsListStatusIcParameterInner) Set(val *VpnTunnelsListStatusIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnTunnelsListStatusIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnTunnelsListStatusIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnTunnelsListStatusIcParameterInner(val *VpnTunnelsListStatusIcParameterInner) *NullableVpnTunnelsListStatusIcParameterInner {
	return &NullableVpnTunnelsListStatusIcParameterInner{value: val, isSet: true}
}

func (v NullableVpnTunnelsListStatusIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnTunnelsListStatusIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

