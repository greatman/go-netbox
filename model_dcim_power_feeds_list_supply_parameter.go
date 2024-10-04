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

// DcimPowerFeedsListSupplyParameter the model 'DcimPowerFeedsListSupplyParameter'
type DcimPowerFeedsListSupplyParameter string

// List of dcim_power_feeds_list_supply_parameter
const (
	DCIMPOWERFEEDSLISTSUPPLYPARAMETER_AC DcimPowerFeedsListSupplyParameter = "ac"
	DCIMPOWERFEEDSLISTSUPPLYPARAMETER_DC DcimPowerFeedsListSupplyParameter = "dc"
)

// All allowed values of DcimPowerFeedsListSupplyParameter enum
var AllowedDcimPowerFeedsListSupplyParameterEnumValues = []DcimPowerFeedsListSupplyParameter{
	"ac",
	"dc",
}

func (v *DcimPowerFeedsListSupplyParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimPowerFeedsListSupplyParameter(value)
	for _, existing := range AllowedDcimPowerFeedsListSupplyParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimPowerFeedsListSupplyParameter", value)
}

// NewDcimPowerFeedsListSupplyParameterFromValue returns a pointer to a valid DcimPowerFeedsListSupplyParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimPowerFeedsListSupplyParameterFromValue(v string) (*DcimPowerFeedsListSupplyParameter, error) {
	ev := DcimPowerFeedsListSupplyParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimPowerFeedsListSupplyParameter: valid values are %v", v, AllowedDcimPowerFeedsListSupplyParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimPowerFeedsListSupplyParameter) IsValid() bool {
	for _, existing := range AllowedDcimPowerFeedsListSupplyParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_power_feeds_list_supply_parameter value
func (v DcimPowerFeedsListSupplyParameter) Ptr() *DcimPowerFeedsListSupplyParameter {
	return &v
}

type NullableDcimPowerFeedsListSupplyParameter struct {
	value *DcimPowerFeedsListSupplyParameter
	isSet bool
}

func (v NullableDcimPowerFeedsListSupplyParameter) Get() *DcimPowerFeedsListSupplyParameter {
	return v.value
}

func (v *NullableDcimPowerFeedsListSupplyParameter) Set(val *DcimPowerFeedsListSupplyParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimPowerFeedsListSupplyParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimPowerFeedsListSupplyParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimPowerFeedsListSupplyParameter(val *DcimPowerFeedsListSupplyParameter) *NullableDcimPowerFeedsListSupplyParameter {
	return &NullableDcimPowerFeedsListSupplyParameter{value: val, isSet: true}
}

func (v NullableDcimPowerFeedsListSupplyParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimPowerFeedsListSupplyParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

