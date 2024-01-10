# Traffic Replicator

This document is translated from Chinese by machine. If you find any problems in the translation, please feel free to contact me.

[简体中文](README.md) | English

## Project Introduction

This project is a traffic replication tool based on Go. It can replicate the traffic of a specified port to the same port on other hosts, supporting both TCP and UDP protocols.

UDP is one-way replication. After the port receives data, it will copy the data to all target hosts.

TCP is two-way replication. After the port receives data, it will copy the data to all target hosts, and will also copy the data returned by the first connected target host back to the source host.

## Usage

### 1. Download

#### 1.1 Download and install using go

```shell
go install github.com/taills/traffic-replicator/cmd/ipcopy@latest
```

#### 1.2 Use binary files

Download the binary file for the corresponding platform from https://github.com/taills/traffic-replicator/releases.

### 2. Run

Copy the UDP traffic of ports 53, 514-550 on this machine to two hosts, 192.168.0.22 and 192.168.0.23, and output the data packets in ASCII format.

```shell
ipcopy -targets 192.168.0.22,192.168.0.23 -ports 53,514-550 -udp -ascii
```

Copy the TCP traffic of port 1024 on this machine to 10.100.0.171.

```shell
ipcopy -targets 10.100.0.171 -ports 1024 -tcp -ascii
```

### 3. Help

```shell
ipcopy -h
```