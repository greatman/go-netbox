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

// checks if the IPSecProposalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecProposalRequest{}

// IPSecProposalRequest Adds support for custom fields and tags.
type IPSecProposalRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	EncryptionAlgorithm IKEProposalEncryptionAlgorithmValue `json:"encryption_algorithm"`
	AuthenticationAlgorithm IKEProposalAuthenticationAlgorithmValue `json:"authentication_algorithm"`
	// Security association lifetime (seconds)
	SaLifetimeSeconds NullableInt64 `json:"sa_lifetime_seconds,omitempty"`
	// Security association lifetime (in kilobytes)
	SaLifetimeData NullableInt64 `json:"sa_lifetime_data,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPSecProposalRequest IPSecProposalRequest

// NewIPSecProposalRequest instantiates a new IPSecProposalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecProposalRequest(name string, encryptionAlgorithm IKEProposalEncryptionAlgorithmValue, authenticationAlgorithm IKEProposalAuthenticationAlgorithmValue) *IPSecProposalRequest {
	this := IPSecProposalRequest{}
	this.Name = name
	this.EncryptionAlgorithm = encryptionAlgorithm
	this.AuthenticationAlgorithm = authenticationAlgorithm
	return &this
}

// NewIPSecProposalRequestWithDefaults instantiates a new IPSecProposalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecProposalRequestWithDefaults() *IPSecProposalRequest {
	this := IPSecProposalRequest{}
	return &this
}

// GetName returns the Name field value
func (o *IPSecProposalRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IPSecProposalRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IPSecProposalRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPSecProposalRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposalRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPSecProposalRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPSecProposalRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value
func (o *IPSecProposalRequest) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithmValue {
	if o == nil {
		var ret IKEProposalEncryptionAlgorithmValue
		return ret
	}

	return o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *IPSecProposalRequest) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithmValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionAlgorithm, true
}

// SetEncryptionAlgorithm sets field value
func (o *IPSecProposalRequest) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithmValue) {
	o.EncryptionAlgorithm = v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value
func (o *IPSecProposalRequest) GetAuthenticationAlgorithm() IKEProposalAuthenticationAlgorithmValue {
	if o == nil {
		var ret IKEProposalAuthenticationAlgorithmValue
		return ret
	}

	return o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value
// and a boolean to check if the value has been set.
func (o *IPSecProposalRequest) GetAuthenticationAlgorithmOk() (*IKEProposalAuthenticationAlgorithmValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationAlgorithm, true
}

// SetAuthenticationAlgorithm sets field value
func (o *IPSecProposalRequest) SetAuthenticationAlgorithm(v IKEProposalAuthenticationAlgorithmValue) {
	o.AuthenticationAlgorithm = v
}

// GetSaLifetimeSeconds returns the SaLifetimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPSecProposalRequest) GetSaLifetimeSeconds() int64 {
	if o == nil || IsNil(o.SaLifetimeSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.SaLifetimeSeconds.Get()
}

// GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProposalRequest) GetSaLifetimeSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeSeconds.Get(), o.SaLifetimeSeconds.IsSet()
}

// HasSaLifetimeSeconds returns a boolean if a field has been set.
func (o *IPSecProposalRequest) HasSaLifetimeSeconds() bool {
	if o != nil && o.SaLifetimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeSeconds gets a reference to the given NullableInt64 and assigns it to the SaLifetimeSeconds field.
func (o *IPSecProposalRequest) SetSaLifetimeSeconds(v int64) {
	o.SaLifetimeSeconds.Set(&v)
}
// SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil
func (o *IPSecProposalRequest) SetSaLifetimeSecondsNil() {
	o.SaLifetimeSeconds.Set(nil)
}

// UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
func (o *IPSecProposalRequest) UnsetSaLifetimeSeconds() {
	o.SaLifetimeSeconds.Unset()
}

// GetSaLifetimeData returns the SaLifetimeData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPSecProposalRequest) GetSaLifetimeData() int64 {
	if o == nil || IsNil(o.SaLifetimeData.Get()) {
		var ret int64
		return ret
	}
	return *o.SaLifetimeData.Get()
}

// GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProposalRequest) GetSaLifetimeDataOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeData.Get(), o.SaLifetimeData.IsSet()
}

// HasSaLifetimeData returns a boolean if a field has been set.
func (o *IPSecProposalRequest) HasSaLifetimeData() bool {
	if o != nil && o.SaLifetimeData.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeData gets a reference to the given NullableInt64 and assigns it to the SaLifetimeData field.
func (o *IPSecProposalRequest) SetSaLifetimeData(v int64) {
	o.SaLifetimeData.Set(&v)
}
// SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil
func (o *IPSecProposalRequest) SetSaLifetimeDataNil() {
	o.SaLifetimeData.Set(nil)
}

// UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
func (o *IPSecProposalRequest) UnsetSaLifetimeData() {
	o.SaLifetimeData.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPSecProposalRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposalRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPSecProposalRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPSecProposalRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPSecProposalRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposalRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPSecProposalRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *IPSecProposalRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPSecProposalRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposalRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPSecProposalRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *IPSecProposalRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o IPSecProposalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecProposalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	if o.SaLifetimeSeconds.IsSet() {
		toSerialize["sa_lifetime_seconds"] = o.SaLifetimeSeconds.Get()
	}
	if o.SaLifetimeData.IsSet() {
		toSerialize["sa_lifetime_data"] = o.SaLifetimeData.Get()
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

func (o *IPSecProposalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"encryption_algorithm",
		"authentication_algorithm",
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

	varIPSecProposalRequest := _IPSecProposalRequest{}

	err = json.Unmarshal(data, &varIPSecProposalRequest)

	if err != nil {
		return err
	}

	*o = IPSecProposalRequest(varIPSecProposalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "sa_lifetime_seconds")
		delete(additionalProperties, "sa_lifetime_data")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPSecProposalRequest struct {
	value *IPSecProposalRequest
	isSet bool
}

func (v NullableIPSecProposalRequest) Get() *IPSecProposalRequest {
	return v.value
}

func (v *NullableIPSecProposalRequest) Set(val *IPSecProposalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProposalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProposalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProposalRequest(val *IPSecProposalRequest) *NullableIPSecProposalRequest {
	return &NullableIPSecProposalRequest{value: val, isSet: true}
}

func (v NullableIPSecProposalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProposalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


