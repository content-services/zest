/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the PythonPythonDistributionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PythonPythonDistributionResponse{}

// PythonPythonDistributionResponse Serializer for Pulp distributions for the Python type.
type PythonPythonDistributionResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	BaseUrl *string `json:"base_url,omitempty"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	// Whether this distribution should be shown in the content app.
	Hidden *bool `json:"hidden,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// Publication to be served
	Publication NullableString `json:"publication,omitempty"`
	// Allow packages to be uploaded to this index.
	AllowUploads *bool `json:"allow_uploads,omitempty"`
	// Remote that can be used to fetch content when using pull-through caching.
	Remote NullableString `json:"remote,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PythonPythonDistributionResponse PythonPythonDistributionResponse

// NewPythonPythonDistributionResponse instantiates a new PythonPythonDistributionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPythonPythonDistributionResponse(basePath string, name string) *PythonPythonDistributionResponse {
	this := PythonPythonDistributionResponse{}
	this.BasePath = basePath
	var hidden bool = false
	this.Hidden = &hidden
	this.Name = name
	var allowUploads bool = true
	this.AllowUploads = &allowUploads
	return &this
}

// NewPythonPythonDistributionResponseWithDefaults instantiates a new PythonPythonDistributionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPythonPythonDistributionResponseWithDefaults() *PythonPythonDistributionResponse {
	this := PythonPythonDistributionResponse{}
	var hidden bool = false
	this.Hidden = &hidden
	var allowUploads bool = true
	this.AllowUploads = &allowUploads
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *PythonPythonDistributionResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *PythonPythonDistributionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *PythonPythonDistributionResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *PythonPythonDistributionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *PythonPythonDistributionResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *PythonPythonDistributionResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetBasePath returns the BasePath field value
func (o *PythonPythonDistributionResponse) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *PythonPythonDistributionResponse) SetBasePath(v string) {
	o.BasePath = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *PythonPythonDistributionResponse) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *PythonPythonDistributionResponse) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonPythonDistributionResponse) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonPythonDistributionResponse) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *PythonPythonDistributionResponse) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *PythonPythonDistributionResponse) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *PythonPythonDistributionResponse) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *PythonPythonDistributionResponse) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *PythonPythonDistributionResponse) SetHidden(v bool) {
	o.Hidden = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PythonPythonDistributionResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PythonPythonDistributionResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *PythonPythonDistributionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PythonPythonDistributionResponse) SetName(v string) {
	o.Name = v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonPythonDistributionResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonPythonDistributionResponse) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *PythonPythonDistributionResponse) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *PythonPythonDistributionResponse) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *PythonPythonDistributionResponse) UnsetRepository() {
	o.Repository.Unset()
}

// GetPublication returns the Publication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonPythonDistributionResponse) GetPublication() string {
	if o == nil || IsNil(o.Publication.Get()) {
		var ret string
		return ret
	}
	return *o.Publication.Get()
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonPythonDistributionResponse) GetPublicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Publication.Get(), o.Publication.IsSet()
}

// HasPublication returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasPublication() bool {
	if o != nil && o.Publication.IsSet() {
		return true
	}

	return false
}

// SetPublication gets a reference to the given NullableString and assigns it to the Publication field.
func (o *PythonPythonDistributionResponse) SetPublication(v string) {
	o.Publication.Set(&v)
}
// SetPublicationNil sets the value for Publication to be an explicit nil
func (o *PythonPythonDistributionResponse) SetPublicationNil() {
	o.Publication.Set(nil)
}

// UnsetPublication ensures that no value is present for Publication, not even an explicit nil
func (o *PythonPythonDistributionResponse) UnsetPublication() {
	o.Publication.Unset()
}

// GetAllowUploads returns the AllowUploads field value if set, zero value otherwise.
func (o *PythonPythonDistributionResponse) GetAllowUploads() bool {
	if o == nil || IsNil(o.AllowUploads) {
		var ret bool
		return ret
	}
	return *o.AllowUploads
}

// GetAllowUploadsOk returns a tuple with the AllowUploads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonDistributionResponse) GetAllowUploadsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUploads) {
		return nil, false
	}
	return o.AllowUploads, true
}

// HasAllowUploads returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasAllowUploads() bool {
	if o != nil && !IsNil(o.AllowUploads) {
		return true
	}

	return false
}

// SetAllowUploads gets a reference to the given bool and assigns it to the AllowUploads field.
func (o *PythonPythonDistributionResponse) SetAllowUploads(v bool) {
	o.AllowUploads = &v
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonPythonDistributionResponse) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonPythonDistributionResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *PythonPythonDistributionResponse) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *PythonPythonDistributionResponse) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *PythonPythonDistributionResponse) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *PythonPythonDistributionResponse) UnsetRemote() {
	o.Remote.Unset()
}

func (o PythonPythonDistributionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PythonPythonDistributionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	toSerialize["base_path"] = o.BasePath
	if !IsNil(o.BaseUrl) {
		toSerialize["base_url"] = o.BaseUrl
	}
	if o.ContentGuard.IsSet() {
		toSerialize["content_guard"] = o.ContentGuard.Get()
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["name"] = o.Name
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.Publication.IsSet() {
		toSerialize["publication"] = o.Publication.Get()
	}
	if !IsNil(o.AllowUploads) {
		toSerialize["allow_uploads"] = o.AllowUploads
	}
	if o.Remote.IsSet() {
		toSerialize["remote"] = o.Remote.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PythonPythonDistributionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base_path",
		"name",
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

	varPythonPythonDistributionResponse := _PythonPythonDistributionResponse{}

	err = json.Unmarshal(data, &varPythonPythonDistributionResponse)

	if err != nil {
		return err
	}

	*o = PythonPythonDistributionResponse(varPythonPythonDistributionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "publication")
		delete(additionalProperties, "allow_uploads")
		delete(additionalProperties, "remote")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePythonPythonDistributionResponse struct {
	value *PythonPythonDistributionResponse
	isSet bool
}

func (v NullablePythonPythonDistributionResponse) Get() *PythonPythonDistributionResponse {
	return v.value
}

func (v *NullablePythonPythonDistributionResponse) Set(val *PythonPythonDistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePythonPythonDistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePythonPythonDistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePythonPythonDistributionResponse(val *PythonPythonDistributionResponse) *NullablePythonPythonDistributionResponse {
	return &NullablePythonPythonDistributionResponse{value: val, isSet: true}
}

func (v NullablePythonPythonDistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePythonPythonDistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


