/*

посылка 51377906

-- ПРИНЦИП РАБОТЫ --
	Производим бинарный поиск по заданному ключу, рекурсивно вызывая данный алгоритм на каждом шаге.
Нетривиальный базовый случай - искомый элемент найден. В таком случае в правом поддереве найдем наименьший элемент,
который является самым левым листом (на английском языке такой элемент называют successor). Далее превращаем
саксессор в вершину поддерева путем переопределения ссылок на правый и левый сыновья.
При выходе из рекурсии заменяем удаляемый элемент на вершину получившегося поддерева.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Процесс "поднятия" саксессора на вершину поддерева формирует BST. После замены удаляемого элемента на
новую вершину BST-поддерева, изначальное дерево не потеряет свойств BST, т.к. BST-поддерево состоит из элементов
исходного дерева. В результате имеем BST дерево с отсутствующим удаляемым элементом - что и требует задача.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	На каждом шаге рекурсии, а также при поиске саксессора мы опускаемся на один уровень вглубь дерева. Таким
образом в худшем случае временная сложность алгоритма равна O(H), где H - высота дерева.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
	Т.к. не создается вспомогательных структур данных, пространственная сложность равна O(1).

*/

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

		return successorToRoot(root)
	}

	return root
}

func successorToRoot(root *Node) *Node {
	successor := root.right
	for successor.left != nil {
		successor = successor.left
	}

	successor.left = root.left

	return root.right
}
