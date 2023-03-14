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

// checks if the ContainerContainerDistributionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerContainerDistributionResponse{}

// ContainerContainerDistributionResponse A serializer for ContainerDistribution.
type ContainerContainerDistributionResponse struct {
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	PulpHref *string `json:"pulp_href,omitempty"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// An optional content-guard. If none is specified, a default one will be used.
	ContentGuard *string `json:"content_guard,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// RepositoryVersion to be served
	RepositoryVersion NullableString `json:"repository_version,omitempty"`
	// The Registry hostname/name/ to use with docker pull command defined by this distribution.
	RegistryPath *string `json:"registry_path,omitempty"`
	// Namespace this distribution belongs to.
	Namespace *string `json:"namespace,omitempty"`
	// Restrict pull access to explicitly authorized users. Defaults to unrestricted pull access.
	Private *bool `json:"private,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
}

// NewContainerContainerDistributionResponse instantiates a new ContainerContainerDistributionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerDistributionResponse(name string, basePath string) *ContainerContainerDistributionResponse {
	this := ContainerContainerDistributionResponse{}
	this.Name = name
	this.BasePath = basePath
	return &this
}

// NewContainerContainerDistributionResponseWithDefaults instantiates a new ContainerContainerDistributionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerDistributionResponseWithDefaults() *ContainerContainerDistributionResponse {
	this := ContainerContainerDistributionResponse{}
	return &this
}

// GetName returns the Name field value
func (o *ContainerContainerDistributionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerDistributionResponse) SetName(v string) {
	o.Name = v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerContainerDistributionResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerContainerDistributionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerDistributionResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerDistributionResponse) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *ContainerContainerDistributionResponse) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *ContainerContainerDistributionResponse) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *ContainerContainerDistributionResponse) UnsetRepository() {
	o.Repository.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *ContainerContainerDistributionResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *ContainerContainerDistributionResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetBasePath returns the BasePath field value
func (o *ContainerContainerDistributionResponse) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *ContainerContainerDistributionResponse) SetBasePath(v string) {
	o.BasePath = v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise.
func (o *ContainerContainerDistributionResponse) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard) {
		var ret string
		return ret
	}
	return *o.ContentGuard
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetContentGuardOk() (*string, bool) {
	if o == nil || IsNil(o.ContentGuard) {
		return nil, false
	}
	return o.ContentGuard, true
}

// HasContentGuard returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasContentGuard() bool {
	if o != nil && !IsNil(o.ContentGuard) {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given string and assigns it to the ContentGuard field.
func (o *ContainerContainerDistributionResponse) SetContentGuard(v string) {
	o.ContentGuard = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerContainerDistributionResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerContainerDistributionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerDistributionResponse) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion.Get()
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerDistributionResponse) GetRepositoryVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryVersion.Get(), o.RepositoryVersion.IsSet()
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion.IsSet() {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given NullableString and assigns it to the RepositoryVersion field.
func (o *ContainerContainerDistributionResponse) SetRepositoryVersion(v string) {
	o.RepositoryVersion.Set(&v)
}
// SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil
func (o *ContainerContainerDistributionResponse) SetRepositoryVersionNil() {
	o.RepositoryVersion.Set(nil)
}

// UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
func (o *ContainerContainerDistributionResponse) UnsetRepositoryVersion() {
	o.RepositoryVersion.Unset()
}

// GetRegistryPath returns the RegistryPath field value if set, zero value otherwise.
func (o *ContainerContainerDistributionResponse) GetRegistryPath() string {
	if o == nil || IsNil(o.RegistryPath) {
		var ret string
		return ret
	}
	return *o.RegistryPath
}

// GetRegistryPathOk returns a tuple with the RegistryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetRegistryPathOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryPath) {
		return nil, false
	}
	return o.RegistryPath, true
}

// HasRegistryPath returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasRegistryPath() bool {
	if o != nil && !IsNil(o.RegistryPath) {
		return true
	}

	return false
}

// SetRegistryPath gets a reference to the given string and assigns it to the RegistryPath field.
func (o *ContainerContainerDistributionResponse) SetRegistryPath(v string) {
	o.RegistryPath = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ContainerContainerDistributionResponse) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ContainerContainerDistributionResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ContainerContainerDistributionResponse) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerDistributionResponse) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ContainerContainerDistributionResponse) SetPrivate(v bool) {
	o.Private = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerDistributionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerDistributionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerContainerDistributionResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerContainerDistributionResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerContainerDistributionResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerContainerDistributionResponse) UnsetDescription() {
	o.Description.Unset()
}

func (o ContainerContainerDistributionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerContainerDistributionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	// skip: pulp_href is readOnly
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["base_path"] = o.BasePath
	if !IsNil(o.ContentGuard) {
		toSerialize["content_guard"] = o.ContentGuard
	}
	// skip: pulp_created is readOnly
	if o.RepositoryVersion.IsSet() {
		toSerialize["repository_version"] = o.RepositoryVersion.Get()
	}
	// skip: registry_path is readOnly
	// skip: namespace is readOnly
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableContainerContainerDistributionResponse struct {
	value *ContainerContainerDistributionResponse
	isSet bool
}

func (v NullableContainerContainerDistributionResponse) Get() *ContainerContainerDistributionResponse {
	return v.value
}

func (v *NullableContainerContainerDistributionResponse) Set(val *ContainerContainerDistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerDistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerDistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerDistributionResponse(val *ContainerContainerDistributionResponse) *NullableContainerContainerDistributionResponse {
	return &NullableContainerContainerDistributionResponse{value: val, isSet: true}
}

func (v NullableContainerContainerDistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerDistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


