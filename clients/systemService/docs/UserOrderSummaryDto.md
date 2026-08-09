# UserOrderSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewUserOrderSummaryDto

`func NewUserOrderSummaryDto() *UserOrderSummaryDto`

NewUserOrderSummaryDto instantiates a new UserOrderSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOrderSummaryDtoWithDefaults

`func NewUserOrderSummaryDtoWithDefaults() *UserOrderSummaryDto`

NewUserOrderSummaryDtoWithDefaults instantiates a new UserOrderSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserOrderSummaryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserOrderSummaryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserOrderSummaryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserOrderSummaryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserOrderSummaryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserOrderSummaryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOrderType

`func (o *UserOrderSummaryDto) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *UserOrderSummaryDto) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *UserOrderSummaryDto) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *UserOrderSummaryDto) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrderStatus

`func (o *UserOrderSummaryDto) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *UserOrderSummaryDto) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *UserOrderSummaryDto) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *UserOrderSummaryDto) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


