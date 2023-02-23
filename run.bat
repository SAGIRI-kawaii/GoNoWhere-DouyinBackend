start cmd /k  "etcd"

start cmd /k  "timeout -nobreak 5 && c: && go run .\service\user\rpc\user.go"

start cmd /k  "timeout -nobreak 10 && c: && go run .\service\video\rpc\video.go"

start cmd /k  "timeout -nobreak 15 && c: && go run .\service\interact\rpc\interact.go"

start cmd /k  "timeout -nobreak 20 && c: && go run .\service\social\rpc\follow.go"

start cmd /k  "timeout -nobreak 25 && c: && go run .\service\message\rpc\message.go"

start cmd /k  "timeout -nobreak 30 && c: && go run .\service\api\api.go"
