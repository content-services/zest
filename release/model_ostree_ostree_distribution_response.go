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

// checks if the OstreeOstreeDistributionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeOstreeDistributionResponse{}

// OstreeOstreeDistributionResponse A Serializer class for an OSTree distribution.
type OstreeOstreeDistributionResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// The URL for accessing the publication as defined by this distribution.
	BaseUrl *string `json:"base_url,omitempty"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	// Timestamp since when the distributed content served by this distribution has not changed. If equals to `null`, no guarantee is provided about content changes.
	NoContentChangeSince *string `json:"no_content_change_since,omitempty"`
	// Whether this distribution should be shown in the content app.
	Hidden *bool `json:"hidden,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// RepositoryVersion to be served
	RepositoryVersion NullableString `json:"repository_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OstreeOstreeDistributionResponse OstreeOstreeDistributionResponse

// NewOstreeOstreeDistributionResponse instantiates a new OstreeOstreeDistributionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeOstreeDistributionResponse(basePath string, name string) *OstreeOstreeDistributionResponse {
	this := OstreeOstreeDistributionResponse{}
	this.BasePath = basePath
	var hidden bool = false
	this.Hidden = &hidden
	this.Name = name
	return &this
}

// NewOstreeOstreeDistributionResponseWithDefaults instantiates a new OstreeOstreeDistributionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeOstreeDistributionResponseWithDefaults() *OstreeOstreeDistributionResponse {
	this := OstreeOstreeDistributionResponse{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *OstreeOstreeDistributionResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *OstreeOstreeDistributionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *OstreeOstreeDistributionResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *OstreeOstreeDistributionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *OstreeOstreeDistributionResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *OstreeOstreeDistributionResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetBasePath returns the BasePath field value
func (o *OstreeOstreeDistributionResponse) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *OstreeOstreeDistributionResponse) SetBasePath(v string) {
	o.BasePath = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *OstreeOstreeDistributionResponse) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *OstreeOstreeDistributionResponse) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OstreeOstreeDistributionResponse) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OstreeOstreeDistributionResponse) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *OstreeOstreeDistributionResponse) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *OstreeOstreeDistributionResponse) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *OstreeOstreeDistributionResponse) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetNoContentChangeSince returns the NoContentChangeSince field value if set, zero value otherwise.
func (o *OstreeOstreeDistributionResponse) GetNoContentChangeSince() string {
	if o == nil || IsNil(o.NoContentChangeSince) {
		var ret string
		return ret
	}
	return *o.NoContentChangeSince
}

// GetNoContentChangeSinceOk returns a tuple with the NoContentChangeSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetNoContentChangeSinceOk() (*string, bool) {
	if o == nil || IsNil(o.NoContentChangeSince) {
		return nil, false
	}
	return o.NoContentChangeSince, true
}

// HasNoContentChangeSince returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasNoContentChangeSince() bool {
	if o != nil && !IsNil(o.NoContentChangeSince) {
		return true
	}

	return false
}

// SetNoContentChangeSince gets a reference to the given string and assigns it to the NoContentChangeSince field.
func (o *OstreeOstreeDistributionResponse) SetNoContentChangeSince(v string) {
	o.NoContentChangeSince = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *OstreeOstreeDistributionResponse) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *OstreeOstreeDistributionResponse) SetHidden(v bool) {
	o.Hidden = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *OstreeOstreeDistributionResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *OstreeOstreeDistributionResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *OstreeOstreeDistributionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeDistributionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OstreeOstreeDistributionResponse) SetName(v string) {
	o.Name = v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OstreeOstreeDistributionResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OstreeOstreeDistributionResponse) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *OstreeOstreeDistributionResponse) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *OstreeOstreeDistributionResponse) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *OstreeOstreeDistributionResponse) UnsetRepository() {
	o.Repository.Unset()
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OstreeOstreeDistributionResponse) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion.Get()
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OstreeOstreeDistributionResponse) GetRepositoryVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryVersion.Get(), o.RepositoryVersion.IsSet()
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *OstreeOstreeDistributionResponse) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion.IsSet() {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given NullableString and assigns it to the RepositoryVersion field.
func (o *OstreeOstreeDistributionResponse) SetRepositoryVersion(v string) {
	o.RepositoryVersion.Set(&v)
}
// SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil
func (o *OstreeOstreeDistributionResponse) SetRepositoryVersionNil() {
	o.RepositoryVersion.Set(nil)
}

// UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
func (o *OstreeOstreeDistributionResponse) UnsetRepositoryVersion() {
	o.RepositoryVersion.Unset()
}

func (o OstreeOstreeDistributionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeOstreeDistributionResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NoContentChangeSince) {
		toSerialize["no_content_change_since"] = o.NoContentChangeSince
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
	if o.RepositoryVersion.IsSet() {
		toSerialize["repository_version"] = o.RepositoryVersion.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OstreeOstreeDistributionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varOstreeOstreeDistributionResponse := _OstreeOstreeDistributionResponse{}

	err = json.Unmarshal(data, &varOstreeOstreeDistributionResponse)

	if err != nil {
		return err
	}

	*o = OstreeOstreeDistributionResponse(varOstreeOstreeDistributionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "no_content_change_since")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "repository_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOstreeOstreeDistributionResponse struct {
	value *OstreeOstreeDistributionResponse
	isSet bool
}

func (v NullableOstreeOstreeDistributionResponse) Get() *OstreeOstreeDistributionResponse {
	return v.value
}

func (v *NullableOstreeOstreeDistributionResponse) Set(val *OstreeOstreeDistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeOstreeDistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeOstreeDistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeOstreeDistributionResponse(val *OstreeOstreeDistributionResponse) *NullableOstreeOstreeDistributionResponse {
	return &NullableOstreeOstreeDistributionResponse{value: val, isSet: true}
}

func (v NullableOstreeOstreeDistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeOstreeDistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


