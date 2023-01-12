package offer

type TreeNode *struct {
	val   int
	left  TreeNode
	right TreeNode
}

func NewTreeNode(val int) TreeNode {
	return &struct {
		val   int
		left  TreeNode
		right TreeNode
	}{val: val}
}

func ConstructCBT(input []int) TreeNode {
	if len(input) == 0 {
		return nil
	}
	root := NewTreeNode(input[0])
	inserter := NewCBTInserter(root)
	for i := 1; i < len(input); i++ {
		inserter.Insert(input[i])
	}
	return root
}

func IsEqual(root1, root2 TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil || root1 != nil && root2 == nil {
		return false
	}
	return root1.val == root2.val && IsEqual(root1.left, root2.left) && IsEqual(root1.right, root2.right)
}
