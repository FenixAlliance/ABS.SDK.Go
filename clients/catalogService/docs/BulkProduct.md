# BulkProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**Supplier** | Pointer to **NullableString** |  | [optional] 
**TaxPolicies** | Pointer to **NullableString** |  | [optional] 
**SupplierCode** | Pointer to **NullableString** |  | [optional] 
**GoogleCategory** | Pointer to **NullableString** |  | [optional] 
**ShippingCountry** | Pointer to **NullableString** |  | [optional] 
**RegularPrice** | Pointer to **NullableFloat64** |  | [optional] 
**DiscountPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**DiscountAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CurrentStock** | Pointer to **NullableFloat64** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**InStock** | Pointer to **NullableBool** |  | [optional] 
**OnDiscount** | Pointer to **NullableBool** |  | [optional] 
**ByRequest** | Pointer to **NullableBool** |  | [optional] 
**IsFixedDiscount** | Pointer to **NullableBool** |  | [optional] 
**ManageInventory** | Pointer to **NullableBool** |  | [optional] 
**IsDeadlineDiscount** | Pointer to **NullableBool** |  | [optional] 
**DeadlineDiscountFromDate** | Pointer to **NullableTime** |  | [optional] 
**DeadlineDiscountDueDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBulkProduct

`func NewBulkProduct() *BulkProduct`

NewBulkProduct instantiates a new BulkProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkProductWithDefaults

`func NewBulkProductWithDefaults() *BulkProduct`

NewBulkProductWithDefaults instantiates a new BulkProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BulkProduct) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BulkProduct) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSku

`func (o *BulkProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *BulkProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *BulkProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *BulkProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *BulkProduct) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *BulkProduct) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetTitle

`func (o *BulkProduct) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BulkProduct) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BulkProduct) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BulkProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BulkProduct) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BulkProduct) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *BulkProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BulkProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BulkProduct) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BulkProduct) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetImage

`func (o *BulkProduct) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BulkProduct) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BulkProduct) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BulkProduct) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *BulkProduct) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *BulkProduct) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetBrand

`func (o *BulkProduct) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *BulkProduct) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *BulkProduct) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *BulkProduct) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *BulkProduct) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *BulkProduct) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCurrency

`func (o *BulkProduct) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BulkProduct) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BulkProduct) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BulkProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *BulkProduct) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *BulkProduct) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSupplier

`func (o *BulkProduct) GetSupplier() string`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *BulkProduct) GetSupplierOk() (*string, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *BulkProduct) SetSupplier(v string)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *BulkProduct) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *BulkProduct) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *BulkProduct) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetTaxPolicies

`func (o *BulkProduct) GetTaxPolicies() string`

GetTaxPolicies returns the TaxPolicies field if non-nil, zero value otherwise.

### GetTaxPoliciesOk

`func (o *BulkProduct) GetTaxPoliciesOk() (*string, bool)`

GetTaxPoliciesOk returns a tuple with the TaxPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicies

`func (o *BulkProduct) SetTaxPolicies(v string)`

SetTaxPolicies sets TaxPolicies field to given value.

### HasTaxPolicies

`func (o *BulkProduct) HasTaxPolicies() bool`

HasTaxPolicies returns a boolean if a field has been set.

### SetTaxPoliciesNil

`func (o *BulkProduct) SetTaxPoliciesNil(b bool)`

 SetTaxPoliciesNil sets the value for TaxPolicies to be an explicit nil

### UnsetTaxPolicies
`func (o *BulkProduct) UnsetTaxPolicies()`

UnsetTaxPolicies ensures that no value is present for TaxPolicies, not even an explicit nil
### GetSupplierCode

`func (o *BulkProduct) GetSupplierCode() string`

GetSupplierCode returns the SupplierCode field if non-nil, zero value otherwise.

### GetSupplierCodeOk

`func (o *BulkProduct) GetSupplierCodeOk() (*string, bool)`

GetSupplierCodeOk returns a tuple with the SupplierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierCode

`func (o *BulkProduct) SetSupplierCode(v string)`

SetSupplierCode sets SupplierCode field to given value.

### HasSupplierCode

`func (o *BulkProduct) HasSupplierCode() bool`

HasSupplierCode returns a boolean if a field has been set.

### SetSupplierCodeNil

`func (o *BulkProduct) SetSupplierCodeNil(b bool)`

 SetSupplierCodeNil sets the value for SupplierCode to be an explicit nil

### UnsetSupplierCode
`func (o *BulkProduct) UnsetSupplierCode()`

UnsetSupplierCode ensures that no value is present for SupplierCode, not even an explicit nil
### GetGoogleCategory

`func (o *BulkProduct) GetGoogleCategory() string`

GetGoogleCategory returns the GoogleCategory field if non-nil, zero value otherwise.

### GetGoogleCategoryOk

`func (o *BulkProduct) GetGoogleCategoryOk() (*string, bool)`

GetGoogleCategoryOk returns a tuple with the GoogleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCategory

`func (o *BulkProduct) SetGoogleCategory(v string)`

SetGoogleCategory sets GoogleCategory field to given value.

### HasGoogleCategory

`func (o *BulkProduct) HasGoogleCategory() bool`

HasGoogleCategory returns a boolean if a field has been set.

### SetGoogleCategoryNil

`func (o *BulkProduct) SetGoogleCategoryNil(b bool)`

 SetGoogleCategoryNil sets the value for GoogleCategory to be an explicit nil

### UnsetGoogleCategory
`func (o *BulkProduct) UnsetGoogleCategory()`

UnsetGoogleCategory ensures that no value is present for GoogleCategory, not even an explicit nil
### GetShippingCountry

`func (o *BulkProduct) GetShippingCountry() string`

GetShippingCountry returns the ShippingCountry field if non-nil, zero value otherwise.

### GetShippingCountryOk

`func (o *BulkProduct) GetShippingCountryOk() (*string, bool)`

GetShippingCountryOk returns a tuple with the ShippingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountry

`func (o *BulkProduct) SetShippingCountry(v string)`

SetShippingCountry sets ShippingCountry field to given value.

### HasShippingCountry

`func (o *BulkProduct) HasShippingCountry() bool`

HasShippingCountry returns a boolean if a field has been set.

### SetShippingCountryNil

`func (o *BulkProduct) SetShippingCountryNil(b bool)`

 SetShippingCountryNil sets the value for ShippingCountry to be an explicit nil

### UnsetShippingCountry
`func (o *BulkProduct) UnsetShippingCountry()`

UnsetShippingCountry ensures that no value is present for ShippingCountry, not even an explicit nil
### GetRegularPrice

`func (o *BulkProduct) GetRegularPrice() float64`

GetRegularPrice returns the RegularPrice field if non-nil, zero value otherwise.

### GetRegularPriceOk

`func (o *BulkProduct) GetRegularPriceOk() (*float64, bool)`

GetRegularPriceOk returns a tuple with the RegularPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPrice

`func (o *BulkProduct) SetRegularPrice(v float64)`

SetRegularPrice sets RegularPrice field to given value.

### HasRegularPrice

`func (o *BulkProduct) HasRegularPrice() bool`

HasRegularPrice returns a boolean if a field has been set.

### SetRegularPriceNil

`func (o *BulkProduct) SetRegularPriceNil(b bool)`

 SetRegularPriceNil sets the value for RegularPrice to be an explicit nil

### UnsetRegularPrice
`func (o *BulkProduct) UnsetRegularPrice()`

UnsetRegularPrice ensures that no value is present for RegularPrice, not even an explicit nil
### GetDiscountPercentage

`func (o *BulkProduct) GetDiscountPercentage() float64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *BulkProduct) GetDiscountPercentageOk() (*float64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *BulkProduct) SetDiscountPercentage(v float64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *BulkProduct) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### SetDiscountPercentageNil

`func (o *BulkProduct) SetDiscountPercentageNil(b bool)`

 SetDiscountPercentageNil sets the value for DiscountPercentage to be an explicit nil

### UnsetDiscountPercentage
`func (o *BulkProduct) UnsetDiscountPercentage()`

UnsetDiscountPercentage ensures that no value is present for DiscountPercentage, not even an explicit nil
### GetDiscountAmount

`func (o *BulkProduct) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *BulkProduct) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *BulkProduct) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *BulkProduct) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### SetDiscountAmountNil

`func (o *BulkProduct) SetDiscountAmountNil(b bool)`

 SetDiscountAmountNil sets the value for DiscountAmount to be an explicit nil

### UnsetDiscountAmount
`func (o *BulkProduct) UnsetDiscountAmount()`

UnsetDiscountAmount ensures that no value is present for DiscountAmount, not even an explicit nil
### GetCurrentStock

`func (o *BulkProduct) GetCurrentStock() float64`

GetCurrentStock returns the CurrentStock field if non-nil, zero value otherwise.

### GetCurrentStockOk

`func (o *BulkProduct) GetCurrentStockOk() (*float64, bool)`

GetCurrentStockOk returns a tuple with the CurrentStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStock

`func (o *BulkProduct) SetCurrentStock(v float64)`

SetCurrentStock sets CurrentStock field to given value.

### HasCurrentStock

`func (o *BulkProduct) HasCurrentStock() bool`

HasCurrentStock returns a boolean if a field has been set.

### SetCurrentStockNil

`func (o *BulkProduct) SetCurrentStockNil(b bool)`

 SetCurrentStockNil sets the value for CurrentStock to be an explicit nil

### UnsetCurrentStock
`func (o *BulkProduct) UnsetCurrentStock()`

UnsetCurrentStock ensures that no value is present for CurrentStock, not even an explicit nil
### GetTaxable

`func (o *BulkProduct) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *BulkProduct) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *BulkProduct) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *BulkProduct) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *BulkProduct) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *BulkProduct) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetInStock

`func (o *BulkProduct) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *BulkProduct) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *BulkProduct) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *BulkProduct) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### SetInStockNil

`func (o *BulkProduct) SetInStockNil(b bool)`

 SetInStockNil sets the value for InStock to be an explicit nil

### UnsetInStock
`func (o *BulkProduct) UnsetInStock()`

UnsetInStock ensures that no value is present for InStock, not even an explicit nil
### GetOnDiscount

`func (o *BulkProduct) GetOnDiscount() bool`

GetOnDiscount returns the OnDiscount field if non-nil, zero value otherwise.

### GetOnDiscountOk

`func (o *BulkProduct) GetOnDiscountOk() (*bool, bool)`

GetOnDiscountOk returns a tuple with the OnDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDiscount

`func (o *BulkProduct) SetOnDiscount(v bool)`

SetOnDiscount sets OnDiscount field to given value.

### HasOnDiscount

`func (o *BulkProduct) HasOnDiscount() bool`

HasOnDiscount returns a boolean if a field has been set.

### SetOnDiscountNil

`func (o *BulkProduct) SetOnDiscountNil(b bool)`

 SetOnDiscountNil sets the value for OnDiscount to be an explicit nil

### UnsetOnDiscount
`func (o *BulkProduct) UnsetOnDiscount()`

UnsetOnDiscount ensures that no value is present for OnDiscount, not even an explicit nil
### GetByRequest

`func (o *BulkProduct) GetByRequest() bool`

GetByRequest returns the ByRequest field if non-nil, zero value otherwise.

### GetByRequestOk

`func (o *BulkProduct) GetByRequestOk() (*bool, bool)`

GetByRequestOk returns a tuple with the ByRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByRequest

`func (o *BulkProduct) SetByRequest(v bool)`

SetByRequest sets ByRequest field to given value.

### HasByRequest

`func (o *BulkProduct) HasByRequest() bool`

HasByRequest returns a boolean if a field has been set.

### SetByRequestNil

`func (o *BulkProduct) SetByRequestNil(b bool)`

 SetByRequestNil sets the value for ByRequest to be an explicit nil

### UnsetByRequest
`func (o *BulkProduct) UnsetByRequest()`

UnsetByRequest ensures that no value is present for ByRequest, not even an explicit nil
### GetIsFixedDiscount

`func (o *BulkProduct) GetIsFixedDiscount() bool`

GetIsFixedDiscount returns the IsFixedDiscount field if non-nil, zero value otherwise.

### GetIsFixedDiscountOk

`func (o *BulkProduct) GetIsFixedDiscountOk() (*bool, bool)`

GetIsFixedDiscountOk returns a tuple with the IsFixedDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixedDiscount

`func (o *BulkProduct) SetIsFixedDiscount(v bool)`

SetIsFixedDiscount sets IsFixedDiscount field to given value.

### HasIsFixedDiscount

`func (o *BulkProduct) HasIsFixedDiscount() bool`

HasIsFixedDiscount returns a boolean if a field has been set.

### SetIsFixedDiscountNil

`func (o *BulkProduct) SetIsFixedDiscountNil(b bool)`

 SetIsFixedDiscountNil sets the value for IsFixedDiscount to be an explicit nil

### UnsetIsFixedDiscount
`func (o *BulkProduct) UnsetIsFixedDiscount()`

UnsetIsFixedDiscount ensures that no value is present for IsFixedDiscount, not even an explicit nil
### GetManageInventory

`func (o *BulkProduct) GetManageInventory() bool`

GetManageInventory returns the ManageInventory field if non-nil, zero value otherwise.

### GetManageInventoryOk

`func (o *BulkProduct) GetManageInventoryOk() (*bool, bool)`

GetManageInventoryOk returns a tuple with the ManageInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageInventory

`func (o *BulkProduct) SetManageInventory(v bool)`

SetManageInventory sets ManageInventory field to given value.

### HasManageInventory

`func (o *BulkProduct) HasManageInventory() bool`

HasManageInventory returns a boolean if a field has been set.

### SetManageInventoryNil

`func (o *BulkProduct) SetManageInventoryNil(b bool)`

 SetManageInventoryNil sets the value for ManageInventory to be an explicit nil

### UnsetManageInventory
`func (o *BulkProduct) UnsetManageInventory()`

UnsetManageInventory ensures that no value is present for ManageInventory, not even an explicit nil
### GetIsDeadlineDiscount

`func (o *BulkProduct) GetIsDeadlineDiscount() bool`

GetIsDeadlineDiscount returns the IsDeadlineDiscount field if non-nil, zero value otherwise.

### GetIsDeadlineDiscountOk

`func (o *BulkProduct) GetIsDeadlineDiscountOk() (*bool, bool)`

GetIsDeadlineDiscountOk returns a tuple with the IsDeadlineDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeadlineDiscount

`func (o *BulkProduct) SetIsDeadlineDiscount(v bool)`

SetIsDeadlineDiscount sets IsDeadlineDiscount field to given value.

### HasIsDeadlineDiscount

`func (o *BulkProduct) HasIsDeadlineDiscount() bool`

HasIsDeadlineDiscount returns a boolean if a field has been set.

### SetIsDeadlineDiscountNil

`func (o *BulkProduct) SetIsDeadlineDiscountNil(b bool)`

 SetIsDeadlineDiscountNil sets the value for IsDeadlineDiscount to be an explicit nil

### UnsetIsDeadlineDiscount
`func (o *BulkProduct) UnsetIsDeadlineDiscount()`

UnsetIsDeadlineDiscount ensures that no value is present for IsDeadlineDiscount, not even an explicit nil
### GetDeadlineDiscountFromDate

`func (o *BulkProduct) GetDeadlineDiscountFromDate() time.Time`

GetDeadlineDiscountFromDate returns the DeadlineDiscountFromDate field if non-nil, zero value otherwise.

### GetDeadlineDiscountFromDateOk

`func (o *BulkProduct) GetDeadlineDiscountFromDateOk() (*time.Time, bool)`

GetDeadlineDiscountFromDateOk returns a tuple with the DeadlineDiscountFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineDiscountFromDate

`func (o *BulkProduct) SetDeadlineDiscountFromDate(v time.Time)`

SetDeadlineDiscountFromDate sets DeadlineDiscountFromDate field to given value.

### HasDeadlineDiscountFromDate

`func (o *BulkProduct) HasDeadlineDiscountFromDate() bool`

HasDeadlineDiscountFromDate returns a boolean if a field has been set.

### SetDeadlineDiscountFromDateNil

`func (o *BulkProduct) SetDeadlineDiscountFromDateNil(b bool)`

 SetDeadlineDiscountFromDateNil sets the value for DeadlineDiscountFromDate to be an explicit nil

### UnsetDeadlineDiscountFromDate
`func (o *BulkProduct) UnsetDeadlineDiscountFromDate()`

UnsetDeadlineDiscountFromDate ensures that no value is present for DeadlineDiscountFromDate, not even an explicit nil
### GetDeadlineDiscountDueDate

`func (o *BulkProduct) GetDeadlineDiscountDueDate() time.Time`

GetDeadlineDiscountDueDate returns the DeadlineDiscountDueDate field if non-nil, zero value otherwise.

### GetDeadlineDiscountDueDateOk

`func (o *BulkProduct) GetDeadlineDiscountDueDateOk() (*time.Time, bool)`

GetDeadlineDiscountDueDateOk returns a tuple with the DeadlineDiscountDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineDiscountDueDate

`func (o *BulkProduct) SetDeadlineDiscountDueDate(v time.Time)`

SetDeadlineDiscountDueDate sets DeadlineDiscountDueDate field to given value.

### HasDeadlineDiscountDueDate

`func (o *BulkProduct) HasDeadlineDiscountDueDate() bool`

HasDeadlineDiscountDueDate returns a boolean if a field has been set.

### SetDeadlineDiscountDueDateNil

`func (o *BulkProduct) SetDeadlineDiscountDueDateNil(b bool)`

 SetDeadlineDiscountDueDateNil sets the value for DeadlineDiscountDueDate to be an explicit nil

### UnsetDeadlineDiscountDueDate
`func (o *BulkProduct) UnsetDeadlineDiscountDueDate()`

UnsetDeadlineDiscountDueDate ensures that no value is present for DeadlineDiscountDueDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


