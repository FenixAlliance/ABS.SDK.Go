# PaymentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TimeStamp** | Pointer to **time.Time** |  | [optional] 
**Test** | Pointer to **bool** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EmisorWalletId** | Pointer to **NullableString** |  | [optional] 
**ReceiverWalletId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**TotalCost** | Pointer to **float64** |  | [optional] 
**TotalTaxes** | Pointer to **float64** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**DataLabel** | Pointer to **NullableString** |  | [optional] 
**Data1** | Pointer to **NullableString** |  | [optional] 
**Data1Label** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to **NullableString** |  | [optional] 
**Authorization** | Pointer to **NullableString** |  | [optional] 
**ReferenceCode** | Pointer to **NullableString** |  | [optional] 
**CorrelationCode** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**OnBehalfOf** | Pointer to **int32** |  | [optional] 
**PaymentType** | Pointer to **int32** |  | [optional] 
**PaymentStatus** | Pointer to **int32** |  | [optional] 
**BaseCost** | Pointer to **float64** |  | [optional] 
**Signature** | Pointer to **NullableString** |  | [optional] 
**SignatureMismatch** | Pointer to **bool** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**MarkedForRevision** | Pointer to **bool** |  | [optional] 
**ForexRatesSnapshot** | Pointer to **NullableString** |  | [optional] 
**OfficialId** | Pointer to **NullableString** |  | [optional] 
**OfficialIdExpeditionDate** | Pointer to **time.Time** |  | [optional] 
**FiscalIdentificationTypeId** | Pointer to **NullableString** |  | [optional] 
**BillingAddress** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Cellphone** | Pointer to **NullableString** |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**LocationId** | Pointer to **NullableString** |  | [optional] 
**EntitlementId** | Pointer to **NullableString** |  | [optional] 
**AntiFraudScore** | Pointer to **float64** |  | [optional] 
**CallRecordURL** | Pointer to **NullableString** |  | [optional] 
**Called** | Pointer to **bool** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**PayerPictureTimestamp** | Pointer to **NullableString** |  | [optional] 
**PayerPicture** | Pointer to **NullableString** |  | [optional] 
**IdentificationPictureTimestamp** | Pointer to **NullableString** |  | [optional] 
**IdentificationPicture** | Pointer to **NullableString** |  | [optional] 
**IdentificationBackPicture** | Pointer to **NullableString** |  | [optional] 
**IdentificationBackPictureTimestamp** | Pointer to **NullableString** |  | [optional] 
**IpLookupId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**AccountingEntryId** | Pointer to **NullableString** |  | [optional] 
**PaymentGatewayId** | Pointer to **NullableString** |  | [optional] 
**BankAccountId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**BankId** | Pointer to **NullableString** |  | [optional] 
**PaymentTokenId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentDto

`func NewPaymentDto() *PaymentDto`

NewPaymentDto instantiates a new PaymentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDtoWithDefaults

`func NewPaymentDtoWithDefaults() *PaymentDto`

NewPaymentDtoWithDefaults instantiates a new PaymentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PaymentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PaymentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PaymentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTimeStamp

`func (o *PaymentDto) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *PaymentDto) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *PaymentDto) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *PaymentDto) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetTest

`func (o *PaymentDto) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PaymentDto) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PaymentDto) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *PaymentDto) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetInvoiceId

`func (o *PaymentDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PaymentDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *PaymentDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *PaymentDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetTenantId

`func (o *PaymentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PaymentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PaymentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PaymentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEmisorWalletId

`func (o *PaymentDto) GetEmisorWalletId() string`

GetEmisorWalletId returns the EmisorWalletId field if non-nil, zero value otherwise.

### GetEmisorWalletIdOk

`func (o *PaymentDto) GetEmisorWalletIdOk() (*string, bool)`

GetEmisorWalletIdOk returns a tuple with the EmisorWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletId

`func (o *PaymentDto) SetEmisorWalletId(v string)`

SetEmisorWalletId sets EmisorWalletId field to given value.

### HasEmisorWalletId

`func (o *PaymentDto) HasEmisorWalletId() bool`

HasEmisorWalletId returns a boolean if a field has been set.

### SetEmisorWalletIdNil

`func (o *PaymentDto) SetEmisorWalletIdNil(b bool)`

 SetEmisorWalletIdNil sets the value for EmisorWalletId to be an explicit nil

### UnsetEmisorWalletId
`func (o *PaymentDto) UnsetEmisorWalletId()`

UnsetEmisorWalletId ensures that no value is present for EmisorWalletId, not even an explicit nil
### GetReceiverWalletId

`func (o *PaymentDto) GetReceiverWalletId() string`

GetReceiverWalletId returns the ReceiverWalletId field if non-nil, zero value otherwise.

### GetReceiverWalletIdOk

`func (o *PaymentDto) GetReceiverWalletIdOk() (*string, bool)`

GetReceiverWalletIdOk returns a tuple with the ReceiverWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletId

`func (o *PaymentDto) SetReceiverWalletId(v string)`

SetReceiverWalletId sets ReceiverWalletId field to given value.

### HasReceiverWalletId

`func (o *PaymentDto) HasReceiverWalletId() bool`

HasReceiverWalletId returns a boolean if a field has been set.

### SetReceiverWalletIdNil

`func (o *PaymentDto) SetReceiverWalletIdNil(b bool)`

 SetReceiverWalletIdNil sets the value for ReceiverWalletId to be an explicit nil

### UnsetReceiverWalletId
`func (o *PaymentDto) UnsetReceiverWalletId()`

UnsetReceiverWalletId ensures that no value is present for ReceiverWalletId, not even an explicit nil
### GetCurrencyId

`func (o *PaymentDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PaymentDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PaymentDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PaymentDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *PaymentDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *PaymentDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetForexRate

`func (o *PaymentDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *PaymentDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *PaymentDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *PaymentDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTotalCost

`func (o *PaymentDto) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *PaymentDto) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *PaymentDto) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *PaymentDto) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *PaymentDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *PaymentDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *PaymentDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *PaymentDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetClosed

`func (o *PaymentDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *PaymentDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *PaymentDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *PaymentDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetData

`func (o *PaymentDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PaymentDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PaymentDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *PaymentDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *PaymentDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *PaymentDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *PaymentDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *PaymentDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *PaymentDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *PaymentDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *PaymentDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *PaymentDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *PaymentDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *PaymentDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *PaymentDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *PaymentDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *PaymentDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *PaymentDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *PaymentDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *PaymentDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *PaymentDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetResponse

`func (o *PaymentDto) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *PaymentDto) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *PaymentDto) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *PaymentDto) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *PaymentDto) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *PaymentDto) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetAuthorization

`func (o *PaymentDto) GetAuthorization() string`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *PaymentDto) GetAuthorizationOk() (*string, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *PaymentDto) SetAuthorization(v string)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *PaymentDto) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### SetAuthorizationNil

`func (o *PaymentDto) SetAuthorizationNil(b bool)`

 SetAuthorizationNil sets the value for Authorization to be an explicit nil

### UnsetAuthorization
`func (o *PaymentDto) UnsetAuthorization()`

UnsetAuthorization ensures that no value is present for Authorization, not even an explicit nil
### GetReferenceCode

`func (o *PaymentDto) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *PaymentDto) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *PaymentDto) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *PaymentDto) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.

### SetReferenceCodeNil

`func (o *PaymentDto) SetReferenceCodeNil(b bool)`

 SetReferenceCodeNil sets the value for ReferenceCode to be an explicit nil

### UnsetReferenceCode
`func (o *PaymentDto) UnsetReferenceCode()`

UnsetReferenceCode ensures that no value is present for ReferenceCode, not even an explicit nil
### GetCorrelationCode

`func (o *PaymentDto) GetCorrelationCode() string`

GetCorrelationCode returns the CorrelationCode field if non-nil, zero value otherwise.

### GetCorrelationCodeOk

`func (o *PaymentDto) GetCorrelationCodeOk() (*string, bool)`

GetCorrelationCodeOk returns a tuple with the CorrelationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationCode

`func (o *PaymentDto) SetCorrelationCode(v string)`

SetCorrelationCode sets CorrelationCode field to given value.

### HasCorrelationCode

`func (o *PaymentDto) HasCorrelationCode() bool`

HasCorrelationCode returns a boolean if a field has been set.

### SetCorrelationCodeNil

`func (o *PaymentDto) SetCorrelationCodeNil(b bool)`

 SetCorrelationCodeNil sets the value for CorrelationCode to be an explicit nil

### UnsetCorrelationCode
`func (o *PaymentDto) UnsetCorrelationCode()`

UnsetCorrelationCode ensures that no value is present for CorrelationCode, not even an explicit nil
### GetLastUpdated

`func (o *PaymentDto) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PaymentDto) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PaymentDto) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PaymentDto) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOnBehalfOf

`func (o *PaymentDto) GetOnBehalfOf() int32`

GetOnBehalfOf returns the OnBehalfOf field if non-nil, zero value otherwise.

### GetOnBehalfOfOk

`func (o *PaymentDto) GetOnBehalfOfOk() (*int32, bool)`

GetOnBehalfOfOk returns a tuple with the OnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOf

`func (o *PaymentDto) SetOnBehalfOf(v int32)`

SetOnBehalfOf sets OnBehalfOf field to given value.

### HasOnBehalfOf

`func (o *PaymentDto) HasOnBehalfOf() bool`

HasOnBehalfOf returns a boolean if a field has been set.

### GetPaymentType

`func (o *PaymentDto) GetPaymentType() int32`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentDto) GetPaymentTypeOk() (*int32, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentDto) SetPaymentType(v int32)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PaymentDto) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *PaymentDto) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PaymentDto) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PaymentDto) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *PaymentDto) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetBaseCost

`func (o *PaymentDto) GetBaseCost() float64`

GetBaseCost returns the BaseCost field if non-nil, zero value otherwise.

### GetBaseCostOk

`func (o *PaymentDto) GetBaseCostOk() (*float64, bool)`

GetBaseCostOk returns a tuple with the BaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCost

`func (o *PaymentDto) SetBaseCost(v float64)`

SetBaseCost sets BaseCost field to given value.

### HasBaseCost

`func (o *PaymentDto) HasBaseCost() bool`

HasBaseCost returns a boolean if a field has been set.

### GetSignature

`func (o *PaymentDto) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentDto) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentDto) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *PaymentDto) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *PaymentDto) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *PaymentDto) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetSignatureMismatch

`func (o *PaymentDto) GetSignatureMismatch() bool`

GetSignatureMismatch returns the SignatureMismatch field if non-nil, zero value otherwise.

### GetSignatureMismatchOk

`func (o *PaymentDto) GetSignatureMismatchOk() (*bool, bool)`

GetSignatureMismatchOk returns a tuple with the SignatureMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureMismatch

`func (o *PaymentDto) SetSignatureMismatch(v bool)`

SetSignatureMismatch sets SignatureMismatch field to given value.

### HasSignatureMismatch

`func (o *PaymentDto) HasSignatureMismatch() bool`

HasSignatureMismatch returns a boolean if a field has been set.

### GetIsExternal

`func (o *PaymentDto) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *PaymentDto) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *PaymentDto) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *PaymentDto) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetMarkedForRevision

`func (o *PaymentDto) GetMarkedForRevision() bool`

GetMarkedForRevision returns the MarkedForRevision field if non-nil, zero value otherwise.

### GetMarkedForRevisionOk

`func (o *PaymentDto) GetMarkedForRevisionOk() (*bool, bool)`

GetMarkedForRevisionOk returns a tuple with the MarkedForRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForRevision

`func (o *PaymentDto) SetMarkedForRevision(v bool)`

SetMarkedForRevision sets MarkedForRevision field to given value.

### HasMarkedForRevision

`func (o *PaymentDto) HasMarkedForRevision() bool`

HasMarkedForRevision returns a boolean if a field has been set.

### GetForexRatesSnapshot

`func (o *PaymentDto) GetForexRatesSnapshot() string`

GetForexRatesSnapshot returns the ForexRatesSnapshot field if non-nil, zero value otherwise.

### GetForexRatesSnapshotOk

`func (o *PaymentDto) GetForexRatesSnapshotOk() (*string, bool)`

GetForexRatesSnapshotOk returns a tuple with the ForexRatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRatesSnapshot

`func (o *PaymentDto) SetForexRatesSnapshot(v string)`

SetForexRatesSnapshot sets ForexRatesSnapshot field to given value.

### HasForexRatesSnapshot

`func (o *PaymentDto) HasForexRatesSnapshot() bool`

HasForexRatesSnapshot returns a boolean if a field has been set.

### SetForexRatesSnapshotNil

`func (o *PaymentDto) SetForexRatesSnapshotNil(b bool)`

 SetForexRatesSnapshotNil sets the value for ForexRatesSnapshot to be an explicit nil

### UnsetForexRatesSnapshot
`func (o *PaymentDto) UnsetForexRatesSnapshot()`

UnsetForexRatesSnapshot ensures that no value is present for ForexRatesSnapshot, not even an explicit nil
### GetOfficialId

`func (o *PaymentDto) GetOfficialId() string`

GetOfficialId returns the OfficialId field if non-nil, zero value otherwise.

### GetOfficialIdOk

`func (o *PaymentDto) GetOfficialIdOk() (*string, bool)`

GetOfficialIdOk returns a tuple with the OfficialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialId

`func (o *PaymentDto) SetOfficialId(v string)`

SetOfficialId sets OfficialId field to given value.

### HasOfficialId

`func (o *PaymentDto) HasOfficialId() bool`

HasOfficialId returns a boolean if a field has been set.

### SetOfficialIdNil

`func (o *PaymentDto) SetOfficialIdNil(b bool)`

 SetOfficialIdNil sets the value for OfficialId to be an explicit nil

### UnsetOfficialId
`func (o *PaymentDto) UnsetOfficialId()`

UnsetOfficialId ensures that no value is present for OfficialId, not even an explicit nil
### GetOfficialIdExpeditionDate

`func (o *PaymentDto) GetOfficialIdExpeditionDate() time.Time`

GetOfficialIdExpeditionDate returns the OfficialIdExpeditionDate field if non-nil, zero value otherwise.

### GetOfficialIdExpeditionDateOk

`func (o *PaymentDto) GetOfficialIdExpeditionDateOk() (*time.Time, bool)`

GetOfficialIdExpeditionDateOk returns a tuple with the OfficialIdExpeditionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialIdExpeditionDate

`func (o *PaymentDto) SetOfficialIdExpeditionDate(v time.Time)`

SetOfficialIdExpeditionDate sets OfficialIdExpeditionDate field to given value.

### HasOfficialIdExpeditionDate

`func (o *PaymentDto) HasOfficialIdExpeditionDate() bool`

HasOfficialIdExpeditionDate returns a boolean if a field has been set.

### GetFiscalIdentificationTypeId

`func (o *PaymentDto) GetFiscalIdentificationTypeId() string`

GetFiscalIdentificationTypeId returns the FiscalIdentificationTypeId field if non-nil, zero value otherwise.

### GetFiscalIdentificationTypeIdOk

`func (o *PaymentDto) GetFiscalIdentificationTypeIdOk() (*string, bool)`

GetFiscalIdentificationTypeIdOk returns a tuple with the FiscalIdentificationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalIdentificationTypeId

`func (o *PaymentDto) SetFiscalIdentificationTypeId(v string)`

SetFiscalIdentificationTypeId sets FiscalIdentificationTypeId field to given value.

### HasFiscalIdentificationTypeId

`func (o *PaymentDto) HasFiscalIdentificationTypeId() bool`

HasFiscalIdentificationTypeId returns a boolean if a field has been set.

### SetFiscalIdentificationTypeIdNil

`func (o *PaymentDto) SetFiscalIdentificationTypeIdNil(b bool)`

 SetFiscalIdentificationTypeIdNil sets the value for FiscalIdentificationTypeId to be an explicit nil

### UnsetFiscalIdentificationTypeId
`func (o *PaymentDto) UnsetFiscalIdentificationTypeId()`

UnsetFiscalIdentificationTypeId ensures that no value is present for FiscalIdentificationTypeId, not even an explicit nil
### GetBillingAddress

`func (o *PaymentDto) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentDto) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentDto) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *PaymentDto) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *PaymentDto) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *PaymentDto) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetPhone

`func (o *PaymentDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PaymentDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PaymentDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PaymentDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PaymentDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PaymentDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCellphone

`func (o *PaymentDto) GetCellphone() string`

GetCellphone returns the Cellphone field if non-nil, zero value otherwise.

### GetCellphoneOk

`func (o *PaymentDto) GetCellphoneOk() (*string, bool)`

GetCellphoneOk returns a tuple with the Cellphone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphone

`func (o *PaymentDto) SetCellphone(v string)`

SetCellphone sets Cellphone field to given value.

### HasCellphone

`func (o *PaymentDto) HasCellphone() bool`

HasCellphone returns a boolean if a field has been set.

### SetCellphoneNil

`func (o *PaymentDto) SetCellphoneNil(b bool)`

 SetCellphoneNil sets the value for Cellphone to be an explicit nil

### UnsetCellphone
`func (o *PaymentDto) UnsetCellphone()`

UnsetCellphone ensures that no value is present for Cellphone, not even an explicit nil
### GetDepartment

`func (o *PaymentDto) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *PaymentDto) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *PaymentDto) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *PaymentDto) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *PaymentDto) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *PaymentDto) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetCity

`func (o *PaymentDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PaymentDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PaymentDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PaymentDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *PaymentDto) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *PaymentDto) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountryId

`func (o *PaymentDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *PaymentDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *PaymentDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *PaymentDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *PaymentDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *PaymentDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetLocationId

`func (o *PaymentDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *PaymentDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *PaymentDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *PaymentDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *PaymentDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *PaymentDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetEntitlementId

`func (o *PaymentDto) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *PaymentDto) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *PaymentDto) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.

### HasEntitlementId

`func (o *PaymentDto) HasEntitlementId() bool`

HasEntitlementId returns a boolean if a field has been set.

### SetEntitlementIdNil

`func (o *PaymentDto) SetEntitlementIdNil(b bool)`

 SetEntitlementIdNil sets the value for EntitlementId to be an explicit nil

### UnsetEntitlementId
`func (o *PaymentDto) UnsetEntitlementId()`

UnsetEntitlementId ensures that no value is present for EntitlementId, not even an explicit nil
### GetAntiFraudScore

`func (o *PaymentDto) GetAntiFraudScore() float64`

GetAntiFraudScore returns the AntiFraudScore field if non-nil, zero value otherwise.

### GetAntiFraudScoreOk

`func (o *PaymentDto) GetAntiFraudScoreOk() (*float64, bool)`

GetAntiFraudScoreOk returns a tuple with the AntiFraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudScore

`func (o *PaymentDto) SetAntiFraudScore(v float64)`

SetAntiFraudScore sets AntiFraudScore field to given value.

### HasAntiFraudScore

`func (o *PaymentDto) HasAntiFraudScore() bool`

HasAntiFraudScore returns a boolean if a field has been set.

### GetCallRecordURL

`func (o *PaymentDto) GetCallRecordURL() string`

GetCallRecordURL returns the CallRecordURL field if non-nil, zero value otherwise.

### GetCallRecordURLOk

`func (o *PaymentDto) GetCallRecordURLOk() (*string, bool)`

GetCallRecordURLOk returns a tuple with the CallRecordURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordURL

`func (o *PaymentDto) SetCallRecordURL(v string)`

SetCallRecordURL sets CallRecordURL field to given value.

### HasCallRecordURL

`func (o *PaymentDto) HasCallRecordURL() bool`

HasCallRecordURL returns a boolean if a field has been set.

### SetCallRecordURLNil

`func (o *PaymentDto) SetCallRecordURLNil(b bool)`

 SetCallRecordURLNil sets the value for CallRecordURL to be an explicit nil

### UnsetCallRecordURL
`func (o *PaymentDto) UnsetCallRecordURL()`

UnsetCallRecordURL ensures that no value is present for CallRecordURL, not even an explicit nil
### GetCalled

`func (o *PaymentDto) GetCalled() bool`

GetCalled returns the Called field if non-nil, zero value otherwise.

### GetCalledOk

`func (o *PaymentDto) GetCalledOk() (*bool, bool)`

GetCalledOk returns a tuple with the Called field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalled

`func (o *PaymentDto) SetCalled(v bool)`

SetCalled sets Called field to given value.

### HasCalled

`func (o *PaymentDto) HasCalled() bool`

HasCalled returns a boolean if a field has been set.

### GetVerified

`func (o *PaymentDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PaymentDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PaymentDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *PaymentDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetPayerPictureTimestamp

`func (o *PaymentDto) GetPayerPictureTimestamp() string`

GetPayerPictureTimestamp returns the PayerPictureTimestamp field if non-nil, zero value otherwise.

### GetPayerPictureTimestampOk

`func (o *PaymentDto) GetPayerPictureTimestampOk() (*string, bool)`

GetPayerPictureTimestampOk returns a tuple with the PayerPictureTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerPictureTimestamp

`func (o *PaymentDto) SetPayerPictureTimestamp(v string)`

SetPayerPictureTimestamp sets PayerPictureTimestamp field to given value.

### HasPayerPictureTimestamp

`func (o *PaymentDto) HasPayerPictureTimestamp() bool`

HasPayerPictureTimestamp returns a boolean if a field has been set.

### SetPayerPictureTimestampNil

`func (o *PaymentDto) SetPayerPictureTimestampNil(b bool)`

 SetPayerPictureTimestampNil sets the value for PayerPictureTimestamp to be an explicit nil

### UnsetPayerPictureTimestamp
`func (o *PaymentDto) UnsetPayerPictureTimestamp()`

UnsetPayerPictureTimestamp ensures that no value is present for PayerPictureTimestamp, not even an explicit nil
### GetPayerPicture

`func (o *PaymentDto) GetPayerPicture() string`

GetPayerPicture returns the PayerPicture field if non-nil, zero value otherwise.

### GetPayerPictureOk

`func (o *PaymentDto) GetPayerPictureOk() (*string, bool)`

GetPayerPictureOk returns a tuple with the PayerPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerPicture

`func (o *PaymentDto) SetPayerPicture(v string)`

SetPayerPicture sets PayerPicture field to given value.

### HasPayerPicture

`func (o *PaymentDto) HasPayerPicture() bool`

HasPayerPicture returns a boolean if a field has been set.

### SetPayerPictureNil

`func (o *PaymentDto) SetPayerPictureNil(b bool)`

 SetPayerPictureNil sets the value for PayerPicture to be an explicit nil

### UnsetPayerPicture
`func (o *PaymentDto) UnsetPayerPicture()`

UnsetPayerPicture ensures that no value is present for PayerPicture, not even an explicit nil
### GetIdentificationPictureTimestamp

`func (o *PaymentDto) GetIdentificationPictureTimestamp() string`

GetIdentificationPictureTimestamp returns the IdentificationPictureTimestamp field if non-nil, zero value otherwise.

### GetIdentificationPictureTimestampOk

`func (o *PaymentDto) GetIdentificationPictureTimestampOk() (*string, bool)`

GetIdentificationPictureTimestampOk returns a tuple with the IdentificationPictureTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationPictureTimestamp

`func (o *PaymentDto) SetIdentificationPictureTimestamp(v string)`

SetIdentificationPictureTimestamp sets IdentificationPictureTimestamp field to given value.

### HasIdentificationPictureTimestamp

`func (o *PaymentDto) HasIdentificationPictureTimestamp() bool`

HasIdentificationPictureTimestamp returns a boolean if a field has been set.

### SetIdentificationPictureTimestampNil

`func (o *PaymentDto) SetIdentificationPictureTimestampNil(b bool)`

 SetIdentificationPictureTimestampNil sets the value for IdentificationPictureTimestamp to be an explicit nil

### UnsetIdentificationPictureTimestamp
`func (o *PaymentDto) UnsetIdentificationPictureTimestamp()`

UnsetIdentificationPictureTimestamp ensures that no value is present for IdentificationPictureTimestamp, not even an explicit nil
### GetIdentificationPicture

`func (o *PaymentDto) GetIdentificationPicture() string`

GetIdentificationPicture returns the IdentificationPicture field if non-nil, zero value otherwise.

### GetIdentificationPictureOk

`func (o *PaymentDto) GetIdentificationPictureOk() (*string, bool)`

GetIdentificationPictureOk returns a tuple with the IdentificationPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationPicture

`func (o *PaymentDto) SetIdentificationPicture(v string)`

SetIdentificationPicture sets IdentificationPicture field to given value.

### HasIdentificationPicture

`func (o *PaymentDto) HasIdentificationPicture() bool`

HasIdentificationPicture returns a boolean if a field has been set.

### SetIdentificationPictureNil

`func (o *PaymentDto) SetIdentificationPictureNil(b bool)`

 SetIdentificationPictureNil sets the value for IdentificationPicture to be an explicit nil

### UnsetIdentificationPicture
`func (o *PaymentDto) UnsetIdentificationPicture()`

UnsetIdentificationPicture ensures that no value is present for IdentificationPicture, not even an explicit nil
### GetIdentificationBackPicture

`func (o *PaymentDto) GetIdentificationBackPicture() string`

GetIdentificationBackPicture returns the IdentificationBackPicture field if non-nil, zero value otherwise.

### GetIdentificationBackPictureOk

`func (o *PaymentDto) GetIdentificationBackPictureOk() (*string, bool)`

GetIdentificationBackPictureOk returns a tuple with the IdentificationBackPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationBackPicture

`func (o *PaymentDto) SetIdentificationBackPicture(v string)`

SetIdentificationBackPicture sets IdentificationBackPicture field to given value.

### HasIdentificationBackPicture

`func (o *PaymentDto) HasIdentificationBackPicture() bool`

HasIdentificationBackPicture returns a boolean if a field has been set.

### SetIdentificationBackPictureNil

`func (o *PaymentDto) SetIdentificationBackPictureNil(b bool)`

 SetIdentificationBackPictureNil sets the value for IdentificationBackPicture to be an explicit nil

### UnsetIdentificationBackPicture
`func (o *PaymentDto) UnsetIdentificationBackPicture()`

UnsetIdentificationBackPicture ensures that no value is present for IdentificationBackPicture, not even an explicit nil
### GetIdentificationBackPictureTimestamp

`func (o *PaymentDto) GetIdentificationBackPictureTimestamp() string`

GetIdentificationBackPictureTimestamp returns the IdentificationBackPictureTimestamp field if non-nil, zero value otherwise.

### GetIdentificationBackPictureTimestampOk

`func (o *PaymentDto) GetIdentificationBackPictureTimestampOk() (*string, bool)`

GetIdentificationBackPictureTimestampOk returns a tuple with the IdentificationBackPictureTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationBackPictureTimestamp

`func (o *PaymentDto) SetIdentificationBackPictureTimestamp(v string)`

SetIdentificationBackPictureTimestamp sets IdentificationBackPictureTimestamp field to given value.

### HasIdentificationBackPictureTimestamp

`func (o *PaymentDto) HasIdentificationBackPictureTimestamp() bool`

HasIdentificationBackPictureTimestamp returns a boolean if a field has been set.

### SetIdentificationBackPictureTimestampNil

`func (o *PaymentDto) SetIdentificationBackPictureTimestampNil(b bool)`

 SetIdentificationBackPictureTimestampNil sets the value for IdentificationBackPictureTimestamp to be an explicit nil

### UnsetIdentificationBackPictureTimestamp
`func (o *PaymentDto) UnsetIdentificationBackPictureTimestamp()`

UnsetIdentificationBackPictureTimestamp ensures that no value is present for IdentificationBackPictureTimestamp, not even an explicit nil
### GetIpLookupId

`func (o *PaymentDto) GetIpLookupId() string`

GetIpLookupId returns the IpLookupId field if non-nil, zero value otherwise.

### GetIpLookupIdOk

`func (o *PaymentDto) GetIpLookupIdOk() (*string, bool)`

GetIpLookupIdOk returns a tuple with the IpLookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLookupId

`func (o *PaymentDto) SetIpLookupId(v string)`

SetIpLookupId sets IpLookupId field to given value.

### HasIpLookupId

`func (o *PaymentDto) HasIpLookupId() bool`

HasIpLookupId returns a boolean if a field has been set.

### SetIpLookupIdNil

`func (o *PaymentDto) SetIpLookupIdNil(b bool)`

 SetIpLookupIdNil sets the value for IpLookupId to be an explicit nil

### UnsetIpLookupId
`func (o *PaymentDto) UnsetIpLookupId()`

UnsetIpLookupId ensures that no value is present for IpLookupId, not even an explicit nil
### GetOrderId

`func (o *PaymentDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *PaymentDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *PaymentDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetAccountingEntryId

`func (o *PaymentDto) GetAccountingEntryId() string`

GetAccountingEntryId returns the AccountingEntryId field if non-nil, zero value otherwise.

### GetAccountingEntryIdOk

`func (o *PaymentDto) GetAccountingEntryIdOk() (*string, bool)`

GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryId

`func (o *PaymentDto) SetAccountingEntryId(v string)`

SetAccountingEntryId sets AccountingEntryId field to given value.

### HasAccountingEntryId

`func (o *PaymentDto) HasAccountingEntryId() bool`

HasAccountingEntryId returns a boolean if a field has been set.

### SetAccountingEntryIdNil

`func (o *PaymentDto) SetAccountingEntryIdNil(b bool)`

 SetAccountingEntryIdNil sets the value for AccountingEntryId to be an explicit nil

### UnsetAccountingEntryId
`func (o *PaymentDto) UnsetAccountingEntryId()`

UnsetAccountingEntryId ensures that no value is present for AccountingEntryId, not even an explicit nil
### GetPaymentGatewayId

`func (o *PaymentDto) GetPaymentGatewayId() string`

GetPaymentGatewayId returns the PaymentGatewayId field if non-nil, zero value otherwise.

### GetPaymentGatewayIdOk

`func (o *PaymentDto) GetPaymentGatewayIdOk() (*string, bool)`

GetPaymentGatewayIdOk returns a tuple with the PaymentGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayId

`func (o *PaymentDto) SetPaymentGatewayId(v string)`

SetPaymentGatewayId sets PaymentGatewayId field to given value.

### HasPaymentGatewayId

`func (o *PaymentDto) HasPaymentGatewayId() bool`

HasPaymentGatewayId returns a boolean if a field has been set.

### SetPaymentGatewayIdNil

`func (o *PaymentDto) SetPaymentGatewayIdNil(b bool)`

 SetPaymentGatewayIdNil sets the value for PaymentGatewayId to be an explicit nil

### UnsetPaymentGatewayId
`func (o *PaymentDto) UnsetPaymentGatewayId()`

UnsetPaymentGatewayId ensures that no value is present for PaymentGatewayId, not even an explicit nil
### GetBankAccountId

`func (o *PaymentDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *PaymentDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *PaymentDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *PaymentDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *PaymentDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *PaymentDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetEnrolmentId

`func (o *PaymentDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *PaymentDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *PaymentDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *PaymentDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *PaymentDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *PaymentDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetBankId

`func (o *PaymentDto) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *PaymentDto) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *PaymentDto) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *PaymentDto) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### SetBankIdNil

`func (o *PaymentDto) SetBankIdNil(b bool)`

 SetBankIdNil sets the value for BankId to be an explicit nil

### UnsetBankId
`func (o *PaymentDto) UnsetBankId()`

UnsetBankId ensures that no value is present for BankId, not even an explicit nil
### GetPaymentTokenId

`func (o *PaymentDto) GetPaymentTokenId() string`

GetPaymentTokenId returns the PaymentTokenId field if non-nil, zero value otherwise.

### GetPaymentTokenIdOk

`func (o *PaymentDto) GetPaymentTokenIdOk() (*string, bool)`

GetPaymentTokenIdOk returns a tuple with the PaymentTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTokenId

`func (o *PaymentDto) SetPaymentTokenId(v string)`

SetPaymentTokenId sets PaymentTokenId field to given value.

### HasPaymentTokenId

`func (o *PaymentDto) HasPaymentTokenId() bool`

HasPaymentTokenId returns a boolean if a field has been set.

### SetPaymentTokenIdNil

`func (o *PaymentDto) SetPaymentTokenIdNil(b bool)`

 SetPaymentTokenIdNil sets the value for PaymentTokenId to be an explicit nil

### UnsetPaymentTokenId
`func (o *PaymentDto) UnsetPaymentTokenId()`

UnsetPaymentTokenId ensures that no value is present for PaymentTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


