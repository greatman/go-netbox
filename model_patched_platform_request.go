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

// checks if the PatchedPlatformRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedPlatformRequest{}

// PatchedPlatformRequest Adds support for custom fields and tags.
type PatchedPlatformRequest struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Manufacturer NullableBriefManufacturerRequest `json:"manufacturer,omitempty"`
	ConfigTemplate NullableBriefConfigTemplateRequest `json:"config_template,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedPlatformRequest PatchedPlatformRequest

// NewPatchedPlatformRequest instantiates a new PatchedPlatformRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPlatformRequest() *PatchedPlatformRequest {
	this := PatchedPlatformRequest{}
	return &this
}

// NewPatchedPlatformRequestWithDefaults instantiates a new PatchedPlatformRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPlatformRequestWithDefaults() *PatchedPlatformRequest {
	this := PatchedPlatformRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPlatformRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlatformRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPlatformRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPlatformRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedPlatformRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlatformRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedPlatformRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedPlatformRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPlatformRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret BriefManufacturerRequest
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPlatformRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *PatchedPlatformRequest) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableBriefManufacturerRequest and assigns it to the Manufacturer field.
func (o *PatchedPlatformRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *PatchedPlatformRequest) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *PatchedPlatformRequest) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPlatformRequest) GetConfigTemplate() BriefConfigTemplateRequest {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret BriefConfigTemplateRequest
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPlatformRequest) GetConfigTemplateOk() (*BriefConfigTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *PatchedPlatformRequest) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableBriefConfigTemplateRequest and assigns it to the ConfigTemplate field.
func (o *PatchedPlatformRequest) SetConfigTemplate(v BriefConfigTemplateRequest) {
	o.ConfigTemplate.Set(&v)
}
// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *PatchedPlatformRequest) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *PatchedPlatformRequest) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedPlatformRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlatformRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedPlatformRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedPlatformRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedPlatformRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlatformRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedPlatformRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedPlatformRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedPlatformRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlatformRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedPlatformRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedPlatformRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedPlatformRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedPlatformRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
	}
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

func (o *PatchedPlatformRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedPlatformRequest := _PatchedPlatformRequest{}

	err = json.Unmarshal(data, &varPatchedPlatformRequest)

	if err != nil {
		return err
	}

	*o = PatchedPlatformRequest(varPatchedPlatformRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedPlatformRequest struct {
	value *PatchedPlatformRequest
	isSet bool
}

func (v NullablePatchedPlatformRequest) Get() *PatchedPlatformRequest {
	return v.value
}

func (v *NullablePatchedPlatformRequest) Set(val *PatchedPlatformRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPlatformRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPlatformRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPlatformRequest(val *PatchedPlatformRequest) *NullablePatchedPlatformRequest {
	return &NullablePatchedPlatformRequest{value: val, isSet: true}
}

func (v NullablePatchedPlatformRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPlatformRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


