package dnsimple

type Record struct {
  Name string
  Ttl int
  CreatedAt string  `json:"created_at"`
  UpdatedAt string  `json:"updated_at"`
  DomainId int      `json:"domain_id"`
  Id int
  Content string
  RecordType string `json:"record_type"`
  Prio int
}

type RecordObj struct {
  Record Record
}