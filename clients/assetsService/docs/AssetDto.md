# AssetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**BusinessName** | Pointer to **NullableString** |  | [optional] 
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
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**ItemName** | Pointer to **NullableString** |  | [optional] 
**AssetCategoryId** | Pointer to **NullableString** |  | [optional] 
**AssetCategoryName** | Pointer to **NullableString** |  | [optional] 
**PurchaseInvoiceId** | Pointer to **NullableString** |  | [optional] 
**PurchaseInvoiceNumber** | Pointer to **NullableString** |  | [optional] 
**AssetLocationId** | Pointer to **NullableString** |  | [optional] 
**AssetLocationName** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 
**OrganizationDepartmentId** | Pointer to **NullableString** |  | [optional] 
**OrganizationDepartmentName** | Pointer to **NullableString** |  | [optional] 
**PurchaseReceiptId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetDto

`func NewAssetDto() *AssetDto`

NewAssetDto instantiates a new AssetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDtoWithDefaults

`func NewAssetDtoWithDefaults() *AssetDto`

NewAssetDtoWithDefaults instantiates a new AssetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AssetDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AssetDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AssetDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *AssetDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssetDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssetDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AssetDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AssetDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AssetDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetBusinessName

`func (o *AssetDto) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *AssetDto) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *AssetDto) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *AssetDto) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### SetBusinessNameNil

`func (o *AssetDto) SetBusinessNameNil(b bool)`

 SetBusinessNameNil sets the value for BusinessName to be an explicit nil

### UnsetBusinessName
`func (o *AssetDto) UnsetBusinessName()`

UnsetBusinessName ensures that no value is present for BusinessName, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *AssetDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *AssetDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *AssetDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *AssetDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *AssetDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *AssetDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetName

`func (o *AssetDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AssetDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AssetDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AssetDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AssetDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AssetDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAssetClass

`func (o *AssetDto) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *AssetDto) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *AssetDto) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *AssetDto) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetAssetOwner

`func (o *AssetDto) GetAssetOwner() string`

GetAssetOwner returns the AssetOwner field if non-nil, zero value otherwise.

### GetAssetOwnerOk

`func (o *AssetDto) GetAssetOwnerOk() (*string, bool)`

GetAssetOwnerOk returns a tuple with the AssetOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetOwner

`func (o *AssetDto) SetAssetOwner(v string)`

SetAssetOwner sets AssetOwner field to given value.

### HasAssetOwner

`func (o *AssetDto) HasAssetOwner() bool`

HasAssetOwner returns a boolean if a field has been set.

### GetIsExistingAsset

`func (o *AssetDto) GetIsExistingAsset() bool`

GetIsExistingAsset returns the IsExistingAsset field if non-nil, zero value otherwise.

### GetIsExistingAssetOk

`func (o *AssetDto) GetIsExistingAssetOk() (*bool, bool)`

GetIsExistingAssetOk returns a tuple with the IsExistingAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExistingAsset

`func (o *AssetDto) SetIsExistingAsset(v bool)`

SetIsExistingAsset sets IsExistingAsset field to given value.

### HasIsExistingAsset

`func (o *AssetDto) HasIsExistingAsset() bool`

HasIsExistingAsset returns a boolean if a field has been set.

### GetCalculateDepreciation

`func (o *AssetDto) GetCalculateDepreciation() bool`

GetCalculateDepreciation returns the CalculateDepreciation field if non-nil, zero value otherwise.

### GetCalculateDepreciationOk

`func (o *AssetDto) GetCalculateDepreciationOk() (*bool, bool)`

GetCalculateDepreciationOk returns a tuple with the CalculateDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateDepreciation

`func (o *AssetDto) SetCalculateDepreciation(v bool)`

SetCalculateDepreciation sets CalculateDepreciation field to given value.

### HasCalculateDepreciation

`func (o *AssetDto) HasCalculateDepreciation() bool`

HasCalculateDepreciation returns a boolean if a field has been set.

### GetAllowMonthlyDepreciation

`func (o *AssetDto) GetAllowMonthlyDepreciation() bool`

GetAllowMonthlyDepreciation returns the AllowMonthlyDepreciation field if non-nil, zero value otherwise.

### GetAllowMonthlyDepreciationOk

`func (o *AssetDto) GetAllowMonthlyDepreciationOk() (*bool, bool)`

GetAllowMonthlyDepreciationOk returns a tuple with the AllowMonthlyDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMonthlyDepreciation

`func (o *AssetDto) SetAllowMonthlyDepreciation(v bool)`

SetAllowMonthlyDepreciation sets AllowMonthlyDepreciation field to given value.

### HasAllowMonthlyDepreciation

`func (o *AssetDto) HasAllowMonthlyDepreciation() bool`

HasAllowMonthlyDepreciation returns a boolean if a field has been set.

### GetOpeningDepreciation

`func (o *AssetDto) GetOpeningDepreciation() float64`

GetOpeningDepreciation returns the OpeningDepreciation field if non-nil, zero value otherwise.

### GetOpeningDepreciationOk

`func (o *AssetDto) GetOpeningDepreciationOk() (*float64, bool)`

GetOpeningDepreciationOk returns a tuple with the OpeningDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDepreciation

`func (o *AssetDto) SetOpeningDepreciation(v float64)`

SetOpeningDepreciation sets OpeningDepreciation field to given value.

### HasOpeningDepreciation

`func (o *AssetDto) HasOpeningDepreciation() bool`

HasOpeningDepreciation returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *AssetDto) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *AssetDto) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *AssetDto) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *AssetDto) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *AssetDto) GetPurchasePrice() float64`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *AssetDto) GetPurchasePriceOk() (*float64, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *AssetDto) SetPurchasePrice(v float64)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *AssetDto) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AssetDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AssetDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AssetDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AssetDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AssetDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AssetDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurrencyCode

`func (o *AssetDto) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AssetDto) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AssetDto) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AssetDto) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *AssetDto) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *AssetDto) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetItemId

`func (o *AssetDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AssetDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AssetDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AssetDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AssetDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AssetDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemName

`func (o *AssetDto) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *AssetDto) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *AssetDto) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *AssetDto) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### SetItemNameNil

`func (o *AssetDto) SetItemNameNil(b bool)`

 SetItemNameNil sets the value for ItemName to be an explicit nil

### UnsetItemName
`func (o *AssetDto) UnsetItemName()`

UnsetItemName ensures that no value is present for ItemName, not even an explicit nil
### GetAssetCategoryId

`func (o *AssetDto) GetAssetCategoryId() string`

GetAssetCategoryId returns the AssetCategoryId field if non-nil, zero value otherwise.

### GetAssetCategoryIdOk

`func (o *AssetDto) GetAssetCategoryIdOk() (*string, bool)`

GetAssetCategoryIdOk returns a tuple with the AssetCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCategoryId

`func (o *AssetDto) SetAssetCategoryId(v string)`

SetAssetCategoryId sets AssetCategoryId field to given value.

### HasAssetCategoryId

`func (o *AssetDto) HasAssetCategoryId() bool`

HasAssetCategoryId returns a boolean if a field has been set.

### SetAssetCategoryIdNil

`func (o *AssetDto) SetAssetCategoryIdNil(b bool)`

 SetAssetCategoryIdNil sets the value for AssetCategoryId to be an explicit nil

### UnsetAssetCategoryId
`func (o *AssetDto) UnsetAssetCategoryId()`

UnsetAssetCategoryId ensures that no value is present for AssetCategoryId, not even an explicit nil
### GetAssetCategoryName

`func (o *AssetDto) GetAssetCategoryName() string`

GetAssetCategoryName returns the AssetCategoryName field if non-nil, zero value otherwise.

### GetAssetCategoryNameOk

`func (o *AssetDto) GetAssetCategoryNameOk() (*string, bool)`

GetAssetCategoryNameOk returns a tuple with the AssetCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCategoryName

`func (o *AssetDto) SetAssetCategoryName(v string)`

SetAssetCategoryName sets AssetCategoryName field to given value.

### HasAssetCategoryName

`func (o *AssetDto) HasAssetCategoryName() bool`

HasAssetCategoryName returns a boolean if a field has been set.

### SetAssetCategoryNameNil

`func (o *AssetDto) SetAssetCategoryNameNil(b bool)`

 SetAssetCategoryNameNil sets the value for AssetCategoryName to be an explicit nil

### UnsetAssetCategoryName
`func (o *AssetDto) UnsetAssetCategoryName()`

UnsetAssetCategoryName ensures that no value is present for AssetCategoryName, not even an explicit nil
### GetPurchaseInvoiceId

`func (o *AssetDto) GetPurchaseInvoiceId() string`

GetPurchaseInvoiceId returns the PurchaseInvoiceId field if non-nil, zero value otherwise.

### GetPurchaseInvoiceIdOk

`func (o *AssetDto) GetPurchaseInvoiceIdOk() (*string, bool)`

GetPurchaseInvoiceIdOk returns a tuple with the PurchaseInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseInvoiceId

`func (o *AssetDto) SetPurchaseInvoiceId(v string)`

SetPurchaseInvoiceId sets PurchaseInvoiceId field to given value.

### HasPurchaseInvoiceId

`func (o *AssetDto) HasPurchaseInvoiceId() bool`

HasPurchaseInvoiceId returns a boolean if a field has been set.

### SetPurchaseInvoiceIdNil

`func (o *AssetDto) SetPurchaseInvoiceIdNil(b bool)`

 SetPurchaseInvoiceIdNil sets the value for PurchaseInvoiceId to be an explicit nil

### UnsetPurchaseInvoiceId
`func (o *AssetDto) UnsetPurchaseInvoiceId()`

UnsetPurchaseInvoiceId ensures that no value is present for PurchaseInvoiceId, not even an explicit nil
### GetPurchaseInvoiceNumber

`func (o *AssetDto) GetPurchaseInvoiceNumber() string`

GetPurchaseInvoiceNumber returns the PurchaseInvoiceNumber field if non-nil, zero value otherwise.

### GetPurchaseInvoiceNumberOk

`func (o *AssetDto) GetPurchaseInvoiceNumberOk() (*string, bool)`

GetPurchaseInvoiceNumberOk returns a tuple with the PurchaseInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseInvoiceNumber

`func (o *AssetDto) SetPurchaseInvoiceNumber(v string)`

SetPurchaseInvoiceNumber sets PurchaseInvoiceNumber field to given value.

### HasPurchaseInvoiceNumber

`func (o *AssetDto) HasPurchaseInvoiceNumber() bool`

HasPurchaseInvoiceNumber returns a boolean if a field has been set.

### SetPurchaseInvoiceNumberNil

`func (o *AssetDto) SetPurchaseInvoiceNumberNil(b bool)`

 SetPurchaseInvoiceNumberNil sets the value for PurchaseInvoiceNumber to be an explicit nil

### UnsetPurchaseInvoiceNumber
`func (o *AssetDto) UnsetPurchaseInvoiceNumber()`

UnsetPurchaseInvoiceNumber ensures that no value is present for PurchaseInvoiceNumber, not even an explicit nil
### GetAssetLocationId

`func (o *AssetDto) GetAssetLocationId() string`

GetAssetLocationId returns the AssetLocationId field if non-nil, zero value otherwise.

### GetAssetLocationIdOk

`func (o *AssetDto) GetAssetLocationIdOk() (*string, bool)`

GetAssetLocationIdOk returns a tuple with the AssetLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLocationId

`func (o *AssetDto) SetAssetLocationId(v string)`

SetAssetLocationId sets AssetLocationId field to given value.

### HasAssetLocationId

`func (o *AssetDto) HasAssetLocationId() bool`

HasAssetLocationId returns a boolean if a field has been set.

### SetAssetLocationIdNil

`func (o *AssetDto) SetAssetLocationIdNil(b bool)`

 SetAssetLocationIdNil sets the value for AssetLocationId to be an explicit nil

### UnsetAssetLocationId
`func (o *AssetDto) UnsetAssetLocationId()`

UnsetAssetLocationId ensures that no value is present for AssetLocationId, not even an explicit nil
### GetAssetLocationName

`func (o *AssetDto) GetAssetLocationName() string`

GetAssetLocationName returns the AssetLocationName field if non-nil, zero value otherwise.

### GetAssetLocationNameOk

`func (o *AssetDto) GetAssetLocationNameOk() (*string, bool)`

GetAssetLocationNameOk returns a tuple with the AssetLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLocationName

`func (o *AssetDto) SetAssetLocationName(v string)`

SetAssetLocationName sets AssetLocationName field to given value.

### HasAssetLocationName

`func (o *AssetDto) HasAssetLocationName() bool`

HasAssetLocationName returns a boolean if a field has been set.

### SetAssetLocationNameNil

`func (o *AssetDto) SetAssetLocationNameNil(b bool)`

 SetAssetLocationNameNil sets the value for AssetLocationName to be an explicit nil

### UnsetAssetLocationName
`func (o *AssetDto) UnsetAssetLocationName()`

UnsetAssetLocationName ensures that no value is present for AssetLocationName, not even an explicit nil
### GetContactId

`func (o *AssetDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *AssetDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *AssetDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *AssetDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *AssetDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *AssetDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetContactName

`func (o *AssetDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *AssetDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *AssetDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *AssetDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *AssetDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *AssetDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetOrganizationDepartmentId

`func (o *AssetDto) GetOrganizationDepartmentId() string`

GetOrganizationDepartmentId returns the OrganizationDepartmentId field if non-nil, zero value otherwise.

### GetOrganizationDepartmentIdOk

`func (o *AssetDto) GetOrganizationDepartmentIdOk() (*string, bool)`

GetOrganizationDepartmentIdOk returns a tuple with the OrganizationDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDepartmentId

`func (o *AssetDto) SetOrganizationDepartmentId(v string)`

SetOrganizationDepartmentId sets OrganizationDepartmentId field to given value.

### HasOrganizationDepartmentId

`func (o *AssetDto) HasOrganizationDepartmentId() bool`

HasOrganizationDepartmentId returns a boolean if a field has been set.

### SetOrganizationDepartmentIdNil

`func (o *AssetDto) SetOrganizationDepartmentIdNil(b bool)`

 SetOrganizationDepartmentIdNil sets the value for OrganizationDepartmentId to be an explicit nil

### UnsetOrganizationDepartmentId
`func (o *AssetDto) UnsetOrganizationDepartmentId()`

UnsetOrganizationDepartmentId ensures that no value is present for OrganizationDepartmentId, not even an explicit nil
### GetOrganizationDepartmentName

`func (o *AssetDto) GetOrganizationDepartmentName() string`

GetOrganizationDepartmentName returns the OrganizationDepartmentName field if non-nil, zero value otherwise.

### GetOrganizationDepartmentNameOk

`func (o *AssetDto) GetOrganizationDepartmentNameOk() (*string, bool)`

GetOrganizationDepartmentNameOk returns a tuple with the OrganizationDepartmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDepartmentName

`func (o *AssetDto) SetOrganizationDepartmentName(v string)`

SetOrganizationDepartmentName sets OrganizationDepartmentName field to given value.

### HasOrganizationDepartmentName

`func (o *AssetDto) HasOrganizationDepartmentName() bool`

HasOrganizationDepartmentName returns a boolean if a field has been set.

### SetOrganizationDepartmentNameNil

`func (o *AssetDto) SetOrganizationDepartmentNameNil(b bool)`

 SetOrganizationDepartmentNameNil sets the value for OrganizationDepartmentName to be an explicit nil

### UnsetOrganizationDepartmentName
`func (o *AssetDto) UnsetOrganizationDepartmentName()`

UnsetOrganizationDepartmentName ensures that no value is present for OrganizationDepartmentName, not even an explicit nil
### GetPurchaseReceiptId

`func (o *AssetDto) GetPurchaseReceiptId() string`

GetPurchaseReceiptId returns the PurchaseReceiptId field if non-nil, zero value otherwise.

### GetPurchaseReceiptIdOk

`func (o *AssetDto) GetPurchaseReceiptIdOk() (*string, bool)`

GetPurchaseReceiptIdOk returns a tuple with the PurchaseReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseReceiptId

`func (o *AssetDto) SetPurchaseReceiptId(v string)`

SetPurchaseReceiptId sets PurchaseReceiptId field to given value.

### HasPurchaseReceiptId

`func (o *AssetDto) HasPurchaseReceiptId() bool`

HasPurchaseReceiptId returns a boolean if a field has been set.

### SetPurchaseReceiptIdNil

`func (o *AssetDto) SetPurchaseReceiptIdNil(b bool)`

 SetPurchaseReceiptIdNil sets the value for PurchaseReceiptId to be an explicit nil

### UnsetPurchaseReceiptId
`func (o *AssetDto) UnsetPurchaseReceiptId()`

UnsetPurchaseReceiptId ensures that no value is present for PurchaseReceiptId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


