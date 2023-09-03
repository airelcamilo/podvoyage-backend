package podvoyageService

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	"gorm.io/gorm"
)

type PodcastService struct {
	DB *gorm.DB
}

func (s *PodcastService) New(db *gorm.DB) PodcastService {
	return PodcastService{db}
}

func (s *PodcastService) SearchPodcasts(request *model.SearchRequest) (model.SearchResult, error) {
	var searchResult model.SearchResult
	baseUrl, err := url.Parse("https://itunes.apple.com/search?media=podcast")
	if err != nil {
		return searchResult, err
	}
	values := baseUrl.Query()
	values.Add("term", request.PodcastName)
	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		return searchResult, err
	}
	defer response.Body.Close()

	bodyByte, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(bodyByte, &searchResult)
	if err != nil {
		return searchResult, err
	}
	return searchResult, nil
}

func (s *PodcastService) SearchPodcast(id int, userId int) (model.Podcast, error) {
	var searchResult model.SearchResult
	var podcast model.Podcast
	response, err := http.Get("https://itunes.apple.com/lookup?id=" + strconv.Itoa(id))

	if err != nil {
		return podcast, err
	}

	bodyByte, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(bodyByte, &searchResult)
	if err != nil {
		return podcast, err
	}
	podcast = *searchResult.Results[0]
	response.Body.Close()

	response, err = http.Get(podcast.FeedUrl)
	if err != nil {
		return podcast, err
	}
	defer response.Body.Close()

	dec := xml.NewDecoder(response.Body)
	err = dec.Decode(&podcast)
	if err != nil {
		return podcast, err
	}

	for _, episode := range podcast.Episodes {
		episode.UserId = userId
		episode.TrackId = podcast.TrackId
	}
	return podcast, nil
}

func (s *PodcastService) GetAllPodcast(userId int) ([]model.Podcast, error) {
	var podcasts []model.Podcast
	if result := s.DB.Where("user_id = ?", userId).Find(&podcasts); result.Error != nil {
		return podcasts, result.Error
	}
	return podcasts, nil
}

func (s *PodcastService) GetPodcast(id int, userId int) (model.Podcast, error) {
	var podcast model.Podcast
	if result := s.DB.Preload("Categories").Preload("Episodes", func(db *gorm.DB) *gorm.DB {
		return s.DB.Order("episodes.id ASC")
	}).Where("user_id = ?", userId).First(&podcast, id); result.Error != nil {
		return podcast, result.Error
	}
	return podcast, nil
}

func (s *PodcastService) SavePodcast(request *model.Podcast, userId int) (model.Podcast, error) {
	podcast := *request
	podcast.UserId = userId
	var temp model.Podcast

	if result := s.DB.Where("track_id = ? AND user_id = ?", podcast.TrackId, userId).First(&temp); result.Error == nil {
		return podcast, errors.New("podcast already saved")
	}

	if result := s.DB.Create(&podcast); result.Error != nil {
		return podcast, result.Error
	}
	return podcast, nil
}

func (s *PodcastService) RemovePodcast(id int, userId int) error {
	var podcast model.Podcast
	if result := s.DB.Where("user_id = ?", userId).First(&podcast, id).Delete(&podcast); result.Error != nil {
		return result.Error
	}
	return nil
}
