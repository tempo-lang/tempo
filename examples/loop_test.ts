import { assertEquals } from "@std/assert";
import { simulate } from "../typescript/simulator.ts";
import { Main_A, Main_B } from "./loop/loop.ts";

Deno.test("simulate loop", async () => {
  const result = await simulate({
    A: Main_A,
    B: Main_B,
  });

  assertEquals(result, {
    A: {
      receives: [
        { sender: "B", value: "ping" },
        { sender: "B", value: "ping" },
        { sender: "B", value: "ping" },
        { sender: "B", value: "ping" },
        { sender: "B", value: "ping" },
      ],
      return: undefined,
      sends: [
        { receivers: ["B"], value: true },
        { receivers: ["B"], value: true },
        { receivers: ["B"], value: true },
        { receivers: ["B"], value: true },
        { receivers: ["B"], value: true },
        { receivers: ["B"], value: false },
      ],
    },
    B: {
      receives: [
        { sender: "A", value: true },
        { sender: "A", value: true },
        { sender: "A", value: true },
        { sender: "A", value: true },
        { sender: "A", value: true },
        { sender: "A", value: false },
      ],
      return: undefined,
      sends: [
        { receivers: ["A"], value: "ping" },
        { receivers: ["A"], value: "ping" },
        { receivers: ["A"], value: "ping" },
        { receivers: ["A"], value: "ping" },
        { receivers: ["A"], value: "ping" },
      ],
    },
  });
});
