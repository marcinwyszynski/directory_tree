Package directory_tree
=====================

	import "github.com/marcinwyszynski/directory_tree"

Package directory_tree provides a way to generate a directory tree.

Example usage:

	tree, err := directory_tree.NewTree("/home/me")

type Node
---------

	type Node struct {
	  FullPath string    `json:"path"`
	  Info     *FileInfo `json:"info"`
	  Children []*Node   `json:"children"`
	}

Node represents a node in a directory tree.

### func NewTree

	func NewTree(root string) *Node, error
	
NewTree creates directory hierarchy starting at provided root.

type FileInfo
-------------

	type FileInfo struct {
	  Name    string      `json:"name"`
	  Size    int64       `json:"size"`
	  Mode    os.FileMode `json:"mode"`
	  ModTime time.Time   `json:"mod_time"`
	  IsDir   bool        `json:"is_dir"`
	}

FileInfo is a struct created from os.FileInfo interface for serialization.