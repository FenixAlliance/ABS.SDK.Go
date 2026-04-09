# InvoiceEnumerationRangeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**CurrentNumeration** | Pointer to **int64** |  | [optional] 
**NumerationTo** | Pointer to **int64** |  | [optional] 
**NumerationFrom** | Pointer to **int64** |  | [optional] 
**ValidFrom** | Pointer to **time.Time** |  | [optional] 
**ValidTo** | Pointer to **time.Time** |  | [optional] 
**FiscalAuthorityId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceEnumerationRangeDto

`func NewInvoiceEnumerationRangeDto() *InvoiceEnumerationRangeDto`

NewInvoiceEnumerationRangeDto instantiates a new InvoiceEnumerationRangeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceEnumerationRangeDtoWithDefaults

`func NewInvoiceEnumerationRangeDtoWithDefaults() *InvoiceEnumerationRangeDto`

NewInvoiceEnumerationRangeDtoWithDefaults instantiates a new InvoiceEnumerationRangeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceEnumerationRangeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceEnumerationRangeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceEnumerationRangeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceEnumerationRangeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InvoiceEnumerationRangeDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceEnumerationRangeDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *InvoiceEnumerationRangeDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceEnumerationRangeDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceEnumerationRangeDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceEnumerationRangeDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InvoiceEnumerationRangeDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InvoiceEnumerationRangeDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetPrefix

`func (o *InvoiceEnumerationRangeDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InvoiceEnumerationRangeDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InvoiceEnumerationRangeDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *InvoiceEnumerationRangeDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *InvoiceEnumerationRangeDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *InvoiceEnumerationRangeDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetSuffix

`func (o *InvoiceEnumerationRangeDto) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *InvoiceEnumerationRangeDto) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *InvoiceEnumerationRangeDto) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *InvoiceEnumerationRangeDto) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *InvoiceEnumerationRangeDto) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *InvoiceEnumerationRangeDto) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIdentifier

`func (o *InvoiceEnumerationRangeDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InvoiceEnumerationRangeDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InvoiceEnumerationRangeDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InvoiceEnumerationRangeDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *InvoiceEnumerationRangeDto) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *InvoiceEnumerationRangeDto) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetQualifiedName

`func (o *InvoiceEnumerationRangeDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *InvoiceEnumerationRangeDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *InvoiceEnumerationRangeDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *InvoiceEnumerationRangeDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *InvoiceEnumerationRangeDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *InvoiceEnumerationRangeDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetCurrentNumeration

`func (o *InvoiceEnumerationRangeDto) GetCurrentNumeration() int64`

GetCurrentNumeration returns the CurrentNumeration field if non-nil, zero value otherwise.

### GetCurrentNumerationOk

`func (o *InvoiceEnumerationRangeDto) GetCurrentNumerationOk() (*int64, bool)`

GetCurrentNumerationOk returns a tuple with the CurrentNumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumeration

`func (o *InvoiceEnumerationRangeDto) SetCurrentNumeration(v int64)`

SetCurrentNumeration sets CurrentNumeration field to given value.

### HasCurrentNumeration

`func (o *InvoiceEnumerationRangeDto) HasCurrentNumeration() bool`

HasCurrentNumeration returns a boolean if a field has been set.

### GetNumerationTo

`func (o *InvoiceEnumerationRangeDto) GetNumerationTo() int64`

GetNumerationTo returns the NumerationTo field if non-nil, zero value otherwise.

### GetNumerationToOk

`func (o *InvoiceEnumerationRangeDto) GetNumerationToOk() (*int64, bool)`

GetNumerationToOk returns a tuple with the NumerationTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerationTo

`func (o *InvoiceEnumerationRangeDto) SetNumerationTo(v int64)`

SetNumerationTo sets NumerationTo field to given value.

### HasNumerationTo

`func (o *InvoiceEnumerationRangeDto) HasNumerationTo() bool`

HasNumerationTo returns a boolean if a field has been set.

### GetNumerationFrom

`func (o *InvoiceEnumerationRangeDto) GetNumerationFrom() int64`

GetNumerationFrom returns the NumerationFrom field if non-nil, zero value otherwise.

### GetNumerationFromOk

`func (o *InvoiceEnumerationRangeDto) GetNumerationFromOk() (*int64, bool)`

GetNumerationFromOk returns a tuple with the NumerationFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerationFrom

`func (o *InvoiceEnumerationRangeDto) SetNumerationFrom(v int64)`

SetNumerationFrom sets NumerationFrom field to given value.

### HasNumerationFrom

`func (o *InvoiceEnumerationRangeDto) HasNumerationFrom() bool`

HasNumerationFrom returns a boolean if a field has been set.

### GetValidFrom

`func (o *InvoiceEnumerationRangeDto) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *InvoiceEnumerationRangeDto) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *InvoiceEnumerationRangeDto) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *InvoiceEnumerationRangeDto) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *InvoiceEnumerationRangeDto) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *InvoiceEnumerationRangeDto) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *InvoiceEnumerationRangeDto) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *InvoiceEnumerationRangeDto) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetFiscalAuthorityId

`func (o *InvoiceEnumerationRangeDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *InvoiceEnumerationRangeDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *InvoiceEnumerationRangeDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *InvoiceEnumerationRangeDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *InvoiceEnumerationRangeDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *InvoiceEnumerationRangeDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetTenantId

`func (o *InvoiceEnumerationRangeDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceEnumerationRangeDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceEnumerationRangeDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceEnumerationRangeDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceEnumerationRangeDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceEnumerationRangeDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceEnumerationRangeDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceEnumerationRangeDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceEnumerationRangeDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceEnumerationRangeDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceEnumerationRangeDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceEnumerationRangeDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetDocumentType

`func (o *InvoiceEnumerationRangeDto) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *InvoiceEnumerationRangeDto) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *InvoiceEnumerationRangeDto) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *InvoiceEnumerationRangeDto) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


