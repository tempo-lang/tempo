import { Transport } from "../runtime.ts";

/**
 * A value that has been sent to a set of receivers.
 */
export type SendValue = {
  value: unknown;
  receivers: string[];
};

/**
 * A value received by a given sender.
 */
export type RecvValue = {
  value: unknown;
  sender: string;
};

/**
 * Recoder wraps another {@link Transport} implementation and records every message sent and received.
 */
export class Recorder implements Transport {
  #inner: Transport;
  #sends: SendValue[];
  #receives: RecvValue[];

  constructor(transport: Transport) {
    this.#inner = transport;
    this.#sends = [];
    this.#receives = [];
  }

  get sends(): SendValue[] {
    return this.#sends;
  }

  get receives(): RecvValue[] {
    return this.#receives;
  }

  send<T>(value: T, ...roles: string[]): void {
    this.#sends.push({
      receivers: roles,
      value,
    });

    this.#inner.send(value, ...roles);
  }

  async recv<T>(role: string): Promise<T> {
    const value: T = await this.#inner.recv(role);

    this.#receives.push({
      sender: role,
      value,
    });

    return value;
  }
}
