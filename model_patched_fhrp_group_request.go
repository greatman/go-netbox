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

// checks if the PatchedFHRPGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedFHRPGroupRequest{}

// PatchedFHRPGroupRequest Adds support for custom fields and tags.
type PatchedFHRPGroupRequest struct {
	Name *string `json:"name,omitempty"`
	Protocol *BriefFHRPGroupProtocol `json:"protocol,omitempty"`
	GroupId *int32 `json:"group_id,omitempty"`
	AuthType *AuthenticationType2 `json:"auth_type,omitempty"`
	AuthKey *string `json:"auth_key,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedFHRPGroupRequest PatchedFHRPGroupRequest

// NewPatchedFHRPGroupRequest instantiates a new PatchedFHRPGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedFHRPGroupRequest() *PatchedFHRPGroupRequest {
	this := PatchedFHRPGroupRequest{}
	return &this
}

// NewPatchedFHRPGroupRequestWithDefaults instantiates a new PatchedFHRPGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedFHRPGroupRequestWithDefaults() *PatchedFHRPGroupRequest {
	this := PatchedFHRPGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedFHRPGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetProtocol() BriefFHRPGroupProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret BriefFHRPGroupProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetProtocolOk() (*BriefFHRPGroupProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given BriefFHRPGroupProtocol and assigns it to the Protocol field.
func (o *PatchedFHRPGroupRequest) SetProtocol(v BriefFHRPGroupProtocol) {
	o.Protocol = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetGroupId() int32 {
	if o == nil || IsNil(o.GroupId) {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *PatchedFHRPGroupRequest) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetAuthType() AuthenticationType2 {
	if o == nil || IsNil(o.AuthType) {
		var ret AuthenticationType2
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetAuthTypeOk() (*AuthenticationType2, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given AuthenticationType2 and assigns it to the AuthType field.
func (o *PatchedFHRPGroupRequest) SetAuthType(v AuthenticationType2) {
	o.AuthType = &v
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetAuthKey() string {
	if o == nil || IsNil(o.AuthKey) {
		var ret string
		return ret
	}
	return *o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetAuthKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuthKey) {
		return nil, false
	}
	return o.AuthKey, true
}

// HasAuthKey returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasAuthKey() bool {
	if o != nil && !IsNil(o.AuthKey) {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given string and assigns it to the AuthKey field.
func (o *PatchedFHRPGroupRequest) SetAuthKey(v string) {
	o.AuthKey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedFHRPGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedFHRPGroupRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedFHRPGroupRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedFHRPGroupRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFHRPGroupRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedFHRPGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *PatchedFHRPGroupRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o PatchedFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedFHRPGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.AuthKey) {
		toSerialize["auth_key"] = o.AuthKey
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

func (o *PatchedFHRPGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedFHRPGroupRequest := _PatchedFHRPGroupRequest{}

	err = json.Unmarshal(data, &varPatchedFHRPGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedFHRPGroupRequest(varPatchedFHRPGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "auth_key")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedFHRPGroupRequest struct {
	value *PatchedFHRPGroupRequest
	isSet bool
}

func (v NullablePatchedFHRPGroupRequest) Get() *PatchedFHRPGroupRequest {
	return v.value
}

func (v *NullablePatchedFHRPGroupRequest) Set(val *PatchedFHRPGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedFHRPGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedFHRPGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedFHRPGroupRequest(val *PatchedFHRPGroupRequest) *NullablePatchedFHRPGroupRequest {
	return &NullablePatchedFHRPGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedFHRPGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


