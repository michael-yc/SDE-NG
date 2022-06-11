## 1. Priority queue

1. Priority queue using unordered array vs ordered array
2. Complete binary tree

- Binary tree: Empty or node with links to left and right binary trees.
- Complete tree: Perfectly balanced, except for bottom level.
- Binary heap: Array representation of a heap-ordered complete binary tree.
- Heap-ordered binary tree.
	+ Keys in nodes.
	+ Parent's key no smaller than children's keys.
	
- Array representation.
	+ Indices start at 1.
	+ Take nodes inlevel order.
	+ No explicit links needed!
	
3. Heapsort

- Basic plan for in-place sort.
	+ Create max-heap with all N keys.
	+ Repeatedly remove the maximum key.
	
- Heap construction. Build max heap using bottom-up method.

## 2. Symbol Table API

1. Key-value pair abstraction.

- Insert a value with specified key.
- Given a key, search for the corresponding value.

2. Associative array abstraction
- Conventions:
	+ Values are not null.
	+ Method get() returns null if key not present.
	+ Method put() returns old value with new value.
	
3. Keys and values

- Value type. Any generic type.
- Key type: several natural assumptions.
	+ Assume keys are Comparable, use compareTo().
	+ Assume keys are any generic type, use equals() to test equality.
	- Assume keys are any generic type, use equals() to test equality; use hashCode() to scramble key.
	
4. Elementary Implenentations

- Data structure: Maintain an (unordered) linked list of key-value pairs.
- Search. Scan through all keys until find a match.
- Insert. Scan through all keys until find a match; if no match add to front.

	1. Unordered list
	2. Ordered list
	
	- Binary search
	
5. Binary search trees
	
	- A BST is a binary tree in **symmetric order**.
		+ A binary tree is either empty or two disjoint binary trees(left and right)
		
	- Symmetric order: Each node has a key, and every node's key is: 
		+ Larger than all keys in its left subtree.
		+ Smaller than all keys in its right subtree.
		

    - Search. If less, go left; if greater, go right; if equal, search hit.
	



