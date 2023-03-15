/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the RepositoryVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryVersionResponse{}

// RepositoryVersionResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type RepositoryVersionResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	Number *int64 `json:"number,omitempty"`
	Repository *string `json:"repository,omitempty"`
	// A repository version whose content was used as the initial set of content for this repository version
	BaseVersion *string `json:"base_version,omitempty"`
	ContentSummary *RepositoryVersionResponseContentSummary `json:"content_summary,omitempty"`
}

// NewRepositoryVersionResponse instantiates a new RepositoryVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryVersionResponse() *RepositoryVersionResponse {
	this := RepositoryVersionResponse{}
	return &this
}

// NewRepositoryVersionResponseWithDefaults instantiates a new RepositoryVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryVersionResponseWithDefaults() *RepositoryVersionResponse {
	this := RepositoryVersionResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RepositoryVersionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RepositoryVersionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetNumber() int64 {
	if o == nil || IsNil(o.Number) {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *RepositoryVersionResponse) SetNumber(v int64) {
	o.Number = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *RepositoryVersionResponse) SetRepository(v string) {
	o.Repository = &v
}

// GetBaseVersion returns the BaseVersion field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetBaseVersion() string {
	if o == nil || IsNil(o.BaseVersion) {
		var ret string
		return ret
	}
	return *o.BaseVersion
}

// GetBaseVersionOk returns a tuple with the BaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetBaseVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BaseVersion) {
		return nil, false
	}
	return o.BaseVersion, true
}

// HasBaseVersion returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasBaseVersion() bool {
	if o != nil && !IsNil(o.BaseVersion) {
		return true
	}

	return false
}

// SetBaseVersion gets a reference to the given string and assigns it to the BaseVersion field.
func (o *RepositoryVersionResponse) SetBaseVersion(v string) {
	o.BaseVersion = &v
}

// GetContentSummary returns the ContentSummary field value if set, zero value otherwise.
func (o *RepositoryVersionResponse) GetContentSummary() RepositoryVersionResponseContentSummary {
	if o == nil || IsNil(o.ContentSummary) {
		var ret RepositoryVersionResponseContentSummary
		return ret
	}
	return *o.ContentSummary
}

// GetContentSummaryOk returns a tuple with the ContentSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryVersionResponse) GetContentSummaryOk() (*RepositoryVersionResponseContentSummary, bool) {
	if o == nil || IsNil(o.ContentSummary) {
		return nil, false
	}
	return o.ContentSummary, true
}

// HasContentSummary returns a boolean if a field has been set.
func (o *RepositoryVersionResponse) HasContentSummary() bool {
	if o != nil && !IsNil(o.ContentSummary) {
		return true
	}

	return false
}

// SetContentSummary gets a reference to the given RepositoryVersionResponseContentSummary and assigns it to the ContentSummary field.
func (o *RepositoryVersionResponse) SetContentSummary(v RepositoryVersionResponseContentSummary) {
	o.ContentSummary = &v
}

func (o RepositoryVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	// skip: number is readOnly
	// skip: repository is readOnly
	if !IsNil(o.BaseVersion) {
		toSerialize["base_version"] = o.BaseVersion
	}
	if !IsNil(o.ContentSummary) {
		toSerialize["content_summary"] = o.ContentSummary
	}
	return toSerialize, nil
}

type NullableRepositoryVersionResponse struct {
	value *RepositoryVersionResponse
	isSet bool
}

func (v NullableRepositoryVersionResponse) Get() *RepositoryVersionResponse {
	return v.value
}

func (v *NullableRepositoryVersionResponse) Set(val *RepositoryVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryVersionResponse(val *RepositoryVersionResponse) *NullableRepositoryVersionResponse {
	return &NullableRepositoryVersionResponse{value: val, isSet: true}
}

func (v NullableRepositoryVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


