package com_haothink_interview

// TreeNode 1 4 6 8
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

var sortedNodeList = make([]*TreeNode, 0)

func MiddleSortQuery(treeNode *TreeNode) *LinkedNode {

	if treeNode == nil {
		return &LinkedNode{}
	}
	sortNode(treeNode)
	// sorted list
	preHead := &LinkedNode{}
	linkedPointer := preHead
	for _, node := range sortedNodeList {
		linkedPointer.Next = &LinkedNode{
			Val: node.Val,
		}
		linkedPointer = linkedPointer.Next
	}
	return preHead.Next
}

func sortNode(treeNode *TreeNode) {

	if treeNode != nil {
		sortNode(treeNode.Left)
		sortedNodeList = append(sortedNodeList, treeNode)
		sortNode(treeNode.Right)
	}
}
