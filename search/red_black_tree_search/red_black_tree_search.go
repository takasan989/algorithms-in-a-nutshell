package red_black_tree_search

import (
    "algorithms-in-a-nutshell/search/red_black_tree"
)

func Search(tree *red_black_tree.Tree, value int) bool {
    return red_black_tree.Search(tree, value)
}

func NewTree(data []int) *red_black_tree.Tree {
    tree := &red_black_tree.Tree{}
    for _, v := range data {
        tree.Append(v)
    }
    return tree
}