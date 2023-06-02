# 主机名管理与查询 

`hostnamectl` 是一个用于管理和查询系统主机名的命令行工具. 它提供了一些方便的功能, 包括显示当前主机名、设置新的主机名以及查询主机名设置的状态. 本文档将介绍如何使用 `hostnamectl` 命令进行主机名的管理与查询.

## 1. 显示当前主机名

要显示当前系统的主机名，可以使用以下命令：

```bash
hostnamectl
```
该命令将显示与主机名相关的信息, 包括静态主机名(Static hostname)、交互式主机名(Transient hostname)以及主机名状态(Hostname status)等.

## 2. 设置新的主机名

您可以使用 `hostnamectl set-hostname` 命令来设置系统的新主机名。以下是设置主机名的示例命令：

```bash
sudo hostnamectl set-hostname ubuntu01
```

