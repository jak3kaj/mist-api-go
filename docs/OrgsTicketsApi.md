# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgTicketComment**](OrgsTicketsApi.md#AddOrgTicketComment) | **Post** /api/v1/orgs/{org_id}/tickets/{ticket_id}/comments | addOrgTicketComment
[**CountOrgTickets**](OrgsTicketsApi.md#CountOrgTickets) | **Get** /api/v1/orgs/{org_id}/tickets/count | countOrgTickets
[**CreateOrgTicket**](OrgsTicketsApi.md#CreateOrgTicket) | **Post** /api/v1/orgs/{org_id}/tickets | createOrgTicket
[**GetOrgTicket**](OrgsTicketsApi.md#GetOrgTicket) | **Get** /api/v1/orgs/{org_id}/tickets/{ticket_id} | getOrgTicket
[**GetOrgTicketAttachment**](OrgsTicketsApi.md#GetOrgTicketAttachment) | **Get** /api/v1/orgs/{org_id}/tickets/{ticket_id}/attachments/{attachment_id} | GetOrgTicketAttachment
[**ListOrgTickets**](OrgsTicketsApi.md#ListOrgTickets) | **Get** /api/v1/orgs/{org_id}/tickets | listOrgTickets
[**UpdateOrgTicket**](OrgsTicketsApi.md#UpdateOrgTicket) | **Put** /api/v1/orgs/{org_id}/tickets/{ticket_id} | updateOrgTicket
[**UploadrgTicketAttachment**](OrgsTicketsApi.md#UploadrgTicketAttachment) | **Post** /api/v1/orgs/{org_id}/tickets/{ticket_id}/attachments | UploadrgTicketAttachment

# **AddOrgTicketComment**
> InlineResponse200104 AddOrgTicketComment(ctx, orgId, ticketId, optional)
addOrgTicketComment

Add Comment to support ticket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ticketId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiAddOrgTicketCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiAddOrgTicketCommentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **comment** | **optional.**|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**InlineResponse200104**](inline_response_200_104.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountOrgTickets**
> InlineResponse20016 CountOrgTickets(ctx, orgId, optional)
countOrgTickets

Count Org Tickets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiCountOrgTicketsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiCountOrgTicketsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgTicketsCountDistinct**](.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrgTicket**
> InlineResponse200104 CreateOrgTicket(ctx, orgId, optional)
createOrgTicket

Create a support ticket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiCreateOrgTicketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiCreateOrgTicketOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdTicketsBody**](OrgIdTicketsBody.md)| Request Body | 

### Return type

[**InlineResponse200104**](inline_response_200_104.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgTicket**
> InlineResponse200104 GetOrgTicket(ctx, orgId, ticketId, optional)
getOrgTicket

Get support ticket details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ticketId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiGetOrgTicketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiGetOrgTicketOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200104**](inline_response_200_104.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgTicketAttachment**
> InlineResponse200105 GetOrgTicketAttachment(ctx, orgId, ticketId, attachmentId, optional)
GetOrgTicketAttachment

Get Org ticket Attachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ticketId** | [**string**](.md)|  | 
  **attachmentId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiGetOrgTicketAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiGetOrgTicketAttachmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200105**](inline_response_200_105.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgTickets**
> []Ticket ListOrgTickets(ctx, orgId, optional)
listOrgTickets

Get List of Tickets of an Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiListOrgTicketsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiListOrgTicketsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**[]Ticket**](ticket.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgTicket**
> InlineResponse200104 UpdateOrgTicket(ctx, orgId, ticketId, optional)
updateOrgTicket

Update support ticket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ticketId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiUpdateOrgTicketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiUpdateOrgTicketOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TicketsTicketIdBody**](TicketsTicketIdBody.md)| Request Body | 

### Return type

[**InlineResponse200104**](inline_response_200_104.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadrgTicketAttachment**
> UploadrgTicketAttachment(ctx, orgId, ticketId, optional)
UploadrgTicketAttachment

Get Org ticket Attachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ticketId** | [**string**](.md)|  | 
 **optional** | ***OrgsTicketsApiUploadrgTicketAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsTicketsApiUploadrgTicketAttachmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

