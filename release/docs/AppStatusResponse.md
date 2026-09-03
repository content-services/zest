# AppStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the worker. | [optional] [readonly] 
**LastHeartbeat** | Pointer to **time.Time** | Timestamp of the last time the worker talked to the service. | [optional] [readonly] 
**Versions** | Pointer to **map[string]string** | Versions of the components installed. | [optional] [readonly] 

## Methods

### NewAppStatusResponse

`func NewAppStatusResponse() *AppStatusResponse`

NewAppStatusResponse instantiates a new AppStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStatusResponseWithDefaults

`func NewAppStatusResponseWithDefaults() *AppStatusResponse`

NewAppStatusResponseWithDefaults instantiates a new AppStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppStatusResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppStatusResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *AppStatusResponse) GetLastHeartbeat() time.Time`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *AppStatusResponse) GetLastHeartbeatOk() (*time.Time, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *AppStatusResponse) SetLastHeartbeat(v time.Time)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *AppStatusResponse) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### GetVersions

`func (o *AppStatusResponse) GetVersions() map[string]*string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AppStatusResponse) GetVersionsOk() (*map[string]*string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AppStatusResponse) SetVersions(v map[string]*string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AppStatusResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


