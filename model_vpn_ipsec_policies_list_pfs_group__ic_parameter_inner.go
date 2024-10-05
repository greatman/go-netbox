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

// VpnIpsecPoliciesListPfsGroupIcParameterInner the model 'VpnIpsecPoliciesListPfsGroupIcParameterInner'
type VpnIpsecPoliciesListPfsGroupIcParameterInner int32

// List of vpn_ipsec_policies_list_pfs_group__ic_parameter_inner
const (
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__1 VpnIpsecPoliciesListPfsGroupIcParameterInner = 1
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__14 VpnIpsecPoliciesListPfsGroupIcParameterInner = 14
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__15 VpnIpsecPoliciesListPfsGroupIcParameterInner = 15
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__16 VpnIpsecPoliciesListPfsGroupIcParameterInner = 16
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__17 VpnIpsecPoliciesListPfsGroupIcParameterInner = 17
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__18 VpnIpsecPoliciesListPfsGroupIcParameterInner = 18
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__19 VpnIpsecPoliciesListPfsGroupIcParameterInner = 19
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__2 VpnIpsecPoliciesListPfsGroupIcParameterInner = 2
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__20 VpnIpsecPoliciesListPfsGroupIcParameterInner = 20
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__21 VpnIpsecPoliciesListPfsGroupIcParameterInner = 21
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__22 VpnIpsecPoliciesListPfsGroupIcParameterInner = 22
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__23 VpnIpsecPoliciesListPfsGroupIcParameterInner = 23
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__24 VpnIpsecPoliciesListPfsGroupIcParameterInner = 24
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__25 VpnIpsecPoliciesListPfsGroupIcParameterInner = 25
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__26 VpnIpsecPoliciesListPfsGroupIcParameterInner = 26
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__27 VpnIpsecPoliciesListPfsGroupIcParameterInner = 27
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__28 VpnIpsecPoliciesListPfsGroupIcParameterInner = 28
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__29 VpnIpsecPoliciesListPfsGroupIcParameterInner = 29
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__30 VpnIpsecPoliciesListPfsGroupIcParameterInner = 30
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__31 VpnIpsecPoliciesListPfsGroupIcParameterInner = 31
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__32 VpnIpsecPoliciesListPfsGroupIcParameterInner = 32
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__33 VpnIpsecPoliciesListPfsGroupIcParameterInner = 33
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__34 VpnIpsecPoliciesListPfsGroupIcParameterInner = 34
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__5 VpnIpsecPoliciesListPfsGroupIcParameterInner = 5
	VPNIPSECPOLICIESLISTPFSGROUPICPARAMETERINNER__null VpnIpsecPoliciesListPfsGroupIcParameterInner = null
)

// All allowed values of VpnIpsecPoliciesListPfsGroupIcParameterInner enum
var AllowedVpnIpsecPoliciesListPfsGroupIcParameterInnerEnumValues = []VpnIpsecPoliciesListPfsGroupIcParameterInner{
	1,
	14,
	15,
	16,
	17,
	18,
	19,
	2,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
	5,
	null,
}

func (v *VpnIpsecPoliciesListPfsGroupIcParameterInner) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VpnIpsecPoliciesListPfsGroupIcParameterInner(value)
	for _, existing := range AllowedVpnIpsecPoliciesListPfsGroupIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VpnIpsecPoliciesListPfsGroupIcParameterInner", value)
}

// NewVpnIpsecPoliciesListPfsGroupIcParameterInnerFromValue returns a pointer to a valid VpnIpsecPoliciesListPfsGroupIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVpnIpsecPoliciesListPfsGroupIcParameterInnerFromValue(v int32) (*VpnIpsecPoliciesListPfsGroupIcParameterInner, error) {
	ev := VpnIpsecPoliciesListPfsGroupIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VpnIpsecPoliciesListPfsGroupIcParameterInner: valid values are %v", v, AllowedVpnIpsecPoliciesListPfsGroupIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VpnIpsecPoliciesListPfsGroupIcParameterInner) IsValid() bool {
	for _, existing := range AllowedVpnIpsecPoliciesListPfsGroupIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vpn_ipsec_policies_list_pfs_group__ic_parameter_inner value
func (v VpnIpsecPoliciesListPfsGroupIcParameterInner) Ptr() *VpnIpsecPoliciesListPfsGroupIcParameterInner {
	return &v
}

type NullableVpnIpsecPoliciesListPfsGroupIcParameterInner struct {
	value *VpnIpsecPoliciesListPfsGroupIcParameterInner
	isSet bool
}

func (v NullableVpnIpsecPoliciesListPfsGroupIcParameterInner) Get() *VpnIpsecPoliciesListPfsGroupIcParameterInner {
	return v.value
}

func (v *NullableVpnIpsecPoliciesListPfsGroupIcParameterInner) Set(val *VpnIpsecPoliciesListPfsGroupIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnIpsecPoliciesListPfsGroupIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnIpsecPoliciesListPfsGroupIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnIpsecPoliciesListPfsGroupIcParameterInner(val *VpnIpsecPoliciesListPfsGroupIcParameterInner) *NullableVpnIpsecPoliciesListPfsGroupIcParameterInner {
	return &NullableVpnIpsecPoliciesListPfsGroupIcParameterInner{value: val, isSet: true}
}

func (v NullableVpnIpsecPoliciesListPfsGroupIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnIpsecPoliciesListPfsGroupIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

