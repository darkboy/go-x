package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
)

var (
	xnode *Node = NewNode()
)

type Node struct {
	discovery.Node
}

func NewNode() *Node {
	node := &Node{}
	node.SetBaseInfoType(uint32(common.Gateway))
	node.InitPolicy(discovery.Random)
	return node
}