# HuggingFaceHuggingFacePublicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 

## Methods

### NewHuggingFaceHuggingFacePublicationResponse

`func NewHuggingFaceHuggingFacePublicationResponse() *HuggingFaceHuggingFacePublicationResponse`

NewHuggingFaceHuggingFacePublicationResponse instantiates a new HuggingFaceHuggingFacePublicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceHuggingFacePublicationResponseWithDefaults

`func NewHuggingFaceHuggingFacePublicationResponseWithDefaults() *HuggingFaceHuggingFacePublicationResponse`

NewHuggingFaceHuggingFacePublicationResponseWithDefaults instantiates a new HuggingFaceHuggingFacePublicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *HuggingFaceHuggingFacePublicationResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *HuggingFaceHuggingFacePublicationResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *HuggingFaceHuggingFacePublicationResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *HuggingFaceHuggingFacePublicationResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *HuggingFaceHuggingFacePublicationResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *HuggingFaceHuggingFacePublicationResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *HuggingFaceHuggingFacePublicationResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *HuggingFaceHuggingFacePublicationResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *HuggingFaceHuggingFacePublicationResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetRepositoryVersion

`func (o *HuggingFaceHuggingFacePublicationResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *HuggingFaceHuggingFacePublicationResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *HuggingFaceHuggingFacePublicationResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *HuggingFaceHuggingFacePublicationResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *HuggingFaceHuggingFacePublicationResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HuggingFaceHuggingFacePublicationResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HuggingFaceHuggingFacePublicationResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HuggingFaceHuggingFacePublicationResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


