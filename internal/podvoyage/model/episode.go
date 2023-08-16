package model

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type Episode struct {
	Id          int              `json:"id" gorm:"primaryKey"`
	PodcastId   int              `json:"podcastId"`
	TrackId     int              `json:"trackId"`
	Title       string           `json:"title" xml:"title"`
	Desc        string           `json:"desc" xml:"encoded"`
	Season      int              `json:"season" xml:"season"`
	Date        customDate       `json:"date" xml:"pubDate"`
	Duration    customDuration   `json:"duration" xml:"duration"`
	Audio       customAudio      `json:"audio" xml:"enclosure"`
	ArtworkUrl  customArtworkUrl `json:"artworkUrl" xml:"image"`
	Played      bool             `json:"played"`
	CurrentTime float64          `json:"currentTime"`
}

type customDuration int
type customDate string
type customAudio string
type customArtworkUrl string

func (c *customDuration) parse(st string) (int, error) {
	var h, m, s int

	stSlice := strings.Split(st, ":")
	if len(stSlice) == 3 {
		h, m, s = c.conv(stSlice[0]), c.conv(stSlice[1]), c.conv(stSlice[2])
		return h*3600 + m*60 + s, nil
	} else {
		m, s = c.conv(stSlice[0]), c.conv(stSlice[1])
		return m*60 + s, nil
	}
}

func (c *customDuration) conv(s string) int {
	s, _ = strings.CutSuffix(s, "0")
	a, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return a
}

func (c *customDuration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data string
	var result int
	var err error
	d.DecodeElement(&data, &start)
	if strings.Contains(data, ":") {
		result, err = c.parse(data)
	} else {
		result, err = strconv.Atoi(data)
	}
	if err != nil {
		return err
	}
	*c = customDuration(result)
	return nil
}

func (c *customDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data string
	d.DecodeElement(&data, &start)
	result, err := time.Parse(time.RFC1123Z, data)
	if err != nil {
		result, err = time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", data)
		if err != nil {
			return err
		}
	}
	*c = customDate(result.String())
	return nil
}

func (c *customAudio) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data string
	d.DecodeElement(&data, &start)
	for _, attr := range start.Attr {
		if attr.Name.Local == "url" {
			*c = customAudio(attr.Value)
		}
	}
	return nil
}

func (c *customArtworkUrl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data string
	d.DecodeElement(&data, &start)
	for _, attr := range start.Attr {
		if attr.Name.Local == "href" {
			*c = customArtworkUrl(attr.Value)
		}
	}
	return nil
}
