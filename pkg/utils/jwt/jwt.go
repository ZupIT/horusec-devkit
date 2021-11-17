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

package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	jwtGO "github.com/form3tech-oss/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/entities"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

func CreateToken(tokenData *entities.TokenData, permissions []string) (string, time.Time, error) {
	expiresAt := time.Now().Add(time.Hour * time.Duration(1))

	tokenSigned, err := newTokenNotSignedWithClaims(tokenData, expiresAt, permissions).SignedString(getHorusecJWTKey())

	return tokenSigned, expiresAt, err
}

func newTokenNotSignedWithClaims(account *entities.TokenData, expiresAt time.Time, permissions []string) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.JWTClaims{
		Email:       account.Email,
		Username:    account.Username,
		Permissions: permissions,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "horusec",
			Subject:   account.AccountID.String(),
		},
	})
}

func DecodeToken(tokenString string) (*entities.JWTClaims, error) {
	token, err := parseStringToToken(strings.ReplaceAll(tokenString, "Bearer ", ""))
	if err != nil {
		return nil, err
	}

	return token.Claims.(*entities.JWTClaims), nil
}

func parseStringToToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &entities.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getHorusecJWTKey(), nil
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	middleware := jwtMiddleware.New(jwtMiddleware.Options{
		ValidationKeyGetter: func(token *jwtGO.Token) (interface{}, error) {
			return getHorusecJWTKey(), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return middleware.Handler(next)
}

func GetAccountIDByJWTToken(token string) (uuid.UUID, error) {
	claims, err := DecodeToken(verifyIfContainsBearer(token))
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(claims.Subject)
}

func getHorusecJWTKey() []byte {
	secretKey := env.GetEnvOrDefault("HORUSEC_JWT_SECRET_KEY", enums.DefaultSecretJWT)
	if secretKey == enums.DefaultSecretJWT {
		logger.LogInfo(enums.MessageWarningDefaultJWTSecretKey)
	}

	return []byte(secretKey)
}

func verifyIfContainsBearer(token string) string {
	if strings.Contains(token, "Bearer") {
		return token
	}

	return fmt.Sprintf("Bearer %s", token)
}

func CreateRefreshToken() string {
	refreshToken := fmt.Sprintf("%s%s", uuid.New(), uuid.New())

	return strings.ReplaceAll(refreshToken, "-", "")
}
