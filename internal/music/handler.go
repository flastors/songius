package music

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/httputil"
	"path"
	"strings"

	"github.com/flastors/songius/internal/apperror"
	"github.com/flastors/songius/internal/handlers"
	"github.com/flastors/songius/internal/music/model"
	"github.com/flastors/songius/internal/music/service"
	"github.com/flastors/songius/pkg/api/filter"
	"github.com/flastors/songius/pkg/utils/logging"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	songsURL = "/api/v1/songs"
	songURL  = "/api/v1/songs/:uuid"
)

type handler struct {
	logger  *logging.Logger
	service *service.Service
}

func NewHandler(service *service.Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	h.logger.Debug("Registering Song Handlers")
	router.HandlerFunc(http.MethodGet, songsURL, filter.Middleware(apperror.Middleware(h.GetList), 10))
	router.HandlerFunc(http.MethodPost, songsURL, apperror.Middleware(h.Create))
	router.HandlerFunc(http.MethodGet, songURL, apperror.Middleware(h.GetByUUID))
	router.HandlerFunc(http.MethodPut, songURL, apperror.Middleware(h.Update))
	router.HandlerFunc(http.MethodDelete, songURL, apperror.Middleware(h.Delete))
	router.HandlerFunc(http.MethodGet, "/swagger/:any", httpSwagger.WrapHandler)
}

// GetList godoc
// @Summary      List musics
// @Description  get musics
// @Tags         music
// @Accept       json
// @Produce      json
// @Param        song   query      string  false  "song filter"
// @Param        group   query      string  false  "group filter"
// @Param        release_date   query      string  false  "release_date filter"
// @Param        link   query      string  false  "link filter"
// @Param        text   query      string  false  "text filter"
// @Param        limit   query      string  false  "set output limit"
// @Param        offset   query      string  false  "set offset"
// @Success      200  {array}  model.Music
// @Failure      418
// @Router       /songs [get]
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	filterOptions := r.Context().Value(filter.OptionsContextKey).(filter.Options)

	song := r.URL.Query().Get("song")
	if song != "" {
		filterOptions.AddField("song", filter.OperatorLike, song, filter.DataTypeStr)
	}
	group := r.URL.Query().Get("group")
	if group != "" {
		filterOptions.AddField("artist", filter.OperatorLike, group, filter.DataTypeStr)
	}
	releaseDate := r.URL.Query().Get("release_date")
	if releaseDate != "" {
		var operator string
		if strings.Index(releaseDate, ":") != -1 {
			operator = filter.OperatorBetween
		} else {
			operator = filter.OperatorEq
		}
		err := filterOptions.AddField("release_date", operator, releaseDate, filter.DataTypeDate)
		if err != nil {
			return err
		}
	}
	link := r.URL.Query().Get("link")
	if link != "" {
		filterOptions.AddField("link", filter.OperatorLike, link, filter.DataTypeStr)
	}
	text := r.URL.Query().Get("text")
	if text != "" {
		filterOptions.AddField("text", filter.OperatorLike, text, filter.DataTypeStr)
	}

	w.Header().Set("Content-Type", "application/json")
	all, err := h.service.GetAll(r.Context(), filterOptions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

// Create godoc
// @Summary      Create music
// @Description  create music
// @Tags         music
// @Accept       json
// @Produce      json
// @Param        music body         model.CreateMusicDTO  true  "Create music"
// @Success      200  {object}  model.Music
// @Failure      418
// @Router       /songs [post]
func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	var mdto model.CreateMusicDTO
	err := json.NewDecoder(r.Body).Decode(&mdto)
	if err != nil {
		return err
	}
	m, err := h.service.Create(r.Context(), &mdto)
	if err != nil {
		return err
	}
	mBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mBytes)

	return nil
}

// GetByUUID godoc
// @Summary      Show a music
// @Description  get music by ID
// @Tags         music
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Music ID"
// @Success      200  {object}  model.Music
// @Failure      418
// @Router       /songs/{id} [get]
func (h *handler) GetByUUID(w http.ResponseWriter, r *http.Request) error {
	m, err := h.service.GetOne(r.Context(), path.Base(r.RequestURI))
	if err != nil {
		return err
	}
	mBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(mBytes)
	return nil
}

// Update godoc
// @Summary      Update music
// @Description  update music
// @Tags         music
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Music ID"
// @Param        music body         model.UpdateMusicDTO  true  "update music"
// @Success      200
// @Failure      418
// @Router       /songs/{id} [put]
func (h *handler) Update(w http.ResponseWriter, r *http.Request) error {
	var mudto model.UpdateMusicDTO
	err := json.NewDecoder(r.Body).Decode(&mudto)
	if err != nil {
		return err
	}
	m := model.NewMusicModel(mudto.Song, mudto.Group, mudto.ReleaseDate, mudto.Link, strings.ReplaceAll(mudto.Text, "'", "''"))
	err = h.service.Update(r.Context(), m)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update is succesfull"))
	return nil
}

// Delete godoc
// @Summary      delete a music
// @Description  delete music by ID
// @Tags         music
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Music ID"
// @Success      200
// @Failure      418
// @Router       /songs/{id} [delete]
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) error {
	err := h.service.Delete(r.Context(), path.Base(r.RequestURI))
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Successfully deleted object with id: %s", path.Base(r.RequestURI))))
	return nil
}
