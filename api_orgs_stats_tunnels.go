
/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

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

type OrgsStatsTunnelsApiService service
/*
OrgsStatsTunnelsApiService countOrgTunnelsStats
Count Mist Tunnels Stats
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId
 * @param optional nil or *OrgsStatsTunnelsApiCountOrgTunnelsStatsOpts - Optional Parameters:
     * @param "Distinct" (optional.Interface of Distinct1) -  - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;: wxtunnel_id / ap / remote_ip / remote_port / state / mxedge_id / mxcluster_id / site_id / peer_mxedge_id; default is wxtunnel_id  - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;: mac / site_id / node / peer_ip / peer_host/ ip / tunnel_name / protocol / auth_algo / encrypt_algo / ike_version / last_event / up
     * @param "Type_" (optional.Interface of OrgTunnelTypeCount) - 
@return InlineResponse20016
*/

type OrgsStatsTunnelsApiCountOrgTunnelsStatsOpts struct {
    Distinct optional.Interface
    Type_ optional.Interface
}

func (a *OrgsStatsTunnelsApiService) CountOrgTunnelsStats(ctx context.Context, orgId string, localVarOptionals *OrgsStatsTunnelsApiCountOrgTunnelsStatsOpts) (InlineResponse20016, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20016
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{org_id}/stats/tunnels/count"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", fmt.Sprintf("%v", orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Distinct.IsSet() {
		localVarQueryParams.Add("distinct", parameterToString(localVarOptionals.Distinct.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
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
OrgsStatsTunnelsApiService searchOrgTunnelsStats
Search Org Tunnels Stats
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId
 * @param optional nil or *OrgsStatsTunnelsApiSearchOrgTunnelsStatsOpts - Optional Parameters:
     * @param "MxclusterId" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;
     * @param "SiteId" (optional.String) - 
     * @param "WxtunnelId" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;
     * @param "Ap" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;
     * @param "Mac" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "Node" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "PeerIp" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "PeerHost" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "Ip" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "TunnelName" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "Protocol" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "AuthAlgo" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "EncryptAlgo" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "IkeVersion" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "Up" (optional.String) -  if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
     * @param "Type_" (optional.Interface of TunnelType) - 
     * @param "Limit" (optional.Int32) - 
     * @param "Start" (optional.Int32) -  start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
     * @param "End" (optional.Int32) -  end datetime, can be epoch or relative time like -1d, -2h; now if not specified
     * @param "Duration" (optional.String) -  duration like 7d, 2w
@return InlineResponse20079
*/

type OrgsStatsTunnelsApiSearchOrgTunnelsStatsOpts struct {
    MxclusterId optional.String
    SiteId optional.String
    WxtunnelId optional.String
    Ap optional.String
    Mac optional.String
    Node optional.String
    PeerIp optional.String
    PeerHost optional.String
    Ip optional.String
    TunnelName optional.String
    Protocol optional.String
    AuthAlgo optional.String
    EncryptAlgo optional.String
    IkeVersion optional.String
    Up optional.String
    Type_ optional.Interface
    Limit optional.Int32
    Start optional.Int32
    End optional.Int32
    Duration optional.String
}

func (a *OrgsStatsTunnelsApiService) SearchOrgTunnelsStats(ctx context.Context, orgId string, localVarOptionals *OrgsStatsTunnelsApiSearchOrgTunnelsStatsOpts) (InlineResponse20079, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20079
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{org_id}/stats/tunnels/search"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", fmt.Sprintf("%v", orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.MxclusterId.IsSet() {
		localVarQueryParams.Add("mxcluster_id", parameterToString(localVarOptionals.MxclusterId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SiteId.IsSet() {
		localVarQueryParams.Add("site_id", parameterToString(localVarOptionals.SiteId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WxtunnelId.IsSet() {
		localVarQueryParams.Add("wxtunnel_id", parameterToString(localVarOptionals.WxtunnelId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ap.IsSet() {
		localVarQueryParams.Add("ap", parameterToString(localVarOptionals.Ap.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mac.IsSet() {
		localVarQueryParams.Add("mac", parameterToString(localVarOptionals.Mac.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Node.IsSet() {
		localVarQueryParams.Add("node", parameterToString(localVarOptionals.Node.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PeerIp.IsSet() {
		localVarQueryParams.Add("peer_ip", parameterToString(localVarOptionals.PeerIp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PeerHost.IsSet() {
		localVarQueryParams.Add("peer_host", parameterToString(localVarOptionals.PeerHost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ip.IsSet() {
		localVarQueryParams.Add("ip", parameterToString(localVarOptionals.Ip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TunnelName.IsSet() {
		localVarQueryParams.Add("tunnel_name", parameterToString(localVarOptionals.TunnelName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Protocol.IsSet() {
		localVarQueryParams.Add("protocol", parameterToString(localVarOptionals.Protocol.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AuthAlgo.IsSet() {
		localVarQueryParams.Add("auth_algo", parameterToString(localVarOptionals.AuthAlgo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EncryptAlgo.IsSet() {
		localVarQueryParams.Add("encrypt_algo", parameterToString(localVarOptionals.EncryptAlgo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IkeVersion.IsSet() {
		localVarQueryParams.Add("ike_version", parameterToString(localVarOptionals.IkeVersion.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Up.IsSet() {
		localVarQueryParams.Add("up", parameterToString(localVarOptionals.Up.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
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
			var v InlineResponse20079
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
