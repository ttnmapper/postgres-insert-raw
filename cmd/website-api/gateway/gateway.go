package gateway

import "github.com/go-chi/chi"

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/radar/{network_id}/{gateway_id}", GetGatewayRadar)

	return router
}
