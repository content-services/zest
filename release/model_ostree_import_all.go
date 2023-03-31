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
)

// checks if the OstreeImportAll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeImportAll{}

// OstreeImportAll A Serializer class for importing all refs and commits to a repository.
type OstreeImportAll struct {
	// An artifact representing OSTree content compressed as a tarball.
	Artifact string `json:"artifact"`
	// The name of a repository that contains the compressed OSTree content.
	RepositoryName string `json:"repository_name"`
}

// NewOstreeImportAll instantiates a new OstreeImportAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeImportAll(artifact string, repositoryName string) *OstreeImportAll {
	this := OstreeImportAll{}
	this.Artifact = artifact
	this.RepositoryName = repositoryName
	return &this
}

// NewOstreeImportAllWithDefaults instantiates a new OstreeImportAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeImportAllWithDefaults() *OstreeImportAll {
	this := OstreeImportAll{}
	return &this
}

// GetArtifact returns the Artifact field value
func (o *OstreeImportAll) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *OstreeImportAll) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *OstreeImportAll) SetArtifact(v string) {
	o.Artifact = v
}

// GetRepositoryName returns the RepositoryName field value
func (o *OstreeImportAll) GetRepositoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value
// and a boolean to check if the value has been set.
func (o *OstreeImportAll) GetRepositoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryName, true
}

// SetRepositoryName sets field value
func (o *OstreeImportAll) SetRepositoryName(v string) {
	o.RepositoryName = v
}

func (o OstreeImportAll) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeImportAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifact"] = o.Artifact
	toSerialize["repository_name"] = o.RepositoryName
	return toSerialize, nil
}

type NullableOstreeImportAll struct {
	value *OstreeImportAll
	isSet bool
}

func (v NullableOstreeImportAll) Get() *OstreeImportAll {
	return v.value
}

func (v *NullableOstreeImportAll) Set(val *OstreeImportAll) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeImportAll) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeImportAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeImportAll(val *OstreeImportAll) *NullableOstreeImportAll {
	return &NullableOstreeImportAll{value: val, isSet: true}
}

func (v NullableOstreeImportAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeImportAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


