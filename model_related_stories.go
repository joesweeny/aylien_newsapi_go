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
// RelatedStories struct for RelatedStories
type RelatedStories struct {
	// An array of related stories for the input story
	RelatedStories []Story `json:"related_stories,omitempty"`
	// The input story body
	StoryBody string `json:"story_body,omitempty"`
	// The input story language
	StoryLanguage string `json:"story_language,omitempty"`
	// The input story title
	StoryTitle string `json:"story_title,omitempty"`
}
