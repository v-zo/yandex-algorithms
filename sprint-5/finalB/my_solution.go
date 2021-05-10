package main

func remove(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.value {
		root.left = remove(root.left, key)
	} else if key > root.value {
		root.right = remove(root.right, key)
	} else {
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		}

		rightSmallest := root.right
		for rightSmallest.left != nil {
			rightSmallest = rightSmallest.left
		}

		rightSmallest.left = root.left

		return root.right
	}

	return root
}
