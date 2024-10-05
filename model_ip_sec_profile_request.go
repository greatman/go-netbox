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

// checks if the IPSecProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecProfileRequest{}

// IPSecProfileRequest Adds support for custom fields and tags.
type IPSecProfileRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Mode IPSecProfileModeValue `json:"mode"`
	IkePolicy BriefIKEPolicyRequest `json:"ike_policy"`
	IpsecPolicy BriefIPSecPolicyRequest `json:"ipsec_policy"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPSecProfileRequest IPSecProfileRequest

// NewIPSecProfileRequest instantiates a new IPSecProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecProfileRequest(name string, mode IPSecProfileModeValue, ikePolicy BriefIKEPolicyRequest, ipsecPolicy BriefIPSecPolicyRequest) *IPSecProfileRequest {
	this := IPSecProfileRequest{}
	this.Name = name
	this.Mode = mode
	this.IkePolicy = ikePolicy
	this.IpsecPolicy = ipsecPolicy
	return &this
}

// NewIPSecProfileRequestWithDefaults instantiates a new IPSecProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecProfileRequestWithDefaults() *IPSecProfileRequest {
	this := IPSecProfileRequest{}
	return &this
}

// GetName returns the Name field value
func (o *IPSecProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IPSecProfileRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPSecProfileRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPSecProfileRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPSecProfileRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value
func (o *IPSecProfileRequest) GetMode() IPSecProfileModeValue {
	if o == nil {
		var ret IPSecProfileModeValue
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetModeOk() (*IPSecProfileModeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *IPSecProfileRequest) SetMode(v IPSecProfileModeValue) {
	o.Mode = v
}

// GetIkePolicy returns the IkePolicy field value
func (o *IPSecProfileRequest) GetIkePolicy() BriefIKEPolicyRequest {
	if o == nil {
		var ret BriefIKEPolicyRequest
		return ret
	}

	return o.IkePolicy
}

// GetIkePolicyOk returns a tuple with the IkePolicy field value
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetIkePolicyOk() (*BriefIKEPolicyRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IkePolicy, true
}

// SetIkePolicy sets field value
func (o *IPSecProfileRequest) SetIkePolicy(v BriefIKEPolicyRequest) {
	o.IkePolicy = v
}

// GetIpsecPolicy returns the IpsecPolicy field value
func (o *IPSecProfileRequest) GetIpsecPolicy() BriefIPSecPolicyRequest {
	if o == nil {
		var ret BriefIPSecPolicyRequest
		return ret
	}

	return o.IpsecPolicy
}

// GetIpsecPolicyOk returns a tuple with the IpsecPolicy field value
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetIpsecPolicyOk() (*BriefIPSecPolicyRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpsecPolicy, true
}

// SetIpsecPolicy sets field value
func (o *IPSecProfileRequest) SetIpsecPolicy(v BriefIPSecPolicyRequest) {
	o.IpsecPolicy = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPSecProfileRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPSecProfileRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPSecProfileRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPSecProfileRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPSecProfileRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *IPSecProfileRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPSecProfileRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfileRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPSecProfileRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *IPSecProfileRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o IPSecProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["mode"] = o.Mode
	toSerialize["ike_policy"] = o.IkePolicy
	toSerialize["ipsec_policy"] = o.IpsecPolicy
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

func (o *IPSecProfileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"mode",
		"ike_policy",
		"ipsec_policy",
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

	varIPSecProfileRequest := _IPSecProfileRequest{}

	err = json.Unmarshal(data, &varIPSecProfileRequest)

	if err != nil {
		return err
	}

	*o = IPSecProfileRequest(varIPSecProfileRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "ike_policy")
		delete(additionalProperties, "ipsec_policy")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPSecProfileRequest struct {
	value *IPSecProfileRequest
	isSet bool
}

func (v NullableIPSecProfileRequest) Get() *IPSecProfileRequest {
	return v.value
}

func (v *NullableIPSecProfileRequest) Set(val *IPSecProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProfileRequest(val *IPSecProfileRequest) *NullableIPSecProfileRequest {
	return &NullableIPSecProfileRequest{value: val, isSet: true}
}

func (v NullableIPSecProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


