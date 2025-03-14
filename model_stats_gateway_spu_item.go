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

type StatsGatewaySpuItem struct {
	SpuCpu int32 `json:"spu_cpu,omitempty"`
	SpuCurrentSession int32 `json:"spu_current_session,omitempty"`
	SpuMaxSession int32 `json:"spu_max_session,omitempty"`
	SpuMemory int32 `json:"spu_memory,omitempty"`
	SpuPendingSession int32 `json:"spu_pending_session,omitempty"`
	SpuValidSession int32 `json:"spu_valid_session,omitempty"`
}
