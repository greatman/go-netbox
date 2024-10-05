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

// checks if the PatchedWritableVirtualChassisRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableVirtualChassisRequest{}

// PatchedWritableVirtualChassisRequest Adds support for custom fields and tags.
type PatchedWritableVirtualChassisRequest struct {
	Name *string `json:"name,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Master NullableInt32 `json:"master,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableVirtualChassisRequest PatchedWritableVirtualChassisRequest

// NewPatchedWritableVirtualChassisRequest instantiates a new PatchedWritableVirtualChassisRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableVirtualChassisRequest() *PatchedWritableVirtualChassisRequest {
	this := PatchedWritableVirtualChassisRequest{}
	return &this
}

// NewPatchedWritableVirtualChassisRequestWithDefaults instantiates a new PatchedWritableVirtualChassisRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableVirtualChassisRequestWithDefaults() *PatchedWritableVirtualChassisRequest {
	this := PatchedWritableVirtualChassisRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableVirtualChassisRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualChassisRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableVirtualChassisRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableVirtualChassisRequest) SetName(v string) {
	o.Name = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PatchedWritableVirtualChassisRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualChassisRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PatchedWritableVirtualChassisRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *PatchedWritableVirtualChassisRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetMaster returns the Master field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVirtualChassisRequest) GetMaster() int32 {
	if o == nil || IsNil(o.Master.Get()) {
		var ret int32
		return ret
	}
	return *o.Master.Get()
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVirtualChassisRequest) GetMasterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Master.Get(), o.Master.IsSet()
}

// HasMaster returns a boolean if a field has been set.
func (o *PatchedWritableVirtualChassisRequest) HasMaster() bool {
	if o != nil && o.Master.IsSet() {
		return true
	}

	return false
}

// SetMaster gets a reference to the given NullableInt32 and assigns it to the Master field.
func (o *PatchedWritableVirtualChassisRequest) SetMaster(v int32) {
	o.Master.Set(&v)
}
// SetMasterNil sets the value for Master to be an explicit nil
func (o *PatchedWritableVirtualChassisRequest) SetMasterNil() {
	o.Master.Set(nil)
}

// UnsetMaster ensures that no value is present for Master, not even an explicit nil
func (o *PatchedWritableVirtualChassisRequest) UnsetMaster() {
	o.Master.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableVirtualChassisRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualChassisRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableVirtualChassisRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableVirtualChassisRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableVirtualChassisRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualChassisRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableVirtualChassisRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableVirtualChassisRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableVirtualChassisRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualChassisRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableVirtualChassisRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableVirtualChassisRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableVirtualChassisRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualChassisRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableVirtualChassisRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *PatchedWritableVirtualChassisRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o PatchedWritableVirtualChassisRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableVirtualChassisRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if o.Master.IsSet() {
		toSerialize["master"] = o.Master.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *PatchedWritableVirtualChassisRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableVirtualChassisRequest := _PatchedWritableVirtualChassisRequest{}

	err = json.Unmarshal(data, &varPatchedWritableVirtualChassisRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableVirtualChassisRequest(varPatchedWritableVirtualChassisRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "master")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableVirtualChassisRequest struct {
	value *PatchedWritableVirtualChassisRequest
	isSet bool
}

func (v NullablePatchedWritableVirtualChassisRequest) Get() *PatchedWritableVirtualChassisRequest {
	return v.value
}

func (v *NullablePatchedWritableVirtualChassisRequest) Set(val *PatchedWritableVirtualChassisRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableVirtualChassisRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableVirtualChassisRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableVirtualChassisRequest(val *PatchedWritableVirtualChassisRequest) *NullablePatchedWritableVirtualChassisRequest {
	return &NullablePatchedWritableVirtualChassisRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableVirtualChassisRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableVirtualChassisRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


