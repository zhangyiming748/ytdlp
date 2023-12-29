package main

type Info struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Formats []struct {
		FormatId   string  `json:"format_id"`
		FormatNote string  `json:"format_note,omitempty"`
		Ext        string  `json:"ext"`
		Protocol   string  `json:"protocol"`
		Acodec     string  `json:"acodec,omitempty"`
		Vcodec     string  `json:"vcodec"`
		Url        string  `json:"url"`
		Width      int     `json:"width,omitempty"`
		Height     int     `json:"height,omitempty"`
		Fps        float64 `json:"fps,omitempty"`
		Rows       int     `json:"rows,omitempty"`
		Columns    int     `json:"columns,omitempty"`
		Fragments  []struct {
			Url      string  `json:"url"`
			Duration float64 `json:"duration"`
		} `json:"fragments,omitempty"`
		Resolution  string  `json:"resolution"`
		AspectRatio float64 `json:"aspect_ratio,omitempty"`
		HttpHeaders struct {
			UserAgent      string `json:"User-Agent"`
			Accept         string `json:"Accept"`
			AcceptLanguage string `json:"Accept-Language"`
			SecFetchMode   string `json:"Sec-Fetch-Mode"`
		} `json:"http_headers"`
		AudioExt           string  `json:"audio_ext"`
		VideoExt           string  `json:"video_ext"`
		Vbr                float64 `json:"vbr,omitempty"`
		Abr                float64 `json:"abr,omitempty"`
		Format             string  `json:"format"`
		ManifestUrl        string  `json:"manifest_url,omitempty"`
		Quality            float64 `json:"quality,omitempty"`
		HasDrm             bool    `json:"has_drm,omitempty"`
		SourcePreference   int     `json:"source_preference,omitempty"`
		Asr                int     `json:"asr,omitempty"`
		Filesize           int     `json:"filesize,omitempty"`
		AudioChannels      int     `json:"audio_channels,omitempty"`
		Tbr                float64 `json:"tbr,omitempty"`
		LanguagePreference int     `json:"language_preference,omitempty"`
		Container          string  `json:"container,omitempty"`
		DownloaderOptions  struct {
			HttpChunkSize int `json:"http_chunk_size"`
		} `json:"downloader_options,omitempty"`
		DynamicRange   string `json:"dynamic_range,omitempty"`
		FilesizeApprox int    `json:"filesize_approx,omitempty"`
	} `json:"formats"`
	Thumbnails []struct {
		Url        string `json:"url"`
		Preference int    `json:"preference"`
		Id         string `json:"id"`
		Height     int    `json:"height,omitempty"`
		Width      int    `json:"width,omitempty"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"thumbnails"`
	Thumbnail         string   `json:"thumbnail"`
	Description       string   `json:"description"`
	ChannelId         string   `json:"channel_id"`
	ChannelUrl        string   `json:"channel_url"`
	Duration          int      `json:"duration"`
	ViewCount         int      `json:"view_count"`
	AgeLimit          int      `json:"age_limit"`
	WebpageUrl        string   `json:"webpage_url"`
	Categories        []string `json:"categories"`
	Tags              []string `json:"tags"`
	PlayableInEmbed   bool     `json:"playable_in_embed"`
	LiveStatus        string   `json:"live_status"`
	FormatSortFields  []string `json:"_format_sort_fields"`
	AutomaticCaptions struct {
	} `json:"automatic_captions"`
	Subtitles struct {
	} `json:"subtitles"`
	CommentCount         int     `json:"comment_count"`
	Channel              string  `json:"channel"`
	ChannelFollowerCount int     `json:"channel_follower_count"`
	Uploader             string  `json:"uploader"`
	UploaderId           string  `json:"uploader_id"`
	UploaderUrl          string  `json:"uploader_url"`
	UploadDate           string  `json:"upload_date"`
	Availability         string  `json:"availability"`
	WebpageUrlBasename   string  `json:"webpage_url_basename"`
	WebpageUrlDomain     string  `json:"webpage_url_domain"`
	Extractor            string  `json:"extractor"`
	ExtractorKey         string  `json:"extractor_key"`
	DisplayId            string  `json:"display_id"`
	Fulltitle            string  `json:"fulltitle"`
	DurationString       string  `json:"duration_string"`
	IsLive               bool    `json:"is_live"`
	WasLive              bool    `json:"was_live"`
	Epoch                int     `json:"epoch"`
	Format               string  `json:"format"`
	FormatId             string  `json:"format_id"`
	Ext                  string  `json:"ext"`
	Protocol             string  `json:"protocol"`
	FormatNote           string  `json:"format_note"`
	FilesizeApprox       int     `json:"filesize_approx"`
	Tbr                  float64 `json:"tbr"`
	Width                int     `json:"width"`
	Height               int     `json:"height"`
	Resolution           string  `json:"resolution"`
	Fps                  int     `json:"fps"`
	DynamicRange         string  `json:"dynamic_range"`
	Vcodec               string  `json:"vcodec"`
	Vbr                  float64 `json:"vbr"`
	AspectRatio          float64 `json:"aspect_ratio"`
	Acodec               string  `json:"acodec"`
	Abr                  float64 `json:"abr"`
	Asr                  int     `json:"asr"`
	AudioChannels        int     `json:"audio_channels"`
	Type                 string  `json:"_type"`
	Version              struct {
		Version        string `json:"version"`
		ReleaseGitHead string `json:"release_git_head"`
		Repository     string `json:"repository"`
	} `json:"_version"`
}
