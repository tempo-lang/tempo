// Code generated by tempo, DO NOT EDIT.

import { Env } from '../../typescript/runtime.ts';

// Projection of choreography compose
export async function compose_A(env: Env, f: (env: Env, arg0: number) => Promise<void>): Promise<(env: Env, arg0: number) => Promise<void>> {
  return async (env: Env, input: number) => {
    await f(env, input);
  };
}
export async function compose_B(env: Env, f: (env: Env) => Promise<number>, g: (env: Env, arg0: number) => Promise<void>): Promise<(env: Env) => Promise<void>> {
  return async (env: Env) => {
    await g(env, await f(env));
  };
}
export async function compose_C(env: Env, g: (env: Env) => Promise<number>): Promise<(env: Env) => Promise<number>> {
  return async (env: Env): Promise<number> => {
    return await g(env);
  };
}

// Projection of choreography incAndSend
export async function incAndSend_X(env: Env, value: number) {
  env.send(value + 1, "Y");
}
export async function incAndSend_Y(env: Env): Promise<number> {
  return await env.recv("X");
}

// Projection of choreography Start
export async function Start_A(env: Env, input: number) {
  let f: (env: Env, arg0: number) => Promise<void> = async (env: Env, value: number) => {
    await incAndSend_X(env.subst("A", "X", "B", "Y"), value);
  };
  let c: (env: Env, arg0: number) => Promise<void> = await compose_A(env.subst("A", "A", "B", "B", "C", "C"), env.copy(f));
  await c(env, input);
}
export async function Start_B(env: Env) {
  let f: (env: Env) => Promise<number> = async (env: Env): Promise<number> => {
    return await incAndSend_Y(env.subst("A", "X", "B", "Y"));
  };
  let g: (env: Env, arg0: number) => Promise<void> = async (env: Env, value: number) => {
    await incAndSend_X(env.subst("B", "X", "C", "Y"), value);
  };
  let c: (env: Env) => Promise<void> = await compose_B(env.subst("A", "A", "B", "B", "C", "C"), env.copy(f), env.copy(g));
  await c(env);
}
export async function Start_C(env: Env): Promise<number> {
  let g: (env: Env) => Promise<number> = async (env: Env): Promise<number> => {
    return await incAndSend_Y(env.subst("B", "X", "C", "Y"));
  };
  let c: (env: Env) => Promise<number> = await compose_C(env.subst("A", "A", "B", "B", "C", "C"), env.copy(g));
  return await c(env);
}

