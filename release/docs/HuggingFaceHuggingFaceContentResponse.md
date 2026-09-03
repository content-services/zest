# HuggingFaceHuggingFaceContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**VulnReport** | Pointer to **string** |  | [optional] [readonly] 
**Artifact** | **string** | Artifact file representing the physical content | 
**RelativePath** | **string** | The relative path within the repository | 
**RepoId** | **string** | The Hugging Face repository ID (e.g., &#39;microsoft/DialoGPT-medium&#39;) | 
**RepoType** | Pointer to [**RepoTypeEnum**](RepoTypeEnum.md) | The type of Hugging Face repository* &#x60;models&#x60; - Models* &#x60;datasets&#x60; - Datasets* &#x60;spaces&#x60; - Spaces | [optional] [default to REPOTYPEENUM_MODELS]
**Revision** | Pointer to **string** | The git revision/branch/tag | [optional] [default to "main"]
**Size** | Pointer to **NullableInt64** | File size in bytes | [optional] 
**Etag** | Pointer to **NullableString** | ETag from Hugging Face | [optional] 
**LastModified** | Pointer to **NullableTime** | Last modified timestamp | [optional] 

## Methods

### NewHuggingFaceHuggingFaceContentResponse

`func NewHuggingFaceHuggingFaceContentResponse(artifact string, relativePath string, repoId string, ) *HuggingFaceHuggingFaceContentResponse`

NewHuggingFaceHuggingFaceContentResponse instantiates a new HuggingFaceHuggingFaceContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceHuggingFaceContentResponseWithDefaults

`func NewHuggingFaceHuggingFaceContentResponseWithDefaults() *HuggingFaceHuggingFaceContentResponse`

NewHuggingFaceHuggingFaceContentResponseWithDefaults instantiates a new HuggingFaceHuggingFaceContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *HuggingFaceHuggingFaceContentResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *HuggingFaceHuggingFaceContentResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *HuggingFaceHuggingFaceContentResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *HuggingFaceHuggingFaceContentResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *HuggingFaceHuggingFaceContentResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *HuggingFaceHuggingFaceContentResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *HuggingFaceHuggingFaceContentResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *HuggingFaceHuggingFaceContentResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *HuggingFaceHuggingFaceContentResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *HuggingFaceHuggingFaceContentResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *HuggingFaceHuggingFaceContentResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetVulnReport

`func (o *HuggingFaceHuggingFaceContentResponse) GetVulnReport() string`

GetVulnReport returns the VulnReport field if non-nil, zero value otherwise.

### GetVulnReportOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetVulnReportOk() (*string, bool)`

GetVulnReportOk returns a tuple with the VulnReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnReport

`func (o *HuggingFaceHuggingFaceContentResponse) SetVulnReport(v string)`

SetVulnReport sets VulnReport field to given value.

### HasVulnReport

`func (o *HuggingFaceHuggingFaceContentResponse) HasVulnReport() bool`

HasVulnReport returns a boolean if a field has been set.

### GetArtifact

`func (o *HuggingFaceHuggingFaceContentResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *HuggingFaceHuggingFaceContentResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRelativePath

`func (o *HuggingFaceHuggingFaceContentResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HuggingFaceHuggingFaceContentResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetRepoId

`func (o *HuggingFaceHuggingFaceContentResponse) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *HuggingFaceHuggingFaceContentResponse) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.


### GetRepoType

`func (o *HuggingFaceHuggingFaceContentResponse) GetRepoType() RepoTypeEnum`

GetRepoType returns the RepoType field if non-nil, zero value otherwise.

### GetRepoTypeOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetRepoTypeOk() (*RepoTypeEnum, bool)`

GetRepoTypeOk returns a tuple with the RepoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoType

`func (o *HuggingFaceHuggingFaceContentResponse) SetRepoType(v RepoTypeEnum)`

SetRepoType sets RepoType field to given value.

### HasRepoType

`func (o *HuggingFaceHuggingFaceContentResponse) HasRepoType() bool`

HasRepoType returns a boolean if a field has been set.

### GetRevision

`func (o *HuggingFaceHuggingFaceContentResponse) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HuggingFaceHuggingFaceContentResponse) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *HuggingFaceHuggingFaceContentResponse) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSize

`func (o *HuggingFaceHuggingFaceContentResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HuggingFaceHuggingFaceContentResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *HuggingFaceHuggingFaceContentResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *HuggingFaceHuggingFaceContentResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *HuggingFaceHuggingFaceContentResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetEtag

`func (o *HuggingFaceHuggingFaceContentResponse) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *HuggingFaceHuggingFaceContentResponse) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *HuggingFaceHuggingFaceContentResponse) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### SetEtagNil

`func (o *HuggingFaceHuggingFaceContentResponse) SetEtagNil(b bool)`

 SetEtagNil sets the value for Etag to be an explicit nil

### UnsetEtag
`func (o *HuggingFaceHuggingFaceContentResponse) UnsetEtag()`

UnsetEtag ensures that no value is present for Etag, not even an explicit nil
### GetLastModified

`func (o *HuggingFaceHuggingFaceContentResponse) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *HuggingFaceHuggingFaceContentResponse) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *HuggingFaceHuggingFaceContentResponse) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *HuggingFaceHuggingFaceContentResponse) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *HuggingFaceHuggingFaceContentResponse) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *HuggingFaceHuggingFaceContentResponse) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


