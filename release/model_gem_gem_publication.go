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

// checks if the GemGemPublication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GemGemPublication{}

// GemGemPublication A Serializer for GemPublication.
type GemGemPublication struct {
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// A URI of the repository to be published.
	Repository *string `json:"repository,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GemGemPublication GemGemPublication

// NewGemGemPublication instantiates a new GemGemPublication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGemGemPublication() *GemGemPublication {
	this := GemGemPublication{}
	return &this
}

// NewGemGemPublicationWithDefaults instantiates a new GemGemPublication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGemGemPublicationWithDefaults() *GemGemPublication {
	this := GemGemPublication{}
	return &this
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *GemGemPublication) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemPublication) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryVersion) {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *GemGemPublication) HasRepositoryVersion() bool {
	if o != nil && !IsNil(o.RepositoryVersion) {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *GemGemPublication) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GemGemPublication) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemPublication) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GemGemPublication) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *GemGemPublication) SetRepository(v string) {
	o.Repository = &v
}

func (o GemGemPublication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GemGemPublication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepositoryVersion) {
		toSerialize["repository_version"] = o.RepositoryVersion
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GemGemPublication) UnmarshalJSON(data []byte) (err error) {
	varGemGemPublication := _GemGemPublication{}

	err = json.Unmarshal(data, &varGemGemPublication)

	if err != nil {
		return err
	}

	*o = GemGemPublication(varGemGemPublication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repository_version")
		delete(additionalProperties, "repository")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGemGemPublication struct {
	value *GemGemPublication
	isSet bool
}

func (v NullableGemGemPublication) Get() *GemGemPublication {
	return v.value
}

func (v *NullableGemGemPublication) Set(val *GemGemPublication) {
	v.value = val
	v.isSet = true
}

func (v NullableGemGemPublication) IsSet() bool {
	return v.isSet
}

func (v *NullableGemGemPublication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGemGemPublication(val *GemGemPublication) *NullableGemGemPublication {
	return &NullableGemGemPublication{value: val, isSet: true}
}

func (v NullableGemGemPublication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGemGemPublication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


