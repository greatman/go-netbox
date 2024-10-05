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

// checks if the EventRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventRule{}

// EventRule Adds support for custom fields and tags.
type EventRule struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	ObjectTypes []string `json:"object_types"`
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	// The types of event which will trigger this rule.
	EventTypes []EventRuleEventTypesInner `json:"event_types"`
	// A set of conditions which determine whether the event will be generated.
	Conditions interface{} `json:"conditions,omitempty"`
	ActionType EventRuleActionType `json:"action_type"`
	ActionObjectType string `json:"action_object_type"`
	ActionObjectId NullableInt64 `json:"action_object_id,omitempty"`
	ActionObject map[string]interface{} `json:"action_object"`
	Description *string `json:"description,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _EventRule EventRule

// NewEventRule instantiates a new EventRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventRule(id int32, url string, displayUrl string, display string, objectTypes []string, name string, eventTypes []EventRuleEventTypesInner, actionType EventRuleActionType, actionObjectType string, actionObject map[string]interface{}, created NullableTime, lastUpdated NullableTime) *EventRule {
	this := EventRule{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.ObjectTypes = objectTypes
	this.Name = name
	this.EventTypes = eventTypes
	this.ActionType = actionType
	this.ActionObjectType = actionObjectType
	this.ActionObject = actionObject
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewEventRuleWithDefaults instantiates a new EventRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventRuleWithDefaults() *EventRule {
	this := EventRule{}
	return &this
}

// GetId returns the Id field value
func (o *EventRule) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventRule) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *EventRule) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *EventRule) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *EventRule) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *EventRule) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *EventRule) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *EventRule) SetDisplay(v string) {
	o.Display = v
}

// GetObjectTypes returns the ObjectTypes field value
func (o *EventRule) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *EventRule) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value
func (o *EventRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventRule) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EventRule) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRule) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EventRule) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EventRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypes returns the EventTypes field value
func (o *EventRule) GetEventTypes() []EventRuleEventTypesInner {
	if o == nil {
		var ret []EventRuleEventTypesInner
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetEventTypesOk() ([]EventRuleEventTypesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *EventRule) SetEventTypes(v []EventRuleEventTypesInner) {
	o.EventTypes = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventRule) GetConditions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventRule) GetConditionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *EventRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given interface{} and assigns it to the Conditions field.
func (o *EventRule) SetConditions(v interface{}) {
	o.Conditions = v
}

// GetActionType returns the ActionType field value
func (o *EventRule) GetActionType() EventRuleActionType {
	if o == nil {
		var ret EventRuleActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetActionTypeOk() (*EventRuleActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *EventRule) SetActionType(v EventRuleActionType) {
	o.ActionType = v
}

// GetActionObjectType returns the ActionObjectType field value
func (o *EventRule) GetActionObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionObjectType
}

// GetActionObjectTypeOk returns a tuple with the ActionObjectType field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetActionObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionObjectType, true
}

// SetActionObjectType sets field value
func (o *EventRule) SetActionObjectType(v string) {
	o.ActionObjectType = v
}

// GetActionObjectId returns the ActionObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventRule) GetActionObjectId() int64 {
	if o == nil || IsNil(o.ActionObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ActionObjectId.Get()
}

// GetActionObjectIdOk returns a tuple with the ActionObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventRule) GetActionObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionObjectId.Get(), o.ActionObjectId.IsSet()
}

// HasActionObjectId returns a boolean if a field has been set.
func (o *EventRule) HasActionObjectId() bool {
	if o != nil && o.ActionObjectId.IsSet() {
		return true
	}

	return false
}

// SetActionObjectId gets a reference to the given NullableInt64 and assigns it to the ActionObjectId field.
func (o *EventRule) SetActionObjectId(v int64) {
	o.ActionObjectId.Set(&v)
}
// SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil
func (o *EventRule) SetActionObjectIdNil() {
	o.ActionObjectId.Set(nil)
}

// UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
func (o *EventRule) UnsetActionObjectId() {
	o.ActionObjectId.Unset()
}

// GetActionObject returns the ActionObject field value
func (o *EventRule) GetActionObject() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ActionObject
}

// GetActionObjectOk returns a tuple with the ActionObject field value
// and a boolean to check if the value has been set.
func (o *EventRule) GetActionObjectOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ActionObject, true
}

// SetActionObject sets field value
func (o *EventRule) SetActionObject(v map[string]interface{}) {
	o.ActionObject = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventRule) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *EventRule) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRule) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *EventRule) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *EventRule) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventRule) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRule) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventRule) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *EventRule) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EventRule) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventRule) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *EventRule) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EventRule) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventRule) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *EventRule) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o EventRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["object_types"] = o.ObjectTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["event_types"] = o.EventTypes
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	toSerialize["action_type"] = o.ActionType
	toSerialize["action_object_type"] = o.ActionObjectType
	if o.ActionObjectId.IsSet() {
		toSerialize["action_object_id"] = o.ActionObjectId.Get()
	}
	toSerialize["action_object"] = o.ActionObject
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"object_types",
		"name",
		"event_types",
		"action_type",
		"action_object_type",
		"action_object",
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

	varEventRule := _EventRule{}

	err = json.Unmarshal(data, &varEventRule)

	if err != nil {
		return err
	}

	*o = EventRule(varEventRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "event_types")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "action_type")
		delete(additionalProperties, "action_object_type")
		delete(additionalProperties, "action_object_id")
		delete(additionalProperties, "action_object")
		delete(additionalProperties, "description")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventRule struct {
	value *EventRule
	isSet bool
}

func (v NullableEventRule) Get() *EventRule {
	return v.value
}

func (v *NullableEventRule) Set(val *EventRule) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRule) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRule(val *EventRule) *NullableEventRule {
	return &NullableEventRule{value: val, isSet: true}
}

func (v NullableEventRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


