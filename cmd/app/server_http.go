package app

import (
	"net/http"

	backendServ "webapp3/pkg/backend"

	"github.com/go-chi/chi"
)

type Server struct {
	mux        chi.Router
	backendSvc *backendServ.Service
}

// NewServer
func NewServer(
	mux chi.Router,
	backendSvc *backendServ.Service,
) *Server {
	return &Server{
		mux:        mux,
		backendSvc: backendSvc,
	}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

// Init
func (s *Server) Init() error {
	s.mux.Route("/api/v1", func(r chi.Router) {
		//
		r.Get("/echo", s.handleEcho)

		// base posts
		r.Get("/allPosts", s.handleAllPosts)
		r.Get("/post/{post_id}", s.handlePostById)
		r.Post("/postUpdate", s.handleSaveUpdates)
		r.Post("/postNew", s.handleInsertNewPost)
		r.Post("/postDelete/{post_id}", s.handleDeletePostById)

		// activities names
		r.Get("/activ_names", s.handleActivitiesNames)
		r.Post("/activ_names_new", s.handleActivitiesNamesNew)
		r.Post("/activ_names_upd", s.handleActivitiesNamesUpd)
		r.Post("/activ_names_del/{del_id}", s.handleActivitiesNamesDel)

		// activities normatives
		r.Get("/activ_normativs", s.handleActivNormativs)
		r.Post("/activ_normativs_new", s.handleActivNormativsNew)
		r.Post("/activ_normativs_upd", s.handleActivNormativsUpd)
		r.Post("/activ_normativs_del/{del_id}", s.handleActivNormativsDel)

		// activities logs
		r.Get("/activ_logs", s.handleActivitiesLogs)
		r.Post("/activ_logs_new", s.handleActivitiesLogsNew)
		r.Post("/activ_logs_upd", s.handleActivitiesLogsUpd)
		r.Post("/activ_logs_del/{del_id}", s.handleActivitiesLogsDel)

		// analytics
		r.Get("/analytic_params", s.handleAnalyticParams)
		r.Get("/activ_3/{date_from}/{date_to}", s.handleActiv3)
		r.Get("/activ_rep_common/{date_from}/{date_to}", s.handleActivRepCommon)
	})

	return nil
}
