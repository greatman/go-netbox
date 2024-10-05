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

// checks if the Prefix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prefix{}

// Prefix Adds support for custom fields and tags.
type Prefix struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Family AggregateFamily `json:"family"`
	Prefix string `json:"prefix"`
	Site NullableBriefSite `json:"site,omitempty"`
	Vrf NullableBriefVRF `json:"vrf,omitempty"`
	Tenant NullableBriefTenant `json:"tenant,omitempty"`
	Vlan NullableBriefVLAN `json:"vlan,omitempty"`
	Status *PrefixStatus `json:"status,omitempty"`
	Role NullableBriefRole `json:"role,omitempty"`
	// All IP addresses within this prefix are considered usable
	IsPool *bool `json:"is_pool,omitempty"`
	// Treat as fully utilized
	MarkUtilized *bool `json:"mark_utilized,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Children int32 `json:"children"`
	Depth int32 `json:"_depth"`
	AdditionalProperties map[string]interface{}
}

type _Prefix Prefix

// NewPrefix instantiates a new Prefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrefix(id int32, url string, displayUrl string, display string, family AggregateFamily, prefix string, created NullableTime, lastUpdated NullableTime, children int32, depth int32) *Prefix {
	this := Prefix{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Family = family
	this.Prefix = prefix
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Children = children
	this.Depth = depth
	return &this
}

// NewPrefixWithDefaults instantiates a new Prefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrefixWithDefaults() *Prefix {
	this := Prefix{}
	return &this
}

// GetId returns the Id field value
func (o *Prefix) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Prefix) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Prefix) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Prefix) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Prefix) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Prefix) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *Prefix) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Prefix) SetDisplay(v string) {
	o.Display = v
}

// GetFamily returns the Family field value
func (o *Prefix) GetFamily() AggregateFamily {
	if o == nil {
		var ret AggregateFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetFamilyOk() (*AggregateFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *Prefix) SetFamily(v AggregateFamily) {
	o.Family = v
}

// GetPrefix returns the Prefix field value
func (o *Prefix) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *Prefix) SetPrefix(v string) {
	o.Prefix = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetSite() BriefSite {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSite
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetSiteOk() (*BriefSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *Prefix) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSite and assigns it to the Site field.
func (o *Prefix) SetSite(v BriefSite) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *Prefix) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *Prefix) UnsetSite() {
	o.Site.Unset()
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetVrf() BriefVRF {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BriefVRF
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetVrfOk() (*BriefVRF, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *Prefix) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBriefVRF and assigns it to the Vrf field.
func (o *Prefix) SetVrf(v BriefVRF) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *Prefix) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *Prefix) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Prefix) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *Prefix) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Prefix) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Prefix) UnsetTenant() {
	o.Tenant.Unset()
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetVlan() BriefVLAN {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret BriefVLAN
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetVlanOk() (*BriefVLAN, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *Prefix) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableBriefVLAN and assigns it to the Vlan field.
func (o *Prefix) SetVlan(v BriefVLAN) {
	o.Vlan.Set(&v)
}
// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *Prefix) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *Prefix) UnsetVlan() {
	o.Vlan.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Prefix) GetStatus() PrefixStatus {
	if o == nil || IsNil(o.Status) {
		var ret PrefixStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetStatusOk() (*PrefixStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Prefix) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PrefixStatus and assigns it to the Status field.
func (o *Prefix) SetStatus(v PrefixStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetRole() BriefRole {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetRoleOk() (*BriefRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *Prefix) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefRole and assigns it to the Role field.
func (o *Prefix) SetRole(v BriefRole) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *Prefix) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *Prefix) UnsetRole() {
	o.Role.Unset()
}

// GetIsPool returns the IsPool field value if set, zero value otherwise.
func (o *Prefix) GetIsPool() bool {
	if o == nil || IsNil(o.IsPool) {
		var ret bool
		return ret
	}
	return *o.IsPool
}

// GetIsPoolOk returns a tuple with the IsPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetIsPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPool) {
		return nil, false
	}
	return o.IsPool, true
}

// HasIsPool returns a boolean if a field has been set.
func (o *Prefix) HasIsPool() bool {
	if o != nil && !IsNil(o.IsPool) {
		return true
	}

	return false
}

// SetIsPool gets a reference to the given bool and assigns it to the IsPool field.
func (o *Prefix) SetIsPool(v bool) {
	o.IsPool = &v
}

// GetMarkUtilized returns the MarkUtilized field value if set, zero value otherwise.
func (o *Prefix) GetMarkUtilized() bool {
	if o == nil || IsNil(o.MarkUtilized) {
		var ret bool
		return ret
	}
	return *o.MarkUtilized
}

// GetMarkUtilizedOk returns a tuple with the MarkUtilized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetMarkUtilizedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkUtilized) {
		return nil, false
	}
	return o.MarkUtilized, true
}

// HasMarkUtilized returns a boolean if a field has been set.
func (o *Prefix) HasMarkUtilized() bool {
	if o != nil && !IsNil(o.MarkUtilized) {
		return true
	}

	return false
}

// SetMarkUtilized gets a reference to the given bool and assigns it to the MarkUtilized field.
func (o *Prefix) SetMarkUtilized(v bool) {
	o.MarkUtilized = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Prefix) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Prefix) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Prefix) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Prefix) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Prefix) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Prefix) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Prefix) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Prefix) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Prefix) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Prefix) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Prefix) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *Prefix) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Prefix) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Prefix) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Prefix) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Prefix) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetChildren returns the Children field value
func (o *Prefix) GetChildren() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetChildrenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Children, true
}

// SetChildren sets field value
func (o *Prefix) SetChildren(v int32) {
	o.Children = v
}

// GetDepth returns the Depth field value
func (o *Prefix) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *Prefix) SetDepth(v int32) {
	o.Depth = v
}

func (o Prefix) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Prefix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["family"] = o.Family
	toSerialize["prefix"] = o.Prefix
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.IsPool) {
		toSerialize["is_pool"] = o.IsPool
	}
	if !IsNil(o.MarkUtilized) {
		toSerialize["mark_utilized"] = o.MarkUtilized
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["children"] = o.Children
	toSerialize["_depth"] = o.Depth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Prefix) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"family",
		"prefix",
		"created",
		"last_updated",
		"children",
		"_depth",
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

	varPrefix := _Prefix{}

	err = json.Unmarshal(data, &varPrefix)

	if err != nil {
		return err
	}

	*o = Prefix(varPrefix)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "family")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "site")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "is_pool")
		delete(additionalProperties, "mark_utilized")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "children")
		delete(additionalProperties, "_depth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrefix struct {
	value *Prefix
	isSet bool
}

func (v NullablePrefix) Get() *Prefix {
	return v.value
}

func (v *NullablePrefix) Set(val *Prefix) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefix) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefix(val *Prefix) *NullablePrefix {
	return &NullablePrefix{value: val, isSet: true}
}

func (v NullablePrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


