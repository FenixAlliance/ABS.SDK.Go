# CartDtoCollectionQueryParameters

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

### NewCartDtoCollectionQueryParameters

`func NewCartDtoCollectionQueryParameters() *CartDtoCollectionQueryParameters`

NewCartDtoCollectionQueryParameters instantiates a new CartDtoCollectionQueryParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartDtoCollectionQueryParametersWithDefaults

`func NewCartDtoCollectionQueryParametersWithDefaults() *CartDtoCollectionQueryParameters`

NewCartDtoCollectionQueryParametersWithDefaults instantiates a new CartDtoCollectionQueryParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTop

`func (o *CartDtoCollectionQueryParameters) GetTop() int32`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *CartDtoCollectionQueryParameters) GetTopOk() (*int32, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *CartDtoCollectionQueryParameters) SetTop(v int32)`

SetTop sets Top field to given value.

### HasTop

`func (o *CartDtoCollectionQueryParameters) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTopNil

`func (o *CartDtoCollectionQueryParameters) SetTopNil(b bool)`

 SetTopNil sets the value for Top to be an explicit nil

### UnsetTop
`func (o *CartDtoCollectionQueryParameters) UnsetTop()`

UnsetTop ensures that no value is present for Top, not even an explicit nil
### GetSkip

`func (o *CartDtoCollectionQueryParameters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *CartDtoCollectionQueryParameters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *CartDtoCollectionQueryParameters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *CartDtoCollectionQueryParameters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### SetSkipNil

`func (o *CartDtoCollectionQueryParameters) SetSkipNil(b bool)`

 SetSkipNil sets the value for Skip to be an explicit nil

### UnsetSkip
`func (o *CartDtoCollectionQueryParameters) UnsetSkip()`

UnsetSkip ensures that no value is present for Skip, not even an explicit nil
### GetCount

`func (o *CartDtoCollectionQueryParameters) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CartDtoCollectionQueryParameters) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CartDtoCollectionQueryParameters) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *CartDtoCollectionQueryParameters) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFilter

`func (o *CartDtoCollectionQueryParameters) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CartDtoCollectionQueryParameters) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CartDtoCollectionQueryParameters) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CartDtoCollectionQueryParameters) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *CartDtoCollectionQueryParameters) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CartDtoCollectionQueryParameters) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetOrderBy

`func (o *CartDtoCollectionQueryParameters) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *CartDtoCollectionQueryParameters) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *CartDtoCollectionQueryParameters) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *CartDtoCollectionQueryParameters) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *CartDtoCollectionQueryParameters) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *CartDtoCollectionQueryParameters) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetSearch

`func (o *CartDtoCollectionQueryParameters) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CartDtoCollectionQueryParameters) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CartDtoCollectionQueryParameters) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CartDtoCollectionQueryParameters) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *CartDtoCollectionQueryParameters) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *CartDtoCollectionQueryParameters) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetSelect

`func (o *CartDtoCollectionQueryParameters) GetSelect() string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *CartDtoCollectionQueryParameters) GetSelectOk() (*string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *CartDtoCollectionQueryParameters) SetSelect(v string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *CartDtoCollectionQueryParameters) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### SetSelectNil

`func (o *CartDtoCollectionQueryParameters) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *CartDtoCollectionQueryParameters) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetExpand

`func (o *CartDtoCollectionQueryParameters) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *CartDtoCollectionQueryParameters) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *CartDtoCollectionQueryParameters) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *CartDtoCollectionQueryParameters) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### SetExpandNil

`func (o *CartDtoCollectionQueryParameters) SetExpandNil(b bool)`

 SetExpandNil sets the value for Expand to be an explicit nil

### UnsetExpand
`func (o *CartDtoCollectionQueryParameters) UnsetExpand()`

UnsetExpand ensures that no value is present for Expand, not even an explicit nil
### GetIsEmpty

`func (o *CartDtoCollectionQueryParameters) GetIsEmpty() bool`

GetIsEmpty returns the IsEmpty field if non-nil, zero value otherwise.

### GetIsEmptyOk

`func (o *CartDtoCollectionQueryParameters) GetIsEmptyOk() (*bool, bool)`

GetIsEmptyOk returns a tuple with the IsEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmpty

`func (o *CartDtoCollectionQueryParameters) SetIsEmpty(v bool)`

SetIsEmpty sets IsEmpty field to given value.

### HasIsEmpty

`func (o *CartDtoCollectionQueryParameters) HasIsEmpty() bool`

HasIsEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


