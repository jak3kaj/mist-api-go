# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSso**](OrgsSSOApi.md#CreateOrgSso) | **Post** /api/v1/orgs/{org_id}/ssos | createOrgSso
[**DeleteOrgSso**](OrgsSSOApi.md#DeleteOrgSso) | **Delete** /api/v1/orgs/{org_id}/ssos/{sso_id} | deleteOrgSso
[**DownloadOrgSsoSamlMetadata**](OrgsSSOApi.md#DownloadOrgSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id}/metadata.xml | downloadOrgSsoSamlMetadata
[**GetOrgSso**](OrgsSSOApi.md#GetOrgSso) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id} | getOrgSso
[**GetOrgSsoSamlMetadata**](OrgsSSOApi.md#GetOrgSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id}/metadata | getOrgSsoSamlMetadata
[**ListOrgSsoLatestFailures**](OrgsSSOApi.md#ListOrgSsoLatestFailures) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id}/failures | listOrgSsoLatestFailures
[**ListOrgSsos**](OrgsSSOApi.md#ListOrgSsos) | **Get** /api/v1/orgs/{org_id}/ssos | listOrgSsos
[**UpdateOrgSso**](OrgsSSOApi.md#UpdateOrgSso) | **Put** /api/v1/orgs/{org_id}/ssos/{sso_id} | updateOrgSso

# **CreateOrgSso**
> InlineResponse20025 CreateOrgSso(ctx, orgId, optional)
createOrgSso

Create Org SSO Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSSOApiCreateOrgSsoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSSOApiCreateOrgSsoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSsosBody**](OrgIdSsosBody.md)| Request Body | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgSso**
> DeleteOrgSso(ctx, orgId, ssoId)
deleteOrgSso

Delete Org SSO Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadOrgSsoSamlMetadata**
> *os.File DownloadOrgSsoSamlMetadata(ctx, orgId, ssoId)
downloadOrgSsoSamlMetadata

Download Org SSO SAML Metdata  Example of metadata.xml: ```xml <?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\">     <md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\">         <md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/5hdF5g/logout\" />         <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>         <md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" index=\"0\" isDefault=\"true\"/>         <md:AttributeConsumingService index=\"0\">             <md:ServiceName xml:lang=\"en-US\">Mist</md:ServiceName>             <md:RequestedAttribute Name=\"Role\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"true\"/>             <md:RequestedAttribute Name=\"FirstName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>             <md:RequestedAttribute Name=\"LastName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>         </md:AttributeConsumingService>     </md:SPSSODescriptor> </md:EntityDescriptor> ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSso**
> InlineResponse20025 GetOrgSso(ctx, orgId, ssoId)
getOrgSso

Get Org SSO Configuration Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSsoSamlMetadata**
> InlineResponse20027 GetOrgSsoSamlMetadata(ctx, orgId, ssoId)
getOrgSsoSamlMetadata

Get Org SSO SAML Metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSsoLatestFailures**
> InlineResponse20026 ListOrgSsoLatestFailures(ctx, orgId, ssoId, optional)
listOrgSsoLatestFailures

Get List of Org SSO Latest Failures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 
 **optional** | ***OrgsSSOApiListOrgSsoLatestFailuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSSOApiListOrgSsoLatestFailuresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSsos**
> []Sso ListOrgSsos(ctx, orgId, optional)
listOrgSsos

Get List of Org SSO Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSSOApiListOrgSsosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSSOApiListOrgSsosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Sso**](sso.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSso**
> InlineResponse20025 UpdateOrgSso(ctx, orgId, ssoId, optional)
updateOrgSso

Update Org SSO Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 
 **optional** | ***OrgsSSOApiUpdateOrgSsoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSSOApiUpdateOrgSsoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SsosSsoIdBody1**](SsosSsoIdBody1.md)| Request Body | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

