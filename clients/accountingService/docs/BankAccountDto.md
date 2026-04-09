# BankAccountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Group** | Pointer to **bool** |  | [optional] 
**Frozen** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Balance** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to **NullableString** |  | [optional] 
**AccountTypeId** | Pointer to **NullableString** |  | [optional] 
**DebitsBalance** | Pointer to **float64** |  | [optional] 
**CreditsBalance** | Pointer to **float64** |  | [optional] 
**ParentAccountId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ChildrenAccountsCount** | Pointer to **int32** |  | [optional] 
**AccountCategory** | Pointer to **string** |  | [optional] 
**BalanceAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**CreditsBalanceAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**DebitsBalanceAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Swift** | Pointer to **NullableString** |  | [optional] 
**BranchCode** | Pointer to **NullableString** |  | [optional] 
**BankAccountNumber** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**BankId** | Pointer to **NullableString** |  | [optional] 
**BankProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankAccountDto

`func NewBankAccountDto() *BankAccountDto`

NewBankAccountDto instantiates a new BankAccountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountDtoWithDefaults

`func NewBankAccountDtoWithDefaults() *BankAccountDto`

NewBankAccountDtoWithDefaults instantiates a new BankAccountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankAccountDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankAccountDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankAccountDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankAccountDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BankAccountDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BankAccountDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BankAccountDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankAccountDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankAccountDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BankAccountDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BankAccountDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BankAccountDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetGroup

`func (o *BankAccountDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BankAccountDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BankAccountDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BankAccountDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetFrozen

`func (o *BankAccountDto) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *BankAccountDto) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *BankAccountDto) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *BankAccountDto) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetName

`func (o *BankAccountDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankAccountDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankAccountDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BankAccountDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BankAccountDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BankAccountDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *BankAccountDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BankAccountDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BankAccountDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BankAccountDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *BankAccountDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BankAccountDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPath

`func (o *BankAccountDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BankAccountDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BankAccountDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BankAccountDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BankAccountDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BankAccountDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTitle

`func (o *BankAccountDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BankAccountDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BankAccountDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BankAccountDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BankAccountDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BankAccountDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPrefix

`func (o *BankAccountDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BankAccountDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BankAccountDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *BankAccountDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *BankAccountDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *BankAccountDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetBalance

`func (o *BankAccountDto) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BankAccountDto) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BankAccountDto) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BankAccountDto) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrencyId

`func (o *BankAccountDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankAccountDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankAccountDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankAccountDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BankAccountDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BankAccountDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetAccountType

`func (o *BankAccountDto) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountDto) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountDto) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankAccountDto) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *BankAccountDto) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *BankAccountDto) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetAccountTypeId

`func (o *BankAccountDto) GetAccountTypeId() string`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *BankAccountDto) GetAccountTypeIdOk() (*string, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *BankAccountDto) SetAccountTypeId(v string)`

SetAccountTypeId sets AccountTypeId field to given value.

### HasAccountTypeId

`func (o *BankAccountDto) HasAccountTypeId() bool`

HasAccountTypeId returns a boolean if a field has been set.

### SetAccountTypeIdNil

`func (o *BankAccountDto) SetAccountTypeIdNil(b bool)`

 SetAccountTypeIdNil sets the value for AccountTypeId to be an explicit nil

### UnsetAccountTypeId
`func (o *BankAccountDto) UnsetAccountTypeId()`

UnsetAccountTypeId ensures that no value is present for AccountTypeId, not even an explicit nil
### GetDebitsBalance

`func (o *BankAccountDto) GetDebitsBalance() float64`

GetDebitsBalance returns the DebitsBalance field if non-nil, zero value otherwise.

### GetDebitsBalanceOk

`func (o *BankAccountDto) GetDebitsBalanceOk() (*float64, bool)`

GetDebitsBalanceOk returns a tuple with the DebitsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitsBalance

`func (o *BankAccountDto) SetDebitsBalance(v float64)`

SetDebitsBalance sets DebitsBalance field to given value.

### HasDebitsBalance

`func (o *BankAccountDto) HasDebitsBalance() bool`

HasDebitsBalance returns a boolean if a field has been set.

### GetCreditsBalance

`func (o *BankAccountDto) GetCreditsBalance() float64`

GetCreditsBalance returns the CreditsBalance field if non-nil, zero value otherwise.

### GetCreditsBalanceOk

`func (o *BankAccountDto) GetCreditsBalanceOk() (*float64, bool)`

GetCreditsBalanceOk returns a tuple with the CreditsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalance

`func (o *BankAccountDto) SetCreditsBalance(v float64)`

SetCreditsBalance sets CreditsBalance field to given value.

### HasCreditsBalance

`func (o *BankAccountDto) HasCreditsBalance() bool`

HasCreditsBalance returns a boolean if a field has been set.

### GetParentAccountId

`func (o *BankAccountDto) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *BankAccountDto) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *BankAccountDto) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *BankAccountDto) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### SetParentAccountIdNil

`func (o *BankAccountDto) SetParentAccountIdNil(b bool)`

 SetParentAccountIdNil sets the value for ParentAccountId to be an explicit nil

### UnsetParentAccountId
`func (o *BankAccountDto) UnsetParentAccountId()`

UnsetParentAccountId ensures that no value is present for ParentAccountId, not even an explicit nil
### GetTenantId

`func (o *BankAccountDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BankAccountDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BankAccountDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BankAccountDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BankAccountDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BankAccountDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *BankAccountDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BankAccountDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BankAccountDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BankAccountDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BankAccountDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BankAccountDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetChildrenAccountsCount

`func (o *BankAccountDto) GetChildrenAccountsCount() int32`

GetChildrenAccountsCount returns the ChildrenAccountsCount field if non-nil, zero value otherwise.

### GetChildrenAccountsCountOk

`func (o *BankAccountDto) GetChildrenAccountsCountOk() (*int32, bool)`

GetChildrenAccountsCountOk returns a tuple with the ChildrenAccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAccountsCount

`func (o *BankAccountDto) SetChildrenAccountsCount(v int32)`

SetChildrenAccountsCount sets ChildrenAccountsCount field to given value.

### HasChildrenAccountsCount

`func (o *BankAccountDto) HasChildrenAccountsCount() bool`

HasChildrenAccountsCount returns a boolean if a field has been set.

### GetAccountCategory

`func (o *BankAccountDto) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *BankAccountDto) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *BankAccountDto) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.

### HasAccountCategory

`func (o *BankAccountDto) HasAccountCategory() bool`

HasAccountCategory returns a boolean if a field has been set.

### GetBalanceAmount

`func (o *BankAccountDto) GetBalanceAmount() Money`

GetBalanceAmount returns the BalanceAmount field if non-nil, zero value otherwise.

### GetBalanceAmountOk

`func (o *BankAccountDto) GetBalanceAmountOk() (*Money, bool)`

GetBalanceAmountOk returns a tuple with the BalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmount

`func (o *BankAccountDto) SetBalanceAmount(v Money)`

SetBalanceAmount sets BalanceAmount field to given value.

### HasBalanceAmount

`func (o *BankAccountDto) HasBalanceAmount() bool`

HasBalanceAmount returns a boolean if a field has been set.

### GetCreditsBalanceAmount

`func (o *BankAccountDto) GetCreditsBalanceAmount() Money`

GetCreditsBalanceAmount returns the CreditsBalanceAmount field if non-nil, zero value otherwise.

### GetCreditsBalanceAmountOk

`func (o *BankAccountDto) GetCreditsBalanceAmountOk() (*Money, bool)`

GetCreditsBalanceAmountOk returns a tuple with the CreditsBalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalanceAmount

`func (o *BankAccountDto) SetCreditsBalanceAmount(v Money)`

SetCreditsBalanceAmount sets CreditsBalanceAmount field to given value.

### HasCreditsBalanceAmount

`func (o *BankAccountDto) HasCreditsBalanceAmount() bool`

HasCreditsBalanceAmount returns a boolean if a field has been set.

### GetDebitsBalanceAmount

`func (o *BankAccountDto) GetDebitsBalanceAmount() Money`

GetDebitsBalanceAmount returns the DebitsBalanceAmount field if non-nil, zero value otherwise.

### GetDebitsBalanceAmountOk

`func (o *BankAccountDto) GetDebitsBalanceAmountOk() (*Money, bool)`

GetDebitsBalanceAmountOk returns a tuple with the DebitsBalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitsBalanceAmount

`func (o *BankAccountDto) SetDebitsBalanceAmount(v Money)`

SetDebitsBalanceAmount sets DebitsBalanceAmount field to given value.

### HasDebitsBalanceAmount

`func (o *BankAccountDto) HasDebitsBalanceAmount() bool`

HasDebitsBalanceAmount returns a boolean if a field has been set.

### GetIban

`func (o *BankAccountDto) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountDto) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountDto) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *BankAccountDto) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *BankAccountDto) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *BankAccountDto) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetSwift

`func (o *BankAccountDto) GetSwift() string`

GetSwift returns the Swift field if non-nil, zero value otherwise.

### GetSwiftOk

`func (o *BankAccountDto) GetSwiftOk() (*string, bool)`

GetSwiftOk returns a tuple with the Swift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwift

`func (o *BankAccountDto) SetSwift(v string)`

SetSwift sets Swift field to given value.

### HasSwift

`func (o *BankAccountDto) HasSwift() bool`

HasSwift returns a boolean if a field has been set.

### SetSwiftNil

`func (o *BankAccountDto) SetSwiftNil(b bool)`

 SetSwiftNil sets the value for Swift to be an explicit nil

### UnsetSwift
`func (o *BankAccountDto) UnsetSwift()`

UnsetSwift ensures that no value is present for Swift, not even an explicit nil
### GetBranchCode

`func (o *BankAccountDto) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *BankAccountDto) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *BankAccountDto) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *BankAccountDto) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### SetBranchCodeNil

`func (o *BankAccountDto) SetBranchCodeNil(b bool)`

 SetBranchCodeNil sets the value for BranchCode to be an explicit nil

### UnsetBranchCode
`func (o *BankAccountDto) UnsetBranchCode()`

UnsetBranchCode ensures that no value is present for BranchCode, not even an explicit nil
### GetBankAccountNumber

`func (o *BankAccountDto) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *BankAccountDto) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *BankAccountDto) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *BankAccountDto) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### SetBankAccountNumberNil

`func (o *BankAccountDto) SetBankAccountNumberNil(b bool)`

 SetBankAccountNumberNil sets the value for BankAccountNumber to be an explicit nil

### UnsetBankAccountNumber
`func (o *BankAccountDto) UnsetBankAccountNumber()`

UnsetBankAccountNumber ensures that no value is present for BankAccountNumber, not even an explicit nil
### GetQualifiedName

`func (o *BankAccountDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *BankAccountDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *BankAccountDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *BankAccountDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *BankAccountDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *BankAccountDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetBankId

`func (o *BankAccountDto) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *BankAccountDto) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *BankAccountDto) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *BankAccountDto) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### SetBankIdNil

`func (o *BankAccountDto) SetBankIdNil(b bool)`

 SetBankIdNil sets the value for BankId to be an explicit nil

### UnsetBankId
`func (o *BankAccountDto) UnsetBankId()`

UnsetBankId ensures that no value is present for BankId, not even an explicit nil
### GetBankProfileId

`func (o *BankAccountDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankAccountDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankAccountDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankAccountDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankAccountDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankAccountDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


