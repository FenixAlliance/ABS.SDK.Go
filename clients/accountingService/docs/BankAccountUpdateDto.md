# BankAccountUpdateDto

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
**Iban** | Pointer to **NullableString** |  | [optional] 
**Swift** | Pointer to **NullableString** |  | [optional] 
**BranchCode** | Pointer to **NullableString** |  | [optional] 
**BankAccountNumber** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**BankId** | Pointer to **NullableString** |  | [optional] 
**BankProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankAccountUpdateDto

`func NewBankAccountUpdateDto(name string, currencyId string, ) *BankAccountUpdateDto`

NewBankAccountUpdateDto instantiates a new BankAccountUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountUpdateDtoWithDefaults

`func NewBankAccountUpdateDtoWithDefaults() *BankAccountUpdateDto`

NewBankAccountUpdateDtoWithDefaults instantiates a new BankAccountUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *BankAccountUpdateDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BankAccountUpdateDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BankAccountUpdateDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BankAccountUpdateDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetFrozen

`func (o *BankAccountUpdateDto) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *BankAccountUpdateDto) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *BankAccountUpdateDto) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *BankAccountUpdateDto) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetName

`func (o *BankAccountUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankAccountUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankAccountUpdateDto) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *BankAccountUpdateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BankAccountUpdateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BankAccountUpdateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BankAccountUpdateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *BankAccountUpdateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BankAccountUpdateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPath

`func (o *BankAccountUpdateDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BankAccountUpdateDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BankAccountUpdateDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BankAccountUpdateDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BankAccountUpdateDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BankAccountUpdateDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPrefix

`func (o *BankAccountUpdateDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BankAccountUpdateDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BankAccountUpdateDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *BankAccountUpdateDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *BankAccountUpdateDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *BankAccountUpdateDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetCurrencyId

`func (o *BankAccountUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankAccountUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankAccountUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetAccountTypeId

`func (o *BankAccountUpdateDto) GetAccountTypeId() string`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *BankAccountUpdateDto) GetAccountTypeIdOk() (*string, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *BankAccountUpdateDto) SetAccountTypeId(v string)`

SetAccountTypeId sets AccountTypeId field to given value.

### HasAccountTypeId

`func (o *BankAccountUpdateDto) HasAccountTypeId() bool`

HasAccountTypeId returns a boolean if a field has been set.

### SetAccountTypeIdNil

`func (o *BankAccountUpdateDto) SetAccountTypeIdNil(b bool)`

 SetAccountTypeIdNil sets the value for AccountTypeId to be an explicit nil

### UnsetAccountTypeId
`func (o *BankAccountUpdateDto) UnsetAccountTypeId()`

UnsetAccountTypeId ensures that no value is present for AccountTypeId, not even an explicit nil
### GetParentAccountId

`func (o *BankAccountUpdateDto) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *BankAccountUpdateDto) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *BankAccountUpdateDto) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *BankAccountUpdateDto) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### SetParentAccountIdNil

`func (o *BankAccountUpdateDto) SetParentAccountIdNil(b bool)`

 SetParentAccountIdNil sets the value for ParentAccountId to be an explicit nil

### UnsetParentAccountId
`func (o *BankAccountUpdateDto) UnsetParentAccountId()`

UnsetParentAccountId ensures that no value is present for ParentAccountId, not even an explicit nil
### GetAccountCategory

`func (o *BankAccountUpdateDto) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *BankAccountUpdateDto) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *BankAccountUpdateDto) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.

### HasAccountCategory

`func (o *BankAccountUpdateDto) HasAccountCategory() bool`

HasAccountCategory returns a boolean if a field has been set.

### GetIban

`func (o *BankAccountUpdateDto) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountUpdateDto) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountUpdateDto) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *BankAccountUpdateDto) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *BankAccountUpdateDto) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *BankAccountUpdateDto) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetSwift

`func (o *BankAccountUpdateDto) GetSwift() string`

GetSwift returns the Swift field if non-nil, zero value otherwise.

### GetSwiftOk

`func (o *BankAccountUpdateDto) GetSwiftOk() (*string, bool)`

GetSwiftOk returns a tuple with the Swift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwift

`func (o *BankAccountUpdateDto) SetSwift(v string)`

SetSwift sets Swift field to given value.

### HasSwift

`func (o *BankAccountUpdateDto) HasSwift() bool`

HasSwift returns a boolean if a field has been set.

### SetSwiftNil

`func (o *BankAccountUpdateDto) SetSwiftNil(b bool)`

 SetSwiftNil sets the value for Swift to be an explicit nil

### UnsetSwift
`func (o *BankAccountUpdateDto) UnsetSwift()`

UnsetSwift ensures that no value is present for Swift, not even an explicit nil
### GetBranchCode

`func (o *BankAccountUpdateDto) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *BankAccountUpdateDto) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *BankAccountUpdateDto) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *BankAccountUpdateDto) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### SetBranchCodeNil

`func (o *BankAccountUpdateDto) SetBranchCodeNil(b bool)`

 SetBranchCodeNil sets the value for BranchCode to be an explicit nil

### UnsetBranchCode
`func (o *BankAccountUpdateDto) UnsetBranchCode()`

UnsetBranchCode ensures that no value is present for BranchCode, not even an explicit nil
### GetBankAccountNumber

`func (o *BankAccountUpdateDto) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *BankAccountUpdateDto) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *BankAccountUpdateDto) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *BankAccountUpdateDto) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### SetBankAccountNumberNil

`func (o *BankAccountUpdateDto) SetBankAccountNumberNil(b bool)`

 SetBankAccountNumberNil sets the value for BankAccountNumber to be an explicit nil

### UnsetBankAccountNumber
`func (o *BankAccountUpdateDto) UnsetBankAccountNumber()`

UnsetBankAccountNumber ensures that no value is present for BankAccountNumber, not even an explicit nil
### GetQualifiedName

`func (o *BankAccountUpdateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *BankAccountUpdateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *BankAccountUpdateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *BankAccountUpdateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *BankAccountUpdateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *BankAccountUpdateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetBankId

`func (o *BankAccountUpdateDto) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *BankAccountUpdateDto) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *BankAccountUpdateDto) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *BankAccountUpdateDto) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### SetBankIdNil

`func (o *BankAccountUpdateDto) SetBankIdNil(b bool)`

 SetBankIdNil sets the value for BankId to be an explicit nil

### UnsetBankId
`func (o *BankAccountUpdateDto) UnsetBankId()`

UnsetBankId ensures that no value is present for BankId, not even an explicit nil
### GetBankProfileId

`func (o *BankAccountUpdateDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankAccountUpdateDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankAccountUpdateDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankAccountUpdateDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankAccountUpdateDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankAccountUpdateDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


