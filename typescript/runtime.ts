/**
 * This module contains the runtime used by Tempo generated source code.
 *
 * @module runtime
 */

/**
 * The Transport interface specifies the methods needed in order for processes in a choreography to communicate.
 *
 * The `@tempo-lang/tempo/transports` module contains a set of default implementations.
 */
export interface Transport {
  send<T>(value: T, ...roles: string[]): void;
  recv<T>(role: string): Promise<T>;
}

/**
 * The environment class keeps track of the state of the choreography at a single process.
 * The first argument of any generated process function will take an `Env` as its first argument.
 * It is used to interact with the environment such as sending and receiving messages.
 */
export class Env {
  private readonly trans: Transport;
  private roleSubst: {
    [key: string]: string;
  };

  /**
   * constructs a new environment to be passed to a process.
   * @param transport The transport implementation used for sending and receiving messages.
   */
  constructor(transport: Transport) {
    this.trans = transport;
    this.roleSubst = {};
  }

  /**
   * Send will use the underlying {@link Transport} implementation to send the value.
   * @param value The value to send.
   * @param roles The roles to receive the value.
   */
  async send<T>(value: T, ...roles: string[]) {
    const subRoles = roles.map((r) => this.role(r));
    await this.trans.send(value, ...subRoles);
  }

  /**
   * Receive will use the underlying {@link Transport} implementation to receive a value.
   * @param role The role to receive the value from.
   * @returns The received value.
   */
  async recv<T>(role: string): Promise<T> {
    const subRole = this.role(role);
    return await this.trans.recv(subRole);
  }

  /**
   * Maps a static role name to the name substituted in the invocation of the current function.
   * @param name The static role.
   * @returns The substituted role.
   */
  role(name: string): string {
    if (this.roleSubst[name]) {
      return this.roleSubst[name];
    } else {
      return name;
    }
  }

  /**
   * Substitute will return a copy of the environment with a new role substitution map.
   * @param roles A pair-wise list of substitutions.
   * @returns A copy of this environment with the new substitution.
   */
  subst(...roles: string[]): Env {
    const newSub: { [key: string]: string } = {};
    for (let i = 0; i < roles.length; i += 2) {
      const old = roles[i];
      const newRole = roles[i + 1];
      newSub[newRole] = this.role(old);
    }

    const copy = new Env(this.trans);
    copy.roleSubst = newSub;
    return copy;
  }
}
