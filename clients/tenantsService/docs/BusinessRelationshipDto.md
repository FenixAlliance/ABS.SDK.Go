# BusinessRelationshipDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ParentTenantId** | Pointer to **NullableString** |  | [optional] 
**ChildTenantId** | Pointer to **NullableString** |  | [optional] 
**OwnershipPercentage** | Pointer to **float64** |  | [optional] 

## Methods

### NewBusinessRelationshipDto

`func NewBusinessRelationshipDto() *BusinessRelationshipDto`

NewBusinessRelationshipDto instantiates a new BusinessRelationshipDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessRelationshipDtoWithDefaults

`func NewBusinessRelationshipDtoWithDefaults() *BusinessRelationshipDto`

NewBusinessRelationshipDtoWithDefaults instantiates a new BusinessRelationshipDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessRelationshipDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessRelationshipDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessRelationshipDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessRelationshipDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BusinessRelationshipDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BusinessRelationshipDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BusinessRelationshipDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BusinessRelationshipDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BusinessRelationshipDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BusinessRelationshipDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BusinessRelationshipDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BusinessRelationshipDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetParentTenantId

`func (o *BusinessRelationshipDto) GetParentTenantId() string`

GetParentTenantId returns the ParentTenantId field if non-nil, zero value otherwise.

### GetParentTenantIdOk

`func (o *BusinessRelationshipDto) GetParentTenantIdOk() (*string, bool)`

GetParentTenantIdOk returns a tuple with the ParentTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTenantId

`func (o *BusinessRelationshipDto) SetParentTenantId(v string)`

SetParentTenantId sets ParentTenantId field to given value.

### HasParentTenantId

`func (o *BusinessRelationshipDto) HasParentTenantId() bool`

HasParentTenantId returns a boolean if a field has been set.

### SetParentTenantIdNil

`func (o *BusinessRelationshipDto) SetParentTenantIdNil(b bool)`

 SetParentTenantIdNil sets the value for ParentTenantId to be an explicit nil

### UnsetParentTenantId
`func (o *BusinessRelationshipDto) UnsetParentTenantId()`

UnsetParentTenantId ensures that no value is present for ParentTenantId, not even an explicit nil
### GetChildTenantId

`func (o *BusinessRelationshipDto) GetChildTenantId() string`

GetChildTenantId returns the ChildTenantId field if non-nil, zero value otherwise.

### GetChildTenantIdOk

`func (o *BusinessRelationshipDto) GetChildTenantIdOk() (*string, bool)`

GetChildTenantIdOk returns a tuple with the ChildTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTenantId

`func (o *BusinessRelationshipDto) SetChildTenantId(v string)`

SetChildTenantId sets ChildTenantId field to given value.

### HasChildTenantId

`func (o *BusinessRelationshipDto) HasChildTenantId() bool`

HasChildTenantId returns a boolean if a field has been set.

### SetChildTenantIdNil

`func (o *BusinessRelationshipDto) SetChildTenantIdNil(b bool)`

 SetChildTenantIdNil sets the value for ChildTenantId to be an explicit nil

### UnsetChildTenantId
`func (o *BusinessRelationshipDto) UnsetChildTenantId()`

UnsetChildTenantId ensures that no value is present for ChildTenantId, not even an explicit nil
### GetOwnershipPercentage

`func (o *BusinessRelationshipDto) GetOwnershipPercentage() float64`

GetOwnershipPercentage returns the OwnershipPercentage field if non-nil, zero value otherwise.

### GetOwnershipPercentageOk

`func (o *BusinessRelationshipDto) GetOwnershipPercentageOk() (*float64, bool)`

GetOwnershipPercentageOk returns a tuple with the OwnershipPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipPercentage

`func (o *BusinessRelationshipDto) SetOwnershipPercentage(v float64)`

SetOwnershipPercentage sets OwnershipPercentage field to given value.

### HasOwnershipPercentage

`func (o *BusinessRelationshipDto) HasOwnershipPercentage() bool`

HasOwnershipPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


