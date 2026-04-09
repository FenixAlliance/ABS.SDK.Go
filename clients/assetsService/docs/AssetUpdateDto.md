# AssetUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AssetType** | Pointer to **NullableString** |  | [optional] 
**AssetOwner** | Pointer to **NullableString** |  | [optional] 
**CalculateDepreciation** | Pointer to **bool** |  | [optional] 
**AllowMonthlyDepreciation** | Pointer to **bool** |  | [optional] 
**OpeningDepreciation** | Pointer to **NullableFloat64** |  | [optional] 
**PurchaseDate** | Pointer to **NullableTime** |  | [optional] 
**PurchasePrice** | Pointer to **NullableFloat64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**AssetCategoryId** | Pointer to **NullableString** |  | [optional] 
**PurchaseInvoiceId** | Pointer to **NullableString** |  | [optional] 
**PurchaseReceiptId** | Pointer to **NullableString** |  | [optional] 
**AssetLocationId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**OrganizationDepartmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetUpdateDto

`func NewAssetUpdateDto() *AssetUpdateDto`

NewAssetUpdateDto instantiates a new AssetUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetUpdateDtoWithDefaults

`func NewAssetUpdateDtoWithDefaults() *AssetUpdateDto`

NewAssetUpdateDtoWithDefaults instantiates a new AssetUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssetUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AssetUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AssetUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AssetUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AssetUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AssetUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAssetType

`func (o *AssetUpdateDto) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetUpdateDto) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetUpdateDto) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *AssetUpdateDto) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### SetAssetTypeNil

`func (o *AssetUpdateDto) SetAssetTypeNil(b bool)`

 SetAssetTypeNil sets the value for AssetType to be an explicit nil

### UnsetAssetType
`func (o *AssetUpdateDto) UnsetAssetType()`

UnsetAssetType ensures that no value is present for AssetType, not even an explicit nil
### GetAssetOwner

`func (o *AssetUpdateDto) GetAssetOwner() string`

GetAssetOwner returns the AssetOwner field if non-nil, zero value otherwise.

### GetAssetOwnerOk

`func (o *AssetUpdateDto) GetAssetOwnerOk() (*string, bool)`

GetAssetOwnerOk returns a tuple with the AssetOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetOwner

`func (o *AssetUpdateDto) SetAssetOwner(v string)`

SetAssetOwner sets AssetOwner field to given value.

### HasAssetOwner

`func (o *AssetUpdateDto) HasAssetOwner() bool`

HasAssetOwner returns a boolean if a field has been set.

### SetAssetOwnerNil

`func (o *AssetUpdateDto) SetAssetOwnerNil(b bool)`

 SetAssetOwnerNil sets the value for AssetOwner to be an explicit nil

### UnsetAssetOwner
`func (o *AssetUpdateDto) UnsetAssetOwner()`

UnsetAssetOwner ensures that no value is present for AssetOwner, not even an explicit nil
### GetCalculateDepreciation

`func (o *AssetUpdateDto) GetCalculateDepreciation() bool`

GetCalculateDepreciation returns the CalculateDepreciation field if non-nil, zero value otherwise.

### GetCalculateDepreciationOk

`func (o *AssetUpdateDto) GetCalculateDepreciationOk() (*bool, bool)`

GetCalculateDepreciationOk returns a tuple with the CalculateDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateDepreciation

`func (o *AssetUpdateDto) SetCalculateDepreciation(v bool)`

SetCalculateDepreciation sets CalculateDepreciation field to given value.

### HasCalculateDepreciation

`func (o *AssetUpdateDto) HasCalculateDepreciation() bool`

HasCalculateDepreciation returns a boolean if a field has been set.

### GetAllowMonthlyDepreciation

`func (o *AssetUpdateDto) GetAllowMonthlyDepreciation() bool`

GetAllowMonthlyDepreciation returns the AllowMonthlyDepreciation field if non-nil, zero value otherwise.

### GetAllowMonthlyDepreciationOk

`func (o *AssetUpdateDto) GetAllowMonthlyDepreciationOk() (*bool, bool)`

GetAllowMonthlyDepreciationOk returns a tuple with the AllowMonthlyDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMonthlyDepreciation

`func (o *AssetUpdateDto) SetAllowMonthlyDepreciation(v bool)`

SetAllowMonthlyDepreciation sets AllowMonthlyDepreciation field to given value.

### HasAllowMonthlyDepreciation

`func (o *AssetUpdateDto) HasAllowMonthlyDepreciation() bool`

HasAllowMonthlyDepreciation returns a boolean if a field has been set.

### GetOpeningDepreciation

`func (o *AssetUpdateDto) GetOpeningDepreciation() float64`

GetOpeningDepreciation returns the OpeningDepreciation field if non-nil, zero value otherwise.

### GetOpeningDepreciationOk

`func (o *AssetUpdateDto) GetOpeningDepreciationOk() (*float64, bool)`

GetOpeningDepreciationOk returns a tuple with the OpeningDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDepreciation

`func (o *AssetUpdateDto) SetOpeningDepreciation(v float64)`

SetOpeningDepreciation sets OpeningDepreciation field to given value.

### HasOpeningDepreciation

`func (o *AssetUpdateDto) HasOpeningDepreciation() bool`

HasOpeningDepreciation returns a boolean if a field has been set.

### SetOpeningDepreciationNil

`func (o *AssetUpdateDto) SetOpeningDepreciationNil(b bool)`

 SetOpeningDepreciationNil sets the value for OpeningDepreciation to be an explicit nil

### UnsetOpeningDepreciation
`func (o *AssetUpdateDto) UnsetOpeningDepreciation()`

UnsetOpeningDepreciation ensures that no value is present for OpeningDepreciation, not even an explicit nil
### GetPurchaseDate

`func (o *AssetUpdateDto) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *AssetUpdateDto) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *AssetUpdateDto) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *AssetUpdateDto) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### SetPurchaseDateNil

`func (o *AssetUpdateDto) SetPurchaseDateNil(b bool)`

 SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil

### UnsetPurchaseDate
`func (o *AssetUpdateDto) UnsetPurchaseDate()`

UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
### GetPurchasePrice

`func (o *AssetUpdateDto) GetPurchasePrice() float64`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *AssetUpdateDto) GetPurchasePriceOk() (*float64, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *AssetUpdateDto) SetPurchasePrice(v float64)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *AssetUpdateDto) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *AssetUpdateDto) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *AssetUpdateDto) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetCurrencyId

`func (o *AssetUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AssetUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AssetUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AssetUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AssetUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AssetUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurrencyCode

`func (o *AssetUpdateDto) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AssetUpdateDto) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AssetUpdateDto) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AssetUpdateDto) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *AssetUpdateDto) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *AssetUpdateDto) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetItemId

`func (o *AssetUpdateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AssetUpdateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AssetUpdateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AssetUpdateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AssetUpdateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AssetUpdateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetAssetCategoryId

`func (o *AssetUpdateDto) GetAssetCategoryId() string`

GetAssetCategoryId returns the AssetCategoryId field if non-nil, zero value otherwise.

### GetAssetCategoryIdOk

`func (o *AssetUpdateDto) GetAssetCategoryIdOk() (*string, bool)`

GetAssetCategoryIdOk returns a tuple with the AssetCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCategoryId

`func (o *AssetUpdateDto) SetAssetCategoryId(v string)`

SetAssetCategoryId sets AssetCategoryId field to given value.

### HasAssetCategoryId

`func (o *AssetUpdateDto) HasAssetCategoryId() bool`

HasAssetCategoryId returns a boolean if a field has been set.

### SetAssetCategoryIdNil

`func (o *AssetUpdateDto) SetAssetCategoryIdNil(b bool)`

 SetAssetCategoryIdNil sets the value for AssetCategoryId to be an explicit nil

### UnsetAssetCategoryId
`func (o *AssetUpdateDto) UnsetAssetCategoryId()`

UnsetAssetCategoryId ensures that no value is present for AssetCategoryId, not even an explicit nil
### GetPurchaseInvoiceId

`func (o *AssetUpdateDto) GetPurchaseInvoiceId() string`

GetPurchaseInvoiceId returns the PurchaseInvoiceId field if non-nil, zero value otherwise.

### GetPurchaseInvoiceIdOk

`func (o *AssetUpdateDto) GetPurchaseInvoiceIdOk() (*string, bool)`

GetPurchaseInvoiceIdOk returns a tuple with the PurchaseInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseInvoiceId

`func (o *AssetUpdateDto) SetPurchaseInvoiceId(v string)`

SetPurchaseInvoiceId sets PurchaseInvoiceId field to given value.

### HasPurchaseInvoiceId

`func (o *AssetUpdateDto) HasPurchaseInvoiceId() bool`

HasPurchaseInvoiceId returns a boolean if a field has been set.

### SetPurchaseInvoiceIdNil

`func (o *AssetUpdateDto) SetPurchaseInvoiceIdNil(b bool)`

 SetPurchaseInvoiceIdNil sets the value for PurchaseInvoiceId to be an explicit nil

### UnsetPurchaseInvoiceId
`func (o *AssetUpdateDto) UnsetPurchaseInvoiceId()`

UnsetPurchaseInvoiceId ensures that no value is present for PurchaseInvoiceId, not even an explicit nil
### GetPurchaseReceiptId

`func (o *AssetUpdateDto) GetPurchaseReceiptId() string`

GetPurchaseReceiptId returns the PurchaseReceiptId field if non-nil, zero value otherwise.

### GetPurchaseReceiptIdOk

`func (o *AssetUpdateDto) GetPurchaseReceiptIdOk() (*string, bool)`

GetPurchaseReceiptIdOk returns a tuple with the PurchaseReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseReceiptId

`func (o *AssetUpdateDto) SetPurchaseReceiptId(v string)`

SetPurchaseReceiptId sets PurchaseReceiptId field to given value.

### HasPurchaseReceiptId

`func (o *AssetUpdateDto) HasPurchaseReceiptId() bool`

HasPurchaseReceiptId returns a boolean if a field has been set.

### SetPurchaseReceiptIdNil

`func (o *AssetUpdateDto) SetPurchaseReceiptIdNil(b bool)`

 SetPurchaseReceiptIdNil sets the value for PurchaseReceiptId to be an explicit nil

### UnsetPurchaseReceiptId
`func (o *AssetUpdateDto) UnsetPurchaseReceiptId()`

UnsetPurchaseReceiptId ensures that no value is present for PurchaseReceiptId, not even an explicit nil
### GetAssetLocationId

`func (o *AssetUpdateDto) GetAssetLocationId() string`

GetAssetLocationId returns the AssetLocationId field if non-nil, zero value otherwise.

### GetAssetLocationIdOk

`func (o *AssetUpdateDto) GetAssetLocationIdOk() (*string, bool)`

GetAssetLocationIdOk returns a tuple with the AssetLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLocationId

`func (o *AssetUpdateDto) SetAssetLocationId(v string)`

SetAssetLocationId sets AssetLocationId field to given value.

### HasAssetLocationId

`func (o *AssetUpdateDto) HasAssetLocationId() bool`

HasAssetLocationId returns a boolean if a field has been set.

### SetAssetLocationIdNil

`func (o *AssetUpdateDto) SetAssetLocationIdNil(b bool)`

 SetAssetLocationIdNil sets the value for AssetLocationId to be an explicit nil

### UnsetAssetLocationId
`func (o *AssetUpdateDto) UnsetAssetLocationId()`

UnsetAssetLocationId ensures that no value is present for AssetLocationId, not even an explicit nil
### GetContactId

`func (o *AssetUpdateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *AssetUpdateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *AssetUpdateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *AssetUpdateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *AssetUpdateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *AssetUpdateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetOrganizationDepartmentId

`func (o *AssetUpdateDto) GetOrganizationDepartmentId() string`

GetOrganizationDepartmentId returns the OrganizationDepartmentId field if non-nil, zero value otherwise.

### GetOrganizationDepartmentIdOk

`func (o *AssetUpdateDto) GetOrganizationDepartmentIdOk() (*string, bool)`

GetOrganizationDepartmentIdOk returns a tuple with the OrganizationDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDepartmentId

`func (o *AssetUpdateDto) SetOrganizationDepartmentId(v string)`

SetOrganizationDepartmentId sets OrganizationDepartmentId field to given value.

### HasOrganizationDepartmentId

`func (o *AssetUpdateDto) HasOrganizationDepartmentId() bool`

HasOrganizationDepartmentId returns a boolean if a field has been set.

### SetOrganizationDepartmentIdNil

`func (o *AssetUpdateDto) SetOrganizationDepartmentIdNil(b bool)`

 SetOrganizationDepartmentIdNil sets the value for OrganizationDepartmentId to be an explicit nil

### UnsetOrganizationDepartmentId
`func (o *AssetUpdateDto) UnsetOrganizationDepartmentId()`

UnsetOrganizationDepartmentId ensures that no value is present for OrganizationDepartmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


