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

// VpnIkePoliciesListVersionIcParameterInner the model 'VpnIkePoliciesListVersionIcParameterInner'
type VpnIkePoliciesListVersionIcParameterInner int32

// List of vpn_ike_policies_list_version__ic_parameter_inner
const (
	VPNIKEPOLICIESLISTVERSIONICPARAMETERINNER__1 VpnIkePoliciesListVersionIcParameterInner = 1
	VPNIKEPOLICIESLISTVERSIONICPARAMETERINNER__2 VpnIkePoliciesListVersionIcParameterInner = 2
)

// All allowed values of VpnIkePoliciesListVersionIcParameterInner enum
var AllowedVpnIkePoliciesListVersionIcParameterInnerEnumValues = []VpnIkePoliciesListVersionIcParameterInner{
	1,
	2,
}

func (v *VpnIkePoliciesListVersionIcParameterInner) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VpnIkePoliciesListVersionIcParameterInner(value)
	for _, existing := range AllowedVpnIkePoliciesListVersionIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VpnIkePoliciesListVersionIcParameterInner", value)
}

// NewVpnIkePoliciesListVersionIcParameterInnerFromValue returns a pointer to a valid VpnIkePoliciesListVersionIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVpnIkePoliciesListVersionIcParameterInnerFromValue(v int32) (*VpnIkePoliciesListVersionIcParameterInner, error) {
	ev := VpnIkePoliciesListVersionIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VpnIkePoliciesListVersionIcParameterInner: valid values are %v", v, AllowedVpnIkePoliciesListVersionIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VpnIkePoliciesListVersionIcParameterInner) IsValid() bool {
	for _, existing := range AllowedVpnIkePoliciesListVersionIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vpn_ike_policies_list_version__ic_parameter_inner value
func (v VpnIkePoliciesListVersionIcParameterInner) Ptr() *VpnIkePoliciesListVersionIcParameterInner {
	return &v
}

type NullableVpnIkePoliciesListVersionIcParameterInner struct {
	value *VpnIkePoliciesListVersionIcParameterInner
	isSet bool
}

func (v NullableVpnIkePoliciesListVersionIcParameterInner) Get() *VpnIkePoliciesListVersionIcParameterInner {
	return v.value
}

func (v *NullableVpnIkePoliciesListVersionIcParameterInner) Set(val *VpnIkePoliciesListVersionIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnIkePoliciesListVersionIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnIkePoliciesListVersionIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnIkePoliciesListVersionIcParameterInner(val *VpnIkePoliciesListVersionIcParameterInner) *NullableVpnIkePoliciesListVersionIcParameterInner {
	return &NullableVpnIkePoliciesListVersionIcParameterInner{value: val, isSet: true}
}

func (v NullableVpnIkePoliciesListVersionIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnIkePoliciesListVersionIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

