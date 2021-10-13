package deploy

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

type NodeConfigIf interface {
	Host() string                     // 节点的主机名
	Domain() OrgConfigIf              // 节点所在域
	Port() nat.Port                   // 节点的服务端口
	JoinOrg(domain OrgConfigIf) error // 节点加入组织
	NetName() string                  // 节点链接的网络名
}

type OrgConfigIf interface {
	Name() string                    // 返回组织的名称
	Domain() string                  // 返回组织所在域
	MSPID() string                   // 返回组织的成员管理ID
	CryptoPath() string              // 返回成员管理密钥目录
	ConfigtxPath() string            // 返回创世区块、通道交易所在路径
	AddNode(node NodeConfigIf) error // 向组织中添加节点
	DelNode(host string) error       // 根据host移除该节点
	GetNodes() []NodeConfigIf        // 获得当前组织中的所有节点
}

type NetIf interface {
	NetName() string                // 返回网络名称
	NetID() string                  // 返回网络ID
	CreateNet() error               // 创建docker网络
	RemoveNet() error               // 移除docker网络
	IsNetExist() bool               // 移除docker网络
	Inspect() types.NetworkResource // 审查docker网络状态
}
