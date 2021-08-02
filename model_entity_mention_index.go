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
// EntityMentionIndex struct for EntityMentionIndex
type EntityMentionIndex struct {
	// End index of a single entity mention in the text (counting from 0)
	End int32 `json:"end"`
	// Start index of a single entity mention in the text (counting from 0)
	Start int32 `json:"start"`
}
