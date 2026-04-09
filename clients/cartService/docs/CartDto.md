# CartDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Ip** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**Taxes** | Pointer to **float64** |  | [optional] 
**Freight** | Pointer to **float64** |  | [optional] 
**SubTotal** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**ItemCartRecordsCount** | Pointer to **NullableInt32** |  | [optional] 
**ItemToCompareRecordsCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCartDto

`func NewCartDto() *CartDto`

NewCartDto instantiates a new CartDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartDtoWithDefaults

`func NewCartDtoWithDefaults() *CartDto`

NewCartDtoWithDefaults instantiates a new CartDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CartDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CartDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CartDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIp

`func (o *CartDto) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CartDto) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CartDto) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CartDto) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *CartDto) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *CartDto) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetType

`func (o *CartDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CartDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CartDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CartDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CartDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CartDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTotal

`func (o *CartDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CartDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CartDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CartDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTaxes

`func (o *CartDto) GetTaxes() float64`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *CartDto) GetTaxesOk() (*float64, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *CartDto) SetTaxes(v float64)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *CartDto) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetFreight

`func (o *CartDto) GetFreight() float64`

GetFreight returns the Freight field if non-nil, zero value otherwise.

### GetFreightOk

`func (o *CartDto) GetFreightOk() (*float64, bool)`

GetFreightOk returns a tuple with the Freight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreight

`func (o *CartDto) SetFreight(v float64)`

SetFreight sets Freight field to given value.

### HasFreight

`func (o *CartDto) HasFreight() bool`

HasFreight returns a boolean if a field has been set.

### GetSubTotal

`func (o *CartDto) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *CartDto) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *CartDto) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *CartDto) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetCurrencyId

`func (o *CartDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CartDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CartDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *CartDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *CartDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *CartDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCountryId

`func (o *CartDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *CartDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *CartDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *CartDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *CartDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *CartDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetItemCartRecordsCount

`func (o *CartDto) GetItemCartRecordsCount() int32`

GetItemCartRecordsCount returns the ItemCartRecordsCount field if non-nil, zero value otherwise.

### GetItemCartRecordsCountOk

`func (o *CartDto) GetItemCartRecordsCountOk() (*int32, bool)`

GetItemCartRecordsCountOk returns a tuple with the ItemCartRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCartRecordsCount

`func (o *CartDto) SetItemCartRecordsCount(v int32)`

SetItemCartRecordsCount sets ItemCartRecordsCount field to given value.

### HasItemCartRecordsCount

`func (o *CartDto) HasItemCartRecordsCount() bool`

HasItemCartRecordsCount returns a boolean if a field has been set.

### SetItemCartRecordsCountNil

`func (o *CartDto) SetItemCartRecordsCountNil(b bool)`

 SetItemCartRecordsCountNil sets the value for ItemCartRecordsCount to be an explicit nil

### UnsetItemCartRecordsCount
`func (o *CartDto) UnsetItemCartRecordsCount()`

UnsetItemCartRecordsCount ensures that no value is present for ItemCartRecordsCount, not even an explicit nil
### GetItemToCompareRecordsCount

`func (o *CartDto) GetItemToCompareRecordsCount() int32`

GetItemToCompareRecordsCount returns the ItemToCompareRecordsCount field if non-nil, zero value otherwise.

### GetItemToCompareRecordsCountOk

`func (o *CartDto) GetItemToCompareRecordsCountOk() (*int32, bool)`

GetItemToCompareRecordsCountOk returns a tuple with the ItemToCompareRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemToCompareRecordsCount

`func (o *CartDto) SetItemToCompareRecordsCount(v int32)`

SetItemToCompareRecordsCount sets ItemToCompareRecordsCount field to given value.

### HasItemToCompareRecordsCount

`func (o *CartDto) HasItemToCompareRecordsCount() bool`

HasItemToCompareRecordsCount returns a boolean if a field has been set.

### SetItemToCompareRecordsCountNil

`func (o *CartDto) SetItemToCompareRecordsCountNil(b bool)`

 SetItemToCompareRecordsCountNil sets the value for ItemToCompareRecordsCount to be an explicit nil

### UnsetItemToCompareRecordsCount
`func (o *CartDto) UnsetItemToCompareRecordsCount()`

UnsetItemToCompareRecordsCount ensures that no value is present for ItemToCompareRecordsCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


