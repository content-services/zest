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

// checks if the PythonPythonPublication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PythonPythonPublication{}

// PythonPythonPublication A Serializer for PythonPublication.
type PythonPythonPublication struct {
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// A URI of the repository to be published.
	Repository *string `json:"repository,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PythonPythonPublication PythonPythonPublication

// NewPythonPythonPublication instantiates a new PythonPythonPublication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPythonPythonPublication() *PythonPythonPublication {
	this := PythonPythonPublication{}
	return &this
}

// NewPythonPythonPublicationWithDefaults instantiates a new PythonPythonPublication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPythonPythonPublicationWithDefaults() *PythonPythonPublication {
	this := PythonPythonPublication{}
	return &this
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *PythonPythonPublication) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublication) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryVersion) {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *PythonPythonPublication) HasRepositoryVersion() bool {
	if o != nil && !IsNil(o.RepositoryVersion) {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *PythonPythonPublication) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *PythonPythonPublication) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublication) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *PythonPythonPublication) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *PythonPythonPublication) SetRepository(v string) {
	o.Repository = &v
}

func (o PythonPythonPublication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PythonPythonPublication) ToMap() (map[string]interface{}, error) {
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

func (o *PythonPythonPublication) UnmarshalJSON(data []byte) (err error) {
	varPythonPythonPublication := _PythonPythonPublication{}

	err = json.Unmarshal(data, &varPythonPythonPublication)

	if err != nil {
		return err
	}

	*o = PythonPythonPublication(varPythonPythonPublication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repository_version")
		delete(additionalProperties, "repository")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePythonPythonPublication struct {
	value *PythonPythonPublication
	isSet bool
}

func (v NullablePythonPythonPublication) Get() *PythonPythonPublication {
	return v.value
}

func (v *NullablePythonPythonPublication) Set(val *PythonPythonPublication) {
	v.value = val
	v.isSet = true
}

func (v NullablePythonPythonPublication) IsSet() bool {
	return v.isSet
}

func (v *NullablePythonPythonPublication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePythonPythonPublication(val *PythonPythonPublication) *NullablePythonPythonPublication {
	return &NullablePythonPythonPublication{value: val, isSet: true}
}

func (v NullablePythonPythonPublication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePythonPythonPublication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


