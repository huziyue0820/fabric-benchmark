package network

import (
	"benchmark/deploy"
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

const (
	ORDERERIMAGE = "hyperledger/fabric-orderer:1.4"
	PEERIMAGE    = "hyperledger/fabric-peer:1.4"
	FABRICNET    = "fabric_net"
	TOOLSIMAGE   = "hyperledger/fabric-tools:1.4"
)

type OperateIf interface {
	Create() error                     // 创建操作，加载配置
	Remove() error                     // 移除操作
	Start() error                      // 启动操作
	Stop() error                       // 停止操作
	ConnectNet(net deploy.NetIf) error // 连接到docker网络
}

type NodeIf interface {
	OperateIf                              // 复用操作接口，用于操作docker容器
	NodeName() string                      // 返回节点ID
	NodeOrg() OrgIf                        // 返回节点所在域
	Inspect() (types.ContainerJSON, error) // 审查容器状态
	PrintLog() (int64, error)              // 打印容器日志
	ContainerID() string                   // 获取容器的唯一标识
	ContainerName() string                 // 获得容器的名称
	ServePort() nat.Port                   // 获得容器向外暴露的服务端口
}

type OrgIf interface {
	OperateIf       // 复用操作接口，用于批量操作组织内的节点
	Name() string   // 返回组织的名称
	Domain() string // 返回组织所在域
	MSPID() string  // 返回组织的成员管理ID
}
