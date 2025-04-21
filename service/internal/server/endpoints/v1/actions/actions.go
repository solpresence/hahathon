package actionsEndpoints

import (
	"encoding/json"
	"hahathon/internal/tabs-client/repo"
	"hahathon/internal/tabs-client/repo/actions"
	"log/slog"
	"net/http"
)

func Create(actionTypes repo.ActionTypesRepo, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Received request", "method", r.Method, "url", r.URL.Path)

		var bodyReq PostReq
		if err := json.NewDecoder(r.Body).Decode(&bodyReq); err != nil {
			log.Warn("failed to decode json", "err", err)
			http.Error(w, "🤡", http.StatusInternalServerError)
			return
		}
		body := actions.PostReq{
			Records: []struct {
				Fields struct {
					IdEmployee []string "json:\"fld4bl4ul8fmz\""
					IdAction   []string "json:\"fldRGb7pEMUb8\""
					IdLocation []string "json:\"fldu9XfPVMGGP\""
				} "json:\"fields\""
			}{
				{Fields: struct {
					IdEmployee []string "json:\"fld4bl4ul8fmz\""
					IdAction   []string "json:\"fldRGb7pEMUb8\""
					IdLocation []string "json:\"fldu9XfPVMGGP\""
				}{
					IdEmployee: []string{
						bodyReq.IdEmployee,
					},
					IdAction: []string{
						bodyReq.IdAction,
					},
					IdLocation: []string{
						bodyReq.IdLocation,
					},
				}},
			},
			FieldKey: "id",
		}
		response, err := actionTypes.Create(body)
		if err != nil {
			log.Warn("failed to decode json", "err", err)
			http.Error(w, "🤡", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
}
