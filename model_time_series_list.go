/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 5.0.1
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
import (
	"time"
)
// TimeSeriesList struct for TimeSeriesList
type TimeSeriesList struct {
	// The size of each date range expressed as an interval to be added to the lower bound. 
	Period string `json:"period,omitempty"`
	// The end published date of the time series
	PublishedAtEnd time.Time `json:"published_at.end,omitempty"`
	// The start published date of the time series
	PublishedAtStart time.Time `json:"published_at.start,omitempty"`
	// An array of time series
	TimeSeries []TimeSeries `json:"time_series,omitempty"`
}
