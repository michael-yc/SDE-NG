## Metaprogramming

1. What is MetaProgramming?
- it was the best collective term we could come up with for the set of things that are more about process than they are about writing code or working more efficiently.
2. Semantic Versioning:
	1. 8.1.2 -->> 8: major   1:minor    7: patch
- If a new release does not change the API, increase the patch version.
- If you add to your API in a backwards-compatible way, increase the minor version.
- If you change the API in a non-backwards-compatible way, increase the major version.

3. When working with dependency management systems, you may also come across the notion of lock files. A lock file is simply a file that lists the exact version you are currently depending on of each dependency. 
4. An extreme version of this kind of dependency locking is vendoring, which is where you copy all the code of your dependencies into your own project. That gives you total control over any changes to it, and lets you introduce your own changes to it, but also means you have to explicitly pull in any updates from the upstream maintainers over time.
5. Continuous integration systems
6. Test:
- Test suite: a collective term for all the tests
- Unit test: a “micro-test” that tests a specific feature in isolation
- Integration test: a “macro-test” that runs a larger part of the system to check that different feature or components work together.
- Regression test: a test that implements a particular pattern that previously caused a bug to ensure that the bug does not resurface.
- Mocking: to replace a function, module, or type with a fake implementation to avoid testing unrelated functionality. For example, you might “mock the network” or “mock the disk”.