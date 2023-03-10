/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
)

// checks if the DebReleaseFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebReleaseFile{}

// DebReleaseFile A serializer for ReleaseFile.
type DebReleaseFile struct {
	// A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {'relative/path': '/artifacts/1/'
	Artifacts map[string]interface{} `json:"artifacts"`
	// Codename of the release, i.e. \"buster\".
	Codename *string `json:"codename,omitempty"`
	// Suite of the release, i.e. \"stable\".
	Suite *string `json:"suite,omitempty"`
	// Distribution of the release, i.e. \"stable/updates\".
	Distribution string `json:"distribution"`
	// Path of file relative to url.
	RelativePath *string `json:"relative_path,omitempty"`
}

// NewDebReleaseFile instantiates a new DebReleaseFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebReleaseFile(artifacts map[string]interface{}, distribution string) *DebReleaseFile {
	this := DebReleaseFile{}
	this.Artifacts = artifacts
	this.Distribution = distribution
	return &this
}

// NewDebReleaseFileWithDefaults instantiates a new DebReleaseFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebReleaseFileWithDefaults() *DebReleaseFile {
	this := DebReleaseFile{}
	return &this
}

// GetArtifacts returns the Artifacts field value
func (o *DebReleaseFile) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *DebReleaseFile) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *DebReleaseFile) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

// GetCodename returns the Codename field value if set, zero value otherwise.
func (o *DebReleaseFile) GetCodename() string {
	if o == nil || IsNil(o.Codename) {
		var ret string
		return ret
	}
	return *o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebReleaseFile) GetCodenameOk() (*string, bool) {
	if o == nil || IsNil(o.Codename) {
		return nil, false
	}
	return o.Codename, true
}

// HasCodename returns a boolean if a field has been set.
func (o *DebReleaseFile) HasCodename() bool {
	if o != nil && !IsNil(o.Codename) {
		return true
	}

	return false
}

// SetCodename gets a reference to the given string and assigns it to the Codename field.
func (o *DebReleaseFile) SetCodename(v string) {
	o.Codename = &v
}

// GetSuite returns the Suite field value if set, zero value otherwise.
func (o *DebReleaseFile) GetSuite() string {
	if o == nil || IsNil(o.Suite) {
		var ret string
		return ret
	}
	return *o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebReleaseFile) GetSuiteOk() (*string, bool) {
	if o == nil || IsNil(o.Suite) {
		return nil, false
	}
	return o.Suite, true
}

// HasSuite returns a boolean if a field has been set.
func (o *DebReleaseFile) HasSuite() bool {
	if o != nil && !IsNil(o.Suite) {
		return true
	}

	return false
}

// SetSuite gets a reference to the given string and assigns it to the Suite field.
func (o *DebReleaseFile) SetSuite(v string) {
	o.Suite = &v
}

// GetDistribution returns the Distribution field value
func (o *DebReleaseFile) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *DebReleaseFile) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *DebReleaseFile) SetDistribution(v string) {
	o.Distribution = v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *DebReleaseFile) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebReleaseFile) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *DebReleaseFile) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *DebReleaseFile) SetRelativePath(v string) {
	o.RelativePath = &v
}

func (o DebReleaseFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebReleaseFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifacts"] = o.Artifacts
	if !IsNil(o.Codename) {
		toSerialize["codename"] = o.Codename
	}
	if !IsNil(o.Suite) {
		toSerialize["suite"] = o.Suite
	}
	toSerialize["distribution"] = o.Distribution
	if !IsNil(o.RelativePath) {
		toSerialize["relative_path"] = o.RelativePath
	}
	return toSerialize, nil
}

type NullableDebReleaseFile struct {
	value *DebReleaseFile
	isSet bool
}

func (v NullableDebReleaseFile) Get() *DebReleaseFile {
	return v.value
}

func (v *NullableDebReleaseFile) Set(val *DebReleaseFile) {
	v.value = val
	v.isSet = true
}

func (v NullableDebReleaseFile) IsSet() bool {
	return v.isSet
}

func (v *NullableDebReleaseFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebReleaseFile(val *DebReleaseFile) *NullableDebReleaseFile {
	return &NullableDebReleaseFile{value: val, isSet: true}
}

func (v NullableDebReleaseFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebReleaseFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


