package dnsimple

import (
  "fmt"
  "net/http"
)

type Authorizer interface {
  // Authorize should add its authorization code to the request
  Authorize(request *http.Request)
}

// Domain authorization
// Authorization for a single domain
type DomainAuth struct {
  domain string
  token string
}

func NewDomainAuth(domain, token string) (auth Authorizer) {
  return DomainAuth{domain: domain, token: token}
}

func (auth DomainAuth) Authorize(request *http.Request) {
  request.Header.Add("X-DNSimple-Domain-Token", auth.token)
}

// Token authorization
// Authorization by token
type TokenAuth struct {
  email string
  token string
}

func NewTokenAuth(email, token string) (auth Authorizer) {
  return TokenAuth{email: email, token: token}
}

func (auth TokenAuth) Authorize(request *http.Request) {
  request.Header.Add("X-DNSimple-Token", fmt.Sprintf("%s:%s", auth.email, auth.token))
}