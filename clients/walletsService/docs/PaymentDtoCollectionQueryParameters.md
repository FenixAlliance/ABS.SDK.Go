# PaymentDtoCollectionQueryParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Top** | Pointer to **NullableInt32** |  | [optional] 
**Skip** | Pointer to **NullableInt32** |  | [optional] 
**Count** | Pointer to **bool** |  | [optional] 
**Filter** | Pointer to **NullableString** |  | [optional] 
**OrderBy** | Pointer to **NullableString** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**Select** | Pointer to **NullableString** |  | [optional] 
**Expand** | Pointer to **NullableString** |  | [optional] 
**IsEmpty** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewPaymentDtoCollectionQueryParameters

`func NewPaymentDtoCollectionQueryParameters() *PaymentDtoCollectionQueryParameters`

NewPaymentDtoCollectionQueryParameters instantiates a new PaymentDtoCollectionQueryParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDtoCollectionQueryParametersWithDefaults

`func NewPaymentDtoCollectionQueryParametersWithDefaults() *PaymentDtoCollectionQueryParameters`

NewPaymentDtoCollectionQueryParametersWithDefaults instantiates a new PaymentDtoCollectionQueryParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTop

`func (o *PaymentDtoCollectionQueryParameters) GetTop() int32`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *PaymentDtoCollectionQueryParameters) GetTopOk() (*int32, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *PaymentDtoCollectionQueryParameters) SetTop(v int32)`

SetTop sets Top field to given value.

### HasTop

`func (o *PaymentDtoCollectionQueryParameters) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTopNil

`func (o *PaymentDtoCollectionQueryParameters) SetTopNil(b bool)`

 SetTopNil sets the value for Top to be an explicit nil

### UnsetTop
`func (o *PaymentDtoCollectionQueryParameters) UnsetTop()`

UnsetTop ensures that no value is present for Top, not even an explicit nil
### GetSkip

`func (o *PaymentDtoCollectionQueryParameters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PaymentDtoCollectionQueryParameters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PaymentDtoCollectionQueryParameters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PaymentDtoCollectionQueryParameters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### SetSkipNil

`func (o *PaymentDtoCollectionQueryParameters) SetSkipNil(b bool)`

 SetSkipNil sets the value for Skip to be an explicit nil

### UnsetSkip
`func (o *PaymentDtoCollectionQueryParameters) UnsetSkip()`

UnsetSkip ensures that no value is present for Skip, not even an explicit nil
### GetCount

`func (o *PaymentDtoCollectionQueryParameters) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaymentDtoCollectionQueryParameters) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaymentDtoCollectionQueryParameters) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaymentDtoCollectionQueryParameters) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFilter

`func (o *PaymentDtoCollectionQueryParameters) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PaymentDtoCollectionQueryParameters) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PaymentDtoCollectionQueryParameters) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PaymentDtoCollectionQueryParameters) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *PaymentDtoCollectionQueryParameters) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *PaymentDtoCollectionQueryParameters) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetOrderBy

`func (o *PaymentDtoCollectionQueryParameters) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *PaymentDtoCollectionQueryParameters) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *PaymentDtoCollectionQueryParameters) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *PaymentDtoCollectionQueryParameters) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *PaymentDtoCollectionQueryParameters) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *PaymentDtoCollectionQueryParameters) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetSearch

`func (o *PaymentDtoCollectionQueryParameters) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *PaymentDtoCollectionQueryParameters) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *PaymentDtoCollectionQueryParameters) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *PaymentDtoCollectionQueryParameters) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *PaymentDtoCollectionQueryParameters) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *PaymentDtoCollectionQueryParameters) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetSelect

`func (o *PaymentDtoCollectionQueryParameters) GetSelect() string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *PaymentDtoCollectionQueryParameters) GetSelectOk() (*string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *PaymentDtoCollectionQueryParameters) SetSelect(v string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *PaymentDtoCollectionQueryParameters) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### SetSelectNil

`func (o *PaymentDtoCollectionQueryParameters) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *PaymentDtoCollectionQueryParameters) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetExpand

`func (o *PaymentDtoCollectionQueryParameters) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *PaymentDtoCollectionQueryParameters) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *PaymentDtoCollectionQueryParameters) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *PaymentDtoCollectionQueryParameters) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### SetExpandNil

`func (o *PaymentDtoCollectionQueryParameters) SetExpandNil(b bool)`

 SetExpandNil sets the value for Expand to be an explicit nil

### UnsetExpand
`func (o *PaymentDtoCollectionQueryParameters) UnsetExpand()`

UnsetExpand ensures that no value is present for Expand, not even an explicit nil
### GetIsEmpty

`func (o *PaymentDtoCollectionQueryParameters) GetIsEmpty() bool`

GetIsEmpty returns the IsEmpty field if non-nil, zero value otherwise.

### GetIsEmptyOk

`func (o *PaymentDtoCollectionQueryParameters) GetIsEmptyOk() (*bool, bool)`

GetIsEmptyOk returns a tuple with the IsEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmpty

`func (o *PaymentDtoCollectionQueryParameters) SetIsEmpty(v bool)`

SetIsEmpty sets IsEmpty field to given value.

### HasIsEmpty

`func (o *PaymentDtoCollectionQueryParameters) HasIsEmpty() bool`

HasIsEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


