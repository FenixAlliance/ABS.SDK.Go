# GoogleIntegrationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**GoogleMaps** | Pointer to [**GoogleMapsIntegrationOptions**](GoogleMapsIntegrationOptions.md) |  | [optional] 
**GoogleMerchantCenter** | Pointer to [**GoogleMerchantCenterIntegrationOptions**](GoogleMerchantCenterIntegrationOptions.md) |  | [optional] 
**GoogleTagManager** | Pointer to [**GoogleTagManagerIntegrationOptions**](GoogleTagManagerIntegrationOptions.md) |  | [optional] 
**GoogleRecaptcha** | Pointer to [**GoogleRecaptchaIntegrationOptions**](GoogleRecaptchaIntegrationOptions.md) |  | [optional] 
**GoogleAnalytics** | Pointer to [**GoogleAnalytics**](GoogleAnalytics.md) |  | [optional] 
**GoogleMyBusiness** | Pointer to [**GoogleMyBusinessIntegrationOptions**](GoogleMyBusinessIntegrationOptions.md) |  | [optional] 

## Methods

### NewGoogleIntegrationOptions

`func NewGoogleIntegrationOptions() *GoogleIntegrationOptions`

NewGoogleIntegrationOptions instantiates a new GoogleIntegrationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleIntegrationOptionsWithDefaults

`func NewGoogleIntegrationOptionsWithDefaults() *GoogleIntegrationOptions`

NewGoogleIntegrationOptionsWithDefaults instantiates a new GoogleIntegrationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GoogleIntegrationOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GoogleIntegrationOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GoogleIntegrationOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GoogleIntegrationOptions) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGoogleMaps

`func (o *GoogleIntegrationOptions) GetGoogleMaps() GoogleMapsIntegrationOptions`

GetGoogleMaps returns the GoogleMaps field if non-nil, zero value otherwise.

### GetGoogleMapsOk

`func (o *GoogleIntegrationOptions) GetGoogleMapsOk() (*GoogleMapsIntegrationOptions, bool)`

GetGoogleMapsOk returns a tuple with the GoogleMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleMaps

`func (o *GoogleIntegrationOptions) SetGoogleMaps(v GoogleMapsIntegrationOptions)`

SetGoogleMaps sets GoogleMaps field to given value.

### HasGoogleMaps

`func (o *GoogleIntegrationOptions) HasGoogleMaps() bool`

HasGoogleMaps returns a boolean if a field has been set.

### GetGoogleMerchantCenter

`func (o *GoogleIntegrationOptions) GetGoogleMerchantCenter() GoogleMerchantCenterIntegrationOptions`

GetGoogleMerchantCenter returns the GoogleMerchantCenter field if non-nil, zero value otherwise.

### GetGoogleMerchantCenterOk

`func (o *GoogleIntegrationOptions) GetGoogleMerchantCenterOk() (*GoogleMerchantCenterIntegrationOptions, bool)`

GetGoogleMerchantCenterOk returns a tuple with the GoogleMerchantCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleMerchantCenter

`func (o *GoogleIntegrationOptions) SetGoogleMerchantCenter(v GoogleMerchantCenterIntegrationOptions)`

SetGoogleMerchantCenter sets GoogleMerchantCenter field to given value.

### HasGoogleMerchantCenter

`func (o *GoogleIntegrationOptions) HasGoogleMerchantCenter() bool`

HasGoogleMerchantCenter returns a boolean if a field has been set.

### GetGoogleTagManager

`func (o *GoogleIntegrationOptions) GetGoogleTagManager() GoogleTagManagerIntegrationOptions`

GetGoogleTagManager returns the GoogleTagManager field if non-nil, zero value otherwise.

### GetGoogleTagManagerOk

`func (o *GoogleIntegrationOptions) GetGoogleTagManagerOk() (*GoogleTagManagerIntegrationOptions, bool)`

GetGoogleTagManagerOk returns a tuple with the GoogleTagManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTagManager

`func (o *GoogleIntegrationOptions) SetGoogleTagManager(v GoogleTagManagerIntegrationOptions)`

SetGoogleTagManager sets GoogleTagManager field to given value.

### HasGoogleTagManager

`func (o *GoogleIntegrationOptions) HasGoogleTagManager() bool`

HasGoogleTagManager returns a boolean if a field has been set.

### GetGoogleRecaptcha

`func (o *GoogleIntegrationOptions) GetGoogleRecaptcha() GoogleRecaptchaIntegrationOptions`

GetGoogleRecaptcha returns the GoogleRecaptcha field if non-nil, zero value otherwise.

### GetGoogleRecaptchaOk

`func (o *GoogleIntegrationOptions) GetGoogleRecaptchaOk() (*GoogleRecaptchaIntegrationOptions, bool)`

GetGoogleRecaptchaOk returns a tuple with the GoogleRecaptcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleRecaptcha

`func (o *GoogleIntegrationOptions) SetGoogleRecaptcha(v GoogleRecaptchaIntegrationOptions)`

SetGoogleRecaptcha sets GoogleRecaptcha field to given value.

### HasGoogleRecaptcha

`func (o *GoogleIntegrationOptions) HasGoogleRecaptcha() bool`

HasGoogleRecaptcha returns a boolean if a field has been set.

### GetGoogleAnalytics

`func (o *GoogleIntegrationOptions) GetGoogleAnalytics() GoogleAnalytics`

GetGoogleAnalytics returns the GoogleAnalytics field if non-nil, zero value otherwise.

### GetGoogleAnalyticsOk

`func (o *GoogleIntegrationOptions) GetGoogleAnalyticsOk() (*GoogleAnalytics, bool)`

GetGoogleAnalyticsOk returns a tuple with the GoogleAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalytics

`func (o *GoogleIntegrationOptions) SetGoogleAnalytics(v GoogleAnalytics)`

SetGoogleAnalytics sets GoogleAnalytics field to given value.

### HasGoogleAnalytics

`func (o *GoogleIntegrationOptions) HasGoogleAnalytics() bool`

HasGoogleAnalytics returns a boolean if a field has been set.

### GetGoogleMyBusiness

`func (o *GoogleIntegrationOptions) GetGoogleMyBusiness() GoogleMyBusinessIntegrationOptions`

GetGoogleMyBusiness returns the GoogleMyBusiness field if non-nil, zero value otherwise.

### GetGoogleMyBusinessOk

`func (o *GoogleIntegrationOptions) GetGoogleMyBusinessOk() (*GoogleMyBusinessIntegrationOptions, bool)`

GetGoogleMyBusinessOk returns a tuple with the GoogleMyBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleMyBusiness

`func (o *GoogleIntegrationOptions) SetGoogleMyBusiness(v GoogleMyBusinessIntegrationOptions)`

SetGoogleMyBusiness sets GoogleMyBusiness field to given value.

### HasGoogleMyBusiness

`func (o *GoogleIntegrationOptions) HasGoogleMyBusiness() bool`

HasGoogleMyBusiness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


