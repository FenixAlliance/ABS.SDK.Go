# ChartOfAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**FileUrl** | Pointer to **NullableString** |  | [optional] 
**Childs** | Pointer to [**[]Account**](Account.md) |  | [optional] 

## Methods

### NewChartOfAccounts

`func NewChartOfAccounts() *ChartOfAccounts`

NewChartOfAccounts instantiates a new ChartOfAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartOfAccountsWithDefaults

`func NewChartOfAccountsWithDefaults() *ChartOfAccounts`

NewChartOfAccountsWithDefaults instantiates a new ChartOfAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChartOfAccounts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChartOfAccounts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChartOfAccounts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChartOfAccounts) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ChartOfAccounts) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ChartOfAccounts) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVerified

`func (o *ChartOfAccounts) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ChartOfAccounts) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ChartOfAccounts) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ChartOfAccounts) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetFileUrl

`func (o *ChartOfAccounts) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ChartOfAccounts) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ChartOfAccounts) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *ChartOfAccounts) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *ChartOfAccounts) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *ChartOfAccounts) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetChilds

`func (o *ChartOfAccounts) GetChilds() []Account`

GetChilds returns the Childs field if non-nil, zero value otherwise.

### GetChildsOk

`func (o *ChartOfAccounts) GetChildsOk() (*[]Account, bool)`

GetChildsOk returns a tuple with the Childs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChilds

`func (o *ChartOfAccounts) SetChilds(v []Account)`

SetChilds sets Childs field to given value.

### HasChilds

`func (o *ChartOfAccounts) HasChilds() bool`

HasChilds returns a boolean if a field has been set.

### SetChildsNil

`func (o *ChartOfAccounts) SetChildsNil(b bool)`

 SetChildsNil sets the value for Childs to be an explicit nil

### UnsetChilds
`func (o *ChartOfAccounts) UnsetChilds()`

UnsetChilds ensures that no value is present for Childs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


