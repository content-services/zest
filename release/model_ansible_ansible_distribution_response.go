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
)

// checks if the AnsibleAnsibleDistributionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleAnsibleDistributionResponse{}

// AnsibleAnsibleDistributionResponse Serializer for Ansible Distributions.
type AnsibleAnsibleDistributionResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// RepositoryVersion to be served
	RepositoryVersion NullableString `json:"repository_version,omitempty"`
	// The URL of a Collection content source.
	ClientUrl *string `json:"client_url,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
}

// NewAnsibleAnsibleDistributionResponse instantiates a new AnsibleAnsibleDistributionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleAnsibleDistributionResponse(basePath string, name string) *AnsibleAnsibleDistributionResponse {
	this := AnsibleAnsibleDistributionResponse{}
	this.BasePath = basePath
	this.Name = name
	return &this
}

// NewAnsibleAnsibleDistributionResponseWithDefaults instantiates a new AnsibleAnsibleDistributionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleAnsibleDistributionResponseWithDefaults() *AnsibleAnsibleDistributionResponse {
	this := AnsibleAnsibleDistributionResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleAnsibleDistributionResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleDistributionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleAnsibleDistributionResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleAnsibleDistributionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleAnsibleDistributionResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleDistributionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleAnsibleDistributionResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleAnsibleDistributionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetBasePath returns the BasePath field value
func (o *AnsibleAnsibleDistributionResponse) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleDistributionResponse) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *AnsibleAnsibleDistributionResponse) SetBasePath(v string) {
	o.BasePath = v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleDistributionResponse) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleDistributionResponse) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *AnsibleAnsibleDistributionResponse) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *AnsibleAnsibleDistributionResponse) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *AnsibleAnsibleDistributionResponse) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *AnsibleAnsibleDistributionResponse) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetName returns the Name field value
func (o *AnsibleAnsibleDistributionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleDistributionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleAnsibleDistributionResponse) SetName(v string) {
	o.Name = v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleDistributionResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleDistributionResponse) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *AnsibleAnsibleDistributionResponse) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *AnsibleAnsibleDistributionResponse) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *AnsibleAnsibleDistributionResponse) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *AnsibleAnsibleDistributionResponse) UnsetRepository() {
	o.Repository.Unset()
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleAnsibleDistributionResponse) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion.Get()
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleAnsibleDistributionResponse) GetRepositoryVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryVersion.Get(), o.RepositoryVersion.IsSet()
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *AnsibleAnsibleDistributionResponse) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion.IsSet() {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given NullableString and assigns it to the RepositoryVersion field.
func (o *AnsibleAnsibleDistributionResponse) SetRepositoryVersion(v string) {
	o.RepositoryVersion.Set(&v)
}
// SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil
func (o *AnsibleAnsibleDistributionResponse) SetRepositoryVersionNil() {
	o.RepositoryVersion.Set(nil)
}

// UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
func (o *AnsibleAnsibleDistributionResponse) UnsetRepositoryVersion() {
	o.RepositoryVersion.Unset()
}

// GetClientUrl returns the ClientUrl field value if set, zero value otherwise.
func (o *AnsibleAnsibleDistributionResponse) GetClientUrl() string {
	if o == nil || IsNil(o.ClientUrl) {
		var ret string
		return ret
	}
	return *o.ClientUrl
}

// GetClientUrlOk returns a tuple with the ClientUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleDistributionResponse) GetClientUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ClientUrl) {
		return nil, false
	}
	return o.ClientUrl, true
}

// HasClientUrl returns a boolean if a field has been set.
func (o *AnsibleAnsibleDistributionResponse) HasClientUrl() bool {
	if o != nil && !IsNil(o.ClientUrl) {
		return true
	}

	return false
}

// SetClientUrl gets a reference to the given string and assigns it to the ClientUrl field.
func (o *AnsibleAnsibleDistributionResponse) SetClientUrl(v string) {
	o.ClientUrl = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *AnsibleAnsibleDistributionResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleAnsibleDistributionResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *AnsibleAnsibleDistributionResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *AnsibleAnsibleDistributionResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

func (o AnsibleAnsibleDistributionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleAnsibleDistributionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["base_path"] = o.BasePath
	if o.ContentGuard.IsSet() {
		toSerialize["content_guard"] = o.ContentGuard.Get()
	}
	toSerialize["name"] = o.Name
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.RepositoryVersion.IsSet() {
		toSerialize["repository_version"] = o.RepositoryVersion.Get()
	}
	// skip: client_url is readOnly
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	return toSerialize, nil
}

type NullableAnsibleAnsibleDistributionResponse struct {
	value *AnsibleAnsibleDistributionResponse
	isSet bool
}

func (v NullableAnsibleAnsibleDistributionResponse) Get() *AnsibleAnsibleDistributionResponse {
	return v.value
}

func (v *NullableAnsibleAnsibleDistributionResponse) Set(val *AnsibleAnsibleDistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleAnsibleDistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleAnsibleDistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleAnsibleDistributionResponse(val *AnsibleAnsibleDistributionResponse) *NullableAnsibleAnsibleDistributionResponse {
	return &NullableAnsibleAnsibleDistributionResponse{value: val, isSet: true}
}

func (v NullableAnsibleAnsibleDistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleAnsibleDistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


