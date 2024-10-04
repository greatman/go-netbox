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

// checks if the PatchedDeviceBayTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedDeviceBayTemplateRequest{}

// PatchedDeviceBayTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedDeviceBayTemplateRequest struct {
	DeviceType *BriefDeviceTypeRequest `json:"device_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedDeviceBayTemplateRequest PatchedDeviceBayTemplateRequest

// NewPatchedDeviceBayTemplateRequest instantiates a new PatchedDeviceBayTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDeviceBayTemplateRequest() *PatchedDeviceBayTemplateRequest {
	this := PatchedDeviceBayTemplateRequest{}
	return &this
}

// NewPatchedDeviceBayTemplateRequestWithDefaults instantiates a new PatchedDeviceBayTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDeviceBayTemplateRequestWithDefaults() *PatchedDeviceBayTemplateRequest {
	this := PatchedDeviceBayTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PatchedDeviceBayTemplateRequest) GetDeviceType() BriefDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType) {
		var ret BriefDeviceTypeRequest
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedDeviceBayTemplateRequest) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given BriefDeviceTypeRequest and assigns it to the DeviceType field.
func (o *PatchedDeviceBayTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest) {
	o.DeviceType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDeviceBayTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDeviceBayTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDeviceBayTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedDeviceBayTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedDeviceBayTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedDeviceBayTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedDeviceBayTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedDeviceBayTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedDeviceBayTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedDeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedDeviceBayTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedDeviceBayTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedDeviceBayTemplateRequest := _PatchedDeviceBayTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedDeviceBayTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedDeviceBayTemplateRequest(varPatchedDeviceBayTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedDeviceBayTemplateRequest struct {
	value *PatchedDeviceBayTemplateRequest
	isSet bool
}

func (v NullablePatchedDeviceBayTemplateRequest) Get() *PatchedDeviceBayTemplateRequest {
	return v.value
}

func (v *NullablePatchedDeviceBayTemplateRequest) Set(val *PatchedDeviceBayTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDeviceBayTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDeviceBayTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDeviceBayTemplateRequest(val *PatchedDeviceBayTemplateRequest) *NullablePatchedDeviceBayTemplateRequest {
	return &NullablePatchedDeviceBayTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedDeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDeviceBayTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


