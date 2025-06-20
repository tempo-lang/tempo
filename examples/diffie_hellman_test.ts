import { assertEquals } from "@std/assert";
import { simulate } from "../typescript/simulator.ts";
import {
  DiffieHellman_A,
  DiffieHellman_B,
  Math_A,
} from "./diffie_hellman/diffie_hellman.ts";
import { Env } from "../typescript/runtime.ts";

Deno.test("simulate diffie hellman", async () => {
  const mathImpl: Math_A = {
    Exp(_: Env, base: number, exp: number): Promise<number> {
      return Promise.resolve(base ** exp);
    },
  };

  const result = await simulate({
    A: async (env) => {
      return await DiffieHellman_A(env, mathImpl);
    },
    B: async (env) => {
      return await DiffieHellman_B(env, mathImpl);
    },
  });

  assertEquals(result, {
    A: {
      receives: [{ sender: "B", value: 10 }],
      return: { A: 18 },
      sends: [{ receivers: ["B"], value: 4 }],
    },
    B: {
      receives: [{ sender: "A", value: 4 }],
      return: { B: 18 },
      sends: [{ receivers: ["A"], value: 10 }],
    },
  });
});
