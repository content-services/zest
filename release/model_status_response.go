/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponse{}

// StatusResponse Serializer for the status information of the app
type StatusResponse struct {
	// Version information of Pulp components
	Versions []VersionResponse `json:"versions"`
	// List of online workers known to the application. An online worker is actively heartbeating and can respond to new work
	OnlineWorkers []WorkerResponse `json:"online_workers"`
	// List of online content apps known to the application. An online content app is actively heartbeating and can serve data to clients
	OnlineContentApps []ContentAppStatusResponse `json:"online_content_apps"`
	DatabaseConnection StatusResponseDatabaseConnection `json:"database_connection"`
	RedisConnection *StatusResponseRedisConnection `json:"redis_connection,omitempty"`
	Storage *StatusResponseStorage `json:"storage,omitempty"`
	ContentSettings StatusResponseContentSettings `json:"content_settings"`
	// Is Domains enabled
	DomainEnabled bool `json:"domain_enabled"`
}

// NewStatusResponse instantiates a new StatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponse(versions []VersionResponse, onlineWorkers []WorkerResponse, onlineContentApps []ContentAppStatusResponse, databaseConnection StatusResponseDatabaseConnection, contentSettings StatusResponseContentSettings, domainEnabled bool) *StatusResponse {
	this := StatusResponse{}
	this.Versions = versions
	this.OnlineWorkers = onlineWorkers
	this.OnlineContentApps = onlineContentApps
	this.DatabaseConnection = databaseConnection
	this.ContentSettings = contentSettings
	this.DomainEnabled = domainEnabled
	return &this
}

// NewStatusResponseWithDefaults instantiates a new StatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseWithDefaults() *StatusResponse {
	this := StatusResponse{}
	return &this
}

// GetVersions returns the Versions field value
func (o *StatusResponse) GetVersions() []VersionResponse {
	if o == nil {
		var ret []VersionResponse
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetVersionsOk() ([]VersionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *StatusResponse) SetVersions(v []VersionResponse) {
	o.Versions = v
}

// GetOnlineWorkers returns the OnlineWorkers field value
func (o *StatusResponse) GetOnlineWorkers() []WorkerResponse {
	if o == nil {
		var ret []WorkerResponse
		return ret
	}

	return o.OnlineWorkers
}

// GetOnlineWorkersOk returns a tuple with the OnlineWorkers field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetOnlineWorkersOk() ([]WorkerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlineWorkers, true
}

// SetOnlineWorkers sets field value
func (o *StatusResponse) SetOnlineWorkers(v []WorkerResponse) {
	o.OnlineWorkers = v
}

// GetOnlineContentApps returns the OnlineContentApps field value
func (o *StatusResponse) GetOnlineContentApps() []ContentAppStatusResponse {
	if o == nil {
		var ret []ContentAppStatusResponse
		return ret
	}

	return o.OnlineContentApps
}

// GetOnlineContentAppsOk returns a tuple with the OnlineContentApps field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetOnlineContentAppsOk() ([]ContentAppStatusResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlineContentApps, true
}

// SetOnlineContentApps sets field value
func (o *StatusResponse) SetOnlineContentApps(v []ContentAppStatusResponse) {
	o.OnlineContentApps = v
}

// GetDatabaseConnection returns the DatabaseConnection field value
func (o *StatusResponse) GetDatabaseConnection() StatusResponseDatabaseConnection {
	if o == nil {
		var ret StatusResponseDatabaseConnection
		return ret
	}

	return o.DatabaseConnection
}

// GetDatabaseConnectionOk returns a tuple with the DatabaseConnection field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetDatabaseConnectionOk() (*StatusResponseDatabaseConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseConnection, true
}

// SetDatabaseConnection sets field value
func (o *StatusResponse) SetDatabaseConnection(v StatusResponseDatabaseConnection) {
	o.DatabaseConnection = v
}

// GetRedisConnection returns the RedisConnection field value if set, zero value otherwise.
func (o *StatusResponse) GetRedisConnection() StatusResponseRedisConnection {
	if o == nil || IsNil(o.RedisConnection) {
		var ret StatusResponseRedisConnection
		return ret
	}
	return *o.RedisConnection
}

// GetRedisConnectionOk returns a tuple with the RedisConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetRedisConnectionOk() (*StatusResponseRedisConnection, bool) {
	if o == nil || IsNil(o.RedisConnection) {
		return nil, false
	}
	return o.RedisConnection, true
}

// HasRedisConnection returns a boolean if a field has been set.
func (o *StatusResponse) HasRedisConnection() bool {
	if o != nil && !IsNil(o.RedisConnection) {
		return true
	}

	return false
}

// SetRedisConnection gets a reference to the given StatusResponseRedisConnection and assigns it to the RedisConnection field.
func (o *StatusResponse) SetRedisConnection(v StatusResponseRedisConnection) {
	o.RedisConnection = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *StatusResponse) GetStorage() StatusResponseStorage {
	if o == nil || IsNil(o.Storage) {
		var ret StatusResponseStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetStorageOk() (*StatusResponseStorage, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *StatusResponse) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given StatusResponseStorage and assigns it to the Storage field.
func (o *StatusResponse) SetStorage(v StatusResponseStorage) {
	o.Storage = &v
}

// GetContentSettings returns the ContentSettings field value
func (o *StatusResponse) GetContentSettings() StatusResponseContentSettings {
	if o == nil {
		var ret StatusResponseContentSettings
		return ret
	}

	return o.ContentSettings
}

// GetContentSettingsOk returns a tuple with the ContentSettings field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetContentSettingsOk() (*StatusResponseContentSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentSettings, true
}

// SetContentSettings sets field value
func (o *StatusResponse) SetContentSettings(v StatusResponseContentSettings) {
	o.ContentSettings = v
}

// GetDomainEnabled returns the DomainEnabled field value
func (o *StatusResponse) GetDomainEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DomainEnabled
}

// GetDomainEnabledOk returns a tuple with the DomainEnabled field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetDomainEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainEnabled, true
}

// SetDomainEnabled sets field value
func (o *StatusResponse) SetDomainEnabled(v bool) {
	o.DomainEnabled = v
}

func (o StatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versions"] = o.Versions
	toSerialize["online_workers"] = o.OnlineWorkers
	toSerialize["online_content_apps"] = o.OnlineContentApps
	toSerialize["database_connection"] = o.DatabaseConnection
	if !IsNil(o.RedisConnection) {
		toSerialize["redis_connection"] = o.RedisConnection
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	toSerialize["content_settings"] = o.ContentSettings
	toSerialize["domain_enabled"] = o.DomainEnabled
	return toSerialize, nil
}

type NullableStatusResponse struct {
	value *StatusResponse
	isSet bool
}

func (v NullableStatusResponse) Get() *StatusResponse {
	return v.value
}

func (v *NullableStatusResponse) Set(val *StatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponse(val *StatusResponse) *NullableStatusResponse {
	return &NullableStatusResponse{value: val, isSet: true}
}

func (v NullableStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


