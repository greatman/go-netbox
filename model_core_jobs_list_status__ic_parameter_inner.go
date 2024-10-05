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

// CoreJobsListStatusIcParameterInner the model 'CoreJobsListStatusIcParameterInner'
type CoreJobsListStatusIcParameterInner string

// List of core_jobs_list_status__ic_parameter_inner
const (
	COREJOBSLISTSTATUSICPARAMETERINNER_COMPLETED CoreJobsListStatusIcParameterInner = "completed"
	COREJOBSLISTSTATUSICPARAMETERINNER_ERRORED CoreJobsListStatusIcParameterInner = "errored"
	COREJOBSLISTSTATUSICPARAMETERINNER_FAILED CoreJobsListStatusIcParameterInner = "failed"
	COREJOBSLISTSTATUSICPARAMETERINNER_PENDING CoreJobsListStatusIcParameterInner = "pending"
	COREJOBSLISTSTATUSICPARAMETERINNER_RUNNING CoreJobsListStatusIcParameterInner = "running"
	COREJOBSLISTSTATUSICPARAMETERINNER_SCHEDULED CoreJobsListStatusIcParameterInner = "scheduled"
)

// All allowed values of CoreJobsListStatusIcParameterInner enum
var AllowedCoreJobsListStatusIcParameterInnerEnumValues = []CoreJobsListStatusIcParameterInner{
	"completed",
	"errored",
	"failed",
	"pending",
	"running",
	"scheduled",
}

func (v *CoreJobsListStatusIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CoreJobsListStatusIcParameterInner(value)
	for _, existing := range AllowedCoreJobsListStatusIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CoreJobsListStatusIcParameterInner", value)
}

// NewCoreJobsListStatusIcParameterInnerFromValue returns a pointer to a valid CoreJobsListStatusIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCoreJobsListStatusIcParameterInnerFromValue(v string) (*CoreJobsListStatusIcParameterInner, error) {
	ev := CoreJobsListStatusIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CoreJobsListStatusIcParameterInner: valid values are %v", v, AllowedCoreJobsListStatusIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CoreJobsListStatusIcParameterInner) IsValid() bool {
	for _, existing := range AllowedCoreJobsListStatusIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to core_jobs_list_status__ic_parameter_inner value
func (v CoreJobsListStatusIcParameterInner) Ptr() *CoreJobsListStatusIcParameterInner {
	return &v
}

type NullableCoreJobsListStatusIcParameterInner struct {
	value *CoreJobsListStatusIcParameterInner
	isSet bool
}

func (v NullableCoreJobsListStatusIcParameterInner) Get() *CoreJobsListStatusIcParameterInner {
	return v.value
}

func (v *NullableCoreJobsListStatusIcParameterInner) Set(val *CoreJobsListStatusIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreJobsListStatusIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreJobsListStatusIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreJobsListStatusIcParameterInner(val *CoreJobsListStatusIcParameterInner) *NullableCoreJobsListStatusIcParameterInner {
	return &NullableCoreJobsListStatusIcParameterInner{value: val, isSet: true}
}

func (v NullableCoreJobsListStatusIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreJobsListStatusIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

