# PulpExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Task** | Pointer to **NullableString** | A URI of the task that ran the Export. | [optional] 
**ExportedResources** | Pointer to **[]string** | Resources that were exported. | [optional] [readonly] 
**Params** | Pointer to **interface{}** | Any additional parameters that were used to create the export. | [optional] [readonly] 
**OutputFileInfo** | Pointer to **interface{}** | Dictionary of filename: sha256hash entries for export-output-file(s) | [optional] [readonly] 
**TocInfo** | Pointer to **interface{}** | Filename and sha256-checksum of table-of-contents for this export | [optional] [readonly] 

## Methods

### NewPulpExportResponse

`func NewPulpExportResponse() *PulpExportResponse`

NewPulpExportResponse instantiates a new PulpExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpExportResponseWithDefaults

`func NewPulpExportResponseWithDefaults() *PulpExportResponse`

NewPulpExportResponseWithDefaults instantiates a new PulpExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PulpExportResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PulpExportResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PulpExportResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PulpExportResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PulpExportResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PulpExportResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PulpExportResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PulpExportResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *PulpExportResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *PulpExportResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *PulpExportResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *PulpExportResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetTask

`func (o *PulpExportResponse) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PulpExportResponse) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PulpExportResponse) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *PulpExportResponse) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *PulpExportResponse) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *PulpExportResponse) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetExportedResources

`func (o *PulpExportResponse) GetExportedResources() []string`

GetExportedResources returns the ExportedResources field if non-nil, zero value otherwise.

### GetExportedResourcesOk

`func (o *PulpExportResponse) GetExportedResourcesOk() (*[]string, bool)`

GetExportedResourcesOk returns a tuple with the ExportedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedResources

`func (o *PulpExportResponse) SetExportedResources(v []string)`

SetExportedResources sets ExportedResources field to given value.

### HasExportedResources

`func (o *PulpExportResponse) HasExportedResources() bool`

HasExportedResources returns a boolean if a field has been set.

### GetParams

`func (o *PulpExportResponse) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PulpExportResponse) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PulpExportResponse) SetParams(v interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *PulpExportResponse) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *PulpExportResponse) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *PulpExportResponse) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetOutputFileInfo

`func (o *PulpExportResponse) GetOutputFileInfo() interface{}`

GetOutputFileInfo returns the OutputFileInfo field if non-nil, zero value otherwise.

### GetOutputFileInfoOk

`func (o *PulpExportResponse) GetOutputFileInfoOk() (*interface{}, bool)`

GetOutputFileInfoOk returns a tuple with the OutputFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileInfo

`func (o *PulpExportResponse) SetOutputFileInfo(v interface{})`

SetOutputFileInfo sets OutputFileInfo field to given value.

### HasOutputFileInfo

`func (o *PulpExportResponse) HasOutputFileInfo() bool`

HasOutputFileInfo returns a boolean if a field has been set.

### SetOutputFileInfoNil

`func (o *PulpExportResponse) SetOutputFileInfoNil(b bool)`

 SetOutputFileInfoNil sets the value for OutputFileInfo to be an explicit nil

### UnsetOutputFileInfo
`func (o *PulpExportResponse) UnsetOutputFileInfo()`

UnsetOutputFileInfo ensures that no value is present for OutputFileInfo, not even an explicit nil
### GetTocInfo

`func (o *PulpExportResponse) GetTocInfo() interface{}`

GetTocInfo returns the TocInfo field if non-nil, zero value otherwise.

### GetTocInfoOk

`func (o *PulpExportResponse) GetTocInfoOk() (*interface{}, bool)`

GetTocInfoOk returns a tuple with the TocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTocInfo

`func (o *PulpExportResponse) SetTocInfo(v interface{})`

SetTocInfo sets TocInfo field to given value.

### HasTocInfo

`func (o *PulpExportResponse) HasTocInfo() bool`

HasTocInfo returns a boolean if a field has been set.

### SetTocInfoNil

`func (o *PulpExportResponse) SetTocInfoNil(b bool)`

 SetTocInfoNil sets the value for TocInfo to be an explicit nil

### UnsetTocInfo
`func (o *PulpExportResponse) UnsetTocInfo()`

UnsetTocInfo ensures that no value is present for TocInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


