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

// checks if the InterfaceDuplex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceDuplex{}

// InterfaceDuplex struct for InterfaceDuplex
type InterfaceDuplex struct {
	Value *InterfaceDuplexValue `json:"value,omitempty"`
	Label *InterfaceDuplexLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceDuplex InterfaceDuplex

// NewInterfaceDuplex instantiates a new InterfaceDuplex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceDuplex() *InterfaceDuplex {
	this := InterfaceDuplex{}
	return &this
}

// NewInterfaceDuplexWithDefaults instantiates a new InterfaceDuplex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceDuplexWithDefaults() *InterfaceDuplex {
	this := InterfaceDuplex{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceDuplex) GetValue() InterfaceDuplexValue {
	if o == nil || IsNil(o.Value) {
		var ret InterfaceDuplexValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDuplex) GetValueOk() (*InterfaceDuplexValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceDuplex) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given InterfaceDuplexValue and assigns it to the Value field.
func (o *InterfaceDuplex) SetValue(v InterfaceDuplexValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceDuplex) GetLabel() InterfaceDuplexLabel {
	if o == nil || IsNil(o.Label) {
		var ret InterfaceDuplexLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceDuplex) GetLabelOk() (*InterfaceDuplexLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceDuplex) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given InterfaceDuplexLabel and assigns it to the Label field.
func (o *InterfaceDuplex) SetLabel(v InterfaceDuplexLabel) {
	o.Label = &v
}

func (o InterfaceDuplex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceDuplex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceDuplex) UnmarshalJSON(data []byte) (err error) {
	varInterfaceDuplex := _InterfaceDuplex{}

	err = json.Unmarshal(data, &varInterfaceDuplex)

	if err != nil {
		return err
	}

	*o = InterfaceDuplex(varInterfaceDuplex)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceDuplex struct {
	value *InterfaceDuplex
	isSet bool
}

func (v NullableInterfaceDuplex) Get() *InterfaceDuplex {
	return v.value
}

func (v *NullableInterfaceDuplex) Set(val *InterfaceDuplex) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceDuplex) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceDuplex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceDuplex(val *InterfaceDuplex) *NullableInterfaceDuplex {
	return &NullableInterfaceDuplex{value: val, isSet: true}
}

func (v NullableInterfaceDuplex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceDuplex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


