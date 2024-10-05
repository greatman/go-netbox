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

// checks if the TunnelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelRequest{}

// TunnelRequest Adds support for custom fields and tags.
type TunnelRequest struct {
	Name string `json:"name"`
	Status PatchedWritableTunnelRequestStatus `json:"status"`
	Group NullableBriefTunnelGroupRequest `json:"group,omitempty"`
	Encapsulation PatchedWritableTunnelRequestEncapsulation `json:"encapsulation"`
	IpsecProfile NullableBriefIPSecProfileRequest `json:"ipsec_profile,omitempty"`
	Tenant NullableBriefTenantRequest `json:"tenant,omitempty"`
	TunnelId NullableInt64 `json:"tunnel_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelRequest TunnelRequest

// NewTunnelRequest instantiates a new TunnelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelRequest(name string, status PatchedWritableTunnelRequestStatus, encapsulation PatchedWritableTunnelRequestEncapsulation) *TunnelRequest {
	this := TunnelRequest{}
	this.Name = name
	this.Status = status
	this.Encapsulation = encapsulation
	return &this
}

// NewTunnelRequestWithDefaults instantiates a new TunnelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelRequestWithDefaults() *TunnelRequest {
	this := TunnelRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TunnelRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TunnelRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TunnelRequest) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *TunnelRequest) GetStatus() PatchedWritableTunnelRequestStatus {
	if o == nil {
		var ret PatchedWritableTunnelRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TunnelRequest) GetStatusOk() (*PatchedWritableTunnelRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TunnelRequest) SetStatus(v PatchedWritableTunnelRequestStatus) {
	o.Status = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunnelRequest) GetGroup() BriefTunnelGroupRequest {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefTunnelGroupRequest
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelRequest) GetGroupOk() (*BriefTunnelGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *TunnelRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefTunnelGroupRequest and assigns it to the Group field.
func (o *TunnelRequest) SetGroup(v BriefTunnelGroupRequest) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *TunnelRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *TunnelRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetEncapsulation returns the Encapsulation field value
func (o *TunnelRequest) GetEncapsulation() PatchedWritableTunnelRequestEncapsulation {
	if o == nil {
		var ret PatchedWritableTunnelRequestEncapsulation
		return ret
	}

	return o.Encapsulation
}

// GetEncapsulationOk returns a tuple with the Encapsulation field value
// and a boolean to check if the value has been set.
func (o *TunnelRequest) GetEncapsulationOk() (*PatchedWritableTunnelRequestEncapsulation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encapsulation, true
}

// SetEncapsulation sets field value
func (o *TunnelRequest) SetEncapsulation(v PatchedWritableTunnelRequestEncapsulation) {
	o.Encapsulation = v
}

// GetIpsecProfile returns the IpsecProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunnelRequest) GetIpsecProfile() BriefIPSecProfileRequest {
	if o == nil || IsNil(o.IpsecProfile.Get()) {
		var ret BriefIPSecProfileRequest
		return ret
	}
	return *o.IpsecProfile.Get()
}

// GetIpsecProfileOk returns a tuple with the IpsecProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelRequest) GetIpsecProfileOk() (*BriefIPSecProfileRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpsecProfile.Get(), o.IpsecProfile.IsSet()
}

// HasIpsecProfile returns a boolean if a field has been set.
func (o *TunnelRequest) HasIpsecProfile() bool {
	if o != nil && o.IpsecProfile.IsSet() {
		return true
	}

	return false
}

// SetIpsecProfile gets a reference to the given NullableBriefIPSecProfileRequest and assigns it to the IpsecProfile field.
func (o *TunnelRequest) SetIpsecProfile(v BriefIPSecProfileRequest) {
	o.IpsecProfile.Set(&v)
}
// SetIpsecProfileNil sets the value for IpsecProfile to be an explicit nil
func (o *TunnelRequest) SetIpsecProfileNil() {
	o.IpsecProfile.Set(nil)
}

// UnsetIpsecProfile ensures that no value is present for IpsecProfile, not even an explicit nil
func (o *TunnelRequest) UnsetIpsecProfile() {
	o.IpsecProfile.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunnelRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *TunnelRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *TunnelRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *TunnelRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *TunnelRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetTunnelId returns the TunnelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunnelRequest) GetTunnelId() int64 {
	if o == nil || IsNil(o.TunnelId.Get()) {
		var ret int64
		return ret
	}
	return *o.TunnelId.Get()
}

// GetTunnelIdOk returns a tuple with the TunnelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunnelRequest) GetTunnelIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TunnelId.Get(), o.TunnelId.IsSet()
}

// HasTunnelId returns a boolean if a field has been set.
func (o *TunnelRequest) HasTunnelId() bool {
	if o != nil && o.TunnelId.IsSet() {
		return true
	}

	return false
}

// SetTunnelId gets a reference to the given NullableInt64 and assigns it to the TunnelId field.
func (o *TunnelRequest) SetTunnelId(v int64) {
	o.TunnelId.Set(&v)
}
// SetTunnelIdNil sets the value for TunnelId to be an explicit nil
func (o *TunnelRequest) SetTunnelIdNil() {
	o.TunnelId.Set(nil)
}

// UnsetTunnelId ensures that no value is present for TunnelId, not even an explicit nil
func (o *TunnelRequest) UnsetTunnelId() {
	o.TunnelId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TunnelRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TunnelRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TunnelRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *TunnelRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *TunnelRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *TunnelRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TunnelRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TunnelRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *TunnelRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TunnelRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TunnelRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *TunnelRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o TunnelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	toSerialize["encapsulation"] = o.Encapsulation
	if o.IpsecProfile.IsSet() {
		toSerialize["ipsec_profile"] = o.IpsecProfile.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.TunnelId.IsSet() {
		toSerialize["tunnel_id"] = o.TunnelId.Get()
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

func (o *TunnelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"status",
		"encapsulation",
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

	varTunnelRequest := _TunnelRequest{}

	err = json.Unmarshal(data, &varTunnelRequest)

	if err != nil {
		return err
	}

	*o = TunnelRequest(varTunnelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "group")
		delete(additionalProperties, "encapsulation")
		delete(additionalProperties, "ipsec_profile")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "tunnel_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelRequest struct {
	value *TunnelRequest
	isSet bool
}

func (v NullableTunnelRequest) Get() *TunnelRequest {
	return v.value
}

func (v *NullableTunnelRequest) Set(val *TunnelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelRequest(val *TunnelRequest) *NullableTunnelRequest {
	return &NullableTunnelRequest{value: val, isSet: true}
}

func (v NullableTunnelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


