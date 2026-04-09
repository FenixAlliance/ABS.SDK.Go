# RecommendationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewWeight** | Pointer to **float64** |  | [optional] 
**AddToCartWeight** | Pointer to **float64** |  | [optional] 
**RemovedFromCartWeight** | Pointer to **float64** |  | [optional] 
**AddedToWishlistWeight** | Pointer to **float64** |  | [optional] 
**AlreadyPurchasedWeight** | Pointer to **float64** |  | [optional] 
**RemovedToWishlistWeight** | Pointer to **float64** |  | [optional] 
**AddedToCompareTableWeight** | Pointer to **float64** |  | [optional] 
**RemovedToCompareTableWeight** | Pointer to **float64** |  | [optional] 
**EnableCrossSelling** | Pointer to **bool** |  | [optional] 
**EnableBundledProducts** | Pointer to **bool** |  | [optional] 
**EnableRecentlyViewedProducts** | Pointer to **bool** |  | [optional] 

## Methods

### NewRecommendationOptions

`func NewRecommendationOptions() *RecommendationOptions`

NewRecommendationOptions instantiates a new RecommendationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationOptionsWithDefaults

`func NewRecommendationOptionsWithDefaults() *RecommendationOptions`

NewRecommendationOptionsWithDefaults instantiates a new RecommendationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewWeight

`func (o *RecommendationOptions) GetViewWeight() float64`

GetViewWeight returns the ViewWeight field if non-nil, zero value otherwise.

### GetViewWeightOk

`func (o *RecommendationOptions) GetViewWeightOk() (*float64, bool)`

GetViewWeightOk returns a tuple with the ViewWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewWeight

`func (o *RecommendationOptions) SetViewWeight(v float64)`

SetViewWeight sets ViewWeight field to given value.

### HasViewWeight

`func (o *RecommendationOptions) HasViewWeight() bool`

HasViewWeight returns a boolean if a field has been set.

### GetAddToCartWeight

`func (o *RecommendationOptions) GetAddToCartWeight() float64`

GetAddToCartWeight returns the AddToCartWeight field if non-nil, zero value otherwise.

### GetAddToCartWeightOk

`func (o *RecommendationOptions) GetAddToCartWeightOk() (*float64, bool)`

GetAddToCartWeightOk returns a tuple with the AddToCartWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToCartWeight

`func (o *RecommendationOptions) SetAddToCartWeight(v float64)`

SetAddToCartWeight sets AddToCartWeight field to given value.

### HasAddToCartWeight

`func (o *RecommendationOptions) HasAddToCartWeight() bool`

HasAddToCartWeight returns a boolean if a field has been set.

### GetRemovedFromCartWeight

`func (o *RecommendationOptions) GetRemovedFromCartWeight() float64`

GetRemovedFromCartWeight returns the RemovedFromCartWeight field if non-nil, zero value otherwise.

### GetRemovedFromCartWeightOk

`func (o *RecommendationOptions) GetRemovedFromCartWeightOk() (*float64, bool)`

GetRemovedFromCartWeightOk returns a tuple with the RemovedFromCartWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedFromCartWeight

`func (o *RecommendationOptions) SetRemovedFromCartWeight(v float64)`

SetRemovedFromCartWeight sets RemovedFromCartWeight field to given value.

### HasRemovedFromCartWeight

`func (o *RecommendationOptions) HasRemovedFromCartWeight() bool`

HasRemovedFromCartWeight returns a boolean if a field has been set.

### GetAddedToWishlistWeight

`func (o *RecommendationOptions) GetAddedToWishlistWeight() float64`

GetAddedToWishlistWeight returns the AddedToWishlistWeight field if non-nil, zero value otherwise.

### GetAddedToWishlistWeightOk

`func (o *RecommendationOptions) GetAddedToWishlistWeightOk() (*float64, bool)`

GetAddedToWishlistWeightOk returns a tuple with the AddedToWishlistWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedToWishlistWeight

`func (o *RecommendationOptions) SetAddedToWishlistWeight(v float64)`

SetAddedToWishlistWeight sets AddedToWishlistWeight field to given value.

### HasAddedToWishlistWeight

`func (o *RecommendationOptions) HasAddedToWishlistWeight() bool`

HasAddedToWishlistWeight returns a boolean if a field has been set.

### GetAlreadyPurchasedWeight

`func (o *RecommendationOptions) GetAlreadyPurchasedWeight() float64`

GetAlreadyPurchasedWeight returns the AlreadyPurchasedWeight field if non-nil, zero value otherwise.

### GetAlreadyPurchasedWeightOk

`func (o *RecommendationOptions) GetAlreadyPurchasedWeightOk() (*float64, bool)`

GetAlreadyPurchasedWeightOk returns a tuple with the AlreadyPurchasedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyPurchasedWeight

`func (o *RecommendationOptions) SetAlreadyPurchasedWeight(v float64)`

SetAlreadyPurchasedWeight sets AlreadyPurchasedWeight field to given value.

### HasAlreadyPurchasedWeight

`func (o *RecommendationOptions) HasAlreadyPurchasedWeight() bool`

HasAlreadyPurchasedWeight returns a boolean if a field has been set.

### GetRemovedToWishlistWeight

`func (o *RecommendationOptions) GetRemovedToWishlistWeight() float64`

GetRemovedToWishlistWeight returns the RemovedToWishlistWeight field if non-nil, zero value otherwise.

### GetRemovedToWishlistWeightOk

`func (o *RecommendationOptions) GetRemovedToWishlistWeightOk() (*float64, bool)`

GetRemovedToWishlistWeightOk returns a tuple with the RemovedToWishlistWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedToWishlistWeight

`func (o *RecommendationOptions) SetRemovedToWishlistWeight(v float64)`

SetRemovedToWishlistWeight sets RemovedToWishlistWeight field to given value.

### HasRemovedToWishlistWeight

`func (o *RecommendationOptions) HasRemovedToWishlistWeight() bool`

HasRemovedToWishlistWeight returns a boolean if a field has been set.

### GetAddedToCompareTableWeight

`func (o *RecommendationOptions) GetAddedToCompareTableWeight() float64`

GetAddedToCompareTableWeight returns the AddedToCompareTableWeight field if non-nil, zero value otherwise.

### GetAddedToCompareTableWeightOk

`func (o *RecommendationOptions) GetAddedToCompareTableWeightOk() (*float64, bool)`

GetAddedToCompareTableWeightOk returns a tuple with the AddedToCompareTableWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedToCompareTableWeight

`func (o *RecommendationOptions) SetAddedToCompareTableWeight(v float64)`

SetAddedToCompareTableWeight sets AddedToCompareTableWeight field to given value.

### HasAddedToCompareTableWeight

`func (o *RecommendationOptions) HasAddedToCompareTableWeight() bool`

HasAddedToCompareTableWeight returns a boolean if a field has been set.

### GetRemovedToCompareTableWeight

`func (o *RecommendationOptions) GetRemovedToCompareTableWeight() float64`

GetRemovedToCompareTableWeight returns the RemovedToCompareTableWeight field if non-nil, zero value otherwise.

### GetRemovedToCompareTableWeightOk

`func (o *RecommendationOptions) GetRemovedToCompareTableWeightOk() (*float64, bool)`

GetRemovedToCompareTableWeightOk returns a tuple with the RemovedToCompareTableWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedToCompareTableWeight

`func (o *RecommendationOptions) SetRemovedToCompareTableWeight(v float64)`

SetRemovedToCompareTableWeight sets RemovedToCompareTableWeight field to given value.

### HasRemovedToCompareTableWeight

`func (o *RecommendationOptions) HasRemovedToCompareTableWeight() bool`

HasRemovedToCompareTableWeight returns a boolean if a field has been set.

### GetEnableCrossSelling

`func (o *RecommendationOptions) GetEnableCrossSelling() bool`

GetEnableCrossSelling returns the EnableCrossSelling field if non-nil, zero value otherwise.

### GetEnableCrossSellingOk

`func (o *RecommendationOptions) GetEnableCrossSellingOk() (*bool, bool)`

GetEnableCrossSellingOk returns a tuple with the EnableCrossSelling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCrossSelling

`func (o *RecommendationOptions) SetEnableCrossSelling(v bool)`

SetEnableCrossSelling sets EnableCrossSelling field to given value.

### HasEnableCrossSelling

`func (o *RecommendationOptions) HasEnableCrossSelling() bool`

HasEnableCrossSelling returns a boolean if a field has been set.

### GetEnableBundledProducts

`func (o *RecommendationOptions) GetEnableBundledProducts() bool`

GetEnableBundledProducts returns the EnableBundledProducts field if non-nil, zero value otherwise.

### GetEnableBundledProductsOk

`func (o *RecommendationOptions) GetEnableBundledProductsOk() (*bool, bool)`

GetEnableBundledProductsOk returns a tuple with the EnableBundledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBundledProducts

`func (o *RecommendationOptions) SetEnableBundledProducts(v bool)`

SetEnableBundledProducts sets EnableBundledProducts field to given value.

### HasEnableBundledProducts

`func (o *RecommendationOptions) HasEnableBundledProducts() bool`

HasEnableBundledProducts returns a boolean if a field has been set.

### GetEnableRecentlyViewedProducts

`func (o *RecommendationOptions) GetEnableRecentlyViewedProducts() bool`

GetEnableRecentlyViewedProducts returns the EnableRecentlyViewedProducts field if non-nil, zero value otherwise.

### GetEnableRecentlyViewedProductsOk

`func (o *RecommendationOptions) GetEnableRecentlyViewedProductsOk() (*bool, bool)`

GetEnableRecentlyViewedProductsOk returns a tuple with the EnableRecentlyViewedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecentlyViewedProducts

`func (o *RecommendationOptions) SetEnableRecentlyViewedProducts(v bool)`

SetEnableRecentlyViewedProducts sets EnableRecentlyViewedProducts field to given value.

### HasEnableRecentlyViewedProducts

`func (o *RecommendationOptions) HasEnableRecentlyViewedProducts() bool`

HasEnableRecentlyViewedProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


