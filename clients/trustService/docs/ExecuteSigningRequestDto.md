# ExecuteSigningRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** |  | 
**ProviderMode** | Pointer to **NullableString** |  | [optional] 
**SigningProfileId** | Pointer to **NullableString** |  | [optional] 
**SigningCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExecuteSigningRequestDto

`func NewExecuteSigningRequestDto(providerName string, ) *ExecuteSigningRequestDto`

NewExecuteSigningRequestDto instantiates a new ExecuteSigningRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteSigningRequestDtoWithDefaults

`func NewExecuteSigningRequestDtoWithDefaults() *ExecuteSigningRequestDto`

NewExecuteSigningRequestDtoWithDefaults instantiates a new ExecuteSigningRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *ExecuteSigningRequestDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ExecuteSigningRequestDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ExecuteSigningRequestDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetProviderMode

`func (o *ExecuteSigningRequestDto) GetProviderMode() string`

GetProviderMode returns the ProviderMode field if non-nil, zero value otherwise.

### GetProviderModeOk

`func (o *ExecuteSigningRequestDto) GetProviderModeOk() (*string, bool)`

GetProviderModeOk returns a tuple with the ProviderMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMode

`func (o *ExecuteSigningRequestDto) SetProviderMode(v string)`

SetProviderMode sets ProviderMode field to given value.

### HasProviderMode

`func (o *ExecuteSigningRequestDto) HasProviderMode() bool`

HasProviderMode returns a boolean if a field has been set.

### SetProviderModeNil

`func (o *ExecuteSigningRequestDto) SetProviderModeNil(b bool)`

 SetProviderModeNil sets the value for ProviderMode to be an explicit nil

### UnsetProviderMode
`func (o *ExecuteSigningRequestDto) UnsetProviderMode()`

UnsetProviderMode ensures that no value is present for ProviderMode, not even an explicit nil
### GetSigningProfileId

`func (o *ExecuteSigningRequestDto) GetSigningProfileId() string`

GetSigningProfileId returns the SigningProfileId field if non-nil, zero value otherwise.

### GetSigningProfileIdOk

`func (o *ExecuteSigningRequestDto) GetSigningProfileIdOk() (*string, bool)`

GetSigningProfileIdOk returns a tuple with the SigningProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileId

`func (o *ExecuteSigningRequestDto) SetSigningProfileId(v string)`

SetSigningProfileId sets SigningProfileId field to given value.

### HasSigningProfileId

`func (o *ExecuteSigningRequestDto) HasSigningProfileId() bool`

HasSigningProfileId returns a boolean if a field has been set.

### SetSigningProfileIdNil

`func (o *ExecuteSigningRequestDto) SetSigningProfileIdNil(b bool)`

 SetSigningProfileIdNil sets the value for SigningProfileId to be an explicit nil

### UnsetSigningProfileId
`func (o *ExecuteSigningRequestDto) UnsetSigningProfileId()`

UnsetSigningProfileId ensures that no value is present for SigningProfileId, not even an explicit nil
### GetSigningCertificateId

`func (o *ExecuteSigningRequestDto) GetSigningCertificateId() string`

GetSigningCertificateId returns the SigningCertificateId field if non-nil, zero value otherwise.

### GetSigningCertificateIdOk

`func (o *ExecuteSigningRequestDto) GetSigningCertificateIdOk() (*string, bool)`

GetSigningCertificateIdOk returns a tuple with the SigningCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificateId

`func (o *ExecuteSigningRequestDto) SetSigningCertificateId(v string)`

SetSigningCertificateId sets SigningCertificateId field to given value.

### HasSigningCertificateId

`func (o *ExecuteSigningRequestDto) HasSigningCertificateId() bool`

HasSigningCertificateId returns a boolean if a field has been set.

### SetSigningCertificateIdNil

`func (o *ExecuteSigningRequestDto) SetSigningCertificateIdNil(b bool)`

 SetSigningCertificateIdNil sets the value for SigningCertificateId to be an explicit nil

### UnsetSigningCertificateId
`func (o *ExecuteSigningRequestDto) UnsetSigningCertificateId()`

UnsetSigningCertificateId ensures that no value is present for SigningCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


