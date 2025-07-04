// Code generated by tempo, DO NOT EDIT.

import { Env } from '../../../typescript/runtime.ts';

// Projection of choreography foo
export async function foo_A(env: Env) {
  let x: number = 10;
  env.send(x, "B");
  let y: number = x;
  let z: number = y;
}
export async function foo_B(env: Env) {
  let y: number = await env.recv<number>("A");
  env.send(y, "C");
  let z: number = y;
}
export async function foo_C(env: Env) {
  let z: number = await env.recv<number>("B");
}

