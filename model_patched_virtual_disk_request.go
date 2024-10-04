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

// checks if the PatchedVirtualDiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedVirtualDiskRequest{}

// PatchedVirtualDiskRequest Adds support for custom fields and tags.
type PatchedVirtualDiskRequest struct {
	VirtualMachine *BriefVirtualMachineRequest `json:"virtual_machine,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedVirtualDiskRequest PatchedVirtualDiskRequest

// NewPatchedVirtualDiskRequest instantiates a new PatchedVirtualDiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedVirtualDiskRequest() *PatchedVirtualDiskRequest {
	this := PatchedVirtualDiskRequest{}
	return &this
}

// NewPatchedVirtualDiskRequestWithDefaults instantiates a new PatchedVirtualDiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedVirtualDiskRequestWithDefaults() *PatchedVirtualDiskRequest {
	this := PatchedVirtualDiskRequest{}
	return &this
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *PatchedVirtualDiskRequest) GetVirtualMachine() BriefVirtualMachineRequest {
	if o == nil || IsNil(o.VirtualMachine) {
		var ret BriefVirtualMachineRequest
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVirtualDiskRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool) {
	if o == nil || IsNil(o.VirtualMachine) {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *PatchedVirtualDiskRequest) HasVirtualMachine() bool {
	if o != nil && !IsNil(o.VirtualMachine) {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given BriefVirtualMachineRequest and assigns it to the VirtualMachine field.
func (o *PatchedVirtualDiskRequest) SetVirtualMachine(v BriefVirtualMachineRequest) {
	o.VirtualMachine = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedVirtualDiskRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVirtualDiskRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedVirtualDiskRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedVirtualDiskRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedVirtualDiskRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVirtualDiskRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedVirtualDiskRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedVirtualDiskRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PatchedVirtualDiskRequest) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVirtualDiskRequest) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PatchedVirtualDiskRequest) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PatchedVirtualDiskRequest) SetSize(v int32) {
	o.Size = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedVirtualDiskRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVirtualDiskRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedVirtualDiskRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedVirtualDiskRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedVirtualDiskRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVirtualDiskRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedVirtualDiskRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedVirtualDiskRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedVirtualDiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedVirtualDiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VirtualMachine) {
		toSerialize["virtual_machine"] = o.VirtualMachine
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
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

func (o *PatchedVirtualDiskRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedVirtualDiskRequest := _PatchedVirtualDiskRequest{}

	err = json.Unmarshal(data, &varPatchedVirtualDiskRequest)

	if err != nil {
		return err
	}

	*o = PatchedVirtualDiskRequest(varPatchedVirtualDiskRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "size")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedVirtualDiskRequest struct {
	value *PatchedVirtualDiskRequest
	isSet bool
}

func (v NullablePatchedVirtualDiskRequest) Get() *PatchedVirtualDiskRequest {
	return v.value
}

func (v *NullablePatchedVirtualDiskRequest) Set(val *PatchedVirtualDiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedVirtualDiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedVirtualDiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedVirtualDiskRequest(val *PatchedVirtualDiskRequest) *NullablePatchedVirtualDiskRequest {
	return &NullablePatchedVirtualDiskRequest{value: val, isSet: true}
}

func (v NullablePatchedVirtualDiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedVirtualDiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


