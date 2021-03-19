package swagger

import (
	"fmt"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/ZupIT/horusec-devkit/pkg/services/swagger/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type ISwagger interface {
	SetupSwagger()
	GetSwaggerHost() string
}

type Swagger struct {
	router *chi.Mux
	port   string
	host   string
}

func NewSwagger(router *chi.Mux, defaultPort string) ISwagger {
	return &Swagger{
		router: router,
		port:   env.GetEnvOrDefault("HORUSEC_PORT", defaultPort),
		host:   env.GetEnvOrDefault("HORUSEC_SWAGGER_HOST", "localhost"),
	}
}

func (s *Swagger) SetupSwagger() {
	s.routerSwagger()

	logger.LogInfo(fmt.Sprintf(enums.SwaggerURLMessage, s.GetSwaggerHost()))
}

func (s *Swagger) routerSwagger() {
	swaggerConfig := httpSwagger.URL(fmt.Sprintf(enums.SwaggerURL, s.host, s.port))

	s.router.Get(enums.SwaggerRoute, httpSwagger.Handler(swaggerConfig))
}

func (s *Swagger) GetSwaggerHost() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}
