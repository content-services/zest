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

// checks if the OstreeOstreeObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeOstreeObjectResponse{}

// OstreeOstreeObjectResponse A Serializer class for OSTree objects (e.g., dirtree, dirmeta, file).
type OstreeOstreeObjectResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// The Pulp Resource Name (PRN).
	Prn *string `json:"prn,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
	Checksum string `json:"checksum"`
	//             The type of an object. All values are described by the mapping declared at            https://lazka.github.io/pgi-docs/OSTree-1.0/enums.html#OSTree.ObjectType            
	Typ int64 `json:"typ"`
	AdditionalProperties map[string]interface{}
}

type _OstreeOstreeObjectResponse OstreeOstreeObjectResponse

// NewOstreeOstreeObjectResponse instantiates a new OstreeOstreeObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeOstreeObjectResponse(artifact string, relativePath string, checksum string, typ int64) *OstreeOstreeObjectResponse {
	this := OstreeOstreeObjectResponse{}
	this.Artifact = artifact
	this.RelativePath = relativePath
	this.Checksum = checksum
	this.Typ = typ
	return &this
}

// NewOstreeOstreeObjectResponseWithDefaults instantiates a new OstreeOstreeObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeOstreeObjectResponseWithDefaults() *OstreeOstreeObjectResponse {
	this := OstreeOstreeObjectResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *OstreeOstreeObjectResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *OstreeOstreeObjectResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *OstreeOstreeObjectResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *OstreeOstreeObjectResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *OstreeOstreeObjectResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *OstreeOstreeObjectResponse) SetPrn(v string) {
	o.Prn = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *OstreeOstreeObjectResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *OstreeOstreeObjectResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *OstreeOstreeObjectResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *OstreeOstreeObjectResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *OstreeOstreeObjectResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *OstreeOstreeObjectResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetArtifact returns the Artifact field value
func (o *OstreeOstreeObjectResponse) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *OstreeOstreeObjectResponse) SetArtifact(v string) {
	o.Artifact = v
}

// GetRelativePath returns the RelativePath field value
func (o *OstreeOstreeObjectResponse) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *OstreeOstreeObjectResponse) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetChecksum returns the Checksum field value
func (o *OstreeOstreeObjectResponse) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value
func (o *OstreeOstreeObjectResponse) SetChecksum(v string) {
	o.Checksum = v
}

// GetTyp returns the Typ field value
func (o *OstreeOstreeObjectResponse) GetTyp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Typ
}

// GetTypOk returns a tuple with the Typ field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeObjectResponse) GetTypOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typ, true
}

// SetTyp sets field value
func (o *OstreeOstreeObjectResponse) SetTyp(v int64) {
	o.Typ = v
}

func (o OstreeOstreeObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeOstreeObjectResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["artifact"] = o.Artifact
	toSerialize["relative_path"] = o.RelativePath
	toSerialize["checksum"] = o.Checksum
	toSerialize["typ"] = o.Typ

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OstreeOstreeObjectResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"artifact",
		"relative_path",
		"checksum",
		"typ",
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

	varOstreeOstreeObjectResponse := _OstreeOstreeObjectResponse{}

	err = json.Unmarshal(data, &varOstreeOstreeObjectResponse)

	if err != nil {
		return err
	}

	*o = OstreeOstreeObjectResponse(varOstreeOstreeObjectResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "artifact")
		delete(additionalProperties, "relative_path")
		delete(additionalProperties, "checksum")
		delete(additionalProperties, "typ")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOstreeOstreeObjectResponse struct {
	value *OstreeOstreeObjectResponse
	isSet bool
}

func (v NullableOstreeOstreeObjectResponse) Get() *OstreeOstreeObjectResponse {
	return v.value
}

func (v *NullableOstreeOstreeObjectResponse) Set(val *OstreeOstreeObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeOstreeObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeOstreeObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeOstreeObjectResponse(val *OstreeOstreeObjectResponse) *NullableOstreeOstreeObjectResponse {
	return &NullableOstreeOstreeObjectResponse{value: val, isSet: true}
}

func (v NullableOstreeOstreeObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeOstreeObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


