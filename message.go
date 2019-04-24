package dingbot

// exported message types
const (
	MsgTypeText       = "text"
	MsgTypeLink       = "link"
	MsgTypeMd         = "markdown"
	MsgTypeActionCard = "actionCard"
	MsgTypeFeedCard   = "feedCard"
)

// TextMsg represents text.
type TextMsg struct {
	Content string `json:"content"`
}

// LinkMsg represents a link.
type LinkMsg struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl,omitempty"`
	MessageURL string `json:"messageUrl"`
}

// MarkdownMsg represents markdown formatted text.
type MarkdownMsg struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

// ActionCardMsg represents a action card.
type ActionCardMsg struct {
	Text           string `json:"text"`
	Title          string `json:"title"`
	HideAvatar     string `json:"hideAvatar"`     // "0" or "1"
	BtnOrientation string `json:"btnOrientation"` // "0" or "1"
	// single card
	SingleTitle string `json:"singleTitle"`
	SingleURL   string `json:"singleURL"`
	// independent cards
	Btns []struct {
		Title     string `json:"title"`
		ActionURL string `json:"actionURL"`
	} `json:"btns"`
}

// FeedLink represents a feed.
type FeedLink struct {
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

// FeedCardMsg represents a feed card.
type FeedCardMsg struct {
	Links []FeedLink `json:"links"`
}

// AtOption is the option for @someone.
type AtOption struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// DingMessage is the JSON-encoded message to send.
type DingMessage struct {
	Msgtype    string         `json:"msgtype"`
	Text       *TextMsg       `json:"text,omitempty"`
	Link       *LinkMsg       `json:"link,omitempty"`
	Markdown   *MarkdownMsg   `json:"markdown,omitempty"`
	ActionCard *ActionCardMsg `json:"actionCard,omitempty"`
	FeedCard   *FeedCardMsg   `json:"feedCard,omitempty"`
	At         *AtOption      `json:"at,omitempty"`
}
