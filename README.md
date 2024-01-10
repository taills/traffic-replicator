# 流量复制者 

简体中文 | [English](./README.en.md)

## 项目介绍

本项目是一个基于Go的流量复制工具，可以将指定端口的流量复制到其他主机上相同的端口，支持TCP和UDP协议。

UDP为单向复制，端口收到数据后，会将数据复制到所有的目标主机上。

TCP为双向复制，端口收到数据后，会将数据复制到所有的目标主机上，同时也会将第一台连接上的目标主机返回的数据复制到源主机上。

## 使用方法

### 1. 下载

#### 1.1 使用 go 下载安装

```shell
go install github.com/taills/traffic-replicator/cmd/traffic-replicator@latest
```

#### 1.2 使用二进制文件

从 https://github.com/taills/traffic-replicator/releases 下载对应平台的二进制文件。


### 2. 运行

复制本机53端口，514-550端口的UDP流量到 192.168.0.22 和 192.168.0.23 两台主机上，并以ASCII格式输出数据包。

```shell
traffic-replicator -targets 192.168.0.22,192.168.0.23 -ports 53,514-550 -udp -ascii
```

复制本机的 1024 端口的TCP流量到 10.100.0.171

```shell
traffic-replicator -targets 10.100.0.171 -ports 1024 -tcp -ascii
```

### 3. 帮助

```shell
traffic-replicator -h
```