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

type AllOfstatsMxedgeMemoryStat struct {
	// The amount of memory, in kibibytes, that has been used more recently and is usually not reclaimed unless absolutely necessary.
	Active int32 `json:"active,omitempty"`
	// An estimate of how much memory is available for starting new applications, without swapping.
	Available int64 `json:"available,omitempty"`
	// The amount, in kibibytes, of temporary storage for raw disk blocks.
	Buffers int32 `json:"buffers,omitempty"`
	// The amount of physical RAM, in kibibytes, used as cache memory.
	Cached int32 `json:"cached,omitempty"`
	// The amount of physical RAM, in kibibytes, left unused by the system
	Free int64 `json:"free,omitempty"`
	// The amount of memory, in kibibytes, that has been used less recently and is more eligible to be reclaimed for other purposes.
	Inactive int32 `json:"inactive,omitempty"`
	// The amount of memory, in kibibytes, that has once been moved into swap, then back into the main memory, but still also remains in the swapfile.
	SwapCached int32 `json:"swap_cached,omitempty"`
	// The total amount of swap free, in kibibytes.
	SwapFree int32 `json:"swap_free,omitempty"`
	// The total amount of swap available, in kibibytes.
	SwapTotal int32 `json:"swap_total,omitempty"`
	// Total amount of usable RAM, in kibibytes, which is physical RAM minus a number of reserved bits and the kernel binary code
	Total int64 `json:"total,omitempty"`
	Usage int32 `json:"usage,omitempty"`
}
