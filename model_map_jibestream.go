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

type MapJibestream struct {
	// the client id
	ClientId string `json:"client_id"`
	// the client secret
	ClientSecret string `json:"client_secret"`
	// the jibestream customer record id
	CustomerId int32 `json:"customer_id"`
	// the map contents endpoint host
	EndpointUrl string `json:"endpoint_url"`
	// the jibestream map id
	MapId string `json:"map_id"`
	// millimeter per pixel
	Mmpp int32 `json:"mmpp"`
	// pixel per meter, same as the map JSON value.
	Ppm float64 `json:"ppm"`
	// the vendor ‘jibestream’. enum: `jibestream`
	VendorName string `json:"vendor_name"`
	// the venue or organization id
	VenueId int32 `json:"venue_id"`
}
