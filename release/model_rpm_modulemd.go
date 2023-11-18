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
	"fmt"
)

// checks if the RpmModulemd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmModulemd{}

// RpmModulemd Modulemd serializer.
type RpmModulemd struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// Modulemd name.
	Name string `json:"name"`
	// Stream name.
	Stream string `json:"stream"`
	// Modulemd version.
	Version string `json:"version"`
	// Modulemd static-context flag.
	StaticContext *bool `json:"static_context,omitempty"`
	// Modulemd context.
	Context string `json:"context"`
	// Modulemd architecture.
	Arch string `json:"arch"`
	// Modulemd artifacts.
	Artifacts map[string]interface{} `json:"artifacts"`
	// Modulemd dependencies.
	Dependencies map[string]interface{} `json:"dependencies"`
	// Modulemd artifacts' packages.
	Packages []*string `json:"packages,omitempty"`
	// Modulemd snippet
	Snippet string `json:"snippet"`
	// Modulemd profiles.
	Profiles map[string]interface{} `json:"profiles"`
	// Description of module.
	Description string `json:"description"`
}

type _RpmModulemd RpmModulemd

// NewRpmModulemd instantiates a new RpmModulemd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmModulemd(name string, stream string, version string, context string, arch string, artifacts map[string]interface{}, dependencies map[string]interface{}, snippet string, profiles map[string]interface{}, description string) *RpmModulemd {
	this := RpmModulemd{}
	this.Name = name
	this.Stream = stream
	this.Version = version
	this.Context = context
	this.Arch = arch
	this.Artifacts = artifacts
	this.Dependencies = dependencies
	this.Snippet = snippet
	this.Profiles = profiles
	this.Description = description
	return &this
}

// NewRpmModulemdWithDefaults instantiates a new RpmModulemd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmModulemdWithDefaults() *RpmModulemd {
	this := RpmModulemd{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RpmModulemd) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RpmModulemd) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *RpmModulemd) SetRepository(v string) {
	o.Repository = &v
}

// GetName returns the Name field value
func (o *RpmModulemd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmModulemd) SetName(v string) {
	o.Name = v
}

// GetStream returns the Stream field value
func (o *RpmModulemd) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *RpmModulemd) SetStream(v string) {
	o.Stream = v
}

// GetVersion returns the Version field value
func (o *RpmModulemd) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RpmModulemd) SetVersion(v string) {
	o.Version = v
}

// GetStaticContext returns the StaticContext field value if set, zero value otherwise.
func (o *RpmModulemd) GetStaticContext() bool {
	if o == nil || IsNil(o.StaticContext) {
		var ret bool
		return ret
	}
	return *o.StaticContext
}

// GetStaticContextOk returns a tuple with the StaticContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetStaticContextOk() (*bool, bool) {
	if o == nil || IsNil(o.StaticContext) {
		return nil, false
	}
	return o.StaticContext, true
}

// HasStaticContext returns a boolean if a field has been set.
func (o *RpmModulemd) HasStaticContext() bool {
	if o != nil && !IsNil(o.StaticContext) {
		return true
	}

	return false
}

// SetStaticContext gets a reference to the given bool and assigns it to the StaticContext field.
func (o *RpmModulemd) SetStaticContext(v bool) {
	o.StaticContext = &v
}

// GetContext returns the Context field value
func (o *RpmModulemd) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *RpmModulemd) SetContext(v string) {
	o.Context = v
}

// GetArch returns the Arch field value
func (o *RpmModulemd) GetArch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arch
}

// GetArchOk returns a tuple with the Arch field value
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetArchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arch, true
}

// SetArch sets field value
func (o *RpmModulemd) SetArch(v string) {
	o.Arch = v
}

// GetArtifacts returns the Artifacts field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmModulemd) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemd) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Artifacts) {
		return map[string]interface{}{}, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *RpmModulemd) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

// GetDependencies returns the Dependencies field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmModulemd) GetDependencies() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemd) GetDependenciesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return map[string]interface{}{}, false
	}
	return o.Dependencies, true
}

// SetDependencies sets field value
func (o *RpmModulemd) SetDependencies(v map[string]interface{}) {
	o.Dependencies = v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *RpmModulemd) GetPackages() []*string {
	if o == nil || IsNil(o.Packages) {
		var ret []*string
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetPackagesOk() ([]*string, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *RpmModulemd) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []*string and assigns it to the Packages field.
func (o *RpmModulemd) SetPackages(v []*string) {
	o.Packages = v
}

// GetSnippet returns the Snippet field value
func (o *RpmModulemd) GetSnippet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetSnippetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snippet, true
}

// SetSnippet sets field value
func (o *RpmModulemd) SetSnippet(v string) {
	o.Snippet = v
}

// GetProfiles returns the Profiles field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmModulemd) GetProfiles() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemd) GetProfilesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profiles) {
		return map[string]interface{}{}, false
	}
	return o.Profiles, true
}

// SetProfiles sets field value
func (o *RpmModulemd) SetProfiles(v map[string]interface{}) {
	o.Profiles = v
}

// GetDescription returns the Description field value
func (o *RpmModulemd) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RpmModulemd) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RpmModulemd) SetDescription(v string) {
	o.Description = v
}

func (o RpmModulemd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmModulemd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	toSerialize["name"] = o.Name
	toSerialize["stream"] = o.Stream
	toSerialize["version"] = o.Version
	if !IsNil(o.StaticContext) {
		toSerialize["static_context"] = o.StaticContext
	}
	toSerialize["context"] = o.Context
	toSerialize["arch"] = o.Arch
	if o.Artifacts != nil {
		toSerialize["artifacts"] = o.Artifacts
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	toSerialize["snippet"] = o.Snippet
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *RpmModulemd) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"stream",
		"version",
		"context",
		"arch",
		"artifacts",
		"dependencies",
		"snippet",
		"profiles",
		"description",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRpmModulemd := _RpmModulemd{}

	err = json.Unmarshal(bytes, &varRpmModulemd)

	if err != nil {
		return err
	}

	*o = RpmModulemd(varRpmModulemd)

	return err
}

type NullableRpmModulemd struct {
	value *RpmModulemd
	isSet bool
}

func (v NullableRpmModulemd) Get() *RpmModulemd {
	return v.value
}

func (v *NullableRpmModulemd) Set(val *RpmModulemd) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmModulemd) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmModulemd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmModulemd(val *RpmModulemd) *NullableRpmModulemd {
	return &NullableRpmModulemd{value: val, isSet: true}
}

func (v NullableRpmModulemd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmModulemd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


