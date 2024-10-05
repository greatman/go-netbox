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

// checks if the PatchedWritableCircuitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableCircuitRequest{}

// PatchedWritableCircuitRequest Adds support for custom fields and tags.
type PatchedWritableCircuitRequest struct {
	// Unique circuit ID
	Cid *string `json:"cid,omitempty"`
	Provider *BriefProviderRequest `json:"provider,omitempty"`
	ProviderAccount NullableBriefProviderAccountRequest `json:"provider_account,omitempty"`
	Type *BriefCircuitTypeRequest `json:"type,omitempty"`
	Status *CircuitStatusValue `json:"status,omitempty"`
	Tenant NullableBriefTenantRequest `json:"tenant,omitempty"`
	InstallDate NullableString `json:"install_date,omitempty"`
	TerminationDate NullableString `json:"termination_date,omitempty"`
	// Committed rate
	CommitRate NullableInt32 `json:"commit_rate,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	Assignments []BriefCircuitGroupAssignmentSerializerRequest `json:"assignments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableCircuitRequest PatchedWritableCircuitRequest

// NewPatchedWritableCircuitRequest instantiates a new PatchedWritableCircuitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCircuitRequest() *PatchedWritableCircuitRequest {
	this := PatchedWritableCircuitRequest{}
	return &this
}

// NewPatchedWritableCircuitRequestWithDefaults instantiates a new PatchedWritableCircuitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCircuitRequestWithDefaults() *PatchedWritableCircuitRequest {
	this := PatchedWritableCircuitRequest{}
	return &this
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetCid() string {
	if o == nil || IsNil(o.Cid) {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetCidOk() (*string, bool) {
	if o == nil || IsNil(o.Cid) {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasCid() bool {
	if o != nil && !IsNil(o.Cid) {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *PatchedWritableCircuitRequest) SetCid(v string) {
	o.Cid = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetProvider() BriefProviderRequest {
	if o == nil || IsNil(o.Provider) {
		var ret BriefProviderRequest
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetProviderOk() (*BriefProviderRequest, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given BriefProviderRequest and assigns it to the Provider field.
func (o *PatchedWritableCircuitRequest) SetProvider(v BriefProviderRequest) {
	o.Provider = &v
}

// GetProviderAccount returns the ProviderAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitRequest) GetProviderAccount() BriefProviderAccountRequest {
	if o == nil || IsNil(o.ProviderAccount.Get()) {
		var ret BriefProviderAccountRequest
		return ret
	}
	return *o.ProviderAccount.Get()
}

// GetProviderAccountOk returns a tuple with the ProviderAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitRequest) GetProviderAccountOk() (*BriefProviderAccountRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderAccount.Get(), o.ProviderAccount.IsSet()
}

// HasProviderAccount returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasProviderAccount() bool {
	if o != nil && o.ProviderAccount.IsSet() {
		return true
	}

	return false
}

// SetProviderAccount gets a reference to the given NullableBriefProviderAccountRequest and assigns it to the ProviderAccount field.
func (o *PatchedWritableCircuitRequest) SetProviderAccount(v BriefProviderAccountRequest) {
	o.ProviderAccount.Set(&v)
}
// SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil
func (o *PatchedWritableCircuitRequest) SetProviderAccountNil() {
	o.ProviderAccount.Set(nil)
}

// UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
func (o *PatchedWritableCircuitRequest) UnsetProviderAccount() {
	o.ProviderAccount.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetType() BriefCircuitTypeRequest {
	if o == nil || IsNil(o.Type) {
		var ret BriefCircuitTypeRequest
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetTypeOk() (*BriefCircuitTypeRequest, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BriefCircuitTypeRequest and assigns it to the Type field.
func (o *PatchedWritableCircuitRequest) SetType(v BriefCircuitTypeRequest) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetStatus() CircuitStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret CircuitStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetStatusOk() (*CircuitStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CircuitStatusValue and assigns it to the Status field.
func (o *PatchedWritableCircuitRequest) SetStatus(v CircuitStatusValue) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *PatchedWritableCircuitRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableCircuitRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableCircuitRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitRequest) GetInstallDate() string {
	if o == nil || IsNil(o.InstallDate.Get()) {
		var ret string
		return ret
	}
	return *o.InstallDate.Get()
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitRequest) GetInstallDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallDate.Get(), o.InstallDate.IsSet()
}

// HasInstallDate returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasInstallDate() bool {
	if o != nil && o.InstallDate.IsSet() {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given NullableString and assigns it to the InstallDate field.
func (o *PatchedWritableCircuitRequest) SetInstallDate(v string) {
	o.InstallDate.Set(&v)
}
// SetInstallDateNil sets the value for InstallDate to be an explicit nil
func (o *PatchedWritableCircuitRequest) SetInstallDateNil() {
	o.InstallDate.Set(nil)
}

// UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
func (o *PatchedWritableCircuitRequest) UnsetInstallDate() {
	o.InstallDate.Unset()
}

// GetTerminationDate returns the TerminationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitRequest) GetTerminationDate() string {
	if o == nil || IsNil(o.TerminationDate.Get()) {
		var ret string
		return ret
	}
	return *o.TerminationDate.Get()
}

// GetTerminationDateOk returns a tuple with the TerminationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitRequest) GetTerminationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationDate.Get(), o.TerminationDate.IsSet()
}

// HasTerminationDate returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasTerminationDate() bool {
	if o != nil && o.TerminationDate.IsSet() {
		return true
	}

	return false
}

// SetTerminationDate gets a reference to the given NullableString and assigns it to the TerminationDate field.
func (o *PatchedWritableCircuitRequest) SetTerminationDate(v string) {
	o.TerminationDate.Set(&v)
}
// SetTerminationDateNil sets the value for TerminationDate to be an explicit nil
func (o *PatchedWritableCircuitRequest) SetTerminationDateNil() {
	o.TerminationDate.Set(nil)
}

// UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
func (o *PatchedWritableCircuitRequest) UnsetTerminationDate() {
	o.TerminationDate.Unset()
}

// GetCommitRate returns the CommitRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitRequest) GetCommitRate() int32 {
	if o == nil || IsNil(o.CommitRate.Get()) {
		var ret int32
		return ret
	}
	return *o.CommitRate.Get()
}

// GetCommitRateOk returns a tuple with the CommitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitRequest) GetCommitRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitRate.Get(), o.CommitRate.IsSet()
}

// HasCommitRate returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasCommitRate() bool {
	if o != nil && o.CommitRate.IsSet() {
		return true
	}

	return false
}

// SetCommitRate gets a reference to the given NullableInt32 and assigns it to the CommitRate field.
func (o *PatchedWritableCircuitRequest) SetCommitRate(v int32) {
	o.CommitRate.Set(&v)
}
// SetCommitRateNil sets the value for CommitRate to be an explicit nil
func (o *PatchedWritableCircuitRequest) SetCommitRateNil() {
	o.CommitRate.Set(nil)
}

// UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
func (o *PatchedWritableCircuitRequest) UnsetCommitRate() {
	o.CommitRate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableCircuitRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableCircuitRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableCircuitRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *PatchedWritableCircuitRequest) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *PatchedWritableCircuitRequest) GetAssignments() []BriefCircuitGroupAssignmentSerializerRequest {
	if o == nil || IsNil(o.Assignments) {
		var ret []BriefCircuitGroupAssignmentSerializerRequest
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitRequest) GetAssignmentsOk() ([]BriefCircuitGroupAssignmentSerializerRequest, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *PatchedWritableCircuitRequest) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []BriefCircuitGroupAssignmentSerializerRequest and assigns it to the Assignments field.
func (o *PatchedWritableCircuitRequest) SetAssignments(v []BriefCircuitGroupAssignmentSerializerRequest) {
	o.Assignments = v
}

func (o PatchedWritableCircuitRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableCircuitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cid) {
		toSerialize["cid"] = o.Cid
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if o.ProviderAccount.IsSet() {
		toSerialize["provider_account"] = o.ProviderAccount.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.InstallDate.IsSet() {
		toSerialize["install_date"] = o.InstallDate.Get()
	}
	if o.TerminationDate.IsSet() {
		toSerialize["termination_date"] = o.TerminationDate.Get()
	}
	if o.CommitRate.IsSet() {
		toSerialize["commit_rate"] = o.CommitRate.Get()
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
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableCircuitRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableCircuitRequest := _PatchedWritableCircuitRequest{}

	err = json.Unmarshal(data, &varPatchedWritableCircuitRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableCircuitRequest(varPatchedWritableCircuitRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "provider_account")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "install_date")
		delete(additionalProperties, "termination_date")
		delete(additionalProperties, "commit_rate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "assignments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableCircuitRequest struct {
	value *PatchedWritableCircuitRequest
	isSet bool
}

func (v NullablePatchedWritableCircuitRequest) Get() *PatchedWritableCircuitRequest {
	return v.value
}

func (v *NullablePatchedWritableCircuitRequest) Set(val *PatchedWritableCircuitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCircuitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCircuitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCircuitRequest(val *PatchedWritableCircuitRequest) *NullablePatchedWritableCircuitRequest {
	return &NullablePatchedWritableCircuitRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableCircuitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCircuitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


