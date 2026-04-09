# JsonWebKeySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]JsonWebKey**](JsonWebKey.md) |  | [optional] 

## Methods

### NewJsonWebKeySet

`func NewJsonWebKeySet() *JsonWebKeySet`

NewJsonWebKeySet instantiates a new JsonWebKeySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeySetWithDefaults

`func NewJsonWebKeySetWithDefaults() *JsonWebKeySet`

NewJsonWebKeySetWithDefaults instantiates a new JsonWebKeySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *JsonWebKeySet) GetKeys() []JsonWebKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *JsonWebKeySet) GetKeysOk() (*[]JsonWebKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *JsonWebKeySet) SetKeys(v []JsonWebKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *JsonWebKeySet) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeysNil

`func (o *JsonWebKeySet) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *JsonWebKeySet) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


