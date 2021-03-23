// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	authEnums "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares/enums"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"
	jwtEnums "github.com/ZupIT/horusec-devkit/pkg/utils/jwt/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type IAuthzMiddleware interface {
	IsApplicationAdmin(next http.Handler) http.Handler
	IsCompanyMember(next http.Handler) http.Handler
	IsCompanyAdmin(next http.Handler) http.Handler
	IsRepositoryMember(next http.Handler) http.Handler
	IsRepositoryAdmin(next http.Handler) http.Handler
	IsRepositorySupervisor(next http.Handler) http.Handler
}

type AuthzMiddleware struct {
	grpcClient auth.AuthServiceClient
	ctx        context.Context
}

func NewAuthzMiddleware(grpcCon grpc.ClientConnInterface) IAuthzMiddleware {
	return &AuthzMiddleware{
		grpcClient: auth.NewAuthServiceClient(grpcCon),
		ctx:        context.Background(),
	}
}

func (a *AuthzMiddleware) IsApplicationAdmin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authConfig, err := a.grpcClient.GetAuthConfig(a.ctx, &auth.GetAuthConfigData{})
		if a.checkGetConfigResponse(err, w) != nil {
			return
		}

		if authConfig.EnableApplicationAdmin {
			response, err := a.grpcClient.IsAuthorized(a.ctx, a.setAuthorizedData(r, authEnums.ApplicationAdmin))
			if a.checkIsAuthorizedResponse(err, response, w, r, authEnums.ApplicationAdmin) != nil {
				return
			}
		}

		handler.ServeHTTP(w, r)
	})
}

func (a *AuthzMiddleware) IsCompanyMember(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := a.grpcClient.IsAuthorized(a.ctx, a.setAuthorizedData(r, authEnums.CompanyMember))
		if a.checkIsAuthorizedResponse(err, response, w, r, authEnums.CompanyMember) != nil {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (a *AuthzMiddleware) IsCompanyAdmin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := a.grpcClient.IsAuthorized(a.ctx, a.setAuthorizedData(r, authEnums.CompanyAdmin))
		if a.checkIsAuthorizedResponse(err, response, w, r, authEnums.CompanyAdmin) != nil {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (a *AuthzMiddleware) IsRepositoryMember(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := a.grpcClient.IsAuthorized(a.ctx, a.setAuthorizedData(r, authEnums.RepositoryMember))
		if a.checkIsAuthorizedResponse(err, response, w, r, authEnums.RepositoryMember) != nil {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (a *AuthzMiddleware) IsRepositorySupervisor(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := a.grpcClient.IsAuthorized(a.ctx, a.setAuthorizedData(r, authEnums.RepositorySupervisor))
		if a.checkIsAuthorizedResponse(err, response, w, r, authEnums.RepositorySupervisor) != nil {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (a *AuthzMiddleware) IsRepositoryAdmin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := a.grpcClient.IsAuthorized(a.ctx, a.setAuthorizedData(r, authEnums.RepositoryAdmin))
		if a.checkIsAuthorizedResponse(err, response, w, r, authEnums.RepositoryAdmin) != nil {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (a *AuthzMiddleware) setAuthorizedData(r *http.Request,
	isAuthorizedType authEnums.IsAuthorizedType) *auth.IsAuthorizedData {
	return &auth.IsAuthorizedData{
		Token:        a.getJWTToken(r),
		Type:         isAuthorizedType.ToString(),
		CompanyID:    chi.URLParam(r, enums.CompanyID),
		RepositoryID: chi.URLParam(r, enums.RepositoryID),
	}
}

func (a *AuthzMiddleware) checkIsAuthorizedResponse(err error, response *auth.IsAuthorizedResponse,
	w http.ResponseWriter, r *http.Request, isAuthorizedType authEnums.IsAuthorizedType) error {
	if err != nil {
		logger.LogError(enums.MessageIsAuthorizedGRPCRequestError, err)
		httpUtil.StatusInternalServerError(w, enums.ErrorFailedToVerifyRequest)
		return enums.ErrorFailedToVerifyRequest
	}

	if !response.GetIsAuthorized() {
		a.logHTTPRequestError(r, isAuthorizedType)
		httpUtil.StatusUnauthorized(w, enums.ErrorUnauthorized)
		return enums.ErrorUnauthorized
	}

	return nil
}

func (a *AuthzMiddleware) logHTTPRequestError(r *http.Request, isAuthorizedType authEnums.IsAuthorizedType) {
	logger.LogWarn(fmt.Sprintf(enums.MessageUnauthorizedHTTPRequest, a.getAccountID(r),
		r.URL, r.Method, isAuthorizedType))
}

func (a *AuthzMiddleware) getAccountID(r *http.Request) string {
	accountID, err := jwt.GetAccountIDByJWTToken(a.getJWTToken(r))
	if err != nil {
		logger.LogError(enums.MessageFailedToGetAccountID, err)
		return uuid.Nil.String()
	}

	return accountID.String()
}

func (a *AuthzMiddleware) getJWTToken(r *http.Request) string {
	return r.Header.Get(jwtEnums.HorusecJWTHeader)
}

func (a *AuthzMiddleware) checkGetConfigResponse(err error, w http.ResponseWriter) error {
	if err != nil {
		logger.LogError(enums.MessageFailedToGetAuthConfig, err)
		httpUtil.StatusInternalServerError(w, enums.ErrorWhenGettingAuthConfig)
		return enums.ErrorWhenGettingAuthConfig
	}

	return nil
}
