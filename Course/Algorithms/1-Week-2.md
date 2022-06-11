## 1. Stacks

1. Stack API

2. Implementation: using linked-list
   
	   inner class:
		private class node {
			String item;
			Node next;
		}
		- save a link to the list:
		Node oldfirst = first;
		- create a new node for the begining
		first = new Node();
		- set the instance variable in the new node;
		first.item = "not";
		first.next = oldfirst;
		
		public class LinkedStackofStrings {
			private Node first = null;
			
			private class node {
				String item;
				Node next;
			}
			
			public boolean isEmpty() {
				return first == null;
			}
			
			public void push(String item) {
				Node oldfirst = first;
				first = new Node();
				first.item = item;
				first.next = oldfirst;
			}
			
			public String pop() {
				String item = first.item;
				first = first.next;
				return item;
			}
		}
		

3. Stack: array implementation

		public class FixedCapacityStackOfStrings {
			private String[] s;
			private int N = 0;
			
			public FixedCapacityStackOfStrings(int capacity) {
				s = new String[capacity];
			}
			
			public boolean isEmpty() {
				return N == 0;
			}
			
			public void push(String item) {
				s[N++] = item;
			}
			
			public String pop() {
				return s[---N];
			}
		}
		// Underflow(throw exception if pop from an empty stack) vs Overflow(use resizing array)
		
4. Resizing Arrays -- Using stack resizing-array

1. How to grow an array?
2. How to shrink array? push or pop(when array is one quarter full)
		
		
## 2. Queues  -- LinkedList implementation  || -- resizing array implementation

## 3. Generics
   Q: How to implement  
1. Implement a separate stack class for each type.  --> not effieicent
2. Implement a stack with items of type Object.  --> cast
3. Java Generics
	1. Avoid casting in client
	2. Discover type mismatch errors at comile-time instead of run-time

## 4. Iterators
1. What is an Iterable?
	1. Has a method that returns an Iterator.

2. What is an Iterator?
	1. Has methods hasNext() and next().

3. Why make data structures Iterable?
	1. Java supports elegant client code.


## 5. Sort
1. Callback: reference to executable code.
	1. Client passes array of objects to sort() function.
	2. The sort() function calls back object's __compareTo()__ method as needed.

2. Goal: sort any type of data.
3. Total Order: is a binary relation <= that satisfies
	 - Antisymmetry
	 - Transitivity 
	 - Totality

4. Comparable API, implement compareTo() method
	- Less() function
	- Exch() method


## 6. Selection Sort
- In iteration i, find index min of smallest remaining entry.
- Swap a[i] and a[min].
- 1/2 N^2 compares and n exchanges
## 7. Insertion Sort
- 1/4 N^2 compares and 1/4 N^2 exchanges(average), both 1/2 worst, best N - 1 compare and 0 exchanges

## 8. Shell Sort
- Move entries more than one position at a time by h-sorting the array.
 
## 9. Shuffling
- Knuth shuffling --> liner time complexity 