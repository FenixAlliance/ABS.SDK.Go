# InventoryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableStockManagement** | Pointer to **bool** |  | [optional] 
**HideOutOfStockProducts** | Pointer to **bool** |  | [optional] 
**EnableLowStockNotifications** | Pointer to **bool** |  | [optional] 
**EnableOutOfStockNotifications** | Pointer to **bool** |  | [optional] 
**StockNotificationRecipients** | Pointer to **NullableString** |  | [optional] 
**HoldStock** | Pointer to **int32** |  | [optional] 
**LowStockThreshold** | Pointer to **int32** |  | [optional] 
**OutOfStockThreshold** | Pointer to **int32** |  | [optional] 
**StockDisplayFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewInventoryOptions

`func NewInventoryOptions() *InventoryOptions`

NewInventoryOptions instantiates a new InventoryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryOptionsWithDefaults

`func NewInventoryOptionsWithDefaults() *InventoryOptions`

NewInventoryOptionsWithDefaults instantiates a new InventoryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableStockManagement

`func (o *InventoryOptions) GetEnableStockManagement() bool`

GetEnableStockManagement returns the EnableStockManagement field if non-nil, zero value otherwise.

### GetEnableStockManagementOk

`func (o *InventoryOptions) GetEnableStockManagementOk() (*bool, bool)`

GetEnableStockManagementOk returns a tuple with the EnableStockManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStockManagement

`func (o *InventoryOptions) SetEnableStockManagement(v bool)`

SetEnableStockManagement sets EnableStockManagement field to given value.

### HasEnableStockManagement

`func (o *InventoryOptions) HasEnableStockManagement() bool`

HasEnableStockManagement returns a boolean if a field has been set.

### GetHideOutOfStockProducts

`func (o *InventoryOptions) GetHideOutOfStockProducts() bool`

GetHideOutOfStockProducts returns the HideOutOfStockProducts field if non-nil, zero value otherwise.

### GetHideOutOfStockProductsOk

`func (o *InventoryOptions) GetHideOutOfStockProductsOk() (*bool, bool)`

GetHideOutOfStockProductsOk returns a tuple with the HideOutOfStockProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideOutOfStockProducts

`func (o *InventoryOptions) SetHideOutOfStockProducts(v bool)`

SetHideOutOfStockProducts sets HideOutOfStockProducts field to given value.

### HasHideOutOfStockProducts

`func (o *InventoryOptions) HasHideOutOfStockProducts() bool`

HasHideOutOfStockProducts returns a boolean if a field has been set.

### GetEnableLowStockNotifications

`func (o *InventoryOptions) GetEnableLowStockNotifications() bool`

GetEnableLowStockNotifications returns the EnableLowStockNotifications field if non-nil, zero value otherwise.

### GetEnableLowStockNotificationsOk

`func (o *InventoryOptions) GetEnableLowStockNotificationsOk() (*bool, bool)`

GetEnableLowStockNotificationsOk returns a tuple with the EnableLowStockNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLowStockNotifications

`func (o *InventoryOptions) SetEnableLowStockNotifications(v bool)`

SetEnableLowStockNotifications sets EnableLowStockNotifications field to given value.

### HasEnableLowStockNotifications

`func (o *InventoryOptions) HasEnableLowStockNotifications() bool`

HasEnableLowStockNotifications returns a boolean if a field has been set.

### GetEnableOutOfStockNotifications

`func (o *InventoryOptions) GetEnableOutOfStockNotifications() bool`

GetEnableOutOfStockNotifications returns the EnableOutOfStockNotifications field if non-nil, zero value otherwise.

### GetEnableOutOfStockNotificationsOk

`func (o *InventoryOptions) GetEnableOutOfStockNotificationsOk() (*bool, bool)`

GetEnableOutOfStockNotificationsOk returns a tuple with the EnableOutOfStockNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOutOfStockNotifications

`func (o *InventoryOptions) SetEnableOutOfStockNotifications(v bool)`

SetEnableOutOfStockNotifications sets EnableOutOfStockNotifications field to given value.

### HasEnableOutOfStockNotifications

`func (o *InventoryOptions) HasEnableOutOfStockNotifications() bool`

HasEnableOutOfStockNotifications returns a boolean if a field has been set.

### GetStockNotificationRecipients

`func (o *InventoryOptions) GetStockNotificationRecipients() string`

GetStockNotificationRecipients returns the StockNotificationRecipients field if non-nil, zero value otherwise.

### GetStockNotificationRecipientsOk

`func (o *InventoryOptions) GetStockNotificationRecipientsOk() (*string, bool)`

GetStockNotificationRecipientsOk returns a tuple with the StockNotificationRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNotificationRecipients

`func (o *InventoryOptions) SetStockNotificationRecipients(v string)`

SetStockNotificationRecipients sets StockNotificationRecipients field to given value.

### HasStockNotificationRecipients

`func (o *InventoryOptions) HasStockNotificationRecipients() bool`

HasStockNotificationRecipients returns a boolean if a field has been set.

### SetStockNotificationRecipientsNil

`func (o *InventoryOptions) SetStockNotificationRecipientsNil(b bool)`

 SetStockNotificationRecipientsNil sets the value for StockNotificationRecipients to be an explicit nil

### UnsetStockNotificationRecipients
`func (o *InventoryOptions) UnsetStockNotificationRecipients()`

UnsetStockNotificationRecipients ensures that no value is present for StockNotificationRecipients, not even an explicit nil
### GetHoldStock

`func (o *InventoryOptions) GetHoldStock() int32`

GetHoldStock returns the HoldStock field if non-nil, zero value otherwise.

### GetHoldStockOk

`func (o *InventoryOptions) GetHoldStockOk() (*int32, bool)`

GetHoldStockOk returns a tuple with the HoldStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldStock

`func (o *InventoryOptions) SetHoldStock(v int32)`

SetHoldStock sets HoldStock field to given value.

### HasHoldStock

`func (o *InventoryOptions) HasHoldStock() bool`

HasHoldStock returns a boolean if a field has been set.

### GetLowStockThreshold

`func (o *InventoryOptions) GetLowStockThreshold() int32`

GetLowStockThreshold returns the LowStockThreshold field if non-nil, zero value otherwise.

### GetLowStockThresholdOk

`func (o *InventoryOptions) GetLowStockThresholdOk() (*int32, bool)`

GetLowStockThresholdOk returns a tuple with the LowStockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowStockThreshold

`func (o *InventoryOptions) SetLowStockThreshold(v int32)`

SetLowStockThreshold sets LowStockThreshold field to given value.

### HasLowStockThreshold

`func (o *InventoryOptions) HasLowStockThreshold() bool`

HasLowStockThreshold returns a boolean if a field has been set.

### GetOutOfStockThreshold

`func (o *InventoryOptions) GetOutOfStockThreshold() int32`

GetOutOfStockThreshold returns the OutOfStockThreshold field if non-nil, zero value otherwise.

### GetOutOfStockThresholdOk

`func (o *InventoryOptions) GetOutOfStockThresholdOk() (*int32, bool)`

GetOutOfStockThresholdOk returns a tuple with the OutOfStockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfStockThreshold

`func (o *InventoryOptions) SetOutOfStockThreshold(v int32)`

SetOutOfStockThreshold sets OutOfStockThreshold field to given value.

### HasOutOfStockThreshold

`func (o *InventoryOptions) HasOutOfStockThreshold() bool`

HasOutOfStockThreshold returns a boolean if a field has been set.

### GetStockDisplayFormat

`func (o *InventoryOptions) GetStockDisplayFormat() string`

GetStockDisplayFormat returns the StockDisplayFormat field if non-nil, zero value otherwise.

### GetStockDisplayFormatOk

`func (o *InventoryOptions) GetStockDisplayFormatOk() (*string, bool)`

GetStockDisplayFormatOk returns a tuple with the StockDisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDisplayFormat

`func (o *InventoryOptions) SetStockDisplayFormat(v string)`

SetStockDisplayFormat sets StockDisplayFormat field to given value.

### HasStockDisplayFormat

`func (o *InventoryOptions) HasStockDisplayFormat() bool`

HasStockDisplayFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


