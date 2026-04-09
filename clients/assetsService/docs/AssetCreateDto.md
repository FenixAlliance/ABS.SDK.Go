# AssetCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessId** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AssetClass** | Pointer to **string** |  | [optional] 
**AssetOwner** | Pointer to **string** |  | [optional] 
**IsExistingAsset** | Pointer to **bool** |  | [optional] 
**CalculateDepreciation** | Pointer to **bool** |  | [optional] 
**AllowMonthlyDepreciation** | Pointer to **bool** |  | [optional] 
**OpeningDepreciation** | Pointer to **float64** |  | [optional] 
**PurchaseDate** | Pointer to **time.Time** |  | [optional] 
**PurchasePrice** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**AssetCategoryId** | Pointer to **NullableString** |  | [optional] 
**PurchaseInvoiceId** | Pointer to **NullableString** |  | [optional] 
**PurchaseReceiptId** | Pointer to **NullableString** |  | [optional] 
**AssetLocationId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**OrganizationDepartmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetCreateDto

`func NewAssetCreateDto() *AssetCreateDto`

NewAssetCreateDto instantiates a new AssetCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetCreateDtoWithDefaults

`func NewAssetCreateDtoWithDefaults() *AssetCreateDto`

NewAssetCreateDtoWithDefaults instantiates a new AssetCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessId

`func (o *AssetCreateDto) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AssetCreateDto) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AssetCreateDto) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AssetCreateDto) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### SetBusinessIdNil

`func (o *AssetCreateDto) SetBusinessIdNil(b bool)`

 SetBusinessIdNil sets the value for BusinessId to be an explicit nil

### UnsetBusinessId
`func (o *AssetCreateDto) UnsetBusinessId()`

UnsetBusinessId ensures that no value is present for BusinessId, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *AssetCreateDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *AssetCreateDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *AssetCreateDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *AssetCreateDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *AssetCreateDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *AssetCreateDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetName

`func (o *AssetCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AssetCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AssetCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AssetCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AssetCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AssetCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAssetClass

`func (o *AssetCreateDto) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *AssetCreateDto) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *AssetCreateDto) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *AssetCreateDto) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetAssetOwner

`func (o *AssetCreateDto) GetAssetOwner() string`

GetAssetOwner returns the AssetOwner field if non-nil, zero value otherwise.

### GetAssetOwnerOk

`func (o *AssetCreateDto) GetAssetOwnerOk() (*string, bool)`

GetAssetOwnerOk returns a tuple with the AssetOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetOwner

`func (o *AssetCreateDto) SetAssetOwner(v string)`

SetAssetOwner sets AssetOwner field to given value.

### HasAssetOwner

`func (o *AssetCreateDto) HasAssetOwner() bool`

HasAssetOwner returns a boolean if a field has been set.

### GetIsExistingAsset

`func (o *AssetCreateDto) GetIsExistingAsset() bool`

GetIsExistingAsset returns the IsExistingAsset field if non-nil, zero value otherwise.

### GetIsExistingAssetOk

`func (o *AssetCreateDto) GetIsExistingAssetOk() (*bool, bool)`

GetIsExistingAssetOk returns a tuple with the IsExistingAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExistingAsset

`func (o *AssetCreateDto) SetIsExistingAsset(v bool)`

SetIsExistingAsset sets IsExistingAsset field to given value.

### HasIsExistingAsset

`func (o *AssetCreateDto) HasIsExistingAsset() bool`

HasIsExistingAsset returns a boolean if a field has been set.

### GetCalculateDepreciation

`func (o *AssetCreateDto) GetCalculateDepreciation() bool`

GetCalculateDepreciation returns the CalculateDepreciation field if non-nil, zero value otherwise.

### GetCalculateDepreciationOk

`func (o *AssetCreateDto) GetCalculateDepreciationOk() (*bool, bool)`

GetCalculateDepreciationOk returns a tuple with the CalculateDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateDepreciation

`func (o *AssetCreateDto) SetCalculateDepreciation(v bool)`

SetCalculateDepreciation sets CalculateDepreciation field to given value.

### HasCalculateDepreciation

`func (o *AssetCreateDto) HasCalculateDepreciation() bool`

HasCalculateDepreciation returns a boolean if a field has been set.

### GetAllowMonthlyDepreciation

`func (o *AssetCreateDto) GetAllowMonthlyDepreciation() bool`

GetAllowMonthlyDepreciation returns the AllowMonthlyDepreciation field if non-nil, zero value otherwise.

### GetAllowMonthlyDepreciationOk

`func (o *AssetCreateDto) GetAllowMonthlyDepreciationOk() (*bool, bool)`

GetAllowMonthlyDepreciationOk returns a tuple with the AllowMonthlyDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMonthlyDepreciation

`func (o *AssetCreateDto) SetAllowMonthlyDepreciation(v bool)`

SetAllowMonthlyDepreciation sets AllowMonthlyDepreciation field to given value.

### HasAllowMonthlyDepreciation

`func (o *AssetCreateDto) HasAllowMonthlyDepreciation() bool`

HasAllowMonthlyDepreciation returns a boolean if a field has been set.

### GetOpeningDepreciation

`func (o *AssetCreateDto) GetOpeningDepreciation() float64`

GetOpeningDepreciation returns the OpeningDepreciation field if non-nil, zero value otherwise.

### GetOpeningDepreciationOk

`func (o *AssetCreateDto) GetOpeningDepreciationOk() (*float64, bool)`

GetOpeningDepreciationOk returns a tuple with the OpeningDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDepreciation

`func (o *AssetCreateDto) SetOpeningDepreciation(v float64)`

SetOpeningDepreciation sets OpeningDepreciation field to given value.

### HasOpeningDepreciation

`func (o *AssetCreateDto) HasOpeningDepreciation() bool`

HasOpeningDepreciation returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *AssetCreateDto) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *AssetCreateDto) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *AssetCreateDto) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *AssetCreateDto) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *AssetCreateDto) GetPurchasePrice() float64`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *AssetCreateDto) GetPurchasePriceOk() (*float64, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *AssetCreateDto) SetPurchasePrice(v float64)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *AssetCreateDto) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AssetCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AssetCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AssetCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AssetCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AssetCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AssetCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetItemId

`func (o *AssetCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AssetCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AssetCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AssetCreateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AssetCreateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AssetCreateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetAssetCategoryId

`func (o *AssetCreateDto) GetAssetCategoryId() string`

GetAssetCategoryId returns the AssetCategoryId field if non-nil, zero value otherwise.

### GetAssetCategoryIdOk

`func (o *AssetCreateDto) GetAssetCategoryIdOk() (*string, bool)`

GetAssetCategoryIdOk returns a tuple with the AssetCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCategoryId

`func (o *AssetCreateDto) SetAssetCategoryId(v string)`

SetAssetCategoryId sets AssetCategoryId field to given value.

### HasAssetCategoryId

`func (o *AssetCreateDto) HasAssetCategoryId() bool`

HasAssetCategoryId returns a boolean if a field has been set.

### SetAssetCategoryIdNil

`func (o *AssetCreateDto) SetAssetCategoryIdNil(b bool)`

 SetAssetCategoryIdNil sets the value for AssetCategoryId to be an explicit nil

### UnsetAssetCategoryId
`func (o *AssetCreateDto) UnsetAssetCategoryId()`

UnsetAssetCategoryId ensures that no value is present for AssetCategoryId, not even an explicit nil
### GetPurchaseInvoiceId

`func (o *AssetCreateDto) GetPurchaseInvoiceId() string`

GetPurchaseInvoiceId returns the PurchaseInvoiceId field if non-nil, zero value otherwise.

### GetPurchaseInvoiceIdOk

`func (o *AssetCreateDto) GetPurchaseInvoiceIdOk() (*string, bool)`

GetPurchaseInvoiceIdOk returns a tuple with the PurchaseInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseInvoiceId

`func (o *AssetCreateDto) SetPurchaseInvoiceId(v string)`

SetPurchaseInvoiceId sets PurchaseInvoiceId field to given value.

### HasPurchaseInvoiceId

`func (o *AssetCreateDto) HasPurchaseInvoiceId() bool`

HasPurchaseInvoiceId returns a boolean if a field has been set.

### SetPurchaseInvoiceIdNil

`func (o *AssetCreateDto) SetPurchaseInvoiceIdNil(b bool)`

 SetPurchaseInvoiceIdNil sets the value for PurchaseInvoiceId to be an explicit nil

### UnsetPurchaseInvoiceId
`func (o *AssetCreateDto) UnsetPurchaseInvoiceId()`

UnsetPurchaseInvoiceId ensures that no value is present for PurchaseInvoiceId, not even an explicit nil
### GetPurchaseReceiptId

`func (o *AssetCreateDto) GetPurchaseReceiptId() string`

GetPurchaseReceiptId returns the PurchaseReceiptId field if non-nil, zero value otherwise.

### GetPurchaseReceiptIdOk

`func (o *AssetCreateDto) GetPurchaseReceiptIdOk() (*string, bool)`

GetPurchaseReceiptIdOk returns a tuple with the PurchaseReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseReceiptId

`func (o *AssetCreateDto) SetPurchaseReceiptId(v string)`

SetPurchaseReceiptId sets PurchaseReceiptId field to given value.

### HasPurchaseReceiptId

`func (o *AssetCreateDto) HasPurchaseReceiptId() bool`

HasPurchaseReceiptId returns a boolean if a field has been set.

### SetPurchaseReceiptIdNil

`func (o *AssetCreateDto) SetPurchaseReceiptIdNil(b bool)`

 SetPurchaseReceiptIdNil sets the value for PurchaseReceiptId to be an explicit nil

### UnsetPurchaseReceiptId
`func (o *AssetCreateDto) UnsetPurchaseReceiptId()`

UnsetPurchaseReceiptId ensures that no value is present for PurchaseReceiptId, not even an explicit nil
### GetAssetLocationId

`func (o *AssetCreateDto) GetAssetLocationId() string`

GetAssetLocationId returns the AssetLocationId field if non-nil, zero value otherwise.

### GetAssetLocationIdOk

`func (o *AssetCreateDto) GetAssetLocationIdOk() (*string, bool)`

GetAssetLocationIdOk returns a tuple with the AssetLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLocationId

`func (o *AssetCreateDto) SetAssetLocationId(v string)`

SetAssetLocationId sets AssetLocationId field to given value.

### HasAssetLocationId

`func (o *AssetCreateDto) HasAssetLocationId() bool`

HasAssetLocationId returns a boolean if a field has been set.

### SetAssetLocationIdNil

`func (o *AssetCreateDto) SetAssetLocationIdNil(b bool)`

 SetAssetLocationIdNil sets the value for AssetLocationId to be an explicit nil

### UnsetAssetLocationId
`func (o *AssetCreateDto) UnsetAssetLocationId()`

UnsetAssetLocationId ensures that no value is present for AssetLocationId, not even an explicit nil
### GetContactId

`func (o *AssetCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *AssetCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *AssetCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *AssetCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *AssetCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *AssetCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetOrganizationDepartmentId

`func (o *AssetCreateDto) GetOrganizationDepartmentId() string`

GetOrganizationDepartmentId returns the OrganizationDepartmentId field if non-nil, zero value otherwise.

### GetOrganizationDepartmentIdOk

`func (o *AssetCreateDto) GetOrganizationDepartmentIdOk() (*string, bool)`

GetOrganizationDepartmentIdOk returns a tuple with the OrganizationDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDepartmentId

`func (o *AssetCreateDto) SetOrganizationDepartmentId(v string)`

SetOrganizationDepartmentId sets OrganizationDepartmentId field to given value.

### HasOrganizationDepartmentId

`func (o *AssetCreateDto) HasOrganizationDepartmentId() bool`

HasOrganizationDepartmentId returns a boolean if a field has been set.

### SetOrganizationDepartmentIdNil

`func (o *AssetCreateDto) SetOrganizationDepartmentIdNil(b bool)`

 SetOrganizationDepartmentIdNil sets the value for OrganizationDepartmentId to be an explicit nil

### UnsetOrganizationDepartmentId
`func (o *AssetCreateDto) UnsetOrganizationDepartmentId()`

UnsetOrganizationDepartmentId ensures that no value is present for OrganizationDepartmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


