# BankAccountCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Group** | Pointer to **bool** |  | [optional] 
**Frozen** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | **string** |  | 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**AccountTypeId** | Pointer to **NullableString** |  | [optional] 
**ParentAccountId** | Pointer to **NullableString** |  | [optional] 
**AccountCategory** | **string** |  | 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Swift** | Pointer to **NullableString** |  | [optional] 
**BranchCode** | Pointer to **NullableString** |  | [optional] 
**BankAccountNumber** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**BankId** | Pointer to **NullableString** |  | [optional] 
**BankProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankAccountCreateDto

`func NewBankAccountCreateDto(name string, currencyId string, accountCategory string, ) *BankAccountCreateDto`

NewBankAccountCreateDto instantiates a new BankAccountCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountCreateDtoWithDefaults

`func NewBankAccountCreateDtoWithDefaults() *BankAccountCreateDto`

NewBankAccountCreateDtoWithDefaults instantiates a new BankAccountCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankAccountCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankAccountCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankAccountCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankAccountCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BankAccountCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankAccountCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankAccountCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BankAccountCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetGroup

`func (o *BankAccountCreateDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BankAccountCreateDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BankAccountCreateDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BankAccountCreateDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetFrozen

`func (o *BankAccountCreateDto) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *BankAccountCreateDto) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *BankAccountCreateDto) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *BankAccountCreateDto) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetName

`func (o *BankAccountCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankAccountCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankAccountCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *BankAccountCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BankAccountCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BankAccountCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BankAccountCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *BankAccountCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BankAccountCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPath

`func (o *BankAccountCreateDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BankAccountCreateDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BankAccountCreateDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BankAccountCreateDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BankAccountCreateDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BankAccountCreateDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPrefix

`func (o *BankAccountCreateDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BankAccountCreateDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BankAccountCreateDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *BankAccountCreateDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *BankAccountCreateDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *BankAccountCreateDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetTenantId

`func (o *BankAccountCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BankAccountCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BankAccountCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BankAccountCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BankAccountCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BankAccountCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCurrencyId

`func (o *BankAccountCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankAccountCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankAccountCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetEnrollmentId

`func (o *BankAccountCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BankAccountCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BankAccountCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BankAccountCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BankAccountCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BankAccountCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAccountTypeId

`func (o *BankAccountCreateDto) GetAccountTypeId() string`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *BankAccountCreateDto) GetAccountTypeIdOk() (*string, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *BankAccountCreateDto) SetAccountTypeId(v string)`

SetAccountTypeId sets AccountTypeId field to given value.

### HasAccountTypeId

`func (o *BankAccountCreateDto) HasAccountTypeId() bool`

HasAccountTypeId returns a boolean if a field has been set.

### SetAccountTypeIdNil

`func (o *BankAccountCreateDto) SetAccountTypeIdNil(b bool)`

 SetAccountTypeIdNil sets the value for AccountTypeId to be an explicit nil

### UnsetAccountTypeId
`func (o *BankAccountCreateDto) UnsetAccountTypeId()`

UnsetAccountTypeId ensures that no value is present for AccountTypeId, not even an explicit nil
### GetParentAccountId

`func (o *BankAccountCreateDto) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *BankAccountCreateDto) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *BankAccountCreateDto) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *BankAccountCreateDto) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### SetParentAccountIdNil

`func (o *BankAccountCreateDto) SetParentAccountIdNil(b bool)`

 SetParentAccountIdNil sets the value for ParentAccountId to be an explicit nil

### UnsetParentAccountId
`func (o *BankAccountCreateDto) UnsetParentAccountId()`

UnsetParentAccountId ensures that no value is present for ParentAccountId, not even an explicit nil
### GetAccountCategory

`func (o *BankAccountCreateDto) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *BankAccountCreateDto) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *BankAccountCreateDto) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.


### GetIban

`func (o *BankAccountCreateDto) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountCreateDto) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountCreateDto) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *BankAccountCreateDto) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *BankAccountCreateDto) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *BankAccountCreateDto) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetSwift

`func (o *BankAccountCreateDto) GetSwift() string`

GetSwift returns the Swift field if non-nil, zero value otherwise.

### GetSwiftOk

`func (o *BankAccountCreateDto) GetSwiftOk() (*string, bool)`

GetSwiftOk returns a tuple with the Swift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwift

`func (o *BankAccountCreateDto) SetSwift(v string)`

SetSwift sets Swift field to given value.

### HasSwift

`func (o *BankAccountCreateDto) HasSwift() bool`

HasSwift returns a boolean if a field has been set.

### SetSwiftNil

`func (o *BankAccountCreateDto) SetSwiftNil(b bool)`

 SetSwiftNil sets the value for Swift to be an explicit nil

### UnsetSwift
`func (o *BankAccountCreateDto) UnsetSwift()`

UnsetSwift ensures that no value is present for Swift, not even an explicit nil
### GetBranchCode

`func (o *BankAccountCreateDto) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *BankAccountCreateDto) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *BankAccountCreateDto) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *BankAccountCreateDto) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### SetBranchCodeNil

`func (o *BankAccountCreateDto) SetBranchCodeNil(b bool)`

 SetBranchCodeNil sets the value for BranchCode to be an explicit nil

### UnsetBranchCode
`func (o *BankAccountCreateDto) UnsetBranchCode()`

UnsetBranchCode ensures that no value is present for BranchCode, not even an explicit nil
### GetBankAccountNumber

`func (o *BankAccountCreateDto) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *BankAccountCreateDto) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *BankAccountCreateDto) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *BankAccountCreateDto) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### SetBankAccountNumberNil

`func (o *BankAccountCreateDto) SetBankAccountNumberNil(b bool)`

 SetBankAccountNumberNil sets the value for BankAccountNumber to be an explicit nil

### UnsetBankAccountNumber
`func (o *BankAccountCreateDto) UnsetBankAccountNumber()`

UnsetBankAccountNumber ensures that no value is present for BankAccountNumber, not even an explicit nil
### GetQualifiedName

`func (o *BankAccountCreateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *BankAccountCreateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *BankAccountCreateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *BankAccountCreateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *BankAccountCreateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *BankAccountCreateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetBankId

`func (o *BankAccountCreateDto) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *BankAccountCreateDto) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *BankAccountCreateDto) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *BankAccountCreateDto) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### SetBankIdNil

`func (o *BankAccountCreateDto) SetBankIdNil(b bool)`

 SetBankIdNil sets the value for BankId to be an explicit nil

### UnsetBankId
`func (o *BankAccountCreateDto) UnsetBankId()`

UnsetBankId ensures that no value is present for BankId, not even an explicit nil
### GetBankProfileId

`func (o *BankAccountCreateDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankAccountCreateDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankAccountCreateDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankAccountCreateDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankAccountCreateDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankAccountCreateDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


