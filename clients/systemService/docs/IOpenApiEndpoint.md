# IOpenApiEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIOpenApiEndpoint

`func NewIOpenApiEndpoint() *IOpenApiEndpoint`

NewIOpenApiEndpoint instantiates a new IOpenApiEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIOpenApiEndpointWithDefaults

`func NewIOpenApiEndpointWithDefaults() *IOpenApiEndpoint`

NewIOpenApiEndpointWithDefaults instantiates a new IOpenApiEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IOpenApiEndpoint) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IOpenApiEndpoint) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IOpenApiEndpoint) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IOpenApiEndpoint) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetName

`func (o *IOpenApiEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IOpenApiEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IOpenApiEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IOpenApiEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IOpenApiEndpoint) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IOpenApiEndpoint) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *IOpenApiEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IOpenApiEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IOpenApiEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IOpenApiEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IOpenApiEndpoint) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IOpenApiEndpoint) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


