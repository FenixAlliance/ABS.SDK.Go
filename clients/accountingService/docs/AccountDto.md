# AccountDto

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
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
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

## Methods

### NewAccountDto

`func NewAccountDto() *AccountDto`

NewAccountDto instantiates a new AccountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDtoWithDefaults

`func NewAccountDtoWithDefaults() *AccountDto`

NewAccountDtoWithDefaults instantiates a new AccountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AccountDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AccountDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AccountDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AccountDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AccountDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetGroup

`func (o *AccountDto) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AccountDto) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AccountDto) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AccountDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetFrozen

`func (o *AccountDto) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *AccountDto) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *AccountDto) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *AccountDto) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetName

`func (o *AccountDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AccountDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccountDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *AccountDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AccountDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AccountDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AccountDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AccountDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AccountDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPath

`func (o *AccountDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AccountDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AccountDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AccountDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AccountDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AccountDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTitle

`func (o *AccountDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccountDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccountDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AccountDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AccountDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AccountDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPrefix

`func (o *AccountDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AccountDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AccountDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AccountDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *AccountDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *AccountDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetBalance

`func (o *AccountDto) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountDto) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountDto) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountDto) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AccountDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AccountDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AccountDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AccountDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AccountDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AccountDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetAccountType

`func (o *AccountDto) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountDto) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountDto) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountDto) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *AccountDto) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *AccountDto) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetQualifiedName

`func (o *AccountDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *AccountDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *AccountDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *AccountDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *AccountDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *AccountDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetAccountTypeId

`func (o *AccountDto) GetAccountTypeId() string`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *AccountDto) GetAccountTypeIdOk() (*string, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *AccountDto) SetAccountTypeId(v string)`

SetAccountTypeId sets AccountTypeId field to given value.

### HasAccountTypeId

`func (o *AccountDto) HasAccountTypeId() bool`

HasAccountTypeId returns a boolean if a field has been set.

### SetAccountTypeIdNil

`func (o *AccountDto) SetAccountTypeIdNil(b bool)`

 SetAccountTypeIdNil sets the value for AccountTypeId to be an explicit nil

### UnsetAccountTypeId
`func (o *AccountDto) UnsetAccountTypeId()`

UnsetAccountTypeId ensures that no value is present for AccountTypeId, not even an explicit nil
### GetDebitsBalance

`func (o *AccountDto) GetDebitsBalance() float64`

GetDebitsBalance returns the DebitsBalance field if non-nil, zero value otherwise.

### GetDebitsBalanceOk

`func (o *AccountDto) GetDebitsBalanceOk() (*float64, bool)`

GetDebitsBalanceOk returns a tuple with the DebitsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitsBalance

`func (o *AccountDto) SetDebitsBalance(v float64)`

SetDebitsBalance sets DebitsBalance field to given value.

### HasDebitsBalance

`func (o *AccountDto) HasDebitsBalance() bool`

HasDebitsBalance returns a boolean if a field has been set.

### GetCreditsBalance

`func (o *AccountDto) GetCreditsBalance() float64`

GetCreditsBalance returns the CreditsBalance field if non-nil, zero value otherwise.

### GetCreditsBalanceOk

`func (o *AccountDto) GetCreditsBalanceOk() (*float64, bool)`

GetCreditsBalanceOk returns a tuple with the CreditsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalance

`func (o *AccountDto) SetCreditsBalance(v float64)`

SetCreditsBalance sets CreditsBalance field to given value.

### HasCreditsBalance

`func (o *AccountDto) HasCreditsBalance() bool`

HasCreditsBalance returns a boolean if a field has been set.

### GetParentAccountId

`func (o *AccountDto) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *AccountDto) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *AccountDto) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *AccountDto) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### SetParentAccountIdNil

`func (o *AccountDto) SetParentAccountIdNil(b bool)`

 SetParentAccountIdNil sets the value for ParentAccountId to be an explicit nil

### UnsetParentAccountId
`func (o *AccountDto) UnsetParentAccountId()`

UnsetParentAccountId ensures that no value is present for ParentAccountId, not even an explicit nil
### GetTenantId

`func (o *AccountDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccountDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccountDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AccountDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AccountDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AccountDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AccountDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AccountDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AccountDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AccountDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AccountDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AccountDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetChildrenAccountsCount

`func (o *AccountDto) GetChildrenAccountsCount() int32`

GetChildrenAccountsCount returns the ChildrenAccountsCount field if non-nil, zero value otherwise.

### GetChildrenAccountsCountOk

`func (o *AccountDto) GetChildrenAccountsCountOk() (*int32, bool)`

GetChildrenAccountsCountOk returns a tuple with the ChildrenAccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAccountsCount

`func (o *AccountDto) SetChildrenAccountsCount(v int32)`

SetChildrenAccountsCount sets ChildrenAccountsCount field to given value.

### HasChildrenAccountsCount

`func (o *AccountDto) HasChildrenAccountsCount() bool`

HasChildrenAccountsCount returns a boolean if a field has been set.

### GetAccountCategory

`func (o *AccountDto) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *AccountDto) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *AccountDto) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.

### HasAccountCategory

`func (o *AccountDto) HasAccountCategory() bool`

HasAccountCategory returns a boolean if a field has been set.

### GetBalanceAmount

`func (o *AccountDto) GetBalanceAmount() Money`

GetBalanceAmount returns the BalanceAmount field if non-nil, zero value otherwise.

### GetBalanceAmountOk

`func (o *AccountDto) GetBalanceAmountOk() (*Money, bool)`

GetBalanceAmountOk returns a tuple with the BalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmount

`func (o *AccountDto) SetBalanceAmount(v Money)`

SetBalanceAmount sets BalanceAmount field to given value.

### HasBalanceAmount

`func (o *AccountDto) HasBalanceAmount() bool`

HasBalanceAmount returns a boolean if a field has been set.

### GetCreditsBalanceAmount

`func (o *AccountDto) GetCreditsBalanceAmount() Money`

GetCreditsBalanceAmount returns the CreditsBalanceAmount field if non-nil, zero value otherwise.

### GetCreditsBalanceAmountOk

`func (o *AccountDto) GetCreditsBalanceAmountOk() (*Money, bool)`

GetCreditsBalanceAmountOk returns a tuple with the CreditsBalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalanceAmount

`func (o *AccountDto) SetCreditsBalanceAmount(v Money)`

SetCreditsBalanceAmount sets CreditsBalanceAmount field to given value.

### HasCreditsBalanceAmount

`func (o *AccountDto) HasCreditsBalanceAmount() bool`

HasCreditsBalanceAmount returns a boolean if a field has been set.

### GetDebitsBalanceAmount

`func (o *AccountDto) GetDebitsBalanceAmount() Money`

GetDebitsBalanceAmount returns the DebitsBalanceAmount field if non-nil, zero value otherwise.

### GetDebitsBalanceAmountOk

`func (o *AccountDto) GetDebitsBalanceAmountOk() (*Money, bool)`

GetDebitsBalanceAmountOk returns a tuple with the DebitsBalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitsBalanceAmount

`func (o *AccountDto) SetDebitsBalanceAmount(v Money)`

SetDebitsBalanceAmount sets DebitsBalanceAmount field to given value.

### HasDebitsBalanceAmount

`func (o *AccountDto) HasDebitsBalanceAmount() bool`

HasDebitsBalanceAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


