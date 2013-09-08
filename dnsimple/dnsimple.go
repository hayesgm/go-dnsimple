package dnsimple

/*
  This is a package to handle interactions with the DNSimple API.
*/

import (
  "net/http"
  "fmt"
  "path"
  "strings"
  "io/ioutil"
  "encoding/json"
  "strconv"
)

const (
  endpoint = "https://dnsimple.com/"

  CNAME_RECORD  = "CNAME"
  A_RECORD      = "A"
  TXT_RECORD    = "TXT"
  POOL_RECORD   = "POOL"
)

// A client is the main interaction with the DNSimple API
type Client struct {
  Auth Authorizer
}

func (cli *Client) makeRequest(method, path string, params map[string]interface{}) (respBody string, err error) {
  httpClient := &http.Client{}

  body, err := json.Marshal(params)
  if err != nil {
    return
  }

  req, err := http.NewRequest(method, fmt.Sprintf("%s%s", endpoint, path), strings.NewReader(string(body)))
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/json")
  cli.Auth.Authorize(req)

  resp, err := httpClient.Do(req)
  if err != nil {
    return
  }

  respBodyBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return
  }
  
  respBody = string(respBodyBytes)

  return
}

func (cli *Client) GetRecords(domain string, subdomain string) (resp []RecordObj, err error) {
  
  body, err := cli.makeRequest("GET", path.Join("domains", domain, "records"), nil)
  if err != nil {
    return
  }

  if err = json.Unmarshal([]byte(body), &resp); err != nil {
    return
  }

  return resp, nil
}

func (cli *Client) CreateRecord(domain string, subdomain string, recordType string, content string, ttl int, priority int) (record RecordObj, err error) {
  params := make(map[string]interface{})

  recordObj := make(map[string]interface{})
  recordObj["name"] = subdomain
  recordObj["record_type"] = recordType
  recordObj["content"] = content
  recordObj["ttl"] = ttl
  recordObj["priority"] = priority

  params["record"] = recordObj

  body, err := cli.makeRequest("POST", path.Join("domains", domain, "records"), params)
  if err != nil {
    return
  }

  if err = json.Unmarshal([]byte(body), &record); err != nil {
    return
  }

  return record, nil
}

func (cli *Client) DeleteRecord(domain string, subdomain string, record Record) (err error) {
  _, err = cli.makeRequest("DELETE", path.Join("domains", domain, "records", strconv.Itoa(record.Id)), nil)

  return
}