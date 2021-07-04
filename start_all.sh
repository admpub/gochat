go build -o gochat.bin -tags=etcd main.go
sudo pkill -9 gochat.bin
sudo pkill -9 etcd
sudo pkill -9 redis-server
etcd &
redis-server &
sleep 5
./gochat.bin -module logic &
./gochat.bin -module connect_websocket &
./gochat.bin -module task &
./gochat.bin -module api &
./gochat.bin -module site