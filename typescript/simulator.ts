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
export type Processes = {
  [key: string]: (env: Env) => Promise<any>;
};

/**
 * The result of a single process in a simulation.
 */
export type Result = {
  return: any;
  sends: SendValue[];
  receives: RecvValue[];
};

/**
 * The result of a simulation, indexed by each role.
 */
export type Results = {
  [key: string]: Result;
};

/**
 * Simulates a choreography locally given a list of processes.
 * @param processes a set of all processes in a choreography.
 * @returns the result of each processes in the same order as the processes given as input.
 */
export async function simulate(processes: Processes): Promise<Results> {
  const queue = new LocalQueue();

  const results: Promise<Result & { role: string }>[] = [];

  for (const role in processes) {
    results.push(
      (async () => {
        const trans = new Recorder(queue.role(role));
        const ret = await processes[role](new Env(trans));

        return {
          role,
          return: ret,
          sends: trans.sends,
          receives: trans.receives,
        };
      })()
    );
  }

  const result: Results = {};

  for await (const res of results) {
    const { role, ...rest } = res;
    result[role] = rest;
  }

  return result;
}
