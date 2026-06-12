# BlockchainBlockDtoODataQueryOptions

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
**IfMatch** | Pointer to [**BlockchainBlockDtoETag**](BlockchainBlockDtoETag.md) |  | [optional] 
**IfNoneMatch** | Pointer to [**BlockchainBlockDtoETag**](BlockchainBlockDtoETag.md) |  | [optional] 

## Methods

### NewBlockchainBlockDtoODataQueryOptions

`func NewBlockchainBlockDtoODataQueryOptions() *BlockchainBlockDtoODataQueryOptions`

NewBlockchainBlockDtoODataQueryOptions instantiates a new BlockchainBlockDtoODataQueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainBlockDtoODataQueryOptionsWithDefaults

`func NewBlockchainBlockDtoODataQueryOptionsWithDefaults() *BlockchainBlockDtoODataQueryOptions`

NewBlockchainBlockDtoODataQueryOptionsWithDefaults instantiates a new BlockchainBlockDtoODataQueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *BlockchainBlockDtoODataQueryOptions) GetRequest() HttpRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetRequestOk() (*HttpRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *BlockchainBlockDtoODataQueryOptions) SetRequest(v HttpRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *BlockchainBlockDtoODataQueryOptions) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetContext

`func (o *BlockchainBlockDtoODataQueryOptions) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BlockchainBlockDtoODataQueryOptions) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *BlockchainBlockDtoODataQueryOptions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawValues

`func (o *BlockchainBlockDtoODataQueryOptions) GetRawValues() ODataRawQueryOptions`

GetRawValues returns the RawValues field if non-nil, zero value otherwise.

### GetRawValuesOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetRawValuesOk() (*ODataRawQueryOptions, bool)`

GetRawValuesOk returns a tuple with the RawValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValues

`func (o *BlockchainBlockDtoODataQueryOptions) SetRawValues(v ODataRawQueryOptions)`

SetRawValues sets RawValues field to given value.

### HasRawValues

`func (o *BlockchainBlockDtoODataQueryOptions) HasRawValues() bool`

HasRawValues returns a boolean if a field has been set.

### GetSelectExpand

`func (o *BlockchainBlockDtoODataQueryOptions) GetSelectExpand() SelectExpandQueryOption`

GetSelectExpand returns the SelectExpand field if non-nil, zero value otherwise.

### GetSelectExpandOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetSelectExpandOk() (*SelectExpandQueryOption, bool)`

GetSelectExpandOk returns a tuple with the SelectExpand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectExpand

`func (o *BlockchainBlockDtoODataQueryOptions) SetSelectExpand(v SelectExpandQueryOption)`

SetSelectExpand sets SelectExpand field to given value.

### HasSelectExpand

`func (o *BlockchainBlockDtoODataQueryOptions) HasSelectExpand() bool`

HasSelectExpand returns a boolean if a field has been set.

### GetApply

`func (o *BlockchainBlockDtoODataQueryOptions) GetApply() ApplyQueryOption`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetApplyOk() (*ApplyQueryOption, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *BlockchainBlockDtoODataQueryOptions) SetApply(v ApplyQueryOption)`

SetApply sets Apply field to given value.

### HasApply

`func (o *BlockchainBlockDtoODataQueryOptions) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetCompute

`func (o *BlockchainBlockDtoODataQueryOptions) GetCompute() ComputeQueryOption`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetComputeOk() (*ComputeQueryOption, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *BlockchainBlockDtoODataQueryOptions) SetCompute(v ComputeQueryOption)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *BlockchainBlockDtoODataQueryOptions) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetFilter

`func (o *BlockchainBlockDtoODataQueryOptions) GetFilter() FilterQueryOption`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetFilterOk() (*FilterQueryOption, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BlockchainBlockDtoODataQueryOptions) SetFilter(v FilterQueryOption)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BlockchainBlockDtoODataQueryOptions) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSearch

`func (o *BlockchainBlockDtoODataQueryOptions) GetSearch() SearchQueryOption`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetSearchOk() (*SearchQueryOption, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *BlockchainBlockDtoODataQueryOptions) SetSearch(v SearchQueryOption)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *BlockchainBlockDtoODataQueryOptions) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetOrderBy

`func (o *BlockchainBlockDtoODataQueryOptions) GetOrderBy() OrderByQueryOption`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetOrderByOk() (*OrderByQueryOption, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *BlockchainBlockDtoODataQueryOptions) SetOrderBy(v OrderByQueryOption)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *BlockchainBlockDtoODataQueryOptions) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSkip

`func (o *BlockchainBlockDtoODataQueryOptions) GetSkip() SkipQueryOption`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetSkipOk() (*SkipQueryOption, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *BlockchainBlockDtoODataQueryOptions) SetSkip(v SkipQueryOption)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *BlockchainBlockDtoODataQueryOptions) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSkipToken

`func (o *BlockchainBlockDtoODataQueryOptions) GetSkipToken() SkipTokenQueryOption`

GetSkipToken returns the SkipToken field if non-nil, zero value otherwise.

### GetSkipTokenOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetSkipTokenOk() (*SkipTokenQueryOption, bool)`

GetSkipTokenOk returns a tuple with the SkipToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipToken

`func (o *BlockchainBlockDtoODataQueryOptions) SetSkipToken(v SkipTokenQueryOption)`

SetSkipToken sets SkipToken field to given value.

### HasSkipToken

`func (o *BlockchainBlockDtoODataQueryOptions) HasSkipToken() bool`

HasSkipToken returns a boolean if a field has been set.

### GetTop

`func (o *BlockchainBlockDtoODataQueryOptions) GetTop() TopQueryOption`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetTopOk() (*TopQueryOption, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *BlockchainBlockDtoODataQueryOptions) SetTop(v TopQueryOption)`

SetTop sets Top field to given value.

### HasTop

`func (o *BlockchainBlockDtoODataQueryOptions) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetCount

`func (o *BlockchainBlockDtoODataQueryOptions) GetCount() CountQueryOption`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetCountOk() (*CountQueryOption, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BlockchainBlockDtoODataQueryOptions) SetCount(v CountQueryOption)`

SetCount sets Count field to given value.

### HasCount

`func (o *BlockchainBlockDtoODataQueryOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetValidator

`func (o *BlockchainBlockDtoODataQueryOptions) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BlockchainBlockDtoODataQueryOptions) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *BlockchainBlockDtoODataQueryOptions) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetIfMatch

`func (o *BlockchainBlockDtoODataQueryOptions) GetIfMatch() BlockchainBlockDtoETag`

GetIfMatch returns the IfMatch field if non-nil, zero value otherwise.

### GetIfMatchOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetIfMatchOk() (*BlockchainBlockDtoETag, bool)`

GetIfMatchOk returns a tuple with the IfMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfMatch

`func (o *BlockchainBlockDtoODataQueryOptions) SetIfMatch(v BlockchainBlockDtoETag)`

SetIfMatch sets IfMatch field to given value.

### HasIfMatch

`func (o *BlockchainBlockDtoODataQueryOptions) HasIfMatch() bool`

HasIfMatch returns a boolean if a field has been set.

### GetIfNoneMatch

`func (o *BlockchainBlockDtoODataQueryOptions) GetIfNoneMatch() BlockchainBlockDtoETag`

GetIfNoneMatch returns the IfNoneMatch field if non-nil, zero value otherwise.

### GetIfNoneMatchOk

`func (o *BlockchainBlockDtoODataQueryOptions) GetIfNoneMatchOk() (*BlockchainBlockDtoETag, bool)`

GetIfNoneMatchOk returns a tuple with the IfNoneMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfNoneMatch

`func (o *BlockchainBlockDtoODataQueryOptions) SetIfNoneMatch(v BlockchainBlockDtoETag)`

SetIfNoneMatch sets IfNoneMatch field to given value.

### HasIfNoneMatch

`func (o *BlockchainBlockDtoODataQueryOptions) HasIfNoneMatch() bool`

HasIfNoneMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


