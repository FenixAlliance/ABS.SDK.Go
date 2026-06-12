# SupportEntitlementUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
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
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**ReceiverTenantId** | Pointer to **NullableString** |  | [optional] 
**PaymentTokenId** | Pointer to **NullableString** |  | [optional] 
**WalletAccountId** | Pointer to **NullableString** |  | [optional] 
**SecurityCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportEntitlementUpdateDto

`func NewSupportEntitlementUpdateDto() *SupportEntitlementUpdateDto`

NewSupportEntitlementUpdateDto instantiates a new SupportEntitlementUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportEntitlementUpdateDtoWithDefaults

`func NewSupportEntitlementUpdateDtoWithDefaults() *SupportEntitlementUpdateDto`

NewSupportEntitlementUpdateDtoWithDefaults instantiates a new SupportEntitlementUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SupportEntitlementUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportEntitlementUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportEntitlementUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportEntitlementUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportEntitlementUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportEntitlementUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportEntitlementUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportEntitlementUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportEntitlementUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportEntitlementUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportEntitlementUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportEntitlementUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndDateTime

`func (o *SupportEntitlementUpdateDto) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *SupportEntitlementUpdateDto) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *SupportEntitlementUpdateDto) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *SupportEntitlementUpdateDto) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetNextInvoiceDateTime

`func (o *SupportEntitlementUpdateDto) GetNextInvoiceDateTime() time.Time`

GetNextInvoiceDateTime returns the NextInvoiceDateTime field if non-nil, zero value otherwise.

### GetNextInvoiceDateTimeOk

`func (o *SupportEntitlementUpdateDto) GetNextInvoiceDateTimeOk() (*time.Time, bool)`

GetNextInvoiceDateTimeOk returns a tuple with the NextInvoiceDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDateTime

`func (o *SupportEntitlementUpdateDto) SetNextInvoiceDateTime(v time.Time)`

SetNextInvoiceDateTime sets NextInvoiceDateTime field to given value.

### HasNextInvoiceDateTime

`func (o *SupportEntitlementUpdateDto) HasNextInvoiceDateTime() bool`

HasNextInvoiceDateTime returns a boolean if a field has been set.

### GetCode

`func (o *SupportEntitlementUpdateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SupportEntitlementUpdateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SupportEntitlementUpdateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SupportEntitlementUpdateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *SupportEntitlementUpdateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SupportEntitlementUpdateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSignature

`func (o *SupportEntitlementUpdateDto) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SupportEntitlementUpdateDto) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SupportEntitlementUpdateDto) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SupportEntitlementUpdateDto) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *SupportEntitlementUpdateDto) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *SupportEntitlementUpdateDto) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetQuantity

`func (o *SupportEntitlementUpdateDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SupportEntitlementUpdateDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SupportEntitlementUpdateDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SupportEntitlementUpdateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRepetitions

`func (o *SupportEntitlementUpdateDto) GetRepetitions() int32`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *SupportEntitlementUpdateDto) GetRepetitionsOk() (*int32, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *SupportEntitlementUpdateDto) SetRepetitions(v int32)`

SetRepetitions sets Repetitions field to given value.

### HasRepetitions

`func (o *SupportEntitlementUpdateDto) HasRepetitions() bool`

HasRepetitions returns a boolean if a field has been set.

### GetChargeAttempts

`func (o *SupportEntitlementUpdateDto) GetChargeAttempts() int32`

GetChargeAttempts returns the ChargeAttempts field if non-nil, zero value otherwise.

### GetChargeAttemptsOk

`func (o *SupportEntitlementUpdateDto) GetChargeAttemptsOk() (*int32, bool)`

GetChargeAttemptsOk returns a tuple with the ChargeAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAttempts

`func (o *SupportEntitlementUpdateDto) SetChargeAttempts(v int32)`

SetChargeAttempts sets ChargeAttempts field to given value.

### HasChargeAttempts

`func (o *SupportEntitlementUpdateDto) HasChargeAttempts() bool`

HasChargeAttempts returns a boolean if a field has been set.

### GetFreeTrialInDays

`func (o *SupportEntitlementUpdateDto) GetFreeTrialInDays() int32`

GetFreeTrialInDays returns the FreeTrialInDays field if non-nil, zero value otherwise.

### GetFreeTrialInDaysOk

`func (o *SupportEntitlementUpdateDto) GetFreeTrialInDaysOk() (*int32, bool)`

GetFreeTrialInDaysOk returns a tuple with the FreeTrialInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrialInDays

`func (o *SupportEntitlementUpdateDto) SetFreeTrialInDays(v int32)`

SetFreeTrialInDays sets FreeTrialInDays field to given value.

### HasFreeTrialInDays

`func (o *SupportEntitlementUpdateDto) HasFreeTrialInDays() bool`

HasFreeTrialInDays returns a boolean if a field has been set.

### GetGracePeriodInDays

`func (o *SupportEntitlementUpdateDto) GetGracePeriodInDays() int32`

GetGracePeriodInDays returns the GracePeriodInDays field if non-nil, zero value otherwise.

### GetGracePeriodInDaysOk

`func (o *SupportEntitlementUpdateDto) GetGracePeriodInDaysOk() (*int32, bool)`

GetGracePeriodInDaysOk returns a tuple with the GracePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodInDays

`func (o *SupportEntitlementUpdateDto) SetGracePeriodInDays(v int32)`

SetGracePeriodInDays sets GracePeriodInDays field to given value.

### HasGracePeriodInDays

`func (o *SupportEntitlementUpdateDto) HasGracePeriodInDays() bool`

HasGracePeriodInDays returns a boolean if a field has been set.

### GetCustomRenewalPeriod

`func (o *SupportEntitlementUpdateDto) GetCustomRenewalPeriod() int32`

GetCustomRenewalPeriod returns the CustomRenewalPeriod field if non-nil, zero value otherwise.

### GetCustomRenewalPeriodOk

`func (o *SupportEntitlementUpdateDto) GetCustomRenewalPeriodOk() (*int32, bool)`

GetCustomRenewalPeriodOk returns a tuple with the CustomRenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRenewalPeriod

`func (o *SupportEntitlementUpdateDto) SetCustomRenewalPeriod(v int32)`

SetCustomRenewalPeriod sets CustomRenewalPeriod field to given value.

### HasCustomRenewalPeriod

`func (o *SupportEntitlementUpdateDto) HasCustomRenewalPeriod() bool`

HasCustomRenewalPeriod returns a boolean if a field has been set.

### GetEnableAutomaticRenew

`func (o *SupportEntitlementUpdateDto) GetEnableAutomaticRenew() bool`

GetEnableAutomaticRenew returns the EnableAutomaticRenew field if non-nil, zero value otherwise.

### GetEnableAutomaticRenewOk

`func (o *SupportEntitlementUpdateDto) GetEnableAutomaticRenewOk() (*bool, bool)`

GetEnableAutomaticRenewOk returns a tuple with the EnableAutomaticRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticRenew

`func (o *SupportEntitlementUpdateDto) SetEnableAutomaticRenew(v bool)`

SetEnableAutomaticRenew sets EnableAutomaticRenew field to given value.

### HasEnableAutomaticRenew

`func (o *SupportEntitlementUpdateDto) HasEnableAutomaticRenew() bool`

HasEnableAutomaticRenew returns a boolean if a field has been set.

### GetEnableProRateBilling

`func (o *SupportEntitlementUpdateDto) GetEnableProRateBilling() bool`

GetEnableProRateBilling returns the EnableProRateBilling field if non-nil, zero value otherwise.

### GetEnableProRateBillingOk

`func (o *SupportEntitlementUpdateDto) GetEnableProRateBillingOk() (*bool, bool)`

GetEnableProRateBillingOk returns a tuple with the EnableProRateBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProRateBilling

`func (o *SupportEntitlementUpdateDto) SetEnableProRateBilling(v bool)`

SetEnableProRateBilling sets EnableProRateBilling field to given value.

### HasEnableProRateBilling

`func (o *SupportEntitlementUpdateDto) HasEnableProRateBilling() bool`

HasEnableProRateBilling returns a boolean if a field has been set.

### GetEnableUsageThreshold

`func (o *SupportEntitlementUpdateDto) GetEnableUsageThreshold() bool`

GetEnableUsageThreshold returns the EnableUsageThreshold field if non-nil, zero value otherwise.

### GetEnableUsageThresholdOk

`func (o *SupportEntitlementUpdateDto) GetEnableUsageThresholdOk() (*bool, bool)`

GetEnableUsageThresholdOk returns a tuple with the EnableUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUsageThreshold

`func (o *SupportEntitlementUpdateDto) SetEnableUsageThreshold(v bool)`

SetEnableUsageThreshold sets EnableUsageThreshold field to given value.

### HasEnableUsageThreshold

`func (o *SupportEntitlementUpdateDto) HasEnableUsageThreshold() bool`

HasEnableUsageThreshold returns a boolean if a field has been set.

### GetEnableAutomaticDisable

`func (o *SupportEntitlementUpdateDto) GetEnableAutomaticDisable() bool`

GetEnableAutomaticDisable returns the EnableAutomaticDisable field if non-nil, zero value otherwise.

### GetEnableAutomaticDisableOk

`func (o *SupportEntitlementUpdateDto) GetEnableAutomaticDisableOk() (*bool, bool)`

GetEnableAutomaticDisableOk returns a tuple with the EnableAutomaticDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticDisable

`func (o *SupportEntitlementUpdateDto) SetEnableAutomaticDisable(v bool)`

SetEnableAutomaticDisable sets EnableAutomaticDisable field to given value.

### HasEnableAutomaticDisable

`func (o *SupportEntitlementUpdateDto) HasEnableAutomaticDisable() bool`

HasEnableAutomaticDisable returns a boolean if a field has been set.

### GetEnableAutomaticPayments

`func (o *SupportEntitlementUpdateDto) GetEnableAutomaticPayments() bool`

GetEnableAutomaticPayments returns the EnableAutomaticPayments field if non-nil, zero value otherwise.

### GetEnableAutomaticPaymentsOk

`func (o *SupportEntitlementUpdateDto) GetEnableAutomaticPaymentsOk() (*bool, bool)`

GetEnableAutomaticPaymentsOk returns a tuple with the EnableAutomaticPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPayments

`func (o *SupportEntitlementUpdateDto) SetEnableAutomaticPayments(v bool)`

SetEnableAutomaticPayments sets EnableAutomaticPayments field to given value.

### HasEnableAutomaticPayments

`func (o *SupportEntitlementUpdateDto) HasEnableAutomaticPayments() bool`

HasEnableAutomaticPayments returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *SupportEntitlementUpdateDto) GetUsageThreshold() int32`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *SupportEntitlementUpdateDto) GetUsageThresholdOk() (*int32, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *SupportEntitlementUpdateDto) SetUsageThreshold(v int32)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *SupportEntitlementUpdateDto) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### GetData

`func (o *SupportEntitlementUpdateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SupportEntitlementUpdateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SupportEntitlementUpdateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SupportEntitlementUpdateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SupportEntitlementUpdateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SupportEntitlementUpdateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *SupportEntitlementUpdateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *SupportEntitlementUpdateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *SupportEntitlementUpdateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *SupportEntitlementUpdateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *SupportEntitlementUpdateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *SupportEntitlementUpdateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *SupportEntitlementUpdateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *SupportEntitlementUpdateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *SupportEntitlementUpdateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *SupportEntitlementUpdateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *SupportEntitlementUpdateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *SupportEntitlementUpdateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *SupportEntitlementUpdateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *SupportEntitlementUpdateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *SupportEntitlementUpdateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *SupportEntitlementUpdateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *SupportEntitlementUpdateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *SupportEntitlementUpdateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *SupportEntitlementUpdateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *SupportEntitlementUpdateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *SupportEntitlementUpdateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *SupportEntitlementUpdateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *SupportEntitlementUpdateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *SupportEntitlementUpdateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *SupportEntitlementUpdateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *SupportEntitlementUpdateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *SupportEntitlementUpdateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *SupportEntitlementUpdateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *SupportEntitlementUpdateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *SupportEntitlementUpdateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *SupportEntitlementUpdateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *SupportEntitlementUpdateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *SupportEntitlementUpdateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *SupportEntitlementUpdateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *SupportEntitlementUpdateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *SupportEntitlementUpdateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *SupportEntitlementUpdateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *SupportEntitlementUpdateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *SupportEntitlementUpdateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *SupportEntitlementUpdateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *SupportEntitlementUpdateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *SupportEntitlementUpdateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *SupportEntitlementUpdateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *SupportEntitlementUpdateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *SupportEntitlementUpdateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *SupportEntitlementUpdateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *SupportEntitlementUpdateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *SupportEntitlementUpdateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *SupportEntitlementUpdateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *SupportEntitlementUpdateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *SupportEntitlementUpdateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *SupportEntitlementUpdateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *SupportEntitlementUpdateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *SupportEntitlementUpdateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *SupportEntitlementUpdateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *SupportEntitlementUpdateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *SupportEntitlementUpdateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *SupportEntitlementUpdateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *SupportEntitlementUpdateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *SupportEntitlementUpdateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *SupportEntitlementUpdateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *SupportEntitlementUpdateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *SupportEntitlementUpdateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *SupportEntitlementUpdateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *SupportEntitlementUpdateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *SupportEntitlementUpdateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *SupportEntitlementUpdateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *SupportEntitlementUpdateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *SupportEntitlementUpdateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *SupportEntitlementUpdateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *SupportEntitlementUpdateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *SupportEntitlementUpdateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *SupportEntitlementUpdateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *SupportEntitlementUpdateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *SupportEntitlementUpdateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *SupportEntitlementUpdateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *SupportEntitlementUpdateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *SupportEntitlementUpdateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *SupportEntitlementUpdateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *SupportEntitlementUpdateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *SupportEntitlementUpdateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *SupportEntitlementUpdateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *SupportEntitlementUpdateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *SupportEntitlementUpdateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *SupportEntitlementUpdateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *SupportEntitlementUpdateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *SupportEntitlementUpdateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *SupportEntitlementUpdateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *SupportEntitlementUpdateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *SupportEntitlementUpdateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *SupportEntitlementUpdateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *SupportEntitlementUpdateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *SupportEntitlementUpdateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *SupportEntitlementUpdateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *SupportEntitlementUpdateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *SupportEntitlementUpdateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *SupportEntitlementUpdateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *SupportEntitlementUpdateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *SupportEntitlementUpdateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *SupportEntitlementUpdateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *SupportEntitlementUpdateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *SupportEntitlementUpdateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *SupportEntitlementUpdateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *SupportEntitlementUpdateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *SupportEntitlementUpdateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *SupportEntitlementUpdateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *SupportEntitlementUpdateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *SupportEntitlementUpdateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *SupportEntitlementUpdateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *SupportEntitlementUpdateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *SupportEntitlementUpdateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *SupportEntitlementUpdateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *SupportEntitlementUpdateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *SupportEntitlementUpdateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetIndividualId

`func (o *SupportEntitlementUpdateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *SupportEntitlementUpdateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *SupportEntitlementUpdateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *SupportEntitlementUpdateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *SupportEntitlementUpdateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *SupportEntitlementUpdateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *SupportEntitlementUpdateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SupportEntitlementUpdateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SupportEntitlementUpdateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SupportEntitlementUpdateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *SupportEntitlementUpdateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *SupportEntitlementUpdateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverTenantId

`func (o *SupportEntitlementUpdateDto) GetReceiverTenantId() string`

GetReceiverTenantId returns the ReceiverTenantId field if non-nil, zero value otherwise.

### GetReceiverTenantIdOk

`func (o *SupportEntitlementUpdateDto) GetReceiverTenantIdOk() (*string, bool)`

GetReceiverTenantIdOk returns a tuple with the ReceiverTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverTenantId

`func (o *SupportEntitlementUpdateDto) SetReceiverTenantId(v string)`

SetReceiverTenantId sets ReceiverTenantId field to given value.

### HasReceiverTenantId

`func (o *SupportEntitlementUpdateDto) HasReceiverTenantId() bool`

HasReceiverTenantId returns a boolean if a field has been set.

### SetReceiverTenantIdNil

`func (o *SupportEntitlementUpdateDto) SetReceiverTenantIdNil(b bool)`

 SetReceiverTenantIdNil sets the value for ReceiverTenantId to be an explicit nil

### UnsetReceiverTenantId
`func (o *SupportEntitlementUpdateDto) UnsetReceiverTenantId()`

UnsetReceiverTenantId ensures that no value is present for ReceiverTenantId, not even an explicit nil
### GetPaymentTokenId

`func (o *SupportEntitlementUpdateDto) GetPaymentTokenId() string`

GetPaymentTokenId returns the PaymentTokenId field if non-nil, zero value otherwise.

### GetPaymentTokenIdOk

`func (o *SupportEntitlementUpdateDto) GetPaymentTokenIdOk() (*string, bool)`

GetPaymentTokenIdOk returns a tuple with the PaymentTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTokenId

`func (o *SupportEntitlementUpdateDto) SetPaymentTokenId(v string)`

SetPaymentTokenId sets PaymentTokenId field to given value.

### HasPaymentTokenId

`func (o *SupportEntitlementUpdateDto) HasPaymentTokenId() bool`

HasPaymentTokenId returns a boolean if a field has been set.

### SetPaymentTokenIdNil

`func (o *SupportEntitlementUpdateDto) SetPaymentTokenIdNil(b bool)`

 SetPaymentTokenIdNil sets the value for PaymentTokenId to be an explicit nil

### UnsetPaymentTokenId
`func (o *SupportEntitlementUpdateDto) UnsetPaymentTokenId()`

UnsetPaymentTokenId ensures that no value is present for PaymentTokenId, not even an explicit nil
### GetWalletAccountId

`func (o *SupportEntitlementUpdateDto) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *SupportEntitlementUpdateDto) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *SupportEntitlementUpdateDto) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.

### HasWalletAccountId

`func (o *SupportEntitlementUpdateDto) HasWalletAccountId() bool`

HasWalletAccountId returns a boolean if a field has been set.

### SetWalletAccountIdNil

`func (o *SupportEntitlementUpdateDto) SetWalletAccountIdNil(b bool)`

 SetWalletAccountIdNil sets the value for WalletAccountId to be an explicit nil

### UnsetWalletAccountId
`func (o *SupportEntitlementUpdateDto) UnsetWalletAccountId()`

UnsetWalletAccountId ensures that no value is present for WalletAccountId, not even an explicit nil
### GetSecurityCertificateId

`func (o *SupportEntitlementUpdateDto) GetSecurityCertificateId() string`

GetSecurityCertificateId returns the SecurityCertificateId field if non-nil, zero value otherwise.

### GetSecurityCertificateIdOk

`func (o *SupportEntitlementUpdateDto) GetSecurityCertificateIdOk() (*string, bool)`

GetSecurityCertificateIdOk returns a tuple with the SecurityCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCertificateId

`func (o *SupportEntitlementUpdateDto) SetSecurityCertificateId(v string)`

SetSecurityCertificateId sets SecurityCertificateId field to given value.

### HasSecurityCertificateId

`func (o *SupportEntitlementUpdateDto) HasSecurityCertificateId() bool`

HasSecurityCertificateId returns a boolean if a field has been set.

### SetSecurityCertificateIdNil

`func (o *SupportEntitlementUpdateDto) SetSecurityCertificateIdNil(b bool)`

 SetSecurityCertificateIdNil sets the value for SecurityCertificateId to be an explicit nil

### UnsetSecurityCertificateId
`func (o *SupportEntitlementUpdateDto) UnsetSecurityCertificateId()`

UnsetSecurityCertificateId ensures that no value is present for SecurityCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


