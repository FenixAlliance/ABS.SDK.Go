# WebsiteThemeDtoODataQueryOptions

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
**IfMatch** | Pointer to [**WebsiteThemeDtoETag**](WebsiteThemeDtoETag.md) |  | [optional] 
**IfNoneMatch** | Pointer to [**WebsiteThemeDtoETag**](WebsiteThemeDtoETag.md) |  | [optional] 

## Methods

### NewWebsiteThemeDtoODataQueryOptions

`func NewWebsiteThemeDtoODataQueryOptions() *WebsiteThemeDtoODataQueryOptions`

NewWebsiteThemeDtoODataQueryOptions instantiates a new WebsiteThemeDtoODataQueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteThemeDtoODataQueryOptionsWithDefaults

`func NewWebsiteThemeDtoODataQueryOptionsWithDefaults() *WebsiteThemeDtoODataQueryOptions`

NewWebsiteThemeDtoODataQueryOptionsWithDefaults instantiates a new WebsiteThemeDtoODataQueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *WebsiteThemeDtoODataQueryOptions) GetRequest() HttpRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetRequestOk() (*HttpRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WebsiteThemeDtoODataQueryOptions) SetRequest(v HttpRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *WebsiteThemeDtoODataQueryOptions) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetContext

`func (o *WebsiteThemeDtoODataQueryOptions) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WebsiteThemeDtoODataQueryOptions) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *WebsiteThemeDtoODataQueryOptions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawValues

`func (o *WebsiteThemeDtoODataQueryOptions) GetRawValues() ODataRawQueryOptions`

GetRawValues returns the RawValues field if non-nil, zero value otherwise.

### GetRawValuesOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetRawValuesOk() (*ODataRawQueryOptions, bool)`

GetRawValuesOk returns a tuple with the RawValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValues

`func (o *WebsiteThemeDtoODataQueryOptions) SetRawValues(v ODataRawQueryOptions)`

SetRawValues sets RawValues field to given value.

### HasRawValues

`func (o *WebsiteThemeDtoODataQueryOptions) HasRawValues() bool`

HasRawValues returns a boolean if a field has been set.

### GetSelectExpand

`func (o *WebsiteThemeDtoODataQueryOptions) GetSelectExpand() SelectExpandQueryOption`

GetSelectExpand returns the SelectExpand field if non-nil, zero value otherwise.

### GetSelectExpandOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetSelectExpandOk() (*SelectExpandQueryOption, bool)`

GetSelectExpandOk returns a tuple with the SelectExpand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectExpand

`func (o *WebsiteThemeDtoODataQueryOptions) SetSelectExpand(v SelectExpandQueryOption)`

SetSelectExpand sets SelectExpand field to given value.

### HasSelectExpand

`func (o *WebsiteThemeDtoODataQueryOptions) HasSelectExpand() bool`

HasSelectExpand returns a boolean if a field has been set.

### GetApply

`func (o *WebsiteThemeDtoODataQueryOptions) GetApply() ApplyQueryOption`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetApplyOk() (*ApplyQueryOption, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *WebsiteThemeDtoODataQueryOptions) SetApply(v ApplyQueryOption)`

SetApply sets Apply field to given value.

### HasApply

`func (o *WebsiteThemeDtoODataQueryOptions) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetCompute

`func (o *WebsiteThemeDtoODataQueryOptions) GetCompute() ComputeQueryOption`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetComputeOk() (*ComputeQueryOption, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *WebsiteThemeDtoODataQueryOptions) SetCompute(v ComputeQueryOption)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *WebsiteThemeDtoODataQueryOptions) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetFilter

`func (o *WebsiteThemeDtoODataQueryOptions) GetFilter() FilterQueryOption`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetFilterOk() (*FilterQueryOption, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *WebsiteThemeDtoODataQueryOptions) SetFilter(v FilterQueryOption)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *WebsiteThemeDtoODataQueryOptions) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSearch

`func (o *WebsiteThemeDtoODataQueryOptions) GetSearch() SearchQueryOption`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetSearchOk() (*SearchQueryOption, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *WebsiteThemeDtoODataQueryOptions) SetSearch(v SearchQueryOption)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *WebsiteThemeDtoODataQueryOptions) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetOrderBy

`func (o *WebsiteThemeDtoODataQueryOptions) GetOrderBy() OrderByQueryOption`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetOrderByOk() (*OrderByQueryOption, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *WebsiteThemeDtoODataQueryOptions) SetOrderBy(v OrderByQueryOption)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *WebsiteThemeDtoODataQueryOptions) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSkip

`func (o *WebsiteThemeDtoODataQueryOptions) GetSkip() SkipQueryOption`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetSkipOk() (*SkipQueryOption, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *WebsiteThemeDtoODataQueryOptions) SetSkip(v SkipQueryOption)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *WebsiteThemeDtoODataQueryOptions) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSkipToken

`func (o *WebsiteThemeDtoODataQueryOptions) GetSkipToken() SkipTokenQueryOption`

GetSkipToken returns the SkipToken field if non-nil, zero value otherwise.

### GetSkipTokenOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetSkipTokenOk() (*SkipTokenQueryOption, bool)`

GetSkipTokenOk returns a tuple with the SkipToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipToken

`func (o *WebsiteThemeDtoODataQueryOptions) SetSkipToken(v SkipTokenQueryOption)`

SetSkipToken sets SkipToken field to given value.

### HasSkipToken

`func (o *WebsiteThemeDtoODataQueryOptions) HasSkipToken() bool`

HasSkipToken returns a boolean if a field has been set.

### GetTop

`func (o *WebsiteThemeDtoODataQueryOptions) GetTop() TopQueryOption`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetTopOk() (*TopQueryOption, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *WebsiteThemeDtoODataQueryOptions) SetTop(v TopQueryOption)`

SetTop sets Top field to given value.

### HasTop

`func (o *WebsiteThemeDtoODataQueryOptions) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetCount

`func (o *WebsiteThemeDtoODataQueryOptions) GetCount() CountQueryOption`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetCountOk() (*CountQueryOption, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WebsiteThemeDtoODataQueryOptions) SetCount(v CountQueryOption)`

SetCount sets Count field to given value.

### HasCount

`func (o *WebsiteThemeDtoODataQueryOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetValidator

`func (o *WebsiteThemeDtoODataQueryOptions) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *WebsiteThemeDtoODataQueryOptions) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *WebsiteThemeDtoODataQueryOptions) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetIfMatch

`func (o *WebsiteThemeDtoODataQueryOptions) GetIfMatch() WebsiteThemeDtoETag`

GetIfMatch returns the IfMatch field if non-nil, zero value otherwise.

### GetIfMatchOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetIfMatchOk() (*WebsiteThemeDtoETag, bool)`

GetIfMatchOk returns a tuple with the IfMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfMatch

`func (o *WebsiteThemeDtoODataQueryOptions) SetIfMatch(v WebsiteThemeDtoETag)`

SetIfMatch sets IfMatch field to given value.

### HasIfMatch

`func (o *WebsiteThemeDtoODataQueryOptions) HasIfMatch() bool`

HasIfMatch returns a boolean if a field has been set.

### GetIfNoneMatch

`func (o *WebsiteThemeDtoODataQueryOptions) GetIfNoneMatch() WebsiteThemeDtoETag`

GetIfNoneMatch returns the IfNoneMatch field if non-nil, zero value otherwise.

### GetIfNoneMatchOk

`func (o *WebsiteThemeDtoODataQueryOptions) GetIfNoneMatchOk() (*WebsiteThemeDtoETag, bool)`

GetIfNoneMatchOk returns a tuple with the IfNoneMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfNoneMatch

`func (o *WebsiteThemeDtoODataQueryOptions) SetIfNoneMatch(v WebsiteThemeDtoETag)`

SetIfNoneMatch sets IfNoneMatch field to given value.

### HasIfNoneMatch

`func (o *WebsiteThemeDtoODataQueryOptions) HasIfNoneMatch() bool`

HasIfNoneMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


