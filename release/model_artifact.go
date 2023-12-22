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
	"os"
	"bytes"
	"fmt"
)

// checks if the Artifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Artifact{}

// Artifact Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type Artifact struct {
	// The stored file.
	File *os.File `json:"file"`
	// The size of the file in bytes.
	Size *int64 `json:"size,omitempty"`
	// The MD5 checksum of the file if available.
	Md5 NullableString `json:"md5,omitempty"`
	// The SHA-1 checksum of the file if available.
	Sha1 NullableString `json:"sha1,omitempty"`
	// The SHA-224 checksum of the file if available.
	Sha224 NullableString `json:"sha224,omitempty"`
	// The SHA-256 checksum of the file if available.
	Sha256 NullableString `json:"sha256,omitempty"`
	// The SHA-384 checksum of the file if available.
	Sha384 NullableString `json:"sha384,omitempty"`
	// The SHA-512 checksum of the file if available.
	Sha512 NullableString `json:"sha512,omitempty"`
}

type _Artifact Artifact

// NewArtifact instantiates a new Artifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifact(file *os.File) *Artifact {
	this := Artifact{}
	this.File = file
	return &this
}

// NewArtifactWithDefaults instantiates a new Artifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactWithDefaults() *Artifact {
	this := Artifact{}
	return &this
}

// GetFile returns the File field value
func (o *Artifact) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetFileOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *Artifact) SetFile(v *os.File) {
	o.File = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Artifact) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Artifact) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Artifact) SetSize(v int64) {
	o.Size = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetMd5() string {
	if o == nil || IsNil(o.Md5.Get()) {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *Artifact) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *Artifact) SetMd5(v string) {
	o.Md5.Set(&v)
}
// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *Artifact) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *Artifact) UnsetMd5() {
	o.Md5.Unset()
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetSha1() string {
	if o == nil || IsNil(o.Sha1.Get()) {
		var ret string
		return ret
	}
	return *o.Sha1.Get()
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetSha1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha1.Get(), o.Sha1.IsSet()
}

// HasSha1 returns a boolean if a field has been set.
func (o *Artifact) HasSha1() bool {
	if o != nil && o.Sha1.IsSet() {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given NullableString and assigns it to the Sha1 field.
func (o *Artifact) SetSha1(v string) {
	o.Sha1.Set(&v)
}
// SetSha1Nil sets the value for Sha1 to be an explicit nil
func (o *Artifact) SetSha1Nil() {
	o.Sha1.Set(nil)
}

// UnsetSha1 ensures that no value is present for Sha1, not even an explicit nil
func (o *Artifact) UnsetSha1() {
	o.Sha1.Unset()
}

// GetSha224 returns the Sha224 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetSha224() string {
	if o == nil || IsNil(o.Sha224.Get()) {
		var ret string
		return ret
	}
	return *o.Sha224.Get()
}

// GetSha224Ok returns a tuple with the Sha224 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetSha224Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha224.Get(), o.Sha224.IsSet()
}

// HasSha224 returns a boolean if a field has been set.
func (o *Artifact) HasSha224() bool {
	if o != nil && o.Sha224.IsSet() {
		return true
	}

	return false
}

// SetSha224 gets a reference to the given NullableString and assigns it to the Sha224 field.
func (o *Artifact) SetSha224(v string) {
	o.Sha224.Set(&v)
}
// SetSha224Nil sets the value for Sha224 to be an explicit nil
func (o *Artifact) SetSha224Nil() {
	o.Sha224.Set(nil)
}

// UnsetSha224 ensures that no value is present for Sha224, not even an explicit nil
func (o *Artifact) UnsetSha224() {
	o.Sha224.Unset()
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetSha256() string {
	if o == nil || IsNil(o.Sha256.Get()) {
		var ret string
		return ret
	}
	return *o.Sha256.Get()
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha256.Get(), o.Sha256.IsSet()
}

// HasSha256 returns a boolean if a field has been set.
func (o *Artifact) HasSha256() bool {
	if o != nil && o.Sha256.IsSet() {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given NullableString and assigns it to the Sha256 field.
func (o *Artifact) SetSha256(v string) {
	o.Sha256.Set(&v)
}
// SetSha256Nil sets the value for Sha256 to be an explicit nil
func (o *Artifact) SetSha256Nil() {
	o.Sha256.Set(nil)
}

// UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
func (o *Artifact) UnsetSha256() {
	o.Sha256.Unset()
}

// GetSha384 returns the Sha384 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetSha384() string {
	if o == nil || IsNil(o.Sha384.Get()) {
		var ret string
		return ret
	}
	return *o.Sha384.Get()
}

// GetSha384Ok returns a tuple with the Sha384 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetSha384Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha384.Get(), o.Sha384.IsSet()
}

// HasSha384 returns a boolean if a field has been set.
func (o *Artifact) HasSha384() bool {
	if o != nil && o.Sha384.IsSet() {
		return true
	}

	return false
}

// SetSha384 gets a reference to the given NullableString and assigns it to the Sha384 field.
func (o *Artifact) SetSha384(v string) {
	o.Sha384.Set(&v)
}
// SetSha384Nil sets the value for Sha384 to be an explicit nil
func (o *Artifact) SetSha384Nil() {
	o.Sha384.Set(nil)
}

// UnsetSha384 ensures that no value is present for Sha384, not even an explicit nil
func (o *Artifact) UnsetSha384() {
	o.Sha384.Unset()
}

// GetSha512 returns the Sha512 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetSha512() string {
	if o == nil || IsNil(o.Sha512.Get()) {
		var ret string
		return ret
	}
	return *o.Sha512.Get()
}

// GetSha512Ok returns a tuple with the Sha512 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetSha512Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha512.Get(), o.Sha512.IsSet()
}

// HasSha512 returns a boolean if a field has been set.
func (o *Artifact) HasSha512() bool {
	if o != nil && o.Sha512.IsSet() {
		return true
	}

	return false
}

// SetSha512 gets a reference to the given NullableString and assigns it to the Sha512 field.
func (o *Artifact) SetSha512(v string) {
	o.Sha512.Set(&v)
}
// SetSha512Nil sets the value for Sha512 to be an explicit nil
func (o *Artifact) SetSha512Nil() {
	o.Sha512.Set(nil)
}

// UnsetSha512 ensures that no value is present for Sha512, not even an explicit nil
func (o *Artifact) UnsetSha512() {
	o.Sha512.Unset()
}

func (o Artifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Artifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if o.Sha1.IsSet() {
		toSerialize["sha1"] = o.Sha1.Get()
	}
	if o.Sha224.IsSet() {
		toSerialize["sha224"] = o.Sha224.Get()
	}
	if o.Sha256.IsSet() {
		toSerialize["sha256"] = o.Sha256.Get()
	}
	if o.Sha384.IsSet() {
		toSerialize["sha384"] = o.Sha384.Get()
	}
	if o.Sha512.IsSet() {
		toSerialize["sha512"] = o.Sha512.Get()
	}
	return toSerialize, nil
}

func (o *Artifact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file",
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

	varArtifact := _Artifact{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArtifact)

	if err != nil {
		return err
	}

	*o = Artifact(varArtifact)

	return err
}

type NullableArtifact struct {
	value *Artifact
	isSet bool
}

func (v NullableArtifact) Get() *Artifact {
	return v.value
}

func (v *NullableArtifact) Set(val *Artifact) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifact(val *Artifact) *NullableArtifact {
	return &NullableArtifact{value: val, isSet: true}
}

func (v NullableArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


