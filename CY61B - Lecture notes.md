## Lecture 1 -- Overview
* Reference books
	* Head First Java
* Goals
	* Learn efficient data structure and algorighms
	* Design and write large programs
	* Understanding and designing data abstraction and interfaces
	* Learn Java
* Class, method
* Inheritance: A class can inherit properties from a more general class.
* Polymorphism: One method call works on several classes, even if the classes need different implementations. e.g. .add() in List, shopingList shoppingCart is different.
* Two ways to get class
	* Use one defined by someone else
	* Define one by yourself

* Must declare variable before using it. Must compile it before it run.(.java -> .javac class files(JVM) -> answer)

## Lecture 2 -- Using Objects and Methods
		String s1 = new String("Hello");
		String s2 = new String("Hello");
		s1, s2 is the reference of String "Hello". They are not the same, but they are two different objects.
* 3 String constructors
	* new String()
	* "Whatever"
	* new String(s1)

Constructors always have some names as their class except "Stuffquotes".
 
String method:
* s2 = s1.toUppercase();
* String s3 = s2.toLowercase();
* s.charAt(): Returns the character at the specified index (position);
* s.compareTo(): Compares two strings lexicographically;
* s.concat(): Appends a string to the end of another string;
* s.contains(): Checks whether a string contains a sequence of characters;
* s.equals(): Compares two strings. Returns true if the strings are equal, and false if not;
* s.indexOf(): Returns the position of the first found occurrence of specified characters in a string;
* s.isEmpty(): Checks whether a string is empty or not;
* s.length(): Returns the length of a specified string;
* s.matchs(): Searches a string for a match against a regular expression, and returns the matches
* s.split(): Splits a string into an array of substrings;
* s.substring(): Returns a new string which is the substring of a specified string;
* s.toCharArray(): Converts this string to a new character array;
* s.toString(): Return the value of a String Object;
* s.trim(): Removes whitespace from both ends of a string;

System.out() is a PrintStream object that outputs to the screen.
System.in() os a InputStream object that reads from the keyboard.
ReadLine() method is defined on BufferedReader objects.
How do we construct a BufferedReader(chars into entire lines of text)?
With InputStreamReader.
... InputStreamReader(Compose into characters, typically two types)? With InputStream.
... InputStream?(Read raw data) System.in is one.

		import java.io.*;
		class SimpleIO {
			public static void main(String[] args) throws Exception{
				BufferedReader keyboard = new BufferedReader(new InputStreamReader(System.in));
				System.out.println(keyboard.readLine());
			}
		}
		
How to read files from the webpage?
		
		import java.net.*;
		import java.IO.*;
		class WWWW{
			public static void main(String[] args) throws Exception {
				URL u = new URL("http://www.whitehouse.gov/");
				InputStream ins = u.openStream();
				InputStreamReader isr = new InputStreamReader(ins);
				BufferReader whiteHouse = new BufferedReader(isr);
				System.out.println(whiteHouse.readLine());
			}
		}

## Lecture 3 -- Define Classes
- Fields: 
	- Variables stored in objects.
	- Ussed to store data.
	- Instance variables.
- Constructors:
	- the same name as the class
	- return the same object of the class;
	- Java has its own default constructor;

 - The "this" keyword: create a reference to the object; You can't change the value of "this" which will cause compile-time error.
 - The "static" keyword: static field:
	  - A single variable shared by a whole class of objects
	  - also called class variables.
	  - main method is always static, in a static method, there is no "this".

- Lifetime of variables
	- A local variable is gone when method ends.
	- An instance variable(non-static field) lasts as long as object exists.
	- A class variable(static field) lasts as long as the program.

## Lecture 4 -- Types and Conditions
- Primitive Types: 
	- byte: 8-bit integer from -128 to 127.
	- short: 16-bit integer from -32748 to 32747.
	- int: 32-bit integer from -2147483648 to 2147483647.
	- long: 64-bit integer.
		- long x = 124L;
	- float: 32-bit integer.
		- float f = 18.9f;
	- double: 64-bit float point number.
	- boolean: true or false.
	- char: single character.
	- double & float values must have a decimal point.
		- double y = 18.0;

||Object types|Primitive types|
|-----|------------|-------------------|
|Contains|Reference|Value|
|How defined|Class definition|build into Java|
|How create|"new"|"6", "3.4"|
|how in|constructor|default(usually zero)|
|how used|mehods|operators("+", "-", etc)|

- java.lang library
	- Math class
		- x = Math.abs(y);
		- Math.abs(); Math.min(x, y); Math.random(0, 1); Math.sqrt(x); Math.round(x); Math.
	- Integer class
		- int x = Integer.parseInt("1984");
		- Integer and int difference
			- Integer是int的包装类；int是基本数据类型；
			- Integer变量必须实例化后才能使用；int变量不需要；
			- Integer实际是对象的引用，指向此new的Integer对象；int是直接存储数据值 ；
			- Integer的默认值是null；int的默认值是0。
		- Integer method:
			- Interger.compareTo();
			- Integer.equals();
			- static int parseInt(int i);
			- Integer.toString();
			- static Integer valueOf(int i);
			- Integer.Max_Value, Integer.Min_Value;
	- Double class
		- couble d = Double.parseDouble("3.14");

- Integers can be assigned to variables of longer types.   CAST.

- Conditions
	- if else, while, do while.
	- switch 

- The "return" keyword
	- "return" causes a method to end immediately.

## Lecture 5 -- Iteration and Arrays I
- while loops

		public static boolean isPrime(int n) {
			int divisor = 2;
			while (divisor < n) {
				if (n % divisor == 0) {
					return false;
				}
				divisor ++;
			}
			return true;
		}
		if n <= 2, the loop won't iterate.
		
- for loops

		for (initialize; test; next) {
			statements;
		}
		
- Arrays: An object constisting of a numbered list of variables. Each is a primitive type of reference.
- Pascel Triangle

## Lecture 6 -- Iteration and Arrays II

- Automatic Array Construction
	- int[][] table = new int[x][y];
	- Creates an array of x references to array.

- Initializers: Human[] b = {kayla, rishi, new Human("Paolo)};

- break and continue statements
	- break exits the innermost "switch" or loop enclosing the "break"
	- continue
		- only applies to loops
		- doesn't end loop

- difference for and while loop		
		
		int i = 0;
		while (i < 10) {
			if (condition(i))
				continue;
			call(i);
			i++;
		}
		
		for (int i = 0; i < 10; i++) {
			if (condition(i))
				continue;
			call(i);
		}
		
- Constraints:
	- final keyword: value that can never be changed.

## Lecture 7 -- Linked List I

- List : Store a list of ints as an array
	- Disadvantages: 
		- Insert item to beginning or middle   -->> time proportional to length of array
		- Arrays have a fixed length.
- Example:
		
		public class List {
			int[] a;
			int lastItem;
			
			public List() {
				a = new int[10];
				lastItem = -1;
			}
			
			public void insertItem(int newItem, int location) {
				int i;
				for (int i = lastItem; i >= location; i--) {
					a[i + 1] = a[i];
				}
				a[location] = newItem;
				lastItem++;
			}
		}
- LinkedList: A recursive data type
	- Made up of "ndoes".
	- Each nodes has
		- an item
		- a reference to next node in list
- Code:

		public class ListNode {
			int item;
			ListNode next;
		}
		ListNode l1 = new ListNode();
		ListNode l2 = new ListNode();
		ListNode l3 = new ListNode();
		11.item = 7;
		12.item = 0;
		13.item = 6;
		l1.next = l2;
		l2.next = l3;
		l3.next = null;
		
		Node operations
		public ListNode(int item, ListNode next) {
			this.item = item;
			this.next = next;
		}
		public ListNode (int item) {
			this(item, null);
		}
		
		ListNode l1 = new ListNode(7, new ListNode(0, new ListNode(6)));
		
- Advantages over Array List:
	- inserting item in the middle of linked list takes constant time if you have reference to previous node.
	- moreover, list can keep growing until memory funs out.

- Insert a new item after "this".

		public void insertAfter(int item) {
			next = new ListNode(item, next);
		}
		l2.insertAfter(3);
		
- Disadvantages:
	- Finding the nth item of a linked list takes time proportional to n  -->> length of list. (const-time on array lists)

- Code: 

		public ListNode nth(int position) {
			if (position == 1) {
				return this;
			} else if ((position < 1) || (next == null)) {
				return null;
			} else {
				return next.nth(position - 1);
			}
		}
		
- Lists of Objects
	- Reference any object by declaring a reference of type object. 

- A List Class
	- 2 problems with SListNodes.
		- Insert new item at the beginning of list.
			- e.g. x = new SListNode("soap", x);
		- How do you represent an empty list?
			- x = null; -->>Run-time error if you call a method on a null object.'
				- Solution: Separate SList class mantains head of list.

- Code Implementation:

		public class SList {
			private SListNode head;
			private int size;
			
			public SList() {
				head = null;
				size = 0;
			}
			
			public void insertFront(Object Item) {
				head = new SListNode(item, head);
				size ++;
			}
		}	

## Lecture 8 -- Linked List II

- A private method or field is invisible & inaccessible to other classes.
- Interface of a class: prototypes for public methods, plus descriptions of their behaviors.
- Abstract Data Type(ADT): a class with a well-defined interface, but implementation details are hidden from other classes.
- Invariant: a fact about a data structure that is always true.
- The SList ADT
	- Another advantage of SList class: SList ADT enforces 2 invariants.
		- "Size" is always correct.
		- List is never circularly linked.
	- Both goals accomplished because only SList methods can change the lists.

- SList ensures this: 
	- The fields of SList(head and size) are "private".
	- No method of SList returns on SListNode.

- Doubly-Linked Lists
	- Insert/deleting at front of list is easy.
	- Inserting/deleting at the end of list takes a long time.

- Code implementation:

		class DlistNode {
			Object item;
			DListNode next;
			DListNode prev;
		}
		
		class DList {
			private DListNode head;
			private DListNode tail;
		} 
		
- insert & delete items at both ends in constant running time.
- Removes the tail node (at least 2 items in DList) 
	- tail.prev.next = null;
	- tail = tail.prev;

- Sentinel: A special node that does not represent an item.
- DList  vs  Doubly-Linked List
- DList invariants with sentinel:
	-  For any DList d, d.head != null.
	-  For any DList x, x.next != null.
	-  x.prev != null
	-  x.next == y, then y.prev == x
	-  x.prev == y, then y.next == x
	-  A DList's "size" variable is # of DListNodes, not counting sentinel.
	-  Empty DList: sentinels prev & next field point to itself.


## Lecture 9 -- Stack Frames

- The Stack and The Heap
	- The heap stores all objects, including arrays, and all class variables.
	- The stack stores all local variables, including parameters.

- When a method is called, Java creates a stack frame, sotes the parameters & local variables.  >> stack of stack frames.
- Illustrate the deep understanding of stack and heap.
- When method finishes, its stack frame erased. -->> Method: Thread.dumpstack() in java.lang library.
- Parameter passing: All parameters are passed by value.
  
		 int a = 1;
		 doNothing(a);
		 private void doNothing(int x) {
			 x = 2;
		 }
		 
		 
		 When A parameter is a reference, the reference is copied, but not the thing the object do.
		 IntBox b = new IntBox();
		 set3(b); --> 3
		 badSet4(b); --> 3
		 
		 public int i;
		 static void set3(IntBox ib) {
			 ib.i = 3;
		 }
		 
		 static void badSet4(IntBox ib) {
			  ib = new IntBox(4);
			  ib.i = 4;
		 }
		 
- Binary Search 
	- Recursion base cases: 
		- findMe = middle element, return the element: return its index
		- Subarray of length zero: return failure.
		- How fast? O(logn)  

- Enough stack space for a few thousands of stack frames.
- Scope and Recursion
	- Scope of a variable: portion of program that can access the variable.
		- Class variables: in scope everywhere in the class, except when shadowed by local variables
		- Fully qualified class variable("System.out") in scope in the class, cannot be shadowed. If public, they are in scope in all classes.
		- Instance variables: in scope in non-static methods of the class, except when shadowed.
		- Fully qualified instance variables("Kayla.name", "this.item") same as 2
		- Local variables & parameters: only in scope in the method that defines them, only for topmost stack frame.

## Leture 10 -- Testing

- equals() -- Every class has an equals() method.
	- Default: r1.equals(r2) same as r1 == r2
	- Many classes define equals() to compare content.
	- e.g. String.
		- s1 --> yes   s2-->yes
		- s1 == s2 -->> false
		- s1.equals(s2) -->> true

- 4 degrees of equality:
	- Reference equality; ==
	- Shallow structural equality; fields are ==
	- Deep structural equality; fields are equals()
	- Logical equality
		- Fractions 1/3 and 3/6 are "equals"
		- "Set" objects are equals if they contain some elements(even stored in differnent orders). 

- Code
	- Test deep structural equality

			public boolean equals(SList other) {
				if (size != other.size) {
					return false;
				}
			
				SListNode n1 = head;
				SListNode n2 = other.head;
			
				while (n1 != null) {
					if (!n1.item.equals(n2.item)) {
						return false;
					}
					n1 = n1.next;
					n2 = n2.next;
				}
			}
			
- Testing 
	- Modular testing
		- Test drives & stubs
		- Test drivers call the code, check the results.

	- Stubs: bits of code called by the code being tested.
		- Fill in for missing method.
		- Determine whether bug lies in calling method or callee by replacing callee with stub
		- Produce test data that real subroutines rarely produce  |  Produces repeatable test data.

	- Integration Testing
		- Define interfaces well! No ambiguity in descriptions of behaviors.
			- Learn to use the debugger.

	- Result verification
		- Data structure integrity checkers.
			- Inspects data structure & verifies that all invariants are satisfied.
		- Algorithm result checker.
			- Sorting -->> check that each result <= its successor.

	- Assertion: Code that tests an invariant or result.

	- Regression Testing: test suite that can be re-run when changes are made.

## Lecture 11 Inheritance

- A subclass can modify a superclass in 3 ways.
	- It can declare new fields.
	- New methods
	- It can override old methods with new implementations.

- TailList is the subclass of SList.
- Code:

		public TailList() {
			//SList(); set size = 0, head = null;
			tail = null;
		}
		To change this:
		public TailList(int x) {
			super(x);   // must be first statement in constructor.-->> SList(x);
			tail = null;
		}
		public void insertFront(Object obj) {
			super.insertFront(obj);
			if (size == 1) {
				tail = head;
			}
		}
		
		
- The "protected" keyword
- Code

		public class SList {
			protected SListNode head;
			protected int size;
		}
		"protected" field/method is visible to declring class & all its subclasses.
		"private" fields aren't visible to subclasses.
		
- __Every TailList Is an SList__
	- SList s = new TailList();
	- TailList t = new SList(); -->> Compile time error.
		- static type: the type of a variable.
		- dynamic type: the class of the object/variable reference.
	- When we invoke overriden method, Java calls method for the object's dynamic type, regardless of static type.
	- SList s = new TailList();
	- s.insertEnd(obg); // calls TailList.inserEnd();
	- s = new SList();
	- s.insertEnd(obj); // calls SListinsertEnd();
	- Why D.M.L matters?
	- Method that sorts an SList using only method calls. Now TailLists as well.


## Lecture 12 -- Abstract Classes

- New method in TailList called eatList().
	- TailList t = new TailList();
	- t.eatList();
		- SList s = new TailList();
		- s.eatList(); //compile-time error.
		- Why?
			- Not every SList has an "eatList()" method.
			- Java can't use dynamic method look up on s.
	- Slist s;
	- TailList t = new TailList();
		- s = t;
		- t = s;  // compile-time error.
		- t = (TailList) s;
		- s = new SList();
		- t = (TailList) s; //Run-time error.
	- "instanceof" operator tells you whether object is of specific class.
		- code
				
				if (s instanceof TailList) {
					t = (TailList)s;
				}
				false, s is null or does not reference to a TailList or subclass.

- Abstract classes:
	- A class whose sole purpose is to be implemented.
		- Code

				public abstract class List {
					protected int size;
					public int length() {
						return size;
					}
					public abstract void insertFront(object item);
				}
				List myList;
				myList = new List();  -->> Compile-time error
				
- __Abstract methods lacks implementation.__

- A non-abstract class may never
	- contain an abstract method,
	- inherit one without providing an implementation
		- List myList = new SList();
		- myList.insertFront();

- Interface vs abstract class
	- A class can inherit from only are class, but can "implement"(inherit from) as many Java interfaces as you like.
	- A java interface cannot be new;
		- implement any methods,
		- include only fields except "static final" constraints. Only contains methos prototypes & constants.

## Lecture 13 -- Java Packages

- Java packages
	- Package: Colletion of classes, Java interfaces, & subpackages.
	- 3 benefits:
		- Packages can contain hidden classes not visible outside package.
		- Classes can have fields & methods visible inside package only.
		- Different packages can have classes with same name.
	- 2 examples:
		- java.io standard library.
		- List package containing DList & DListNode.

- Using packages
	- Every program implicitly imports java.lang.*
	- A class/variable with package protection is visible to any class in same package, not outside package(files outside directory).
	- Files outside only see public classes/methods/fields.

- Compiling & running must be done from outside package.
	- javac -g   list/SList.java
	- java   list.SList

## Lecture 14 -- Exceptions

- Run-time error: Java throws an exception.
- Prevent the error by "catching" the exception.
	- Purpose #1: Surviving errors.
	- Code explanation about open a file and read with try and catch.
		- Executes the code inside "Try".
		- If "try" code executes normally, skip "catch" clauses.
		- If "try" code throws exception, do not finish "try" code. Jump to first "catch" clause that matches exception, execute " Matches" exception object thrown is some class/subclass of type in "catch" clause.