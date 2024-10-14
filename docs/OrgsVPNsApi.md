# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgVpns**](OrgsVPNsApi.md#CreateOrgVpns) | **Post** /api/v1/orgs/{org_id}/vpns | createOrgVpns
[**DeleteOrgVpn**](OrgsVPNsApi.md#DeleteOrgVpn) | **Delete** /api/v1/orgs/{org_id}/vpns/{vpn_id} | deleteOrgVpn
[**GetOrgVpn**](OrgsVPNsApi.md#GetOrgVpn) | **Get** /api/v1/orgs/{org_id}/vpns/{vpn_id} | getOrgVpn
[**ListOrgsVpns**](OrgsVPNsApi.md#ListOrgsVpns) | **Get** /api/v1/orgs/{org_id}/vpns | listOrgsVpns
[**UpdateOrgVpn**](OrgsVPNsApi.md#UpdateOrgVpn) | **Put** /api/v1/orgs/{org_id}/vpns/{vpn_id} | updateOrgVpn

# **CreateOrgVpns**
> InlineResponse200110 CreateOrgVpns(ctx, orgId, optional)
createOrgVpns

Create Org VPN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsVPNsApiCreateOrgVpnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsVPNsApiCreateOrgVpnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdVpnsBody**](OrgIdVpnsBody.md)|  | 

### Return type

[**InlineResponse200110**](inline_response_200_110.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgVpn**
> DeleteOrgVpn(ctx, orgId, vpnId)
deleteOrgVpn

delete Org Vpn

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **vpnId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgVpn**
> InlineResponse200110 GetOrgVpn(ctx, orgId, vpnId)
getOrgVpn

getOrgVpn

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **vpnId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200110**](inline_response_200_110.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgsVpns**
> []Vpn ListOrgsVpns(ctx, orgId, optional)
listOrgsVpns

Get List of Org VPNs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsVPNsApiListOrgsVpnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsVPNsApiListOrgsVpnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Vpn**](vpn.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgVpn**
> InlineResponse200110 UpdateOrgVpn(ctx, orgId, vpnId, optional)
updateOrgVpn

update Org Vpn

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **vpnId** | [**string**](.md)|  | 
 **optional** | ***OrgsVPNsApiUpdateOrgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsVPNsApiUpdateOrgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VpnsVpnIdBody**](VpnsVpnIdBody.md)|  | 

### Return type

[**InlineResponse200110**](inline_response_200_110.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

