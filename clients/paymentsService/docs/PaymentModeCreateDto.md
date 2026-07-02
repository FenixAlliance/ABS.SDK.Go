# PaymentModeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**PaymentMeansCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentModeCreateDto

`func NewPaymentModeCreateDto(name string, ) *PaymentModeCreateDto`

NewPaymentModeCreateDto instantiates a new PaymentModeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentModeCreateDtoWithDefaults

`func NewPaymentModeCreateDtoWithDefaults() *PaymentModeCreateDto`

NewPaymentModeCreateDtoWithDefaults instantiates a new PaymentModeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentModeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentModeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentModeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentModeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PaymentModeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentModeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentModeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentModeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *PaymentModeCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentModeCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentModeCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PaymentModeCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentModeCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentModeCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentModeCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentModeCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentModeCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPaymentMeansCode

`func (o *PaymentModeCreateDto) GetPaymentMeansCode() string`

GetPaymentMeansCode returns the PaymentMeansCode field if non-nil, zero value otherwise.

### GetPaymentMeansCodeOk

`func (o *PaymentModeCreateDto) GetPaymentMeansCodeOk() (*string, bool)`

GetPaymentMeansCodeOk returns a tuple with the PaymentMeansCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMeansCode

`func (o *PaymentModeCreateDto) SetPaymentMeansCode(v string)`

SetPaymentMeansCode sets PaymentMeansCode field to given value.

### HasPaymentMeansCode

`func (o *PaymentModeCreateDto) HasPaymentMeansCode() bool`

HasPaymentMeansCode returns a boolean if a field has been set.

### SetPaymentMeansCodeNil

`func (o *PaymentModeCreateDto) SetPaymentMeansCodeNil(b bool)`

 SetPaymentMeansCodeNil sets the value for PaymentMeansCode to be an explicit nil

### UnsetPaymentMeansCode
`func (o *PaymentModeCreateDto) UnsetPaymentMeansCode()`

UnsetPaymentMeansCode ensures that no value is present for PaymentMeansCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


