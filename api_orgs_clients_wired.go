
/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mist

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type OrgsClientsWiredApiService service
/*
OrgsClientsWiredApiService countOrgWiredClients
Count by Distinct Attributes of Clients  Note: For list of avaialable &#x60;type&#x60; values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId
 * @param optional nil or *OrgsClientsWiredApiCountOrgWiredClientsOpts - Optional Parameters:
     * @param "Distinct" (optional.Interface of OrgWiredClientsCountDistinct) - 
     * @param "Start" (optional.Int32) -  start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
     * @param "End" (optional.Int32) -  end datetime, can be epoch or relative time like -1d, -2h; now if not specified
     * @param "Duration" (optional.String) -  duration like 7d, 2w
     * @param "Limit" (optional.Int32) - 
     * @param "Page" (optional.Int32) - 
@return InlineResponse20016
*/

type OrgsClientsWiredApiCountOrgWiredClientsOpts struct {
    Distinct optional.Interface
    Start optional.Int32
    End optional.Int32
    Duration optional.String
    Limit optional.Int32
    Page optional.Int32
}

func (a *OrgsClientsWiredApiService) CountOrgWiredClients(ctx context.Context, orgId string, localVarOptionals *OrgsClientsWiredApiCountOrgWiredClientsOpts) (InlineResponse20016, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20016
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{org_id}/wired_clients/count"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", fmt.Sprintf("%v", orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Distinct.IsSet() {
		localVarQueryParams.Add("distinct", parameterToString(localVarOptionals.Distinct.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Duration.IsSet() {
		localVarQueryParams.Add("duration", parameterToString(localVarOptionals.Duration.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-CSRFToken"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20016
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v InlineResponse429
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OrgsClientsWiredApiService searchOrgWiredClients
Search for Wired Clients in org  Note: For list of avaialable &#x60;type&#x60; values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId
 * @param optional nil or *OrgsClientsWiredApiSearchOrgWiredClientsOpts - Optional Parameters:
     * @param "SiteId" (optional.String) -  Site ID
     * @param "DeviceMac" (optional.String) -  device mac
     * @param "Mac" (optional.String) -  client mac
     * @param "PortId" (optional.String) -  port id
     * @param "Vlan" (optional.Int32) -  vlan
     * @param "Ip" (optional.String) -  ip
     * @param "Manufacture" (optional.String) -  client manufacturer
     * @param "Text" (optional.String) -  single entry of hostname/mac
     * @param "NacruleId" (optional.String) -  nacrule_id
     * @param "DhcpHostname" (optional.String) -  DHCP Hostname
     * @param "DhcpFqdn" (optional.String) -  DHCP FQDN
     * @param "DhcpClientIdentifier" (optional.String) -  DHCP Client Identifier
     * @param "DhcpVendorClassIdentifier" (optional.String) -  DHCP Vendor Class Identifier
     * @param "DhcpRequestParams" (optional.String) -  DHCP Request Parameters
     * @param "Limit" (optional.Int32) - 
     * @param "Start" (optional.Int32) -  start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
     * @param "End" (optional.Int32) -  end datetime, can be epoch or relative time like -1d, -2h; now if not specified
     * @param "Duration" (optional.String) -  duration like 7d, 2w
@return InlineResponse20041
*/

type OrgsClientsWiredApiSearchOrgWiredClientsOpts struct {
    SiteId optional.String
    DeviceMac optional.String
    Mac optional.String
    PortId optional.String
    Vlan optional.Int32
    Ip optional.String
    Manufacture optional.String
    Text optional.String
    NacruleId optional.String
    DhcpHostname optional.String
    DhcpFqdn optional.String
    DhcpClientIdentifier optional.String
    DhcpVendorClassIdentifier optional.String
    DhcpRequestParams optional.String
    Limit optional.Int32
    Start optional.Int32
    End optional.Int32
    Duration optional.String
}

func (a *OrgsClientsWiredApiService) SearchOrgWiredClients(ctx context.Context, orgId string, localVarOptionals *OrgsClientsWiredApiSearchOrgWiredClientsOpts) (InlineResponse20041, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20041
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{org_id}/wired_clients/search"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", fmt.Sprintf("%v", orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SiteId.IsSet() {
		localVarQueryParams.Add("site_id", parameterToString(localVarOptionals.SiteId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DeviceMac.IsSet() {
		localVarQueryParams.Add("device_mac", parameterToString(localVarOptionals.DeviceMac.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mac.IsSet() {
		localVarQueryParams.Add("mac", parameterToString(localVarOptionals.Mac.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PortId.IsSet() {
		localVarQueryParams.Add("port_id", parameterToString(localVarOptionals.PortId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Vlan.IsSet() {
		localVarQueryParams.Add("vlan", parameterToString(localVarOptionals.Vlan.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ip.IsSet() {
		localVarQueryParams.Add("ip", parameterToString(localVarOptionals.Ip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Manufacture.IsSet() {
		localVarQueryParams.Add("manufacture", parameterToString(localVarOptionals.Manufacture.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Text.IsSet() {
		localVarQueryParams.Add("text", parameterToString(localVarOptionals.Text.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NacruleId.IsSet() {
		localVarQueryParams.Add("nacrule_id", parameterToString(localVarOptionals.NacruleId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DhcpHostname.IsSet() {
		localVarQueryParams.Add("dhcp_hostname", parameterToString(localVarOptionals.DhcpHostname.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DhcpFqdn.IsSet() {
		localVarQueryParams.Add("dhcp_fqdn", parameterToString(localVarOptionals.DhcpFqdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DhcpClientIdentifier.IsSet() {
		localVarQueryParams.Add("dhcp_client_identifier", parameterToString(localVarOptionals.DhcpClientIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DhcpVendorClassIdentifier.IsSet() {
		localVarQueryParams.Add("dhcp_vendor_class_identifier", parameterToString(localVarOptionals.DhcpVendorClassIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DhcpRequestParams.IsSet() {
		localVarQueryParams.Add("dhcp_request_params", parameterToString(localVarOptionals.DhcpRequestParams.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Duration.IsSet() {
		localVarQueryParams.Add("duration", parameterToString(localVarOptionals.Duration.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-CSRFToken"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20041
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v InlineResponse429
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
