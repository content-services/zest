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

// checks if the MultipleArtifactContentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipleArtifactContentResponse{}

// MultipleArtifactContentResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type MultipleArtifactContentResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {'relative/path': '/artifacts/1/'
	Artifacts map[string]interface{} `json:"artifacts"`
}

// NewMultipleArtifactContentResponse instantiates a new MultipleArtifactContentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleArtifactContentResponse(artifacts map[string]interface{}) *MultipleArtifactContentResponse {
	this := MultipleArtifactContentResponse{}
	this.Artifacts = artifacts
	return &this
}

// NewMultipleArtifactContentResponseWithDefaults instantiates a new MultipleArtifactContentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleArtifactContentResponseWithDefaults() *MultipleArtifactContentResponse {
	this := MultipleArtifactContentResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *MultipleArtifactContentResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleArtifactContentResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *MultipleArtifactContentResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *MultipleArtifactContentResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *MultipleArtifactContentResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleArtifactContentResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *MultipleArtifactContentResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *MultipleArtifactContentResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArtifacts returns the Artifacts field value
func (o *MultipleArtifactContentResponse) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *MultipleArtifactContentResponse) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *MultipleArtifactContentResponse) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

func (o MultipleArtifactContentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipleArtifactContentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["artifacts"] = o.Artifacts
	return toSerialize, nil
}

type NullableMultipleArtifactContentResponse struct {
	value *MultipleArtifactContentResponse
	isSet bool
}

func (v NullableMultipleArtifactContentResponse) Get() *MultipleArtifactContentResponse {
	return v.value
}

func (v *NullableMultipleArtifactContentResponse) Set(val *MultipleArtifactContentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleArtifactContentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleArtifactContentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleArtifactContentResponse(val *MultipleArtifactContentResponse) *NullableMultipleArtifactContentResponse {
	return &NullableMultipleArtifactContentResponse{value: val, isSet: true}
}

func (v NullableMultipleArtifactContentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleArtifactContentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


