package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	es_secret_name         = "es-credentials"
	es_secret_user_key     = "user"
	es_secret_password_key = "password"
)

type EsCredential struct {
	UserName string
	Password string
}

func getESCredentialWithCache(env Env) (*EsCredential, error) {
	cacheKey := fmt.Sprintf("%s:%s", env, "es_credential")
	return GetOrCalculate(cacheKey, func(string) (*EsCredential, error) {
		return getESCredential(env)
	})
}

func getESCredential(env Env) (*EsCredential, error) {
	s, err := GetSecret(context.TODO(), string(env), es_secret_name)
	if err != nil {
		return nil, err
	}
	credential := EsCredential{UserName: string(s[es_secret_user_key]), Password: string(s[es_secret_password_key])}
	return &credential, nil
}

func CallES(env Env, method string, path string) (string, error) {
	credential, err := getESCredentialWithCache(env)
	logger.Info().Msgf("Env: %s. ES credentials: %+v", env, credential)
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Timeout: time.Second * 100,
	}
	port := ES_STAGING_LOCAL_PORT
	if (env == Production) {
		port = ES_PRODUCTION_LOCAL_PORT
	}
	url := fmt.Sprintf("http://localhost:%d/%s", port, strings.TrimPrefix(path, "/"))
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(credential.UserName, credential.Password)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return string(bodyBytes), nil
}
