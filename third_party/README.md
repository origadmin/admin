# third_party

The current directory is `buf.build` backup dependent on the third-party proto file. You are not advised to change the
directory if necessary

```shell
make deps
# or use buf command to export
buf export buf.build/bufbuild/protovalidate -o third_party
buf export buf.build/protocolbuffers/wellknowntypes -o third_party
buf export buf.build/googleapis/googleapis -o third_party
buf export buf.build/envoyproxy/protoc-gen-validate -o third_party
buf export buf.build/gnostic/gnostic -o third_party
buf export buf.build/kratos/apis -o third_party
buf export buf.build/origadmin/rpcerr -o third_party
buf export buf.build/origadmin/runtime -o third_party
buf export buf.build/origadmin/entgen -o third_party
```