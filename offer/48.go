package offer

import (
	"strconv"
	"strings"
)

//面试题48：序列化和反序列化二叉树

func Serialize(root TreeNode) string {
	if root == nil {
		return "#"
	}
	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(root.val))
	builder.WriteString(",")
	builder.WriteString(Serialize(root.left))
	builder.WriteString(",")
	builder.WriteString(Serialize(root.right))
	return builder.String()
}

func Deserialize(str string) TreeNode {
	s := strings.Split(str, ",")
	index := 0
	return DeserializeHelper(s, &index)
}

func DeserializeHelper(strings []string, index *int) TreeNode {
	if strings[*index] == "#" {
		*index++
		return nil
	}
	val, _ := strconv.ParseInt(strings[*index], 10, 64)
	node := NewTreeNode(int(val))
	*index++
	node.left = DeserializeHelper(strings, index)
	node.right = DeserializeHelper(strings, index)
	return node
}
