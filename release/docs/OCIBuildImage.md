# OCIBuildImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerfileName** | Pointer to **string** | Name of the Containerfile, from build_context, that should be used to run podman-build. | [optional] 
**Containerfile** | Pointer to ***os.File** | An uploaded Containerfile that should be used to run podman-build. | [optional] 
**Tag** | Pointer to **string** | A tag name for the new image being built. | [optional] [default to "latest"]
**BuildContext** | Pointer to **NullableString** | RepositoryVersion to be used as the build context for container images. | [optional] 

## Methods

### NewOCIBuildImage

`func NewOCIBuildImage() *OCIBuildImage`

NewOCIBuildImage instantiates a new OCIBuildImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCIBuildImageWithDefaults

`func NewOCIBuildImageWithDefaults() *OCIBuildImage`

NewOCIBuildImageWithDefaults instantiates a new OCIBuildImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerfileName

`func (o *OCIBuildImage) GetContainerfileName() string`

GetContainerfileName returns the ContainerfileName field if non-nil, zero value otherwise.

### GetContainerfileNameOk

`func (o *OCIBuildImage) GetContainerfileNameOk() (*string, bool)`

GetContainerfileNameOk returns a tuple with the ContainerfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerfileName

`func (o *OCIBuildImage) SetContainerfileName(v string)`

SetContainerfileName sets ContainerfileName field to given value.

### HasContainerfileName

`func (o *OCIBuildImage) HasContainerfileName() bool`

HasContainerfileName returns a boolean if a field has been set.

### GetContainerfile

`func (o *OCIBuildImage) GetContainerfile() *os.File`

GetContainerfile returns the Containerfile field if non-nil, zero value otherwise.

### GetContainerfileOk

`func (o *OCIBuildImage) GetContainerfileOk() (**os.File, bool)`

GetContainerfileOk returns a tuple with the Containerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerfile

`func (o *OCIBuildImage) SetContainerfile(v *os.File)`

SetContainerfile sets Containerfile field to given value.

### HasContainerfile

`func (o *OCIBuildImage) HasContainerfile() bool`

HasContainerfile returns a boolean if a field has been set.

### GetTag

`func (o *OCIBuildImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OCIBuildImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OCIBuildImage) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *OCIBuildImage) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetBuildContext

`func (o *OCIBuildImage) GetBuildContext() string`

GetBuildContext returns the BuildContext field if non-nil, zero value otherwise.

### GetBuildContextOk

`func (o *OCIBuildImage) GetBuildContextOk() (*string, bool)`

GetBuildContextOk returns a tuple with the BuildContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildContext

`func (o *OCIBuildImage) SetBuildContext(v string)`

SetBuildContext sets BuildContext field to given value.

### HasBuildContext

`func (o *OCIBuildImage) HasBuildContext() bool`

HasBuildContext returns a boolean if a field has been set.

### SetBuildContextNil

`func (o *OCIBuildImage) SetBuildContextNil(b bool)`

 SetBuildContextNil sets the value for BuildContext to be an explicit nil

### UnsetBuildContext
`func (o *OCIBuildImage) UnsetBuildContext()`

UnsetBuildContext ensures that no value is present for BuildContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


