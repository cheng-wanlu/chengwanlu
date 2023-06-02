# etcd

## 安装

```bash
ETCD_VER=v3.5.0

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

mkdir -p /data/etcd
cd /data/etcd/
wget ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf etcd-${ETCD_VER}-linux-amd64.tar.gz

/data/etcd/etcd-${ETCD_VER}-linux-amd64/etcd --version
/data/etcd/etcd-${ETCD_VER}-linux-amd64/etcdctl version
/data/etcd/etcd-${ETCD_VER}-linux-amd64/etcdutl version
```

## 启动

```bash
# start a local etcd server
/data/etcd/etcd-${ETCD_VER}-linux-amd64/etcd
```

## write,read to etcd
```bash
/data/etcd/etcd-${ETCD_VER}-linux-amd64/etcdctl --endpoints=localhost:2379 put foo bar
/data/etcd/etcd-${ETCD_VER}-linux-amd64/etcdctl --endpoints=localhost:2379 get foo
```


## 修改 /data/etcd/confs/etcd.conf

```text
NAME: etcd_sg
DATA-DIR: "/data/etcd/etcd_data"
LISTEN-CLIENT-URLS: "http://10.0.2.15:2379,http://127.0.0.1:2379"
ADVERTISE-CLIENT-URLS: "http://10.0.2.15:2379,http://127.0.0.1:2379"
INITIAL-CLUSTER-TOKEN: "sg"
INITIAL-CLUSTER: "etcd_sg=http://10.0.2.15:2380"
INITIAL-ADVERTISE-PEER-URLS: "http://10.0.2.15:2380"
```

```bash
./etcdctl user add root
Password of root: J7dNelJ3&LyQV9KR

./etcdctl auth enable

./etcdctl user add server --user=root
Password of server: w6VAeLsmFqOt%aec

./etcdctl user add client --user=root
Password of client: knni0VK%8t&Vva8k

./etcdctl role add read --user=root

./etcdctl --user=root role grant-permission read read /rpc/ --prefix=true

./etcdctl --user=root user grant-role server root
./etcdctl --user=root user grant-role client read

./etcdctl --user=root --endpoints=101.32.216.100:2379 get "" --prefix --keys-only
```