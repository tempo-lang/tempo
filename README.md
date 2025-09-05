# Tempo Choreographic Programming Language

[![Go Reference](https://pkg.go.dev/badge/github.com/tempo-lang/tempo.svg)](https://pkg.go.dev/github.com/tempo-lang/tempo)
[![JSR](https://jsr.io/badges/@tempo-lang)](https://jsr.io/@tempo-lang)

Tempo is a practical choreographic programming language that is compiled to Go source code.

The language is under development, up-to-date examples can be found in the [`examples`](./examples/) directory.

Here is a brief list of implemented things.

- [x] Primitive values and types
- [x] Shared variables
- [x] Asynchronous values
- [x] Channels
- [x] Control flow (if and while statements)
- [x] Function calls
- [x] Struct types
- [x] Interfaces
- [x] Closures
- [x] Lists
- [x] Methods
- [ ] Enums
- [ ] Maps
- [ ] Co-routines
- [ ] Unit tests
- [ ] Generics

## Installation

Make sure you have the latest version of Go installed.
Then run the following command to install the tempo compiler.

```sh
$ go install github.com/tempo-lang/tempo@latest
```

If everything went well, you can now run the compiler by executing the `tempo` command.

For a better developer experience, install the language support [extension for vscode](https://marketplace.visualstudio.com/items?itemName=tempo-lang.vscode-tempo).

## Values and Types

Values are statically typed with the addition of roles.

```tempo
struct@(A,B) Pair {
  left: Int@A;
  right: Int@B;
}

let x: Bool@A = true; // local
let y: String@[A,B] = "hello"; // shared
let z: Pair@(A,B) = Pair@(A,B) { left: 1@A, right: 2@B }; // distributed
```

### Asynchronous types

An asynchronous value, indicates that the underlying value is not necessarily present yet.
Use the `await` expression to get the underlying value, which will wait until it arrives before continuing with the underlying value.

Normal types can be coerced into asynchronous types that immediately return the result when `await` is used.

```tempo
let x: async Int@A = 1;
let y: Int@A = await x; // value is already present

let z: async Int@A = 3 + x; // expression will be coerced to async

let callback = func@(A,B) () async Int@A {
  return B->A 10;
};

let list: [async Int@A] = [1, callback(), 3];
await list; // will wait for all elements in the list to finish
```

## Channels

Channels are built-in primitives.
All roles can communicate with each other by writing `A->B` where `A` and `B` are roles.

```tempo
// local value at A.
let x = 10@A;

// send value from A to B.
let y: Int@B = await A -> B x;

// send value from A to B and C to obtain a shared value.
let z: Int@[A,B,C] = await A->[B,C] x;
```

## Shared variables

A value can be co-located at multiple roles and is then called a _shared variable_ and is denoted with square brackets when declaring the type.
A shared variable can be coerced to a shared variable of a subset of the roles or even a single role.
Constant literals are automatically shared between all roles and coerced to the subset needed.

```tempo
let x: Bool@[A,B,C] = true;
let y: Bool[A,B] = x;
```

Shared variables is an alternative to traditional labels for determining knowledge of choice.
Instead, when a choice is made, all participants of the choice will calculate it independently using shared variables.

```tempo
let x: Int@[A,B] = 3;
if x > 0 {
  // both A and B knows choice
}
```

Shared variables can only be mutated by expressions that are of the same shared roles.
This ensures that shared variables always agree on the same value across the roles.

Shared variables can be expanded by sending it to further participants transitively.

```tempo
func@(A,B,C) shareTrans() {
  let x: Int@A = 42;
  let y: Int@[A,B] = await A->B x;
  let z: Int@[A,B,C] = await B->C y;
}
```

## Functions

All choreographies start with a function defined over a set of roles.
Other functions can be called as long as they use a subset of the roles.

```tempo
func@(A,B,C) hello() {
  let hello: String@B = await A->B "Hello";
  let greeting: String@C = await B->C (hello + ", World!");
}
```

If a function exists only at a single role, the role parameters can be omitted.

```tempo
func greet(name: String) {
  print("Hello " + name);
}
```

## Interfaces

Functions from the host language can be called from Tempo through interfaces.

```tempo
interface Printer {
  func print(value: String);
}

func@(A,B) hello(printA: Printer@A, printB: Printer@B) {
  printA.print("Hello from A");
  printB.print("Hello from B");
}
```
