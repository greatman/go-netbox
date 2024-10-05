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

// checks if the NestedWirelessLANGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedWirelessLANGroup{}

// NestedWirelessLANGroup Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedWirelessLANGroup struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Depth int32 `json:"_depth"`
	AdditionalProperties map[string]interface{}
}

type _NestedWirelessLANGroup NestedWirelessLANGroup

// NewNestedWirelessLANGroup instantiates a new NestedWirelessLANGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedWirelessLANGroup(id int32, url string, displayUrl string, display string, name string, slug string, depth int32) *NestedWirelessLANGroup {
	this := NestedWirelessLANGroup{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Depth = depth
	return &this
}

// NewNestedWirelessLANGroupWithDefaults instantiates a new NestedWirelessLANGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedWirelessLANGroupWithDefaults() *NestedWirelessLANGroup {
	this := NestedWirelessLANGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NestedWirelessLANGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedWirelessLANGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedWirelessLANGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedWirelessLANGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *NestedWirelessLANGroup) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroup) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *NestedWirelessLANGroup) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *NestedWirelessLANGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedWirelessLANGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedWirelessLANGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedWirelessLANGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedWirelessLANGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedWirelessLANGroup) SetSlug(v string) {
	o.Slug = v
}

// GetDepth returns the Depth field value
func (o *NestedWirelessLANGroup) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroup) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *NestedWirelessLANGroup) SetDepth(v int32) {
	o.Depth = v
}

func (o NestedWirelessLANGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedWirelessLANGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["_depth"] = o.Depth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedWirelessLANGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
		"slug",
		"_depth",
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

	varNestedWirelessLANGroup := _NestedWirelessLANGroup{}

	err = json.Unmarshal(data, &varNestedWirelessLANGroup)

	if err != nil {
		return err
	}

	*o = NestedWirelessLANGroup(varNestedWirelessLANGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "_depth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedWirelessLANGroup struct {
	value *NestedWirelessLANGroup
	isSet bool
}

func (v NullableNestedWirelessLANGroup) Get() *NestedWirelessLANGroup {
	return v.value
}

func (v *NullableNestedWirelessLANGroup) Set(val *NestedWirelessLANGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedWirelessLANGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedWirelessLANGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedWirelessLANGroup(val *NestedWirelessLANGroup) *NullableNestedWirelessLANGroup {
	return &NullableNestedWirelessLANGroup{value: val, isSet: true}
}

func (v NullableNestedWirelessLANGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedWirelessLANGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


