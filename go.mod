module dev

go 1.15

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go v0.0.0-20190204201341-e444a5086c43
replace vbom.ml/util => github.com/fvbommel/util v0.0.0-20160121211510-db5cfe13f5cc

require (
	github.com/gogo/protobuf v1.3.2
	github.com/stretchr/testify v1.7.0
	vietnt.me/core/golang-sdk v1.2.21 // indirect
	vietnt.me/core/sen-kit v0.1.38
	vietnt.me/protobuf/internal-apis-go v1.25.68
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.40.0
)
