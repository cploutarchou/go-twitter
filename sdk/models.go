package sdk

type Tweet struct {
	ID                  string   `json:"id"`
	Text                string   `json:"text"`
	Truncated           bool     `json:"truncated"`
	Entities            Entities `json:"entities"`
	ExtendedEntities    Entities `json:"extended_entities"`
	Source              string   `json:"source"`
	InReplyToStatusID   string   `json:"in_reply_to_status_id"`
	InReplyToUserID     string   `json:"in_reply_to_user_id"`
	InReplyToScreenName string   `json:"in_reply_to_screen_name"`
	User                User     `json:"user"`
	Coordinates         struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"coordinates"`
	Place             Place   `json:"place"`
	Contributors      []int64 `json:"contributors"`
	IsQuoteStatus     bool    `json:"is_quote_status"`
	RetweetCount      int     `json:"retweet_count"`
	FavoriteCount     int     `json:"favorite_count"`
	Favorited         bool    `json:"favorited"`
	Retweeted         bool    `json:"retweeted"`
	PossiblySensitive bool    `json:"possibly_sensitive"`
	Lang              string  `json:"lang"`
	DisplayTextRange  []int   `json:"display_text_range"`
	ExtendedTweet     struct {
		FullText            string   `json:"full_text"`
		DisplayTextRange    []int    `json:"display_text_range"`
		Entities            Entities `json:"entities"`
		ExtendedEntities    Entities `json:"extended_entities"`
		InReplyToStatusID   string   `json:"in_reply_to_status_id"`
		InReplyToUserID     string   `json:"in_reply_to_user_id"`
		InReplyToScreenName string   `json:"in_reply_to_screen_name"`
	} `json:"extended_tweet"`
	QuotedStatusID           string `json:"quoted_status_id"`
	QuotedStatusIDStr        string `json:"quoted_status_id_str"`
	QuotedStatus             *Tweet `json:"quoted_status"`
	QuotedStatusPermalinkURL string `json:"quoted_status_permalink"`
	RetweetedStatus          *Tweet `json:"retweeted_status"`
	CurrentUserRetweet       *struct {
		ID int64 `json:"id"`
	} `json:"current_user_retweet"`
}

type Entities struct {
	Hashtags []Hashtag `json:"hashtags"`
	URLs     []URL     `json:"urls"`
	Mentions []Mention `json:"user_mentions"`
	Media    []Media   `json:"media"`
	Symbols  []Symbol  `json:"symbols"`
	Polls    []Poll    `json:"polls"`
}

type Hashtag struct {
	Indices []int  `json:"indices"`
	Text    string `json:"text"`
}

type URL struct {
	Indices      []int   `json:"indices"`
	URL          string  `json:"url"`
	DisplayURL   string  `json:"display_url"`
	ExpandedURL  string  `json:"expanded_url"`
	UnwoundURL   string  `json:"unwound_url"`
	Description  string  `json:"description"`
	Images       []Image `json:"images"`
	Video        Video   `json:"video"`
	Statuses     []int64 `json:"statuses"`
	Title        string  `json:"title"`
	Audio        Audio   `json:"audio"`
	PreviewImage Image   `json:"preview_image"`
}

type Mention struct {
	Indices    []int  `json:"indices"`
	Name       string `json:"name"`
	ID         int64  `json:"id"`
	IDStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

type Media struct {
	ID              int64     `json:"id"`
	IDStr           string    `json:"id_str"`
	Indices         []int     `json:"indices"`
	MediaURL        string    `json:"media_url"`
	MediaURLHTTPS   string    `json:"media_url_https"`
	URL             string    `json:"url"`
	DisplayURL      string    `json:"display_url"`
	ExpandedURL     string    `json:"expanded_url"`
	Type            string    `json:"type"`
	Sizes           Sizes     `json:"sizes"`
	Features        Features  `json:"features"`
	VideoInfo       VideoInfo `json:"video_info"`
	AdditionalMedia []Media   `json:"additional_media_info"`
}

type Symbol struct {
	Indices []int  `json:"indices"`
	Text    string `json:"text"`
}

type Poll struct {
	ID                 int64        `json:"id"`
	Options            []PollOption `json:"options"`
	DurationMinutes    int          `json:"duration_minutes"`
	EndDatetime        string       `json:"end_datetime"`
	VotingStatus       string       `json:"voting_status"`
	IsExpired          bool         `json:"is_expired"`
	Disclaimer         string       `json:"disclaimer"`
	VotingInstructions string       `json:"voting_instructions"`
	VisibilityStatus   string       `json:"visibility_status"`
}

type PollOption struct {
	ID         int64   `json:"id"`
	Position   int     `json:"position"`
	Label      string  `json:"label"`
	Votes      int     `json:"votes"`
	Percentage float64 `json:"percentage"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Video struct {
	ID            int64  `json:"id"`
	IDStr         string `json:"id_str"`
	DurationMs    int    `json:"duration_ms"`
	AspectRatio   []int  `json:"aspect_ratio"`
	URL           string `json:"url"`
	MediaURL      string `json:"media_url"`
	MediaURLHTTPS string `json:"media_url_https"`
	PreviewImage  Image  `json:"preview_image"`
}

type Audio struct {
	ID          int64  `json:"id"`
	IDStr       string `json:"id_str"`
	DurationMs  int    `json:"duration_ms"`
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
}

type Sizes struct {
	Thumb  Size `json:"thumb"`
	Small  Size `json:"small"`
	Medium Size `json:"medium"`
	Large  Size `json:"large"`
}

type Size struct {
	W      int    `json:"w"`
	H      int    `json:"h"`
	Resize string `json:"resize"`
}

type User struct {
	ID          int64  `json:"id"`
	IDStr       string `json:"id_str"`
	Name        string `json:"name"`
	ScreenName  string `json:"screen_name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Entities    struct {
		URL struct {
			Urls []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"url"`
		Description struct {
			Urls []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	Protected                      bool     `json:"protected"`
	FollowersCount                 int      `json:"followers_count"`
	FriendsCount                   int      `json:"friends_count"`
	ListedCount                    int      `json:"listed_count"`
	CreatedAt                      string   `json:"created_at"`
	FavouritesCount                int      `json:"favourites_count"`
	UtcOffset                      int      `json:"utc_offset"`
	TimeZone                       string   `json:"time_zone"`
	GeoEnabled                     bool     `json:"geo_enabled"`
	Verified                       bool     `json:"verified"`
	StatusesCount                  int      `json:"statuses_count"`
	Lang                           string   `json:"lang"`
	ContributorsEnabled            bool     `json:"contributors_enabled"`
	IsTranslator                   bool     `json:"is_translator"`
	IsTranslationEnabled           bool     `json:"is_translation_enabled"`
	ProfileBackgroundColor         string   `json:"profile_background_color"`
	ProfileBackgroundImageURL      string   `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string   `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool     `json:"profile_background_tile"`
	ProfileImageURL                string   `json:"profile_image_url"`
	ProfileImageURLHTTPS           string   `json:"profile_image_url_https"`
	ProfileBannerURL               string   `json:"profile_banner_url"`
	ProfileLinkColor               string   `json:"profile_link_color"`
	ProfileSidebarBorderColor      string   `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string   `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string   `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool     `json:"profile_use_background_image"`
	DefaultProfile                 bool     `json:"default_profile"`
	DefaultProfileImage            bool     `json:"default_profile_image"`
	WithheldInCountries            []string `json:"withheld_in_countries"`
	WithheldScope                  string   `json:"withheld_scope"`
}

type Place struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	PlaceType   string `json:"place_type"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	BoundingBox struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"bounding_box"`
	Attributes struct {
	} `json:"attributes"`
}

type Features struct {
	Small    *Size `json:"small"`
	Medium   *Size `json:"medium"`
	Large    *Size `json:"large"`
	Orig     *Size `json:"orig"`
	W        int   `json:"w"`
	H        int   `json:"h"`
	Hashtags int   `json:"hashtags"`
	Mentions int   `json:"mentions"`
	URLs     int   `json:"urls"`
	Media    int   `json:"media"`
}

type VideoInfo struct {
	AspectRatio    []int     `json:"aspect_ratio"`
	DurationMillis int64     `json:"duration_millis"`
	Variants       []Variant `json:"variants"`
}

type Variant struct {
	Bitrate     int64  `json:"bitrate,omitempty"`
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
}
