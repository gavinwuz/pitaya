$cmd = $args[0]

Switch ($cmd)
{
    "init" {
        echo "git submodule init..."
        git submodule update --init --recursive
        break
    }
    "setup" {        
        echo "go get..."
        go get ./...
        break
    }
    "run-cluster-grpc-example-connector" {
        cd examples/demo/cluster_grpc;
        go run main.go
    }
    "run-cluster-grpc-example-room" {
        cd examples/demo/cluster_grpc;
        go run main.go --port 3251 --rpcsvport 3435 --type room --frontend=false
    }
    "protos-compile-demo" {
        protoc -I examples/demo/protos examples/demo/protos/*.proto --go_out=.
    }
    "protos-compile"{
        cd benchmark/testdata
        ./gen_proto.bat
        echo "build test.proto"
        cd ../../
        # protoc -I pitaya-protos/ pitaya-protos/*.proto --go_out=plugins=grpc:protos #旧版本
        protoc -I pitaya-protos/ pitaya-protos/*.proto --go-grpc_out=. --go_out=. # protos
        echo "build out dir ./protos"
        #
        protoc -I pitaya-protos/test pitaya-protos/test/*.proto --go_out=. # protos/test
        echo "build out dir ./protos/test"
    }
    default {
        echo "help cmd, ex: .\make.ps1 setup";
        echo "  setup"
        echo "  run-cluster-grpc-example-connector"
        break 
    }    
}