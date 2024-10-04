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

// checks if the NestedTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedTagRequest{}

// NestedTagRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedTagRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-\\\\w]+$"`
	Color *string `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	AdditionalProperties map[string]interface{}
}

type _NestedTagRequest NestedTagRequest

// NewNestedTagRequest instantiates a new NestedTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTagRequest(name string, slug string) *NestedTagRequest {
	this := NestedTagRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedTagRequestWithDefaults instantiates a new NestedTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTagRequestWithDefaults() *NestedTagRequest {
	this := NestedTagRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedTagRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedTagRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedTagRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedTagRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedTagRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedTagRequest) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *NestedTagRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTagRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *NestedTagRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *NestedTagRequest) SetColor(v string) {
	o.Color = &v
}

func (o NestedTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedTagRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varNestedTagRequest := _NestedTagRequest{}

	err = json.Unmarshal(data, &varNestedTagRequest)

	if err != nil {
		return err
	}

	*o = NestedTagRequest(varNestedTagRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedTagRequest struct {
	value *NestedTagRequest
	isSet bool
}

func (v NullableNestedTagRequest) Get() *NestedTagRequest {
	return v.value
}

func (v *NullableNestedTagRequest) Set(val *NestedTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTagRequest(val *NestedTagRequest) *NullableNestedTagRequest {
	return &NullableNestedTagRequest{value: val, isSet: true}
}

func (v NullableNestedTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


