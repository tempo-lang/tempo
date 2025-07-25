import { assertEquals } from "@std/assert";
import { simulate } from "../typescript/simulator.ts";
import { Start_A, Start_B } from "./ping_pong/ping_pong.ts";

Deno.test("simulate ping pong", async () => {
  const result = await simulate({
    A: Start_A,
    B: Start_B,
  });

  assertEquals(result, {
    A: {
      receives: [
        { sender: "B", value: 3 },
        { sender: "B", value: 1 },
      ],
      return: undefined,
      sends: [
        { receivers: ["B"], value: 4 },
        { receivers: ["B"], value: 2 },
      ],
    },
    B: {
      receives: [
        { sender: "A", value: 4 },
        { sender: "A", value: 2 },
      ],
      return: undefined,
      sends: [
        { receivers: ["A"], value: 3 },
        { receivers: ["A"], value: 1 },
      ],
    },
  });
});
