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

// EventRuleEventTypesInner * `object_created` - Object created * `object_updated` - Object updated * `object_deleted` - Object deleted * `job_started` - Job started * `job_completed` - Job completed * `job_failed` - Job failed * `job_errored` - Job errored
type EventRuleEventTypesInner string

// List of EventRule_event_types_inner
const (
	EVENTRULEEVENTTYPESINNER_OBJECT_CREATED EventRuleEventTypesInner = "object_created"
	EVENTRULEEVENTTYPESINNER_OBJECT_UPDATED EventRuleEventTypesInner = "object_updated"
	EVENTRULEEVENTTYPESINNER_OBJECT_DELETED EventRuleEventTypesInner = "object_deleted"
	EVENTRULEEVENTTYPESINNER_JOB_STARTED EventRuleEventTypesInner = "job_started"
	EVENTRULEEVENTTYPESINNER_JOB_COMPLETED EventRuleEventTypesInner = "job_completed"
	EVENTRULEEVENTTYPESINNER_JOB_FAILED EventRuleEventTypesInner = "job_failed"
	EVENTRULEEVENTTYPESINNER_JOB_ERRORED EventRuleEventTypesInner = "job_errored"
)

// All allowed values of EventRuleEventTypesInner enum
var AllowedEventRuleEventTypesInnerEnumValues = []EventRuleEventTypesInner{
	"object_created",
	"object_updated",
	"object_deleted",
	"job_started",
	"job_completed",
	"job_failed",
	"job_errored",
}

func (v *EventRuleEventTypesInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventRuleEventTypesInner(value)
	for _, existing := range AllowedEventRuleEventTypesInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventRuleEventTypesInner", value)
}

// NewEventRuleEventTypesInnerFromValue returns a pointer to a valid EventRuleEventTypesInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventRuleEventTypesInnerFromValue(v string) (*EventRuleEventTypesInner, error) {
	ev := EventRuleEventTypesInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventRuleEventTypesInner: valid values are %v", v, AllowedEventRuleEventTypesInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventRuleEventTypesInner) IsValid() bool {
	for _, existing := range AllowedEventRuleEventTypesInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventRule_event_types_inner value
func (v EventRuleEventTypesInner) Ptr() *EventRuleEventTypesInner {
	return &v
}

type NullableEventRuleEventTypesInner struct {
	value *EventRuleEventTypesInner
	isSet bool
}

func (v NullableEventRuleEventTypesInner) Get() *EventRuleEventTypesInner {
	return v.value
}

func (v *NullableEventRuleEventTypesInner) Set(val *EventRuleEventTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRuleEventTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRuleEventTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRuleEventTypesInner(val *EventRuleEventTypesInner) *NullableEventRuleEventTypesInner {
	return &NullableEventRuleEventTypesInner{value: val, isSet: true}
}

func (v NullableEventRuleEventTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRuleEventTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

