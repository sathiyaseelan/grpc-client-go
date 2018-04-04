
# gRPC Client in go


### steps

 1. Copy the proto file from the source to the proto folder

 2. Generate the stubs
 ./protoc -I proto/ proto/userdata.proto --go_out=plugins=grpc:userdata_v1

 3. Run the client
 go run main.go sathya@example.com
