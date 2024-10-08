# WorkerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the worker. | [optional] [readonly] 
**LastHeartbeat** | Pointer to **time.Time** | Timestamp of the last time the worker talked to the service. | [optional] [readonly] 
**Versions** | Pointer to **map[string]string** | Versions of the components installed. | [optional] [readonly] 
**CurrentTask** | Pointer to **string** | The task this worker is currently executing, or empty if the worker is not currently assigned to a task. | [optional] [readonly] 

## Methods

### NewWorkerResponse

`func NewWorkerResponse() *WorkerResponse`

NewWorkerResponse instantiates a new WorkerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerResponseWithDefaults

`func NewWorkerResponseWithDefaults() *WorkerResponse`

NewWorkerResponseWithDefaults instantiates a new WorkerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *WorkerResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *WorkerResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *WorkerResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *WorkerResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *WorkerResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *WorkerResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *WorkerResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *WorkerResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *WorkerResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *WorkerResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *WorkerResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *WorkerResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *WorkerResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *WorkerResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *WorkerResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *WorkerResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *WorkerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkerResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkerResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *WorkerResponse) GetLastHeartbeat() time.Time`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *WorkerResponse) GetLastHeartbeatOk() (*time.Time, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *WorkerResponse) SetLastHeartbeat(v time.Time)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *WorkerResponse) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### GetVersions

`func (o *WorkerResponse) GetVersions() map[string]string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *WorkerResponse) GetVersionsOk() (*map[string]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *WorkerResponse) SetVersions(v map[string]string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *WorkerResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetCurrentTask

`func (o *WorkerResponse) GetCurrentTask() string`

GetCurrentTask returns the CurrentTask field if non-nil, zero value otherwise.

### GetCurrentTaskOk

`func (o *WorkerResponse) GetCurrentTaskOk() (*string, bool)`

GetCurrentTaskOk returns a tuple with the CurrentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTask

`func (o *WorkerResponse) SetCurrentTask(v string)`

SetCurrentTask sets CurrentTask field to given value.

### HasCurrentTask

`func (o *WorkerResponse) HasCurrentTask() bool`

HasCurrentTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


