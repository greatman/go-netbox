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

// checks if the ConsoleServerPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsoleServerPort{}

// ConsoleServerPort Adds support for custom fields and tags.
type ConsoleServerPort struct {
	Id int64 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Device BriefDevice `json:"device"`
	Module NullableBriefModule `json:"module,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *ConsolePortType `json:"type,omitempty"`
	Speed NullableConsolePortSpeed `json:"speed,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool `json:"mark_connected,omitempty"`
	Cable NullableBriefCable `json:"cable"`
	CableEnd string `json:"cable_end"`
	LinkPeers []interface{} `json:"link_peers"`
	// Return the type of the peer link terminations, or None.
	LinkPeersType NullableString `json:"link_peers_type"`
	ConnectedEndpoints []interface{} `json:"connected_endpoints"`
	ConnectedEndpointsType NullableString `json:"connected_endpoints_type"`
	ConnectedEndpointsReachable bool `json:"connected_endpoints_reachable"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *map[string]string `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Occupied bool `json:"_occupied"`
	AdditionalProperties map[string]interface{}
}

type _ConsoleServerPort ConsoleServerPort

// NewConsoleServerPort instantiates a new ConsoleServerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleServerPort(id int64, url string, displayUrl string, display string, device BriefDevice, name string, cable NullableBriefCable, cableEnd string, linkPeers []interface{}, linkPeersType NullableString, connectedEndpoints []interface{}, connectedEndpointsType NullableString, connectedEndpointsReachable bool, created NullableTime, lastUpdated NullableTime, occupied bool) *ConsoleServerPort {
	this := ConsoleServerPort{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Device = device
	this.Name = name
	this.Cable = cable
	this.CableEnd = cableEnd
	this.LinkPeers = linkPeers
	this.LinkPeersType = linkPeersType
	this.ConnectedEndpoints = connectedEndpoints
	this.ConnectedEndpointsType = connectedEndpointsType
	this.ConnectedEndpointsReachable = connectedEndpointsReachable
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Occupied = occupied
	return &this
}

// NewConsoleServerPortWithDefaults instantiates a new ConsoleServerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleServerPortWithDefaults() *ConsoleServerPort {
	this := ConsoleServerPort{}
	return &this
}

// GetId returns the Id field value
func (o *ConsoleServerPort) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsoleServerPort) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ConsoleServerPort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConsoleServerPort) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *ConsoleServerPort) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *ConsoleServerPort) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *ConsoleServerPort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConsoleServerPort) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *ConsoleServerPort) GetDevice() BriefDevice {
	if o == nil {
		var ret BriefDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ConsoleServerPort) SetDevice(v BriefDevice) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsoleServerPort) GetModule() BriefModule {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModule
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetModuleOk() (*BriefModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModule and assigns it to the Module field.
func (o *ConsoleServerPort) SetModule(v BriefModule) {
	o.Module.Set(&v)
}
// SetModuleNil sets the value for Module to be an explicit nil
func (o *ConsoleServerPort) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *ConsoleServerPort) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *ConsoleServerPort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsoleServerPort) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsoleServerPort) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConsoleServerPort) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsoleServerPort) GetType() ConsolePortType {
	if o == nil || IsNil(o.Type) {
		var ret ConsolePortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetTypeOk() (*ConsolePortType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsolePortType and assigns it to the Type field.
func (o *ConsoleServerPort) SetType(v ConsolePortType) {
	o.Type = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsoleServerPort) GetSpeed() ConsolePortSpeed {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret ConsolePortSpeed
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetSpeedOk() (*ConsolePortSpeed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableConsolePortSpeed and assigns it to the Speed field.
func (o *ConsoleServerPort) SetSpeed(v ConsolePortSpeed) {
	o.Speed.Set(&v)
}
// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *ConsoleServerPort) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *ConsoleServerPort) UnsetSpeed() {
	o.Speed.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsoleServerPort) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsoleServerPort) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *ConsoleServerPort) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *ConsoleServerPort) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetCable returns the Cable field value
// If the value is explicit nil, the zero value for BriefCable will be returned
func (o *ConsoleServerPort) GetCable() BriefCable {
	if o == nil || o.Cable.Get() == nil {
		var ret BriefCable
		return ret
	}

	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetCableOk() (*BriefCable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// SetCable sets field value
func (o *ConsoleServerPort) SetCable(v BriefCable) {
	o.Cable.Set(&v)
}

// GetCableEnd returns the CableEnd field value
func (o *ConsoleServerPort) GetCableEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CableEnd
}

// GetCableEndOk returns a tuple with the CableEnd field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetCableEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CableEnd, true
}

// SetCableEnd sets field value
func (o *ConsoleServerPort) SetCableEnd(v string) {
	o.CableEnd = v
}

// GetLinkPeers returns the LinkPeers field value
func (o *ConsoleServerPort) GetLinkPeers() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.LinkPeers
}

// GetLinkPeersOk returns a tuple with the LinkPeers field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetLinkPeersOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkPeers, true
}

// SetLinkPeers sets field value
func (o *ConsoleServerPort) SetLinkPeers(v []interface{}) {
	o.LinkPeers = v
}

// GetLinkPeersType returns the LinkPeersType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConsoleServerPort) GetLinkPeersType() string {
	if o == nil || o.LinkPeersType.Get() == nil {
		var ret string
		return ret
	}

	return *o.LinkPeersType.Get()
}

// GetLinkPeersTypeOk returns a tuple with the LinkPeersType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetLinkPeersTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkPeersType.Get(), o.LinkPeersType.IsSet()
}

// SetLinkPeersType sets field value
func (o *ConsoleServerPort) SetLinkPeersType(v string) {
	o.LinkPeersType.Set(&v)
}

// GetConnectedEndpoints returns the ConnectedEndpoints field value
// If the value is explicit nil, the zero value for []interface{} will be returned
func (o *ConsoleServerPort) GetConnectedEndpoints() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.ConnectedEndpoints
}

// GetConnectedEndpointsOk returns a tuple with the ConnectedEndpoints field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetConnectedEndpointsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ConnectedEndpoints) {
		return nil, false
	}
	return o.ConnectedEndpoints, true
}

// SetConnectedEndpoints sets field value
func (o *ConsoleServerPort) SetConnectedEndpoints(v []interface{}) {
	o.ConnectedEndpoints = v
}

// GetConnectedEndpointsType returns the ConnectedEndpointsType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConsoleServerPort) GetConnectedEndpointsType() string {
	if o == nil || o.ConnectedEndpointsType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ConnectedEndpointsType.Get()
}

// GetConnectedEndpointsTypeOk returns a tuple with the ConnectedEndpointsType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetConnectedEndpointsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointsType.Get(), o.ConnectedEndpointsType.IsSet()
}

// SetConnectedEndpointsType sets field value
func (o *ConsoleServerPort) SetConnectedEndpointsType(v string) {
	o.ConnectedEndpointsType.Set(&v)
}

// GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable field value
func (o *ConsoleServerPort) GetConnectedEndpointsReachable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ConnectedEndpointsReachable
}

// GetConnectedEndpointsReachableOk returns a tuple with the ConnectedEndpointsReachable field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetConnectedEndpointsReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectedEndpointsReachable, true
}

// SetConnectedEndpointsReachable sets field value
func (o *ConsoleServerPort) SetConnectedEndpointsReachable(v bool) {
	o.ConnectedEndpointsReachable = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConsoleServerPort) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ConsoleServerPort) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ConsoleServerPort) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ConsoleServerPort) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *ConsoleServerPort) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConsoleServerPort) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ConsoleServerPort) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ConsoleServerPort) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPort) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ConsoleServerPort) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetOccupied returns the Occupied field value
func (o *ConsoleServerPort) GetOccupied() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPort) GetOccupiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupied, true
}

// SetOccupied sets field value
func (o *ConsoleServerPort) SetOccupied(v bool) {
	o.Occupied = v
}

func (o ConsoleServerPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsoleServerPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	toSerialize["cable"] = o.Cable.Get()
	toSerialize["cable_end"] = o.CableEnd
	toSerialize["link_peers"] = o.LinkPeers
	toSerialize["link_peers_type"] = o.LinkPeersType.Get()
	if o.ConnectedEndpoints != nil {
		toSerialize["connected_endpoints"] = o.ConnectedEndpoints
	}
	toSerialize["connected_endpoints_type"] = o.ConnectedEndpointsType.Get()
	toSerialize["connected_endpoints_reachable"] = o.ConnectedEndpointsReachable
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["_occupied"] = o.Occupied

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConsoleServerPort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"device",
		"name",
		"cable",
		"cable_end",
		"link_peers",
		"link_peers_type",
		"connected_endpoints",
		"connected_endpoints_type",
		"connected_endpoints_reachable",
		"created",
		"last_updated",
		"_occupied",
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

	varConsoleServerPort := _ConsoleServerPort{}

	err = json.Unmarshal(data, &varConsoleServerPort)

	if err != nil {
		return err
	}

	*o = ConsoleServerPort(varConsoleServerPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "cable_end")
		delete(additionalProperties, "link_peers")
		delete(additionalProperties, "link_peers_type")
		delete(additionalProperties, "connected_endpoints")
		delete(additionalProperties, "connected_endpoints_type")
		delete(additionalProperties, "connected_endpoints_reachable")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "_occupied")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsoleServerPort struct {
	value *ConsoleServerPort
	isSet bool
}

func (v NullableConsoleServerPort) Get() *ConsoleServerPort {
	return v.value
}

func (v *NullableConsoleServerPort) Set(val *ConsoleServerPort) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleServerPort) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleServerPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleServerPort(val *ConsoleServerPort) *NullableConsoleServerPort {
	return &NullableConsoleServerPort{value: val, isSet: true}
}

func (v NullableConsoleServerPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleServerPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


