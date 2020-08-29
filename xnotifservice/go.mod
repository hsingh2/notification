module cto-github.cisco.com/NFV-BU/xnotifservice

go 1.14

require (
	cto-github.cisco.com/NFV-BU/go-msx-build v0.1.6
	github.com/cockroachdb/cockroach-go v2.0.1+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/consul/api v1.6.0
	github.com/lib/pq v1.8.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/hashicorp/consul v1.4.0 => github.com/hashicorp/consul v1.6.0
