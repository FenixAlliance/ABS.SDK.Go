# VoidSigningRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoidedReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoidSigningRequestDto

`func NewVoidSigningRequestDto() *VoidSigningRequestDto`

NewVoidSigningRequestDto instantiates a new VoidSigningRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidSigningRequestDtoWithDefaults

`func NewVoidSigningRequestDtoWithDefaults() *VoidSigningRequestDto`

NewVoidSigningRequestDtoWithDefaults instantiates a new VoidSigningRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoidedReason

`func (o *VoidSigningRequestDto) GetVoidedReason() string`

GetVoidedReason returns the VoidedReason field if non-nil, zero value otherwise.

### GetVoidedReasonOk

`func (o *VoidSigningRequestDto) GetVoidedReasonOk() (*string, bool)`

GetVoidedReasonOk returns a tuple with the VoidedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedReason

`func (o *VoidSigningRequestDto) SetVoidedReason(v string)`

SetVoidedReason sets VoidedReason field to given value.

### HasVoidedReason

`func (o *VoidSigningRequestDto) HasVoidedReason() bool`

HasVoidedReason returns a boolean if a field has been set.

### SetVoidedReasonNil

`func (o *VoidSigningRequestDto) SetVoidedReasonNil(b bool)`

 SetVoidedReasonNil sets the value for VoidedReason to be an explicit nil

### UnsetVoidedReason
`func (o *VoidSigningRequestDto) UnsetVoidedReason()`

UnsetVoidedReason ensures that no value is present for VoidedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


