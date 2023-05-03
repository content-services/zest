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

// checks if the OstreeOstreeRefResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeOstreeRefResponse{}

// OstreeOstreeRefResponse A Serializer class for OSTree head commits.
type OstreeOstreeRefResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
	Commit string `json:"commit"`
	Checksum *string `json:"checksum,omitempty"`
	Name string `json:"name"`
}

// NewOstreeOstreeRefResponse instantiates a new OstreeOstreeRefResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeOstreeRefResponse(artifact string, relativePath string, commit string, name string) *OstreeOstreeRefResponse {
	this := OstreeOstreeRefResponse{}
	this.Artifact = artifact
	this.RelativePath = relativePath
	this.Commit = commit
	this.Name = name
	return &this
}

// NewOstreeOstreeRefResponseWithDefaults instantiates a new OstreeOstreeRefResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeOstreeRefResponseWithDefaults() *OstreeOstreeRefResponse {
	this := OstreeOstreeRefResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *OstreeOstreeRefResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRefResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *OstreeOstreeRefResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *OstreeOstreeRefResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *OstreeOstreeRefResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRefResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *OstreeOstreeRefResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *OstreeOstreeRefResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArtifact returns the Artifact field value
func (o *OstreeOstreeRefResponse) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRefResponse) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *OstreeOstreeRefResponse) SetArtifact(v string) {
	o.Artifact = v
}

// GetRelativePath returns the RelativePath field value
func (o *OstreeOstreeRefResponse) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRefResponse) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *OstreeOstreeRefResponse) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetCommit returns the Commit field value
func (o *OstreeOstreeRefResponse) GetCommit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRefResponse) GetCommitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *OstreeOstreeRefResponse) SetCommit(v string) {
	o.Commit = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *OstreeOstreeRefResponse) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRefResponse) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *OstreeOstreeRefResponse) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *OstreeOstreeRefResponse) SetChecksum(v string) {
	o.Checksum = &v
}

// GetName returns the Name field value
func (o *OstreeOstreeRefResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRefResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OstreeOstreeRefResponse) SetName(v string) {
	o.Name = v
}

func (o OstreeOstreeRefResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeOstreeRefResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["artifact"] = o.Artifact
	toSerialize["relative_path"] = o.RelativePath
	toSerialize["commit"] = o.Commit
	// skip: checksum is readOnly
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableOstreeOstreeRefResponse struct {
	value *OstreeOstreeRefResponse
	isSet bool
}

func (v NullableOstreeOstreeRefResponse) Get() *OstreeOstreeRefResponse {
	return v.value
}

func (v *NullableOstreeOstreeRefResponse) Set(val *OstreeOstreeRefResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeOstreeRefResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeOstreeRefResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeOstreeRefResponse(val *OstreeOstreeRefResponse) *NullableOstreeOstreeRefResponse {
	return &NullableOstreeOstreeRefResponse{value: val, isSet: true}
}

func (v NullableOstreeOstreeRefResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeOstreeRefResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

