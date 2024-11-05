# PaymentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**EmisorWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**ReceiverWalletAccountId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentUpdateDto

`func NewPaymentUpdateDto() *PaymentUpdateDto`

NewPaymentUpdateDto instantiates a new PaymentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentUpdateDtoWithDefaults

`func NewPaymentUpdateDtoWithDefaults() *PaymentUpdateDto`

NewPaymentUpdateDtoWithDefaults instantiates a new PaymentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *PaymentUpdateDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentUpdateDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentUpdateDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PaymentUpdateDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *PaymentUpdateDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *PaymentUpdateDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetTenantId

`func (o *PaymentUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PaymentUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PaymentUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PaymentUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEmisorWalletId

`func (o *PaymentUpdateDto) GetEmisorWalletId() string`

GetEmisorWalletId returns the EmisorWalletId field if non-nil, zero value otherwise.

### GetEmisorWalletIdOk

`func (o *PaymentUpdateDto) GetEmisorWalletIdOk() (*string, bool)`

GetEmisorWalletIdOk returns a tuple with the EmisorWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletId

`func (o *PaymentUpdateDto) SetEmisorWalletId(v string)`

SetEmisorWalletId sets EmisorWalletId field to given value.

### HasEmisorWalletId

`func (o *PaymentUpdateDto) HasEmisorWalletId() bool`

HasEmisorWalletId returns a boolean if a field has been set.

### SetEmisorWalletIdNil

`func (o *PaymentUpdateDto) SetEmisorWalletIdNil(b bool)`

 SetEmisorWalletIdNil sets the value for EmisorWalletId to be an explicit nil

### UnsetEmisorWalletId
`func (o *PaymentUpdateDto) UnsetEmisorWalletId()`

UnsetEmisorWalletId ensures that no value is present for EmisorWalletId, not even an explicit nil
### GetReceiverWalletId

`func (o *PaymentUpdateDto) GetReceiverWalletId() string`

GetReceiverWalletId returns the ReceiverWalletId field if non-nil, zero value otherwise.

### GetReceiverWalletIdOk

`func (o *PaymentUpdateDto) GetReceiverWalletIdOk() (*string, bool)`

GetReceiverWalletIdOk returns a tuple with the ReceiverWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletId

`func (o *PaymentUpdateDto) SetReceiverWalletId(v string)`

SetReceiverWalletId sets ReceiverWalletId field to given value.

### HasReceiverWalletId

`func (o *PaymentUpdateDto) HasReceiverWalletId() bool`

HasReceiverWalletId returns a boolean if a field has been set.

### SetReceiverWalletIdNil

`func (o *PaymentUpdateDto) SetReceiverWalletIdNil(b bool)`

 SetReceiverWalletIdNil sets the value for ReceiverWalletId to be an explicit nil

### UnsetReceiverWalletId
`func (o *PaymentUpdateDto) UnsetReceiverWalletId()`

UnsetReceiverWalletId ensures that no value is present for ReceiverWalletId, not even an explicit nil
### GetCurrencyId

`func (o *PaymentUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PaymentUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PaymentUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PaymentUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *PaymentUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *PaymentUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetForexRate

`func (o *PaymentUpdateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *PaymentUpdateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *PaymentUpdateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *PaymentUpdateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTotalCost

`func (o *PaymentUpdateDto) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *PaymentUpdateDto) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *PaymentUpdateDto) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *PaymentUpdateDto) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *PaymentUpdateDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *PaymentUpdateDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *PaymentUpdateDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *PaymentUpdateDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetClosed

`func (o *PaymentUpdateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *PaymentUpdateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *PaymentUpdateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *PaymentUpdateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetData

`func (o *PaymentUpdateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentUpdateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentUpdateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentUpdateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PaymentUpdateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PaymentUpdateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *PaymentUpdateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *PaymentUpdateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *PaymentUpdateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *PaymentUpdateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *PaymentUpdateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *PaymentUpdateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *PaymentUpdateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *PaymentUpdateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *PaymentUpdateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *PaymentUpdateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *PaymentUpdateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *PaymentUpdateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *PaymentUpdateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *PaymentUpdateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *PaymentUpdateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *PaymentUpdateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *PaymentUpdateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *PaymentUpdateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetResponse

`func (o *PaymentUpdateDto) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *PaymentUpdateDto) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *PaymentUpdateDto) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *PaymentUpdateDto) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *PaymentUpdateDto) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *PaymentUpdateDto) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetAuthorization

`func (o *PaymentUpdateDto) GetAuthorization() string`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *PaymentUpdateDto) GetAuthorizationOk() (*string, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *PaymentUpdateDto) SetAuthorization(v string)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *PaymentUpdateDto) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### SetAuthorizationNil

`func (o *PaymentUpdateDto) SetAuthorizationNil(b bool)`

 SetAuthorizationNil sets the value for Authorization to be an explicit nil

### UnsetAuthorization
`func (o *PaymentUpdateDto) UnsetAuthorization()`

UnsetAuthorization ensures that no value is present for Authorization, not even an explicit nil
### GetReferenceCode

`func (o *PaymentUpdateDto) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *PaymentUpdateDto) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *PaymentUpdateDto) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *PaymentUpdateDto) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.

### SetReferenceCodeNil

`func (o *PaymentUpdateDto) SetReferenceCodeNil(b bool)`

 SetReferenceCodeNil sets the value for ReferenceCode to be an explicit nil

### UnsetReferenceCode
`func (o *PaymentUpdateDto) UnsetReferenceCode()`

UnsetReferenceCode ensures that no value is present for ReferenceCode, not even an explicit nil
### GetCorrelationCode

`func (o *PaymentUpdateDto) GetCorrelationCode() string`

GetCorrelationCode returns the CorrelationCode field if non-nil, zero value otherwise.

### GetCorrelationCodeOk

`func (o *PaymentUpdateDto) GetCorrelationCodeOk() (*string, bool)`

GetCorrelationCodeOk returns a tuple with the CorrelationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationCode

`func (o *PaymentUpdateDto) SetCorrelationCode(v string)`

SetCorrelationCode sets CorrelationCode field to given value.

### HasCorrelationCode

`func (o *PaymentUpdateDto) HasCorrelationCode() bool`

HasCorrelationCode returns a boolean if a field has been set.

### SetCorrelationCodeNil

`func (o *PaymentUpdateDto) SetCorrelationCodeNil(b bool)`

 SetCorrelationCodeNil sets the value for CorrelationCode to be an explicit nil

### UnsetCorrelationCode
`func (o *PaymentUpdateDto) UnsetCorrelationCode()`

UnsetCorrelationCode ensures that no value is present for CorrelationCode, not even an explicit nil
### GetLastUpdated

`func (o *PaymentUpdateDto) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PaymentUpdateDto) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PaymentUpdateDto) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PaymentUpdateDto) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOnBehalfOf

`func (o *PaymentUpdateDto) GetOnBehalfOf() int32`

GetOnBehalfOf returns the OnBehalfOf field if non-nil, zero value otherwise.

### GetOnBehalfOfOk

`func (o *PaymentUpdateDto) GetOnBehalfOfOk() (*int32, bool)`

GetOnBehalfOfOk returns a tuple with the OnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOf

`func (o *PaymentUpdateDto) SetOnBehalfOf(v int32)`

SetOnBehalfOf sets OnBehalfOf field to given value.

### HasOnBehalfOf

`func (o *PaymentUpdateDto) HasOnBehalfOf() bool`

HasOnBehalfOf returns a boolean if a field has been set.

### GetPaymentType

`func (o *PaymentUpdateDto) GetPaymentType() int32`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentUpdateDto) GetPaymentTypeOk() (*int32, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentUpdateDto) SetPaymentType(v int32)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PaymentUpdateDto) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *PaymentUpdateDto) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PaymentUpdateDto) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PaymentUpdateDto) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *PaymentUpdateDto) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetBaseCost

`func (o *PaymentUpdateDto) GetBaseCost() float64`

GetBaseCost returns the BaseCost field if non-nil, zero value otherwise.

### GetBaseCostOk

`func (o *PaymentUpdateDto) GetBaseCostOk() (*float64, bool)`

GetBaseCostOk returns a tuple with the BaseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCost

`func (o *PaymentUpdateDto) SetBaseCost(v float64)`

SetBaseCost sets BaseCost field to given value.

### HasBaseCost

`func (o *PaymentUpdateDto) HasBaseCost() bool`

HasBaseCost returns a boolean if a field has been set.

### GetSignature

`func (o *PaymentUpdateDto) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentUpdateDto) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentUpdateDto) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *PaymentUpdateDto) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *PaymentUpdateDto) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *PaymentUpdateDto) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetSignatureMismatch

`func (o *PaymentUpdateDto) GetSignatureMismatch() bool`

GetSignatureMismatch returns the SignatureMismatch field if non-nil, zero value otherwise.

### GetSignatureMismatchOk

`func (o *PaymentUpdateDto) GetSignatureMismatchOk() (*bool, bool)`

GetSignatureMismatchOk returns a tuple with the SignatureMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureMismatch

`func (o *PaymentUpdateDto) SetSignatureMismatch(v bool)`

SetSignatureMismatch sets SignatureMismatch field to given value.

### HasSignatureMismatch

`func (o *PaymentUpdateDto) HasSignatureMismatch() bool`

HasSignatureMismatch returns a boolean if a field has been set.

### GetIsExternal

`func (o *PaymentUpdateDto) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *PaymentUpdateDto) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *PaymentUpdateDto) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *PaymentUpdateDto) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetMarkedForRevision

`func (o *PaymentUpdateDto) GetMarkedForRevision() bool`

GetMarkedForRevision returns the MarkedForRevision field if non-nil, zero value otherwise.

### GetMarkedForRevisionOk

`func (o *PaymentUpdateDto) GetMarkedForRevisionOk() (*bool, bool)`

GetMarkedForRevisionOk returns a tuple with the MarkedForRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForRevision

`func (o *PaymentUpdateDto) SetMarkedForRevision(v bool)`

SetMarkedForRevision sets MarkedForRevision field to given value.

### HasMarkedForRevision

`func (o *PaymentUpdateDto) HasMarkedForRevision() bool`

HasMarkedForRevision returns a boolean if a field has been set.

### GetForexRatesSnapshot

`func (o *PaymentUpdateDto) GetForexRatesSnapshot() string`

GetForexRatesSnapshot returns the ForexRatesSnapshot field if non-nil, zero value otherwise.

### GetForexRatesSnapshotOk

`func (o *PaymentUpdateDto) GetForexRatesSnapshotOk() (*string, bool)`

GetForexRatesSnapshotOk returns a tuple with the ForexRatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRatesSnapshot

`func (o *PaymentUpdateDto) SetForexRatesSnapshot(v string)`

SetForexRatesSnapshot sets ForexRatesSnapshot field to given value.

### HasForexRatesSnapshot

`func (o *PaymentUpdateDto) HasForexRatesSnapshot() bool`

HasForexRatesSnapshot returns a boolean if a field has been set.

### SetForexRatesSnapshotNil

`func (o *PaymentUpdateDto) SetForexRatesSnapshotNil(b bool)`

 SetForexRatesSnapshotNil sets the value for ForexRatesSnapshot to be an explicit nil

### UnsetForexRatesSnapshot
`func (o *PaymentUpdateDto) UnsetForexRatesSnapshot()`

UnsetForexRatesSnapshot ensures that no value is present for ForexRatesSnapshot, not even an explicit nil
### GetOfficialId

`func (o *PaymentUpdateDto) GetOfficialId() string`

GetOfficialId returns the OfficialId field if non-nil, zero value otherwise.

### GetOfficialIdOk

`func (o *PaymentUpdateDto) GetOfficialIdOk() (*string, bool)`

GetOfficialIdOk returns a tuple with the OfficialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialId

`func (o *PaymentUpdateDto) SetOfficialId(v string)`

SetOfficialId sets OfficialId field to given value.

### HasOfficialId

`func (o *PaymentUpdateDto) HasOfficialId() bool`

HasOfficialId returns a boolean if a field has been set.

### SetOfficialIdNil

`func (o *PaymentUpdateDto) SetOfficialIdNil(b bool)`

 SetOfficialIdNil sets the value for OfficialId to be an explicit nil

### UnsetOfficialId
`func (o *PaymentUpdateDto) UnsetOfficialId()`

UnsetOfficialId ensures that no value is present for OfficialId, not even an explicit nil
### GetOfficialIdExpeditionDate

`func (o *PaymentUpdateDto) GetOfficialIdExpeditionDate() time.Time`

GetOfficialIdExpeditionDate returns the OfficialIdExpeditionDate field if non-nil, zero value otherwise.

### GetOfficialIdExpeditionDateOk

`func (o *PaymentUpdateDto) GetOfficialIdExpeditionDateOk() (*time.Time, bool)`

GetOfficialIdExpeditionDateOk returns a tuple with the OfficialIdExpeditionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialIdExpeditionDate

`func (o *PaymentUpdateDto) SetOfficialIdExpeditionDate(v time.Time)`

SetOfficialIdExpeditionDate sets OfficialIdExpeditionDate field to given value.

### HasOfficialIdExpeditionDate

`func (o *PaymentUpdateDto) HasOfficialIdExpeditionDate() bool`

HasOfficialIdExpeditionDate returns a boolean if a field has been set.

### GetFiscalIdentificationTypeId

`func (o *PaymentUpdateDto) GetFiscalIdentificationTypeId() string`

GetFiscalIdentificationTypeId returns the FiscalIdentificationTypeId field if non-nil, zero value otherwise.

### GetFiscalIdentificationTypeIdOk

`func (o *PaymentUpdateDto) GetFiscalIdentificationTypeIdOk() (*string, bool)`

GetFiscalIdentificationTypeIdOk returns a tuple with the FiscalIdentificationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalIdentificationTypeId

`func (o *PaymentUpdateDto) SetFiscalIdentificationTypeId(v string)`

SetFiscalIdentificationTypeId sets FiscalIdentificationTypeId field to given value.

### HasFiscalIdentificationTypeId

`func (o *PaymentUpdateDto) HasFiscalIdentificationTypeId() bool`

HasFiscalIdentificationTypeId returns a boolean if a field has been set.

### SetFiscalIdentificationTypeIdNil

`func (o *PaymentUpdateDto) SetFiscalIdentificationTypeIdNil(b bool)`

 SetFiscalIdentificationTypeIdNil sets the value for FiscalIdentificationTypeId to be an explicit nil

### UnsetFiscalIdentificationTypeId
`func (o *PaymentUpdateDto) UnsetFiscalIdentificationTypeId()`

UnsetFiscalIdentificationTypeId ensures that no value is present for FiscalIdentificationTypeId, not even an explicit nil
### GetBillingAddress

`func (o *PaymentUpdateDto) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentUpdateDto) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentUpdateDto) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *PaymentUpdateDto) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *PaymentUpdateDto) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *PaymentUpdateDto) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetPhone

`func (o *PaymentUpdateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PaymentUpdateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PaymentUpdateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PaymentUpdateDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PaymentUpdateDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PaymentUpdateDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCellphone

`func (o *PaymentUpdateDto) GetCellphone() string`

GetCellphone returns the Cellphone field if non-nil, zero value otherwise.

### GetCellphoneOk

`func (o *PaymentUpdateDto) GetCellphoneOk() (*string, bool)`

GetCellphoneOk returns a tuple with the Cellphone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphone

`func (o *PaymentUpdateDto) SetCellphone(v string)`

SetCellphone sets Cellphone field to given value.

### HasCellphone

`func (o *PaymentUpdateDto) HasCellphone() bool`

HasCellphone returns a boolean if a field has been set.

### SetCellphoneNil

`func (o *PaymentUpdateDto) SetCellphoneNil(b bool)`

 SetCellphoneNil sets the value for Cellphone to be an explicit nil

### UnsetCellphone
`func (o *PaymentUpdateDto) UnsetCellphone()`

UnsetCellphone ensures that no value is present for Cellphone, not even an explicit nil
### GetDepartment

`func (o *PaymentUpdateDto) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *PaymentUpdateDto) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *PaymentUpdateDto) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *PaymentUpdateDto) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *PaymentUpdateDto) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *PaymentUpdateDto) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetCity

`func (o *PaymentUpdateDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PaymentUpdateDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PaymentUpdateDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PaymentUpdateDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *PaymentUpdateDto) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *PaymentUpdateDto) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountryId

`func (o *PaymentUpdateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *PaymentUpdateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *PaymentUpdateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *PaymentUpdateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *PaymentUpdateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *PaymentUpdateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetLocationId

`func (o *PaymentUpdateDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *PaymentUpdateDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *PaymentUpdateDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *PaymentUpdateDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *PaymentUpdateDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *PaymentUpdateDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetEntitlementId

`func (o *PaymentUpdateDto) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *PaymentUpdateDto) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *PaymentUpdateDto) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.

### HasEntitlementId

`func (o *PaymentUpdateDto) HasEntitlementId() bool`

HasEntitlementId returns a boolean if a field has been set.

### SetEntitlementIdNil

`func (o *PaymentUpdateDto) SetEntitlementIdNil(b bool)`

 SetEntitlementIdNil sets the value for EntitlementId to be an explicit nil

### UnsetEntitlementId
`func (o *PaymentUpdateDto) UnsetEntitlementId()`

UnsetEntitlementId ensures that no value is present for EntitlementId, not even an explicit nil
### GetAntiFraudScore

`func (o *PaymentUpdateDto) GetAntiFraudScore() float64`

GetAntiFraudScore returns the AntiFraudScore field if non-nil, zero value otherwise.

### GetAntiFraudScoreOk

`func (o *PaymentUpdateDto) GetAntiFraudScoreOk() (*float64, bool)`

GetAntiFraudScoreOk returns a tuple with the AntiFraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudScore

`func (o *PaymentUpdateDto) SetAntiFraudScore(v float64)`

SetAntiFraudScore sets AntiFraudScore field to given value.

### HasAntiFraudScore

`func (o *PaymentUpdateDto) HasAntiFraudScore() bool`

HasAntiFraudScore returns a boolean if a field has been set.

### GetCallRecordURL

`func (o *PaymentUpdateDto) GetCallRecordURL() string`

GetCallRecordURL returns the CallRecordURL field if non-nil, zero value otherwise.

### GetCallRecordURLOk

`func (o *PaymentUpdateDto) GetCallRecordURLOk() (*string, bool)`

GetCallRecordURLOk returns a tuple with the CallRecordURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordURL

`func (o *PaymentUpdateDto) SetCallRecordURL(v string)`

SetCallRecordURL sets CallRecordURL field to given value.

### HasCallRecordURL

`func (o *PaymentUpdateDto) HasCallRecordURL() bool`

HasCallRecordURL returns a boolean if a field has been set.

### SetCallRecordURLNil

`func (o *PaymentUpdateDto) SetCallRecordURLNil(b bool)`

 SetCallRecordURLNil sets the value for CallRecordURL to be an explicit nil

### UnsetCallRecordURL
`func (o *PaymentUpdateDto) UnsetCallRecordURL()`

UnsetCallRecordURL ensures that no value is present for CallRecordURL, not even an explicit nil
### GetCalled

`func (o *PaymentUpdateDto) GetCalled() bool`

GetCalled returns the Called field if non-nil, zero value otherwise.

### GetCalledOk

`func (o *PaymentUpdateDto) GetCalledOk() (*bool, bool)`

GetCalledOk returns a tuple with the Called field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalled

`func (o *PaymentUpdateDto) SetCalled(v bool)`

SetCalled sets Called field to given value.

### HasCalled

`func (o *PaymentUpdateDto) HasCalled() bool`

HasCalled returns a boolean if a field has been set.

### GetVerified

`func (o *PaymentUpdateDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PaymentUpdateDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PaymentUpdateDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *PaymentUpdateDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetPayerPictureTimestamp

`func (o *PaymentUpdateDto) GetPayerPictureTimestamp() string`

GetPayerPictureTimestamp returns the PayerPictureTimestamp field if non-nil, zero value otherwise.

### GetPayerPictureTimestampOk

`func (o *PaymentUpdateDto) GetPayerPictureTimestampOk() (*string, bool)`

GetPayerPictureTimestampOk returns a tuple with the PayerPictureTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerPictureTimestamp

`func (o *PaymentUpdateDto) SetPayerPictureTimestamp(v string)`

SetPayerPictureTimestamp sets PayerPictureTimestamp field to given value.

### HasPayerPictureTimestamp

`func (o *PaymentUpdateDto) HasPayerPictureTimestamp() bool`

HasPayerPictureTimestamp returns a boolean if a field has been set.

### SetPayerPictureTimestampNil

`func (o *PaymentUpdateDto) SetPayerPictureTimestampNil(b bool)`

 SetPayerPictureTimestampNil sets the value for PayerPictureTimestamp to be an explicit nil

### UnsetPayerPictureTimestamp
`func (o *PaymentUpdateDto) UnsetPayerPictureTimestamp()`

UnsetPayerPictureTimestamp ensures that no value is present for PayerPictureTimestamp, not even an explicit nil
### GetPayerPicture

`func (o *PaymentUpdateDto) GetPayerPicture() string`

GetPayerPicture returns the PayerPicture field if non-nil, zero value otherwise.

### GetPayerPictureOk

`func (o *PaymentUpdateDto) GetPayerPictureOk() (*string, bool)`

GetPayerPictureOk returns a tuple with the PayerPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerPicture

`func (o *PaymentUpdateDto) SetPayerPicture(v string)`

SetPayerPicture sets PayerPicture field to given value.

### HasPayerPicture

`func (o *PaymentUpdateDto) HasPayerPicture() bool`

HasPayerPicture returns a boolean if a field has been set.

### SetPayerPictureNil

`func (o *PaymentUpdateDto) SetPayerPictureNil(b bool)`

 SetPayerPictureNil sets the value for PayerPicture to be an explicit nil

### UnsetPayerPicture
`func (o *PaymentUpdateDto) UnsetPayerPicture()`

UnsetPayerPicture ensures that no value is present for PayerPicture, not even an explicit nil
### GetIdentificationPictureTimestamp

`func (o *PaymentUpdateDto) GetIdentificationPictureTimestamp() string`

GetIdentificationPictureTimestamp returns the IdentificationPictureTimestamp field if non-nil, zero value otherwise.

### GetIdentificationPictureTimestampOk

`func (o *PaymentUpdateDto) GetIdentificationPictureTimestampOk() (*string, bool)`

GetIdentificationPictureTimestampOk returns a tuple with the IdentificationPictureTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationPictureTimestamp

`func (o *PaymentUpdateDto) SetIdentificationPictureTimestamp(v string)`

SetIdentificationPictureTimestamp sets IdentificationPictureTimestamp field to given value.

### HasIdentificationPictureTimestamp

`func (o *PaymentUpdateDto) HasIdentificationPictureTimestamp() bool`

HasIdentificationPictureTimestamp returns a boolean if a field has been set.

### SetIdentificationPictureTimestampNil

`func (o *PaymentUpdateDto) SetIdentificationPictureTimestampNil(b bool)`

 SetIdentificationPictureTimestampNil sets the value for IdentificationPictureTimestamp to be an explicit nil

### UnsetIdentificationPictureTimestamp
`func (o *PaymentUpdateDto) UnsetIdentificationPictureTimestamp()`

UnsetIdentificationPictureTimestamp ensures that no value is present for IdentificationPictureTimestamp, not even an explicit nil
### GetIdentificationPicture

`func (o *PaymentUpdateDto) GetIdentificationPicture() string`

GetIdentificationPicture returns the IdentificationPicture field if non-nil, zero value otherwise.

### GetIdentificationPictureOk

`func (o *PaymentUpdateDto) GetIdentificationPictureOk() (*string, bool)`

GetIdentificationPictureOk returns a tuple with the IdentificationPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationPicture

`func (o *PaymentUpdateDto) SetIdentificationPicture(v string)`

SetIdentificationPicture sets IdentificationPicture field to given value.

### HasIdentificationPicture

`func (o *PaymentUpdateDto) HasIdentificationPicture() bool`

HasIdentificationPicture returns a boolean if a field has been set.

### SetIdentificationPictureNil

`func (o *PaymentUpdateDto) SetIdentificationPictureNil(b bool)`

 SetIdentificationPictureNil sets the value for IdentificationPicture to be an explicit nil

### UnsetIdentificationPicture
`func (o *PaymentUpdateDto) UnsetIdentificationPicture()`

UnsetIdentificationPicture ensures that no value is present for IdentificationPicture, not even an explicit nil
### GetIdentificationBackPicture

`func (o *PaymentUpdateDto) GetIdentificationBackPicture() string`

GetIdentificationBackPicture returns the IdentificationBackPicture field if non-nil, zero value otherwise.

### GetIdentificationBackPictureOk

`func (o *PaymentUpdateDto) GetIdentificationBackPictureOk() (*string, bool)`

GetIdentificationBackPictureOk returns a tuple with the IdentificationBackPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationBackPicture

`func (o *PaymentUpdateDto) SetIdentificationBackPicture(v string)`

SetIdentificationBackPicture sets IdentificationBackPicture field to given value.

### HasIdentificationBackPicture

`func (o *PaymentUpdateDto) HasIdentificationBackPicture() bool`

HasIdentificationBackPicture returns a boolean if a field has been set.

### SetIdentificationBackPictureNil

`func (o *PaymentUpdateDto) SetIdentificationBackPictureNil(b bool)`

 SetIdentificationBackPictureNil sets the value for IdentificationBackPicture to be an explicit nil

### UnsetIdentificationBackPicture
`func (o *PaymentUpdateDto) UnsetIdentificationBackPicture()`

UnsetIdentificationBackPicture ensures that no value is present for IdentificationBackPicture, not even an explicit nil
### GetIdentificationBackPictureTimestamp

`func (o *PaymentUpdateDto) GetIdentificationBackPictureTimestamp() string`

GetIdentificationBackPictureTimestamp returns the IdentificationBackPictureTimestamp field if non-nil, zero value otherwise.

### GetIdentificationBackPictureTimestampOk

`func (o *PaymentUpdateDto) GetIdentificationBackPictureTimestampOk() (*string, bool)`

GetIdentificationBackPictureTimestampOk returns a tuple with the IdentificationBackPictureTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationBackPictureTimestamp

`func (o *PaymentUpdateDto) SetIdentificationBackPictureTimestamp(v string)`

SetIdentificationBackPictureTimestamp sets IdentificationBackPictureTimestamp field to given value.

### HasIdentificationBackPictureTimestamp

`func (o *PaymentUpdateDto) HasIdentificationBackPictureTimestamp() bool`

HasIdentificationBackPictureTimestamp returns a boolean if a field has been set.

### SetIdentificationBackPictureTimestampNil

`func (o *PaymentUpdateDto) SetIdentificationBackPictureTimestampNil(b bool)`

 SetIdentificationBackPictureTimestampNil sets the value for IdentificationBackPictureTimestamp to be an explicit nil

### UnsetIdentificationBackPictureTimestamp
`func (o *PaymentUpdateDto) UnsetIdentificationBackPictureTimestamp()`

UnsetIdentificationBackPictureTimestamp ensures that no value is present for IdentificationBackPictureTimestamp, not even an explicit nil
### GetIpLookupId

`func (o *PaymentUpdateDto) GetIpLookupId() string`

GetIpLookupId returns the IpLookupId field if non-nil, zero value otherwise.

### GetIpLookupIdOk

`func (o *PaymentUpdateDto) GetIpLookupIdOk() (*string, bool)`

GetIpLookupIdOk returns a tuple with the IpLookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLookupId

`func (o *PaymentUpdateDto) SetIpLookupId(v string)`

SetIpLookupId sets IpLookupId field to given value.

### HasIpLookupId

`func (o *PaymentUpdateDto) HasIpLookupId() bool`

HasIpLookupId returns a boolean if a field has been set.

### SetIpLookupIdNil

`func (o *PaymentUpdateDto) SetIpLookupIdNil(b bool)`

 SetIpLookupIdNil sets the value for IpLookupId to be an explicit nil

### UnsetIpLookupId
`func (o *PaymentUpdateDto) UnsetIpLookupId()`

UnsetIpLookupId ensures that no value is present for IpLookupId, not even an explicit nil
### GetOrderId

`func (o *PaymentUpdateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentUpdateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentUpdateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentUpdateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *PaymentUpdateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *PaymentUpdateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetAccountingEntryId

`func (o *PaymentUpdateDto) GetAccountingEntryId() string`

GetAccountingEntryId returns the AccountingEntryId field if non-nil, zero value otherwise.

### GetAccountingEntryIdOk

`func (o *PaymentUpdateDto) GetAccountingEntryIdOk() (*string, bool)`

GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryId

`func (o *PaymentUpdateDto) SetAccountingEntryId(v string)`

SetAccountingEntryId sets AccountingEntryId field to given value.

### HasAccountingEntryId

`func (o *PaymentUpdateDto) HasAccountingEntryId() bool`

HasAccountingEntryId returns a boolean if a field has been set.

### SetAccountingEntryIdNil

`func (o *PaymentUpdateDto) SetAccountingEntryIdNil(b bool)`

 SetAccountingEntryIdNil sets the value for AccountingEntryId to be an explicit nil

### UnsetAccountingEntryId
`func (o *PaymentUpdateDto) UnsetAccountingEntryId()`

UnsetAccountingEntryId ensures that no value is present for AccountingEntryId, not even an explicit nil
### GetPaymentGatewayId

`func (o *PaymentUpdateDto) GetPaymentGatewayId() string`

GetPaymentGatewayId returns the PaymentGatewayId field if non-nil, zero value otherwise.

### GetPaymentGatewayIdOk

`func (o *PaymentUpdateDto) GetPaymentGatewayIdOk() (*string, bool)`

GetPaymentGatewayIdOk returns a tuple with the PaymentGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayId

`func (o *PaymentUpdateDto) SetPaymentGatewayId(v string)`

SetPaymentGatewayId sets PaymentGatewayId field to given value.

### HasPaymentGatewayId

`func (o *PaymentUpdateDto) HasPaymentGatewayId() bool`

HasPaymentGatewayId returns a boolean if a field has been set.

### SetPaymentGatewayIdNil

`func (o *PaymentUpdateDto) SetPaymentGatewayIdNil(b bool)`

 SetPaymentGatewayIdNil sets the value for PaymentGatewayId to be an explicit nil

### UnsetPaymentGatewayId
`func (o *PaymentUpdateDto) UnsetPaymentGatewayId()`

UnsetPaymentGatewayId ensures that no value is present for PaymentGatewayId, not even an explicit nil
### GetBankAccountId

`func (o *PaymentUpdateDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *PaymentUpdateDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *PaymentUpdateDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *PaymentUpdateDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *PaymentUpdateDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *PaymentUpdateDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetEnrolmentId

`func (o *PaymentUpdateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *PaymentUpdateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *PaymentUpdateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *PaymentUpdateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *PaymentUpdateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *PaymentUpdateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetBankId

`func (o *PaymentUpdateDto) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *PaymentUpdateDto) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *PaymentUpdateDto) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *PaymentUpdateDto) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### SetBankIdNil

`func (o *PaymentUpdateDto) SetBankIdNil(b bool)`

 SetBankIdNil sets the value for BankId to be an explicit nil

### UnsetBankId
`func (o *PaymentUpdateDto) UnsetBankId()`

UnsetBankId ensures that no value is present for BankId, not even an explicit nil
### GetPaymentTokenId

`func (o *PaymentUpdateDto) GetPaymentTokenId() string`

GetPaymentTokenId returns the PaymentTokenId field if non-nil, zero value otherwise.

### GetPaymentTokenIdOk

`func (o *PaymentUpdateDto) GetPaymentTokenIdOk() (*string, bool)`

GetPaymentTokenIdOk returns a tuple with the PaymentTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTokenId

`func (o *PaymentUpdateDto) SetPaymentTokenId(v string)`

SetPaymentTokenId sets PaymentTokenId field to given value.

### HasPaymentTokenId

`func (o *PaymentUpdateDto) HasPaymentTokenId() bool`

HasPaymentTokenId returns a boolean if a field has been set.

### SetPaymentTokenIdNil

`func (o *PaymentUpdateDto) SetPaymentTokenIdNil(b bool)`

 SetPaymentTokenIdNil sets the value for PaymentTokenId to be an explicit nil

### UnsetPaymentTokenId
`func (o *PaymentUpdateDto) UnsetPaymentTokenId()`

UnsetPaymentTokenId ensures that no value is present for PaymentTokenId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *PaymentUpdateDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *PaymentUpdateDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *PaymentUpdateDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *PaymentUpdateDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *PaymentUpdateDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *PaymentUpdateDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *PaymentUpdateDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *PaymentUpdateDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *PaymentUpdateDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *PaymentUpdateDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *PaymentUpdateDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *PaymentUpdateDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


