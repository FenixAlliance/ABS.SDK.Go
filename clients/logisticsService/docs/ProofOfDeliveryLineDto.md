# ProofOfDeliveryLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ProofOfDeliveryId** | Pointer to **NullableString** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**QuantityExpected** | Pointer to **NullableInt32** |  | [optional] 
**QuantityReceived** | Pointer to **NullableInt32** |  | [optional] 
**QuantityRejected** | Pointer to **NullableInt32** |  | [optional] 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**HsCode** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProofOfDeliveryLineDto

`func NewProofOfDeliveryLineDto() *ProofOfDeliveryLineDto`

NewProofOfDeliveryLineDto instantiates a new ProofOfDeliveryLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfDeliveryLineDtoWithDefaults

`func NewProofOfDeliveryLineDtoWithDefaults() *ProofOfDeliveryLineDto`

NewProofOfDeliveryLineDtoWithDefaults instantiates a new ProofOfDeliveryLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProofOfDeliveryLineDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProofOfDeliveryLineDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProofOfDeliveryLineDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProofOfDeliveryLineDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProofOfDeliveryLineDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProofOfDeliveryLineDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ProofOfDeliveryLineDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProofOfDeliveryLineDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProofOfDeliveryLineDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProofOfDeliveryLineDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProofOfDeliveryLineDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProofOfDeliveryLineDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetProofOfDeliveryId

`func (o *ProofOfDeliveryLineDto) GetProofOfDeliveryId() string`

GetProofOfDeliveryId returns the ProofOfDeliveryId field if non-nil, zero value otherwise.

### GetProofOfDeliveryIdOk

`func (o *ProofOfDeliveryLineDto) GetProofOfDeliveryIdOk() (*string, bool)`

GetProofOfDeliveryIdOk returns a tuple with the ProofOfDeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDeliveryId

`func (o *ProofOfDeliveryLineDto) SetProofOfDeliveryId(v string)`

SetProofOfDeliveryId sets ProofOfDeliveryId field to given value.

### HasProofOfDeliveryId

`func (o *ProofOfDeliveryLineDto) HasProofOfDeliveryId() bool`

HasProofOfDeliveryId returns a boolean if a field has been set.

### SetProofOfDeliveryIdNil

`func (o *ProofOfDeliveryLineDto) SetProofOfDeliveryIdNil(b bool)`

 SetProofOfDeliveryIdNil sets the value for ProofOfDeliveryId to be an explicit nil

### UnsetProofOfDeliveryId
`func (o *ProofOfDeliveryLineDto) UnsetProofOfDeliveryId()`

UnsetProofOfDeliveryId ensures that no value is present for ProofOfDeliveryId, not even an explicit nil
### GetLineNumber

`func (o *ProofOfDeliveryLineDto) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ProofOfDeliveryLineDto) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ProofOfDeliveryLineDto) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *ProofOfDeliveryLineDto) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetDescription

`func (o *ProofOfDeliveryLineDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProofOfDeliveryLineDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProofOfDeliveryLineDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProofOfDeliveryLineDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProofOfDeliveryLineDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProofOfDeliveryLineDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantityExpected

`func (o *ProofOfDeliveryLineDto) GetQuantityExpected() int32`

GetQuantityExpected returns the QuantityExpected field if non-nil, zero value otherwise.

### GetQuantityExpectedOk

`func (o *ProofOfDeliveryLineDto) GetQuantityExpectedOk() (*int32, bool)`

GetQuantityExpectedOk returns a tuple with the QuantityExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityExpected

`func (o *ProofOfDeliveryLineDto) SetQuantityExpected(v int32)`

SetQuantityExpected sets QuantityExpected field to given value.

### HasQuantityExpected

`func (o *ProofOfDeliveryLineDto) HasQuantityExpected() bool`

HasQuantityExpected returns a boolean if a field has been set.

### SetQuantityExpectedNil

`func (o *ProofOfDeliveryLineDto) SetQuantityExpectedNil(b bool)`

 SetQuantityExpectedNil sets the value for QuantityExpected to be an explicit nil

### UnsetQuantityExpected
`func (o *ProofOfDeliveryLineDto) UnsetQuantityExpected()`

UnsetQuantityExpected ensures that no value is present for QuantityExpected, not even an explicit nil
### GetQuantityReceived

`func (o *ProofOfDeliveryLineDto) GetQuantityReceived() int32`

GetQuantityReceived returns the QuantityReceived field if non-nil, zero value otherwise.

### GetQuantityReceivedOk

`func (o *ProofOfDeliveryLineDto) GetQuantityReceivedOk() (*int32, bool)`

GetQuantityReceivedOk returns a tuple with the QuantityReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityReceived

`func (o *ProofOfDeliveryLineDto) SetQuantityReceived(v int32)`

SetQuantityReceived sets QuantityReceived field to given value.

### HasQuantityReceived

`func (o *ProofOfDeliveryLineDto) HasQuantityReceived() bool`

HasQuantityReceived returns a boolean if a field has been set.

### SetQuantityReceivedNil

`func (o *ProofOfDeliveryLineDto) SetQuantityReceivedNil(b bool)`

 SetQuantityReceivedNil sets the value for QuantityReceived to be an explicit nil

### UnsetQuantityReceived
`func (o *ProofOfDeliveryLineDto) UnsetQuantityReceived()`

UnsetQuantityReceived ensures that no value is present for QuantityReceived, not even an explicit nil
### GetQuantityRejected

`func (o *ProofOfDeliveryLineDto) GetQuantityRejected() int32`

GetQuantityRejected returns the QuantityRejected field if non-nil, zero value otherwise.

### GetQuantityRejectedOk

`func (o *ProofOfDeliveryLineDto) GetQuantityRejectedOk() (*int32, bool)`

GetQuantityRejectedOk returns a tuple with the QuantityRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRejected

`func (o *ProofOfDeliveryLineDto) SetQuantityRejected(v int32)`

SetQuantityRejected sets QuantityRejected field to given value.

### HasQuantityRejected

`func (o *ProofOfDeliveryLineDto) HasQuantityRejected() bool`

HasQuantityRejected returns a boolean if a field has been set.

### SetQuantityRejectedNil

`func (o *ProofOfDeliveryLineDto) SetQuantityRejectedNil(b bool)`

 SetQuantityRejectedNil sets the value for QuantityRejected to be an explicit nil

### UnsetQuantityRejected
`func (o *ProofOfDeliveryLineDto) UnsetQuantityRejected()`

UnsetQuantityRejected ensures that no value is present for QuantityRejected, not even an explicit nil
### GetCondition

`func (o *ProofOfDeliveryLineDto) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProofOfDeliveryLineDto) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProofOfDeliveryLineDto) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProofOfDeliveryLineDto) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ProofOfDeliveryLineDto) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ProofOfDeliveryLineDto) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetRemarks

`func (o *ProofOfDeliveryLineDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ProofOfDeliveryLineDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ProofOfDeliveryLineDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ProofOfDeliveryLineDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *ProofOfDeliveryLineDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *ProofOfDeliveryLineDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetHsCode

`func (o *ProofOfDeliveryLineDto) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProofOfDeliveryLineDto) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProofOfDeliveryLineDto) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProofOfDeliveryLineDto) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *ProofOfDeliveryLineDto) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *ProofOfDeliveryLineDto) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetTenantId

`func (o *ProofOfDeliveryLineDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProofOfDeliveryLineDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProofOfDeliveryLineDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ProofOfDeliveryLineDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ProofOfDeliveryLineDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ProofOfDeliveryLineDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


