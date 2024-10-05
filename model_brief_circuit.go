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

// checks if the BriefCircuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCircuit{}

// BriefCircuit Adds support for custom fields and tags.
type BriefCircuit struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	// Unique circuit ID
	Cid string `json:"cid"`
	Provider BriefProvider `json:"provider"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefCircuit BriefCircuit

// NewBriefCircuit instantiates a new BriefCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCircuit(id int32, url string, display string, cid string, provider BriefProvider) *BriefCircuit {
	this := BriefCircuit{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Cid = cid
	this.Provider = provider
	return &this
}

// NewBriefCircuitWithDefaults instantiates a new BriefCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCircuitWithDefaults() *BriefCircuit {
	this := BriefCircuit{}
	return &this
}

// GetId returns the Id field value
func (o *BriefCircuit) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefCircuit) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefCircuit) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefCircuit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefCircuit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefCircuit) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefCircuit) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefCircuit) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefCircuit) SetDisplay(v string) {
	o.Display = v
}

// GetCid returns the Cid field value
func (o *BriefCircuit) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *BriefCircuit) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *BriefCircuit) SetCid(v string) {
	o.Cid = v
}

// GetProvider returns the Provider field value
func (o *BriefCircuit) GetProvider() BriefProvider {
	if o == nil {
		var ret BriefProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *BriefCircuit) GetProviderOk() (*BriefProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *BriefCircuit) SetProvider(v BriefProvider) {
	o.Provider = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefCircuit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefCircuit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefCircuit) SetDescription(v string) {
	o.Description = &v
}

func (o BriefCircuit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCircuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["cid"] = o.Cid
	toSerialize["provider"] = o.Provider
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefCircuit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"cid",
		"provider",
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

	varBriefCircuit := _BriefCircuit{}

	err = json.Unmarshal(data, &varBriefCircuit)

	if err != nil {
		return err
	}

	*o = BriefCircuit(varBriefCircuit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCircuit struct {
	value *BriefCircuit
	isSet bool
}

func (v NullableBriefCircuit) Get() *BriefCircuit {
	return v.value
}

func (v *NullableBriefCircuit) Set(val *BriefCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCircuit(val *BriefCircuit) *NullableBriefCircuit {
	return &NullableBriefCircuit{value: val, isSet: true}
}

func (v NullableBriefCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


