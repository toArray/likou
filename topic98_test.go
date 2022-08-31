package luoqiangLikou

import (
	"fmt"
	"testing"
)

type TreeNode98 struct {
	Val   int
	Left  *TreeNode98
	Right *TreeNode98
}

//19. 删除链表的倒数第 N 个结点
func Test98(t *testing.T) {
	head := &TreeNode98{Val: 5, Left: &TreeNode98{
		Val:   1,
		Left:  nil,
		Right: nil,
	}, Right: &TreeNode98{
		Val: 4,
		Left: &TreeNode98{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode98{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}}

	//head := &TreeNode98{Val: 2, Left: &TreeNode98{
	//	Val:   1,
	//	Left:  nil,
	//	Right: nil,
	//}, Right: &TreeNode98{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}}

	data := isValidBST(head)
	fmt.Println(data)
}

var stack []int

func isValidBST(root *TreeNode98) bool {
	stack = []int{}
	dfs98(root)
	for i := 0; i < len(stack); i++ {
		if stack[i] >= stack[i+1] {
			return false
		}
	}
	return true
}

func dfs98(root *TreeNode98) {
	if root == nil {
		return
	}

	dfs98(root.Left)
	stack = append(stack, root.Val)
	dfs98(root.Right)
}
