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

// checks if the GemGemContentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GemGemContentResponse{}

// GemGemContentResponse A Serializer for GemContent.
type GemGemContentResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// The Pulp Resource Name (PRN).
	Prn *string `json:"prn,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// A dictionary of arbitrary key/value pairs used to describe a specific Content instance.
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {'relative/path': '/artifacts/1/'
	Artifacts map[string]interface{} `json:"artifacts"`
	// SHA256 checksum of the gem
	Checksum *string `json:"checksum,omitempty"`
	// Name of the gem
	Name *string `json:"name,omitempty"`
	// Version of the gem
	Version *string `json:"version,omitempty"`
	// Platform of the gem
	Platform *string `json:"platform,omitempty"`
	// Whether the gem is a prerelease
	Prerelease *bool `json:"prerelease,omitempty"`
	Dependencies *map[string]string `json:"dependencies,omitempty"`
	// Required ruby version of the gem
	RequiredRubyVersion *string `json:"required_ruby_version,omitempty"`
	// Required rubygems version of the gem
	RequiredRubygemsVersion *string `json:"required_rubygems_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GemGemContentResponse GemGemContentResponse

// NewGemGemContentResponse instantiates a new GemGemContentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGemGemContentResponse(artifacts map[string]interface{}) *GemGemContentResponse {
	this := GemGemContentResponse{}
	this.Artifacts = artifacts
	return &this
}

// NewGemGemContentResponseWithDefaults instantiates a new GemGemContentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGemGemContentResponseWithDefaults() *GemGemContentResponse {
	this := GemGemContentResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *GemGemContentResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *GemGemContentResponse) SetPrn(v string) {
	o.Prn = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *GemGemContentResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *GemGemContentResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *GemGemContentResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetArtifacts returns the Artifacts field value
func (o *GemGemContentResponse) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *GemGemContentResponse) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *GemGemContentResponse) SetChecksum(v string) {
	o.Checksum = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GemGemContentResponse) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GemGemContentResponse) SetVersion(v string) {
	o.Version = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *GemGemContentResponse) SetPlatform(v string) {
	o.Platform = &v
}

// GetPrerelease returns the Prerelease field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetPrerelease() bool {
	if o == nil || IsNil(o.Prerelease) {
		var ret bool
		return ret
	}
	return *o.Prerelease
}

// GetPrereleaseOk returns a tuple with the Prerelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetPrereleaseOk() (*bool, bool) {
	if o == nil || IsNil(o.Prerelease) {
		return nil, false
	}
	return o.Prerelease, true
}

// HasPrerelease returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasPrerelease() bool {
	if o != nil && !IsNil(o.Prerelease) {
		return true
	}

	return false
}

// SetPrerelease gets a reference to the given bool and assigns it to the Prerelease field.
func (o *GemGemContentResponse) SetPrerelease(v bool) {
	o.Prerelease = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetDependencies() map[string]string {
	if o == nil || IsNil(o.Dependencies) {
		var ret map[string]string
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetDependenciesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given map[string]string and assigns it to the Dependencies field.
func (o *GemGemContentResponse) SetDependencies(v map[string]string) {
	o.Dependencies = &v
}

// GetRequiredRubyVersion returns the RequiredRubyVersion field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetRequiredRubyVersion() string {
	if o == nil || IsNil(o.RequiredRubyVersion) {
		var ret string
		return ret
	}
	return *o.RequiredRubyVersion
}

// GetRequiredRubyVersionOk returns a tuple with the RequiredRubyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetRequiredRubyVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RequiredRubyVersion) {
		return nil, false
	}
	return o.RequiredRubyVersion, true
}

// HasRequiredRubyVersion returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasRequiredRubyVersion() bool {
	if o != nil && !IsNil(o.RequiredRubyVersion) {
		return true
	}

	return false
}

// SetRequiredRubyVersion gets a reference to the given string and assigns it to the RequiredRubyVersion field.
func (o *GemGemContentResponse) SetRequiredRubyVersion(v string) {
	o.RequiredRubyVersion = &v
}

// GetRequiredRubygemsVersion returns the RequiredRubygemsVersion field value if set, zero value otherwise.
func (o *GemGemContentResponse) GetRequiredRubygemsVersion() string {
	if o == nil || IsNil(o.RequiredRubygemsVersion) {
		var ret string
		return ret
	}
	return *o.RequiredRubygemsVersion
}

// GetRequiredRubygemsVersionOk returns a tuple with the RequiredRubygemsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemContentResponse) GetRequiredRubygemsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RequiredRubygemsVersion) {
		return nil, false
	}
	return o.RequiredRubygemsVersion, true
}

// HasRequiredRubygemsVersion returns a boolean if a field has been set.
func (o *GemGemContentResponse) HasRequiredRubygemsVersion() bool {
	if o != nil && !IsNil(o.RequiredRubygemsVersion) {
		return true
	}

	return false
}

// SetRequiredRubygemsVersion gets a reference to the given string and assigns it to the RequiredRubygemsVersion field.
func (o *GemGemContentResponse) SetRequiredRubygemsVersion(v string) {
	o.RequiredRubygemsVersion = &v
}

func (o GemGemContentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GemGemContentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.Prn) {
		toSerialize["prn"] = o.Prn
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["artifacts"] = o.Artifacts
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Prerelease) {
		toSerialize["prerelease"] = o.Prerelease
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	if !IsNil(o.RequiredRubyVersion) {
		toSerialize["required_ruby_version"] = o.RequiredRubyVersion
	}
	if !IsNil(o.RequiredRubygemsVersion) {
		toSerialize["required_rubygems_version"] = o.RequiredRubygemsVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GemGemContentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"artifacts",
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

	varGemGemContentResponse := _GemGemContentResponse{}

	err = json.Unmarshal(data, &varGemGemContentResponse)

	if err != nil {
		return err
	}

	*o = GemGemContentResponse(varGemGemContentResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "artifacts")
		delete(additionalProperties, "checksum")
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "prerelease")
		delete(additionalProperties, "dependencies")
		delete(additionalProperties, "required_ruby_version")
		delete(additionalProperties, "required_rubygems_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGemGemContentResponse struct {
	value *GemGemContentResponse
	isSet bool
}

func (v NullableGemGemContentResponse) Get() *GemGemContentResponse {
	return v.value
}

func (v *NullableGemGemContentResponse) Set(val *GemGemContentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGemGemContentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGemGemContentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGemGemContentResponse(val *GemGemContentResponse) *NullableGemGemContentResponse {
	return &NullableGemGemContentResponse{value: val, isSet: true}
}

func (v NullableGemGemContentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGemGemContentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


