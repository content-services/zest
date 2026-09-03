# PythonPackageYankResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**VulnReport** | Pointer to **string** |  | [optional] [readonly] 
**NameNormalized** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**YankedReason** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPythonPackageYankResponse

`func NewPythonPackageYankResponse() *PythonPackageYankResponse`

NewPythonPackageYankResponse instantiates a new PythonPackageYankResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPackageYankResponseWithDefaults

`func NewPythonPackageYankResponseWithDefaults() *PythonPackageYankResponse`

NewPythonPackageYankResponseWithDefaults instantiates a new PythonPackageYankResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PythonPackageYankResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PythonPackageYankResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PythonPackageYankResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PythonPackageYankResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *PythonPackageYankResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *PythonPackageYankResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *PythonPackageYankResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *PythonPackageYankResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PythonPackageYankResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PythonPackageYankResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PythonPackageYankResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PythonPackageYankResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *PythonPackageYankResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *PythonPackageYankResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *PythonPackageYankResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *PythonPackageYankResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PythonPackageYankResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PythonPackageYankResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PythonPackageYankResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PythonPackageYankResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetVulnReport

`func (o *PythonPackageYankResponse) GetVulnReport() string`

GetVulnReport returns the VulnReport field if non-nil, zero value otherwise.

### GetVulnReportOk

`func (o *PythonPackageYankResponse) GetVulnReportOk() (*string, bool)`

GetVulnReportOk returns a tuple with the VulnReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnReport

`func (o *PythonPackageYankResponse) SetVulnReport(v string)`

SetVulnReport sets VulnReport field to given value.

### HasVulnReport

`func (o *PythonPackageYankResponse) HasVulnReport() bool`

HasVulnReport returns a boolean if a field has been set.

### GetNameNormalized

`func (o *PythonPackageYankResponse) GetNameNormalized() string`

GetNameNormalized returns the NameNormalized field if non-nil, zero value otherwise.

### GetNameNormalizedOk

`func (o *PythonPackageYankResponse) GetNameNormalizedOk() (*string, bool)`

GetNameNormalizedOk returns a tuple with the NameNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameNormalized

`func (o *PythonPackageYankResponse) SetNameNormalized(v string)`

SetNameNormalized sets NameNormalized field to given value.

### HasNameNormalized

`func (o *PythonPackageYankResponse) HasNameNormalized() bool`

HasNameNormalized returns a boolean if a field has been set.

### GetVersion

`func (o *PythonPackageYankResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PythonPackageYankResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PythonPackageYankResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PythonPackageYankResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetYankedReason

`func (o *PythonPackageYankResponse) GetYankedReason() string`

GetYankedReason returns the YankedReason field if non-nil, zero value otherwise.

### GetYankedReasonOk

`func (o *PythonPackageYankResponse) GetYankedReasonOk() (*string, bool)`

GetYankedReasonOk returns a tuple with the YankedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYankedReason

`func (o *PythonPackageYankResponse) SetYankedReason(v string)`

SetYankedReason sets YankedReason field to given value.

### HasYankedReason

`func (o *PythonPackageYankResponse) HasYankedReason() bool`

HasYankedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


