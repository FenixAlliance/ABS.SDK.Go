# ItemPackingSlipEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Quantity** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**ItemPackingSlipId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemPackingSlipEntryDto

`func NewItemPackingSlipEntryDto() *ItemPackingSlipEntryDto`

NewItemPackingSlipEntryDto instantiates a new ItemPackingSlipEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPackingSlipEntryDtoWithDefaults

`func NewItemPackingSlipEntryDtoWithDefaults() *ItemPackingSlipEntryDto`

NewItemPackingSlipEntryDtoWithDefaults instantiates a new ItemPackingSlipEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPackingSlipEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPackingSlipEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPackingSlipEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPackingSlipEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemPackingSlipEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemPackingSlipEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemPackingSlipEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPackingSlipEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPackingSlipEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPackingSlipEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemPackingSlipEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemPackingSlipEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQuantity

`func (o *ItemPackingSlipEntryDto) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemPackingSlipEntryDto) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemPackingSlipEntryDto) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ItemPackingSlipEntryDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *ItemPackingSlipEntryDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *ItemPackingSlipEntryDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetItemId

`func (o *ItemPackingSlipEntryDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPackingSlipEntryDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPackingSlipEntryDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemPackingSlipEntryDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemPackingSlipEntryDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemPackingSlipEntryDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemPackingSlipId

`func (o *ItemPackingSlipEntryDto) GetItemPackingSlipId() string`

GetItemPackingSlipId returns the ItemPackingSlipId field if non-nil, zero value otherwise.

### GetItemPackingSlipIdOk

`func (o *ItemPackingSlipEntryDto) GetItemPackingSlipIdOk() (*string, bool)`

GetItemPackingSlipIdOk returns a tuple with the ItemPackingSlipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPackingSlipId

`func (o *ItemPackingSlipEntryDto) SetItemPackingSlipId(v string)`

SetItemPackingSlipId sets ItemPackingSlipId field to given value.

### HasItemPackingSlipId

`func (o *ItemPackingSlipEntryDto) HasItemPackingSlipId() bool`

HasItemPackingSlipId returns a boolean if a field has been set.

### SetItemPackingSlipIdNil

`func (o *ItemPackingSlipEntryDto) SetItemPackingSlipIdNil(b bool)`

 SetItemPackingSlipIdNil sets the value for ItemPackingSlipId to be an explicit nil

### UnsetItemPackingSlipId
`func (o *ItemPackingSlipEntryDto) UnsetItemPackingSlipId()`

UnsetItemPackingSlipId ensures that no value is present for ItemPackingSlipId, not even an explicit nil
### GetTenantId

`func (o *ItemPackingSlipEntryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemPackingSlipEntryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemPackingSlipEntryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemPackingSlipEntryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemPackingSlipEntryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemPackingSlipEntryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


