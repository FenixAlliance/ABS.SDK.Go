# LicenseKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**EntitlementId** | Pointer to **NullableString** |  | [optional] 
**Seats** | Pointer to **int32** |  | [optional] 
**LicenseType** | Pointer to **int32** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Features** | Pointer to [**[]LicenseFeature**](LicenseFeature.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]AdditionalAttribute**](AdditionalAttribute.md) |  | [optional] 

## Methods

### NewLicenseKeyRequest

`func NewLicenseKeyRequest() *LicenseKeyRequest`

NewLicenseKeyRequest instantiates a new LicenseKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseKeyRequestWithDefaults

`func NewLicenseKeyRequestWithDefaults() *LicenseKeyRequest`

NewLicenseKeyRequestWithDefaults instantiates a new LicenseKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *LicenseKeyRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LicenseKeyRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LicenseKeyRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LicenseKeyRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTenantId

`func (o *LicenseKeyRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LicenseKeyRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LicenseKeyRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LicenseKeyRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LicenseKeyRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LicenseKeyRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetOrderId

`func (o *LicenseKeyRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *LicenseKeyRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *LicenseKeyRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *LicenseKeyRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *LicenseKeyRequest) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *LicenseKeyRequest) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetPaymentId

`func (o *LicenseKeyRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *LicenseKeyRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *LicenseKeyRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *LicenseKeyRequest) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *LicenseKeyRequest) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *LicenseKeyRequest) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetInvoiceId

`func (o *LicenseKeyRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *LicenseKeyRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *LicenseKeyRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *LicenseKeyRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *LicenseKeyRequest) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *LicenseKeyRequest) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetEnrollmentId

`func (o *LicenseKeyRequest) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LicenseKeyRequest) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LicenseKeyRequest) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LicenseKeyRequest) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LicenseKeyRequest) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LicenseKeyRequest) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetEntitlementId

`func (o *LicenseKeyRequest) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *LicenseKeyRequest) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *LicenseKeyRequest) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.

### HasEntitlementId

`func (o *LicenseKeyRequest) HasEntitlementId() bool`

HasEntitlementId returns a boolean if a field has been set.

### SetEntitlementIdNil

`func (o *LicenseKeyRequest) SetEntitlementIdNil(b bool)`

 SetEntitlementIdNil sets the value for EntitlementId to be an explicit nil

### UnsetEntitlementId
`func (o *LicenseKeyRequest) UnsetEntitlementId()`

UnsetEntitlementId ensures that no value is present for EntitlementId, not even an explicit nil
### GetSeats

`func (o *LicenseKeyRequest) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *LicenseKeyRequest) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *LicenseKeyRequest) SetSeats(v int32)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *LicenseKeyRequest) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetLicenseType

`func (o *LicenseKeyRequest) GetLicenseType() int32`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicenseKeyRequest) GetLicenseTypeOk() (*int32, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicenseKeyRequest) SetLicenseType(v int32)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicenseKeyRequest) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetExpirationDate

`func (o *LicenseKeyRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *LicenseKeyRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *LicenseKeyRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *LicenseKeyRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetFeatures

`func (o *LicenseKeyRequest) GetFeatures() []LicenseFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LicenseKeyRequest) GetFeaturesOk() (*[]LicenseFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LicenseKeyRequest) SetFeatures(v []LicenseFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LicenseKeyRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *LicenseKeyRequest) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *LicenseKeyRequest) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetAdditionalAttributes

`func (o *LicenseKeyRequest) GetAdditionalAttributes() []AdditionalAttribute`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *LicenseKeyRequest) GetAdditionalAttributesOk() (*[]AdditionalAttribute, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *LicenseKeyRequest) SetAdditionalAttributes(v []AdditionalAttribute)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *LicenseKeyRequest) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### SetAdditionalAttributesNil

`func (o *LicenseKeyRequest) SetAdditionalAttributesNil(b bool)`

 SetAdditionalAttributesNil sets the value for AdditionalAttributes to be an explicit nil

### UnsetAdditionalAttributes
`func (o *LicenseKeyRequest) UnsetAdditionalAttributes()`

UnsetAdditionalAttributes ensures that no value is present for AdditionalAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


