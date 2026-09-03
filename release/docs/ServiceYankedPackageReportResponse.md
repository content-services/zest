# ServiceYankedPackageReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Report** | **interface{}** |  | 
**RepositoryName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceYankedPackageReportResponse

`func NewServiceYankedPackageReportResponse(report interface{}, ) *ServiceYankedPackageReportResponse`

NewServiceYankedPackageReportResponse instantiates a new ServiceYankedPackageReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceYankedPackageReportResponseWithDefaults

`func NewServiceYankedPackageReportResponseWithDefaults() *ServiceYankedPackageReportResponse`

NewServiceYankedPackageReportResponseWithDefaults instantiates a new ServiceYankedPackageReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpCreated

`func (o *ServiceYankedPackageReportResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ServiceYankedPackageReportResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ServiceYankedPackageReportResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ServiceYankedPackageReportResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *ServiceYankedPackageReportResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *ServiceYankedPackageReportResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *ServiceYankedPackageReportResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *ServiceYankedPackageReportResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetReport

`func (o *ServiceYankedPackageReportResponse) GetReport() interface{}`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ServiceYankedPackageReportResponse) GetReportOk() (*interface{}, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ServiceYankedPackageReportResponse) SetReport(v interface{})`

SetReport sets Report field to given value.


### SetReportNil

`func (o *ServiceYankedPackageReportResponse) SetReportNil(b bool)`

 SetReportNil sets the value for Report to be an explicit nil

### UnsetReport
`func (o *ServiceYankedPackageReportResponse) UnsetReport()`

UnsetReport ensures that no value is present for Report, not even an explicit nil
### GetRepositoryName

`func (o *ServiceYankedPackageReportResponse) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *ServiceYankedPackageReportResponse) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *ServiceYankedPackageReportResponse) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *ServiceYankedPackageReportResponse) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### SetRepositoryNameNil

`func (o *ServiceYankedPackageReportResponse) SetRepositoryNameNil(b bool)`

 SetRepositoryNameNil sets the value for RepositoryName to be an explicit nil

### UnsetRepositoryName
`func (o *ServiceYankedPackageReportResponse) UnsetRepositoryName()`

UnsetRepositoryName ensures that no value is present for RepositoryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


