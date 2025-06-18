export interface Transport {
  send<T>(value: T, ...roles: string[]): Promise<void>;
  recv<T>(role: string): Promise<T>;
}

export class Env {
  private readonly trans: Transport;
  private roleSubst: {
    [key: string]: string;
  };

  constructor(transport: Transport) {
    this.trans = transport;
    this.roleSubst = {};
  }

  async send<T>(value: T, ...roles: string[]) {
    await this.trans.send(value, ...roles);
  }

  async recv<T>(role: string): Promise<T> {
    return await this.trans.recv(role);
  }

  role(name: string): string {
    if (this.roleSubst[name]) {
      return this.roleSubst[name];
    } else {
      return name;
    }
  }

  subst(...roles: [string]): Env {
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
