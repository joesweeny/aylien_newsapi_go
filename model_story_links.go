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
// StoryLinks struct for StoryLinks
type StoryLinks struct {
	// The story canonical URL
	Canonical string `json:"canonical,omitempty"`
	// The coverages URL
	Coverages string `json:"coverages,omitempty"`
	// The story permalink URL
	Permalink string `json:"permalink,omitempty"`
	// The related stories URL
	RelatedStories string `json:"related_stories,omitempty"`
}
