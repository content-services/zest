# ArtifactDistributionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The URL for accessing the publication as defined by this distribution. | [optional] [readonly] 
**NoContentChangeSince** | Pointer to **string** | Timestamp since when the distributed content served by this distribution has not changed. If equals to &#x60;null&#x60;, no guarantee is provided about content changes. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 

## Methods

### NewArtifactDistributionResponse

`func NewArtifactDistributionResponse(name string, basePath string, ) *ArtifactDistributionResponse`

NewArtifactDistributionResponse instantiates a new ArtifactDistributionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDistributionResponseWithDefaults

`func NewArtifactDistributionResponseWithDefaults() *ArtifactDistributionResponse`

NewArtifactDistributionResponseWithDefaults instantiates a new ArtifactDistributionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *ArtifactDistributionResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ArtifactDistributionResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ArtifactDistributionResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ArtifactDistributionResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetNoContentChangeSince

`func (o *ArtifactDistributionResponse) GetNoContentChangeSince() string`

GetNoContentChangeSince returns the NoContentChangeSince field if non-nil, zero value otherwise.

### GetNoContentChangeSinceOk

`func (o *ArtifactDistributionResponse) GetNoContentChangeSinceOk() (*string, bool)`

GetNoContentChangeSinceOk returns a tuple with the NoContentChangeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoContentChangeSince

`func (o *ArtifactDistributionResponse) SetNoContentChangeSince(v string)`

SetNoContentChangeSince sets NoContentChangeSince field to given value.

### HasNoContentChangeSince

`func (o *ArtifactDistributionResponse) HasNoContentChangeSince() bool`

HasNoContentChangeSince returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *ArtifactDistributionResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *ArtifactDistributionResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *ArtifactDistributionResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *ArtifactDistributionResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ArtifactDistributionResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ArtifactDistributionResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ArtifactDistributionResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ArtifactDistributionResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpHref

`func (o *ArtifactDistributionResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ArtifactDistributionResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ArtifactDistributionResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ArtifactDistributionResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetContentGuard

`func (o *ArtifactDistributionResponse) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *ArtifactDistributionResponse) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *ArtifactDistributionResponse) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *ArtifactDistributionResponse) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *ArtifactDistributionResponse) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *ArtifactDistributionResponse) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetPrn

`func (o *ArtifactDistributionResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *ArtifactDistributionResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *ArtifactDistributionResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *ArtifactDistributionResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpLabels

`func (o *ArtifactDistributionResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ArtifactDistributionResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ArtifactDistributionResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ArtifactDistributionResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *ArtifactDistributionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactDistributionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactDistributionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetHidden

`func (o *ArtifactDistributionResponse) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ArtifactDistributionResponse) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ArtifactDistributionResponse) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ArtifactDistributionResponse) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetBasePath

`func (o *ArtifactDistributionResponse) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *ArtifactDistributionResponse) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *ArtifactDistributionResponse) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


