package reddit

import (
	"time"
)

// RedditResponse represents the top-level Reddit API response
type RedditResponse struct {
	Kind string      `json:"kind"`
	Data ListingData `json:"data"`
}

// ListingData represents the data field of a Reddit listing
type ListingData struct {
	After     string  `json:"after"`
	Dist      int     `json:"dist"`
	Modhash   string  `json:"modhash"`
	GeoFilter string  `json:"geo_filter"`
	Children  []Child `json:"children"`
	Before    string  `json:"before"`
}

// Child represents individual posts in the listing
type Child struct {
	Kind string   `json:"kind"`
	Data PostData `json:"data"`
}

// PostData represents the data of a Reddit post
type PostData struct {
	ApprovedAtUTC         *int64                 `json:"approved_at_utc"`
	Subreddit             string                 `json:"subreddit"`
	Selftext              string                 `json:"selftext"`
	AuthorFullname        string                 `json:"author_fullname"`
	Saved                 bool                   `json:"saved"`
	ModReasonTitle        *string                `json:"mod_reason_title"`
	Gilded                int                    `json:"gilded"`
	Clicked               bool                   `json:"clicked"`
	Title                 string                 `json:"title"`
	LinkFlairRichtext     []FlairRichtext        `json:"link_flair_richtext"`
	SubredditNamePrefixed string                 `json:"subreddit_name_prefixed"`
	Hidden                bool                   `json:"hidden"`
	Pwls                  int                    `json:"pwls"`
	LinkFlairCssClass     *string                `json:"link_flair_css_class"`
	Downs                 int                    `json:"downs"`
	ThumbnailHeight       *int                   `json:"thumbnail_height"`
	TopAwardedType        *string                `json:"top_awarded_type"`
	HideScore             bool                   `json:"hide_score"`
	Name                  string                 `json:"name"`
	Quarantine            bool                   `json:"quarantine"`
	LinkFlairTextColor    string                 `json:"link_flair_text_color"`
	UpvoteRatio           float64                `json:"upvote_ratio"`
	AuthorFlairBgColor    *string                `json:"author_flair_background_color"`
	SubredditType         string                 `json:"subreddit_type"`
	Ups                   int                    `json:"ups"`
	TotalAwardsReceived   int                    `json:"total_awards_received"`
	MediaEmbed            map[string]interface{} `json:"media_embed"`
	ThumbnailWidth        *int                   `json:"thumbnail_width"`
	AuthorFlairTemplateID *string                `json:"author_flair_template_id"`
	IsOriginalContent     bool                   `json:"is_original_content"`
	UserReports           []interface{}          `json:"user_reports"`
	SecureMedia           *interface{}           `json:"secure_media"`
	IsRedditMediaDomain   bool                   `json:"is_reddit_media_domain"`
	IsMeta                bool                   `json:"is_meta"`
	Category              *string                `json:"category"`
	SecureMediaEmbed      map[string]interface{} `json:"secure_media_embed"`
	LinkFlairText         *string                `json:"link_flair_text"`
	CanModPost            bool                   `json:"can_mod_post"`
	Score                 int                    `json:"score"`
	ApprovedBy            *string                `json:"approved_by"`
	IsCreatedFromAdsUI    bool                   `json:"is_created_from_ads_ui"`
	AuthorPremium         bool                   `json:"author_premium"`
	Thumbnail             string                 `json:"thumbnail"`
	Edited                interface{}            `json:"edited"` // Can be bool or timestamp
	AuthorFlairCssClass   *string                `json:"author_flair_css_class"`
	AuthorFlairRichtext   []interface{}          `json:"author_flair_richtext"`
	Gildings              map[string]interface{} `json:"gildings"`
	PostHint              *string                `json:"post_hint"`
	ContentCategories     *interface{}           `json:"content_categories"`
	IsSelf                bool                   `json:"is_self"`
	ModNote               *string                `json:"mod_note"`
	Created               float64                `json:"created"`
	LinkFlairType         string                 `json:"link_flair_type"`
	Wls                   int                    `json:"wls"`
	RemovedByCategory     *string                `json:"removed_by_category"`
	BannedBy              *string                `json:"banned_by"`
	AuthorFlairType       string                 `json:"author_flair_type"`
	Domain                string                 `json:"domain"`
	AllowLiveComments     bool                   `json:"allow_live_comments"`
	SelftextHTML          *string                `json:"selftext_html"`
	Likes                 *interface{}           `json:"likes"`
	SuggestedSort         string                 `json:"suggested_sort"`
	BannedAtUTC           *int64                 `json:"banned_at_utc"`
	URLOverriddenByDest   *string                `json:"url_overridden_by_dest"`
	ViewCount             *int                   `json:"view_count"`
	Archived              bool                   `json:"archived"`
	NoFollow              bool                   `json:"no_follow"`
	IsCrosspostable       bool                   `json:"is_crosspostable"`
	Pinned                bool                   `json:"pinned"`
	Over18                bool                   `json:"over_18"`
	Preview               *Preview               `json:"preview"`
	AllAwardings          []interface{}          `json:"all_awardings"`
	Awarders              []interface{}          `json:"awarders"`
	MediaOnly             bool                   `json:"media_only"`
	LinkFlairTemplateID   *string                `json:"link_flair_template_id"`
	CanGild               bool                   `json:"can_gild"`
	Spoiler               bool                   `json:"spoiler"`
	Locked                bool                   `json:"locked"`
	AuthorFlairText       *string                `json:"author_flair_text"`
	TreatmentTags         []interface{}          `json:"treatment_tags"`
	Visited               bool                   `json:"visited"`
	RemovedBy             *string                `json:"removed_by"`
	NumReports            *int                   `json:"num_reports"`
	Distinguished         *string                `json:"distinguished"`
	SubredditID           string                 `json:"subreddit_id"`
	AuthorIsBlocked       bool                   `json:"author_is_blocked"`
	ModReasonBy           *string                `json:"mod_reason_by"`
	RemovalReason         *string                `json:"removal_reason"`
	LinkFlairBgColor      *string                `json:"link_flair_background_color"`
	ID                    string                 `json:"id"`
	IsRobotIndexable      bool                   `json:"is_robot_indexable"`
	ReportReasons         *interface{}           `json:"report_reasons"`
	Author                string                 `json:"author"`
	DiscussionType        *string                `json:"discussion_type"`
	NumComments           int                    `json:"num_comments"`
	SendReplies           bool                   `json:"send_replies"`
	ContestMode           bool                   `json:"contest_mode"`
	ModReports            []interface{}          `json:"mod_reports"`
	AuthorPatreonFlair    bool                   `json:"author_patreon_flair"`
	AuthorFlairTextColor  *string                `json:"author_flair_text_color"`
	Permalink             string                 `json:"permalink"`
	Stickied              bool                   `json:"stickied"`
	URL                   string                 `json:"url"`
	SubredditSubscribers  int                    `json:"subreddit_subscribers"`
	CreatedUTC            float64                `json:"created_utc"`
	NumCrossposts         int                    `json:"num_crossposts"`
	Media                 *interface{}           `json:"media"`
	IsVideo               bool                   `json:"is_video"`

	// Gallery-specific fields
	IsGallery     *bool          `json:"is_gallery,omitempty"`
	MediaMetadata *MediaMetadata `json:"media_metadata,omitempty"`
	GalleryData   *GalleryData   `json:"gallery_data,omitempty"`
}

// FlairRichtext represents flair richtext elements
type FlairRichtext struct {
	E string `json:"e"`
	T string `json:"t"`
}

// Preview represents preview images for posts
type Preview struct {
	Images  []PreviewImage `json:"images"`
	Enabled bool           `json:"enabled"`
}

// PreviewImage represents individual preview images
type PreviewImage struct {
	Source      ImageSource            `json:"source"`
	Resolutions []ImageSource          `json:"resolutions"`
	Variants    map[string]interface{} `json:"variants"`
	ID          string                 `json:"id"`
}

// ImageSource represents image source data
type ImageSource struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// MediaMetadata represents metadata for gallery images
type MediaMetadata map[string]MediaMetadataItem

// MediaMetadataItem represents individual media metadata
type MediaMetadataItem struct {
	Status string        `json:"status"`
	E      string        `json:"e"`
	M      string        `json:"m"`
	P      []ImageSource `json:"p"`
	S      ImageSource   `json:"s"`
	ID     string        `json:"id"`
}

// GalleryData represents gallery-specific data
type GalleryData struct {
	Items []GalleryItem `json:"items"`
}

// GalleryItem represents individual gallery items
type GalleryItem struct {
	MediaID string `json:"media_id"`
	ID      int    `json:"id"`
}

// Helper methods

// GetCreatedTime returns the created time as a time.Time
func (p *PostData) GetCreatedTime() time.Time {
	return time.Unix(int64(p.Created), 0)
}

// GetCreatedUTCTime returns the created UTC time as a time.Time
func (p *PostData) GetCreatedUTCTime() time.Time {
	return time.Unix(int64(p.CreatedUTC), 0)
}

// IsTextPost returns true if this is a self/text post
func (p *PostData) IsTextPost() bool {
	return p.IsSelf
}

// IsImagePost returns true if this is an image post
func (p *PostData) IsImagePost() bool {
	return p.PostHint != nil && *p.PostHint == "image"
}

// IsGalleryPost returns true if this is a gallery post
func (p *PostData) IsGalleryPost() bool {
	return p.IsGallery != nil && *p.IsGallery
}
