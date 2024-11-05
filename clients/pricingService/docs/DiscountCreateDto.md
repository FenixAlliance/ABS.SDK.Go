# DiscountCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BeginQuantity** | Pointer to **float64** |  | [optional] 
**EndQuantity** | Pointer to **float64** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDiscountCreateDto

`func NewDiscountCreateDto() *DiscountCreateDto`

NewDiscountCreateDto instantiates a new DiscountCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountCreateDtoWithDefaults

`func NewDiscountCreateDtoWithDefaults() *DiscountCreateDto`

NewDiscountCreateDtoWithDefaults instantiates a new DiscountCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscountCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscountCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscountCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscountCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *DiscountCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscountCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscountCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscountCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *DiscountCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscountCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscountCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscountCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscountCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscountCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBeginQuantity

`func (o *DiscountCreateDto) GetBeginQuantity() float64`

GetBeginQuantity returns the BeginQuantity field if non-nil, zero value otherwise.

### GetBeginQuantityOk

`func (o *DiscountCreateDto) GetBeginQuantityOk() (*float64, bool)`

GetBeginQuantityOk returns a tuple with the BeginQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginQuantity

`func (o *DiscountCreateDto) SetBeginQuantity(v float64)`

SetBeginQuantity sets BeginQuantity field to given value.

### HasBeginQuantity

`func (o *DiscountCreateDto) HasBeginQuantity() bool`

HasBeginQuantity returns a boolean if a field has been set.

### GetEndQuantity

`func (o *DiscountCreateDto) GetEndQuantity() float64`

GetEndQuantity returns the EndQuantity field if non-nil, zero value otherwise.

### GetEndQuantityOk

`func (o *DiscountCreateDto) GetEndQuantityOk() (*float64, bool)`

GetEndQuantityOk returns a tuple with the EndQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndQuantity

`func (o *DiscountCreateDto) SetEndQuantity(v float64)`

SetEndQuantity sets EndQuantity field to given value.

### HasEndQuantity

`func (o *DiscountCreateDto) HasEndQuantity() bool`

HasEndQuantity returns a boolean if a field has been set.

### GetPercent

`func (o *DiscountCreateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *DiscountCreateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *DiscountCreateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *DiscountCreateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetValue

`func (o *DiscountCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DiscountCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DiscountCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DiscountCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTenantId

`func (o *DiscountCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DiscountCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DiscountCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DiscountCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DiscountCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DiscountCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *DiscountCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *DiscountCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *DiscountCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *DiscountCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *DiscountCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *DiscountCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetDiscountListId

`func (o *DiscountCreateDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *DiscountCreateDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *DiscountCreateDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *DiscountCreateDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *DiscountCreateDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *DiscountCreateDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


