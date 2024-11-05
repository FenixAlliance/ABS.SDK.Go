# SupportEntitlementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StartDateTime** | Pointer to **time.Time** |  | [optional] 
**EndDateTime** | Pointer to **time.Time** |  | [optional] 
**NextInvoiceDateTime** | Pointer to **time.Time** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Signature** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Repetitions** | Pointer to **int32** |  | [optional] 
**ChargeAttempts** | Pointer to **int32** |  | [optional] 
**FreeTrialInDays** | Pointer to **int32** |  | [optional] 
**GracePeriodInDays** | Pointer to **int32** |  | [optional] 
**CustomRenewalPeriod** | Pointer to **int32** |  | [optional] 
**EnableAutomaticRenew** | Pointer to **bool** |  | [optional] 
**EnableProRateBilling** | Pointer to **bool** |  | [optional] 
**EnableUsageThreshold** | Pointer to **bool** |  | [optional] 
**EnableAutomaticDisable** | Pointer to **bool** |  | [optional] 
**EnableAutomaticPayments** | Pointer to **bool** |  | [optional] 
**UsageThreshold** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**DataLabel** | Pointer to **NullableString** |  | [optional] 
**Data1** | Pointer to **NullableString** |  | [optional] 
**Data1Label** | Pointer to **NullableString** |  | [optional] 
**Data2** | Pointer to **NullableString** |  | [optional] 
**Data2Label** | Pointer to **NullableString** |  | [optional] 
**Data3** | Pointer to **NullableString** |  | [optional] 
**Data3Label** | Pointer to **NullableString** |  | [optional] 
**Data4** | Pointer to **NullableString** |  | [optional] 
**Data4Label** | Pointer to **NullableString** |  | [optional] 
**Data5** | Pointer to **NullableString** |  | [optional] 
**Data5Label** | Pointer to **NullableString** |  | [optional] 
**Data6** | Pointer to **NullableString** |  | [optional] 
**Data6Label** | Pointer to **NullableString** |  | [optional] 
**Data7** | Pointer to **NullableString** |  | [optional] 
**Data7Label** | Pointer to **NullableString** |  | [optional] 
**Data8** | Pointer to **NullableString** |  | [optional] 
**Data8Label** | Pointer to **NullableString** |  | [optional] 
**Data9** | Pointer to **NullableString** |  | [optional] 
**Data9Label** | Pointer to **NullableString** |  | [optional] 
**AccountHolderID** | Pointer to **NullableString** |  | [optional] 
**IndividualID** | Pointer to **NullableString** |  | [optional] 
**OrganizationID** | Pointer to **NullableString** |  | [optional] 
**ReceiverBusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**PaymentTokenID** | Pointer to **NullableString** |  | [optional] 
**WalletAccountID** | Pointer to **NullableString** |  | [optional] 
**SecurityCertificateID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportEntitlementDto

`func NewSupportEntitlementDto() *SupportEntitlementDto`

NewSupportEntitlementDto instantiates a new SupportEntitlementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportEntitlementDtoWithDefaults

`func NewSupportEntitlementDtoWithDefaults() *SupportEntitlementDto`

NewSupportEntitlementDtoWithDefaults instantiates a new SupportEntitlementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportEntitlementDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportEntitlementDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportEntitlementDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportEntitlementDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SupportEntitlementDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SupportEntitlementDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SupportEntitlementDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportEntitlementDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportEntitlementDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportEntitlementDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SupportEntitlementDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SupportEntitlementDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SupportEntitlementDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportEntitlementDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportEntitlementDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportEntitlementDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportEntitlementDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportEntitlementDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportEntitlementDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportEntitlementDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportEntitlementDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportEntitlementDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportEntitlementDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportEntitlementDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartDateTime

`func (o *SupportEntitlementDto) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *SupportEntitlementDto) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *SupportEntitlementDto) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *SupportEntitlementDto) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *SupportEntitlementDto) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *SupportEntitlementDto) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *SupportEntitlementDto) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *SupportEntitlementDto) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetNextInvoiceDateTime

`func (o *SupportEntitlementDto) GetNextInvoiceDateTime() time.Time`

GetNextInvoiceDateTime returns the NextInvoiceDateTime field if non-nil, zero value otherwise.

### GetNextInvoiceDateTimeOk

`func (o *SupportEntitlementDto) GetNextInvoiceDateTimeOk() (*time.Time, bool)`

GetNextInvoiceDateTimeOk returns a tuple with the NextInvoiceDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDateTime

`func (o *SupportEntitlementDto) SetNextInvoiceDateTime(v time.Time)`

SetNextInvoiceDateTime sets NextInvoiceDateTime field to given value.

### HasNextInvoiceDateTime

`func (o *SupportEntitlementDto) HasNextInvoiceDateTime() bool`

HasNextInvoiceDateTime returns a boolean if a field has been set.

### GetCode

`func (o *SupportEntitlementDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SupportEntitlementDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SupportEntitlementDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SupportEntitlementDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *SupportEntitlementDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SupportEntitlementDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSignature

`func (o *SupportEntitlementDto) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SupportEntitlementDto) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SupportEntitlementDto) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SupportEntitlementDto) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SupportEntitlementDto) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SupportEntitlementDto) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetQuantity

`func (o *SupportEntitlementDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SupportEntitlementDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SupportEntitlementDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SupportEntitlementDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRepetitions

`func (o *SupportEntitlementDto) GetRepetitions() int32`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *SupportEntitlementDto) GetRepetitionsOk() (*int32, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *SupportEntitlementDto) SetRepetitions(v int32)`

SetRepetitions sets Repetitions field to given value.

### HasRepetitions

`func (o *SupportEntitlementDto) HasRepetitions() bool`

HasRepetitions returns a boolean if a field has been set.

### GetChargeAttempts

`func (o *SupportEntitlementDto) GetChargeAttempts() int32`

GetChargeAttempts returns the ChargeAttempts field if non-nil, zero value otherwise.

### GetChargeAttemptsOk

`func (o *SupportEntitlementDto) GetChargeAttemptsOk() (*int32, bool)`

GetChargeAttemptsOk returns a tuple with the ChargeAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAttempts

`func (o *SupportEntitlementDto) SetChargeAttempts(v int32)`

SetChargeAttempts sets ChargeAttempts field to given value.

### HasChargeAttempts

`func (o *SupportEntitlementDto) HasChargeAttempts() bool`

HasChargeAttempts returns a boolean if a field has been set.

### GetFreeTrialInDays

`func (o *SupportEntitlementDto) GetFreeTrialInDays() int32`

GetFreeTrialInDays returns the FreeTrialInDays field if non-nil, zero value otherwise.

### GetFreeTrialInDaysOk

`func (o *SupportEntitlementDto) GetFreeTrialInDaysOk() (*int32, bool)`

GetFreeTrialInDaysOk returns a tuple with the FreeTrialInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrialInDays

`func (o *SupportEntitlementDto) SetFreeTrialInDays(v int32)`

SetFreeTrialInDays sets FreeTrialInDays field to given value.

### HasFreeTrialInDays

`func (o *SupportEntitlementDto) HasFreeTrialInDays() bool`

HasFreeTrialInDays returns a boolean if a field has been set.

### GetGracePeriodInDays

`func (o *SupportEntitlementDto) GetGracePeriodInDays() int32`

GetGracePeriodInDays returns the GracePeriodInDays field if non-nil, zero value otherwise.

### GetGracePeriodInDaysOk

`func (o *SupportEntitlementDto) GetGracePeriodInDaysOk() (*int32, bool)`

GetGracePeriodInDaysOk returns a tuple with the GracePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodInDays

`func (o *SupportEntitlementDto) SetGracePeriodInDays(v int32)`

SetGracePeriodInDays sets GracePeriodInDays field to given value.

### HasGracePeriodInDays

`func (o *SupportEntitlementDto) HasGracePeriodInDays() bool`

HasGracePeriodInDays returns a boolean if a field has been set.

### GetCustomRenewalPeriod

`func (o *SupportEntitlementDto) GetCustomRenewalPeriod() int32`

GetCustomRenewalPeriod returns the CustomRenewalPeriod field if non-nil, zero value otherwise.

### GetCustomRenewalPeriodOk

`func (o *SupportEntitlementDto) GetCustomRenewalPeriodOk() (*int32, bool)`

GetCustomRenewalPeriodOk returns a tuple with the CustomRenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRenewalPeriod

`func (o *SupportEntitlementDto) SetCustomRenewalPeriod(v int32)`

SetCustomRenewalPeriod sets CustomRenewalPeriod field to given value.

### HasCustomRenewalPeriod

`func (o *SupportEntitlementDto) HasCustomRenewalPeriod() bool`

HasCustomRenewalPeriod returns a boolean if a field has been set.

### GetEnableAutomaticRenew

`func (o *SupportEntitlementDto) GetEnableAutomaticRenew() bool`

GetEnableAutomaticRenew returns the EnableAutomaticRenew field if non-nil, zero value otherwise.

### GetEnableAutomaticRenewOk

`func (o *SupportEntitlementDto) GetEnableAutomaticRenewOk() (*bool, bool)`

GetEnableAutomaticRenewOk returns a tuple with the EnableAutomaticRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticRenew

`func (o *SupportEntitlementDto) SetEnableAutomaticRenew(v bool)`

SetEnableAutomaticRenew sets EnableAutomaticRenew field to given value.

### HasEnableAutomaticRenew

`func (o *SupportEntitlementDto) HasEnableAutomaticRenew() bool`

HasEnableAutomaticRenew returns a boolean if a field has been set.

### GetEnableProRateBilling

`func (o *SupportEntitlementDto) GetEnableProRateBilling() bool`

GetEnableProRateBilling returns the EnableProRateBilling field if non-nil, zero value otherwise.

### GetEnableProRateBillingOk

`func (o *SupportEntitlementDto) GetEnableProRateBillingOk() (*bool, bool)`

GetEnableProRateBillingOk returns a tuple with the EnableProRateBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProRateBilling

`func (o *SupportEntitlementDto) SetEnableProRateBilling(v bool)`

SetEnableProRateBilling sets EnableProRateBilling field to given value.

### HasEnableProRateBilling

`func (o *SupportEntitlementDto) HasEnableProRateBilling() bool`

HasEnableProRateBilling returns a boolean if a field has been set.

### GetEnableUsageThreshold

`func (o *SupportEntitlementDto) GetEnableUsageThreshold() bool`

GetEnableUsageThreshold returns the EnableUsageThreshold field if non-nil, zero value otherwise.

### GetEnableUsageThresholdOk

`func (o *SupportEntitlementDto) GetEnableUsageThresholdOk() (*bool, bool)`

GetEnableUsageThresholdOk returns a tuple with the EnableUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUsageThreshold

`func (o *SupportEntitlementDto) SetEnableUsageThreshold(v bool)`

SetEnableUsageThreshold sets EnableUsageThreshold field to given value.

### HasEnableUsageThreshold

`func (o *SupportEntitlementDto) HasEnableUsageThreshold() bool`

HasEnableUsageThreshold returns a boolean if a field has been set.

### GetEnableAutomaticDisable

`func (o *SupportEntitlementDto) GetEnableAutomaticDisable() bool`

GetEnableAutomaticDisable returns the EnableAutomaticDisable field if non-nil, zero value otherwise.

### GetEnableAutomaticDisableOk

`func (o *SupportEntitlementDto) GetEnableAutomaticDisableOk() (*bool, bool)`

GetEnableAutomaticDisableOk returns a tuple with the EnableAutomaticDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticDisable

`func (o *SupportEntitlementDto) SetEnableAutomaticDisable(v bool)`

SetEnableAutomaticDisable sets EnableAutomaticDisable field to given value.

### HasEnableAutomaticDisable

`func (o *SupportEntitlementDto) HasEnableAutomaticDisable() bool`

HasEnableAutomaticDisable returns a boolean if a field has been set.

### GetEnableAutomaticPayments

`func (o *SupportEntitlementDto) GetEnableAutomaticPayments() bool`

GetEnableAutomaticPayments returns the EnableAutomaticPayments field if non-nil, zero value otherwise.

### GetEnableAutomaticPaymentsOk

`func (o *SupportEntitlementDto) GetEnableAutomaticPaymentsOk() (*bool, bool)`

GetEnableAutomaticPaymentsOk returns a tuple with the EnableAutomaticPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPayments

`func (o *SupportEntitlementDto) SetEnableAutomaticPayments(v bool)`

SetEnableAutomaticPayments sets EnableAutomaticPayments field to given value.

### HasEnableAutomaticPayments

`func (o *SupportEntitlementDto) HasEnableAutomaticPayments() bool`

HasEnableAutomaticPayments returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *SupportEntitlementDto) GetUsageThreshold() int32`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *SupportEntitlementDto) GetUsageThresholdOk() (*int32, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *SupportEntitlementDto) SetUsageThreshold(v int32)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *SupportEntitlementDto) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### GetData

`func (o *SupportEntitlementDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SupportEntitlementDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SupportEntitlementDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SupportEntitlementDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SupportEntitlementDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SupportEntitlementDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *SupportEntitlementDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *SupportEntitlementDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *SupportEntitlementDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *SupportEntitlementDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *SupportEntitlementDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *SupportEntitlementDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *SupportEntitlementDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *SupportEntitlementDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *SupportEntitlementDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *SupportEntitlementDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *SupportEntitlementDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *SupportEntitlementDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *SupportEntitlementDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *SupportEntitlementDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *SupportEntitlementDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *SupportEntitlementDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *SupportEntitlementDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *SupportEntitlementDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *SupportEntitlementDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *SupportEntitlementDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *SupportEntitlementDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *SupportEntitlementDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *SupportEntitlementDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *SupportEntitlementDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *SupportEntitlementDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *SupportEntitlementDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *SupportEntitlementDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *SupportEntitlementDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *SupportEntitlementDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *SupportEntitlementDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *SupportEntitlementDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *SupportEntitlementDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *SupportEntitlementDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *SupportEntitlementDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *SupportEntitlementDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *SupportEntitlementDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *SupportEntitlementDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *SupportEntitlementDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *SupportEntitlementDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *SupportEntitlementDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *SupportEntitlementDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *SupportEntitlementDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *SupportEntitlementDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *SupportEntitlementDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *SupportEntitlementDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *SupportEntitlementDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *SupportEntitlementDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *SupportEntitlementDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *SupportEntitlementDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *SupportEntitlementDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *SupportEntitlementDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *SupportEntitlementDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *SupportEntitlementDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *SupportEntitlementDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *SupportEntitlementDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *SupportEntitlementDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *SupportEntitlementDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *SupportEntitlementDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *SupportEntitlementDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *SupportEntitlementDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *SupportEntitlementDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *SupportEntitlementDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *SupportEntitlementDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *SupportEntitlementDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *SupportEntitlementDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *SupportEntitlementDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *SupportEntitlementDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *SupportEntitlementDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *SupportEntitlementDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *SupportEntitlementDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *SupportEntitlementDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *SupportEntitlementDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *SupportEntitlementDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *SupportEntitlementDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *SupportEntitlementDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *SupportEntitlementDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *SupportEntitlementDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *SupportEntitlementDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *SupportEntitlementDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *SupportEntitlementDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *SupportEntitlementDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *SupportEntitlementDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *SupportEntitlementDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *SupportEntitlementDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *SupportEntitlementDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *SupportEntitlementDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *SupportEntitlementDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *SupportEntitlementDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *SupportEntitlementDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *SupportEntitlementDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *SupportEntitlementDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *SupportEntitlementDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *SupportEntitlementDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *SupportEntitlementDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *SupportEntitlementDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *SupportEntitlementDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *SupportEntitlementDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *SupportEntitlementDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *SupportEntitlementDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *SupportEntitlementDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *SupportEntitlementDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *SupportEntitlementDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *SupportEntitlementDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *SupportEntitlementDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *SupportEntitlementDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *SupportEntitlementDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *SupportEntitlementDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *SupportEntitlementDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *SupportEntitlementDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *SupportEntitlementDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *SupportEntitlementDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *SupportEntitlementDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *SupportEntitlementDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *SupportEntitlementDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetAccountHolderID

`func (o *SupportEntitlementDto) GetAccountHolderID() string`

GetAccountHolderID returns the AccountHolderID field if non-nil, zero value otherwise.

### GetAccountHolderIDOk

`func (o *SupportEntitlementDto) GetAccountHolderIDOk() (*string, bool)`

GetAccountHolderIDOk returns a tuple with the AccountHolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderID

`func (o *SupportEntitlementDto) SetAccountHolderID(v string)`

SetAccountHolderID sets AccountHolderID field to given value.

### HasAccountHolderID

`func (o *SupportEntitlementDto) HasAccountHolderID() bool`

HasAccountHolderID returns a boolean if a field has been set.

### SetAccountHolderIDNil

`func (o *SupportEntitlementDto) SetAccountHolderIDNil(b bool)`

 SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil

### UnsetAccountHolderID
`func (o *SupportEntitlementDto) UnsetAccountHolderID()`

UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil
### GetIndividualID

`func (o *SupportEntitlementDto) GetIndividualID() string`

GetIndividualID returns the IndividualID field if non-nil, zero value otherwise.

### GetIndividualIDOk

`func (o *SupportEntitlementDto) GetIndividualIDOk() (*string, bool)`

GetIndividualIDOk returns a tuple with the IndividualID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualID

`func (o *SupportEntitlementDto) SetIndividualID(v string)`

SetIndividualID sets IndividualID field to given value.

### HasIndividualID

`func (o *SupportEntitlementDto) HasIndividualID() bool`

HasIndividualID returns a boolean if a field has been set.

### SetIndividualIDNil

`func (o *SupportEntitlementDto) SetIndividualIDNil(b bool)`

 SetIndividualIDNil sets the value for IndividualID to be an explicit nil

### UnsetIndividualID
`func (o *SupportEntitlementDto) UnsetIndividualID()`

UnsetIndividualID ensures that no value is present for IndividualID, not even an explicit nil
### GetOrganizationID

`func (o *SupportEntitlementDto) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *SupportEntitlementDto) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *SupportEntitlementDto) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *SupportEntitlementDto) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### SetOrganizationIDNil

`func (o *SupportEntitlementDto) SetOrganizationIDNil(b bool)`

 SetOrganizationIDNil sets the value for OrganizationID to be an explicit nil

### UnsetOrganizationID
`func (o *SupportEntitlementDto) UnsetOrganizationID()`

UnsetOrganizationID ensures that no value is present for OrganizationID, not even an explicit nil
### GetReceiverBusinessID

`func (o *SupportEntitlementDto) GetReceiverBusinessID() string`

GetReceiverBusinessID returns the ReceiverBusinessID field if non-nil, zero value otherwise.

### GetReceiverBusinessIDOk

`func (o *SupportEntitlementDto) GetReceiverBusinessIDOk() (*string, bool)`

GetReceiverBusinessIDOk returns a tuple with the ReceiverBusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessID

`func (o *SupportEntitlementDto) SetReceiverBusinessID(v string)`

SetReceiverBusinessID sets ReceiverBusinessID field to given value.

### HasReceiverBusinessID

`func (o *SupportEntitlementDto) HasReceiverBusinessID() bool`

HasReceiverBusinessID returns a boolean if a field has been set.

### SetReceiverBusinessIDNil

`func (o *SupportEntitlementDto) SetReceiverBusinessIDNil(b bool)`

 SetReceiverBusinessIDNil sets the value for ReceiverBusinessID to be an explicit nil

### UnsetReceiverBusinessID
`func (o *SupportEntitlementDto) UnsetReceiverBusinessID()`

UnsetReceiverBusinessID ensures that no value is present for ReceiverBusinessID, not even an explicit nil
### GetBusinessID

`func (o *SupportEntitlementDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportEntitlementDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportEntitlementDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportEntitlementDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportEntitlementDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportEntitlementDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *SupportEntitlementDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *SupportEntitlementDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *SupportEntitlementDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *SupportEntitlementDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *SupportEntitlementDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *SupportEntitlementDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetPaymentTokenID

`func (o *SupportEntitlementDto) GetPaymentTokenID() string`

GetPaymentTokenID returns the PaymentTokenID field if non-nil, zero value otherwise.

### GetPaymentTokenIDOk

`func (o *SupportEntitlementDto) GetPaymentTokenIDOk() (*string, bool)`

GetPaymentTokenIDOk returns a tuple with the PaymentTokenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTokenID

`func (o *SupportEntitlementDto) SetPaymentTokenID(v string)`

SetPaymentTokenID sets PaymentTokenID field to given value.

### HasPaymentTokenID

`func (o *SupportEntitlementDto) HasPaymentTokenID() bool`

HasPaymentTokenID returns a boolean if a field has been set.

### SetPaymentTokenIDNil

`func (o *SupportEntitlementDto) SetPaymentTokenIDNil(b bool)`

 SetPaymentTokenIDNil sets the value for PaymentTokenID to be an explicit nil

### UnsetPaymentTokenID
`func (o *SupportEntitlementDto) UnsetPaymentTokenID()`

UnsetPaymentTokenID ensures that no value is present for PaymentTokenID, not even an explicit nil
### GetWalletAccountID

`func (o *SupportEntitlementDto) GetWalletAccountID() string`

GetWalletAccountID returns the WalletAccountID field if non-nil, zero value otherwise.

### GetWalletAccountIDOk

`func (o *SupportEntitlementDto) GetWalletAccountIDOk() (*string, bool)`

GetWalletAccountIDOk returns a tuple with the WalletAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountID

`func (o *SupportEntitlementDto) SetWalletAccountID(v string)`

SetWalletAccountID sets WalletAccountID field to given value.

### HasWalletAccountID

`func (o *SupportEntitlementDto) HasWalletAccountID() bool`

HasWalletAccountID returns a boolean if a field has been set.

### SetWalletAccountIDNil

`func (o *SupportEntitlementDto) SetWalletAccountIDNil(b bool)`

 SetWalletAccountIDNil sets the value for WalletAccountID to be an explicit nil

### UnsetWalletAccountID
`func (o *SupportEntitlementDto) UnsetWalletAccountID()`

UnsetWalletAccountID ensures that no value is present for WalletAccountID, not even an explicit nil
### GetSecurityCertificateID

`func (o *SupportEntitlementDto) GetSecurityCertificateID() string`

GetSecurityCertificateID returns the SecurityCertificateID field if non-nil, zero value otherwise.

### GetSecurityCertificateIDOk

`func (o *SupportEntitlementDto) GetSecurityCertificateIDOk() (*string, bool)`

GetSecurityCertificateIDOk returns a tuple with the SecurityCertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCertificateID

`func (o *SupportEntitlementDto) SetSecurityCertificateID(v string)`

SetSecurityCertificateID sets SecurityCertificateID field to given value.

### HasSecurityCertificateID

`func (o *SupportEntitlementDto) HasSecurityCertificateID() bool`

HasSecurityCertificateID returns a boolean if a field has been set.

### SetSecurityCertificateIDNil

`func (o *SupportEntitlementDto) SetSecurityCertificateIDNil(b bool)`

 SetSecurityCertificateIDNil sets the value for SecurityCertificateID to be an explicit nil

### UnsetSecurityCertificateID
`func (o *SupportEntitlementDto) UnsetSecurityCertificateID()`

UnsetSecurityCertificateID ensures that no value is present for SecurityCertificateID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


