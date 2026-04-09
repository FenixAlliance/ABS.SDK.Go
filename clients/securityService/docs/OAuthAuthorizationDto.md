# OAuthAuthorizationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ConcurrencyToken** | Pointer to **NullableString** |  | [optional] 
**CreationDate** | Pointer to **NullableTime** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**Scopes** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ApplicationId** | Pointer to **NullableString** |  | [optional] 
**TokensCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewOAuthAuthorizationDto

`func NewOAuthAuthorizationDto() *OAuthAuthorizationDto`

NewOAuthAuthorizationDto instantiates a new OAuthAuthorizationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAuthorizationDtoWithDefaults

`func NewOAuthAuthorizationDtoWithDefaults() *OAuthAuthorizationDto`

NewOAuthAuthorizationDtoWithDefaults instantiates a new OAuthAuthorizationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuthAuthorizationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuthAuthorizationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuthAuthorizationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuthAuthorizationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OAuthAuthorizationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OAuthAuthorizationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetConcurrencyToken

`func (o *OAuthAuthorizationDto) GetConcurrencyToken() string`

GetConcurrencyToken returns the ConcurrencyToken field if non-nil, zero value otherwise.

### GetConcurrencyTokenOk

`func (o *OAuthAuthorizationDto) GetConcurrencyTokenOk() (*string, bool)`

GetConcurrencyTokenOk returns a tuple with the ConcurrencyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyToken

`func (o *OAuthAuthorizationDto) SetConcurrencyToken(v string)`

SetConcurrencyToken sets ConcurrencyToken field to given value.

### HasConcurrencyToken

`func (o *OAuthAuthorizationDto) HasConcurrencyToken() bool`

HasConcurrencyToken returns a boolean if a field has been set.

### SetConcurrencyTokenNil

`func (o *OAuthAuthorizationDto) SetConcurrencyTokenNil(b bool)`

 SetConcurrencyTokenNil sets the value for ConcurrencyToken to be an explicit nil

### UnsetConcurrencyToken
`func (o *OAuthAuthorizationDto) UnsetConcurrencyToken()`

UnsetConcurrencyToken ensures that no value is present for ConcurrencyToken, not even an explicit nil
### GetCreationDate

`func (o *OAuthAuthorizationDto) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *OAuthAuthorizationDto) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *OAuthAuthorizationDto) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *OAuthAuthorizationDto) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *OAuthAuthorizationDto) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *OAuthAuthorizationDto) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
### GetProperties

`func (o *OAuthAuthorizationDto) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *OAuthAuthorizationDto) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *OAuthAuthorizationDto) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *OAuthAuthorizationDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *OAuthAuthorizationDto) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *OAuthAuthorizationDto) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetScopes

`func (o *OAuthAuthorizationDto) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuthAuthorizationDto) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuthAuthorizationDto) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuthAuthorizationDto) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *OAuthAuthorizationDto) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *OAuthAuthorizationDto) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetStatus

`func (o *OAuthAuthorizationDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuthAuthorizationDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuthAuthorizationDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuthAuthorizationDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *OAuthAuthorizationDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *OAuthAuthorizationDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubject

`func (o *OAuthAuthorizationDto) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAuthAuthorizationDto) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAuthAuthorizationDto) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OAuthAuthorizationDto) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *OAuthAuthorizationDto) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *OAuthAuthorizationDto) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetType

`func (o *OAuthAuthorizationDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthAuthorizationDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthAuthorizationDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuthAuthorizationDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OAuthAuthorizationDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OAuthAuthorizationDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetApplicationId

`func (o *OAuthAuthorizationDto) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *OAuthAuthorizationDto) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *OAuthAuthorizationDto) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *OAuthAuthorizationDto) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *OAuthAuthorizationDto) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *OAuthAuthorizationDto) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetTokensCount

`func (o *OAuthAuthorizationDto) GetTokensCount() int32`

GetTokensCount returns the TokensCount field if non-nil, zero value otherwise.

### GetTokensCountOk

`func (o *OAuthAuthorizationDto) GetTokensCountOk() (*int32, bool)`

GetTokensCountOk returns a tuple with the TokensCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensCount

`func (o *OAuthAuthorizationDto) SetTokensCount(v int32)`

SetTokensCount sets TokensCount field to given value.

### HasTokensCount

`func (o *OAuthAuthorizationDto) HasTokensCount() bool`

HasTokensCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


