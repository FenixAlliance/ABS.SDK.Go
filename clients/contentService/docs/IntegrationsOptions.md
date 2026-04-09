# IntegrationsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**EPayco** | Pointer to [**EPaycoIntegrationOptions**](EPaycoIntegrationOptions.md) |  | [optional] 
**Google** | Pointer to [**GoogleIntegrationOptions**](GoogleIntegrationOptions.md) |  | [optional] 
**Facebook** | Pointer to [**FacebookIntegrationOptions**](FacebookIntegrationOptions.md) |  | [optional] 
**Sendgrid** | Pointer to [**SendgridIntegrationsOptions**](SendgridIntegrationsOptions.md) |  | [optional] 
**FreeGeoIP** | Pointer to [**FreeGeoIPIntegrationOptions**](FreeGeoIPIntegrationOptions.md) |  | [optional] 
**Microsoft** | Pointer to [**MicrosoftIntegrationsOptions**](MicrosoftIntegrationsOptions.md) |  | [optional] 
**FenixAlliance** | Pointer to [**FenixAllianceIntegrationsOptions**](FenixAllianceIntegrationsOptions.md) |  | [optional] 
**OpenExchangeRates** | Pointer to [**OpenExchangeRatesIntegrationsOptions**](OpenExchangeRatesIntegrationsOptions.md) |  | [optional] 

## Methods

### NewIntegrationsOptions

`func NewIntegrationsOptions() *IntegrationsOptions`

NewIntegrationsOptions instantiates a new IntegrationsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsOptionsWithDefaults

`func NewIntegrationsOptionsWithDefaults() *IntegrationsOptions`

NewIntegrationsOptionsWithDefaults instantiates a new IntegrationsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IntegrationsOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IntegrationsOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IntegrationsOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IntegrationsOptions) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEPayco

`func (o *IntegrationsOptions) GetEPayco() EPaycoIntegrationOptions`

GetEPayco returns the EPayco field if non-nil, zero value otherwise.

### GetEPaycoOk

`func (o *IntegrationsOptions) GetEPaycoOk() (*EPaycoIntegrationOptions, bool)`

GetEPaycoOk returns a tuple with the EPayco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPayco

`func (o *IntegrationsOptions) SetEPayco(v EPaycoIntegrationOptions)`

SetEPayco sets EPayco field to given value.

### HasEPayco

`func (o *IntegrationsOptions) HasEPayco() bool`

HasEPayco returns a boolean if a field has been set.

### GetGoogle

`func (o *IntegrationsOptions) GetGoogle() GoogleIntegrationOptions`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *IntegrationsOptions) GetGoogleOk() (*GoogleIntegrationOptions, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *IntegrationsOptions) SetGoogle(v GoogleIntegrationOptions)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *IntegrationsOptions) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetFacebook

`func (o *IntegrationsOptions) GetFacebook() FacebookIntegrationOptions`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *IntegrationsOptions) GetFacebookOk() (*FacebookIntegrationOptions, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *IntegrationsOptions) SetFacebook(v FacebookIntegrationOptions)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *IntegrationsOptions) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetSendgrid

`func (o *IntegrationsOptions) GetSendgrid() SendgridIntegrationsOptions`

GetSendgrid returns the Sendgrid field if non-nil, zero value otherwise.

### GetSendgridOk

`func (o *IntegrationsOptions) GetSendgridOk() (*SendgridIntegrationsOptions, bool)`

GetSendgridOk returns a tuple with the Sendgrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendgrid

`func (o *IntegrationsOptions) SetSendgrid(v SendgridIntegrationsOptions)`

SetSendgrid sets Sendgrid field to given value.

### HasSendgrid

`func (o *IntegrationsOptions) HasSendgrid() bool`

HasSendgrid returns a boolean if a field has been set.

### GetFreeGeoIP

`func (o *IntegrationsOptions) GetFreeGeoIP() FreeGeoIPIntegrationOptions`

GetFreeGeoIP returns the FreeGeoIP field if non-nil, zero value otherwise.

### GetFreeGeoIPOk

`func (o *IntegrationsOptions) GetFreeGeoIPOk() (*FreeGeoIPIntegrationOptions, bool)`

GetFreeGeoIPOk returns a tuple with the FreeGeoIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeGeoIP

`func (o *IntegrationsOptions) SetFreeGeoIP(v FreeGeoIPIntegrationOptions)`

SetFreeGeoIP sets FreeGeoIP field to given value.

### HasFreeGeoIP

`func (o *IntegrationsOptions) HasFreeGeoIP() bool`

HasFreeGeoIP returns a boolean if a field has been set.

### GetMicrosoft

`func (o *IntegrationsOptions) GetMicrosoft() MicrosoftIntegrationsOptions`

GetMicrosoft returns the Microsoft field if non-nil, zero value otherwise.

### GetMicrosoftOk

`func (o *IntegrationsOptions) GetMicrosoftOk() (*MicrosoftIntegrationsOptions, bool)`

GetMicrosoftOk returns a tuple with the Microsoft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoft

`func (o *IntegrationsOptions) SetMicrosoft(v MicrosoftIntegrationsOptions)`

SetMicrosoft sets Microsoft field to given value.

### HasMicrosoft

`func (o *IntegrationsOptions) HasMicrosoft() bool`

HasMicrosoft returns a boolean if a field has been set.

### GetFenixAlliance

`func (o *IntegrationsOptions) GetFenixAlliance() FenixAllianceIntegrationsOptions`

GetFenixAlliance returns the FenixAlliance field if non-nil, zero value otherwise.

### GetFenixAllianceOk

`func (o *IntegrationsOptions) GetFenixAllianceOk() (*FenixAllianceIntegrationsOptions, bool)`

GetFenixAllianceOk returns a tuple with the FenixAlliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFenixAlliance

`func (o *IntegrationsOptions) SetFenixAlliance(v FenixAllianceIntegrationsOptions)`

SetFenixAlliance sets FenixAlliance field to given value.

### HasFenixAlliance

`func (o *IntegrationsOptions) HasFenixAlliance() bool`

HasFenixAlliance returns a boolean if a field has been set.

### GetOpenExchangeRates

`func (o *IntegrationsOptions) GetOpenExchangeRates() OpenExchangeRatesIntegrationsOptions`

GetOpenExchangeRates returns the OpenExchangeRates field if non-nil, zero value otherwise.

### GetOpenExchangeRatesOk

`func (o *IntegrationsOptions) GetOpenExchangeRatesOk() (*OpenExchangeRatesIntegrationsOptions, bool)`

GetOpenExchangeRatesOk returns a tuple with the OpenExchangeRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenExchangeRates

`func (o *IntegrationsOptions) SetOpenExchangeRates(v OpenExchangeRatesIntegrationsOptions)`

SetOpenExchangeRates sets OpenExchangeRates field to given value.

### HasOpenExchangeRates

`func (o *IntegrationsOptions) HasOpenExchangeRates() bool`

HasOpenExchangeRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


