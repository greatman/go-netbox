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

// DcimInterfaceTemplatesListPoeModeIcParameterInner the model 'DcimInterfaceTemplatesListPoeModeIcParameterInner'
type DcimInterfaceTemplatesListPoeModeIcParameterInner string

// List of dcim_interface_templates_list_poe_mode__ic_parameter_inner
const (
	DCIMINTERFACETEMPLATESLISTPOEMODEICPARAMETERINNER_EMPTY DcimInterfaceTemplatesListPoeModeIcParameterInner = ""
	DCIMINTERFACETEMPLATESLISTPOEMODEICPARAMETERINNER_PD DcimInterfaceTemplatesListPoeModeIcParameterInner = "pd"
	DCIMINTERFACETEMPLATESLISTPOEMODEICPARAMETERINNER_PSE DcimInterfaceTemplatesListPoeModeIcParameterInner = "pse"
)

// All allowed values of DcimInterfaceTemplatesListPoeModeIcParameterInner enum
var AllowedDcimInterfaceTemplatesListPoeModeIcParameterInnerEnumValues = []DcimInterfaceTemplatesListPoeModeIcParameterInner{
	"",
	"pd",
	"pse",
}

func (v *DcimInterfaceTemplatesListPoeModeIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimInterfaceTemplatesListPoeModeIcParameterInner(value)
	for _, existing := range AllowedDcimInterfaceTemplatesListPoeModeIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimInterfaceTemplatesListPoeModeIcParameterInner", value)
}

// NewDcimInterfaceTemplatesListPoeModeIcParameterInnerFromValue returns a pointer to a valid DcimInterfaceTemplatesListPoeModeIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimInterfaceTemplatesListPoeModeIcParameterInnerFromValue(v string) (*DcimInterfaceTemplatesListPoeModeIcParameterInner, error) {
	ev := DcimInterfaceTemplatesListPoeModeIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimInterfaceTemplatesListPoeModeIcParameterInner: valid values are %v", v, AllowedDcimInterfaceTemplatesListPoeModeIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimInterfaceTemplatesListPoeModeIcParameterInner) IsValid() bool {
	for _, existing := range AllowedDcimInterfaceTemplatesListPoeModeIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_interface_templates_list_poe_mode__ic_parameter_inner value
func (v DcimInterfaceTemplatesListPoeModeIcParameterInner) Ptr() *DcimInterfaceTemplatesListPoeModeIcParameterInner {
	return &v
}

type NullableDcimInterfaceTemplatesListPoeModeIcParameterInner struct {
	value *DcimInterfaceTemplatesListPoeModeIcParameterInner
	isSet bool
}

func (v NullableDcimInterfaceTemplatesListPoeModeIcParameterInner) Get() *DcimInterfaceTemplatesListPoeModeIcParameterInner {
	return v.value
}

func (v *NullableDcimInterfaceTemplatesListPoeModeIcParameterInner) Set(val *DcimInterfaceTemplatesListPoeModeIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimInterfaceTemplatesListPoeModeIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimInterfaceTemplatesListPoeModeIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimInterfaceTemplatesListPoeModeIcParameterInner(val *DcimInterfaceTemplatesListPoeModeIcParameterInner) *NullableDcimInterfaceTemplatesListPoeModeIcParameterInner {
	return &NullableDcimInterfaceTemplatesListPoeModeIcParameterInner{value: val, isSet: true}
}

func (v NullableDcimInterfaceTemplatesListPoeModeIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimInterfaceTemplatesListPoeModeIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

