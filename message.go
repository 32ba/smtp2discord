package main

type DiscordWebhook struct {
	Username string       `json:"username,omitempty"`
	Avatar   string       `json:"avatar_url,omitempty"`
	Content  string       `json:"content,omitempty"`
	Embeds   DiscordEmbed `json:"embeds,omitempty"`
}

type DiscordEmbed []struct {
	Title       string                `json:"title"`
	Type        string                `json:"type,omitempty"`
	Description string                `json:"description,omitempty"`
	URL         string                `json:"url,omitempty"`
	Timestamp   string                `json:"timestamp,omitempty"`
	Color       int                   `json:"color,omitempty"`
	Footer      DiscordEmbedFooter    `json:"footer,omitempty"`
	Image       DiscordEmbedImage     `json:"image,omitempty"`
	Thumbnail   DiscordEmbedThumbnail `json:"thumbnail,omitempty"`
	Video       DiscordEmbedVideo     `json:"video,omitempty"`
	Provider    DiscordEmbedProvider  `json:"provider,omitempty"`
	Author      DiscordEmbedAuthor    `json:"author,omitempty"`
	Fields      []DiscordEmbedField   `json:"fields,omitempty"`
}

type DiscordEmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type DiscordEmbedImage struct {
	URL          string `json:"url"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
}

type DiscordEmbedThumbnail struct {
	URL          string `json:"url"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
}

type DiscordEmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type DiscordEmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type DiscordEmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
