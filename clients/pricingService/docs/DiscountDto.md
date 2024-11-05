# DiscountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 
**EndQuantity** | Pointer to **float64** |  | [optional] 
**BeginQuantity** | Pointer to **float64** |  | [optional] 

## Methods

### NewDiscountDto

`func NewDiscountDto() *DiscountDto`

NewDiscountDto instantiates a new DiscountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountDtoWithDefaults

`func NewDiscountDtoWithDefaults() *DiscountDto`

NewDiscountDtoWithDefaults instantiates a new DiscountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscountDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscountDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscountDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscountDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DiscountDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DiscountDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *DiscountDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscountDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscountDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscountDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DiscountDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DiscountDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDescription

`func (o *DiscountDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscountDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscountDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscountDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscountDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscountDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetValue

`func (o *DiscountDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DiscountDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DiscountDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DiscountDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPercent

`func (o *DiscountDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *DiscountDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *DiscountDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *DiscountDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetItemId

`func (o *DiscountDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DiscountDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DiscountDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DiscountDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *DiscountDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *DiscountDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetTenantId

`func (o *DiscountDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DiscountDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DiscountDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DiscountDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DiscountDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DiscountDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *DiscountDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *DiscountDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *DiscountDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *DiscountDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *DiscountDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *DiscountDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetDiscountListId

`func (o *DiscountDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *DiscountDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *DiscountDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *DiscountDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *DiscountDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *DiscountDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil
### GetEndQuantity

`func (o *DiscountDto) GetEndQuantity() float64`

GetEndQuantity returns the EndQuantity field if non-nil, zero value otherwise.

### GetEndQuantityOk

`func (o *DiscountDto) GetEndQuantityOk() (*float64, bool)`

GetEndQuantityOk returns a tuple with the EndQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndQuantity

`func (o *DiscountDto) SetEndQuantity(v float64)`

SetEndQuantity sets EndQuantity field to given value.

### HasEndQuantity

`func (o *DiscountDto) HasEndQuantity() bool`

HasEndQuantity returns a boolean if a field has been set.

### GetBeginQuantity

`func (o *DiscountDto) GetBeginQuantity() float64`

GetBeginQuantity returns the BeginQuantity field if non-nil, zero value otherwise.

### GetBeginQuantityOk

`func (o *DiscountDto) GetBeginQuantityOk() (*float64, bool)`

GetBeginQuantityOk returns a tuple with the BeginQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginQuantity

`func (o *DiscountDto) SetBeginQuantity(v float64)`

SetBeginQuantity sets BeginQuantity field to given value.

### HasBeginQuantity

`func (o *DiscountDto) HasBeginQuantity() bool`

HasBeginQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


