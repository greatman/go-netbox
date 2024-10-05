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

// checks if the WritableVMInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableVMInterfaceRequest{}

// WritableVMInterfaceRequest Adds support for custom fields and tags.
type WritableVMInterfaceRequest struct {
	VirtualMachine BriefVirtualMachineRequest `json:"virtual_machine"`
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	Parent NullableInt32 `json:"parent,omitempty"`
	Bridge NullableInt32 `json:"bridge,omitempty"`
	Mtu NullableInt32 `json:"mtu,omitempty"`
	MacAddress NullableString `json:"mac_address,omitempty"`
	Description *string `json:"description,omitempty"`
	Mode *PatchedWritableInterfaceRequestMode `json:"mode,omitempty"`
	UntaggedVlan NullableBriefVLANRequest `json:"untagged_vlan,omitempty"`
	TaggedVlans []int32 `json:"tagged_vlans,omitempty"`
	Vrf NullableBriefVRFRequest `json:"vrf,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableVMInterfaceRequest WritableVMInterfaceRequest

// NewWritableVMInterfaceRequest instantiates a new WritableVMInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVMInterfaceRequest(virtualMachine BriefVirtualMachineRequest, name string) *WritableVMInterfaceRequest {
	this := WritableVMInterfaceRequest{}
	this.VirtualMachine = virtualMachine
	this.Name = name
	return &this
}

// NewWritableVMInterfaceRequestWithDefaults instantiates a new WritableVMInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVMInterfaceRequestWithDefaults() *WritableVMInterfaceRequest {
	this := WritableVMInterfaceRequest{}
	return &this
}

// GetVirtualMachine returns the VirtualMachine field value
func (o *WritableVMInterfaceRequest) GetVirtualMachine() BriefVirtualMachineRequest {
	if o == nil {
		var ret BriefVirtualMachineRequest
		return ret
	}

	return o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualMachine, true
}

// SetVirtualMachine sets field value
func (o *WritableVMInterfaceRequest) SetVirtualMachine(v BriefVirtualMachineRequest) {
	o.VirtualMachine = v
}

// GetName returns the Name field value
func (o *WritableVMInterfaceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVMInterfaceRequest) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WritableVMInterfaceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WritableVMInterfaceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterfaceRequest) GetParent() int32 {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret int32
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterfaceRequest) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInt32 and assigns it to the Parent field.
func (o *WritableVMInterfaceRequest) SetParent(v int32) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *WritableVMInterfaceRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *WritableVMInterfaceRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterfaceRequest) GetBridge() int32 {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret int32
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterfaceRequest) GetBridgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableInt32 and assigns it to the Bridge field.
func (o *WritableVMInterfaceRequest) SetBridge(v int32) {
	o.Bridge.Set(&v)
}
// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *WritableVMInterfaceRequest) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *WritableVMInterfaceRequest) UnsetBridge() {
	o.Bridge.Unset()
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterfaceRequest) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu.Get()) {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterfaceRequest) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *WritableVMInterfaceRequest) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *WritableVMInterfaceRequest) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *WritableVMInterfaceRequest) UnsetMtu() {
	o.Mtu.Unset()
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterfaceRequest) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress.Get()) {
		var ret string
		return ret
	}
	return *o.MacAddress.Get()
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterfaceRequest) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAddress.Get(), o.MacAddress.IsSet()
}

// HasMacAddress returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasMacAddress() bool {
	if o != nil && o.MacAddress.IsSet() {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given NullableString and assigns it to the MacAddress field.
func (o *WritableVMInterfaceRequest) SetMacAddress(v string) {
	o.MacAddress.Set(&v)
}
// SetMacAddressNil sets the value for MacAddress to be an explicit nil
func (o *WritableVMInterfaceRequest) SetMacAddressNil() {
	o.MacAddress.Set(nil)
}

// UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
func (o *WritableVMInterfaceRequest) UnsetMacAddress() {
	o.MacAddress.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVMInterfaceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVMInterfaceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *WritableVMInterfaceRequest) GetMode() PatchedWritableInterfaceRequestMode {
	if o == nil || IsNil(o.Mode) {
		var ret PatchedWritableInterfaceRequestMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetModeOk() (*PatchedWritableInterfaceRequestMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given PatchedWritableInterfaceRequestMode and assigns it to the Mode field.
func (o *WritableVMInterfaceRequest) SetMode(v PatchedWritableInterfaceRequestMode) {
	o.Mode = &v
}

// GetUntaggedVlan returns the UntaggedVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterfaceRequest) GetUntaggedVlan() BriefVLANRequest {
	if o == nil || IsNil(o.UntaggedVlan.Get()) {
		var ret BriefVLANRequest
		return ret
	}
	return *o.UntaggedVlan.Get()
}

// GetUntaggedVlanOk returns a tuple with the UntaggedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterfaceRequest) GetUntaggedVlanOk() (*BriefVLANRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.UntaggedVlan.Get(), o.UntaggedVlan.IsSet()
}

// HasUntaggedVlan returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasUntaggedVlan() bool {
	if o != nil && o.UntaggedVlan.IsSet() {
		return true
	}

	return false
}

// SetUntaggedVlan gets a reference to the given NullableBriefVLANRequest and assigns it to the UntaggedVlan field.
func (o *WritableVMInterfaceRequest) SetUntaggedVlan(v BriefVLANRequest) {
	o.UntaggedVlan.Set(&v)
}
// SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil
func (o *WritableVMInterfaceRequest) SetUntaggedVlanNil() {
	o.UntaggedVlan.Set(nil)
}

// UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
func (o *WritableVMInterfaceRequest) UnsetUntaggedVlan() {
	o.UntaggedVlan.Unset()
}

// GetTaggedVlans returns the TaggedVlans field value if set, zero value otherwise.
func (o *WritableVMInterfaceRequest) GetTaggedVlans() []int32 {
	if o == nil || IsNil(o.TaggedVlans) {
		var ret []int32
		return ret
	}
	return o.TaggedVlans
}

// GetTaggedVlansOk returns a tuple with the TaggedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetTaggedVlansOk() ([]int32, bool) {
	if o == nil || IsNil(o.TaggedVlans) {
		return nil, false
	}
	return o.TaggedVlans, true
}

// HasTaggedVlans returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasTaggedVlans() bool {
	if o != nil && !IsNil(o.TaggedVlans) {
		return true
	}

	return false
}

// SetTaggedVlans gets a reference to the given []int32 and assigns it to the TaggedVlans field.
func (o *WritableVMInterfaceRequest) SetTaggedVlans(v []int32) {
	o.TaggedVlans = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterfaceRequest) GetVrf() BriefVRFRequest {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BriefVRFRequest
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterfaceRequest) GetVrfOk() (*BriefVRFRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBriefVRFRequest and assigns it to the Vrf field.
func (o *WritableVMInterfaceRequest) SetVrf(v BriefVRFRequest) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *WritableVMInterfaceRequest) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *WritableVMInterfaceRequest) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVMInterfaceRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableVMInterfaceRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVMInterfaceRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterfaceRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVMInterfaceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *WritableVMInterfaceRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o WritableVMInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableVMInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["virtual_machine"] = o.VirtualMachine
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if o.MacAddress.IsSet() {
		toSerialize["mac_address"] = o.MacAddress.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if o.UntaggedVlan.IsSet() {
		toSerialize["untagged_vlan"] = o.UntaggedVlan.Get()
	}
	if !IsNil(o.TaggedVlans) {
		toSerialize["tagged_vlans"] = o.TaggedVlans
	}
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
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

func (o *WritableVMInterfaceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"virtual_machine",
		"name",
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

	varWritableVMInterfaceRequest := _WritableVMInterfaceRequest{}

	err = json.Unmarshal(data, &varWritableVMInterfaceRequest)

	if err != nil {
		return err
	}

	*o = WritableVMInterfaceRequest(varWritableVMInterfaceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "mac_address")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "untagged_vlan")
		delete(additionalProperties, "tagged_vlans")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableVMInterfaceRequest struct {
	value *WritableVMInterfaceRequest
	isSet bool
}

func (v NullableWritableVMInterfaceRequest) Get() *WritableVMInterfaceRequest {
	return v.value
}

func (v *NullableWritableVMInterfaceRequest) Set(val *WritableVMInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVMInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVMInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVMInterfaceRequest(val *WritableVMInterfaceRequest) *NullableWritableVMInterfaceRequest {
	return &NullableWritableVMInterfaceRequest{value: val, isSet: true}
}

func (v NullableWritableVMInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVMInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


