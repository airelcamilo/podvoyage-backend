package router

import (
	pc "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/controller"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PodvoyageRouter(r *mux.Router, DB *gorm.DB) {
	var podcastController pc.PodcastController
	podcastController = podcastController.New(DB)
	PodcastRouter(r, podcastController)

	var itemController pc.ItemController
	itemController = itemController.New(DB)
	ItemRouter(r, itemController)

	var folderController pc.FolderController
	folderController = folderController.New(DB)
	FolderRouter(r, folderController)

	var queueController pc.QueueController
	queueController = queueController.New(DB)
	QueueRouter(r, queueController)

	var episodeController pc.EpisodeController
	episodeController = episodeController.New(DB)
	EpisodeRouter(r, episodeController)
}

func PodcastRouter(r *mux.Router, podcastController pc.PodcastController) {
	r.HandleFunc("/api/search-all", podcastController.SearchPodcasts).Methods("POST")
	r.HandleFunc("/api/search-pod/{id}", podcastController.SearchPodcast).Methods("GET")
	r.HandleFunc("/api/podcasts", podcastController.GetAllPodcast).Methods("GET")
	r.HandleFunc("/api/podcast/{id}", podcastController.GetPodcast).Methods("GET")
	r.HandleFunc("/api/podcast", podcastController.SavePodcast).Methods("POST")
	r.HandleFunc("/api/podcast/{id}", podcastController.RemovePodcast).Methods("DELETE")
}

func ItemRouter(r *mux.Router, itemController pc.ItemController) {
	r.HandleFunc("/api/all", itemController.GetAllItem).Methods("GET")
}

func FolderRouter(r *mux.Router, folderController pc.FolderController) {
	r.HandleFunc("/api/folders", folderController.GetAllFolder).Methods("GET")
	r.HandleFunc("/api/folder/{id}", folderController.GetFolder).Methods("GET")
	r.HandleFunc("/api/folder", folderController.SaveFolder).Methods("POST")
	r.HandleFunc("/api/in-folder/{id}", folderController.CheckInFolder).Methods("GET")
	r.HandleFunc("/api/change-folder/{folderId}/{podId}", folderController.ChangeFolder).Methods("GET")
	r.HandleFunc("/api/folder/{id}", folderController.RemoveFolder).Methods("DELETE")
}

func QueueRouter(r *mux.Router, queueController pc.QueueController) {
	r.HandleFunc("/api/queue", queueController.GetAllQueue).Methods("GET")
	r.HandleFunc("/api/queue", queueController.AddToQueue).Methods("POST")
	r.HandleFunc("/api/queue/{id}", queueController.RemoveInQueue).Methods("DELETE")
}

func EpisodeRouter(r *mux.Router, episodeController pc.EpisodeController) {
	r.HandleFunc("/api/played/{id}", episodeController.MarkAsPlayed).Methods("POST")
	r.HandleFunc("/api/current-time/{id}", episodeController.SetCurrentTime).Methods("POST")
}
