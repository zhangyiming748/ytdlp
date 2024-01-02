package sql

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type PH struct {
	Id           string `json:"id"`
	Uploader     string `json:"uploader"`
	UploaderId   string `json:"uploader_id"`
	UploadDate   string `json:"upload_date"`
	Title        string `json:"title"`
	Thumbnail    string `json:"thumbnail"`
	Duration     int    `json:"duration"`
	LikeCount    int    `json:"like_count"`
	DislikeCount int    `json:"dislike_count"`
	CommentCount int    `json:"comment_count"`
	Formats      []struct {
		Url          string `json:"url"`
		FormatId     string `json:"format_id"`
		Height       int    `json:"height"`
		Ext          string `json:"ext"`
		Protocol     string `json:"protocol"`
		Resolution   string `json:"resolution"`
		DynamicRange string `json:"dynamic_range"`
		HttpHeaders  struct {
			UserAgent      string `json:"User-Agent"`
			Accept         string `json:"Accept"`
			AcceptLanguage string `json:"Accept-Language"`
			SecFetchMode   string `json:"Sec-Fetch-Mode"`
		} `json:"http_headers"`
		VideoExt    string  `json:"video_ext"`
		AudioExt    string  `json:"audio_ext"`
		Format      string  `json:"format"`
		ManifestUrl string  `json:"manifest_url,omitempty"`
		Tbr         float64 `json:"tbr,omitempty"`
		Fps         float64 `json:"fps,omitempty"`
		HasDrm      bool    `json:"has_drm,omitempty"`
		Width       int     `json:"width,omitempty"`
		Vcodec      string  `json:"vcodec,omitempty"`
		Acodec      string  `json:"acodec,omitempty"`
		AspectRatio float64 `json:"aspect_ratio,omitempty"`
	} `json:"formats"`
	AgeLimit   int      `json:"age_limit"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
	Cast       []string `json:"cast"`
	Subtitles  struct {
	} `json:"subtitles"`
	Thumbnails []struct {
		Url string `json:"url"`
		Id  string `json:"id"`
	} `json:"thumbnails"`
	Timestamp          int     `json:"timestamp"`
	ViewCount          int     `json:"view_count"`
	WebpageUrl         string  `json:"webpage_url"`
	WebpageUrlBasename string  `json:"webpage_url_basename"`
	WebpageUrlDomain   string  `json:"webpage_url_domain"`
	Extractor          string  `json:"extractor"`
	ExtractorKey       string  `json:"extractor_key"`
	DisplayId          string  `json:"display_id"`
	Fulltitle          string  `json:"fulltitle"`
	DurationString     string  `json:"duration_string"`
	Epoch              int     `json:"epoch"`
	FormatId           string  `json:"format_id"`
	Url                string  `json:"url"`
	ManifestUrl        string  `json:"manifest_url"`
	Tbr                float64 `json:"tbr"`
	Ext                string  `json:"ext"`
	Fps                float64 `json:"fps"`
	Protocol           string  `json:"protocol"`
	HasDrm             bool    `json:"has_drm"`
	Width              int     `json:"width"`
	Height             int     `json:"height"`
	Vcodec             string  `json:"vcodec"`
	Acodec             string  `json:"acodec"`
	DynamicRange       string  `json:"dynamic_range"`
	Resolution         string  `json:"resolution"`
	AspectRatio        float64 `json:"aspect_ratio"`
	HttpHeaders        struct {
		UserAgent      string `json:"User-Agent"`
		Accept         string `json:"Accept"`
		AcceptLanguage string `json:"Accept-Language"`
		SecFetchMode   string `json:"Sec-Fetch-Mode"`
	} `json:"http_headers"`
	VideoExt string `json:"video_ext"`
	AudioExt string `json:"audio_ext"`
	Format   string `json:"format"`
	Type     string `json:"_type"`
	Version  struct {
		Version        string `json:"version"`
		ReleaseGitHead string `json:"release_git_head"`
		Repository     string `json:"repository"`
	} `json:"_version"`
}
type Info struct {
	gorm.Model
	ID        uint        `gorm:"primaryKey"`
	Title     string      `gorm:"title"`
	JSON      interface{} `gorm:"js"`
	CreatedAt time.Time
}

func (i *Info) FindOneByTitle() *gorm.DB {
	return GetEngine().Where("title = ?", i.Title).First(&i)
}
func (i *Info) SetOne() *gorm.DB {
	return GetEngine().Create(&i)
}

func (i *Info) FindDupByName() int {
	var is []Info
	GetEngine().Table("ytdlps").Where("title = ?", i.Title).Find(&is)
	fmt.Println(is)
	return len(is)
}
