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

// checks if the OstreeOstreeCommitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeOstreeCommitResponse{}

// OstreeOstreeCommitResponse A Serializer class for OSTree commits.
type OstreeOstreeCommitResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// The Pulp Resource Name (PRN).
	Prn *string `json:"prn,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// A dictionary of arbitrary key/value pairs used to describe a specific Content instance.
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
	ParentCommit NullableString `json:"parent_commit,omitempty"`
	Checksum string `json:"checksum"`
	Objs []string `json:"objs"`
	AdditionalProperties map[string]interface{}
}

type _OstreeOstreeCommitResponse OstreeOstreeCommitResponse

// NewOstreeOstreeCommitResponse instantiates a new OstreeOstreeCommitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeOstreeCommitResponse(artifact string, relativePath string, checksum string, objs []string) *OstreeOstreeCommitResponse {
	this := OstreeOstreeCommitResponse{}
	this.Artifact = artifact
	this.RelativePath = relativePath
	this.Checksum = checksum
	this.Objs = objs
	return &this
}

// NewOstreeOstreeCommitResponseWithDefaults instantiates a new OstreeOstreeCommitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeOstreeCommitResponseWithDefaults() *OstreeOstreeCommitResponse {
	this := OstreeOstreeCommitResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *OstreeOstreeCommitResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *OstreeOstreeCommitResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *OstreeOstreeCommitResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *OstreeOstreeCommitResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *OstreeOstreeCommitResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *OstreeOstreeCommitResponse) SetPrn(v string) {
	o.Prn = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *OstreeOstreeCommitResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *OstreeOstreeCommitResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *OstreeOstreeCommitResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *OstreeOstreeCommitResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *OstreeOstreeCommitResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *OstreeOstreeCommitResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *OstreeOstreeCommitResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *OstreeOstreeCommitResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *OstreeOstreeCommitResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetArtifact returns the Artifact field value
func (o *OstreeOstreeCommitResponse) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *OstreeOstreeCommitResponse) SetArtifact(v string) {
	o.Artifact = v
}

// GetRelativePath returns the RelativePath field value
func (o *OstreeOstreeCommitResponse) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *OstreeOstreeCommitResponse) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetParentCommit returns the ParentCommit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OstreeOstreeCommitResponse) GetParentCommit() string {
	if o == nil || IsNil(o.ParentCommit.Get()) {
		var ret string
		return ret
	}
	return *o.ParentCommit.Get()
}

// GetParentCommitOk returns a tuple with the ParentCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OstreeOstreeCommitResponse) GetParentCommitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentCommit.Get(), o.ParentCommit.IsSet()
}

// HasParentCommit returns a boolean if a field has been set.
func (o *OstreeOstreeCommitResponse) HasParentCommit() bool {
	if o != nil && o.ParentCommit.IsSet() {
		return true
	}

	return false
}

// SetParentCommit gets a reference to the given NullableString and assigns it to the ParentCommit field.
func (o *OstreeOstreeCommitResponse) SetParentCommit(v string) {
	o.ParentCommit.Set(&v)
}
// SetParentCommitNil sets the value for ParentCommit to be an explicit nil
func (o *OstreeOstreeCommitResponse) SetParentCommitNil() {
	o.ParentCommit.Set(nil)
}

// UnsetParentCommit ensures that no value is present for ParentCommit, not even an explicit nil
func (o *OstreeOstreeCommitResponse) UnsetParentCommit() {
	o.ParentCommit.Unset()
}

// GetChecksum returns the Checksum field value
func (o *OstreeOstreeCommitResponse) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value
func (o *OstreeOstreeCommitResponse) SetChecksum(v string) {
	o.Checksum = v
}

// GetObjs returns the Objs field value
func (o *OstreeOstreeCommitResponse) GetObjs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Objs
}

// GetObjsOk returns a tuple with the Objs field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeCommitResponse) GetObjsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objs, true
}

// SetObjs sets field value
func (o *OstreeOstreeCommitResponse) SetObjs(v []string) {
	o.Objs = v
}

func (o OstreeOstreeCommitResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeOstreeCommitResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["artifact"] = o.Artifact
	toSerialize["relative_path"] = o.RelativePath
	if o.ParentCommit.IsSet() {
		toSerialize["parent_commit"] = o.ParentCommit.Get()
	}
	toSerialize["checksum"] = o.Checksum
	toSerialize["objs"] = o.Objs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OstreeOstreeCommitResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"artifact",
		"relative_path",
		"checksum",
		"objs",
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

	varOstreeOstreeCommitResponse := _OstreeOstreeCommitResponse{}

	err = json.Unmarshal(data, &varOstreeOstreeCommitResponse)

	if err != nil {
		return err
	}

	*o = OstreeOstreeCommitResponse(varOstreeOstreeCommitResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "artifact")
		delete(additionalProperties, "relative_path")
		delete(additionalProperties, "parent_commit")
		delete(additionalProperties, "checksum")
		delete(additionalProperties, "objs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOstreeOstreeCommitResponse struct {
	value *OstreeOstreeCommitResponse
	isSet bool
}

func (v NullableOstreeOstreeCommitResponse) Get() *OstreeOstreeCommitResponse {
	return v.value
}

func (v *NullableOstreeOstreeCommitResponse) Set(val *OstreeOstreeCommitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeOstreeCommitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeOstreeCommitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeOstreeCommitResponse(val *OstreeOstreeCommitResponse) *NullableOstreeOstreeCommitResponse {
	return &NullableOstreeOstreeCommitResponse{value: val, isSet: true}
}

func (v NullableOstreeOstreeCommitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeOstreeCommitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


