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

// checks if the ServiceTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTemplateRequest{}

// ServiceTemplateRequest Adds support for custom fields and tags.
type ServiceTemplateRequest struct {
	Name string `json:"name"`
	Protocol *PatchedWritableServiceRequestProtocol `json:"protocol,omitempty"`
	Ports []int32 `json:"ports"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceTemplateRequest ServiceTemplateRequest

// NewServiceTemplateRequest instantiates a new ServiceTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTemplateRequest(name string, ports []int32) *ServiceTemplateRequest {
	this := ServiceTemplateRequest{}
	this.Name = name
	this.Ports = ports
	return &this
}

// NewServiceTemplateRequestWithDefaults instantiates a new ServiceTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTemplateRequestWithDefaults() *ServiceTemplateRequest {
	this := ServiceTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ServiceTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ServiceTemplateRequest) GetProtocol() PatchedWritableServiceRequestProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret PatchedWritableServiceRequestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateRequest) GetProtocolOk() (*PatchedWritableServiceRequestProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ServiceTemplateRequest) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given PatchedWritableServiceRequestProtocol and assigns it to the Protocol field.
func (o *ServiceTemplateRequest) SetProtocol(v PatchedWritableServiceRequestProtocol) {
	o.Protocol = &v
}

// GetPorts returns the Ports field value
func (o *ServiceTemplateRequest) GetPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *ServiceTemplateRequest) GetPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *ServiceTemplateRequest) SetPorts(v []int32) {
	o.Ports = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ServiceTemplateRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ServiceTemplateRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ServiceTemplateRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceTemplateRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceTemplateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ServiceTemplateRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ServiceTemplateRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ServiceTemplateRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ServiceTemplateRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ServiceTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	toSerialize["ports"] = o.Ports
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

func (o *ServiceTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"ports",
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

	varServiceTemplateRequest := _ServiceTemplateRequest{}

	err = json.Unmarshal(data, &varServiceTemplateRequest)

	if err != nil {
		return err
	}

	*o = ServiceTemplateRequest(varServiceTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceTemplateRequest struct {
	value *ServiceTemplateRequest
	isSet bool
}

func (v NullableServiceTemplateRequest) Get() *ServiceTemplateRequest {
	return v.value
}

func (v *NullableServiceTemplateRequest) Set(val *ServiceTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTemplateRequest(val *ServiceTemplateRequest) *NullableServiceTemplateRequest {
	return &NullableServiceTemplateRequest{value: val, isSet: true}
}

func (v NullableServiceTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


