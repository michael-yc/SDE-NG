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
		