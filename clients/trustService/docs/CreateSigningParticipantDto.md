# CreateSigningParticipantDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] 
**RoutingOrder** | Pointer to **int32** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateSigningParticipantDto

`func NewCreateSigningParticipantDto(contactId string, ) *CreateSigningParticipantDto`

NewCreateSigningParticipantDto instantiates a new CreateSigningParticipantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSigningParticipantDtoWithDefaults

`func NewCreateSigningParticipantDtoWithDefaults() *CreateSigningParticipantDto`

NewCreateSigningParticipantDtoWithDefaults instantiates a new CreateSigningParticipantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *CreateSigningParticipantDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *CreateSigningParticipantDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *CreateSigningParticipantDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetRole

`func (o *CreateSigningParticipantDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateSigningParticipantDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateSigningParticipantDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateSigningParticipantDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoutingOrder

`func (o *CreateSigningParticipantDto) GetRoutingOrder() int32`

GetRoutingOrder returns the RoutingOrder field if non-nil, zero value otherwise.

### GetRoutingOrderOk

`func (o *CreateSigningParticipantDto) GetRoutingOrderOk() (*int32, bool)`

GetRoutingOrderOk returns a tuple with the RoutingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingOrder

`func (o *CreateSigningParticipantDto) SetRoutingOrder(v int32)`

SetRoutingOrder sets RoutingOrder field to given value.

### HasRoutingOrder

`func (o *CreateSigningParticipantDto) HasRoutingOrder() bool`

HasRoutingOrder returns a boolean if a field has been set.

### GetExternalReference

`func (o *CreateSigningParticipantDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *CreateSigningParticipantDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *CreateSigningParticipantDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *CreateSigningParticipantDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *CreateSigningParticipantDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *CreateSigningParticipantDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


