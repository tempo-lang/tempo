// Code generated by tempo, DO NOT EDIT.

import { Env } from '../../../typescript/runtime.ts';

// Projection of interface Sum
export interface Sum_X {
  sum(env: Env, a: number, b: number): Promise<number>;
}

// Projection of choreography foo
export async function foo_A(env: Env, x: Sum_X) {
  let value: number = await x.sum(env.subst("A", "X"), 10, 20);
}

