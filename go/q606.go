package program

import (
	"bytes"
	"strconv"
)

func tree2str(t *TreeNode) string {
	var bt bytes.Buffer
	tree2strIn(t, &bt)
	return bt.String()
}

func tree2strIn(t *TreeNode, bt *bytes.Buffer) {
	if t == nil {
		return
	}
	bt.WriteString(strconv.Itoa(t.Val))
	if t.Left != nil || t.Right != nil {
		bt.WriteRune('(')
		tree2strIn(t.Left, bt)
		bt.WriteRune(')')
	}
	if t.Right != nil {
		bt.WriteRune('(')
		tree2strIn(t.Right, bt)
		bt.WriteRune(')')
	}

}
