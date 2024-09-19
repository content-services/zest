# DomainBackendMigrator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageClass** | [**StorageClassEnum**](StorageClassEnum.md) | The new backend storage class to migrate to.* &#x60;pulpcore.app.models.storage.FileSystem&#x60; - Use local filesystem as storage* &#x60;storages.backends.s3boto3.S3Boto3Storage&#x60; - Use Amazon S3 as storage* &#x60;storages.backends.azure_storage.AzureStorage&#x60; - Use Azure Blob as storage | 
**StorageSettings** | **map[string]interface{}** | The settings for the new storage class to migrate to. | 

## Methods

### NewDomainBackendMigrator

`func NewDomainBackendMigrator(storageClass StorageClassEnum, storageSettings map[string]interface{}, ) *DomainBackendMigrator`

NewDomainBackendMigrator instantiates a new DomainBackendMigrator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBackendMigratorWithDefaults

`func NewDomainBackendMigratorWithDefaults() *DomainBackendMigrator`

NewDomainBackendMigratorWithDefaults instantiates a new DomainBackendMigrator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageClass

`func (o *DomainBackendMigrator) GetStorageClass() StorageClassEnum`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *DomainBackendMigrator) GetStorageClassOk() (*StorageClassEnum, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *DomainBackendMigrator) SetStorageClass(v StorageClassEnum)`

SetStorageClass sets StorageClass field to given value.


### GetStorageSettings

`func (o *DomainBackendMigrator) GetStorageSettings() map[string]interface{}`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *DomainBackendMigrator) GetStorageSettingsOk() (*map[string]interface{}, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *DomainBackendMigrator) SetStorageSettings(v map[string]interface{})`

SetStorageSettings sets StorageSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


