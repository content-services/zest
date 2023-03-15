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
)

// checks if the ContainerRepositorySyncURL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRepositorySyncURL{}

// ContainerRepositorySyncURL Serializer for Container Sync.
type ContainerRepositorySyncURL struct {
	// A remote to sync from. This will override a remote set on repository.
	Remote *string `json:"remote,omitempty"`
	// If ``True``, synchronization will remove all content that is not present in the remote repository. If ``False``, sync will be additive only.
	Mirror *bool `json:"mirror,omitempty"`
	// If ``True``, only signed content will be synced. Signatures are not verified.
	SignedOnly *bool `json:"signed_only,omitempty"`
}

// NewContainerRepositorySyncURL instantiates a new ContainerRepositorySyncURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRepositorySyncURL() *ContainerRepositorySyncURL {
	this := ContainerRepositorySyncURL{}
	var mirror bool = false
	this.Mirror = &mirror
	var signedOnly bool = false
	this.SignedOnly = &signedOnly
	return &this
}

// NewContainerRepositorySyncURLWithDefaults instantiates a new ContainerRepositorySyncURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRepositorySyncURLWithDefaults() *ContainerRepositorySyncURL {
	this := ContainerRepositorySyncURL{}
	var mirror bool = false
	this.Mirror = &mirror
	var signedOnly bool = false
	this.SignedOnly = &signedOnly
	return &this
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *ContainerRepositorySyncURL) GetRemote() string {
	if o == nil || IsNil(o.Remote) {
		var ret string
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRepositorySyncURL) GetRemoteOk() (*string, bool) {
	if o == nil || IsNil(o.Remote) {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *ContainerRepositorySyncURL) HasRemote() bool {
	if o != nil && !IsNil(o.Remote) {
		return true
	}

	return false
}

// SetRemote gets a reference to the given string and assigns it to the Remote field.
func (o *ContainerRepositorySyncURL) SetRemote(v string) {
	o.Remote = &v
}

// GetMirror returns the Mirror field value if set, zero value otherwise.
func (o *ContainerRepositorySyncURL) GetMirror() bool {
	if o == nil || IsNil(o.Mirror) {
		var ret bool
		return ret
	}
	return *o.Mirror
}

// GetMirrorOk returns a tuple with the Mirror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRepositorySyncURL) GetMirrorOk() (*bool, bool) {
	if o == nil || IsNil(o.Mirror) {
		return nil, false
	}
	return o.Mirror, true
}

// HasMirror returns a boolean if a field has been set.
func (o *ContainerRepositorySyncURL) HasMirror() bool {
	if o != nil && !IsNil(o.Mirror) {
		return true
	}

	return false
}

// SetMirror gets a reference to the given bool and assigns it to the Mirror field.
func (o *ContainerRepositorySyncURL) SetMirror(v bool) {
	o.Mirror = &v
}

// GetSignedOnly returns the SignedOnly field value if set, zero value otherwise.
func (o *ContainerRepositorySyncURL) GetSignedOnly() bool {
	if o == nil || IsNil(o.SignedOnly) {
		var ret bool
		return ret
	}
	return *o.SignedOnly
}

// GetSignedOnlyOk returns a tuple with the SignedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRepositorySyncURL) GetSignedOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.SignedOnly) {
		return nil, false
	}
	return o.SignedOnly, true
}

// HasSignedOnly returns a boolean if a field has been set.
func (o *ContainerRepositorySyncURL) HasSignedOnly() bool {
	if o != nil && !IsNil(o.SignedOnly) {
		return true
	}

	return false
}

// SetSignedOnly gets a reference to the given bool and assigns it to the SignedOnly field.
func (o *ContainerRepositorySyncURL) SetSignedOnly(v bool) {
	o.SignedOnly = &v
}

func (o ContainerRepositorySyncURL) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRepositorySyncURL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Remote) {
		toSerialize["remote"] = o.Remote
	}
	if !IsNil(o.Mirror) {
		toSerialize["mirror"] = o.Mirror
	}
	if !IsNil(o.SignedOnly) {
		toSerialize["signed_only"] = o.SignedOnly
	}
	return toSerialize, nil
}

type NullableContainerRepositorySyncURL struct {
	value *ContainerRepositorySyncURL
	isSet bool
}

func (v NullableContainerRepositorySyncURL) Get() *ContainerRepositorySyncURL {
	return v.value
}

func (v *NullableContainerRepositorySyncURL) Set(val *ContainerRepositorySyncURL) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRepositorySyncURL) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRepositorySyncURL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRepositorySyncURL(val *ContainerRepositorySyncURL) *NullableContainerRepositorySyncURL {
	return &NullableContainerRepositorySyncURL{value: val, isSet: true}
}

func (v NullableContainerRepositorySyncURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRepositorySyncURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


