/**
 * Simulator implements a convenice environment for running processes locally.
 *
 * @module sim
 */

import { Env } from "./runtime.ts";
import { LocalQueue } from "./transports/mod.ts";
import { Recorder, RecvValue, SendValue } from "./transports/record.ts";

/**
 * The configuration of a process given to {@link simulate}.
 */
export type Process = {
  role: string;
  run(env: Env): Promise<any>;
};

/**
 * The result of a simulation.
 */
export type Result = {
  return: any;
  sends: SendValue[];
  receives: RecvValue[];
};

/**
 * Simulates a choreography locally given a list of processes.
 * @param processes a list of all processes in a choreography.
 * @returns the result of each processes in the same order as the processes given as input.
 */
export async function simulate(...processes: Process[]): Promise<Result[]> {
  const queue = new LocalQueue();

  const results: Promise<Result>[] = [];

  for (const proc of processes) {
    results.push(
      (async () => {
        const trans = new Recorder(queue.role(proc.role));
        const ret = await proc.run(new Env(trans));

        return {
          return: ret,
          sends: trans.sends,
          receives: trans.receives,
        };
      })()
    );
  }

  return await Promise.all(results);
}
