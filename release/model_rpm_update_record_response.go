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

// checks if the RpmUpdateRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmUpdateRecordResponse{}

// RpmUpdateRecordResponse A Serializer for UpdateRecord.
type RpmUpdateRecordResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Update id (short update name, e.g. RHEA-2013:1777)
	Id *string `json:"id,omitempty"`
	// Date when the update was updated (e.g. '2013-12-02 00:00:00')
	UpdatedDate *string `json:"updated_date,omitempty"`
	// Update description
	Description *string `json:"description,omitempty"`
	// Date when the update was issued (e.g. '2013-12-02 00:00:00')
	IssuedDate *string `json:"issued_date,omitempty"`
	// Source of the update (e.g. security@redhat.com)
	Fromstr *string `json:"fromstr,omitempty"`
	// Update status ('final', ...)
	Status *string `json:"status,omitempty"`
	// Update name
	Title *string `json:"title,omitempty"`
	// Short summary
	Summary *string `json:"summary,omitempty"`
	// Update version (probably always an integer number)
	Version *string `json:"version,omitempty"`
	// Update type ('enhancement', 'bugfix', ...)
	Type *string `json:"type,omitempty"`
	// Severity
	Severity *string `json:"severity,omitempty"`
	// Solution
	Solution *string `json:"solution,omitempty"`
	// Update release
	Release *string `json:"release,omitempty"`
	// Copyrights
	Rights *string `json:"rights,omitempty"`
	// Push count
	Pushcount *string `json:"pushcount,omitempty"`
	// List of packages
	Pkglist []RpmUpdateCollectionResponse `json:"pkglist,omitempty"`
	// List of references
	References []map[string]interface{} `json:"references,omitempty"`
	// Reboot suggested
	RebootSuggested *bool `json:"reboot_suggested,omitempty"`
}

// NewRpmUpdateRecordResponse instantiates a new RpmUpdateRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmUpdateRecordResponse() *RpmUpdateRecordResponse {
	this := RpmUpdateRecordResponse{}
	return &this
}

// NewRpmUpdateRecordResponseWithDefaults instantiates a new RpmUpdateRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmUpdateRecordResponseWithDefaults() *RpmUpdateRecordResponse {
	this := RpmUpdateRecordResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmUpdateRecordResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmUpdateRecordResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RpmUpdateRecordResponse) SetId(v string) {
	o.Id = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetUpdatedDate() string {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetUpdatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *RpmUpdateRecordResponse) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RpmUpdateRecordResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIssuedDate returns the IssuedDate field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetIssuedDate() string {
	if o == nil || IsNil(o.IssuedDate) {
		var ret string
		return ret
	}
	return *o.IssuedDate
}

// GetIssuedDateOk returns a tuple with the IssuedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetIssuedDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssuedDate) {
		return nil, false
	}
	return o.IssuedDate, true
}

// HasIssuedDate returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasIssuedDate() bool {
	if o != nil && !IsNil(o.IssuedDate) {
		return true
	}

	return false
}

// SetIssuedDate gets a reference to the given string and assigns it to the IssuedDate field.
func (o *RpmUpdateRecordResponse) SetIssuedDate(v string) {
	o.IssuedDate = &v
}

// GetFromstr returns the Fromstr field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetFromstr() string {
	if o == nil || IsNil(o.Fromstr) {
		var ret string
		return ret
	}
	return *o.Fromstr
}

// GetFromstrOk returns a tuple with the Fromstr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetFromstrOk() (*string, bool) {
	if o == nil || IsNil(o.Fromstr) {
		return nil, false
	}
	return o.Fromstr, true
}

// HasFromstr returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasFromstr() bool {
	if o != nil && !IsNil(o.Fromstr) {
		return true
	}

	return false
}

// SetFromstr gets a reference to the given string and assigns it to the Fromstr field.
func (o *RpmUpdateRecordResponse) SetFromstr(v string) {
	o.Fromstr = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RpmUpdateRecordResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RpmUpdateRecordResponse) SetTitle(v string) {
	o.Title = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *RpmUpdateRecordResponse) SetSummary(v string) {
	o.Summary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RpmUpdateRecordResponse) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RpmUpdateRecordResponse) SetType(v string) {
	o.Type = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *RpmUpdateRecordResponse) SetSeverity(v string) {
	o.Severity = &v
}

// GetSolution returns the Solution field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetSolution() string {
	if o == nil || IsNil(o.Solution) {
		var ret string
		return ret
	}
	return *o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetSolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Solution) {
		return nil, false
	}
	return o.Solution, true
}

// HasSolution returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasSolution() bool {
	if o != nil && !IsNil(o.Solution) {
		return true
	}

	return false
}

// SetSolution gets a reference to the given string and assigns it to the Solution field.
func (o *RpmUpdateRecordResponse) SetSolution(v string) {
	o.Solution = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *RpmUpdateRecordResponse) SetRelease(v string) {
	o.Release = &v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetRights() string {
	if o == nil || IsNil(o.Rights) {
		var ret string
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetRightsOk() (*string, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given string and assigns it to the Rights field.
func (o *RpmUpdateRecordResponse) SetRights(v string) {
	o.Rights = &v
}

// GetPushcount returns the Pushcount field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetPushcount() string {
	if o == nil || IsNil(o.Pushcount) {
		var ret string
		return ret
	}
	return *o.Pushcount
}

// GetPushcountOk returns a tuple with the Pushcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetPushcountOk() (*string, bool) {
	if o == nil || IsNil(o.Pushcount) {
		return nil, false
	}
	return o.Pushcount, true
}

// HasPushcount returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasPushcount() bool {
	if o != nil && !IsNil(o.Pushcount) {
		return true
	}

	return false
}

// SetPushcount gets a reference to the given string and assigns it to the Pushcount field.
func (o *RpmUpdateRecordResponse) SetPushcount(v string) {
	o.Pushcount = &v
}

// GetPkglist returns the Pkglist field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetPkglist() []RpmUpdateCollectionResponse {
	if o == nil || IsNil(o.Pkglist) {
		var ret []RpmUpdateCollectionResponse
		return ret
	}
	return o.Pkglist
}

// GetPkglistOk returns a tuple with the Pkglist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetPkglistOk() ([]RpmUpdateCollectionResponse, bool) {
	if o == nil || IsNil(o.Pkglist) {
		return nil, false
	}
	return o.Pkglist, true
}

// HasPkglist returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasPkglist() bool {
	if o != nil && !IsNil(o.Pkglist) {
		return true
	}

	return false
}

// SetPkglist gets a reference to the given []RpmUpdateCollectionResponse and assigns it to the Pkglist field.
func (o *RpmUpdateRecordResponse) SetPkglist(v []RpmUpdateCollectionResponse) {
	o.Pkglist = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetReferences() []map[string]interface{} {
	if o == nil || IsNil(o.References) {
		var ret []map[string]interface{}
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetReferencesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []map[string]interface{} and assigns it to the References field.
func (o *RpmUpdateRecordResponse) SetReferences(v []map[string]interface{}) {
	o.References = v
}

// GetRebootSuggested returns the RebootSuggested field value if set, zero value otherwise.
func (o *RpmUpdateRecordResponse) GetRebootSuggested() bool {
	if o == nil || IsNil(o.RebootSuggested) {
		var ret bool
		return ret
	}
	return *o.RebootSuggested
}

// GetRebootSuggestedOk returns a tuple with the RebootSuggested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateRecordResponse) GetRebootSuggestedOk() (*bool, bool) {
	if o == nil || IsNil(o.RebootSuggested) {
		return nil, false
	}
	return o.RebootSuggested, true
}

// HasRebootSuggested returns a boolean if a field has been set.
func (o *RpmUpdateRecordResponse) HasRebootSuggested() bool {
	if o != nil && !IsNil(o.RebootSuggested) {
		return true
	}

	return false
}

// SetRebootSuggested gets a reference to the given bool and assigns it to the RebootSuggested field.
func (o *RpmUpdateRecordResponse) SetRebootSuggested(v bool) {
	o.RebootSuggested = &v
}

func (o RpmUpdateRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmUpdateRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	// skip: id is readOnly
	// skip: updated_date is readOnly
	// skip: description is readOnly
	// skip: issued_date is readOnly
	// skip: fromstr is readOnly
	// skip: status is readOnly
	// skip: title is readOnly
	// skip: summary is readOnly
	// skip: version is readOnly
	// skip: type is readOnly
	// skip: severity is readOnly
	// skip: solution is readOnly
	// skip: release is readOnly
	// skip: rights is readOnly
	// skip: pushcount is readOnly
	// skip: pkglist is readOnly
	// skip: references is readOnly
	// skip: reboot_suggested is readOnly
	return toSerialize, nil
}

type NullableRpmUpdateRecordResponse struct {
	value *RpmUpdateRecordResponse
	isSet bool
}

func (v NullableRpmUpdateRecordResponse) Get() *RpmUpdateRecordResponse {
	return v.value
}

func (v *NullableRpmUpdateRecordResponse) Set(val *RpmUpdateRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmUpdateRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmUpdateRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmUpdateRecordResponse(val *RpmUpdateRecordResponse) *NullableRpmUpdateRecordResponse {
	return &NullableRpmUpdateRecordResponse{value: val, isSet: true}
}

func (v NullableRpmUpdateRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmUpdateRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


