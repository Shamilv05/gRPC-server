# gRPC-server
Simple gRPC server written in C++

To compile protoc files, run the following command:
```
$ protoc -I <path_to_folder_with_protoc> --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` <path_to_folder_with_protoc>/<protoc_filename>
$ protoc -I <path_to_folder_with_protoc> --cpp_out=. <path_to_folder_with_protoc>/<protoc_filename>
```
