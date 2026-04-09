# AccountUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **bool** |  | [optional] 
**Frozen** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | **string** |  | 
**AccountTypeId** | Pointer to **NullableString** |  | [optional] 
**ParentAccountId** | Pointer to **NullableString** |  | [optional] 
**AccountCategory** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountUpdateDto

`func NewAccountUpdateDto(name string, currencyId string, ) *AccountUpdateDto`

NewAccountUpdateDto instantiates a new AccountUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateDtoWithDefaults

`func NewAccountUpdateDtoWithDefaults() *AccountUpdateDto`

NewAccountUpdateDtoWithDefaults instantiates a new AccountUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *AccountUpdateDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AccountUpdateDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AccountUpdateDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AccountUpdateDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetFrozen

`func (o *AccountUpdateDto) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *AccountUpdateDto) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *AccountUpdateDto) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *AccountUpdateDto) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetName

`func (o *AccountUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountUpdateDto) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AccountUpdateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AccountUpdateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AccountUpdateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AccountUpdateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AccountUpdateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AccountUpdateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPath

`func (o *AccountUpdateDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AccountUpdateDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AccountUpdateDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AccountUpdateDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AccountUpdateDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AccountUpdateDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPrefix

`func (o *AccountUpdateDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AccountUpdateDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AccountUpdateDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AccountUpdateDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *AccountUpdateDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *AccountUpdateDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetCurrencyId

`func (o *AccountUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AccountUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AccountUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetAccountTypeId

`func (o *AccountUpdateDto) GetAccountTypeId() string`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *AccountUpdateDto) GetAccountTypeIdOk() (*string, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *AccountUpdateDto) SetAccountTypeId(v string)`

SetAccountTypeId sets AccountTypeId field to given value.

### HasAccountTypeId

`func (o *AccountUpdateDto) HasAccountTypeId() bool`

HasAccountTypeId returns a boolean if a field has been set.

### SetAccountTypeIdNil

`func (o *AccountUpdateDto) SetAccountTypeIdNil(b bool)`

 SetAccountTypeIdNil sets the value for AccountTypeId to be an explicit nil

### UnsetAccountTypeId
`func (o *AccountUpdateDto) UnsetAccountTypeId()`

UnsetAccountTypeId ensures that no value is present for AccountTypeId, not even an explicit nil
### GetParentAccountId

`func (o *AccountUpdateDto) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *AccountUpdateDto) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *AccountUpdateDto) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *AccountUpdateDto) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### SetParentAccountIdNil

`func (o *AccountUpdateDto) SetParentAccountIdNil(b bool)`

 SetParentAccountIdNil sets the value for ParentAccountId to be an explicit nil

### UnsetParentAccountId
`func (o *AccountUpdateDto) UnsetParentAccountId()`

UnsetParentAccountId ensures that no value is present for ParentAccountId, not even an explicit nil
### GetAccountCategory

`func (o *AccountUpdateDto) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *AccountUpdateDto) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *AccountUpdateDto) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.

### HasAccountCategory

`func (o *AccountUpdateDto) HasAccountCategory() bool`

HasAccountCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


