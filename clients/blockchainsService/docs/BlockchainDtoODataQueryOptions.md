# BlockchainDtoODataQueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**HttpRequest**](HttpRequest.md) |  | [optional] 
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**RawValues** | Pointer to [**ODataRawQueryOptions**](ODataRawQueryOptions.md) |  | [optional] 
**SelectExpand** | Pointer to [**SelectExpandQueryOption**](SelectExpandQueryOption.md) |  | [optional] 
**Apply** | Pointer to [**ApplyQueryOption**](ApplyQueryOption.md) |  | [optional] 
**Compute** | Pointer to [**ComputeQueryOption**](ComputeQueryOption.md) |  | [optional] 
**Filter** | Pointer to [**FilterQueryOption**](FilterQueryOption.md) |  | [optional] 
**Search** | Pointer to [**SearchQueryOption**](SearchQueryOption.md) |  | [optional] 
**OrderBy** | Pointer to [**OrderByQueryOption**](OrderByQueryOption.md) |  | [optional] 
**Skip** | Pointer to [**SkipQueryOption**](SkipQueryOption.md) |  | [optional] 
**SkipToken** | Pointer to [**SkipTokenQueryOption**](SkipTokenQueryOption.md) |  | [optional] 
**Top** | Pointer to [**TopQueryOption**](TopQueryOption.md) |  | [optional] 
**Count** | Pointer to [**CountQueryOption**](CountQueryOption.md) |  | [optional] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 
**IfMatch** | Pointer to [**BlockchainDtoETag**](BlockchainDtoETag.md) |  | [optional] 
**IfNoneMatch** | Pointer to [**BlockchainDtoETag**](BlockchainDtoETag.md) |  | [optional] 

## Methods

### NewBlockchainDtoODataQueryOptions

`func NewBlockchainDtoODataQueryOptions() *BlockchainDtoODataQueryOptions`

NewBlockchainDtoODataQueryOptions instantiates a new BlockchainDtoODataQueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainDtoODataQueryOptionsWithDefaults

`func NewBlockchainDtoODataQueryOptionsWithDefaults() *BlockchainDtoODataQueryOptions`

NewBlockchainDtoODataQueryOptionsWithDefaults instantiates a new BlockchainDtoODataQueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *BlockchainDtoODataQueryOptions) GetRequest() HttpRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *BlockchainDtoODataQueryOptions) GetRequestOk() (*HttpRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *BlockchainDtoODataQueryOptions) SetRequest(v HttpRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *BlockchainDtoODataQueryOptions) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetContext

`func (o *BlockchainDtoODataQueryOptions) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BlockchainDtoODataQueryOptions) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BlockchainDtoODataQueryOptions) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *BlockchainDtoODataQueryOptions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawValues

`func (o *BlockchainDtoODataQueryOptions) GetRawValues() ODataRawQueryOptions`

GetRawValues returns the RawValues field if non-nil, zero value otherwise.

### GetRawValuesOk

`func (o *BlockchainDtoODataQueryOptions) GetRawValuesOk() (*ODataRawQueryOptions, bool)`

GetRawValuesOk returns a tuple with the RawValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValues

`func (o *BlockchainDtoODataQueryOptions) SetRawValues(v ODataRawQueryOptions)`

SetRawValues sets RawValues field to given value.

### HasRawValues

`func (o *BlockchainDtoODataQueryOptions) HasRawValues() bool`

HasRawValues returns a boolean if a field has been set.

### GetSelectExpand

`func (o *BlockchainDtoODataQueryOptions) GetSelectExpand() SelectExpandQueryOption`

GetSelectExpand returns the SelectExpand field if non-nil, zero value otherwise.

### GetSelectExpandOk

`func (o *BlockchainDtoODataQueryOptions) GetSelectExpandOk() (*SelectExpandQueryOption, bool)`

GetSelectExpandOk returns a tuple with the SelectExpand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectExpand

`func (o *BlockchainDtoODataQueryOptions) SetSelectExpand(v SelectExpandQueryOption)`

SetSelectExpand sets SelectExpand field to given value.

### HasSelectExpand

`func (o *BlockchainDtoODataQueryOptions) HasSelectExpand() bool`

HasSelectExpand returns a boolean if a field has been set.

### GetApply

`func (o *BlockchainDtoODataQueryOptions) GetApply() ApplyQueryOption`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *BlockchainDtoODataQueryOptions) GetApplyOk() (*ApplyQueryOption, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *BlockchainDtoODataQueryOptions) SetApply(v ApplyQueryOption)`

SetApply sets Apply field to given value.

### HasApply

`func (o *BlockchainDtoODataQueryOptions) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetCompute

`func (o *BlockchainDtoODataQueryOptions) GetCompute() ComputeQueryOption`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *BlockchainDtoODataQueryOptions) GetComputeOk() (*ComputeQueryOption, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *BlockchainDtoODataQueryOptions) SetCompute(v ComputeQueryOption)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *BlockchainDtoODataQueryOptions) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetFilter

`func (o *BlockchainDtoODataQueryOptions) GetFilter() FilterQueryOption`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BlockchainDtoODataQueryOptions) GetFilterOk() (*FilterQueryOption, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BlockchainDtoODataQueryOptions) SetFilter(v FilterQueryOption)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BlockchainDtoODataQueryOptions) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSearch

`func (o *BlockchainDtoODataQueryOptions) GetSearch() SearchQueryOption`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *BlockchainDtoODataQueryOptions) GetSearchOk() (*SearchQueryOption, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *BlockchainDtoODataQueryOptions) SetSearch(v SearchQueryOption)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *BlockchainDtoODataQueryOptions) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetOrderBy

`func (o *BlockchainDtoODataQueryOptions) GetOrderBy() OrderByQueryOption`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *BlockchainDtoODataQueryOptions) GetOrderByOk() (*OrderByQueryOption, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *BlockchainDtoODataQueryOptions) SetOrderBy(v OrderByQueryOption)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *BlockchainDtoODataQueryOptions) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSkip

`func (o *BlockchainDtoODataQueryOptions) GetSkip() SkipQueryOption`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *BlockchainDtoODataQueryOptions) GetSkipOk() (*SkipQueryOption, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *BlockchainDtoODataQueryOptions) SetSkip(v SkipQueryOption)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *BlockchainDtoODataQueryOptions) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSkipToken

`func (o *BlockchainDtoODataQueryOptions) GetSkipToken() SkipTokenQueryOption`

GetSkipToken returns the SkipToken field if non-nil, zero value otherwise.

### GetSkipTokenOk

`func (o *BlockchainDtoODataQueryOptions) GetSkipTokenOk() (*SkipTokenQueryOption, bool)`

GetSkipTokenOk returns a tuple with the SkipToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipToken

`func (o *BlockchainDtoODataQueryOptions) SetSkipToken(v SkipTokenQueryOption)`

SetSkipToken sets SkipToken field to given value.

### HasSkipToken

`func (o *BlockchainDtoODataQueryOptions) HasSkipToken() bool`

HasSkipToken returns a boolean if a field has been set.

### GetTop

`func (o *BlockchainDtoODataQueryOptions) GetTop() TopQueryOption`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *BlockchainDtoODataQueryOptions) GetTopOk() (*TopQueryOption, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *BlockchainDtoODataQueryOptions) SetTop(v TopQueryOption)`

SetTop sets Top field to given value.

### HasTop

`func (o *BlockchainDtoODataQueryOptions) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetCount

`func (o *BlockchainDtoODataQueryOptions) GetCount() CountQueryOption`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BlockchainDtoODataQueryOptions) GetCountOk() (*CountQueryOption, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BlockchainDtoODataQueryOptions) SetCount(v CountQueryOption)`

SetCount sets Count field to given value.

### HasCount

`func (o *BlockchainDtoODataQueryOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetValidator

`func (o *BlockchainDtoODataQueryOptions) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BlockchainDtoODataQueryOptions) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BlockchainDtoODataQueryOptions) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *BlockchainDtoODataQueryOptions) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetIfMatch

`func (o *BlockchainDtoODataQueryOptions) GetIfMatch() BlockchainDtoETag`

GetIfMatch returns the IfMatch field if non-nil, zero value otherwise.

### GetIfMatchOk

`func (o *BlockchainDtoODataQueryOptions) GetIfMatchOk() (*BlockchainDtoETag, bool)`

GetIfMatchOk returns a tuple with the IfMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfMatch

`func (o *BlockchainDtoODataQueryOptions) SetIfMatch(v BlockchainDtoETag)`

SetIfMatch sets IfMatch field to given value.

### HasIfMatch

`func (o *BlockchainDtoODataQueryOptions) HasIfMatch() bool`

HasIfMatch returns a boolean if a field has been set.

### GetIfNoneMatch

`func (o *BlockchainDtoODataQueryOptions) GetIfNoneMatch() BlockchainDtoETag`

GetIfNoneMatch returns the IfNoneMatch field if non-nil, zero value otherwise.

### GetIfNoneMatchOk

`func (o *BlockchainDtoODataQueryOptions) GetIfNoneMatchOk() (*BlockchainDtoETag, bool)`

GetIfNoneMatchOk returns a tuple with the IfNoneMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfNoneMatch

`func (o *BlockchainDtoODataQueryOptions) SetIfNoneMatch(v BlockchainDtoETag)`

SetIfNoneMatch sets IfNoneMatch field to given value.

### HasIfNoneMatch

`func (o *BlockchainDtoODataQueryOptions) HasIfNoneMatch() bool`

HasIfNoneMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


