# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearOrgCertificates**](OrgsCertApi.md#ClearOrgCertificates) | **Post** /api/v1/orgs/{org_id}/cert/regenerate | clearOrgCertificates
[**GetOrgCertificates**](OrgsCertApi.md#GetOrgCertificates) | **Get** /api/v1/orgs/{org_id}/cert | getOrgCertificates
[**GetOrgSslProxyCert**](OrgsCertApi.md#GetOrgSslProxyCert) | **Get** /api/v1/orgs/{org_id}/ssl_proxy_cert | getOrgSslProxyCert
[**TruncateOrgCrlFile**](OrgsCertApi.md#TruncateOrgCrlFile) | **Post** /api/v1/orgs/{org_id}/crl/truncate | truncateOrgCrlFile

# **ClearOrgCertificates**
> InlineResponse20034 ClearOrgCertificates(ctx, orgId)
clearOrgCertificates

Clear Org Certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgCertificates**
> InlineResponse20034 GetOrgCertificates(ctx, orgId)
getOrgCertificates

Get Org Certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSslProxyCert**
> InlineResponse20035 GetOrgSslProxyCert(ctx, orgId)
getOrgSslProxyCert

Get Org SSL proxy Certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TruncateOrgCrlFile**
> TruncateOrgCrlFile(ctx, orgId, optional)
truncateOrgCrlFile

By default, all certs used by recently unclaimed devices within 9 month will be included in CRL. If the list grows too big, you can truncate it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsCertApiTruncateOrgCrlFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsCertApiTruncateOrgCrlFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DaysNumber**](DaysNumber.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

