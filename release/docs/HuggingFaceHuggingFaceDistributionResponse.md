# HuggingFaceHuggingFaceDistributionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 
**BaseUrl** | Pointer to **string** | The URL for accessing the publication as defined by this distribution. | [optional] [readonly] 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**ContentGuardPrn** | Pointer to **string** | The Pulp Resource Name (PRN) of the associated optional content guard. | [optional] [readonly] 
**NoContentChangeSince** | Pointer to **string** | Timestamp since when the distributed content served by this distribution has not changed. If equals to &#x60;null&#x60;, no guarantee is provided about content changes. | [optional] [readonly] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 
**Publication** | Pointer to **NullableString** | Publication to be served | [optional] 
**Remote** | Pointer to **NullableString** | Remote that can be used to fetch content when using pull-through caching. | [optional] 

## Methods

### NewHuggingFaceHuggingFaceDistributionResponse

`func NewHuggingFaceHuggingFaceDistributionResponse(basePath string, name string, ) *HuggingFaceHuggingFaceDistributionResponse`

NewHuggingFaceHuggingFaceDistributionResponse instantiates a new HuggingFaceHuggingFaceDistributionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceHuggingFaceDistributionResponseWithDefaults

`func NewHuggingFaceHuggingFaceDistributionResponseWithDefaults() *HuggingFaceHuggingFaceDistributionResponse`

NewHuggingFaceHuggingFaceDistributionResponseWithDefaults instantiates a new HuggingFaceHuggingFaceDistributionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetBasePath

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetBaseUrl

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetContentGuard

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *HuggingFaceHuggingFaceDistributionResponse) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetContentGuardPrn

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetContentGuardPrn() string`

GetContentGuardPrn returns the ContentGuardPrn field if non-nil, zero value otherwise.

### GetContentGuardPrnOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetContentGuardPrnOk() (*string, bool)`

GetContentGuardPrnOk returns a tuple with the ContentGuardPrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuardPrn

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetContentGuardPrn(v string)`

SetContentGuardPrn sets ContentGuardPrn field to given value.

### HasContentGuardPrn

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasContentGuardPrn() bool`

HasContentGuardPrn returns a boolean if a field has been set.

### GetNoContentChangeSince

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetNoContentChangeSince() string`

GetNoContentChangeSince returns the NoContentChangeSince field if non-nil, zero value otherwise.

### GetNoContentChangeSinceOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetNoContentChangeSinceOk() (*string, bool)`

GetNoContentChangeSinceOk returns a tuple with the NoContentChangeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoContentChangeSince

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetNoContentChangeSince(v string)`

SetNoContentChangeSince sets NoContentChangeSince field to given value.

### HasNoContentChangeSince

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasNoContentChangeSince() bool`

HasNoContentChangeSince returns a boolean if a field has been set.

### GetHidden

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRepository

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *HuggingFaceHuggingFaceDistributionResponse) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRepositoryVersion

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *HuggingFaceHuggingFaceDistributionResponse) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetPublication

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPublication() string`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetPublicationOk() (*string, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetPublication(v string)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *HuggingFaceHuggingFaceDistributionResponse) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetRemote

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *HuggingFaceHuggingFaceDistributionResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *HuggingFaceHuggingFaceDistributionResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *HuggingFaceHuggingFaceDistributionResponse) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *HuggingFaceHuggingFaceDistributionResponse) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


