# EmployeeProfileDtoCollectionQueryParameters

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

### NewEmployeeProfileDtoCollectionQueryParameters

`func NewEmployeeProfileDtoCollectionQueryParameters() *EmployeeProfileDtoCollectionQueryParameters`

NewEmployeeProfileDtoCollectionQueryParameters instantiates a new EmployeeProfileDtoCollectionQueryParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeProfileDtoCollectionQueryParametersWithDefaults

`func NewEmployeeProfileDtoCollectionQueryParametersWithDefaults() *EmployeeProfileDtoCollectionQueryParameters`

NewEmployeeProfileDtoCollectionQueryParametersWithDefaults instantiates a new EmployeeProfileDtoCollectionQueryParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTop

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetTop() int32`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetTopOk() (*int32, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetTop(v int32)`

SetTop sets Top field to given value.

### HasTop

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTopNil

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetTopNil(b bool)`

 SetTopNil sets the value for Top to be an explicit nil

### UnsetTop
`func (o *EmployeeProfileDtoCollectionQueryParameters) UnsetTop()`

UnsetTop ensures that no value is present for Top, not even an explicit nil
### GetSkip

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### SetSkipNil

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetSkipNil(b bool)`

 SetSkipNil sets the value for Skip to be an explicit nil

### UnsetSkip
`func (o *EmployeeProfileDtoCollectionQueryParameters) UnsetSkip()`

UnsetSkip ensures that no value is present for Skip, not even an explicit nil
### GetCount

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFilter

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *EmployeeProfileDtoCollectionQueryParameters) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetOrderBy

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *EmployeeProfileDtoCollectionQueryParameters) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetSearch

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *EmployeeProfileDtoCollectionQueryParameters) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetSelect

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetSelect() string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetSelectOk() (*string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetSelect(v string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### SetSelectNil

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *EmployeeProfileDtoCollectionQueryParameters) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetExpand

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### SetExpandNil

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetExpandNil(b bool)`

 SetExpandNil sets the value for Expand to be an explicit nil

### UnsetExpand
`func (o *EmployeeProfileDtoCollectionQueryParameters) UnsetExpand()`

UnsetExpand ensures that no value is present for Expand, not even an explicit nil
### GetIsEmpty

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetIsEmpty() bool`

GetIsEmpty returns the IsEmpty field if non-nil, zero value otherwise.

### GetIsEmptyOk

`func (o *EmployeeProfileDtoCollectionQueryParameters) GetIsEmptyOk() (*bool, bool)`

GetIsEmptyOk returns a tuple with the IsEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmpty

`func (o *EmployeeProfileDtoCollectionQueryParameters) SetIsEmpty(v bool)`

SetIsEmpty sets IsEmpty field to given value.

### HasIsEmpty

`func (o *EmployeeProfileDtoCollectionQueryParameters) HasIsEmpty() bool`

HasIsEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


