# ModuleGrantDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**GrantedAtUtc** | Pointer to **time.Time** |  | [optional] 
**GrantedBy** | Pointer to **NullableString** |  | [optional] 
**Note** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModuleGrantDto

`func NewModuleGrantDto() *ModuleGrantDto`

NewModuleGrantDto instantiates a new ModuleGrantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleGrantDtoWithDefaults

`func NewModuleGrantDtoWithDefaults() *ModuleGrantDto`

NewModuleGrantDtoWithDefaults instantiates a new ModuleGrantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *ModuleGrantDto) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ModuleGrantDto) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ModuleGrantDto) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ModuleGrantDto) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *ModuleGrantDto) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *ModuleGrantDto) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetExpiresAt

`func (o *ModuleGrantDto) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ModuleGrantDto) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ModuleGrantDto) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ModuleGrantDto) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ModuleGrantDto) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ModuleGrantDto) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetGrantedAtUtc

`func (o *ModuleGrantDto) GetGrantedAtUtc() time.Time`

GetGrantedAtUtc returns the GrantedAtUtc field if non-nil, zero value otherwise.

### GetGrantedAtUtcOk

`func (o *ModuleGrantDto) GetGrantedAtUtcOk() (*time.Time, bool)`

GetGrantedAtUtcOk returns a tuple with the GrantedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAtUtc

`func (o *ModuleGrantDto) SetGrantedAtUtc(v time.Time)`

SetGrantedAtUtc sets GrantedAtUtc field to given value.

### HasGrantedAtUtc

`func (o *ModuleGrantDto) HasGrantedAtUtc() bool`

HasGrantedAtUtc returns a boolean if a field has been set.

### GetGrantedBy

`func (o *ModuleGrantDto) GetGrantedBy() string`

GetGrantedBy returns the GrantedBy field if non-nil, zero value otherwise.

### GetGrantedByOk

`func (o *ModuleGrantDto) GetGrantedByOk() (*string, bool)`

GetGrantedByOk returns a tuple with the GrantedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedBy

`func (o *ModuleGrantDto) SetGrantedBy(v string)`

SetGrantedBy sets GrantedBy field to given value.

### HasGrantedBy

`func (o *ModuleGrantDto) HasGrantedBy() bool`

HasGrantedBy returns a boolean if a field has been set.

### SetGrantedByNil

`func (o *ModuleGrantDto) SetGrantedByNil(b bool)`

 SetGrantedByNil sets the value for GrantedBy to be an explicit nil

### UnsetGrantedBy
`func (o *ModuleGrantDto) UnsetGrantedBy()`

UnsetGrantedBy ensures that no value is present for GrantedBy, not even an explicit nil
### GetNote

`func (o *ModuleGrantDto) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ModuleGrantDto) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ModuleGrantDto) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ModuleGrantDto) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ModuleGrantDto) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ModuleGrantDto) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


