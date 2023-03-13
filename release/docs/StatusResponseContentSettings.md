# StatusResponseContentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentOrigin** | **string** | The CONTENT_ORIGIN setting for this Pulp instance | 
**ContentPathPrefix** | **string** | The CONTENT_PATH_PREFIX setting for this Pulp instance | 

## Methods

### NewStatusResponseContentSettings

`func NewStatusResponseContentSettings(contentOrigin string, contentPathPrefix string, ) *StatusResponseContentSettings`

NewStatusResponseContentSettings instantiates a new StatusResponseContentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseContentSettingsWithDefaults

`func NewStatusResponseContentSettingsWithDefaults() *StatusResponseContentSettings`

NewStatusResponseContentSettingsWithDefaults instantiates a new StatusResponseContentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentOrigin

`func (o *StatusResponseContentSettings) GetContentOrigin() string`

GetContentOrigin returns the ContentOrigin field if non-nil, zero value otherwise.

### GetContentOriginOk

`func (o *StatusResponseContentSettings) GetContentOriginOk() (*string, bool)`

GetContentOriginOk returns a tuple with the ContentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentOrigin

`func (o *StatusResponseContentSettings) SetContentOrigin(v string)`

SetContentOrigin sets ContentOrigin field to given value.


### GetContentPathPrefix

`func (o *StatusResponseContentSettings) GetContentPathPrefix() string`

GetContentPathPrefix returns the ContentPathPrefix field if non-nil, zero value otherwise.

### GetContentPathPrefixOk

`func (o *StatusResponseContentSettings) GetContentPathPrefixOk() (*string, bool)`

GetContentPathPrefixOk returns a tuple with the ContentPathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPathPrefix

`func (o *StatusResponseContentSettings) SetContentPathPrefix(v string)`

SetContentPathPrefix sets ContentPathPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


