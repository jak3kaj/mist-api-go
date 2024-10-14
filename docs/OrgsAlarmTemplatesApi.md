# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAlarmTemplate**](OrgsAlarmTemplatesApi.md#CreateOrgAlarmTemplate) | **Post** /api/v1/orgs/{org_id}/alarmtemplates | createOrgAlarmTemplate
[**DeleteOrgAlarmTemplate**](OrgsAlarmTemplatesApi.md#DeleteOrgAlarmTemplate) | **Delete** /api/v1/orgs/{org_id}/alarmtemplates/{alarmtemplate_id} | deleteOrgAlarmTemplate
[**GetOrgAlarmTemplate**](OrgsAlarmTemplatesApi.md#GetOrgAlarmTemplate) | **Get** /api/v1/orgs/{org_id}/alarmtemplates/{alarmtemplate_id} | getOrgAlarmTemplate
[**ListOrgAlarmTemplates**](OrgsAlarmTemplatesApi.md#ListOrgAlarmTemplates) | **Get** /api/v1/orgs/{org_id}/alarmtemplates | listOrgAlarmTemplates
[**ListOrgSuppressedAlarms**](OrgsAlarmTemplatesApi.md#ListOrgSuppressedAlarms) | **Get** /api/v1/orgs/{org_id}/alarmtemplates/suppress | listOrgSuppressedAlarms
[**SuppressOrgAlarm**](OrgsAlarmTemplatesApi.md#SuppressOrgAlarm) | **Post** /api/v1/orgs/{org_id}/alarmtemplates/suppress | suppressOrgAlarm
[**UnsuppressOrgSuppressedAlarms**](OrgsAlarmTemplatesApi.md#UnsuppressOrgSuppressedAlarms) | **Delete** /api/v1/orgs/{org_id}/alarmtemplates/suppress | unsuppressOrgSuppressedAlarms
[**UpdateOrgAlarmTemplate**](OrgsAlarmTemplatesApi.md#UpdateOrgAlarmTemplate) | **Put** /api/v1/orgs/{org_id}/alarmtemplates/{alarmtemplate_id} | updateOrgAlarmTemplate

# **CreateOrgAlarmTemplate**
> InlineResponse20028 CreateOrgAlarmTemplate(ctx, orgId, optional)
createOrgAlarmTemplate

Available rules can be found in Orgs>Consts>getAlarmDefs  The delivery dict is only required if different from the template delivery settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmTemplatesApiCreateOrgAlarmTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmTemplatesApiCreateOrgAlarmTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdAlarmtemplatesBody**](OrgIdAlarmtemplatesBody.md)| Request Body | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgAlarmTemplate**
> DeleteOrgAlarmTemplate(ctx, orgId, alarmtemplateId)
deleteOrgAlarmTemplate

Delete Org Alarm Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **alarmtemplateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgAlarmTemplate**
> InlineResponse20028 GetOrgAlarmTemplate(ctx, orgId, alarmtemplateId)
getOrgAlarmTemplate

Get Org Alarm Template Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **alarmtemplateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAlarmTemplates**
> []AlarmTemplate ListOrgAlarmTemplates(ctx, orgId, optional)
listOrgAlarmTemplates

Get List of Org Alarm Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmTemplatesApiListOrgAlarmTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmTemplatesApiListOrgAlarmTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]AlarmTemplate**](alarm_template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSuppressedAlarms**
> InlineResponse20029 ListOrgSuppressedAlarms(ctx, orgId, optional)
listOrgSuppressedAlarms

Get List of Org Alarms Currently Suppressed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmTemplatesApiListOrgSuppressedAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmTemplatesApiListOrgSuppressedAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**optional.Interface of Scope**](.md)| returns both scopes if not specified | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuppressOrgAlarm**
> SuppressOrgAlarm(ctx, orgId, optional)
suppressOrgAlarm

In certain situations, for example, scheduled maintenance, you may want to suspend alarms to be triggered against Sites for a period of time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmTemplatesApiSuppressOrgAlarmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmTemplatesApiSuppressOrgAlarmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmtemplatesSuppressBody**](AlarmtemplatesSuppressBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsuppressOrgSuppressedAlarms**
> UnsuppressOrgSuppressedAlarms(ctx, orgId)
unsuppressOrgSuppressedAlarms

Un-Suppress Suppressed Alarms

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

# **UpdateOrgAlarmTemplate**
> InlineResponse20028 UpdateOrgAlarmTemplate(ctx, orgId, alarmtemplateId, optional)
updateOrgAlarmTemplate

Update Org Alarm Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **alarmtemplateId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmTemplatesApiUpdateOrgAlarmTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmTemplatesApiUpdateOrgAlarmTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AlarmtemplatesAlarmtemplateIdBody**](AlarmtemplatesAlarmtemplateIdBody.md)| Request Body | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

