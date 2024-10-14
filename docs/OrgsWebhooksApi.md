# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgWebhooksDeliveries**](OrgsWebhooksApi.md#CountOrgWebhooksDeliveries) | **Get** /api/v1/orgs/{org_id}/webhooks/{webhook_id}/events/count | countOrgWebhooksDeliveries
[**CreateOrgWebhook**](OrgsWebhooksApi.md#CreateOrgWebhook) | **Post** /api/v1/orgs/{org_id}/webhooks | createOrgWebhook
[**DeleteOrgWebhook**](OrgsWebhooksApi.md#DeleteOrgWebhook) | **Delete** /api/v1/orgs/{org_id}/webhooks/{webhook_id} | deleteOrgWebhook
[**GetOrgWebhook**](OrgsWebhooksApi.md#GetOrgWebhook) | **Get** /api/v1/orgs/{org_id}/webhooks/{webhook_id} | getOrgWebhook
[**ListOrgWebhooks**](OrgsWebhooksApi.md#ListOrgWebhooks) | **Get** /api/v1/orgs/{org_id}/webhooks | listOrgWebhooks
[**PingOrgWebhook**](OrgsWebhooksApi.md#PingOrgWebhook) | **Post** /api/v1/orgs/{org_id}/webhooks/{webhook_id}/ping | pingOrgWebhook
[**SearchOrgWebhooksDeliveries**](OrgsWebhooksApi.md#SearchOrgWebhooksDeliveries) | **Get** /api/v1/orgs/{org_id}/webhooks/{webhook_id}/events/search | searchOrgWebhooksDeliveries
[**UpdateOrgWebhook**](OrgsWebhooksApi.md#UpdateOrgWebhook) | **Put** /api/v1/orgs/{org_id}/webhooks/{webhook_id} | updateOrgWebhook

# **CountOrgWebhooksDeliveries**
> InlineResponse20016 CountOrgWebhooksDeliveries(ctx, orgId, webhookId, optional)
countOrgWebhooksDeliveries

Count Org Webhooks deliveries   Topics Supported: - alarms - audits - device-updowns - occupancy-alerts - ping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 
 **optional** | ***OrgsWebhooksApiCountOrgWebhooksDeliveriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWebhooksApiCountOrgWebhooksDeliveriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **error_** | **optional.String**|  | 
 **statusCode** | **optional.Int32**|  | 
 **status** | [**optional.Interface of Status**](.md)| webhook delivery status | 
 **topic** | [**optional.Interface of Topic**](.md)| webhook topic | 
 **distinct** | [**optional.Interface of Distinct2**](.md)|  | 
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

# **CreateOrgWebhook**
> InlineResponse200111 CreateOrgWebhook(ctx, orgId, optional)
createOrgWebhook

Create Org Webhook  **N.B**. For org webhooks, only alarms/audits/client-info/client-join/client-sessions/device_events/device-updowns/mxedge_events Infrastructure topics are supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWebhooksApiCreateOrgWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWebhooksApiCreateOrgWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdWebhooksBody**](OrgIdWebhooksBody.md)| Request Body | 

### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgWebhook**
> DeleteOrgWebhook(ctx, orgId, webhookId)
deleteOrgWebhook

Delete Org Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgWebhook**
> InlineResponse200111 GetOrgWebhook(ctx, orgId, webhookId)
getOrgWebhook

Get Org Webhook Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgWebhooks**
> []Webhook ListOrgWebhooks(ctx, orgId, optional)
listOrgWebhooks

Get List of Org Webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWebhooksApiListOrgWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWebhooksApiListOrgWebhooksOpts struct
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

# **PingOrgWebhook**
> PingOrgWebhook(ctx, orgId, webhookId)
pingOrgWebhook

send a Ping event to the webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgWebhooksDeliveries**
> InlineResponse200112 SearchOrgWebhooksDeliveries(ctx, orgId, webhookId, optional)
searchOrgWebhooksDeliveries

Search Org Webhooks deliveries   Topics Supported: - alarms - audits - device-updowns - occupancy-alerts - ping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 
 **optional** | ***OrgsWebhooksApiSearchOrgWebhooksDeliveriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWebhooksApiSearchOrgWebhooksDeliveriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **error_** | **optional.String**|  | 
 **statusCode** | **optional.Int32**|  | 
 **status** | [**optional.Interface of Status1**](.md)| webhook delivery status | 
 **topic** | [**optional.Interface of Topic1**](.md)| webhook topic | 
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

# **UpdateOrgWebhook**
> InlineResponse200111 UpdateOrgWebhook(ctx, orgId, webhookId, optional)
updateOrgWebhook

Update Org Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **webhookId** | [**string**](.md)|  | 
 **optional** | ***OrgsWebhooksApiUpdateOrgWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWebhooksApiUpdateOrgWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WebhooksWebhookIdBody**](WebhooksWebhookIdBody.md)| Request Body | 

### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

