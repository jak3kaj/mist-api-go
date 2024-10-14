# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWebhooksDeliveries**](SitesWebhooksApi.md#CountSiteWebhooksDeliveries) | **Get** /api/v1/sites/{site_id}/webhooks/{webhook_id}/events/count | countSiteWebhooksDeliveries
[**CreateSiteWebhook**](SitesWebhooksApi.md#CreateSiteWebhook) | **Post** /api/v1/sites/{site_id}/webhooks | createSiteWebhook
[**DeleteSiteWebhook**](SitesWebhooksApi.md#DeleteSiteWebhook) | **Delete** /api/v1/sites/{site_id}/webhooks/{webhook_id} | deleteSiteWebhook
[**GetSiteWebhook**](SitesWebhooksApi.md#GetSiteWebhook) | **Get** /api/v1/sites/{site_id}/webhooks/{webhook_id} | getSiteWebhook
[**ListSiteWebhooks**](SitesWebhooksApi.md#ListSiteWebhooks) | **Get** /api/v1/sites/{site_id}/webhooks | listSiteWebhooks
[**PingSiteWebhook**](SitesWebhooksApi.md#PingSiteWebhook) | **Post** /api/v1/sites/{site_id}/webhooks/{webhook_id}/ping | pingSiteWebhook
[**SearchSiteWebhooksDeliveries**](SitesWebhooksApi.md#SearchSiteWebhooksDeliveries) | **Get** /api/v1/sites/{site_id}/webhooks/{webhook_id}/events/search | searchSiteWebhooksDeliveries
[**UpdateSiteWebhook**](SitesWebhooksApi.md#UpdateSiteWebhook) | **Put** /api/v1/sites/{site_id}/webhooks/{webhook_id} | updateSiteWebhook

# **CountSiteWebhooksDeliveries**
> InlineResponse20016 CountSiteWebhooksDeliveries(ctx, siteId, webhookId, optional)
countSiteWebhooksDeliveries

Count Site Webhooks deliveries   Topics Supported: - alarms - audits - device-updowns - occupancy-alerts - ping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 
 **optional** | ***SitesWebhooksApiCountSiteWebhooksDeliveriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWebhooksApiCountSiteWebhooksDeliveriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **error_** | **optional.String**|  | 
 **statusCode** | **optional.Int32**|  | 
 **status** | [**optional.Interface of Status2**](.md)| webhook delivery status | 
 **topic** | [**optional.Interface of Topic2**](.md)| webhook topic | 
 **distinct** | [**optional.Interface of Distinct8**](.md)|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSiteWebhook**
> InlineResponse200111 CreateSiteWebhook(ctx, siteId, optional)
createSiteWebhook

Webhook defines a webhook, modeled after [github’s model](https://developer.github.com/webhooks/).  There is two types of webhooks: * webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-be01e694-7253-4195-8563-78e2a745e114)) * raw data webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-e2d5d5f8-4bdb-4efc-93e4-90f4b33d0b2b))  ##### Webhooks Webhooks can be configured at the org level (subset of topics only) and at the site level. It is possible to have multiple topics in the same webhook configuration and/or to have multiple webhooks configured at the same time.  ##### Client Raw Data Webhooks Raw data webhooks are a special subset of webhooks that provide insight into raw data packets emitted by a client, identified by their advertising MAC address (assets, discovered ble, connected wifi, unconnected wifi). The data that client raw data webhooks encompasses are reporting AP information, RSSI Data, and any special packets/telemetry packets that the client may emit. Note that client raw webhooks are the raw data coming from the client and do not contain the X,Y location data of the client. In order to get the location data for a client please see our location webhooks. Clients can be identified uniquely across these client raw data topics and location webhook topic using MAC address as the Unique identifier (client identifier).  ###### Client Raw Data Webhooks Topics Topics that correspond to client raw data for different client types.  * `asset-raw-rssi` - Raw data from packets emitted by named and filtered assets  * `discovered-raw-rssi` - Raw data from packets emitted by passive BLE devices  * `wifi-conn-raw` - Raw data from packets emitted by connected devices  * `wifi-unconn-raw` - Raw data from packets emitted by unconnected devices (passive)  ###### Rules for configuring client raw data webhooks 1. Only one instance of a webhook object containing a client raw data webhook topic is allowed. (a site level entry will override an org level entry for the client raw data webhook topic in question) 2. Only one client raw data webhook topic is allowed per `http-post` message to webhooks api 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWebhooksApiCreateSiteWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWebhooksApiCreateSiteWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdWebhooksBody**](SiteIdWebhooksBody.md)| Request Body | 

### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWebhook**
> DeleteSiteWebhook(ctx, siteId, webhookId)
deleteSiteWebhook

Delete Site Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteWebhook**
> InlineResponse200111 GetSiteWebhook(ctx, siteId, webhookId)
getSiteWebhook

Get Site Webhook Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteWebhooks**
> []Webhook ListSiteWebhooks(ctx, siteId, optional)
listSiteWebhooks

Get List of Site Webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWebhooksApiListSiteWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWebhooksApiListSiteWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Webhook**](webhook.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingSiteWebhook**
> PingSiteWebhook(ctx, siteId, webhookId)
pingSiteWebhook

send a Ping event to the webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteWebhooksDeliveries**
> InlineResponse200112 SearchSiteWebhooksDeliveries(ctx, siteId, webhookId, optional)
searchSiteWebhooksDeliveries

Search Site Webhooks deliveries   Topics Supported: - alarms - audits - device-updowns - occupancy-alerts - ping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 
 **optional** | ***SitesWebhooksApiSearchSiteWebhooksDeliveriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWebhooksApiSearchSiteWebhooksDeliveriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **error_** | **optional.String**|  | 
 **statusCode** | **optional.Int32**|  | 
 **status** | [**optional.Interface of Status3**](.md)| webhook delivery status | 
 **topic** | [**optional.Interface of Topic3**](.md)| webhook topic | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse200112**](inline_response_200_112.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteWebhook**
> InlineResponse200111 UpdateSiteWebhook(ctx, siteId, webhookId, optional)
updateSiteWebhook

Update Site Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 
 **optional** | ***SitesWebhooksApiUpdateSiteWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWebhooksApiUpdateSiteWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WebhooksWebhookIdBody1**](WebhooksWebhookIdBody1.md)| Request Body | 

### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

