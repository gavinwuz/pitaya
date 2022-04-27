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
    
    "run-cluster-protobuf-frontend-example"{
        cd examples/demo/cluster
        go run main.go
        break
    }
    "run-cluster-protobuf-backend-example"{
        cd examples/demo/cluster
        go run main.go --port 3251 --type room --frontend=false
        break
    }
    "run-cluster-grpc-example-connector" {
        cd examples/demo/cluster_grpc;
        go run main.go
        break
    }
    "run-cluster-grpc-example-room" {
        cd examples/demo/cluster_grpc;
        go run main.go --port 3251 --rpcsvport 3435 --type room --frontend=false
        break
    }
    "protos-compile-demo-worker" {
        protoc -I examples/demo/worker/protos examples/demo/worker/protos/*.proto --go_out=.
        break
    }
    "run-cluster-worker-example-room" {
        cd examples/demo/worker;
        go run main.go --type room --frontend=true
        break
    }
    "run-cluster-worker-example-metagame" {
        cd examples/demo/worker;
        go run main.go --type metagame --frontend=false
        break
    }
    "run-cluster-worker-example-worker" {
        cd examples/demo/worker;
        go run main.go --type worker --frontend=false
        break
    }
    "run-rate-limiting-example" {
        cd examples/demo/rate_limiting;
        go run main.go
        break
    }
    "run-pipeline-example" {
        cd examples/demo/pipeline;
        go run main.go
        break
    }    
    "run-custom-metrics-example"{
        cd examples/demo/custom_metrics
        go run main.go --port 3250
        break
    }

    "protos-compile-demo" {
        protoc -I examples/demo/protos examples/demo/protos/*.proto --go_out=.
        break
    }
    "protos-compile" {
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
        break
    }
    default {
        echo "help cmd, ex: .\make.ps1 setup";
        echo "  setup"
        echo "  $cmd"
        break 
    }    
}