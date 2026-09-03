# ContentViewErrataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Id** | **string** | Update id (e.g. RHEA-2013:1777) | 
**UpdatedDate** | **string** | Date when the update was updated | 
**IssuedDate** | **string** | Date when the update was issued | 
**Description** | **string** | Update description | 
**Title** | **string** | Update name | 
**Summary** | **string** | Short summary | 
**Version** | **string** | Update version | 
**Type** | **string** | Update type (&#39;enhancement&#39;, &#39;bugfix&#39;, ...) | 
**Severity** | **string** | Severity | 
**Solution** | **string** | Solution | 
**Release** | **string** | Update release | 
**Rights** | **string** | Copyrights | 
**RebootSuggested** | **bool** | Whether a reboot is suggested | 
**Cves** | Pointer to **string** | CVE references attached to this errata | [optional] [readonly] 

## Methods

### NewContentViewErrataResponse

`func NewContentViewErrataResponse(id string, updatedDate string, issuedDate string, description string, title string, summary string, version string, type_ string, severity string, solution string, release string, rights string, rebootSuggested bool, ) *ContentViewErrataResponse`

NewContentViewErrataResponse instantiates a new ContentViewErrataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentViewErrataResponseWithDefaults

`func NewContentViewErrataResponseWithDefaults() *ContentViewErrataResponse`

NewContentViewErrataResponseWithDefaults instantiates a new ContentViewErrataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContentViewErrataResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContentViewErrataResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContentViewErrataResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContentViewErrataResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetId

`func (o *ContentViewErrataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentViewErrataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentViewErrataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedDate

`func (o *ContentViewErrataResponse) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ContentViewErrataResponse) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ContentViewErrataResponse) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.


### GetIssuedDate

`func (o *ContentViewErrataResponse) GetIssuedDate() string`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *ContentViewErrataResponse) GetIssuedDateOk() (*string, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *ContentViewErrataResponse) SetIssuedDate(v string)`

SetIssuedDate sets IssuedDate field to given value.


### GetDescription

`func (o *ContentViewErrataResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentViewErrataResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentViewErrataResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTitle

`func (o *ContentViewErrataResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentViewErrataResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentViewErrataResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *ContentViewErrataResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ContentViewErrataResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ContentViewErrataResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetVersion

`func (o *ContentViewErrataResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentViewErrataResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentViewErrataResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *ContentViewErrataResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentViewErrataResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentViewErrataResponse) SetType(v string)`

SetType sets Type field to given value.


### GetSeverity

`func (o *ContentViewErrataResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ContentViewErrataResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ContentViewErrataResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSolution

`func (o *ContentViewErrataResponse) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *ContentViewErrataResponse) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *ContentViewErrataResponse) SetSolution(v string)`

SetSolution sets Solution field to given value.


### GetRelease

`func (o *ContentViewErrataResponse) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ContentViewErrataResponse) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ContentViewErrataResponse) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetRights

`func (o *ContentViewErrataResponse) GetRights() string`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *ContentViewErrataResponse) GetRightsOk() (*string, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *ContentViewErrataResponse) SetRights(v string)`

SetRights sets Rights field to given value.


### GetRebootSuggested

`func (o *ContentViewErrataResponse) GetRebootSuggested() bool`

GetRebootSuggested returns the RebootSuggested field if non-nil, zero value otherwise.

### GetRebootSuggestedOk

`func (o *ContentViewErrataResponse) GetRebootSuggestedOk() (*bool, bool)`

GetRebootSuggestedOk returns a tuple with the RebootSuggested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootSuggested

`func (o *ContentViewErrataResponse) SetRebootSuggested(v bool)`

SetRebootSuggested sets RebootSuggested field to given value.


### GetCves

`func (o *ContentViewErrataResponse) GetCves() string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *ContentViewErrataResponse) GetCvesOk() (*string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *ContentViewErrataResponse) SetCves(v string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *ContentViewErrataResponse) HasCves() bool`

HasCves returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


