# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCurrency

`func NewCurrency() *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Currency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Currency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Currency) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Currency) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Currency) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Currency) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCountry

`func (o *Currency) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Currency) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Currency) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Currency) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Currency) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Currency) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


