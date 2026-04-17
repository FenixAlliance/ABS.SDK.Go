# InvoiceEnumerationRangeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**CurrentNumeration** | Pointer to **int64** |  | [optional] 
**NumerationFrom** | Pointer to **int64** |  | [optional] 
**NumerationTo** | Pointer to **int64** |  | [optional] 
**ValidFrom** | **time.Time** |  | 
**ValidTo** | **time.Time** |  | 
**FiscalAuthorityId** | Pointer to **NullableString** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceEnumerationRangeCreateDto

`func NewInvoiceEnumerationRangeCreateDto(validFrom time.Time, validTo time.Time, ) *InvoiceEnumerationRangeCreateDto`

NewInvoiceEnumerationRangeCreateDto instantiates a new InvoiceEnumerationRangeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceEnumerationRangeCreateDtoWithDefaults

`func NewInvoiceEnumerationRangeCreateDtoWithDefaults() *InvoiceEnumerationRangeCreateDto`

NewInvoiceEnumerationRangeCreateDtoWithDefaults instantiates a new InvoiceEnumerationRangeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceEnumerationRangeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceEnumerationRangeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceEnumerationRangeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceEnumerationRangeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *InvoiceEnumerationRangeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceEnumerationRangeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceEnumerationRangeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceEnumerationRangeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPrefix

`func (o *InvoiceEnumerationRangeCreateDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InvoiceEnumerationRangeCreateDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InvoiceEnumerationRangeCreateDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *InvoiceEnumerationRangeCreateDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *InvoiceEnumerationRangeCreateDto) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *InvoiceEnumerationRangeCreateDto) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetSuffix

`func (o *InvoiceEnumerationRangeCreateDto) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *InvoiceEnumerationRangeCreateDto) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *InvoiceEnumerationRangeCreateDto) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *InvoiceEnumerationRangeCreateDto) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *InvoiceEnumerationRangeCreateDto) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *InvoiceEnumerationRangeCreateDto) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetIdentifier

`func (o *InvoiceEnumerationRangeCreateDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InvoiceEnumerationRangeCreateDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InvoiceEnumerationRangeCreateDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InvoiceEnumerationRangeCreateDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *InvoiceEnumerationRangeCreateDto) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *InvoiceEnumerationRangeCreateDto) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetQualifiedName

`func (o *InvoiceEnumerationRangeCreateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *InvoiceEnumerationRangeCreateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *InvoiceEnumerationRangeCreateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *InvoiceEnumerationRangeCreateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *InvoiceEnumerationRangeCreateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *InvoiceEnumerationRangeCreateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetCurrentNumeration

`func (o *InvoiceEnumerationRangeCreateDto) GetCurrentNumeration() int64`

GetCurrentNumeration returns the CurrentNumeration field if non-nil, zero value otherwise.

### GetCurrentNumerationOk

`func (o *InvoiceEnumerationRangeCreateDto) GetCurrentNumerationOk() (*int64, bool)`

GetCurrentNumerationOk returns a tuple with the CurrentNumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumeration

`func (o *InvoiceEnumerationRangeCreateDto) SetCurrentNumeration(v int64)`

SetCurrentNumeration sets CurrentNumeration field to given value.

### HasCurrentNumeration

`func (o *InvoiceEnumerationRangeCreateDto) HasCurrentNumeration() bool`

HasCurrentNumeration returns a boolean if a field has been set.

### GetNumerationFrom

`func (o *InvoiceEnumerationRangeCreateDto) GetNumerationFrom() int64`

GetNumerationFrom returns the NumerationFrom field if non-nil, zero value otherwise.

### GetNumerationFromOk

`func (o *InvoiceEnumerationRangeCreateDto) GetNumerationFromOk() (*int64, bool)`

GetNumerationFromOk returns a tuple with the NumerationFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerationFrom

`func (o *InvoiceEnumerationRangeCreateDto) SetNumerationFrom(v int64)`

SetNumerationFrom sets NumerationFrom field to given value.

### HasNumerationFrom

`func (o *InvoiceEnumerationRangeCreateDto) HasNumerationFrom() bool`

HasNumerationFrom returns a boolean if a field has been set.

### GetNumerationTo

`func (o *InvoiceEnumerationRangeCreateDto) GetNumerationTo() int64`

GetNumerationTo returns the NumerationTo field if non-nil, zero value otherwise.

### GetNumerationToOk

`func (o *InvoiceEnumerationRangeCreateDto) GetNumerationToOk() (*int64, bool)`

GetNumerationToOk returns a tuple with the NumerationTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerationTo

`func (o *InvoiceEnumerationRangeCreateDto) SetNumerationTo(v int64)`

SetNumerationTo sets NumerationTo field to given value.

### HasNumerationTo

`func (o *InvoiceEnumerationRangeCreateDto) HasNumerationTo() bool`

HasNumerationTo returns a boolean if a field has been set.

### GetValidFrom

`func (o *InvoiceEnumerationRangeCreateDto) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *InvoiceEnumerationRangeCreateDto) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *InvoiceEnumerationRangeCreateDto) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetValidTo

`func (o *InvoiceEnumerationRangeCreateDto) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *InvoiceEnumerationRangeCreateDto) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *InvoiceEnumerationRangeCreateDto) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.


### GetFiscalAuthorityId

`func (o *InvoiceEnumerationRangeCreateDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *InvoiceEnumerationRangeCreateDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *InvoiceEnumerationRangeCreateDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *InvoiceEnumerationRangeCreateDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *InvoiceEnumerationRangeCreateDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *InvoiceEnumerationRangeCreateDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetDocumentType

`func (o *InvoiceEnumerationRangeCreateDto) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *InvoiceEnumerationRangeCreateDto) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *InvoiceEnumerationRangeCreateDto) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *InvoiceEnumerationRangeCreateDto) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


