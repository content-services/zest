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

// checks if the ContainerContainerPullThroughDistributionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerContainerPullThroughDistributionResponse{}

// ContainerContainerPullThroughDistributionResponse A serializer for a specialized pull-through distribution referencing sub-distributions.
type ContainerContainerPullThroughDistributionResponse struct {
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp since when the distributed content served by this distribution has not changed. If equals to `null`, no guarantee is provided about content changes.
	NoContentChangeSince *string `json:"no_content_change_since,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// The Pulp Resource Name (PRN).
	Prn *string `json:"prn,omitempty"`
	// Whether this distribution should be shown in the content app.
	Hidden *bool `json:"hidden,omitempty"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// An optional content-guard. If none is specified, a default one will be used.
	ContentGuard *string `json:"content_guard,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// Remote that can be used to fetch content when using pull-through caching.
	Remote string `json:"remote"`
	// Distributions created after pulling content through cache
	Distributions []string `json:"distributions,omitempty"`
	// Namespace this distribution belongs to.
	Namespace *string `json:"namespace,omitempty"`
	// Restrict pull access to explicitly authorized users. Related distributions inherit this value. Defaults to unrestricted pull access.
	Private *bool `json:"private,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContainerContainerPullThroughDistributionResponse ContainerContainerPullThroughDistributionResponse

// NewContainerContainerPullThroughDistributionResponse instantiates a new ContainerContainerPullThroughDistributionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerPullThroughDistributionResponse(name string, basePath string, remote string) *ContainerContainerPullThroughDistributionResponse {
	this := ContainerContainerPullThroughDistributionResponse{}
	this.Name = name
	var hidden bool = false
	this.Hidden = &hidden
	this.BasePath = basePath
	this.Remote = remote
	return &this
}

// NewContainerContainerPullThroughDistributionResponseWithDefaults instantiates a new ContainerContainerPullThroughDistributionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerPullThroughDistributionResponseWithDefaults() *ContainerContainerPullThroughDistributionResponse {
	this := ContainerContainerPullThroughDistributionResponse{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerContainerPullThroughDistributionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerContainerPullThroughDistributionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetNoContentChangeSince returns the NoContentChangeSince field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetNoContentChangeSince() string {
	if o == nil || IsNil(o.NoContentChangeSince) {
		var ret string
		return ret
	}
	return *o.NoContentChangeSince
}

// GetNoContentChangeSinceOk returns a tuple with the NoContentChangeSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetNoContentChangeSinceOk() (*string, bool) {
	if o == nil || IsNil(o.NoContentChangeSince) {
		return nil, false
	}
	return o.NoContentChangeSince, true
}

// HasNoContentChangeSince returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasNoContentChangeSince() bool {
	if o != nil && !IsNil(o.NoContentChangeSince) {
		return true
	}

	return false
}

// SetNoContentChangeSince gets a reference to the given string and assigns it to the NoContentChangeSince field.
func (o *ContainerContainerPullThroughDistributionResponse) SetNoContentChangeSince(v string) {
	o.NoContentChangeSince = &v
}

// GetName returns the Name field value
func (o *ContainerContainerPullThroughDistributionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerPullThroughDistributionResponse) SetName(v string) {
	o.Name = v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *ContainerContainerPullThroughDistributionResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *ContainerContainerPullThroughDistributionResponse) SetPrn(v string) {
	o.Prn = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *ContainerContainerPullThroughDistributionResponse) SetHidden(v bool) {
	o.Hidden = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerPullThroughDistributionResponse) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerPullThroughDistributionResponse) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *ContainerContainerPullThroughDistributionResponse) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *ContainerContainerPullThroughDistributionResponse) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *ContainerContainerPullThroughDistributionResponse) UnsetRepository() {
	o.Repository.Unset()
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard) {
		var ret string
		return ret
	}
	return *o.ContentGuard
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetContentGuardOk() (*string, bool) {
	if o == nil || IsNil(o.ContentGuard) {
		return nil, false
	}
	return o.ContentGuard, true
}

// HasContentGuard returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasContentGuard() bool {
	if o != nil && !IsNil(o.ContentGuard) {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given string and assigns it to the ContentGuard field.
func (o *ContainerContainerPullThroughDistributionResponse) SetContentGuard(v string) {
	o.ContentGuard = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *ContainerContainerPullThroughDistributionResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetBasePath returns the BasePath field value
func (o *ContainerContainerPullThroughDistributionResponse) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *ContainerContainerPullThroughDistributionResponse) SetBasePath(v string) {
	o.BasePath = v
}

// GetRemote returns the Remote field value
func (o *ContainerContainerPullThroughDistributionResponse) GetRemote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remote, true
}

// SetRemote sets field value
func (o *ContainerContainerPullThroughDistributionResponse) SetRemote(v string) {
	o.Remote = v
}

// GetDistributions returns the Distributions field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetDistributions() []string {
	if o == nil || IsNil(o.Distributions) {
		var ret []string
		return ret
	}
	return o.Distributions
}

// GetDistributionsOk returns a tuple with the Distributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetDistributionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Distributions) {
		return nil, false
	}
	return o.Distributions, true
}

// HasDistributions returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasDistributions() bool {
	if o != nil && !IsNil(o.Distributions) {
		return true
	}

	return false
}

// SetDistributions gets a reference to the given []string and assigns it to the Distributions field.
func (o *ContainerContainerPullThroughDistributionResponse) SetDistributions(v []string) {
	o.Distributions = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ContainerContainerPullThroughDistributionResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistributionResponse) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistributionResponse) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ContainerContainerPullThroughDistributionResponse) SetPrivate(v bool) {
	o.Private = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerPullThroughDistributionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerPullThroughDistributionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistributionResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerContainerPullThroughDistributionResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerContainerPullThroughDistributionResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerContainerPullThroughDistributionResponse) UnsetDescription() {
	o.Description.Unset()
}

func (o ContainerContainerPullThroughDistributionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerContainerPullThroughDistributionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.NoContentChangeSince) {
		toSerialize["no_content_change_since"] = o.NoContentChangeSince
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if !IsNil(o.Prn) {
		toSerialize["prn"] = o.Prn
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if !IsNil(o.ContentGuard) {
		toSerialize["content_guard"] = o.ContentGuard
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	toSerialize["base_path"] = o.BasePath
	toSerialize["remote"] = o.Remote
	if !IsNil(o.Distributions) {
		toSerialize["distributions"] = o.Distributions
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContainerContainerPullThroughDistributionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"base_path",
		"remote",
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

	varContainerContainerPullThroughDistributionResponse := _ContainerContainerPullThroughDistributionResponse{}

	err = json.Unmarshal(data, &varContainerContainerPullThroughDistributionResponse)

	if err != nil {
		return err
	}

	*o = ContainerContainerPullThroughDistributionResponse(varContainerContainerPullThroughDistributionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "no_content_change_since")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "remote")
		delete(additionalProperties, "distributions")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "private")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainerContainerPullThroughDistributionResponse struct {
	value *ContainerContainerPullThroughDistributionResponse
	isSet bool
}

func (v NullableContainerContainerPullThroughDistributionResponse) Get() *ContainerContainerPullThroughDistributionResponse {
	return v.value
}

func (v *NullableContainerContainerPullThroughDistributionResponse) Set(val *ContainerContainerPullThroughDistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerPullThroughDistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerPullThroughDistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerPullThroughDistributionResponse(val *ContainerContainerPullThroughDistributionResponse) *NullableContainerContainerPullThroughDistributionResponse {
	return &NullableContainerContainerPullThroughDistributionResponse{value: val, isSet: true}
}

func (v NullableContainerContainerPullThroughDistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerPullThroughDistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


