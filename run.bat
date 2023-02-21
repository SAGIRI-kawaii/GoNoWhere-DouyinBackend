start cmd /k  "etcd"

start cmd /k  "timeout -nobreak 3 && c: && cd .\service\user\rpc && go run user.go -f etc/user.yaml"

start cmd /k  "timeout -nobreak 7 && c: && cd .\service\video\rpc && go run video.go -f etc/video.yaml"

start cmd /k  "timeout -nobreak 10 && c: && cd .\service\interact\rpc && go run interact.go -f etc/interact.yaml"

start cmd /k  "timeout -nobreak 13 && c: && cd .\service\social\rpc && go run follow.go -f etc/follow.yaml"

start cmd /k  "timeout -nobreak 15 && c: && cd .\service\message\rpc && go run message.go -f etc/message.yaml"

start cmd /k  "timeout -nobreak 18 && c: && cd .\service\api && go run api.go -f etc/api.yaml"


