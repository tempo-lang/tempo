/**
 * The transport module provides a set of default {@link Transport} implementations.
 *
 * @module
 */
import { Transport } from "../runtime.ts";

export * from "./record.ts";

/**
 * A local queue that is used to obtain {@link Transport} implementations for processes to communicate locally.
 */
export class LocalQueue {
  #channels: { [key: string]: LocalChan };

  /**
   * Constructs a new local queue.
   */
  constructor() {
    this.#channels = {};
  }

  /**
   * Get the {@link LocalChan} between `from` and `to`.
   * @param from The sending role to the channel.
   * @param to The role receiving from the channel.
   * @returns The channel between `from` and `to`.
   */
  get(from: string, to: string): LocalChan {
    const key = from + "." + to;

    const chan = this.#channels[key];
    if (chan) {
      return chan;
    } else {
      const newChan = new LocalChan();
      this.#channels[key] = newChan;
      return newChan;
    }
  }

  /**
   * Get a transport implementation for a specific role.
   * @param role The role of which the transport is in context of.
   * @returns The transport implementation.
   */
  role(role: string): Transport {
    return new LocalTransport(role, this);
  }
}

type ChanPromise<T> = {
  promise: Promise<T>;
  resolve: (value: T) => void;
};

/**
 * A local directed channel from a specific role to another.
 */
class LocalChan {
  #sendBuf: Promise<any>[];
  #recvBuf: ChanPromise<any>[];

  constructor() {
    this.#sendBuf = [];
    this.#recvBuf = [];
  }

  send<T>(value: T) {
    const recvProm = this.#recvBuf.shift();
    if (recvProm) {
      recvProm.resolve(value);
    } else {
      this.#sendBuf.push(Promise.resolve(value));
    }
  }

  recv<T>(): Promise<T> {
    const sendProm = this.#sendBuf.shift();
    if (sendProm) {
      return sendProm;
    } else {
      let resolveCallback: (value: unknown) => void = () => {};
      const promise = new Promise((resolve) => {
        resolveCallback = resolve;
      });

      this.#recvBuf.push({
        resolve: resolveCallback,
        promise,
      });

      return promise as Promise<T>;
    }
  }
}

/**
 * Implementation of the {@link Transport} interface for a locally communicating process.
 */
class LocalTransport implements Transport {
  #role: string;
  #queue: LocalQueue;

  constructor(role: string, queue: LocalQueue) {
    this.#role = role;
    this.#queue = queue;
  }

  send<T>(value: T, ...roles: string[]) {
    for (const receiver of roles) {
      const chan = this.#queue.get(this.#role, receiver);
      chan.send(value);
    }
  }

  recv<T>(sender: string): Promise<T> {
    return this.#queue.get(sender, this.#role).recv();
  }
}
