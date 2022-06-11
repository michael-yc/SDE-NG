## Data Structure and Algorithms Overview

|Topic|Data structures adn algorithms|
|------|-------|
|data types|stack, queue, bag, union-fond, priority queue|
|sorting|quicksort, mergesort, heapsort, radix sorts|
|searching|BST, read-black BST, hash table|
|graphs|BFS, DFS, Prim, Kruskal, Dijkstra|
|strings|KMP, regular expressions, TST, Huffman, LZW|
|advanced|B-tree, suffix array, maxflow|

## Undirected graphs
1. Set of vertices connected pairwise by edges.
	1. Path: Sequence of vertices connected by edges.
	2. Cycle: Path whose first and last vertices are the same.
	3. Problems? path(shortest path) | cycle(Euler tour | Hamilton tour) | connectivity(MST | Biconnectivity) | planarity(Graph isomorphism)
2. Graph API 
3. Graph representation: Use adjacency-lists representation.
4. Depth-First Search
	+ Goal: Systematically search through a graph.
	+ Mimic maze exploration.
		* DFS(ro visit a vertex v)
		* Mark v as visited.
		* Recursively visit all unmarked vertices w adjacent to V.
	- Typical Applications:	
		- Fina all vertices connected to a given source vertex.
		- Find a path between two vertices.
		
	- Design pattern: Decouple graph data type from graph processing.
		+ Create a Graph Object.
		+ Pass the Graph to graph-processing routine.
		+ Query the graph-processing routine for information.
		
	- Proposition: DFS marks all vertices vonnected to __s__ in time proportional to the sum of their degrees. 
	- After DFS, can find vertices connected to s in constant time and can find a path to s(if one exists) in time proportional to its length.
	
5. Breadth-First Search  
- Queue
- Repeat until queue is empty:
	+ Remove vertex _v_ from queue.
	+ Add to queue all unmarked vertices adjacent to _v_ and mark them.

6. Difference:

|DFS|BFS|
|---|---|
|Put unvisited vertices on a stack|Put unvisited vertices on queue|

- Shortest path: Find path from _s_ to _t_ that uses fewest number of edges.
- BFS(from source vertex s)
	+ Put _s_ onto a FIFO queue, and mark _s_ as visited.
	+ Repeat until the queue is empty:
		* remove the least recently added vertex v
		* add each of v's unvisited neighbors to the queue, and mark them as visited.
		
- Intuition: BDS examines vertices in increasing distance from _s_.
- Proposition: BFS computes shortest paths(fewest number of edges) from _s_ to all other vertices in a graph in time proportional to __E + V__.

7. Connected components

- Defination: Vertices _v_ and _w_ are connected if there is a path between them.
- Goal: Preprocess graph to answer queries of the fom is _v_ connected to _w_ in constant time.
- DFS solution.
## Directed graphs