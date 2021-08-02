/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 3.0
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// DeprecatedEntities struct for DeprecatedEntities
type DeprecatedEntities struct {
	// An array of extracted entities from the story body
	Body []DeprecatedEntity `json:"body,omitempty"`
	// An array of extracted entities from the story title
	Title []DeprecatedEntity `json:"title,omitempty"`
}
