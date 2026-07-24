# BusinessRelationshipCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ChildTenantId** | Pointer to **NullableString** |  | [optional] 
**OwnershipPercentage** | Pointer to **float64** |  | [optional] 

## Methods

### NewBusinessRelationshipCreateDto

`func NewBusinessRelationshipCreateDto() *BusinessRelationshipCreateDto`

NewBusinessRelationshipCreateDto instantiates a new BusinessRelationshipCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessRelationshipCreateDtoWithDefaults

`func NewBusinessRelationshipCreateDtoWithDefaults() *BusinessRelationshipCreateDto`

NewBusinessRelationshipCreateDtoWithDefaults instantiates a new BusinessRelationshipCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessRelationshipCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessRelationshipCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessRelationshipCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessRelationshipCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BusinessRelationshipCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BusinessRelationshipCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BusinessRelationshipCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BusinessRelationshipCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetChildTenantId

`func (o *BusinessRelationshipCreateDto) GetChildTenantId() string`

GetChildTenantId returns the ChildTenantId field if non-nil, zero value otherwise.

### GetChildTenantIdOk

`func (o *BusinessRelationshipCreateDto) GetChildTenantIdOk() (*string, bool)`

GetChildTenantIdOk returns a tuple with the ChildTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTenantId

`func (o *BusinessRelationshipCreateDto) SetChildTenantId(v string)`

SetChildTenantId sets ChildTenantId field to given value.

### HasChildTenantId

`func (o *BusinessRelationshipCreateDto) HasChildTenantId() bool`

HasChildTenantId returns a boolean if a field has been set.

### SetChildTenantIdNil

`func (o *BusinessRelationshipCreateDto) SetChildTenantIdNil(b bool)`

 SetChildTenantIdNil sets the value for ChildTenantId to be an explicit nil

### UnsetChildTenantId
`func (o *BusinessRelationshipCreateDto) UnsetChildTenantId()`

UnsetChildTenantId ensures that no value is present for ChildTenantId, not even an explicit nil
### GetOwnershipPercentage

`func (o *BusinessRelationshipCreateDto) GetOwnershipPercentage() float64`

GetOwnershipPercentage returns the OwnershipPercentage field if non-nil, zero value otherwise.

### GetOwnershipPercentageOk

`func (o *BusinessRelationshipCreateDto) GetOwnershipPercentageOk() (*float64, bool)`

GetOwnershipPercentageOk returns a tuple with the OwnershipPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipPercentage

`func (o *BusinessRelationshipCreateDto) SetOwnershipPercentage(v float64)`

SetOwnershipPercentage sets OwnershipPercentage field to given value.

### HasOwnershipPercentage

`func (o *BusinessRelationshipCreateDto) HasOwnershipPercentage() bool`

HasOwnershipPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


