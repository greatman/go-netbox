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

// checks if the ObjectPermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectPermissionRequest{}

// ObjectPermissionRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ObjectPermissionRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ObjectTypes []string `json:"object_types"`
	// The list of actions granted by this permission
	Actions []string `json:"actions"`
	// Queryset filter matching the applicable objects of the selected type(s)
	Constraints interface{} `json:"constraints,omitempty"`
	Groups []int64 `json:"groups,omitempty"`
	Users []int64 `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ObjectPermissionRequest ObjectPermissionRequest

// NewObjectPermissionRequest instantiates a new ObjectPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectPermissionRequest(name string, objectTypes []string, actions []string) *ObjectPermissionRequest {
	this := ObjectPermissionRequest{}
	this.Name = name
	this.ObjectTypes = objectTypes
	this.Actions = actions
	return &this
}

// NewObjectPermissionRequestWithDefaults instantiates a new ObjectPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectPermissionRequestWithDefaults() *ObjectPermissionRequest {
	this := ObjectPermissionRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ObjectPermissionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjectPermissionRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjectPermissionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPermissionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjectPermissionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjectPermissionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ObjectPermissionRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPermissionRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ObjectPermissionRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ObjectPermissionRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetObjectTypes returns the ObjectTypes field value
func (o *ObjectPermissionRequest) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissionRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *ObjectPermissionRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetActions returns the Actions field value
func (o *ObjectPermissionRequest) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissionRequest) GetActionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *ObjectPermissionRequest) SetActions(v []string) {
	o.Actions = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectPermissionRequest) GetConstraints() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectPermissionRequest) GetConstraintsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return &o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *ObjectPermissionRequest) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given interface{} and assigns it to the Constraints field.
func (o *ObjectPermissionRequest) SetConstraints(v interface{}) {
	o.Constraints = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ObjectPermissionRequest) GetGroups() []int64 {
	if o == nil || IsNil(o.Groups) {
		var ret []int64
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPermissionRequest) GetGroupsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ObjectPermissionRequest) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int64 and assigns it to the Groups field.
func (o *ObjectPermissionRequest) SetGroups(v []int64) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ObjectPermissionRequest) GetUsers() []int64 {
	if o == nil || IsNil(o.Users) {
		var ret []int64
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPermissionRequest) GetUsersOk() ([]int64, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ObjectPermissionRequest) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int64 and assigns it to the Users field.
func (o *ObjectPermissionRequest) SetUsers(v []int64) {
	o.Users = v
}

func (o ObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["object_types"] = o.ObjectTypes
	toSerialize["actions"] = o.Actions
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ObjectPermissionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"object_types",
		"actions",
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

	varObjectPermissionRequest := _ObjectPermissionRequest{}

	err = json.Unmarshal(data, &varObjectPermissionRequest)

	if err != nil {
		return err
	}

	*o = ObjectPermissionRequest(varObjectPermissionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "actions")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectPermissionRequest struct {
	value *ObjectPermissionRequest
	isSet bool
}

func (v NullableObjectPermissionRequest) Get() *ObjectPermissionRequest {
	return v.value
}

func (v *NullableObjectPermissionRequest) Set(val *ObjectPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectPermissionRequest(val *ObjectPermissionRequest) *NullableObjectPermissionRequest {
	return &NullableObjectPermissionRequest{value: val, isSet: true}
}

func (v NullableObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


