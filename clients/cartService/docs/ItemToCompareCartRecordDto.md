# ItemToCompareCartRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**ItemTitle** | Pointer to **NullableString** |  | [optional] 
**ItemShortDescription** | Pointer to **NullableString** |  | [optional] 
**ItemPrimaryImageUrl** | Pointer to **NullableString** |  | [optional] 
**BrandId** | Pointer to **NullableString** |  | [optional] 
**BrandName** | Pointer to **NullableString** |  | [optional] 
**CategoryName** | Pointer to **NullableString** |  | [optional] 
**CategoryId** | Pointer to **NullableString** |  | [optional] 
**GoogleCategoryId** | Pointer to **NullableString** |  | [optional] 
**GoogleCategoryName** | Pointer to **NullableString** |  | [optional] 
**TotalTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalPriceInUsd** | Pointer to **float64** |  | [optional] 
**ShippingCountryId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemToCompareCartRecordDto

`func NewItemToCompareCartRecordDto() *ItemToCompareCartRecordDto`

NewItemToCompareCartRecordDto instantiates a new ItemToCompareCartRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemToCompareCartRecordDtoWithDefaults

`func NewItemToCompareCartRecordDtoWithDefaults() *ItemToCompareCartRecordDto`

NewItemToCompareCartRecordDtoWithDefaults instantiates a new ItemToCompareCartRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemToCompareCartRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemToCompareCartRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemToCompareCartRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemToCompareCartRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemToCompareCartRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemToCompareCartRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemToCompareCartRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemToCompareCartRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemToCompareCartRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemToCompareCartRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemToCompareCartRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemToCompareCartRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCartId

`func (o *ItemToCompareCartRecordDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *ItemToCompareCartRecordDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *ItemToCompareCartRecordDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *ItemToCompareCartRecordDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *ItemToCompareCartRecordDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *ItemToCompareCartRecordDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetItemId

`func (o *ItemToCompareCartRecordDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemToCompareCartRecordDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemToCompareCartRecordDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemToCompareCartRecordDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemToCompareCartRecordDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemToCompareCartRecordDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemTitle

`func (o *ItemToCompareCartRecordDto) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *ItemToCompareCartRecordDto) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *ItemToCompareCartRecordDto) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *ItemToCompareCartRecordDto) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *ItemToCompareCartRecordDto) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *ItemToCompareCartRecordDto) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemShortDescription

`func (o *ItemToCompareCartRecordDto) GetItemShortDescription() string`

GetItemShortDescription returns the ItemShortDescription field if non-nil, zero value otherwise.

### GetItemShortDescriptionOk

`func (o *ItemToCompareCartRecordDto) GetItemShortDescriptionOk() (*string, bool)`

GetItemShortDescriptionOk returns a tuple with the ItemShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemShortDescription

`func (o *ItemToCompareCartRecordDto) SetItemShortDescription(v string)`

SetItemShortDescription sets ItemShortDescription field to given value.

### HasItemShortDescription

`func (o *ItemToCompareCartRecordDto) HasItemShortDescription() bool`

HasItemShortDescription returns a boolean if a field has been set.

### SetItemShortDescriptionNil

`func (o *ItemToCompareCartRecordDto) SetItemShortDescriptionNil(b bool)`

 SetItemShortDescriptionNil sets the value for ItemShortDescription to be an explicit nil

### UnsetItemShortDescription
`func (o *ItemToCompareCartRecordDto) UnsetItemShortDescription()`

UnsetItemShortDescription ensures that no value is present for ItemShortDescription, not even an explicit nil
### GetItemPrimaryImageUrl

`func (o *ItemToCompareCartRecordDto) GetItemPrimaryImageUrl() string`

GetItemPrimaryImageUrl returns the ItemPrimaryImageUrl field if non-nil, zero value otherwise.

### GetItemPrimaryImageUrlOk

`func (o *ItemToCompareCartRecordDto) GetItemPrimaryImageUrlOk() (*string, bool)`

GetItemPrimaryImageUrlOk returns a tuple with the ItemPrimaryImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrimaryImageUrl

`func (o *ItemToCompareCartRecordDto) SetItemPrimaryImageUrl(v string)`

SetItemPrimaryImageUrl sets ItemPrimaryImageUrl field to given value.

### HasItemPrimaryImageUrl

`func (o *ItemToCompareCartRecordDto) HasItemPrimaryImageUrl() bool`

HasItemPrimaryImageUrl returns a boolean if a field has been set.

### SetItemPrimaryImageUrlNil

`func (o *ItemToCompareCartRecordDto) SetItemPrimaryImageUrlNil(b bool)`

 SetItemPrimaryImageUrlNil sets the value for ItemPrimaryImageUrl to be an explicit nil

### UnsetItemPrimaryImageUrl
`func (o *ItemToCompareCartRecordDto) UnsetItemPrimaryImageUrl()`

UnsetItemPrimaryImageUrl ensures that no value is present for ItemPrimaryImageUrl, not even an explicit nil
### GetBrandId

`func (o *ItemToCompareCartRecordDto) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *ItemToCompareCartRecordDto) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *ItemToCompareCartRecordDto) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *ItemToCompareCartRecordDto) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *ItemToCompareCartRecordDto) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *ItemToCompareCartRecordDto) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandName

`func (o *ItemToCompareCartRecordDto) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *ItemToCompareCartRecordDto) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *ItemToCompareCartRecordDto) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *ItemToCompareCartRecordDto) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### SetBrandNameNil

`func (o *ItemToCompareCartRecordDto) SetBrandNameNil(b bool)`

 SetBrandNameNil sets the value for BrandName to be an explicit nil

### UnsetBrandName
`func (o *ItemToCompareCartRecordDto) UnsetBrandName()`

UnsetBrandName ensures that no value is present for BrandName, not even an explicit nil
### GetCategoryName

`func (o *ItemToCompareCartRecordDto) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *ItemToCompareCartRecordDto) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *ItemToCompareCartRecordDto) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *ItemToCompareCartRecordDto) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### SetCategoryNameNil

`func (o *ItemToCompareCartRecordDto) SetCategoryNameNil(b bool)`

 SetCategoryNameNil sets the value for CategoryName to be an explicit nil

### UnsetCategoryName
`func (o *ItemToCompareCartRecordDto) UnsetCategoryName()`

UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
### GetCategoryId

`func (o *ItemToCompareCartRecordDto) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ItemToCompareCartRecordDto) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ItemToCompareCartRecordDto) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ItemToCompareCartRecordDto) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *ItemToCompareCartRecordDto) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *ItemToCompareCartRecordDto) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetGoogleCategoryId

`func (o *ItemToCompareCartRecordDto) GetGoogleCategoryId() string`

GetGoogleCategoryId returns the GoogleCategoryId field if non-nil, zero value otherwise.

### GetGoogleCategoryIdOk

`func (o *ItemToCompareCartRecordDto) GetGoogleCategoryIdOk() (*string, bool)`

GetGoogleCategoryIdOk returns a tuple with the GoogleCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCategoryId

`func (o *ItemToCompareCartRecordDto) SetGoogleCategoryId(v string)`

SetGoogleCategoryId sets GoogleCategoryId field to given value.

### HasGoogleCategoryId

`func (o *ItemToCompareCartRecordDto) HasGoogleCategoryId() bool`

HasGoogleCategoryId returns a boolean if a field has been set.

### SetGoogleCategoryIdNil

`func (o *ItemToCompareCartRecordDto) SetGoogleCategoryIdNil(b bool)`

 SetGoogleCategoryIdNil sets the value for GoogleCategoryId to be an explicit nil

### UnsetGoogleCategoryId
`func (o *ItemToCompareCartRecordDto) UnsetGoogleCategoryId()`

UnsetGoogleCategoryId ensures that no value is present for GoogleCategoryId, not even an explicit nil
### GetGoogleCategoryName

`func (o *ItemToCompareCartRecordDto) GetGoogleCategoryName() string`

GetGoogleCategoryName returns the GoogleCategoryName field if non-nil, zero value otherwise.

### GetGoogleCategoryNameOk

`func (o *ItemToCompareCartRecordDto) GetGoogleCategoryNameOk() (*string, bool)`

GetGoogleCategoryNameOk returns a tuple with the GoogleCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCategoryName

`func (o *ItemToCompareCartRecordDto) SetGoogleCategoryName(v string)`

SetGoogleCategoryName sets GoogleCategoryName field to given value.

### HasGoogleCategoryName

`func (o *ItemToCompareCartRecordDto) HasGoogleCategoryName() bool`

HasGoogleCategoryName returns a boolean if a field has been set.

### SetGoogleCategoryNameNil

`func (o *ItemToCompareCartRecordDto) SetGoogleCategoryNameNil(b bool)`

 SetGoogleCategoryNameNil sets the value for GoogleCategoryName to be an explicit nil

### UnsetGoogleCategoryName
`func (o *ItemToCompareCartRecordDto) UnsetGoogleCategoryName()`

UnsetGoogleCategoryName ensures that no value is present for GoogleCategoryName, not even an explicit nil
### GetTotalTaxesInUsd

`func (o *ItemToCompareCartRecordDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *ItemToCompareCartRecordDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *ItemToCompareCartRecordDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *ItemToCompareCartRecordDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalPriceInUsd

`func (o *ItemToCompareCartRecordDto) GetTotalPriceInUsd() float64`

GetTotalPriceInUsd returns the TotalPriceInUsd field if non-nil, zero value otherwise.

### GetTotalPriceInUsdOk

`func (o *ItemToCompareCartRecordDto) GetTotalPriceInUsdOk() (*float64, bool)`

GetTotalPriceInUsdOk returns a tuple with the TotalPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceInUsd

`func (o *ItemToCompareCartRecordDto) SetTotalPriceInUsd(v float64)`

SetTotalPriceInUsd sets TotalPriceInUsd field to given value.

### HasTotalPriceInUsd

`func (o *ItemToCompareCartRecordDto) HasTotalPriceInUsd() bool`

HasTotalPriceInUsd returns a boolean if a field has been set.

### GetShippingCountryId

`func (o *ItemToCompareCartRecordDto) GetShippingCountryId() string`

GetShippingCountryId returns the ShippingCountryId field if non-nil, zero value otherwise.

### GetShippingCountryIdOk

`func (o *ItemToCompareCartRecordDto) GetShippingCountryIdOk() (*string, bool)`

GetShippingCountryIdOk returns a tuple with the ShippingCountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryId

`func (o *ItemToCompareCartRecordDto) SetShippingCountryId(v string)`

SetShippingCountryId sets ShippingCountryId field to given value.

### HasShippingCountryId

`func (o *ItemToCompareCartRecordDto) HasShippingCountryId() bool`

HasShippingCountryId returns a boolean if a field has been set.

### SetShippingCountryIdNil

`func (o *ItemToCompareCartRecordDto) SetShippingCountryIdNil(b bool)`

 SetShippingCountryIdNil sets the value for ShippingCountryId to be an explicit nil

### UnsetShippingCountryId
`func (o *ItemToCompareCartRecordDto) UnsetShippingCountryId()`

UnsetShippingCountryId ensures that no value is present for ShippingCountryId, not even an explicit nil
### GetTenantId

`func (o *ItemToCompareCartRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemToCompareCartRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemToCompareCartRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemToCompareCartRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemToCompareCartRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemToCompareCartRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


