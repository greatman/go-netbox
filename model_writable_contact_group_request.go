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

// checks if the WritableContactGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableContactGroupRequest{}

// WritableContactGroupRequest Extends PrimaryModelSerializer to include MPTT support.
type WritableContactGroupRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Parent NullableInt64 `json:"parent"`
	Description *string `json:"description,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableContactGroupRequest WritableContactGroupRequest

// NewWritableContactGroupRequest instantiates a new WritableContactGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableContactGroupRequest(name string, slug string, parent NullableInt64) *WritableContactGroupRequest {
	this := WritableContactGroupRequest{}
	this.Name = name
	this.Slug = slug
	this.Parent = parent
	return &this
}

// NewWritableContactGroupRequestWithDefaults instantiates a new WritableContactGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableContactGroupRequestWithDefaults() *WritableContactGroupRequest {
	this := WritableContactGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableContactGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableContactGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableContactGroupRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WritableContactGroupRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritableContactGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritableContactGroupRequest) SetSlug(v string) {
	o.Slug = v
}

// GetParent returns the Parent field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *WritableContactGroupRequest) GetParent() int64 {
	if o == nil || o.Parent.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableContactGroupRequest) GetParentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// SetParent sets field value
func (o *WritableContactGroupRequest) SetParent(v int64) {
	o.Parent.Set(&v)
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableContactGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableContactGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableContactGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableContactGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableContactGroupRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableContactGroupRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableContactGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableContactGroupRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableContactGroupRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableContactGroupRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableContactGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *WritableContactGroupRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o WritableContactGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableContactGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["parent"] = o.Parent.Get()
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableContactGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
		"parent",
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

	varWritableContactGroupRequest := _WritableContactGroupRequest{}

	err = json.Unmarshal(data, &varWritableContactGroupRequest)

	if err != nil {
		return err
	}

	*o = WritableContactGroupRequest(varWritableContactGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableContactGroupRequest struct {
	value *WritableContactGroupRequest
	isSet bool
}

func (v NullableWritableContactGroupRequest) Get() *WritableContactGroupRequest {
	return v.value
}

func (v *NullableWritableContactGroupRequest) Set(val *WritableContactGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableContactGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableContactGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableContactGroupRequest(val *WritableContactGroupRequest) *NullableWritableContactGroupRequest {
	return &NullableWritableContactGroupRequest{value: val, isSet: true}
}

func (v NullableWritableContactGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableContactGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


