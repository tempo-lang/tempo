# Tempo Choreographic Programming Language

Tempo is a practical choreographic programming language that is compiled to Go source code.

The language is under development, here is a brief list of implemented things.

- [x] Primitive values and types
- [x] Shared variables
- [x] Asynchronous values
- [x] Channels
- [x] If statements
- [x] Function calls
- [x] Struct types
- [x] Interfaces
- [ ] Closures
- [ ] Lists
- [ ] Maps
- [ ] Methods
- [ ] Co-routines

## Installation

Make sure you have the latest version of Go installed.
Then run the following command to install the tempo compiler.

```sh
$ go install github.com/tempo-lang/tempo@latest
```

If everything went well, you can now run the compiler by executing the `tempo` command.

## Values and Types

Values are statically typed with the addition of roles.

```
struct@(A,B) Pair {
  left: Int@A
  right: Int@B
}

let x: Bool@A = true
let y: String@[A,B] // shared
let z: Pair@(A,B) // distributed
```

### Asynchronous types

An asynchronous value, indicates that the underlying value is not necessarily present yet.
To get the underlying value the `await` expression is used, which will wait for the value to arrive before continuing with the underlying value.

Normal types can be coerced into asynchronous types that immediately return the result when `await` is used.

```
let x: async Bool@A = true
let y: Bool@A = await x // value is already present

let z: async Bool@A = 3 + x // expression will be coerced to async

func@(A,B) callback() async Int@A {
  return B->A 10;
}

let list: async@(A,B) [Int]@A = [1, callback(), 3]
await list // will wait for all elements in the list to finish
```

## Channels

Channels are built-in primitives.
All roles can communicate with each other by writing `A->B` where `A` and `B` are roles.

```
// send value from A to B.
let y@B = await A -> B x@A

// send value from A to B and C to obtain a shared value.
let y@[A,B,C] = await A->[B,C] x@A
```

## Shared variables

A value can be co-located at multiple roles and is then called a _shared variable_ and is denoted with square brackets when declaring the type.
A shared variable can be coerced to a shared variable of a subset of the roles or even a single role.
Constant literals are automatically shared between all roles and coerced to the subset needed.

```
let x: Bool@[A,B,C] = true
let y: Bool[A,B] = x
```

Shared variables is an alternative to traditional labels for determining knowledge of choice.
Instead, when a choice is made, all participants of the choice will calculate it independently using shared variables.

```
let x: Int@[A,B] = 3
if x > 0 {
  // both A and B knows choice
}
```

Shared variables can only be mutated by expressions that are of the same shared roles.
This ensures that shared variables always agree on the same value across the roles.

Shared variables can be expanded by sending it to further participants transitively.

```
func@(A,B,C) shareTrans() {
  let x: Int@A = 42
  let y: Int@[A,B] = await A->B x
  let z: Int@[A,B,C] = await B->C y
}
```

## Functions

All choreographies start with a function defined over a set of roles.
Other functions can be called as long as they use a subset of the roles.

```
func@(A,B,C) hello() {
  let hello: String@B = A->B "Hello"
  let greeting: String@C = B->C (hello + ", World!")
}
```

If a function exists only at a single role, the role parameters can be omitted.

```
func greet(name: String) {
  print("Hello " + name)
}
```
