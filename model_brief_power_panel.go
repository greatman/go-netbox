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

// checks if the BriefPowerPanel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefPowerPanel{}

// BriefPowerPanel Adds support for custom fields and tags.
type BriefPowerPanel struct {
	Id int64 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	PowerfeedCount int64 `json:"powerfeed_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefPowerPanel BriefPowerPanel

// NewBriefPowerPanel instantiates a new BriefPowerPanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefPowerPanel(id int64, url string, display string, name string, powerfeedCount int64) *BriefPowerPanel {
	this := BriefPowerPanel{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.PowerfeedCount = powerfeedCount
	return &this
}

// NewBriefPowerPanelWithDefaults instantiates a new BriefPowerPanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefPowerPanelWithDefaults() *BriefPowerPanel {
	this := BriefPowerPanel{}
	return &this
}

// GetId returns the Id field value
func (o *BriefPowerPanel) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPanel) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefPowerPanel) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefPowerPanel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPanel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefPowerPanel) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefPowerPanel) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPanel) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefPowerPanel) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefPowerPanel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPanel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefPowerPanel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefPowerPanel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefPowerPanel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefPowerPanel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefPowerPanel) SetDescription(v string) {
	o.Description = &v
}

// GetPowerfeedCount returns the PowerfeedCount field value
func (o *BriefPowerPanel) GetPowerfeedCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PowerfeedCount
}

// GetPowerfeedCountOk returns a tuple with the PowerfeedCount field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPanel) GetPowerfeedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerfeedCount, true
}

// SetPowerfeedCount sets field value
func (o *BriefPowerPanel) SetPowerfeedCount(v int64) {
	o.PowerfeedCount = v
}

func (o BriefPowerPanel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefPowerPanel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["powerfeed_count"] = o.PowerfeedCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefPowerPanel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"powerfeed_count",
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

	varBriefPowerPanel := _BriefPowerPanel{}

	err = json.Unmarshal(data, &varBriefPowerPanel)

	if err != nil {
		return err
	}

	*o = BriefPowerPanel(varBriefPowerPanel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "powerfeed_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefPowerPanel struct {
	value *BriefPowerPanel
	isSet bool
}

func (v NullableBriefPowerPanel) Get() *BriefPowerPanel {
	return v.value
}

func (v *NullableBriefPowerPanel) Set(val *BriefPowerPanel) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefPowerPanel) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefPowerPanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefPowerPanel(val *BriefPowerPanel) *NullableBriefPowerPanel {
	return &NullableBriefPowerPanel{value: val, isSet: true}
}

func (v NullableBriefPowerPanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefPowerPanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


