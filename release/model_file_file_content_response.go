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

// checks if the FileFileContentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileFileContentResponse{}

// FileFileContentResponse Serializer for File Content.
type FileFileContentResponse struct {
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
	Artifact *string `json:"artifact,omitempty"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
	// The MD5 checksum if available.
	Md5 *string `json:"md5,omitempty"`
	// The SHA-1 checksum if available.
	Sha1 *string `json:"sha1,omitempty"`
	// The SHA-224 checksum if available.
	Sha224 *string `json:"sha224,omitempty"`
	// The SHA-256 checksum if available.
	Sha256 *string `json:"sha256,omitempty"`
	// The SHA-384 checksum if available.
	Sha384 *string `json:"sha384,omitempty"`
	// The SHA-512 checksum if available.
	Sha512 *string `json:"sha512,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FileFileContentResponse FileFileContentResponse

// NewFileFileContentResponse instantiates a new FileFileContentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileFileContentResponse(relativePath string) *FileFileContentResponse {
	this := FileFileContentResponse{}
	this.RelativePath = relativePath
	return &this
}

// NewFileFileContentResponseWithDefaults instantiates a new FileFileContentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileFileContentResponseWithDefaults() *FileFileContentResponse {
	this := FileFileContentResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *FileFileContentResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *FileFileContentResponse) SetPrn(v string) {
	o.Prn = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *FileFileContentResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *FileFileContentResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *FileFileContentResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *FileFileContentResponse) SetArtifact(v string) {
	o.Artifact = &v
}

// GetRelativePath returns the RelativePath field value
func (o *FileFileContentResponse) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *FileFileContentResponse) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *FileFileContentResponse) SetMd5(v string) {
	o.Md5 = &v
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetSha1() string {
	if o == nil || IsNil(o.Sha1) {
		var ret string
		return ret
	}
	return *o.Sha1
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetSha1Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha1) {
		return nil, false
	}
	return o.Sha1, true
}

// HasSha1 returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasSha1() bool {
	if o != nil && !IsNil(o.Sha1) {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given string and assigns it to the Sha1 field.
func (o *FileFileContentResponse) SetSha1(v string) {
	o.Sha1 = &v
}

// GetSha224 returns the Sha224 field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetSha224() string {
	if o == nil || IsNil(o.Sha224) {
		var ret string
		return ret
	}
	return *o.Sha224
}

// GetSha224Ok returns a tuple with the Sha224 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetSha224Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha224) {
		return nil, false
	}
	return o.Sha224, true
}

// HasSha224 returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasSha224() bool {
	if o != nil && !IsNil(o.Sha224) {
		return true
	}

	return false
}

// SetSha224 gets a reference to the given string and assigns it to the Sha224 field.
func (o *FileFileContentResponse) SetSha224(v string) {
	o.Sha224 = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetSha256() string {
	if o == nil || IsNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha256) {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasSha256() bool {
	if o != nil && !IsNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *FileFileContentResponse) SetSha256(v string) {
	o.Sha256 = &v
}

// GetSha384 returns the Sha384 field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetSha384() string {
	if o == nil || IsNil(o.Sha384) {
		var ret string
		return ret
	}
	return *o.Sha384
}

// GetSha384Ok returns a tuple with the Sha384 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetSha384Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha384) {
		return nil, false
	}
	return o.Sha384, true
}

// HasSha384 returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasSha384() bool {
	if o != nil && !IsNil(o.Sha384) {
		return true
	}

	return false
}

// SetSha384 gets a reference to the given string and assigns it to the Sha384 field.
func (o *FileFileContentResponse) SetSha384(v string) {
	o.Sha384 = &v
}

// GetSha512 returns the Sha512 field value if set, zero value otherwise.
func (o *FileFileContentResponse) GetSha512() string {
	if o == nil || IsNil(o.Sha512) {
		var ret string
		return ret
	}
	return *o.Sha512
}

// GetSha512Ok returns a tuple with the Sha512 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContentResponse) GetSha512Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha512) {
		return nil, false
	}
	return o.Sha512, true
}

// HasSha512 returns a boolean if a field has been set.
func (o *FileFileContentResponse) HasSha512() bool {
	if o != nil && !IsNil(o.Sha512) {
		return true
	}

	return false
}

// SetSha512 gets a reference to the given string and assigns it to the Sha512 field.
func (o *FileFileContentResponse) SetSha512(v string) {
	o.Sha512 = &v
}

func (o FileFileContentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileFileContentResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	toSerialize["relative_path"] = o.RelativePath
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.Sha1) {
		toSerialize["sha1"] = o.Sha1
	}
	if !IsNil(o.Sha224) {
		toSerialize["sha224"] = o.Sha224
	}
	if !IsNil(o.Sha256) {
		toSerialize["sha256"] = o.Sha256
	}
	if !IsNil(o.Sha384) {
		toSerialize["sha384"] = o.Sha384
	}
	if !IsNil(o.Sha512) {
		toSerialize["sha512"] = o.Sha512
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileFileContentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"relative_path",
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

	varFileFileContentResponse := _FileFileContentResponse{}

	err = json.Unmarshal(data, &varFileFileContentResponse)

	if err != nil {
		return err
	}

	*o = FileFileContentResponse(varFileFileContentResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "artifact")
		delete(additionalProperties, "relative_path")
		delete(additionalProperties, "md5")
		delete(additionalProperties, "sha1")
		delete(additionalProperties, "sha224")
		delete(additionalProperties, "sha256")
		delete(additionalProperties, "sha384")
		delete(additionalProperties, "sha512")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileFileContentResponse struct {
	value *FileFileContentResponse
	isSet bool
}

func (v NullableFileFileContentResponse) Get() *FileFileContentResponse {
	return v.value
}

func (v *NullableFileFileContentResponse) Set(val *FileFileContentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileFileContentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileFileContentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileFileContentResponse(val *FileFileContentResponse) *NullableFileFileContentResponse {
	return &NullableFileFileContentResponse{value: val, isSet: true}
}

func (v NullableFileFileContentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileFileContentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


