## Union find

1. Goal: Design efficient data structure for union-find.
	1. Number of objects N can be huge.
	2. Number of operations M can be huge.
	3. Find queries and union commands may be intermixed.
	
2. Quick-find
- Data structure:
	- Integer array id[] of size N.
	- Interpretation: p and q are connected if they have the same id.
	
3. Quick-union
- Dta structure:
	+ Integer array id[] of size N.
	+ Interpretation: id[i] is parent of i.
	+ Root of i is is[id[...id[i]...]]].
4. Improvement: weighing

- Weighted quick-union.
	+ Modify quick-union to avoid all trees.
	+ Keep track of size of each tree(number of objects).
	+ Balance by linking root of smaller tree to root of larger tree.
	
- Path compression
- Weighted QU + Path compression

## Analysis of Algorithms Introduction
#### Time Complexicity
1. Tilde notation
2. O(1)<O(logn)<O(n)<O(nlogn)<O(n^2)<O(n^3)<O(2^N)
- Best case:
- Worest case:
- Average case:
Three asymptotic notation

#### Space Complexicity

1. Bit: 0 or 1.
2. Byte: 8 bits.
3. Megabyte: 1 million or 2^20 bytes.
4. Gigabyte: 1 billion or 2^30 bytes.
5. boolean/byte: 1 byte
6. char: 2 bytes
7. int/float: 4 bytes
8. long/double: 8 bytes