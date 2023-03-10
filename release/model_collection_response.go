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

// checks if the CollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponse{}

// CollectionResponse A serializer for a Collection.
type CollectionResponse struct {
	Href *string `json:"href,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Name *string `json:"name,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	VersionsUrl *string `json:"versions_url,omitempty"`
	HighestVersion map[string]interface{} `json:"highest_version,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewCollectionResponse instantiates a new CollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponse() *CollectionResponse {
	this := CollectionResponse{}
	return &this
}

// NewCollectionResponseWithDefaults instantiates a new CollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseWithDefaults() *CollectionResponse {
	this := CollectionResponse{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *CollectionResponse) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *CollectionResponse) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *CollectionResponse) SetHref(v string) {
	o.Href = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CollectionResponse) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CollectionResponse) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CollectionResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectionResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectionResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CollectionResponse) SetName(v string) {
	o.Name = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *CollectionResponse) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *CollectionResponse) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *CollectionResponse) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetVersionsUrl returns the VersionsUrl field value if set, zero value otherwise.
func (o *CollectionResponse) GetVersionsUrl() string {
	if o == nil || IsNil(o.VersionsUrl) {
		var ret string
		return ret
	}
	return *o.VersionsUrl
}

// GetVersionsUrlOk returns a tuple with the VersionsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetVersionsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VersionsUrl) {
		return nil, false
	}
	return o.VersionsUrl, true
}

// HasVersionsUrl returns a boolean if a field has been set.
func (o *CollectionResponse) HasVersionsUrl() bool {
	if o != nil && !IsNil(o.VersionsUrl) {
		return true
	}

	return false
}

// SetVersionsUrl gets a reference to the given string and assigns it to the VersionsUrl field.
func (o *CollectionResponse) SetVersionsUrl(v string) {
	o.VersionsUrl = &v
}

// GetHighestVersion returns the HighestVersion field value if set, zero value otherwise.
func (o *CollectionResponse) GetHighestVersion() map[string]interface{} {
	if o == nil || IsNil(o.HighestVersion) {
		var ret map[string]interface{}
		return ret
	}
	return o.HighestVersion
}

// GetHighestVersionOk returns a tuple with the HighestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetHighestVersionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.HighestVersion) {
		return map[string]interface{}{}, false
	}
	return o.HighestVersion, true
}

// HasHighestVersion returns a boolean if a field has been set.
func (o *CollectionResponse) HasHighestVersion() bool {
	if o != nil && !IsNil(o.HighestVersion) {
		return true
	}

	return false
}

// SetHighestVersion gets a reference to the given map[string]interface{} and assigns it to the HighestVersion field.
func (o *CollectionResponse) SetHighestVersion(v map[string]interface{}) {
	o.HighestVersion = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CollectionResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CollectionResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CollectionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CollectionResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CollectionResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CollectionResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: href is readOnly
	// skip: namespace is readOnly
	// skip: name is readOnly
	// skip: deprecated is readOnly
	// skip: versions_url is readOnly
	// skip: highest_version is readOnly
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	return toSerialize, nil
}

type NullableCollectionResponse struct {
	value *CollectionResponse
	isSet bool
}

func (v NullableCollectionResponse) Get() *CollectionResponse {
	return v.value
}

func (v *NullableCollectionResponse) Set(val *CollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponse(val *CollectionResponse) *NullableCollectionResponse {
	return &NullableCollectionResponse{value: val, isSet: true}
}

func (v NullableCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


