# JsonWebTokenHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **NullableString** |  | [optional] 
**Jku** | Pointer to **NullableString** |  | [optional] 
**Kid** | Pointer to **NullableString** |  | [optional] 
**Typ** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJsonWebTokenHeader

`func NewJsonWebTokenHeader() *JsonWebTokenHeader`

NewJsonWebTokenHeader instantiates a new JsonWebTokenHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebTokenHeaderWithDefaults

`func NewJsonWebTokenHeaderWithDefaults() *JsonWebTokenHeader`

NewJsonWebTokenHeaderWithDefaults instantiates a new JsonWebTokenHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *JsonWebTokenHeader) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *JsonWebTokenHeader) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *JsonWebTokenHeader) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *JsonWebTokenHeader) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### SetAlgNil

`func (o *JsonWebTokenHeader) SetAlgNil(b bool)`

 SetAlgNil sets the value for Alg to be an explicit nil

### UnsetAlg
`func (o *JsonWebTokenHeader) UnsetAlg()`

UnsetAlg ensures that no value is present for Alg, not even an explicit nil
### GetJku

`func (o *JsonWebTokenHeader) GetJku() string`

GetJku returns the Jku field if non-nil, zero value otherwise.

### GetJkuOk

`func (o *JsonWebTokenHeader) GetJkuOk() (*string, bool)`

GetJkuOk returns a tuple with the Jku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJku

`func (o *JsonWebTokenHeader) SetJku(v string)`

SetJku sets Jku field to given value.

### HasJku

`func (o *JsonWebTokenHeader) HasJku() bool`

HasJku returns a boolean if a field has been set.

### SetJkuNil

`func (o *JsonWebTokenHeader) SetJkuNil(b bool)`

 SetJkuNil sets the value for Jku to be an explicit nil

### UnsetJku
`func (o *JsonWebTokenHeader) UnsetJku()`

UnsetJku ensures that no value is present for Jku, not even an explicit nil
### GetKid

`func (o *JsonWebTokenHeader) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *JsonWebTokenHeader) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *JsonWebTokenHeader) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *JsonWebTokenHeader) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *JsonWebTokenHeader) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *JsonWebTokenHeader) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetTyp

`func (o *JsonWebTokenHeader) GetTyp() string`

GetTyp returns the Typ field if non-nil, zero value otherwise.

### GetTypOk

`func (o *JsonWebTokenHeader) GetTypOk() (*string, bool)`

GetTypOk returns a tuple with the Typ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTyp

`func (o *JsonWebTokenHeader) SetTyp(v string)`

SetTyp sets Typ field to given value.

### HasTyp

`func (o *JsonWebTokenHeader) HasTyp() bool`

HasTyp returns a boolean if a field has been set.

### SetTypNil

`func (o *JsonWebTokenHeader) SetTypNil(b bool)`

 SetTypNil sets the value for Typ to be an explicit nil

### UnsetTyp
`func (o *JsonWebTokenHeader) UnsetTyp()`

UnsetTyp ensures that no value is present for Typ, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


