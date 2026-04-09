# InvoiceEnumerationRangeUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**CurrentNumeration** | Pointer to **int64** |  | [optional] 
**NumerationFrom** | Pointer to **int64** |  | [optional] 
**NumerationTo** | Pointer to **int64** |  | [optional] 
**ValidFrom** | Pointer to **time.Time** |  | [optional] 
**ValidTo** | Pointer to **time.Time** |  | [optional] 
**FiscalAuthorityId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceEnumerationRangeUpdateDto

`func NewInvoiceEnumerationRangeUpdateDto() *InvoiceEnumerationRangeUpdateDto`

NewInvoiceEnumerationRangeUpdateDto instantiates a new InvoiceEnumerationRangeUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceEnumerationRangeUpdateDtoWithDefaults

`func NewInvoiceEnumerationRangeUpdateDtoWithDefaults() *InvoiceEnumerationRangeUpdateDto`

NewInvoiceEnumerationRangeUpdateDtoWithDefaults instantiates a new InvoiceEnumerationRangeUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *InvoiceEnumerationRangeUpdateDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InvoiceEnumerationRangeUpdateDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *InvoiceEnumerationRangeUpdateDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *InvoiceEnumerationRangeUpdateDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *InvoiceEnumerationRangeUpdateDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetSuffix

`func (o *InvoiceEnumerationRangeUpdateDto) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *InvoiceEnumerationRangeUpdateDto) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *InvoiceEnumerationRangeUpdateDto) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *InvoiceEnumerationRangeUpdateDto) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *InvoiceEnumerationRangeUpdateDto) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIdentifier

`func (o *InvoiceEnumerationRangeUpdateDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InvoiceEnumerationRangeUpdateDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InvoiceEnumerationRangeUpdateDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *InvoiceEnumerationRangeUpdateDto) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *InvoiceEnumerationRangeUpdateDto) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetQualifiedName

`func (o *InvoiceEnumerationRangeUpdateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *InvoiceEnumerationRangeUpdateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *InvoiceEnumerationRangeUpdateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *InvoiceEnumerationRangeUpdateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *InvoiceEnumerationRangeUpdateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetCurrentNumeration

`func (o *InvoiceEnumerationRangeUpdateDto) GetCurrentNumeration() int64`

GetCurrentNumeration returns the CurrentNumeration field if non-nil, zero value otherwise.

### GetCurrentNumerationOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetCurrentNumerationOk() (*int64, bool)`

GetCurrentNumerationOk returns a tuple with the CurrentNumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumeration

`func (o *InvoiceEnumerationRangeUpdateDto) SetCurrentNumeration(v int64)`

SetCurrentNumeration sets CurrentNumeration field to given value.

### HasCurrentNumeration

`func (o *InvoiceEnumerationRangeUpdateDto) HasCurrentNumeration() bool`

HasCurrentNumeration returns a boolean if a field has been set.

### GetNumerationFrom

`func (o *InvoiceEnumerationRangeUpdateDto) GetNumerationFrom() int64`

GetNumerationFrom returns the NumerationFrom field if non-nil, zero value otherwise.

### GetNumerationFromOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetNumerationFromOk() (*int64, bool)`

GetNumerationFromOk returns a tuple with the NumerationFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerationFrom

`func (o *InvoiceEnumerationRangeUpdateDto) SetNumerationFrom(v int64)`

SetNumerationFrom sets NumerationFrom field to given value.

### HasNumerationFrom

`func (o *InvoiceEnumerationRangeUpdateDto) HasNumerationFrom() bool`

HasNumerationFrom returns a boolean if a field has been set.

### GetNumerationTo

`func (o *InvoiceEnumerationRangeUpdateDto) GetNumerationTo() int64`

GetNumerationTo returns the NumerationTo field if non-nil, zero value otherwise.

### GetNumerationToOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetNumerationToOk() (*int64, bool)`

GetNumerationToOk returns a tuple with the NumerationTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerationTo

`func (o *InvoiceEnumerationRangeUpdateDto) SetNumerationTo(v int64)`

SetNumerationTo sets NumerationTo field to given value.

### HasNumerationTo

`func (o *InvoiceEnumerationRangeUpdateDto) HasNumerationTo() bool`

HasNumerationTo returns a boolean if a field has been set.

### GetValidFrom

`func (o *InvoiceEnumerationRangeUpdateDto) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *InvoiceEnumerationRangeUpdateDto) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *InvoiceEnumerationRangeUpdateDto) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *InvoiceEnumerationRangeUpdateDto) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *InvoiceEnumerationRangeUpdateDto) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *InvoiceEnumerationRangeUpdateDto) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetFiscalAuthorityId

`func (o *InvoiceEnumerationRangeUpdateDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *InvoiceEnumerationRangeUpdateDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *InvoiceEnumerationRangeUpdateDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *InvoiceEnumerationRangeUpdateDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *InvoiceEnumerationRangeUpdateDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetTenantId

`func (o *InvoiceEnumerationRangeUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceEnumerationRangeUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceEnumerationRangeUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceEnumerationRangeUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceEnumerationRangeUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceEnumerationRangeUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceEnumerationRangeUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceEnumerationRangeUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceEnumerationRangeUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceEnumerationRangeUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetDocumentType

`func (o *InvoiceEnumerationRangeUpdateDto) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *InvoiceEnumerationRangeUpdateDto) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *InvoiceEnumerationRangeUpdateDto) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *InvoiceEnumerationRangeUpdateDto) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


