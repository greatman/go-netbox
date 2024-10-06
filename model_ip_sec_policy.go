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

// checks if the IPSecPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecPolicy{}

// IPSecPolicy Adds support for custom fields and tags.
type IPSecPolicy struct {
	Id int64 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Proposals []IPSecProposal `json:"proposals,omitempty"`
	PfsGroup *IKEProposalGroup `json:"pfs_group,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _IPSecPolicy IPSecPolicy

// NewIPSecPolicy instantiates a new IPSecPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecPolicy(id int64, url string, displayUrl string, display string, name string, created NullableTime, lastUpdated NullableTime) *IPSecPolicy {
	this := IPSecPolicy{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewIPSecPolicyWithDefaults instantiates a new IPSecPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecPolicyWithDefaults() *IPSecPolicy {
	this := IPSecPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *IPSecPolicy) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IPSecPolicy) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *IPSecPolicy) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IPSecPolicy) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *IPSecPolicy) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *IPSecPolicy) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *IPSecPolicy) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *IPSecPolicy) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *IPSecPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IPSecPolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPSecPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPSecPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPSecPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetProposals returns the Proposals field value if set, zero value otherwise.
func (o *IPSecPolicy) GetProposals() []IPSecProposal {
	if o == nil || IsNil(o.Proposals) {
		var ret []IPSecProposal
		return ret
	}
	return o.Proposals
}

// GetProposalsOk returns a tuple with the Proposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetProposalsOk() ([]IPSecProposal, bool) {
	if o == nil || IsNil(o.Proposals) {
		return nil, false
	}
	return o.Proposals, true
}

// HasProposals returns a boolean if a field has been set.
func (o *IPSecPolicy) HasProposals() bool {
	if o != nil && !IsNil(o.Proposals) {
		return true
	}

	return false
}

// SetProposals gets a reference to the given []IPSecProposal and assigns it to the Proposals field.
func (o *IPSecPolicy) SetProposals(v []IPSecProposal) {
	o.Proposals = v
}

// GetPfsGroup returns the PfsGroup field value if set, zero value otherwise.
func (o *IPSecPolicy) GetPfsGroup() IKEProposalGroup {
	if o == nil || IsNil(o.PfsGroup) {
		var ret IKEProposalGroup
		return ret
	}
	return *o.PfsGroup
}

// GetPfsGroupOk returns a tuple with the PfsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetPfsGroupOk() (*IKEProposalGroup, bool) {
	if o == nil || IsNil(o.PfsGroup) {
		return nil, false
	}
	return o.PfsGroup, true
}

// HasPfsGroup returns a boolean if a field has been set.
func (o *IPSecPolicy) HasPfsGroup() bool {
	if o != nil && !IsNil(o.PfsGroup) {
		return true
	}

	return false
}

// SetPfsGroup gets a reference to the given IKEProposalGroup and assigns it to the PfsGroup field.
func (o *IPSecPolicy) SetPfsGroup(v IKEProposalGroup) {
	o.PfsGroup = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPSecPolicy) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPSecPolicy) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPSecPolicy) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPSecPolicy) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPSecPolicy) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *IPSecPolicy) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPSecPolicy) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicy) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPSecPolicy) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *IPSecPolicy) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPSecPolicy) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecPolicy) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *IPSecPolicy) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPSecPolicy) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecPolicy) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *IPSecPolicy) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o IPSecPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Proposals) {
		toSerialize["proposals"] = o.Proposals
	}
	if !IsNil(o.PfsGroup) {
		toSerialize["pfs_group"] = o.PfsGroup
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

func (o *IPSecPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
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

	varIPSecPolicy := _IPSecPolicy{}

	err = json.Unmarshal(data, &varIPSecPolicy)

	if err != nil {
		return err
	}

	*o = IPSecPolicy(varIPSecPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "proposals")
		delete(additionalProperties, "pfs_group")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPSecPolicy struct {
	value *IPSecPolicy
	isSet bool
}

func (v NullableIPSecPolicy) Get() *IPSecPolicy {
	return v.value
}

func (v *NullableIPSecPolicy) Set(val *IPSecPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecPolicy(val *IPSecPolicy) *NullableIPSecPolicy {
	return &NullableIPSecPolicy{value: val, isSet: true}
}

func (v NullableIPSecPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


