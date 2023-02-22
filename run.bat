start cmd /k  "etcd"

start cmd /k  "timeout -nobreak 3 && c: && cd .\service\user\rpc && go run user.go"

start cmd /k  "timeout -nobreak 7 && c: && cd .\service\video\rpc && go run video.go"

start cmd /k  "timeout -nobreak 10 && c: && cd .\service\interact\rpc && go run interact.go"

start cmd /k  "timeout -nobreak 13 && c: && cd .\service\social\rpc && go run follow.go"

start cmd /k  "timeout -nobreak 15 && c: && cd .\service\message\rpc && go run message.go"

start cmd /k  "timeout -nobreak 18 && c: && cd .\service\api && go run api.go"


