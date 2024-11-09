# DomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Name** | **string** | A name for this domain. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**StorageClass** | [**StorageClassEnum**](StorageClassEnum.md) | Backend storage class for domain.* &#x60;pulpcore.app.models.storage.FileSystem&#x60; - Use local filesystem as storage* &#x60;storages.backends.s3boto3.S3Boto3Storage&#x60; - Use Amazon S3 as storage* &#x60;storages.backends.azure_storage.AzureStorage&#x60; - Use Azure Blob as storage* &#x60;pulp_service.app.storage.OCIStorage&#x60; - Use OCI as storage | 
**StorageSettings** | **map[string]interface{}** | Settings for storage class. | 
**RedirectToObjectStorage** | Pointer to **bool** | Boolean to have the content app redirect to object storage. | [optional] [default to true]
**HideGuardedDistributions** | Pointer to **bool** | Boolean to hide distributions with a content guard in the content app. | [optional] [default to false]

## Methods

### NewDomainResponse

`func NewDomainResponse(name string, storageClass StorageClassEnum, storageSettings map[string]interface{}, ) *DomainResponse`

NewDomainResponse instantiates a new DomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithDefaults

`func NewDomainResponseWithDefaults() *DomainResponse`

NewDomainResponseWithDefaults instantiates a new DomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DomainResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DomainResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DomainResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DomainResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *DomainResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *DomainResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *DomainResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *DomainResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DomainResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DomainResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DomainResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DomainResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *DomainResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *DomainResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *DomainResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *DomainResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *DomainResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DomainResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DomainResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DomainResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStorageClass

`func (o *DomainResponse) GetStorageClass() StorageClassEnum`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *DomainResponse) GetStorageClassOk() (*StorageClassEnum, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *DomainResponse) SetStorageClass(v StorageClassEnum)`

SetStorageClass sets StorageClass field to given value.


### GetStorageSettings

`func (o *DomainResponse) GetStorageSettings() map[string]interface{}`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *DomainResponse) GetStorageSettingsOk() (*map[string]interface{}, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *DomainResponse) SetStorageSettings(v map[string]interface{})`

SetStorageSettings sets StorageSettings field to given value.


### GetRedirectToObjectStorage

`func (o *DomainResponse) GetRedirectToObjectStorage() bool`

GetRedirectToObjectStorage returns the RedirectToObjectStorage field if non-nil, zero value otherwise.

### GetRedirectToObjectStorageOk

`func (o *DomainResponse) GetRedirectToObjectStorageOk() (*bool, bool)`

GetRedirectToObjectStorageOk returns a tuple with the RedirectToObjectStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToObjectStorage

`func (o *DomainResponse) SetRedirectToObjectStorage(v bool)`

SetRedirectToObjectStorage sets RedirectToObjectStorage field to given value.

### HasRedirectToObjectStorage

`func (o *DomainResponse) HasRedirectToObjectStorage() bool`

HasRedirectToObjectStorage returns a boolean if a field has been set.

### GetHideGuardedDistributions

`func (o *DomainResponse) GetHideGuardedDistributions() bool`

GetHideGuardedDistributions returns the HideGuardedDistributions field if non-nil, zero value otherwise.

### GetHideGuardedDistributionsOk

`func (o *DomainResponse) GetHideGuardedDistributionsOk() (*bool, bool)`

GetHideGuardedDistributionsOk returns a tuple with the HideGuardedDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideGuardedDistributions

`func (o *DomainResponse) SetHideGuardedDistributions(v bool)`

SetHideGuardedDistributions sets HideGuardedDistributions field to given value.

### HasHideGuardedDistributions

`func (o *DomainResponse) HasHideGuardedDistributions() bool`

HasHideGuardedDistributions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


