# ContentViewUpdateReferenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Reference URL | 
**RefId** | **NullableString** | ID of the reference | 
**Title** | **NullableString** | Title of the reference | 
**RefType** | **string** | Type of the reference, e.g. &#39;cve&#39; | 

## Methods

### NewContentViewUpdateReferenceResponse

`func NewContentViewUpdateReferenceResponse(href string, refId NullableString, title NullableString, refType string, ) *ContentViewUpdateReferenceResponse`

NewContentViewUpdateReferenceResponse instantiates a new ContentViewUpdateReferenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentViewUpdateReferenceResponseWithDefaults

`func NewContentViewUpdateReferenceResponseWithDefaults() *ContentViewUpdateReferenceResponse`

NewContentViewUpdateReferenceResponseWithDefaults instantiates a new ContentViewUpdateReferenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ContentViewUpdateReferenceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ContentViewUpdateReferenceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ContentViewUpdateReferenceResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetRefId

`func (o *ContentViewUpdateReferenceResponse) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ContentViewUpdateReferenceResponse) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ContentViewUpdateReferenceResponse) SetRefId(v string)`

SetRefId sets RefId field to given value.


### SetRefIdNil

`func (o *ContentViewUpdateReferenceResponse) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ContentViewUpdateReferenceResponse) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetTitle

`func (o *ContentViewUpdateReferenceResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentViewUpdateReferenceResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentViewUpdateReferenceResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ContentViewUpdateReferenceResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ContentViewUpdateReferenceResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetRefType

`func (o *ContentViewUpdateReferenceResponse) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ContentViewUpdateReferenceResponse) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ContentViewUpdateReferenceResponse) SetRefType(v string)`

SetRefType sets RefType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


