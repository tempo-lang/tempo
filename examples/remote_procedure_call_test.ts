import { assertEquals } from "@std/assert";
import { simulate } from "../typescript/simulator.ts";
import {
  Start_A,
  Start_B,
} from "./remote_procedure_call/remote_procedure_call.ts";

Deno.test("simulate remote procedure call", async () => {
  const result = await simulate({
    A: Start_A,
    B: Start_B,
  });

  assertEquals(result, {
    A: {
      receives: [{ sender: "B", value: 20 }],
      return: undefined,
      sends: [{ receivers: ["B"], value: 10 }],
    },
    B: {
      receives: [{ sender: "A", value: 10 }],
      return: undefined,
      sends: [{ receivers: ["A"], value: 20 }],
    },
  });
});
