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

// checks if the PythonPythonPublicationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PythonPythonPublicationResponse{}

// PythonPythonPublicationResponse A Serializer for PythonPublication.
type PythonPythonPublicationResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// A URI of the repository to be published.
	Repository *string `json:"repository,omitempty"`
	// This publication is currently being hosted as configured by these distributions.
	Distributions []string `json:"distributions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PythonPythonPublicationResponse PythonPythonPublicationResponse

// NewPythonPythonPublicationResponse instantiates a new PythonPythonPublicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPythonPythonPublicationResponse() *PythonPythonPublicationResponse {
	this := PythonPythonPublicationResponse{}
	return &this
}

// NewPythonPythonPublicationResponseWithDefaults instantiates a new PythonPythonPublicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPythonPythonPublicationResponseWithDefaults() *PythonPythonPublicationResponse {
	this := PythonPythonPublicationResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *PythonPythonPublicationResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublicationResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *PythonPythonPublicationResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *PythonPythonPublicationResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *PythonPythonPublicationResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublicationResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *PythonPythonPublicationResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *PythonPythonPublicationResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *PythonPythonPublicationResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublicationResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *PythonPythonPublicationResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *PythonPythonPublicationResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *PythonPythonPublicationResponse) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublicationResponse) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryVersion) {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *PythonPythonPublicationResponse) HasRepositoryVersion() bool {
	if o != nil && !IsNil(o.RepositoryVersion) {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *PythonPythonPublicationResponse) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *PythonPythonPublicationResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublicationResponse) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *PythonPythonPublicationResponse) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *PythonPythonPublicationResponse) SetRepository(v string) {
	o.Repository = &v
}

// GetDistributions returns the Distributions field value if set, zero value otherwise.
func (o *PythonPythonPublicationResponse) GetDistributions() []string {
	if o == nil || IsNil(o.Distributions) {
		var ret []string
		return ret
	}
	return o.Distributions
}

// GetDistributionsOk returns a tuple with the Distributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonPythonPublicationResponse) GetDistributionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Distributions) {
		return nil, false
	}
	return o.Distributions, true
}

// HasDistributions returns a boolean if a field has been set.
func (o *PythonPythonPublicationResponse) HasDistributions() bool {
	if o != nil && !IsNil(o.Distributions) {
		return true
	}

	return false
}

// SetDistributions gets a reference to the given []string and assigns it to the Distributions field.
func (o *PythonPythonPublicationResponse) SetDistributions(v []string) {
	o.Distributions = v
}

func (o PythonPythonPublicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PythonPythonPublicationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	if !IsNil(o.RepositoryVersion) {
		toSerialize["repository_version"] = o.RepositoryVersion
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Distributions) {
		toSerialize["distributions"] = o.Distributions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PythonPythonPublicationResponse) UnmarshalJSON(data []byte) (err error) {
	varPythonPythonPublicationResponse := _PythonPythonPublicationResponse{}

	err = json.Unmarshal(data, &varPythonPythonPublicationResponse)

	if err != nil {
		return err
	}

	*o = PythonPythonPublicationResponse(varPythonPythonPublicationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "repository_version")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "distributions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePythonPythonPublicationResponse struct {
	value *PythonPythonPublicationResponse
	isSet bool
}

func (v NullablePythonPythonPublicationResponse) Get() *PythonPythonPublicationResponse {
	return v.value
}

func (v *NullablePythonPythonPublicationResponse) Set(val *PythonPythonPublicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePythonPythonPublicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePythonPythonPublicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePythonPythonPublicationResponse(val *PythonPythonPublicationResponse) *NullablePythonPythonPublicationResponse {
	return &NullablePythonPythonPublicationResponse{value: val, isSet: true}
}

func (v NullablePythonPythonPublicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePythonPythonPublicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


