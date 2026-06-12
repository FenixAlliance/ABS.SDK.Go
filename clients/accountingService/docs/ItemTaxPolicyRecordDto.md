# ItemTaxPolicyRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 
**ItemPriceId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemTaxPolicyRecordDto

`func NewItemTaxPolicyRecordDto() *ItemTaxPolicyRecordDto`

NewItemTaxPolicyRecordDto instantiates a new ItemTaxPolicyRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemTaxPolicyRecordDtoWithDefaults

`func NewItemTaxPolicyRecordDtoWithDefaults() *ItemTaxPolicyRecordDto`

NewItemTaxPolicyRecordDtoWithDefaults instantiates a new ItemTaxPolicyRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemTaxPolicyRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemTaxPolicyRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemTaxPolicyRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemTaxPolicyRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemTaxPolicyRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemTaxPolicyRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemTaxPolicyRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemTaxPolicyRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemTaxPolicyRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemTaxPolicyRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemTaxPolicyRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemTaxPolicyRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *ItemTaxPolicyRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemTaxPolicyRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemTaxPolicyRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemTaxPolicyRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemTaxPolicyRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemTaxPolicyRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ItemTaxPolicyRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ItemTaxPolicyRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ItemTaxPolicyRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ItemTaxPolicyRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ItemTaxPolicyRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ItemTaxPolicyRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTaxPolicyId

`func (o *ItemTaxPolicyRecordDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *ItemTaxPolicyRecordDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *ItemTaxPolicyRecordDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *ItemTaxPolicyRecordDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *ItemTaxPolicyRecordDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *ItemTaxPolicyRecordDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetItemPriceId

`func (o *ItemTaxPolicyRecordDto) GetItemPriceId() string`

GetItemPriceId returns the ItemPriceId field if non-nil, zero value otherwise.

### GetItemPriceIdOk

`func (o *ItemTaxPolicyRecordDto) GetItemPriceIdOk() (*string, bool)`

GetItemPriceIdOk returns a tuple with the ItemPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceId

`func (o *ItemTaxPolicyRecordDto) SetItemPriceId(v string)`

SetItemPriceId sets ItemPriceId field to given value.

### HasItemPriceId

`func (o *ItemTaxPolicyRecordDto) HasItemPriceId() bool`

HasItemPriceId returns a boolean if a field has been set.

### SetItemPriceIdNil

`func (o *ItemTaxPolicyRecordDto) SetItemPriceIdNil(b bool)`

 SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil

### UnsetItemPriceId
`func (o *ItemTaxPolicyRecordDto) UnsetItemPriceId()`

UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
### GetItemId

`func (o *ItemTaxPolicyRecordDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemTaxPolicyRecordDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemTaxPolicyRecordDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemTaxPolicyRecordDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemTaxPolicyRecordDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemTaxPolicyRecordDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


