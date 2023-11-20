package feeds

// rss support
// validation done according to spec here:
//    http://cyber.law.harvard.edu/rss/rss.html

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

const googleAdditionalImageSeparator = ","

// private wrapper around the RssFeed which gives us the <rss>..</rss> xml
type RssFeedXml struct {
	XMLName                xml.Name `xml:"rss"`
	Version                string   `xml:"version,attr"`
	ContentNamespace       string   `xml:"xmlns:content,attr,omitempty"`
	GoogleContentNamespace string   `xml:"xmlns:g,attr,omitempty"`
	MediaContentNamespace  string   `xml:"xmlns:media,attr,omitempty"`

	Channel *RssFeed
}

type RssContent struct {
	XMLName xml.Name `xml:"content:encoded"`
	Content string   `xml:",cdata"`
}

type RssImage struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   int      `xml:"width,omitempty"`
	Height  int      `xml:"height,omitempty"`
}

type RssTextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}

type RssFeed struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`       // required
	Link           string   `xml:"link"`        // required
	Description    string   `xml:"description"` // required
	Language       string   `xml:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"` // Author used
	WebMaster      string   `xml:"webMaster,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"`       // created or updated
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"` // updated used
	Category       string   `xml:"category,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Docs           string   `xml:"docs,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"`
	Ttl            int      `xml:"ttl,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"`
	Image          *RssImage
	TextInput      *RssTextInput
	Items          []*RssItem `xml:"item"`
	Atom           *ChannelAtom
}

type ChannelAtom struct {
	XMLName xml.Name `xml:"atom"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title,omitempty"`
	Link        string   `xml:"link,omitempty"`
	Description string   `xml:"description,omitempty"`
	Content     *RssContent
	Author      string `xml:"author,omitempty"`
	Category    string `xml:"category,omitempty"`
	Comments    string `xml:"comments,omitempty"`
	Enclosure   *RssEnclosure
	Guid        string `xml:"guid,omitempty"`    // Id used
	PubDate     string `xml:"pubDate,omitempty"` // created or updated
	Source      string `xml:"source,omitempty"`

	MediaContent *MediaContent
	// Google Merchant Center
	GoogleId              string   `xml:"g:id,omitempty"`
	GoogleTitle           string   `xml:"g:title,omitempty"`
	GoogleDesc            string   `xml:"g:description,omitempty"`
	GoogleLink            string   `xml:"g:link,omitempty"`
	GoogleCond            string   `xml:"g:condition,omitempty"`
	GooglePrice           string   `xml:"g:price,omitempty"`
	GoogleSale            string   `xml:"g:sale_price,omitempty"`
	GoogleAvail           string   `xml:"g:availability,omitempty"`
	GoogleImage           string   `xml:"g:image_link,omitempty"`
	GoogleAdditionalImage []string `xml:"g:additional_image_link,omitempty"`
	GoogleGtin            string   `xml:"g:gtin,omitempty"`
	GoogleMpn             string   `xml:"g:mpn,omitempty"`
	GoogleBrand           string   `xml:"g:brand,omitempty"`
	GoogleCat             string   `xml:"g:google_product_category,omitempty"`
	GoogleShip            string   `xml:"g:shipping,omitempty"`
	GoogleInv             string   `xml:"g:inventory,omitempty"`
	GoogleColor           string   `xml:"g:color,omitempty"`
	GoogleType            string   `xml:"g:product_type,omitempty"`
	GoogleLabel0          string   `xml:"g:custom_label_0,omitempty"`
	GoogleLabel1          string   `xml:"g:custom_label_1,omitempty"`
	GoogleLabel2          string   `xml:"g:custom_label_2,omitempty"`
	GoogleLabel3          string   `xml:"g:custom_label_3,omitempty"`
	GoogleLabel4          string   `xml:"g:custom_label_4,omitempty"`
	GoogleGroup           string   `xml:"g:item_group_id,omitempty"`
	GooglePromotionId     string   `xml:"g:promotion_id,omitempty"`
}

type MediaContent struct {
	XMLName xml.Name `xml:"media:content"`
	Url     string   `xml:"url,attr"`
}

type RssEnclosure struct {
	// RSS 2.0 <enclosure url="http://example.com/file.mp3" length="123456789" type="audio/mpeg" />
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type Rss struct {
	*Feed
}

// create a new RssItem with a generic Item struct's data
func newRssItem(i *Item) *RssItem {
	item := &RssItem{
		Title:       i.Title,
		Description: i.Description,
		Guid:        i.Id,
		PubDate:     i.PubDate,

		MediaContent:          i.MediaContent,
		GoogleId:              i.GoogleId,
		GoogleTitle:           i.GoogleTitle,
		GoogleDesc:            i.GoogleDesc,
		GoogleLink:            i.GoogleLink,
		GoogleCond:            i.GoogleCond,
		GooglePrice:           i.GooglePrice,
		GoogleSale:            i.GoogleSale,
		GoogleAvail:           i.GoogleAvail,
		GoogleImage:           i.GoogleImage,
		GoogleAdditionalImage: strings.Split(i.GoogleAdditionalImage, googleAdditionalImageSeparator),
		GoogleGtin:            i.GoogleGtin,
		GoogleMpn:             i.GoogleMpn,
		GoogleBrand:           i.GoogleBrand,
		GoogleCat:             i.GoogleCat,
		GoogleShip:            i.GoogleShip,
		GoogleInv:             i.GoogleInv,
		GoogleColor:           i.GoogleColor,
		GoogleType:            i.GoogleType,
		GoogleLabel0:          i.GoogleLabel0,
		GoogleLabel1:          i.GoogleLabel1,
		GoogleLabel2:          i.GoogleLabel2,
		GoogleLabel3:          i.GoogleLabel3,
		GoogleLabel4:          i.GoogleLabel4,
		GoogleGroup:           i.GoogleGroup,
		GooglePromotionId:     i.GooglePromotionId,
	}
	if i.Link != nil {
		item.Link = i.Link.Href
	}
	if len(i.Content) > 0 {
		item.Content = &RssContent{Content: i.Content}
	}
	if i.Source != nil {
		item.Source = i.Source.Href
	}

	// Define a closure
	if i.Enclosure != nil && i.Enclosure.Type != "" && i.Enclosure.Length != "" {
		item.Enclosure = &RssEnclosure{Url: i.Enclosure.Url, Type: i.Enclosure.Type, Length: i.Enclosure.Length}
	}

	if i.Author != nil {
		item.Author = i.Author.Name
	}
	return item
}

// create a new RssFeed with a generic Feed struct's data
func (r *Rss) RssFeed() *RssFeed {
	pub := anyTimeFormat(time.RFC1123Z, r.Created, r.Updated)
	build := anyTimeFormat(time.RFC1123Z, r.Updated)
	author := ""
	if r.Author != nil {
		author = r.Author.Email
		if len(r.Author.Name) > 0 {
			author = fmt.Sprintf("%s (%s)", r.Author.Email, r.Author.Name)
		}
	}

	var image *RssImage
	if r.Image != nil {
		image = &RssImage{Url: r.Image.Url, Title: r.Image.Title, Link: r.Image.Link, Width: r.Image.Width, Height: r.Image.Height}
	}

	channel := &RssFeed{
		Title:          r.Title,
		Link:           r.Link.Href,
		Description:    r.Description,
		ManagingEditor: author,
		PubDate:        pub,
		LastBuildDate:  build,
		Copyright:      r.Copyright,
		Image:          image,
		Atom:           r.Atom,
	}
	for _, i := range r.Items {
		channel.Items = append(channel.Items, newRssItem(i))
	}
	return channel
}

// FeedXml returns an XML-Ready object for an Rss object
func (r *Rss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.RssFeed().FeedXml()

}

// FeedXml returns an XML-ready object for an RssFeed object
func (r *RssFeed) FeedXml() interface{} {
	rssFeedXml := &RssFeedXml{
		Version: "2.0",
		Channel: r,
	}

	if len(r.Items) > 0 {
		if r.Items[0].GoogleId != "" {
			rssFeedXml.GoogleContentNamespace = "http://base.google.com/ns/1.0"
		}

		if r.Items[0].MediaContent != nil {
			rssFeedXml.MediaContentNamespace = "http://search.yahoo.com/mrss/"
		}
	}

	return rssFeedXml
}
