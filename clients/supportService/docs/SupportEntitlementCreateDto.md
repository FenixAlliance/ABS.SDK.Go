# SupportEntitlementCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
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

### NewSupportEntitlementCreateDto

`func NewSupportEntitlementCreateDto() *SupportEntitlementCreateDto`

NewSupportEntitlementCreateDto instantiates a new SupportEntitlementCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportEntitlementCreateDtoWithDefaults

`func NewSupportEntitlementCreateDtoWithDefaults() *SupportEntitlementCreateDto`

NewSupportEntitlementCreateDtoWithDefaults instantiates a new SupportEntitlementCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportEntitlementCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportEntitlementCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportEntitlementCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportEntitlementCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SupportEntitlementCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportEntitlementCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportEntitlementCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportEntitlementCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *SupportEntitlementCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportEntitlementCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportEntitlementCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportEntitlementCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportEntitlementCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportEntitlementCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportEntitlementCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportEntitlementCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportEntitlementCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportEntitlementCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportEntitlementCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportEntitlementCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartDateTime

`func (o *SupportEntitlementCreateDto) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *SupportEntitlementCreateDto) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *SupportEntitlementCreateDto) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *SupportEntitlementCreateDto) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *SupportEntitlementCreateDto) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *SupportEntitlementCreateDto) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *SupportEntitlementCreateDto) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *SupportEntitlementCreateDto) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetNextInvoiceDateTime

`func (o *SupportEntitlementCreateDto) GetNextInvoiceDateTime() time.Time`

GetNextInvoiceDateTime returns the NextInvoiceDateTime field if non-nil, zero value otherwise.

### GetNextInvoiceDateTimeOk

`func (o *SupportEntitlementCreateDto) GetNextInvoiceDateTimeOk() (*time.Time, bool)`

GetNextInvoiceDateTimeOk returns a tuple with the NextInvoiceDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDateTime

`func (o *SupportEntitlementCreateDto) SetNextInvoiceDateTime(v time.Time)`

SetNextInvoiceDateTime sets NextInvoiceDateTime field to given value.

### HasNextInvoiceDateTime

`func (o *SupportEntitlementCreateDto) HasNextInvoiceDateTime() bool`

HasNextInvoiceDateTime returns a boolean if a field has been set.

### GetCode

`func (o *SupportEntitlementCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SupportEntitlementCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SupportEntitlementCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SupportEntitlementCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *SupportEntitlementCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SupportEntitlementCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSignature

`func (o *SupportEntitlementCreateDto) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SupportEntitlementCreateDto) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SupportEntitlementCreateDto) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SupportEntitlementCreateDto) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SupportEntitlementCreateDto) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SupportEntitlementCreateDto) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetQuantity

`func (o *SupportEntitlementCreateDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SupportEntitlementCreateDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SupportEntitlementCreateDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SupportEntitlementCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRepetitions

`func (o *SupportEntitlementCreateDto) GetRepetitions() int32`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *SupportEntitlementCreateDto) GetRepetitionsOk() (*int32, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *SupportEntitlementCreateDto) SetRepetitions(v int32)`

SetRepetitions sets Repetitions field to given value.

### HasRepetitions

`func (o *SupportEntitlementCreateDto) HasRepetitions() bool`

HasRepetitions returns a boolean if a field has been set.

### GetChargeAttempts

`func (o *SupportEntitlementCreateDto) GetChargeAttempts() int32`

GetChargeAttempts returns the ChargeAttempts field if non-nil, zero value otherwise.

### GetChargeAttemptsOk

`func (o *SupportEntitlementCreateDto) GetChargeAttemptsOk() (*int32, bool)`

GetChargeAttemptsOk returns a tuple with the ChargeAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAttempts

`func (o *SupportEntitlementCreateDto) SetChargeAttempts(v int32)`

SetChargeAttempts sets ChargeAttempts field to given value.

### HasChargeAttempts

`func (o *SupportEntitlementCreateDto) HasChargeAttempts() bool`

HasChargeAttempts returns a boolean if a field has been set.

### GetFreeTrialInDays

`func (o *SupportEntitlementCreateDto) GetFreeTrialInDays() int32`

GetFreeTrialInDays returns the FreeTrialInDays field if non-nil, zero value otherwise.

### GetFreeTrialInDaysOk

`func (o *SupportEntitlementCreateDto) GetFreeTrialInDaysOk() (*int32, bool)`

GetFreeTrialInDaysOk returns a tuple with the FreeTrialInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrialInDays

`func (o *SupportEntitlementCreateDto) SetFreeTrialInDays(v int32)`

SetFreeTrialInDays sets FreeTrialInDays field to given value.

### HasFreeTrialInDays

`func (o *SupportEntitlementCreateDto) HasFreeTrialInDays() bool`

HasFreeTrialInDays returns a boolean if a field has been set.

### GetGracePeriodInDays

`func (o *SupportEntitlementCreateDto) GetGracePeriodInDays() int32`

GetGracePeriodInDays returns the GracePeriodInDays field if non-nil, zero value otherwise.

### GetGracePeriodInDaysOk

`func (o *SupportEntitlementCreateDto) GetGracePeriodInDaysOk() (*int32, bool)`

GetGracePeriodInDaysOk returns a tuple with the GracePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodInDays

`func (o *SupportEntitlementCreateDto) SetGracePeriodInDays(v int32)`

SetGracePeriodInDays sets GracePeriodInDays field to given value.

### HasGracePeriodInDays

`func (o *SupportEntitlementCreateDto) HasGracePeriodInDays() bool`

HasGracePeriodInDays returns a boolean if a field has been set.

### GetCustomRenewalPeriod

`func (o *SupportEntitlementCreateDto) GetCustomRenewalPeriod() int32`

GetCustomRenewalPeriod returns the CustomRenewalPeriod field if non-nil, zero value otherwise.

### GetCustomRenewalPeriodOk

`func (o *SupportEntitlementCreateDto) GetCustomRenewalPeriodOk() (*int32, bool)`

GetCustomRenewalPeriodOk returns a tuple with the CustomRenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRenewalPeriod

`func (o *SupportEntitlementCreateDto) SetCustomRenewalPeriod(v int32)`

SetCustomRenewalPeriod sets CustomRenewalPeriod field to given value.

### HasCustomRenewalPeriod

`func (o *SupportEntitlementCreateDto) HasCustomRenewalPeriod() bool`

HasCustomRenewalPeriod returns a boolean if a field has been set.

### GetEnableAutomaticRenew

`func (o *SupportEntitlementCreateDto) GetEnableAutomaticRenew() bool`

GetEnableAutomaticRenew returns the EnableAutomaticRenew field if non-nil, zero value otherwise.

### GetEnableAutomaticRenewOk

`func (o *SupportEntitlementCreateDto) GetEnableAutomaticRenewOk() (*bool, bool)`

GetEnableAutomaticRenewOk returns a tuple with the EnableAutomaticRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticRenew

`func (o *SupportEntitlementCreateDto) SetEnableAutomaticRenew(v bool)`

SetEnableAutomaticRenew sets EnableAutomaticRenew field to given value.

### HasEnableAutomaticRenew

`func (o *SupportEntitlementCreateDto) HasEnableAutomaticRenew() bool`

HasEnableAutomaticRenew returns a boolean if a field has been set.

### GetEnableProRateBilling

`func (o *SupportEntitlementCreateDto) GetEnableProRateBilling() bool`

GetEnableProRateBilling returns the EnableProRateBilling field if non-nil, zero value otherwise.

### GetEnableProRateBillingOk

`func (o *SupportEntitlementCreateDto) GetEnableProRateBillingOk() (*bool, bool)`

GetEnableProRateBillingOk returns a tuple with the EnableProRateBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProRateBilling

`func (o *SupportEntitlementCreateDto) SetEnableProRateBilling(v bool)`

SetEnableProRateBilling sets EnableProRateBilling field to given value.

### HasEnableProRateBilling

`func (o *SupportEntitlementCreateDto) HasEnableProRateBilling() bool`

HasEnableProRateBilling returns a boolean if a field has been set.

### GetEnableUsageThreshold

`func (o *SupportEntitlementCreateDto) GetEnableUsageThreshold() bool`

GetEnableUsageThreshold returns the EnableUsageThreshold field if non-nil, zero value otherwise.

### GetEnableUsageThresholdOk

`func (o *SupportEntitlementCreateDto) GetEnableUsageThresholdOk() (*bool, bool)`

GetEnableUsageThresholdOk returns a tuple with the EnableUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUsageThreshold

`func (o *SupportEntitlementCreateDto) SetEnableUsageThreshold(v bool)`

SetEnableUsageThreshold sets EnableUsageThreshold field to given value.

### HasEnableUsageThreshold

`func (o *SupportEntitlementCreateDto) HasEnableUsageThreshold() bool`

HasEnableUsageThreshold returns a boolean if a field has been set.

### GetEnableAutomaticDisable

`func (o *SupportEntitlementCreateDto) GetEnableAutomaticDisable() bool`

GetEnableAutomaticDisable returns the EnableAutomaticDisable field if non-nil, zero value otherwise.

### GetEnableAutomaticDisableOk

`func (o *SupportEntitlementCreateDto) GetEnableAutomaticDisableOk() (*bool, bool)`

GetEnableAutomaticDisableOk returns a tuple with the EnableAutomaticDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticDisable

`func (o *SupportEntitlementCreateDto) SetEnableAutomaticDisable(v bool)`

SetEnableAutomaticDisable sets EnableAutomaticDisable field to given value.

### HasEnableAutomaticDisable

`func (o *SupportEntitlementCreateDto) HasEnableAutomaticDisable() bool`

HasEnableAutomaticDisable returns a boolean if a field has been set.

### GetEnableAutomaticPayments

`func (o *SupportEntitlementCreateDto) GetEnableAutomaticPayments() bool`

GetEnableAutomaticPayments returns the EnableAutomaticPayments field if non-nil, zero value otherwise.

### GetEnableAutomaticPaymentsOk

`func (o *SupportEntitlementCreateDto) GetEnableAutomaticPaymentsOk() (*bool, bool)`

GetEnableAutomaticPaymentsOk returns a tuple with the EnableAutomaticPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPayments

`func (o *SupportEntitlementCreateDto) SetEnableAutomaticPayments(v bool)`

SetEnableAutomaticPayments sets EnableAutomaticPayments field to given value.

### HasEnableAutomaticPayments

`func (o *SupportEntitlementCreateDto) HasEnableAutomaticPayments() bool`

HasEnableAutomaticPayments returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *SupportEntitlementCreateDto) GetUsageThreshold() int32`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *SupportEntitlementCreateDto) GetUsageThresholdOk() (*int32, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *SupportEntitlementCreateDto) SetUsageThreshold(v int32)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *SupportEntitlementCreateDto) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### GetData

`func (o *SupportEntitlementCreateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SupportEntitlementCreateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SupportEntitlementCreateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SupportEntitlementCreateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SupportEntitlementCreateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SupportEntitlementCreateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *SupportEntitlementCreateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *SupportEntitlementCreateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *SupportEntitlementCreateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *SupportEntitlementCreateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *SupportEntitlementCreateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *SupportEntitlementCreateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *SupportEntitlementCreateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *SupportEntitlementCreateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *SupportEntitlementCreateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *SupportEntitlementCreateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *SupportEntitlementCreateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *SupportEntitlementCreateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *SupportEntitlementCreateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *SupportEntitlementCreateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *SupportEntitlementCreateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *SupportEntitlementCreateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *SupportEntitlementCreateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *SupportEntitlementCreateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *SupportEntitlementCreateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *SupportEntitlementCreateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *SupportEntitlementCreateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *SupportEntitlementCreateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *SupportEntitlementCreateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *SupportEntitlementCreateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *SupportEntitlementCreateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *SupportEntitlementCreateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *SupportEntitlementCreateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *SupportEntitlementCreateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *SupportEntitlementCreateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *SupportEntitlementCreateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *SupportEntitlementCreateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *SupportEntitlementCreateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *SupportEntitlementCreateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *SupportEntitlementCreateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *SupportEntitlementCreateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *SupportEntitlementCreateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *SupportEntitlementCreateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *SupportEntitlementCreateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *SupportEntitlementCreateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *SupportEntitlementCreateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *SupportEntitlementCreateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *SupportEntitlementCreateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *SupportEntitlementCreateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *SupportEntitlementCreateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *SupportEntitlementCreateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *SupportEntitlementCreateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *SupportEntitlementCreateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *SupportEntitlementCreateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *SupportEntitlementCreateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *SupportEntitlementCreateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *SupportEntitlementCreateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *SupportEntitlementCreateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *SupportEntitlementCreateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *SupportEntitlementCreateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *SupportEntitlementCreateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *SupportEntitlementCreateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *SupportEntitlementCreateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *SupportEntitlementCreateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *SupportEntitlementCreateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *SupportEntitlementCreateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *SupportEntitlementCreateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *SupportEntitlementCreateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *SupportEntitlementCreateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *SupportEntitlementCreateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *SupportEntitlementCreateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *SupportEntitlementCreateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *SupportEntitlementCreateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *SupportEntitlementCreateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *SupportEntitlementCreateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *SupportEntitlementCreateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *SupportEntitlementCreateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *SupportEntitlementCreateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *SupportEntitlementCreateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *SupportEntitlementCreateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *SupportEntitlementCreateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *SupportEntitlementCreateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *SupportEntitlementCreateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *SupportEntitlementCreateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *SupportEntitlementCreateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *SupportEntitlementCreateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *SupportEntitlementCreateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *SupportEntitlementCreateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *SupportEntitlementCreateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *SupportEntitlementCreateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *SupportEntitlementCreateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *SupportEntitlementCreateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *SupportEntitlementCreateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *SupportEntitlementCreateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *SupportEntitlementCreateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *SupportEntitlementCreateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *SupportEntitlementCreateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *SupportEntitlementCreateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *SupportEntitlementCreateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *SupportEntitlementCreateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *SupportEntitlementCreateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *SupportEntitlementCreateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *SupportEntitlementCreateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *SupportEntitlementCreateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *SupportEntitlementCreateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *SupportEntitlementCreateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *SupportEntitlementCreateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *SupportEntitlementCreateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *SupportEntitlementCreateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *SupportEntitlementCreateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *SupportEntitlementCreateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *SupportEntitlementCreateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *SupportEntitlementCreateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *SupportEntitlementCreateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *SupportEntitlementCreateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *SupportEntitlementCreateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *SupportEntitlementCreateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *SupportEntitlementCreateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *SupportEntitlementCreateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *SupportEntitlementCreateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetAccountHolderID

`func (o *SupportEntitlementCreateDto) GetAccountHolderID() string`

GetAccountHolderID returns the AccountHolderID field if non-nil, zero value otherwise.

### GetAccountHolderIDOk

`func (o *SupportEntitlementCreateDto) GetAccountHolderIDOk() (*string, bool)`

GetAccountHolderIDOk returns a tuple with the AccountHolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderID

`func (o *SupportEntitlementCreateDto) SetAccountHolderID(v string)`

SetAccountHolderID sets AccountHolderID field to given value.

### HasAccountHolderID

`func (o *SupportEntitlementCreateDto) HasAccountHolderID() bool`

HasAccountHolderID returns a boolean if a field has been set.

### SetAccountHolderIDNil

`func (o *SupportEntitlementCreateDto) SetAccountHolderIDNil(b bool)`

 SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil

### UnsetAccountHolderID
`func (o *SupportEntitlementCreateDto) UnsetAccountHolderID()`

UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil
### GetIndividualID

`func (o *SupportEntitlementCreateDto) GetIndividualID() string`

GetIndividualID returns the IndividualID field if non-nil, zero value otherwise.

### GetIndividualIDOk

`func (o *SupportEntitlementCreateDto) GetIndividualIDOk() (*string, bool)`

GetIndividualIDOk returns a tuple with the IndividualID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualID

`func (o *SupportEntitlementCreateDto) SetIndividualID(v string)`

SetIndividualID sets IndividualID field to given value.

### HasIndividualID

`func (o *SupportEntitlementCreateDto) HasIndividualID() bool`

HasIndividualID returns a boolean if a field has been set.

### SetIndividualIDNil

`func (o *SupportEntitlementCreateDto) SetIndividualIDNil(b bool)`

 SetIndividualIDNil sets the value for IndividualID to be an explicit nil

### UnsetIndividualID
`func (o *SupportEntitlementCreateDto) UnsetIndividualID()`

UnsetIndividualID ensures that no value is present for IndividualID, not even an explicit nil
### GetOrganizationID

`func (o *SupportEntitlementCreateDto) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *SupportEntitlementCreateDto) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *SupportEntitlementCreateDto) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *SupportEntitlementCreateDto) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### SetOrganizationIDNil

`func (o *SupportEntitlementCreateDto) SetOrganizationIDNil(b bool)`

 SetOrganizationIDNil sets the value for OrganizationID to be an explicit nil

### UnsetOrganizationID
`func (o *SupportEntitlementCreateDto) UnsetOrganizationID()`

UnsetOrganizationID ensures that no value is present for OrganizationID, not even an explicit nil
### GetReceiverBusinessID

`func (o *SupportEntitlementCreateDto) GetReceiverBusinessID() string`

GetReceiverBusinessID returns the ReceiverBusinessID field if non-nil, zero value otherwise.

### GetReceiverBusinessIDOk

`func (o *SupportEntitlementCreateDto) GetReceiverBusinessIDOk() (*string, bool)`

GetReceiverBusinessIDOk returns a tuple with the ReceiverBusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessID

`func (o *SupportEntitlementCreateDto) SetReceiverBusinessID(v string)`

SetReceiverBusinessID sets ReceiverBusinessID field to given value.

### HasReceiverBusinessID

`func (o *SupportEntitlementCreateDto) HasReceiverBusinessID() bool`

HasReceiverBusinessID returns a boolean if a field has been set.

### SetReceiverBusinessIDNil

`func (o *SupportEntitlementCreateDto) SetReceiverBusinessIDNil(b bool)`

 SetReceiverBusinessIDNil sets the value for ReceiverBusinessID to be an explicit nil

### UnsetReceiverBusinessID
`func (o *SupportEntitlementCreateDto) UnsetReceiverBusinessID()`

UnsetReceiverBusinessID ensures that no value is present for ReceiverBusinessID, not even an explicit nil
### GetBusinessID

`func (o *SupportEntitlementCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportEntitlementCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportEntitlementCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportEntitlementCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportEntitlementCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportEntitlementCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *SupportEntitlementCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *SupportEntitlementCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *SupportEntitlementCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *SupportEntitlementCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *SupportEntitlementCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *SupportEntitlementCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetPaymentTokenID

`func (o *SupportEntitlementCreateDto) GetPaymentTokenID() string`

GetPaymentTokenID returns the PaymentTokenID field if non-nil, zero value otherwise.

### GetPaymentTokenIDOk

`func (o *SupportEntitlementCreateDto) GetPaymentTokenIDOk() (*string, bool)`

GetPaymentTokenIDOk returns a tuple with the PaymentTokenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTokenID

`func (o *SupportEntitlementCreateDto) SetPaymentTokenID(v string)`

SetPaymentTokenID sets PaymentTokenID field to given value.

### HasPaymentTokenID

`func (o *SupportEntitlementCreateDto) HasPaymentTokenID() bool`

HasPaymentTokenID returns a boolean if a field has been set.

### SetPaymentTokenIDNil

`func (o *SupportEntitlementCreateDto) SetPaymentTokenIDNil(b bool)`

 SetPaymentTokenIDNil sets the value for PaymentTokenID to be an explicit nil

### UnsetPaymentTokenID
`func (o *SupportEntitlementCreateDto) UnsetPaymentTokenID()`

UnsetPaymentTokenID ensures that no value is present for PaymentTokenID, not even an explicit nil
### GetWalletAccountID

`func (o *SupportEntitlementCreateDto) GetWalletAccountID() string`

GetWalletAccountID returns the WalletAccountID field if non-nil, zero value otherwise.

### GetWalletAccountIDOk

`func (o *SupportEntitlementCreateDto) GetWalletAccountIDOk() (*string, bool)`

GetWalletAccountIDOk returns a tuple with the WalletAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountID

`func (o *SupportEntitlementCreateDto) SetWalletAccountID(v string)`

SetWalletAccountID sets WalletAccountID field to given value.

### HasWalletAccountID

`func (o *SupportEntitlementCreateDto) HasWalletAccountID() bool`

HasWalletAccountID returns a boolean if a field has been set.

### SetWalletAccountIDNil

`func (o *SupportEntitlementCreateDto) SetWalletAccountIDNil(b bool)`

 SetWalletAccountIDNil sets the value for WalletAccountID to be an explicit nil

### UnsetWalletAccountID
`func (o *SupportEntitlementCreateDto) UnsetWalletAccountID()`

UnsetWalletAccountID ensures that no value is present for WalletAccountID, not even an explicit nil
### GetSecurityCertificateID

`func (o *SupportEntitlementCreateDto) GetSecurityCertificateID() string`

GetSecurityCertificateID returns the SecurityCertificateID field if non-nil, zero value otherwise.

### GetSecurityCertificateIDOk

`func (o *SupportEntitlementCreateDto) GetSecurityCertificateIDOk() (*string, bool)`

GetSecurityCertificateIDOk returns a tuple with the SecurityCertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCertificateID

`func (o *SupportEntitlementCreateDto) SetSecurityCertificateID(v string)`

SetSecurityCertificateID sets SecurityCertificateID field to given value.

### HasSecurityCertificateID

`func (o *SupportEntitlementCreateDto) HasSecurityCertificateID() bool`

HasSecurityCertificateID returns a boolean if a field has been set.

### SetSecurityCertificateIDNil

`func (o *SupportEntitlementCreateDto) SetSecurityCertificateIDNil(b bool)`

 SetSecurityCertificateIDNil sets the value for SecurityCertificateID to be an explicit nil

### UnsetSecurityCertificateID
`func (o *SupportEntitlementCreateDto) UnsetSecurityCertificateID()`

UnsetSecurityCertificateID ensures that no value is present for SecurityCertificateID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


