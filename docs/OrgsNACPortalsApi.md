# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNacPortal**](OrgsNACPortalsApi.md#CreateOrgNacPortal) | **Post** /api/v1/orgs/{org_id}/nacportals | createOrgNacPortal
[**DeleteOrgNacPortal**](OrgsNACPortalsApi.md#DeleteOrgNacPortal) | **Delete** /api/v1/orgs/{org_id}/nacportals/{nacportal_id} | deleteOrgNacPortal
[**DeleteOrgNacPortalImage**](OrgsNACPortalsApi.md#DeleteOrgNacPortalImage) | **Delete** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/portal_image | deleteOrgNacPortalImage
[**DownloadOrgNacPortalSsoSamlMetadata**](OrgsNACPortalsApi.md#DownloadOrgNacPortalSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/saml_metadata.xml | downloadOrgNacPortalSsoSamlMetadata
[**GetOrgNacPortal**](OrgsNACPortalsApi.md#GetOrgNacPortal) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id} | getOrgNacPortal
[**GetOrgNacPortalSsoSamlMetadata**](OrgsNACPortalsApi.md#GetOrgNacPortalSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/saml_metadata | getOrgNacPortalSsoSamlMetadata
[**ListOrgNacPortalSsoLatestFailures**](OrgsNACPortalsApi.md#ListOrgNacPortalSsoLatestFailures) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/failures | listOrgNacPortalSsoLatestFailures
[**ListOrgNacPortals**](OrgsNACPortalsApi.md#ListOrgNacPortals) | **Get** /api/v1/orgs/{org_id}/nacportals | listOrgNacPortals
[**UpdateOrgNacPortal**](OrgsNACPortalsApi.md#UpdateOrgNacPortal) | **Put** /api/v1/orgs/{org_id}/nacportals/{nacportal_id} | updateOrgNacPortal
[**UpdateOrgNacPortalTempalte**](OrgsNACPortalsApi.md#UpdateOrgNacPortalTempalte) | **Put** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/portal_template | updateOrgNacPortalTempalte
[**UploadOrgNacPortalImage**](OrgsNACPortalsApi.md#UploadOrgNacPortalImage) | **Post** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/portal_image | uploadOrgNacPortalImage

# **CreateOrgNacPortal**
> InlineResponse20073 CreateOrgNacPortal(ctx, orgId, optional)
createOrgNacPortal

Create Org NAC Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACPortalsApiCreateOrgNacPortalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACPortalsApiCreateOrgNacPortalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NacPortal**](NacPortal.md)|  | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgNacPortal**
> DeleteOrgNacPortal(ctx, orgId, nacportalId)
deleteOrgNacPortal

Delete Org NAC Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgNacPortalImage**
> DeleteOrgNacPortalImage(ctx, orgId, nacportalId)
deleteOrgNacPortalImage

Delete background image for NAC Portal   If image is not uploaded or is deleted, NAC Portal will use default image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadOrgNacPortalSsoSamlMetadata**
> *os.File DownloadOrgNacPortalSsoSamlMetadata(ctx, orgId, nacportalId)
downloadOrgNacPortalSsoSamlMetadata

Download Org NAC Portal SSO SAML Metdata  Example of metadata.xml: ```xml <?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\">     <md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\">         <md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/5hdF5g/logout\" />         <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>         <md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" index=\"0\" isDefault=\"true\"/>         <md:AttributeConsumingService index=\"0\">             <md:ServiceName xml:lang=\"en-US\">Mist</md:ServiceName>             <md:RequestedAttribute Name=\"Role\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"true\"/>             <md:RequestedAttribute Name=\"FirstName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>             <md:RequestedAttribute Name=\"LastName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>         </md:AttributeConsumingService>     </md:SPSSODescriptor> </md:EntityDescriptor> ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgNacPortal**
> InlineResponse20073 GetOrgNacPortal(ctx, orgId, nacportalId)
getOrgNacPortal

Get Org NAC Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgNacPortalSsoSamlMetadata**
> InlineResponse20027 GetOrgNacPortalSsoSamlMetadata(ctx, orgId, nacportalId)
getOrgNacPortalSsoSamlMetadata

Get Org NAC Portal SSO SAML Metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgNacPortalSsoLatestFailures**
> InlineResponse20026 ListOrgNacPortalSsoLatestFailures(ctx, orgId, nacportalId, optional)
listOrgNacPortalSsoLatestFailures

Get List of Org NAC Portal SSO Latest Failures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACPortalsApiListOrgNacPortalSsoLatestFailuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACPortalsApiListOrgNacPortalSsoLatestFailuresOpts struct
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

# **ListOrgNacPortals**
> []NacPortal ListOrgNacPortals(ctx, orgId, optional)
listOrgNacPortals

List Org NAC Portals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACPortalsApiListOrgNacPortalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACPortalsApiListOrgNacPortalsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]NacPortal**](nac_portal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgNacPortal**
> InlineResponse20073 UpdateOrgNacPortal(ctx, orgId, nacportalId, optional)
updateOrgNacPortal

Update Org NAC Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACPortalsApiUpdateOrgNacPortalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACPortalsApiUpdateOrgNacPortalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NacPortal**](NacPortal.md)|  | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgNacPortalTempalte**
> UpdateOrgNacPortalTempalte(ctx, orgId, nacportalId, optional)
updateOrgNacPortalTempalte

Update Org NAC Portal Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACPortalsApiUpdateOrgNacPortalTempalteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACPortalsApiUpdateOrgNacPortalTempalteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NacPortalTemplate**](NacPortalTemplate.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadOrgNacPortalImage**
> UploadOrgNacPortalImage(ctx, orgId, nacportalId, optional)
uploadOrgNacPortalImage

Upload background image for NAC Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacportalId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACPortalsApiUploadOrgNacPortalImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACPortalsApiUploadOrgNacPortalImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

