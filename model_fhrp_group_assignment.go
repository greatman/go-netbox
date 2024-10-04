/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the FHRPGroupAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FHRPGroupAssignment{}

// FHRPGroupAssignment Adds support for custom fields and tags.
type FHRPGroupAssignment struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Group BriefFHRPGroup `json:"group"`
	InterfaceType string `json:"interface_type"`
	InterfaceId int64 `json:"interface_id"`
	Interface interface{} `json:"interface"`
	Priority int32 `json:"priority"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _FHRPGroupAssignment FHRPGroupAssignment

// NewFHRPGroupAssignment instantiates a new FHRPGroupAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFHRPGroupAssignment(id int32, url string, display string, group BriefFHRPGroup, interfaceType string, interfaceId int64, interface_ interface{}, priority int32, created NullableTime, lastUpdated NullableTime) *FHRPGroupAssignment {
	this := FHRPGroupAssignment{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Group = group
	this.InterfaceType = interfaceType
	this.InterfaceId = interfaceId
	this.Interface = interface_
	this.Priority = priority
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewFHRPGroupAssignmentWithDefaults instantiates a new FHRPGroupAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFHRPGroupAssignmentWithDefaults() *FHRPGroupAssignment {
	this := FHRPGroupAssignment{}
	return &this
}

// GetId returns the Id field value
func (o *FHRPGroupAssignment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FHRPGroupAssignment) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *FHRPGroupAssignment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FHRPGroupAssignment) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *FHRPGroupAssignment) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignment) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *FHRPGroupAssignment) SetDisplay(v string) {
	o.Display = v
}

// GetGroup returns the Group field value
func (o *FHRPGroupAssignment) GetGroup() BriefFHRPGroup {
	if o == nil {
		var ret BriefFHRPGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignment) GetGroupOk() (*BriefFHRPGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *FHRPGroupAssignment) SetGroup(v BriefFHRPGroup) {
	o.Group = v
}

// GetInterfaceType returns the InterfaceType field value
func (o *FHRPGroupAssignment) GetInterfaceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignment) GetInterfaceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceType, true
}

// SetInterfaceType sets field value
func (o *FHRPGroupAssignment) SetInterfaceType(v string) {
	o.InterfaceType = v
}

// GetInterfaceId returns the InterfaceId field value
func (o *FHRPGroupAssignment) GetInterfaceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignment) GetInterfaceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceId, true
}

// SetInterfaceId sets field value
func (o *FHRPGroupAssignment) SetInterfaceId(v int64) {
	o.InterfaceId = v
}

// GetInterface returns the Interface field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FHRPGroupAssignment) GetInterface() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FHRPGroupAssignment) GetInterfaceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *FHRPGroupAssignment) SetInterface(v interface{}) {
	o.Interface = v
}

// GetPriority returns the Priority field value
func (o *FHRPGroupAssignment) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupAssignment) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *FHRPGroupAssignment) SetPriority(v int32) {
	o.Priority = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FHRPGroupAssignment) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FHRPGroupAssignment) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *FHRPGroupAssignment) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FHRPGroupAssignment) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FHRPGroupAssignment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *FHRPGroupAssignment) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o FHRPGroupAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FHRPGroupAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["group"] = o.Group
	toSerialize["interface_type"] = o.InterfaceType
	toSerialize["interface_id"] = o.InterfaceId
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	toSerialize["priority"] = o.Priority
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FHRPGroupAssignment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"group",
		"interface_type",
		"interface_id",
		"interface",
		"priority",
		"created",
		"last_updated",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFHRPGroupAssignment := _FHRPGroupAssignment{}

	err = json.Unmarshal(data, &varFHRPGroupAssignment)

	if err != nil {
		return err
	}

	*o = FHRPGroupAssignment(varFHRPGroupAssignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "group")
		delete(additionalProperties, "interface_type")
		delete(additionalProperties, "interface_id")
		delete(additionalProperties, "interface")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFHRPGroupAssignment struct {
	value *FHRPGroupAssignment
	isSet bool
}

func (v NullableFHRPGroupAssignment) Get() *FHRPGroupAssignment {
	return v.value
}

func (v *NullableFHRPGroupAssignment) Set(val *FHRPGroupAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableFHRPGroupAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableFHRPGroupAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFHRPGroupAssignment(val *FHRPGroupAssignment) *NullableFHRPGroupAssignment {
	return &NullableFHRPGroupAssignment{value: val, isSet: true}
}

func (v NullableFHRPGroupAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFHRPGroupAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


