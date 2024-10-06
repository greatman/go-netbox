/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedFHRPGroupAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedFHRPGroupAssignmentRequest{}

// PatchedFHRPGroupAssignmentRequest Adds support for custom fields and tags.
type PatchedFHRPGroupAssignmentRequest struct {
	Group *BriefFHRPGroupRequest `json:"group,omitempty"`
	InterfaceType *string `json:"interface_type,omitempty"`
	InterfaceId *int64 `json:"interface_id,omitempty"`
	Priority *int64 `json:"priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedFHRPGroupAssignmentRequest PatchedFHRPGroupAssignmentRequest

// NewPatchedFHRPGroupAssignmentRequest instantiates a new PatchedFHRPGroupAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedFHRPGroupAssignmentRequest() *PatchedFHRPGroupAssignmentRequest {
	this := PatchedFHRPGroupAssignmentRequest{}
	return &this
}

// NewPatchedFHRPGroupAssignmentRequestWithDefaults instantiates a new PatchedFHRPGroupAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedFHRPGroupAssignmentRequestWithDefaults() *PatchedFHRPGroupAssignmentRequest {
	this := PatchedFHRPGroupAssignmentRequest{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedFHRPGroupAssignmentRequest) GetGroup() BriefFHRPGroupRequest {
	if o == nil || IsNil(o.Group) {
		var ret BriefFHRPGroupRequest
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupAssignmentRequest) GetGroupOk() (*BriefFHRPGroupRequest, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedFHRPGroupAssignmentRequest) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given BriefFHRPGroupRequest and assigns it to the Group field.
func (o *PatchedFHRPGroupAssignmentRequest) SetGroup(v BriefFHRPGroupRequest) {
	o.Group = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *PatchedFHRPGroupAssignmentRequest) GetInterfaceType() string {
	if o == nil || IsNil(o.InterfaceType) {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupAssignmentRequest) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceType) {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *PatchedFHRPGroupAssignmentRequest) HasInterfaceType() bool {
	if o != nil && !IsNil(o.InterfaceType) {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *PatchedFHRPGroupAssignmentRequest) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *PatchedFHRPGroupAssignmentRequest) GetInterfaceId() int64 {
	if o == nil || IsNil(o.InterfaceId) {
		var ret int64
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupAssignmentRequest) GetInterfaceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *PatchedFHRPGroupAssignmentRequest) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given int64 and assigns it to the InterfaceId field.
func (o *PatchedFHRPGroupAssignmentRequest) SetInterfaceId(v int64) {
	o.InterfaceId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PatchedFHRPGroupAssignmentRequest) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupAssignmentRequest) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PatchedFHRPGroupAssignmentRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *PatchedFHRPGroupAssignmentRequest) SetPriority(v int64) {
	o.Priority = &v
}

func (o PatchedFHRPGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedFHRPGroupAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.InterfaceType) {
		toSerialize["interface_type"] = o.InterfaceType
	}
	if !IsNil(o.InterfaceId) {
		toSerialize["interface_id"] = o.InterfaceId
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedFHRPGroupAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedFHRPGroupAssignmentRequest := _PatchedFHRPGroupAssignmentRequest{}

	err = json.Unmarshal(data, &varPatchedFHRPGroupAssignmentRequest)

	if err != nil {
		return err
	}

	*o = PatchedFHRPGroupAssignmentRequest(varPatchedFHRPGroupAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "interface_type")
		delete(additionalProperties, "interface_id")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedFHRPGroupAssignmentRequest struct {
	value *PatchedFHRPGroupAssignmentRequest
	isSet bool
}

func (v NullablePatchedFHRPGroupAssignmentRequest) Get() *PatchedFHRPGroupAssignmentRequest {
	return v.value
}

func (v *NullablePatchedFHRPGroupAssignmentRequest) Set(val *PatchedFHRPGroupAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedFHRPGroupAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedFHRPGroupAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedFHRPGroupAssignmentRequest(val *PatchedFHRPGroupAssignmentRequest) *NullablePatchedFHRPGroupAssignmentRequest {
	return &NullablePatchedFHRPGroupAssignmentRequest{value: val, isSet: true}
}

func (v NullablePatchedFHRPGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedFHRPGroupAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


