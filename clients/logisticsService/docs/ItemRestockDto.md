# ItemRestockDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EntryCount** | Pointer to **int32** |  | [optional] 
**SellerBillingProfileId** | Pointer to **NullableString** |  | [optional] 
**BuyerBillingProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemRestockDto

`func NewItemRestockDto() *ItemRestockDto`

NewItemRestockDto instantiates a new ItemRestockDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemRestockDtoWithDefaults

`func NewItemRestockDtoWithDefaults() *ItemRestockDto`

NewItemRestockDtoWithDefaults instantiates a new ItemRestockDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemRestockDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemRestockDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemRestockDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemRestockDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemRestockDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemRestockDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemRestockDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemRestockDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemRestockDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemRestockDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemRestockDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemRestockDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *ItemRestockDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemRestockDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemRestockDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemRestockDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ItemRestockDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ItemRestockDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ItemRestockDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemRestockDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemRestockDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemRestockDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemRestockDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemRestockDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTenantId

`func (o *ItemRestockDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemRestockDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemRestockDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemRestockDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemRestockDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemRestockDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEntryCount

`func (o *ItemRestockDto) GetEntryCount() int32`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *ItemRestockDto) GetEntryCountOk() (*int32, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *ItemRestockDto) SetEntryCount(v int32)`

SetEntryCount sets EntryCount field to given value.

### HasEntryCount

`func (o *ItemRestockDto) HasEntryCount() bool`

HasEntryCount returns a boolean if a field has been set.

### GetSellerBillingProfileId

`func (o *ItemRestockDto) GetSellerBillingProfileId() string`

GetSellerBillingProfileId returns the SellerBillingProfileId field if non-nil, zero value otherwise.

### GetSellerBillingProfileIdOk

`func (o *ItemRestockDto) GetSellerBillingProfileIdOk() (*string, bool)`

GetSellerBillingProfileIdOk returns a tuple with the SellerBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerBillingProfileId

`func (o *ItemRestockDto) SetSellerBillingProfileId(v string)`

SetSellerBillingProfileId sets SellerBillingProfileId field to given value.

### HasSellerBillingProfileId

`func (o *ItemRestockDto) HasSellerBillingProfileId() bool`

HasSellerBillingProfileId returns a boolean if a field has been set.

### SetSellerBillingProfileIdNil

`func (o *ItemRestockDto) SetSellerBillingProfileIdNil(b bool)`

 SetSellerBillingProfileIdNil sets the value for SellerBillingProfileId to be an explicit nil

### UnsetSellerBillingProfileId
`func (o *ItemRestockDto) UnsetSellerBillingProfileId()`

UnsetSellerBillingProfileId ensures that no value is present for SellerBillingProfileId, not even an explicit nil
### GetBuyerBillingProfileId

`func (o *ItemRestockDto) GetBuyerBillingProfileId() string`

GetBuyerBillingProfileId returns the BuyerBillingProfileId field if non-nil, zero value otherwise.

### GetBuyerBillingProfileIdOk

`func (o *ItemRestockDto) GetBuyerBillingProfileIdOk() (*string, bool)`

GetBuyerBillingProfileIdOk returns a tuple with the BuyerBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerBillingProfileId

`func (o *ItemRestockDto) SetBuyerBillingProfileId(v string)`

SetBuyerBillingProfileId sets BuyerBillingProfileId field to given value.

### HasBuyerBillingProfileId

`func (o *ItemRestockDto) HasBuyerBillingProfileId() bool`

HasBuyerBillingProfileId returns a boolean if a field has been set.

### SetBuyerBillingProfileIdNil

`func (o *ItemRestockDto) SetBuyerBillingProfileIdNil(b bool)`

 SetBuyerBillingProfileIdNil sets the value for BuyerBillingProfileId to be an explicit nil

### UnsetBuyerBillingProfileId
`func (o *ItemRestockDto) UnsetBuyerBillingProfileId()`

UnsetBuyerBillingProfileId ensures that no value is present for BuyerBillingProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


