# go-envoymesher

#cds.proto protoc
import_path:D:/protoc-3.7.0-rc-3-win64/include
import_path:D:/go/src/github.com/lyft/protoc-gen-validate
import_paht:D:/go/src/github.com/envoyproxy/data-plane-api/envoy/api
import_path:D:\go\src\github.com\googleapis\googleapis
protoc ./cds.proto -I. -ID:\go\src\github.com\googleapis\googleapis -ID:/protoc-3.7.0-rc-3-win64/include -ID:/go/src/github.com/lyft/protoc-gen-validate -ID:/go/src/github.com/envoyproxy/data-plane-api/ --go_out=grpc:.

#eds.proto protoc
protoc ./eds.proto -I. -ID:\go\src\github.com\googleapis\googleapis -ID:/protoc-3.7.0-rc-3-win64/include -ID:/go/src/github.com/lyft/protoc-gen-validate -ID:/go/src/github.com/envoyproxy/data-plane-api/ --go_out=grpc:.

#lds.proto protoc
protoc ./lds.proto -I. -ID:\go\src\github.com\googleapis\googleapis -ID:/protoc-3.7.0-rc-3-win64/include -ID:/go/src/github.com/lyft/protoc-gen-validate -ID:/go/src/github.com/envoyproxy/data-plane-api/ --go_out=grpc:.

#rds.proto protoc
protoc ./rds.proto -I. -ID:\go\src\github.com\googleapis\googleapis -ID:/protoc-3.7.0-rc-3-win64/include -ID:/go/src/github.com/lyft/protoc-gen-validate -ID:/go/src/github.com/envoyproxy/data-plane-api/ --go_out=grpc:.

#discovery.proto protoc
protoc ./discovery.proto -I. -ID:\go\src\github.com\googleapis\googleapis -ID:/protoc-3.7.0-rc-3-win64/include -ID:/go/src/github.com/lyft/protoc-gen-validate -ID:/go/src/github.com/envoyproxy/data-plane-api/ --go_out=grpc:.

