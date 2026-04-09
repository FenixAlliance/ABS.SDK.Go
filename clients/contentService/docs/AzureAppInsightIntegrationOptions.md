# AzureAppInsightIntegrationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**EnableDiagnosticsTelemetryModule** | Pointer to **bool** |  | [optional] 
**EnableAuthenticationTrackingJavaScript** | Pointer to **bool** |  | [optional] 
**EnableRequestTrackingTelemetryModule** | Pointer to **bool** |  | [optional] 
**InjectResponseHeaders** | Pointer to **bool** |  | [optional] 
**TrackExceptions** | Pointer to **bool** |  | [optional] 
**EnableW3CDistributedTracing** | Pointer to **bool** |  | [optional] 
**AddAutoCollectedMetricExtractor** | Pointer to **bool** |  | [optional] 
**EnableHeartbeat** | Pointer to **bool** |  | [optional] 
**EnableDebugLogger** | Pointer to **bool** |  | [optional] 
**EndpointAddress** | Pointer to **NullableString** |  | [optional] 
**DeveloperMode** | Pointer to **NullableBool** |  | [optional] 
**ApplicationVersion** | Pointer to **NullableString** |  | [optional] 
**ConnectionString** | Pointer to **NullableString** |  | [optional] 
**InstrumentationKey** | Pointer to **NullableString** |  | [optional] 
**EnableAdaptiveSampling** | Pointer to **bool** |  | [optional] 
**EnableEventCounterCollectionModule** | Pointer to **bool** |  | [optional] 
**EnableDependencyTrackingTelemetryModule** | Pointer to **bool** |  | [optional] 
**EnableAzureInstanceMetadataTelemetryModule** | Pointer to **bool** |  | [optional] 
**EnableAppServicesHeartbeatTelemetryModule** | Pointer to **bool** |  | [optional] 
**EnablePerformanceCounterCollectionModule** | Pointer to **bool** |  | [optional] 
**EnableQuickPulseMetricStream** | Pointer to **bool** |  | [optional] 
**EnableLegacyCorrelationHeadersInjection** | Pointer to **bool** |  | [optional] 
**EnableActiveTelemetryConfigurationSetup** | Pointer to **bool** |  | [optional] 

## Methods

### NewAzureAppInsightIntegrationOptions

`func NewAzureAppInsightIntegrationOptions() *AzureAppInsightIntegrationOptions`

NewAzureAppInsightIntegrationOptions instantiates a new AzureAppInsightIntegrationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAppInsightIntegrationOptionsWithDefaults

`func NewAzureAppInsightIntegrationOptionsWithDefaults() *AzureAppInsightIntegrationOptions`

NewAzureAppInsightIntegrationOptionsWithDefaults instantiates a new AzureAppInsightIntegrationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AzureAppInsightIntegrationOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AzureAppInsightIntegrationOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AzureAppInsightIntegrationOptions) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableDiagnosticsTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) GetEnableDiagnosticsTelemetryModule() bool`

GetEnableDiagnosticsTelemetryModule returns the EnableDiagnosticsTelemetryModule field if non-nil, zero value otherwise.

### GetEnableDiagnosticsTelemetryModuleOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableDiagnosticsTelemetryModuleOk() (*bool, bool)`

GetEnableDiagnosticsTelemetryModuleOk returns a tuple with the EnableDiagnosticsTelemetryModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiagnosticsTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) SetEnableDiagnosticsTelemetryModule(v bool)`

SetEnableDiagnosticsTelemetryModule sets EnableDiagnosticsTelemetryModule field to given value.

### HasEnableDiagnosticsTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) HasEnableDiagnosticsTelemetryModule() bool`

HasEnableDiagnosticsTelemetryModule returns a boolean if a field has been set.

### GetEnableAuthenticationTrackingJavaScript

`func (o *AzureAppInsightIntegrationOptions) GetEnableAuthenticationTrackingJavaScript() bool`

GetEnableAuthenticationTrackingJavaScript returns the EnableAuthenticationTrackingJavaScript field if non-nil, zero value otherwise.

### GetEnableAuthenticationTrackingJavaScriptOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableAuthenticationTrackingJavaScriptOk() (*bool, bool)`

GetEnableAuthenticationTrackingJavaScriptOk returns a tuple with the EnableAuthenticationTrackingJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuthenticationTrackingJavaScript

`func (o *AzureAppInsightIntegrationOptions) SetEnableAuthenticationTrackingJavaScript(v bool)`

SetEnableAuthenticationTrackingJavaScript sets EnableAuthenticationTrackingJavaScript field to given value.

### HasEnableAuthenticationTrackingJavaScript

`func (o *AzureAppInsightIntegrationOptions) HasEnableAuthenticationTrackingJavaScript() bool`

HasEnableAuthenticationTrackingJavaScript returns a boolean if a field has been set.

### GetEnableRequestTrackingTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) GetEnableRequestTrackingTelemetryModule() bool`

GetEnableRequestTrackingTelemetryModule returns the EnableRequestTrackingTelemetryModule field if non-nil, zero value otherwise.

### GetEnableRequestTrackingTelemetryModuleOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableRequestTrackingTelemetryModuleOk() (*bool, bool)`

GetEnableRequestTrackingTelemetryModuleOk returns a tuple with the EnableRequestTrackingTelemetryModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRequestTrackingTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) SetEnableRequestTrackingTelemetryModule(v bool)`

SetEnableRequestTrackingTelemetryModule sets EnableRequestTrackingTelemetryModule field to given value.

### HasEnableRequestTrackingTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) HasEnableRequestTrackingTelemetryModule() bool`

HasEnableRequestTrackingTelemetryModule returns a boolean if a field has been set.

### GetInjectResponseHeaders

`func (o *AzureAppInsightIntegrationOptions) GetInjectResponseHeaders() bool`

GetInjectResponseHeaders returns the InjectResponseHeaders field if non-nil, zero value otherwise.

### GetInjectResponseHeadersOk

`func (o *AzureAppInsightIntegrationOptions) GetInjectResponseHeadersOk() (*bool, bool)`

GetInjectResponseHeadersOk returns a tuple with the InjectResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectResponseHeaders

`func (o *AzureAppInsightIntegrationOptions) SetInjectResponseHeaders(v bool)`

SetInjectResponseHeaders sets InjectResponseHeaders field to given value.

### HasInjectResponseHeaders

`func (o *AzureAppInsightIntegrationOptions) HasInjectResponseHeaders() bool`

HasInjectResponseHeaders returns a boolean if a field has been set.

### GetTrackExceptions

`func (o *AzureAppInsightIntegrationOptions) GetTrackExceptions() bool`

GetTrackExceptions returns the TrackExceptions field if non-nil, zero value otherwise.

### GetTrackExceptionsOk

`func (o *AzureAppInsightIntegrationOptions) GetTrackExceptionsOk() (*bool, bool)`

GetTrackExceptionsOk returns a tuple with the TrackExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackExceptions

`func (o *AzureAppInsightIntegrationOptions) SetTrackExceptions(v bool)`

SetTrackExceptions sets TrackExceptions field to given value.

### HasTrackExceptions

`func (o *AzureAppInsightIntegrationOptions) HasTrackExceptions() bool`

HasTrackExceptions returns a boolean if a field has been set.

### GetEnableW3CDistributedTracing

`func (o *AzureAppInsightIntegrationOptions) GetEnableW3CDistributedTracing() bool`

GetEnableW3CDistributedTracing returns the EnableW3CDistributedTracing field if non-nil, zero value otherwise.

### GetEnableW3CDistributedTracingOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableW3CDistributedTracingOk() (*bool, bool)`

GetEnableW3CDistributedTracingOk returns a tuple with the EnableW3CDistributedTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableW3CDistributedTracing

`func (o *AzureAppInsightIntegrationOptions) SetEnableW3CDistributedTracing(v bool)`

SetEnableW3CDistributedTracing sets EnableW3CDistributedTracing field to given value.

### HasEnableW3CDistributedTracing

`func (o *AzureAppInsightIntegrationOptions) HasEnableW3CDistributedTracing() bool`

HasEnableW3CDistributedTracing returns a boolean if a field has been set.

### GetAddAutoCollectedMetricExtractor

`func (o *AzureAppInsightIntegrationOptions) GetAddAutoCollectedMetricExtractor() bool`

GetAddAutoCollectedMetricExtractor returns the AddAutoCollectedMetricExtractor field if non-nil, zero value otherwise.

### GetAddAutoCollectedMetricExtractorOk

`func (o *AzureAppInsightIntegrationOptions) GetAddAutoCollectedMetricExtractorOk() (*bool, bool)`

GetAddAutoCollectedMetricExtractorOk returns a tuple with the AddAutoCollectedMetricExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAutoCollectedMetricExtractor

`func (o *AzureAppInsightIntegrationOptions) SetAddAutoCollectedMetricExtractor(v bool)`

SetAddAutoCollectedMetricExtractor sets AddAutoCollectedMetricExtractor field to given value.

### HasAddAutoCollectedMetricExtractor

`func (o *AzureAppInsightIntegrationOptions) HasAddAutoCollectedMetricExtractor() bool`

HasAddAutoCollectedMetricExtractor returns a boolean if a field has been set.

### GetEnableHeartbeat

`func (o *AzureAppInsightIntegrationOptions) GetEnableHeartbeat() bool`

GetEnableHeartbeat returns the EnableHeartbeat field if non-nil, zero value otherwise.

### GetEnableHeartbeatOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableHeartbeatOk() (*bool, bool)`

GetEnableHeartbeatOk returns a tuple with the EnableHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHeartbeat

`func (o *AzureAppInsightIntegrationOptions) SetEnableHeartbeat(v bool)`

SetEnableHeartbeat sets EnableHeartbeat field to given value.

### HasEnableHeartbeat

`func (o *AzureAppInsightIntegrationOptions) HasEnableHeartbeat() bool`

HasEnableHeartbeat returns a boolean if a field has been set.

### GetEnableDebugLogger

`func (o *AzureAppInsightIntegrationOptions) GetEnableDebugLogger() bool`

GetEnableDebugLogger returns the EnableDebugLogger field if non-nil, zero value otherwise.

### GetEnableDebugLoggerOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableDebugLoggerOk() (*bool, bool)`

GetEnableDebugLoggerOk returns a tuple with the EnableDebugLogger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugLogger

`func (o *AzureAppInsightIntegrationOptions) SetEnableDebugLogger(v bool)`

SetEnableDebugLogger sets EnableDebugLogger field to given value.

### HasEnableDebugLogger

`func (o *AzureAppInsightIntegrationOptions) HasEnableDebugLogger() bool`

HasEnableDebugLogger returns a boolean if a field has been set.

### GetEndpointAddress

`func (o *AzureAppInsightIntegrationOptions) GetEndpointAddress() string`

GetEndpointAddress returns the EndpointAddress field if non-nil, zero value otherwise.

### GetEndpointAddressOk

`func (o *AzureAppInsightIntegrationOptions) GetEndpointAddressOk() (*string, bool)`

GetEndpointAddressOk returns a tuple with the EndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAddress

`func (o *AzureAppInsightIntegrationOptions) SetEndpointAddress(v string)`

SetEndpointAddress sets EndpointAddress field to given value.

### HasEndpointAddress

`func (o *AzureAppInsightIntegrationOptions) HasEndpointAddress() bool`

HasEndpointAddress returns a boolean if a field has been set.

### SetEndpointAddressNil

`func (o *AzureAppInsightIntegrationOptions) SetEndpointAddressNil(b bool)`

 SetEndpointAddressNil sets the value for EndpointAddress to be an explicit nil

### UnsetEndpointAddress
`func (o *AzureAppInsightIntegrationOptions) UnsetEndpointAddress()`

UnsetEndpointAddress ensures that no value is present for EndpointAddress, not even an explicit nil
### GetDeveloperMode

`func (o *AzureAppInsightIntegrationOptions) GetDeveloperMode() bool`

GetDeveloperMode returns the DeveloperMode field if non-nil, zero value otherwise.

### GetDeveloperModeOk

`func (o *AzureAppInsightIntegrationOptions) GetDeveloperModeOk() (*bool, bool)`

GetDeveloperModeOk returns a tuple with the DeveloperMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperMode

`func (o *AzureAppInsightIntegrationOptions) SetDeveloperMode(v bool)`

SetDeveloperMode sets DeveloperMode field to given value.

### HasDeveloperMode

`func (o *AzureAppInsightIntegrationOptions) HasDeveloperMode() bool`

HasDeveloperMode returns a boolean if a field has been set.

### SetDeveloperModeNil

`func (o *AzureAppInsightIntegrationOptions) SetDeveloperModeNil(b bool)`

 SetDeveloperModeNil sets the value for DeveloperMode to be an explicit nil

### UnsetDeveloperMode
`func (o *AzureAppInsightIntegrationOptions) UnsetDeveloperMode()`

UnsetDeveloperMode ensures that no value is present for DeveloperMode, not even an explicit nil
### GetApplicationVersion

`func (o *AzureAppInsightIntegrationOptions) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *AzureAppInsightIntegrationOptions) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *AzureAppInsightIntegrationOptions) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *AzureAppInsightIntegrationOptions) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersionNil

`func (o *AzureAppInsightIntegrationOptions) SetApplicationVersionNil(b bool)`

 SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil

### UnsetApplicationVersion
`func (o *AzureAppInsightIntegrationOptions) UnsetApplicationVersion()`

UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
### GetConnectionString

`func (o *AzureAppInsightIntegrationOptions) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *AzureAppInsightIntegrationOptions) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *AzureAppInsightIntegrationOptions) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *AzureAppInsightIntegrationOptions) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### SetConnectionStringNil

`func (o *AzureAppInsightIntegrationOptions) SetConnectionStringNil(b bool)`

 SetConnectionStringNil sets the value for ConnectionString to be an explicit nil

### UnsetConnectionString
`func (o *AzureAppInsightIntegrationOptions) UnsetConnectionString()`

UnsetConnectionString ensures that no value is present for ConnectionString, not even an explicit nil
### GetInstrumentationKey

`func (o *AzureAppInsightIntegrationOptions) GetInstrumentationKey() string`

GetInstrumentationKey returns the InstrumentationKey field if non-nil, zero value otherwise.

### GetInstrumentationKeyOk

`func (o *AzureAppInsightIntegrationOptions) GetInstrumentationKeyOk() (*string, bool)`

GetInstrumentationKeyOk returns a tuple with the InstrumentationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentationKey

`func (o *AzureAppInsightIntegrationOptions) SetInstrumentationKey(v string)`

SetInstrumentationKey sets InstrumentationKey field to given value.

### HasInstrumentationKey

`func (o *AzureAppInsightIntegrationOptions) HasInstrumentationKey() bool`

HasInstrumentationKey returns a boolean if a field has been set.

### SetInstrumentationKeyNil

`func (o *AzureAppInsightIntegrationOptions) SetInstrumentationKeyNil(b bool)`

 SetInstrumentationKeyNil sets the value for InstrumentationKey to be an explicit nil

### UnsetInstrumentationKey
`func (o *AzureAppInsightIntegrationOptions) UnsetInstrumentationKey()`

UnsetInstrumentationKey ensures that no value is present for InstrumentationKey, not even an explicit nil
### GetEnableAdaptiveSampling

`func (o *AzureAppInsightIntegrationOptions) GetEnableAdaptiveSampling() bool`

GetEnableAdaptiveSampling returns the EnableAdaptiveSampling field if non-nil, zero value otherwise.

### GetEnableAdaptiveSamplingOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableAdaptiveSamplingOk() (*bool, bool)`

GetEnableAdaptiveSamplingOk returns a tuple with the EnableAdaptiveSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdaptiveSampling

`func (o *AzureAppInsightIntegrationOptions) SetEnableAdaptiveSampling(v bool)`

SetEnableAdaptiveSampling sets EnableAdaptiveSampling field to given value.

### HasEnableAdaptiveSampling

`func (o *AzureAppInsightIntegrationOptions) HasEnableAdaptiveSampling() bool`

HasEnableAdaptiveSampling returns a boolean if a field has been set.

### GetEnableEventCounterCollectionModule

`func (o *AzureAppInsightIntegrationOptions) GetEnableEventCounterCollectionModule() bool`

GetEnableEventCounterCollectionModule returns the EnableEventCounterCollectionModule field if non-nil, zero value otherwise.

### GetEnableEventCounterCollectionModuleOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableEventCounterCollectionModuleOk() (*bool, bool)`

GetEnableEventCounterCollectionModuleOk returns a tuple with the EnableEventCounterCollectionModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEventCounterCollectionModule

`func (o *AzureAppInsightIntegrationOptions) SetEnableEventCounterCollectionModule(v bool)`

SetEnableEventCounterCollectionModule sets EnableEventCounterCollectionModule field to given value.

### HasEnableEventCounterCollectionModule

`func (o *AzureAppInsightIntegrationOptions) HasEnableEventCounterCollectionModule() bool`

HasEnableEventCounterCollectionModule returns a boolean if a field has been set.

### GetEnableDependencyTrackingTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) GetEnableDependencyTrackingTelemetryModule() bool`

GetEnableDependencyTrackingTelemetryModule returns the EnableDependencyTrackingTelemetryModule field if non-nil, zero value otherwise.

### GetEnableDependencyTrackingTelemetryModuleOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableDependencyTrackingTelemetryModuleOk() (*bool, bool)`

GetEnableDependencyTrackingTelemetryModuleOk returns a tuple with the EnableDependencyTrackingTelemetryModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDependencyTrackingTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) SetEnableDependencyTrackingTelemetryModule(v bool)`

SetEnableDependencyTrackingTelemetryModule sets EnableDependencyTrackingTelemetryModule field to given value.

### HasEnableDependencyTrackingTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) HasEnableDependencyTrackingTelemetryModule() bool`

HasEnableDependencyTrackingTelemetryModule returns a boolean if a field has been set.

### GetEnableAzureInstanceMetadataTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) GetEnableAzureInstanceMetadataTelemetryModule() bool`

GetEnableAzureInstanceMetadataTelemetryModule returns the EnableAzureInstanceMetadataTelemetryModule field if non-nil, zero value otherwise.

### GetEnableAzureInstanceMetadataTelemetryModuleOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableAzureInstanceMetadataTelemetryModuleOk() (*bool, bool)`

GetEnableAzureInstanceMetadataTelemetryModuleOk returns a tuple with the EnableAzureInstanceMetadataTelemetryModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAzureInstanceMetadataTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) SetEnableAzureInstanceMetadataTelemetryModule(v bool)`

SetEnableAzureInstanceMetadataTelemetryModule sets EnableAzureInstanceMetadataTelemetryModule field to given value.

### HasEnableAzureInstanceMetadataTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) HasEnableAzureInstanceMetadataTelemetryModule() bool`

HasEnableAzureInstanceMetadataTelemetryModule returns a boolean if a field has been set.

### GetEnableAppServicesHeartbeatTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) GetEnableAppServicesHeartbeatTelemetryModule() bool`

GetEnableAppServicesHeartbeatTelemetryModule returns the EnableAppServicesHeartbeatTelemetryModule field if non-nil, zero value otherwise.

### GetEnableAppServicesHeartbeatTelemetryModuleOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableAppServicesHeartbeatTelemetryModuleOk() (*bool, bool)`

GetEnableAppServicesHeartbeatTelemetryModuleOk returns a tuple with the EnableAppServicesHeartbeatTelemetryModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAppServicesHeartbeatTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) SetEnableAppServicesHeartbeatTelemetryModule(v bool)`

SetEnableAppServicesHeartbeatTelemetryModule sets EnableAppServicesHeartbeatTelemetryModule field to given value.

### HasEnableAppServicesHeartbeatTelemetryModule

`func (o *AzureAppInsightIntegrationOptions) HasEnableAppServicesHeartbeatTelemetryModule() bool`

HasEnableAppServicesHeartbeatTelemetryModule returns a boolean if a field has been set.

### GetEnablePerformanceCounterCollectionModule

`func (o *AzureAppInsightIntegrationOptions) GetEnablePerformanceCounterCollectionModule() bool`

GetEnablePerformanceCounterCollectionModule returns the EnablePerformanceCounterCollectionModule field if non-nil, zero value otherwise.

### GetEnablePerformanceCounterCollectionModuleOk

`func (o *AzureAppInsightIntegrationOptions) GetEnablePerformanceCounterCollectionModuleOk() (*bool, bool)`

GetEnablePerformanceCounterCollectionModuleOk returns a tuple with the EnablePerformanceCounterCollectionModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePerformanceCounterCollectionModule

`func (o *AzureAppInsightIntegrationOptions) SetEnablePerformanceCounterCollectionModule(v bool)`

SetEnablePerformanceCounterCollectionModule sets EnablePerformanceCounterCollectionModule field to given value.

### HasEnablePerformanceCounterCollectionModule

`func (o *AzureAppInsightIntegrationOptions) HasEnablePerformanceCounterCollectionModule() bool`

HasEnablePerformanceCounterCollectionModule returns a boolean if a field has been set.

### GetEnableQuickPulseMetricStream

`func (o *AzureAppInsightIntegrationOptions) GetEnableQuickPulseMetricStream() bool`

GetEnableQuickPulseMetricStream returns the EnableQuickPulseMetricStream field if non-nil, zero value otherwise.

### GetEnableQuickPulseMetricStreamOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableQuickPulseMetricStreamOk() (*bool, bool)`

GetEnableQuickPulseMetricStreamOk returns a tuple with the EnableQuickPulseMetricStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuickPulseMetricStream

`func (o *AzureAppInsightIntegrationOptions) SetEnableQuickPulseMetricStream(v bool)`

SetEnableQuickPulseMetricStream sets EnableQuickPulseMetricStream field to given value.

### HasEnableQuickPulseMetricStream

`func (o *AzureAppInsightIntegrationOptions) HasEnableQuickPulseMetricStream() bool`

HasEnableQuickPulseMetricStream returns a boolean if a field has been set.

### GetEnableLegacyCorrelationHeadersInjection

`func (o *AzureAppInsightIntegrationOptions) GetEnableLegacyCorrelationHeadersInjection() bool`

GetEnableLegacyCorrelationHeadersInjection returns the EnableLegacyCorrelationHeadersInjection field if non-nil, zero value otherwise.

### GetEnableLegacyCorrelationHeadersInjectionOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableLegacyCorrelationHeadersInjectionOk() (*bool, bool)`

GetEnableLegacyCorrelationHeadersInjectionOk returns a tuple with the EnableLegacyCorrelationHeadersInjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLegacyCorrelationHeadersInjection

`func (o *AzureAppInsightIntegrationOptions) SetEnableLegacyCorrelationHeadersInjection(v bool)`

SetEnableLegacyCorrelationHeadersInjection sets EnableLegacyCorrelationHeadersInjection field to given value.

### HasEnableLegacyCorrelationHeadersInjection

`func (o *AzureAppInsightIntegrationOptions) HasEnableLegacyCorrelationHeadersInjection() bool`

HasEnableLegacyCorrelationHeadersInjection returns a boolean if a field has been set.

### GetEnableActiveTelemetryConfigurationSetup

`func (o *AzureAppInsightIntegrationOptions) GetEnableActiveTelemetryConfigurationSetup() bool`

GetEnableActiveTelemetryConfigurationSetup returns the EnableActiveTelemetryConfigurationSetup field if non-nil, zero value otherwise.

### GetEnableActiveTelemetryConfigurationSetupOk

`func (o *AzureAppInsightIntegrationOptions) GetEnableActiveTelemetryConfigurationSetupOk() (*bool, bool)`

GetEnableActiveTelemetryConfigurationSetupOk returns a tuple with the EnableActiveTelemetryConfigurationSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableActiveTelemetryConfigurationSetup

`func (o *AzureAppInsightIntegrationOptions) SetEnableActiveTelemetryConfigurationSetup(v bool)`

SetEnableActiveTelemetryConfigurationSetup sets EnableActiveTelemetryConfigurationSetup field to given value.

### HasEnableActiveTelemetryConfigurationSetup

`func (o *AzureAppInsightIntegrationOptions) HasEnableActiveTelemetryConfigurationSetup() bool`

HasEnableActiveTelemetryConfigurationSetup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


