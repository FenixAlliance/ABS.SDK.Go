# DiscountUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**BeginQuantity** | Pointer to **float64** |  | [optional] 
**EndQuantity** | Pointer to **float64** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**DiscountListId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDiscountUpdateDto

`func NewDiscountUpdateDto() *DiscountUpdateDto`

NewDiscountUpdateDto instantiates a new DiscountUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountUpdateDtoWithDefaults

`func NewDiscountUpdateDtoWithDefaults() *DiscountUpdateDto`

NewDiscountUpdateDtoWithDefaults instantiates a new DiscountUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DiscountUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscountUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscountUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscountUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscountUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscountUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBeginQuantity

`func (o *DiscountUpdateDto) GetBeginQuantity() float64`

GetBeginQuantity returns the BeginQuantity field if non-nil, zero value otherwise.

### GetBeginQuantityOk

`func (o *DiscountUpdateDto) GetBeginQuantityOk() (*float64, bool)`

GetBeginQuantityOk returns a tuple with the BeginQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginQuantity

`func (o *DiscountUpdateDto) SetBeginQuantity(v float64)`

SetBeginQuantity sets BeginQuantity field to given value.

### HasBeginQuantity

`func (o *DiscountUpdateDto) HasBeginQuantity() bool`

HasBeginQuantity returns a boolean if a field has been set.

### GetEndQuantity

`func (o *DiscountUpdateDto) GetEndQuantity() float64`

GetEndQuantity returns the EndQuantity field if non-nil, zero value otherwise.

### GetEndQuantityOk

`func (o *DiscountUpdateDto) GetEndQuantityOk() (*float64, bool)`

GetEndQuantityOk returns a tuple with the EndQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndQuantity

`func (o *DiscountUpdateDto) SetEndQuantity(v float64)`

SetEndQuantity sets EndQuantity field to given value.

### HasEndQuantity

`func (o *DiscountUpdateDto) HasEndQuantity() bool`

HasEndQuantity returns a boolean if a field has been set.

### GetPercent

`func (o *DiscountUpdateDto) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *DiscountUpdateDto) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *DiscountUpdateDto) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *DiscountUpdateDto) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetValue

`func (o *DiscountUpdateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DiscountUpdateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DiscountUpdateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DiscountUpdateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTenantId

`func (o *DiscountUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DiscountUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DiscountUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DiscountUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DiscountUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DiscountUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *DiscountUpdateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *DiscountUpdateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *DiscountUpdateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *DiscountUpdateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *DiscountUpdateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *DiscountUpdateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetDiscountListId

`func (o *DiscountUpdateDto) GetDiscountListId() string`

GetDiscountListId returns the DiscountListId field if non-nil, zero value otherwise.

### GetDiscountListIdOk

`func (o *DiscountUpdateDto) GetDiscountListIdOk() (*string, bool)`

GetDiscountListIdOk returns a tuple with the DiscountListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountListId

`func (o *DiscountUpdateDto) SetDiscountListId(v string)`

SetDiscountListId sets DiscountListId field to given value.

### HasDiscountListId

`func (o *DiscountUpdateDto) HasDiscountListId() bool`

HasDiscountListId returns a boolean if a field has been set.

### SetDiscountListIdNil

`func (o *DiscountUpdateDto) SetDiscountListIdNil(b bool)`

 SetDiscountListIdNil sets the value for DiscountListId to be an explicit nil

### UnsetDiscountListId
`func (o *DiscountUpdateDto) UnsetDiscountListId()`

UnsetDiscountListId ensures that no value is present for DiscountListId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


