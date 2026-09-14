# MavenMavenIndexPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**VulnReport** | Pointer to **string** |  | [optional] [readonly] 
**Path** | Pointer to **string** | Directory path this page indexes. | [optional] [readonly] 
**Sha256** | Pointer to **string** | SHA256 digest of the HTML content. | [optional] [readonly] 

## Methods

### NewMavenMavenIndexPageResponse

`func NewMavenMavenIndexPageResponse() *MavenMavenIndexPageResponse`

NewMavenMavenIndexPageResponse instantiates a new MavenMavenIndexPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenIndexPageResponseWithDefaults

`func NewMavenMavenIndexPageResponseWithDefaults() *MavenMavenIndexPageResponse`

NewMavenMavenIndexPageResponseWithDefaults instantiates a new MavenMavenIndexPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *MavenMavenIndexPageResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *MavenMavenIndexPageResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *MavenMavenIndexPageResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *MavenMavenIndexPageResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *MavenMavenIndexPageResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *MavenMavenIndexPageResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *MavenMavenIndexPageResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *MavenMavenIndexPageResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *MavenMavenIndexPageResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *MavenMavenIndexPageResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *MavenMavenIndexPageResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *MavenMavenIndexPageResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *MavenMavenIndexPageResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *MavenMavenIndexPageResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *MavenMavenIndexPageResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *MavenMavenIndexPageResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *MavenMavenIndexPageResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *MavenMavenIndexPageResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *MavenMavenIndexPageResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *MavenMavenIndexPageResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetVulnReport

`func (o *MavenMavenIndexPageResponse) GetVulnReport() string`

GetVulnReport returns the VulnReport field if non-nil, zero value otherwise.

### GetVulnReportOk

`func (o *MavenMavenIndexPageResponse) GetVulnReportOk() (*string, bool)`

GetVulnReportOk returns a tuple with the VulnReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnReport

`func (o *MavenMavenIndexPageResponse) SetVulnReport(v string)`

SetVulnReport sets VulnReport field to given value.

### HasVulnReport

`func (o *MavenMavenIndexPageResponse) HasVulnReport() bool`

HasVulnReport returns a boolean if a field has been set.

### GetPath

`func (o *MavenMavenIndexPageResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MavenMavenIndexPageResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MavenMavenIndexPageResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MavenMavenIndexPageResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSha256

`func (o *MavenMavenIndexPageResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *MavenMavenIndexPageResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *MavenMavenIndexPageResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *MavenMavenIndexPageResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


