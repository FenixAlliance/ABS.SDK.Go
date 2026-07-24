# BusinessRelationshipUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildTenantId** | Pointer to **NullableString** |  | [optional] 
**OwnershipPercentage** | Pointer to **float64** |  | [optional] 

## Methods

### NewBusinessRelationshipUpdateDto

`func NewBusinessRelationshipUpdateDto() *BusinessRelationshipUpdateDto`

NewBusinessRelationshipUpdateDto instantiates a new BusinessRelationshipUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessRelationshipUpdateDtoWithDefaults

`func NewBusinessRelationshipUpdateDtoWithDefaults() *BusinessRelationshipUpdateDto`

NewBusinessRelationshipUpdateDtoWithDefaults instantiates a new BusinessRelationshipUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildTenantId

`func (o *BusinessRelationshipUpdateDto) GetChildTenantId() string`

GetChildTenantId returns the ChildTenantId field if non-nil, zero value otherwise.

### GetChildTenantIdOk

`func (o *BusinessRelationshipUpdateDto) GetChildTenantIdOk() (*string, bool)`

GetChildTenantIdOk returns a tuple with the ChildTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTenantId

`func (o *BusinessRelationshipUpdateDto) SetChildTenantId(v string)`

SetChildTenantId sets ChildTenantId field to given value.

### HasChildTenantId

`func (o *BusinessRelationshipUpdateDto) HasChildTenantId() bool`

HasChildTenantId returns a boolean if a field has been set.

### SetChildTenantIdNil

`func (o *BusinessRelationshipUpdateDto) SetChildTenantIdNil(b bool)`

 SetChildTenantIdNil sets the value for ChildTenantId to be an explicit nil

### UnsetChildTenantId
`func (o *BusinessRelationshipUpdateDto) UnsetChildTenantId()`

UnsetChildTenantId ensures that no value is present for ChildTenantId, not even an explicit nil
### GetOwnershipPercentage

`func (o *BusinessRelationshipUpdateDto) GetOwnershipPercentage() float64`

GetOwnershipPercentage returns the OwnershipPercentage field if non-nil, zero value otherwise.

### GetOwnershipPercentageOk

`func (o *BusinessRelationshipUpdateDto) GetOwnershipPercentageOk() (*float64, bool)`

GetOwnershipPercentageOk returns a tuple with the OwnershipPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipPercentage

`func (o *BusinessRelationshipUpdateDto) SetOwnershipPercentage(v float64)`

SetOwnershipPercentage sets OwnershipPercentage field to given value.

### HasOwnershipPercentage

`func (o *BusinessRelationshipUpdateDto) HasOwnershipPercentage() bool`

HasOwnershipPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


