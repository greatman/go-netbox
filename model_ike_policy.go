/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the IKEPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEPolicy{}

// IKEPolicy Adds support for custom fields and tags.
type IKEPolicy struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Version IKEPolicyVersion `json:"version"`
	Mode *IKEPolicyMode `json:"mode,omitempty"`
	Proposals []IKEProposal `json:"proposals,omitempty"`
	PresharedKey *string `json:"preshared_key,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _IKEPolicy IKEPolicy

// NewIKEPolicy instantiates a new IKEPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEPolicy(id int32, url string, displayUrl string, display string, name string, version IKEPolicyVersion, created NullableTime, lastUpdated NullableTime) *IKEPolicy {
	this := IKEPolicy{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Version = version
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewIKEPolicyWithDefaults instantiates a new IKEPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEPolicyWithDefaults() *IKEPolicy {
	this := IKEPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *IKEPolicy) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IKEPolicy) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *IKEPolicy) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IKEPolicy) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *IKEPolicy) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *IKEPolicy) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *IKEPolicy) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *IKEPolicy) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *IKEPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IKEPolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IKEPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IKEPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IKEPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *IKEPolicy) GetVersion() IKEPolicyVersion {
	if o == nil {
		var ret IKEPolicyVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetVersionOk() (*IKEPolicyVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *IKEPolicy) SetVersion(v IKEPolicyVersion) {
	o.Version = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *IKEPolicy) GetMode() IKEPolicyMode {
	if o == nil || IsNil(o.Mode) {
		var ret IKEPolicyMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetModeOk() (*IKEPolicyMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *IKEPolicy) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given IKEPolicyMode and assigns it to the Mode field.
func (o *IKEPolicy) SetMode(v IKEPolicyMode) {
	o.Mode = &v
}

// GetProposals returns the Proposals field value if set, zero value otherwise.
func (o *IKEPolicy) GetProposals() []IKEProposal {
	if o == nil || IsNil(o.Proposals) {
		var ret []IKEProposal
		return ret
	}
	return o.Proposals
}

// GetProposalsOk returns a tuple with the Proposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetProposalsOk() ([]IKEProposal, bool) {
	if o == nil || IsNil(o.Proposals) {
		return nil, false
	}
	return o.Proposals, true
}

// HasProposals returns a boolean if a field has been set.
func (o *IKEPolicy) HasProposals() bool {
	if o != nil && !IsNil(o.Proposals) {
		return true
	}

	return false
}

// SetProposals gets a reference to the given []IKEProposal and assigns it to the Proposals field.
func (o *IKEPolicy) SetProposals(v []IKEProposal) {
	o.Proposals = v
}

// GetPresharedKey returns the PresharedKey field value if set, zero value otherwise.
func (o *IKEPolicy) GetPresharedKey() string {
	if o == nil || IsNil(o.PresharedKey) {
		var ret string
		return ret
	}
	return *o.PresharedKey
}

// GetPresharedKeyOk returns a tuple with the PresharedKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetPresharedKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PresharedKey) {
		return nil, false
	}
	return o.PresharedKey, true
}

// HasPresharedKey returns a boolean if a field has been set.
func (o *IKEPolicy) HasPresharedKey() bool {
	if o != nil && !IsNil(o.PresharedKey) {
		return true
	}

	return false
}

// SetPresharedKey gets a reference to the given string and assigns it to the PresharedKey field.
func (o *IKEPolicy) SetPresharedKey(v string) {
	o.PresharedKey = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IKEPolicy) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IKEPolicy) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IKEPolicy) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IKEPolicy) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IKEPolicy) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *IKEPolicy) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IKEPolicy) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicy) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IKEPolicy) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *IKEPolicy) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IKEPolicy) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IKEPolicy) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *IKEPolicy) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IKEPolicy) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IKEPolicy) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *IKEPolicy) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o IKEPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Proposals) {
		toSerialize["proposals"] = o.Proposals
	}
	if !IsNil(o.PresharedKey) {
		toSerialize["preshared_key"] = o.PresharedKey
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IKEPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
		"version",
		"created",
		"last_updated",
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

	varIKEPolicy := _IKEPolicy{}

	err = json.Unmarshal(data, &varIKEPolicy)

	if err != nil {
		return err
	}

	*o = IKEPolicy(varIKEPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "version")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "proposals")
		delete(additionalProperties, "preshared_key")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIKEPolicy struct {
	value *IKEPolicy
	isSet bool
}

func (v NullableIKEPolicy) Get() *IKEPolicy {
	return v.value
}

func (v *NullableIKEPolicy) Set(val *IKEPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEPolicy(val *IKEPolicy) *NullableIKEPolicy {
	return &NullableIKEPolicy{value: val, isSet: true}
}

func (v NullableIKEPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


