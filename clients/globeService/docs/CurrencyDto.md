# CurrencyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Symbol** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to [**CountryDto**](CountryDto.md) |  | [optional] 

## Methods

### NewCurrencyDto

`func NewCurrencyDto() *CurrencyDto`

NewCurrencyDto instantiates a new CurrencyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyDtoWithDefaults

`func NewCurrencyDtoWithDefaults() *CurrencyDto`

NewCurrencyDtoWithDefaults instantiates a new CurrencyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrencyDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurrencyDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CurrencyDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CurrencyDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCode

`func (o *CurrencyDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CurrencyDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CurrencyDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CurrencyDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CurrencyDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CurrencyDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *CurrencyDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrencyDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrencyDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrencyDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CurrencyDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CurrencyDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSymbol

`func (o *CurrencyDto) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrencyDto) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrencyDto) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CurrencyDto) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *CurrencyDto) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *CurrencyDto) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
### GetCountry

`func (o *CurrencyDto) GetCountry() CountryDto`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CurrencyDto) GetCountryOk() (*CountryDto, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CurrencyDto) SetCountry(v CountryDto)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CurrencyDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


