# RepositoryVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Number** | Pointer to **int64** |  | [optional] [readonly] 
**Repository** | Pointer to **string** |  | [optional] [readonly] 
**BaseVersion** | Pointer to **string** | A repository version whose content was used as the initial set of content for this repository version | [optional] 
**ContentSummary** | Pointer to [**ContentSummaryResponse**](ContentSummaryResponse.md) | Various count summaries of the content in the version and the HREF to view them. | [optional] [readonly] 

## Methods

### NewRepositoryVersionResponse

`func NewRepositoryVersionResponse() *RepositoryVersionResponse`

NewRepositoryVersionResponse instantiates a new RepositoryVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryVersionResponseWithDefaults

`func NewRepositoryVersionResponseWithDefaults() *RepositoryVersionResponse`

NewRepositoryVersionResponseWithDefaults instantiates a new RepositoryVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RepositoryVersionResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RepositoryVersionResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RepositoryVersionResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RepositoryVersionResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *RepositoryVersionResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *RepositoryVersionResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *RepositoryVersionResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *RepositoryVersionResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RepositoryVersionResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RepositoryVersionResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RepositoryVersionResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RepositoryVersionResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *RepositoryVersionResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *RepositoryVersionResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *RepositoryVersionResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *RepositoryVersionResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetNumber

`func (o *RepositoryVersionResponse) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *RepositoryVersionResponse) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *RepositoryVersionResponse) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *RepositoryVersionResponse) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRepository

`func (o *RepositoryVersionResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepositoryVersionResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepositoryVersionResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RepositoryVersionResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetBaseVersion

`func (o *RepositoryVersionResponse) GetBaseVersion() string`

GetBaseVersion returns the BaseVersion field if non-nil, zero value otherwise.

### GetBaseVersionOk

`func (o *RepositoryVersionResponse) GetBaseVersionOk() (*string, bool)`

GetBaseVersionOk returns a tuple with the BaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersion

`func (o *RepositoryVersionResponse) SetBaseVersion(v string)`

SetBaseVersion sets BaseVersion field to given value.

### HasBaseVersion

`func (o *RepositoryVersionResponse) HasBaseVersion() bool`

HasBaseVersion returns a boolean if a field has been set.

### GetContentSummary

`func (o *RepositoryVersionResponse) GetContentSummary() ContentSummaryResponse`

GetContentSummary returns the ContentSummary field if non-nil, zero value otherwise.

### GetContentSummaryOk

`func (o *RepositoryVersionResponse) GetContentSummaryOk() (*ContentSummaryResponse, bool)`

GetContentSummaryOk returns a tuple with the ContentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSummary

`func (o *RepositoryVersionResponse) SetContentSummary(v ContentSummaryResponse)`

SetContentSummary sets ContentSummary field to given value.

### HasContentSummary

`func (o *RepositoryVersionResponse) HasContentSummary() bool`

HasContentSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


