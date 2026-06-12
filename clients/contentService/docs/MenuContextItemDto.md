# MenuContextItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Text** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Target** | Pointer to **NullableString** |  | [optional] 
**Tooltip** | Pointer to **NullableString** |  | [optional] 
**ParentMenuContextItemId** | Pointer to **NullableString** |  | [optional] 
**MenuContextId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMenuContextItemDto

`func NewMenuContextItemDto() *MenuContextItemDto`

NewMenuContextItemDto instantiates a new MenuContextItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuContextItemDtoWithDefaults

`func NewMenuContextItemDtoWithDefaults() *MenuContextItemDto`

NewMenuContextItemDtoWithDefaults instantiates a new MenuContextItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuContextItemDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuContextItemDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuContextItemDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuContextItemDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MenuContextItemDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MenuContextItemDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *MenuContextItemDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MenuContextItemDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MenuContextItemDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MenuContextItemDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MenuContextItemDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MenuContextItemDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetOrder

`func (o *MenuContextItemDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MenuContextItemDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MenuContextItemDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MenuContextItemDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *MenuContextItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuContextItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuContextItemDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MenuContextItemDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MenuContextItemDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MenuContextItemDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetText

`func (o *MenuContextItemDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MenuContextItemDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MenuContextItemDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MenuContextItemDto) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *MenuContextItemDto) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *MenuContextItemDto) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetUrl

`func (o *MenuContextItemDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MenuContextItemDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MenuContextItemDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MenuContextItemDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MenuContextItemDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MenuContextItemDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetIcon

`func (o *MenuContextItemDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *MenuContextItemDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *MenuContextItemDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *MenuContextItemDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *MenuContextItemDto) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *MenuContextItemDto) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetTarget

`func (o *MenuContextItemDto) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MenuContextItemDto) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MenuContextItemDto) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MenuContextItemDto) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MenuContextItemDto) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MenuContextItemDto) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetTooltip

`func (o *MenuContextItemDto) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *MenuContextItemDto) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *MenuContextItemDto) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *MenuContextItemDto) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### SetTooltipNil

`func (o *MenuContextItemDto) SetTooltipNil(b bool)`

 SetTooltipNil sets the value for Tooltip to be an explicit nil

### UnsetTooltip
`func (o *MenuContextItemDto) UnsetTooltip()`

UnsetTooltip ensures that no value is present for Tooltip, not even an explicit nil
### GetParentMenuContextItemId

`func (o *MenuContextItemDto) GetParentMenuContextItemId() string`

GetParentMenuContextItemId returns the ParentMenuContextItemId field if non-nil, zero value otherwise.

### GetParentMenuContextItemIdOk

`func (o *MenuContextItemDto) GetParentMenuContextItemIdOk() (*string, bool)`

GetParentMenuContextItemIdOk returns a tuple with the ParentMenuContextItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMenuContextItemId

`func (o *MenuContextItemDto) SetParentMenuContextItemId(v string)`

SetParentMenuContextItemId sets ParentMenuContextItemId field to given value.

### HasParentMenuContextItemId

`func (o *MenuContextItemDto) HasParentMenuContextItemId() bool`

HasParentMenuContextItemId returns a boolean if a field has been set.

### SetParentMenuContextItemIdNil

`func (o *MenuContextItemDto) SetParentMenuContextItemIdNil(b bool)`

 SetParentMenuContextItemIdNil sets the value for ParentMenuContextItemId to be an explicit nil

### UnsetParentMenuContextItemId
`func (o *MenuContextItemDto) UnsetParentMenuContextItemId()`

UnsetParentMenuContextItemId ensures that no value is present for ParentMenuContextItemId, not even an explicit nil
### GetMenuContextId

`func (o *MenuContextItemDto) GetMenuContextId() string`

GetMenuContextId returns the MenuContextId field if non-nil, zero value otherwise.

### GetMenuContextIdOk

`func (o *MenuContextItemDto) GetMenuContextIdOk() (*string, bool)`

GetMenuContextIdOk returns a tuple with the MenuContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuContextId

`func (o *MenuContextItemDto) SetMenuContextId(v string)`

SetMenuContextId sets MenuContextId field to given value.

### HasMenuContextId

`func (o *MenuContextItemDto) HasMenuContextId() bool`

HasMenuContextId returns a boolean if a field has been set.

### SetMenuContextIdNil

`func (o *MenuContextItemDto) SetMenuContextIdNil(b bool)`

 SetMenuContextIdNil sets the value for MenuContextId to be an explicit nil

### UnsetMenuContextId
`func (o *MenuContextItemDto) UnsetMenuContextId()`

UnsetMenuContextId ensures that no value is present for MenuContextId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


