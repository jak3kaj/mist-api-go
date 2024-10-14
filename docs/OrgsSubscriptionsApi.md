# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribeOrgAlarmsReports**](OrgsSubscriptionsApi.md#SubscribeOrgAlarmsReports) | **Post** /api/v1/orgs/{org_id}/subscriptions | subscribeOrgAlarmsReports
[**UnsubscribeOrgAlarmsReports**](OrgsSubscriptionsApi.md#UnsubscribeOrgAlarmsReports) | **Delete** /api/v1/orgs/{org_id}/subscriptions | unsubscribeOrgAlarmsReports

# **SubscribeOrgAlarmsReports**
> SubscribeOrgAlarmsReports(ctx, orgId)
subscribeOrgAlarmsReports

Subscribe to Org Alarms/Reports Subscriptions define how Org Alarms/Reports are delivered to whom

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeOrgAlarmsReports**
> UnsubscribeOrgAlarmsReports(ctx, orgId)
unsubscribeOrgAlarmsReports

Unsubscribe from Org Alarms/Reports Subscriptions define how Org Alarms/Reports are delivered to whom

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

