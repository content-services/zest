# ContentSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentOrigin** | Pointer to **NullableString** | The CONTENT_ORIGIN setting for this Pulp instance | [optional] 
**ContentPathPrefix** | **string** | The CONTENT_PATH_PREFIX setting for this Pulp instance | 

## Methods

### NewContentSettingsResponse

`func NewContentSettingsResponse(contentPathPrefix string, ) *ContentSettingsResponse`

NewContentSettingsResponse instantiates a new ContentSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSettingsResponseWithDefaults

`func NewContentSettingsResponseWithDefaults() *ContentSettingsResponse`

NewContentSettingsResponseWithDefaults instantiates a new ContentSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentOrigin

`func (o *ContentSettingsResponse) GetContentOrigin() string`

GetContentOrigin returns the ContentOrigin field if non-nil, zero value otherwise.

### GetContentOriginOk

`func (o *ContentSettingsResponse) GetContentOriginOk() (*string, bool)`

GetContentOriginOk returns a tuple with the ContentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentOrigin

`func (o *ContentSettingsResponse) SetContentOrigin(v string)`

SetContentOrigin sets ContentOrigin field to given value.

### HasContentOrigin

`func (o *ContentSettingsResponse) HasContentOrigin() bool`

HasContentOrigin returns a boolean if a field has been set.

### SetContentOriginNil

`func (o *ContentSettingsResponse) SetContentOriginNil(b bool)`

 SetContentOriginNil sets the value for ContentOrigin to be an explicit nil

### UnsetContentOrigin
`func (o *ContentSettingsResponse) UnsetContentOrigin()`

UnsetContentOrigin ensures that no value is present for ContentOrigin, not even an explicit nil
### GetContentPathPrefix

`func (o *ContentSettingsResponse) GetContentPathPrefix() string`

GetContentPathPrefix returns the ContentPathPrefix field if non-nil, zero value otherwise.

### GetContentPathPrefixOk

`func (o *ContentSettingsResponse) GetContentPathPrefixOk() (*string, bool)`

GetContentPathPrefixOk returns a tuple with the ContentPathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPathPrefix

`func (o *ContentSettingsResponse) SetContentPathPrefix(v string)`

SetContentPathPrefix sets ContentPathPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


