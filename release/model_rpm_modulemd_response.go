/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest/release/v3

import (
	"encoding/json"
	"time"
)

// checks if the RpmModulemdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmModulemdResponse{}

// RpmModulemdResponse Modulemd serializer.
type RpmModulemdResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
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
	// Modulemd profiles.
	Profiles map[string]interface{} `json:"profiles"`
	// Description of module.
	Description string `json:"description"`
}

// NewRpmModulemdResponse instantiates a new RpmModulemdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmModulemdResponse(name string, stream string, version string, context string, arch string, artifacts map[string]interface{}, dependencies map[string]interface{}, profiles map[string]interface{}, description string) *RpmModulemdResponse {
	this := RpmModulemdResponse{}
	this.Name = name
	this.Stream = stream
	this.Version = version
	this.Context = context
	this.Arch = arch
	this.Artifacts = artifacts
	this.Dependencies = dependencies
	this.Profiles = profiles
	this.Description = description
	return &this
}

// NewRpmModulemdResponseWithDefaults instantiates a new RpmModulemdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmModulemdResponseWithDefaults() *RpmModulemdResponse {
	this := RpmModulemdResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmModulemdResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmModulemdResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmModulemdResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmModulemdResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmModulemdResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmModulemdResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *RpmModulemdResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmModulemdResponse) SetName(v string) {
	o.Name = v
}

// GetStream returns the Stream field value
func (o *RpmModulemdResponse) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *RpmModulemdResponse) SetStream(v string) {
	o.Stream = v
}

// GetVersion returns the Version field value
func (o *RpmModulemdResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RpmModulemdResponse) SetVersion(v string) {
	o.Version = v
}

// GetStaticContext returns the StaticContext field value if set, zero value otherwise.
func (o *RpmModulemdResponse) GetStaticContext() bool {
	if o == nil || IsNil(o.StaticContext) {
		var ret bool
		return ret
	}
	return *o.StaticContext
}

// GetStaticContextOk returns a tuple with the StaticContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetStaticContextOk() (*bool, bool) {
	if o == nil || IsNil(o.StaticContext) {
		return nil, false
	}
	return o.StaticContext, true
}

// HasStaticContext returns a boolean if a field has been set.
func (o *RpmModulemdResponse) HasStaticContext() bool {
	if o != nil && !IsNil(o.StaticContext) {
		return true
	}

	return false
}

// SetStaticContext gets a reference to the given bool and assigns it to the StaticContext field.
func (o *RpmModulemdResponse) SetStaticContext(v bool) {
	o.StaticContext = &v
}

// GetContext returns the Context field value
func (o *RpmModulemdResponse) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *RpmModulemdResponse) SetContext(v string) {
	o.Context = v
}

// GetArch returns the Arch field value
func (o *RpmModulemdResponse) GetArch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arch
}

// GetArchOk returns a tuple with the Arch field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetArchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arch, true
}

// SetArch sets field value
func (o *RpmModulemdResponse) SetArch(v string) {
	o.Arch = v
}

// GetArtifacts returns the Artifacts field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmModulemdResponse) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdResponse) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Artifacts) {
		return map[string]interface{}{}, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *RpmModulemdResponse) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

// GetDependencies returns the Dependencies field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmModulemdResponse) GetDependencies() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdResponse) GetDependenciesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return map[string]interface{}{}, false
	}
	return o.Dependencies, true
}

// SetDependencies sets field value
func (o *RpmModulemdResponse) SetDependencies(v map[string]interface{}) {
	o.Dependencies = v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *RpmModulemdResponse) GetPackages() []*string {
	if o == nil || IsNil(o.Packages) {
		var ret []*string
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetPackagesOk() ([]*string, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *RpmModulemdResponse) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []*string and assigns it to the Packages field.
func (o *RpmModulemdResponse) SetPackages(v []*string) {
	o.Packages = v
}

// GetProfiles returns the Profiles field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmModulemdResponse) GetProfiles() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdResponse) GetProfilesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profiles) {
		return map[string]interface{}{}, false
	}
	return o.Profiles, true
}

// SetProfiles sets field value
func (o *RpmModulemdResponse) SetProfiles(v map[string]interface{}) {
	o.Profiles = v
}

// GetDescription returns the Description field value
func (o *RpmModulemdResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RpmModulemdResponse) SetDescription(v string) {
	o.Description = v
}

func (o RpmModulemdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmModulemdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
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
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableRpmModulemdResponse struct {
	value *RpmModulemdResponse
	isSet bool
}

func (v NullableRpmModulemdResponse) Get() *RpmModulemdResponse {
	return v.value
}

func (v *NullableRpmModulemdResponse) Set(val *RpmModulemdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmModulemdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmModulemdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmModulemdResponse(val *RpmModulemdResponse) *NullableRpmModulemdResponse {
	return &NullableRpmModulemdResponse{value: val, isSet: true}
}

func (v NullableRpmModulemdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmModulemdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


