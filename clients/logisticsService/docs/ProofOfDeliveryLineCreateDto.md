# ProofOfDeliveryLineCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**QuantityExpected** | Pointer to **NullableInt32** |  | [optional] 
**QuantityReceived** | Pointer to **NullableInt32** |  | [optional] 
**QuantityRejected** | Pointer to **NullableInt32** |  | [optional] 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**HsCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProofOfDeliveryLineCreateDto

`func NewProofOfDeliveryLineCreateDto() *ProofOfDeliveryLineCreateDto`

NewProofOfDeliveryLineCreateDto instantiates a new ProofOfDeliveryLineCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfDeliveryLineCreateDtoWithDefaults

`func NewProofOfDeliveryLineCreateDtoWithDefaults() *ProofOfDeliveryLineCreateDto`

NewProofOfDeliveryLineCreateDtoWithDefaults instantiates a new ProofOfDeliveryLineCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProofOfDeliveryLineCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProofOfDeliveryLineCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProofOfDeliveryLineCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProofOfDeliveryLineCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProofOfDeliveryLineCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProofOfDeliveryLineCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProofOfDeliveryLineCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProofOfDeliveryLineCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *ProofOfDeliveryLineCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProofOfDeliveryLineCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProofOfDeliveryLineCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProofOfDeliveryLineCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProofOfDeliveryLineCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProofOfDeliveryLineCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantityExpected

`func (o *ProofOfDeliveryLineCreateDto) GetQuantityExpected() int32`

GetQuantityExpected returns the QuantityExpected field if non-nil, zero value otherwise.

### GetQuantityExpectedOk

`func (o *ProofOfDeliveryLineCreateDto) GetQuantityExpectedOk() (*int32, bool)`

GetQuantityExpectedOk returns a tuple with the QuantityExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityExpected

`func (o *ProofOfDeliveryLineCreateDto) SetQuantityExpected(v int32)`

SetQuantityExpected sets QuantityExpected field to given value.

### HasQuantityExpected

`func (o *ProofOfDeliveryLineCreateDto) HasQuantityExpected() bool`

HasQuantityExpected returns a boolean if a field has been set.

### SetQuantityExpectedNil

`func (o *ProofOfDeliveryLineCreateDto) SetQuantityExpectedNil(b bool)`

 SetQuantityExpectedNil sets the value for QuantityExpected to be an explicit nil

### UnsetQuantityExpected
`func (o *ProofOfDeliveryLineCreateDto) UnsetQuantityExpected()`

UnsetQuantityExpected ensures that no value is present for QuantityExpected, not even an explicit nil
### GetQuantityReceived

`func (o *ProofOfDeliveryLineCreateDto) GetQuantityReceived() int32`

GetQuantityReceived returns the QuantityReceived field if non-nil, zero value otherwise.

### GetQuantityReceivedOk

`func (o *ProofOfDeliveryLineCreateDto) GetQuantityReceivedOk() (*int32, bool)`

GetQuantityReceivedOk returns a tuple with the QuantityReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityReceived

`func (o *ProofOfDeliveryLineCreateDto) SetQuantityReceived(v int32)`

SetQuantityReceived sets QuantityReceived field to given value.

### HasQuantityReceived

`func (o *ProofOfDeliveryLineCreateDto) HasQuantityReceived() bool`

HasQuantityReceived returns a boolean if a field has been set.

### SetQuantityReceivedNil

`func (o *ProofOfDeliveryLineCreateDto) SetQuantityReceivedNil(b bool)`

 SetQuantityReceivedNil sets the value for QuantityReceived to be an explicit nil

### UnsetQuantityReceived
`func (o *ProofOfDeliveryLineCreateDto) UnsetQuantityReceived()`

UnsetQuantityReceived ensures that no value is present for QuantityReceived, not even an explicit nil
### GetQuantityRejected

`func (o *ProofOfDeliveryLineCreateDto) GetQuantityRejected() int32`

GetQuantityRejected returns the QuantityRejected field if non-nil, zero value otherwise.

### GetQuantityRejectedOk

`func (o *ProofOfDeliveryLineCreateDto) GetQuantityRejectedOk() (*int32, bool)`

GetQuantityRejectedOk returns a tuple with the QuantityRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRejected

`func (o *ProofOfDeliveryLineCreateDto) SetQuantityRejected(v int32)`

SetQuantityRejected sets QuantityRejected field to given value.

### HasQuantityRejected

`func (o *ProofOfDeliveryLineCreateDto) HasQuantityRejected() bool`

HasQuantityRejected returns a boolean if a field has been set.

### SetQuantityRejectedNil

`func (o *ProofOfDeliveryLineCreateDto) SetQuantityRejectedNil(b bool)`

 SetQuantityRejectedNil sets the value for QuantityRejected to be an explicit nil

### UnsetQuantityRejected
`func (o *ProofOfDeliveryLineCreateDto) UnsetQuantityRejected()`

UnsetQuantityRejected ensures that no value is present for QuantityRejected, not even an explicit nil
### GetCondition

`func (o *ProofOfDeliveryLineCreateDto) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProofOfDeliveryLineCreateDto) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProofOfDeliveryLineCreateDto) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProofOfDeliveryLineCreateDto) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ProofOfDeliveryLineCreateDto) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ProofOfDeliveryLineCreateDto) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetRemarks

`func (o *ProofOfDeliveryLineCreateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ProofOfDeliveryLineCreateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ProofOfDeliveryLineCreateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ProofOfDeliveryLineCreateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *ProofOfDeliveryLineCreateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *ProofOfDeliveryLineCreateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetHsCode

`func (o *ProofOfDeliveryLineCreateDto) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProofOfDeliveryLineCreateDto) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProofOfDeliveryLineCreateDto) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProofOfDeliveryLineCreateDto) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *ProofOfDeliveryLineCreateDto) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *ProofOfDeliveryLineCreateDto) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


