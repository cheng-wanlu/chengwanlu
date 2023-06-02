# 挂载磁盘

```bash
# 查看磁盘使用情况
df -h
# 列出系统中所有的磁盘分区信息
fdisk -l
# 列出系统中的块设备信息
lsblk
# 查看磁盘分区的详细信息
parted -l
# 安装 XFS 文件系统的工具和库
apt install xfsprogs
# 在设备 /dev/sdb 上创建 XFS 文件系统. -f 参数表示强制格式化.
mkfs.xfs -f /dev/sdb
# 创建 /data 的目录, 作为挂载点使用
mkdir /data
# 将设备 /dev/sdb 挂载到 /data 目录, 使其可以访问该设备的文件系统。
mount /dev/sdb /data
```

- 修改配置文件: /etc/fstab, 添加一行

```text
/dev/sdb                /data                   xfs     defaults        0 0
```














