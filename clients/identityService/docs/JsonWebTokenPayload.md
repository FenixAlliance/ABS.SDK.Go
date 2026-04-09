# JsonWebTokenPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | Pointer to **NullableString** |  | [optional] 
**Cid** | Pointer to **NullableString** |  | [optional] 
**Iss** | Pointer to **NullableString** |  | [optional] 
**Aid** | Pointer to **NullableString** |  | [optional] 
**Sub** | Pointer to **NullableString** |  | [optional] 
**Act** | Pointer to **NullableString** |  | [optional] 
**Iat** | Pointer to **int64** |  | [optional] 
**Nbf** | Pointer to **int64** |  | [optional] 
**Exp** | Pointer to **int64** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJsonWebTokenPayload

`func NewJsonWebTokenPayload() *JsonWebTokenPayload`

NewJsonWebTokenPayload instantiates a new JsonWebTokenPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebTokenPayloadWithDefaults

`func NewJsonWebTokenPayloadWithDefaults() *JsonWebTokenPayload`

NewJsonWebTokenPayloadWithDefaults instantiates a new JsonWebTokenPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *JsonWebTokenPayload) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *JsonWebTokenPayload) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *JsonWebTokenPayload) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *JsonWebTokenPayload) HasAud() bool`

HasAud returns a boolean if a field has been set.

### SetAudNil

`func (o *JsonWebTokenPayload) SetAudNil(b bool)`

 SetAudNil sets the value for Aud to be an explicit nil

### UnsetAud
`func (o *JsonWebTokenPayload) UnsetAud()`

UnsetAud ensures that no value is present for Aud, not even an explicit nil
### GetCid

`func (o *JsonWebTokenPayload) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *JsonWebTokenPayload) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *JsonWebTokenPayload) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *JsonWebTokenPayload) HasCid() bool`

HasCid returns a boolean if a field has been set.

### SetCidNil

`func (o *JsonWebTokenPayload) SetCidNil(b bool)`

 SetCidNil sets the value for Cid to be an explicit nil

### UnsetCid
`func (o *JsonWebTokenPayload) UnsetCid()`

UnsetCid ensures that no value is present for Cid, not even an explicit nil
### GetIss

`func (o *JsonWebTokenPayload) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *JsonWebTokenPayload) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *JsonWebTokenPayload) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *JsonWebTokenPayload) HasIss() bool`

HasIss returns a boolean if a field has been set.

### SetIssNil

`func (o *JsonWebTokenPayload) SetIssNil(b bool)`

 SetIssNil sets the value for Iss to be an explicit nil

### UnsetIss
`func (o *JsonWebTokenPayload) UnsetIss()`

UnsetIss ensures that no value is present for Iss, not even an explicit nil
### GetAid

`func (o *JsonWebTokenPayload) GetAid() string`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *JsonWebTokenPayload) GetAidOk() (*string, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *JsonWebTokenPayload) SetAid(v string)`

SetAid sets Aid field to given value.

### HasAid

`func (o *JsonWebTokenPayload) HasAid() bool`

HasAid returns a boolean if a field has been set.

### SetAidNil

`func (o *JsonWebTokenPayload) SetAidNil(b bool)`

 SetAidNil sets the value for Aid to be an explicit nil

### UnsetAid
`func (o *JsonWebTokenPayload) UnsetAid()`

UnsetAid ensures that no value is present for Aid, not even an explicit nil
### GetSub

`func (o *JsonWebTokenPayload) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *JsonWebTokenPayload) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *JsonWebTokenPayload) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *JsonWebTokenPayload) HasSub() bool`

HasSub returns a boolean if a field has been set.

### SetSubNil

`func (o *JsonWebTokenPayload) SetSubNil(b bool)`

 SetSubNil sets the value for Sub to be an explicit nil

### UnsetSub
`func (o *JsonWebTokenPayload) UnsetSub()`

UnsetSub ensures that no value is present for Sub, not even an explicit nil
### GetAct

`func (o *JsonWebTokenPayload) GetAct() string`

GetAct returns the Act field if non-nil, zero value otherwise.

### GetActOk

`func (o *JsonWebTokenPayload) GetActOk() (*string, bool)`

GetActOk returns a tuple with the Act field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAct

`func (o *JsonWebTokenPayload) SetAct(v string)`

SetAct sets Act field to given value.

### HasAct

`func (o *JsonWebTokenPayload) HasAct() bool`

HasAct returns a boolean if a field has been set.

### SetActNil

`func (o *JsonWebTokenPayload) SetActNil(b bool)`

 SetActNil sets the value for Act to be an explicit nil

### UnsetAct
`func (o *JsonWebTokenPayload) UnsetAct()`

UnsetAct ensures that no value is present for Act, not even an explicit nil
### GetIat

`func (o *JsonWebTokenPayload) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *JsonWebTokenPayload) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *JsonWebTokenPayload) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *JsonWebTokenPayload) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetNbf

`func (o *JsonWebTokenPayload) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *JsonWebTokenPayload) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *JsonWebTokenPayload) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *JsonWebTokenPayload) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetExp

`func (o *JsonWebTokenPayload) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *JsonWebTokenPayload) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *JsonWebTokenPayload) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *JsonWebTokenPayload) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetScopes

`func (o *JsonWebTokenPayload) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *JsonWebTokenPayload) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *JsonWebTokenPayload) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *JsonWebTokenPayload) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *JsonWebTokenPayload) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *JsonWebTokenPayload) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


