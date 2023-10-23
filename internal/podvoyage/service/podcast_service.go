package podvoyageService

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	"gorm.io/gorm"
)

type PodcastService struct {
	DB                  *gorm.DB
	MAX_CONCURRENT_JOBS int
}

func (s *PodcastService) New(db *gorm.DB) PodcastService {
	return PodcastService{db, 1000}
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

	wg := &sync.WaitGroup{}
	semaphore := make(chan struct{}, s.MAX_CONCURRENT_JOBS)
	for i := range podcast.Episodes {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			podcast.Episodes[i].UserId = userId
			podcast.Episodes[i].TrackId = podcast.TrackId
			<-semaphore
		}(i)
	}
	wg.Wait()
	return podcast, nil
}

func (s *PodcastService) GetAllPodcast(userId int) ([]model.Podcast, error) {
	var podcasts []model.Podcast
	if result := s.DB.Debug().Preload("Episodes", func(db *gorm.DB) *gorm.DB {
		return s.DB.Joins(`
		JOIN (
			SELECT MAX(date) AS max_date, podcast_id 
			FROM episodes WHERE user_id = ? 
			GROUP BY podcast_id
		) e1 ON episodes.podcast_id = e1.podcast_id AND date = e1.max_date`, userId)
	}).Where("user_id = ?", userId).Find(&podcasts); result.Error != nil {
		return podcasts, result.Error
	}

	wg := &sync.WaitGroup{}
	semaphore := make(chan struct{}, s.MAX_CONCURRENT_JOBS)

	for i := range podcasts {
		lastEpisodeDate, err := time.Parse("2006-01-02 15:04:05 -0700 -0700", string(podcasts[i].Episodes[0].Date))
		if err != nil {
			return podcasts, err
		}
		if time.Until(lastEpisodeDate).Abs().Hours()/24 < 7 {
			continue
		}

		semaphore <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if err := s.updatePodcast(userId, podcasts[i], wg); err != nil {
				<-semaphore
				return
			}
			<-semaphore
		}(i)
	}
	wg.Wait()
	return podcasts, nil
}

func (s *PodcastService) updatePodcast(userId int, oldPodcast model.Podcast, wg *sync.WaitGroup) error {
	var newPodcast model.Podcast
	var episodes []model.Episode
	response, err := http.Get(oldPodcast.FeedUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	dec := xml.NewDecoder(response.Body)
	err = dec.Decode(&newPodcast)
	if err != nil {
		return err
	}

	for i := range newPodcast.Episodes {
		if newPodcast.Episodes[i].Title == oldPodcast.Episodes[0].Title {
			if len(episodes) == 0 {
				break
			}
			if result := s.DB.Create(&episodes); result.Error != nil {
				return result.Error
			}
			break
		}

		newPodcast.Episodes[i].PodcastId = oldPodcast.Id
		newPodcast.Episodes[i].TrackId = oldPodcast.TrackId
		newPodcast.Episodes[i].UserId = userId
		episodes = append(episodes, *newPodcast.Episodes[i])
	}
	return nil
}

func (s *PodcastService) GetPodcast(id int, userId int) (model.Podcast, error) {
	var podcast model.Podcast
	if result := s.DB.Preload("Categories").Preload("Episodes", func(db *gorm.DB) *gorm.DB {
		return s.DB.Order("episodes.date DESC")
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
