# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Alarms**](SamplesWebhooksApi.md#Alarms) | **Post** /webhook_example/_alarm_ | alarms
[**AssetRaw**](SamplesWebhooksApi.md#AssetRaw) | **Post** /webhook_example/_asset_raw_ | assetRaw
[**Audits**](SamplesWebhooksApi.md#Audits) | **Post** /webhook_example/_audit_ | audits
[**ClientInfo**](SamplesWebhooksApi.md#ClientInfo) | **Post** /webhook_example/_client_info_ | clientInfo
[**ClientJoin**](SamplesWebhooksApi.md#ClientJoin) | **Post** /webhook_example/_client_join_ | clientJoin
[**ClientLatency**](SamplesWebhooksApi.md#ClientLatency) | **Post** /webhook_example/_client_latency_ | client_latency
[**ClientSessions**](SamplesWebhooksApi.md#ClientSessions) | **Post** /webhook_example/_client_sessions_ | clientSessions
[**DeviceEvents**](SamplesWebhooksApi.md#DeviceEvents) | **Post** /webhook_example/_device_events_ | deviceEvents
[**DeviceUpDown**](SamplesWebhooksApi.md#DeviceUpDown) | **Post** /webhook_example/_device_updowns_ | deviceUpDown
[**DiscoveredRawRssi**](SamplesWebhooksApi.md#DiscoveredRawRssi) | **Post** /webhook_example/_discovered_raw_rssi_ | discovered-raw-rssi
[**Location**](SamplesWebhooksApi.md#Location) | **Post** /webhook_example/_location_ | location
[**LocationAsset**](SamplesWebhooksApi.md#LocationAsset) | **Post** /webhook_example/_location_asset_ | location_asset
[**LocationCentrak**](SamplesWebhooksApi.md#LocationCentrak) | **Post** /webhook_example/_location_centrak_ | location_centrak
[**LocationClient**](SamplesWebhooksApi.md#LocationClient) | **Post** /webhook_example/_location_client_ | location_client
[**LocationSdk**](SamplesWebhooksApi.md#LocationSdk) | **Post** /webhook_example/_location_sdk_ | location_sdk
[**LocationUnclient**](SamplesWebhooksApi.md#LocationUnclient) | **Post** /webhook_example/_location_unclient_ | location_unclient
[**NacAccounting**](SamplesWebhooksApi.md#NacAccounting) | **Post** /webhook_example/_nac_accounting_ | nacAccounting
[**NacEvents**](SamplesWebhooksApi.md#NacEvents) | **Post** /webhook_example/_nac_events_ | nac_events
[**OccupancyAlerts**](SamplesWebhooksApi.md#OccupancyAlerts) | **Post** /webhook_example/_occupancy_alerts_ | occupancyAlerts
[**Ping**](SamplesWebhooksApi.md#Ping) | **Post** /webhook_example/_ping_ | ping
[**SdkclientScanData**](SamplesWebhooksApi.md#SdkclientScanData) | **Post** /webhook_example/_sdkclient_scan_data | sdkclientScanData
[**SiteSle**](SamplesWebhooksApi.md#SiteSle) | **Post** /webhook_example/_site_sle_ | site_sle
[**Zone**](SamplesWebhooksApi.md#Zone) | **Post** /webhook_example/_zone_ | zone

# **Alarms**
> Alarms(ctx, optional)
alarms

Webhook sample for `alarm` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleAlarmBody**](WebhookExampleAlarmBody.md)| **N.B.**: Fields like &#x60;aps&#x60;, &#x60;bssids&#x60;, &#x60;ssids&#x60; are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose &#x60;details&#x60; to include more event specific details.

Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetRaw**
> AssetRaw(ctx, optional)
assetRaw

Webhook sample for `asset_raw` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages  **will be deprecated after 06/30/2024**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiAssetRawOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiAssetRawOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleAssetRawBody**](WebhookExampleAssetRawBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Audits**
> Audits(ctx, optional)
audits

Webhook sample for `audit` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiAuditsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiAuditsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleAuditBody**](WebhookExampleAuditBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientInfo**
> ClientInfo(ctx, optional)
clientInfo

Webhook sample for `client_info` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiClientInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiClientInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleClientInfoBody**](WebhookExampleClientInfoBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientJoin**
> ClientJoin(ctx, optional)
clientJoin

Webhook sample for `client_join` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiClientJoinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiClientJoinOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleClientJoinBody**](WebhookExampleClientJoinBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientLatency**
> ClientLatency(ctx, optional)
client_latency

Webhook sample for `client-latency` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiClientLatencyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiClientLatencyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleClientLatencyBody**](WebhookExampleClientLatencyBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientSessions**
> ClientSessions(ctx, optional)
clientSessions

Webhook sample for `client_sessions` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiClientSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiClientSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleClientSessionsBody**](WebhookExampleClientSessionsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeviceEvents**
> DeviceEvents(ctx, optional)
deviceEvents

Webhook sample for `device_events` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleDeviceEventsBody**](WebhookExampleDeviceEventsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeviceUpDown**
> DeviceUpDown(ctx, optional)
deviceUpDown

Webhook sample for `device_updowns` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiDeviceUpDownOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiDeviceUpDownOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleDeviceUpdownsBody**](WebhookExampleDeviceUpdownsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoveredRawRssi**
> DiscoveredRawRssi(ctx, optional)
discovered-raw-rssi

Webhook sample for `discovered-raw-rssi` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiDiscoveredRawRssiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiDiscoveredRawRssiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleDiscoveredRawRssiBody**](WebhookExampleDiscoveredRawRssiBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Location**
> Location(ctx, optional)
location

Webhook sample for `location` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiLocationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleLocationBody**](WebhookExampleLocationBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationAsset**
> LocationAsset(ctx, optional)
location_asset

Webhook sample for `location_asset` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiLocationAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiLocationAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleLocationAssetBody**](WebhookExampleLocationAssetBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationCentrak**
> LocationCentrak(ctx, optional)
location_centrak

Webhook sample for `location_centrak` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiLocationCentrakOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiLocationCentrakOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleLocationCentrakBody**](WebhookExampleLocationCentrakBody.md)| **N.B.**: Fields like &#x60;aps&#x60;, &#x60;bssids&#x60;, &#x60;ssids&#x60; are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose &#x60;details&#x60; to include more event specific details.

Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object. | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationClient**
> LocationClient(ctx, optional)
location_client

Webhook sample for `location_client` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiLocationClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiLocationClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleLocationClientBody**](WebhookExampleLocationClientBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationSdk**
> LocationSdk(ctx, optional)
location_sdk

Webhook sample for `location_sdk` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiLocationSdkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiLocationSdkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleLocationSdkBody**](WebhookExampleLocationSdkBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationUnclient**
> LocationUnclient(ctx, optional)
location_unclient

Webhook sample for `location_unclient` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiLocationUnclientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiLocationUnclientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleLocationUnclientBody**](WebhookExampleLocationUnclientBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NacAccounting**
> NacAccounting(ctx, optional)
nacAccounting

Webhook sample for `nac-accounting` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiNacAccountingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiNacAccountingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleNacAccountingBody**](WebhookExampleNacAccountingBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NacEvents**
> NacEvents(ctx, optional)
nac_events

Example Delivery of nac_events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiNacEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiNacEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleNacEventsBody**](WebhookExampleNacEventsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OccupancyAlerts**
> OccupancyAlerts(ctx, optional)
occupancyAlerts

Webhook sample for `occupancy_alerts` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiOccupancyAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiOccupancyAlertsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleOccupancyAlertsBody**](WebhookExampleOccupancyAlertsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> Ping(ctx, optional)
ping

Webhook sample for `ping` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiPingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiPingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExamplePingBody**](WebhookExamplePingBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SdkclientScanData**
> SdkclientScanData(ctx, optional)
sdkclientScanData

Webhook sample for `sdkclient_scan_data` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiSdkclientScanDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiSdkclientScanDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleSdkclientScanDataBody**](WebhookExampleSdkclientScanDataBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SiteSle**
> SiteSle(ctx, optional)
site_sle

Webhook sample for `site_sle` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiSiteSleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiSiteSleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleSiteSleBody**](WebhookExampleSiteSleBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Zone**
> Zone(ctx, optional)
zone

Webhook sample for `zone` topic  **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SamplesWebhooksApiZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SamplesWebhooksApiZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookExampleZoneBody**](WebhookExampleZoneBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

