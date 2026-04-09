# CartOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableGuestCart** | Pointer to **bool** |  | [optional] 
**ProductPlaceholderImage** | Pointer to **NullableString** |  | [optional] 
**RedirectToCartPageAfterAddingProducts** | Pointer to **bool** |  | [optional] 

## Methods

### NewCartOptions

`func NewCartOptions() *CartOptions`

NewCartOptions instantiates a new CartOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartOptionsWithDefaults

`func NewCartOptionsWithDefaults() *CartOptions`

NewCartOptionsWithDefaults instantiates a new CartOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableGuestCart

`func (o *CartOptions) GetEnableGuestCart() bool`

GetEnableGuestCart returns the EnableGuestCart field if non-nil, zero value otherwise.

### GetEnableGuestCartOk

`func (o *CartOptions) GetEnableGuestCartOk() (*bool, bool)`

GetEnableGuestCartOk returns a tuple with the EnableGuestCart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGuestCart

`func (o *CartOptions) SetEnableGuestCart(v bool)`

SetEnableGuestCart sets EnableGuestCart field to given value.

### HasEnableGuestCart

`func (o *CartOptions) HasEnableGuestCart() bool`

HasEnableGuestCart returns a boolean if a field has been set.

### GetProductPlaceholderImage

`func (o *CartOptions) GetProductPlaceholderImage() string`

GetProductPlaceholderImage returns the ProductPlaceholderImage field if non-nil, zero value otherwise.

### GetProductPlaceholderImageOk

`func (o *CartOptions) GetProductPlaceholderImageOk() (*string, bool)`

GetProductPlaceholderImageOk returns a tuple with the ProductPlaceholderImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlaceholderImage

`func (o *CartOptions) SetProductPlaceholderImage(v string)`

SetProductPlaceholderImage sets ProductPlaceholderImage field to given value.

### HasProductPlaceholderImage

`func (o *CartOptions) HasProductPlaceholderImage() bool`

HasProductPlaceholderImage returns a boolean if a field has been set.

### SetProductPlaceholderImageNil

`func (o *CartOptions) SetProductPlaceholderImageNil(b bool)`

 SetProductPlaceholderImageNil sets the value for ProductPlaceholderImage to be an explicit nil

### UnsetProductPlaceholderImage
`func (o *CartOptions) UnsetProductPlaceholderImage()`

UnsetProductPlaceholderImage ensures that no value is present for ProductPlaceholderImage, not even an explicit nil
### GetRedirectToCartPageAfterAddingProducts

`func (o *CartOptions) GetRedirectToCartPageAfterAddingProducts() bool`

GetRedirectToCartPageAfterAddingProducts returns the RedirectToCartPageAfterAddingProducts field if non-nil, zero value otherwise.

### GetRedirectToCartPageAfterAddingProductsOk

`func (o *CartOptions) GetRedirectToCartPageAfterAddingProductsOk() (*bool, bool)`

GetRedirectToCartPageAfterAddingProductsOk returns a tuple with the RedirectToCartPageAfterAddingProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToCartPageAfterAddingProducts

`func (o *CartOptions) SetRedirectToCartPageAfterAddingProducts(v bool)`

SetRedirectToCartPageAfterAddingProducts sets RedirectToCartPageAfterAddingProducts field to given value.

### HasRedirectToCartPageAfterAddingProducts

`func (o *CartOptions) HasRedirectToCartPageAfterAddingProducts() bool`

HasRedirectToCartPageAfterAddingProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


