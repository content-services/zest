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

// checks if the PulpImportCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PulpImportCheckResponse{}

// PulpImportCheckResponse Return the response to a PulpImport import-check call.
type PulpImportCheckResponse struct {
	Toc *PulpImportCheckResponseToc `json:"toc,omitempty"`
	Path *PulpImportCheckResponsePath `json:"path,omitempty"`
	RepoMapping *PulpImportCheckResponseRepoMapping `json:"repo_mapping,omitempty"`
}

// NewPulpImportCheckResponse instantiates a new PulpImportCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPulpImportCheckResponse() *PulpImportCheckResponse {
	this := PulpImportCheckResponse{}
	return &this
}

// NewPulpImportCheckResponseWithDefaults instantiates a new PulpImportCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPulpImportCheckResponseWithDefaults() *PulpImportCheckResponse {
	this := PulpImportCheckResponse{}
	return &this
}

// GetToc returns the Toc field value if set, zero value otherwise.
func (o *PulpImportCheckResponse) GetToc() PulpImportCheckResponseToc {
	if o == nil || IsNil(o.Toc) {
		var ret PulpImportCheckResponseToc
		return ret
	}
	return *o.Toc
}

// GetTocOk returns a tuple with the Toc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponse) GetTocOk() (*PulpImportCheckResponseToc, bool) {
	if o == nil || IsNil(o.Toc) {
		return nil, false
	}
	return o.Toc, true
}

// HasToc returns a boolean if a field has been set.
func (o *PulpImportCheckResponse) HasToc() bool {
	if o != nil && !IsNil(o.Toc) {
		return true
	}

	return false
}

// SetToc gets a reference to the given PulpImportCheckResponseToc and assigns it to the Toc field.
func (o *PulpImportCheckResponse) SetToc(v PulpImportCheckResponseToc) {
	o.Toc = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PulpImportCheckResponse) GetPath() PulpImportCheckResponsePath {
	if o == nil || IsNil(o.Path) {
		var ret PulpImportCheckResponsePath
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponse) GetPathOk() (*PulpImportCheckResponsePath, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PulpImportCheckResponse) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given PulpImportCheckResponsePath and assigns it to the Path field.
func (o *PulpImportCheckResponse) SetPath(v PulpImportCheckResponsePath) {
	o.Path = &v
}

// GetRepoMapping returns the RepoMapping field value if set, zero value otherwise.
func (o *PulpImportCheckResponse) GetRepoMapping() PulpImportCheckResponseRepoMapping {
	if o == nil || IsNil(o.RepoMapping) {
		var ret PulpImportCheckResponseRepoMapping
		return ret
	}
	return *o.RepoMapping
}

// GetRepoMappingOk returns a tuple with the RepoMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponse) GetRepoMappingOk() (*PulpImportCheckResponseRepoMapping, bool) {
	if o == nil || IsNil(o.RepoMapping) {
		return nil, false
	}
	return o.RepoMapping, true
}

// HasRepoMapping returns a boolean if a field has been set.
func (o *PulpImportCheckResponse) HasRepoMapping() bool {
	if o != nil && !IsNil(o.RepoMapping) {
		return true
	}

	return false
}

// SetRepoMapping gets a reference to the given PulpImportCheckResponseRepoMapping and assigns it to the RepoMapping field.
func (o *PulpImportCheckResponse) SetRepoMapping(v PulpImportCheckResponseRepoMapping) {
	o.RepoMapping = &v
}

func (o PulpImportCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PulpImportCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Toc) {
		toSerialize["toc"] = o.Toc
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.RepoMapping) {
		toSerialize["repo_mapping"] = o.RepoMapping
	}
	return toSerialize, nil
}

type NullablePulpImportCheckResponse struct {
	value *PulpImportCheckResponse
	isSet bool
}

func (v NullablePulpImportCheckResponse) Get() *PulpImportCheckResponse {
	return v.value
}

func (v *NullablePulpImportCheckResponse) Set(val *PulpImportCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePulpImportCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePulpImportCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePulpImportCheckResponse(val *PulpImportCheckResponse) *NullablePulpImportCheckResponse {
	return &NullablePulpImportCheckResponse{value: val, isSet: true}
}

func (v NullablePulpImportCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePulpImportCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


