import { assertEquals } from "@std/assert";
import { simulate } from "../typescript/simulator.ts";
import {
  Authenticate_Client,
  Authenticate_IP,
  Authenticate_Service,
  ClientRegistry,
  Credentials,
  Hasher,
  TokenGenerator,
  AuthResult_S,
  AuthResult_C,
} from "./dist_auth/dist_auth.ts";
import { Env } from "../typescript/runtime.ts";
import { crypto } from "@std/crypto";
import { encodeBase64 } from "@std/encoding";
import { randomSeeded, Prng } from "@std/random";

class MockHasher implements Hasher {
  hashes: string[];

  constructor() {
    this.hashes = [];
  }

  async CalcHash(_: Env, salt: string, password: string): Promise<string> {
    const buf = new TextEncoder().encode(salt + password);
    const hashBytes = await crypto.subtle.digest("SHA-256", buf);
    const hash = encodeBase64(hashBytes);
    this.hashes.push(hash);
    return hash;
  }
}

class MockClientRegistry implements ClientRegistry {
  checkCalls: string[];
  getSalts: string[];

  constructor() {
    this.checkCalls = [];
    this.getSalts = [];
  }

  async GetSalt(_: Env, username: string): Promise<string> {
    const buf = new TextEncoder().encode(username);
    const hashBytes = await crypto.subtle.digest("SHA-256", buf);
    const hash = encodeBase64(hashBytes).slice(0, 8);
    this.getSalts.push(hash);
    return hash;
  }

  Check(_: Env, hash: string): Promise<boolean> {
    this.checkCalls.push(hash);
    return Promise.resolve(true);
  }
}

class MockTokenGenerator implements TokenGenerator {
  tokens: string[];
  rng: Prng;

  constructor() {
    this.tokens = [];
    this.rng = randomSeeded(1n);
  }

  GenerateToken(_: Env): Promise<string> {
    const number = Math.floor(this.rng() * 1000000);
    const token = number + "";
    this.tokens.push(token);
    return Promise.resolve(token);
  }
}

Deno.test("simulate distributed authentication", async () => {
  const credentials: Credentials = {
    Username: "username",
    Password: "password",
  };

  const hasher = new MockHasher();
  const registry = new MockClientRegistry();
  const tokenGen = new MockTokenGenerator();

  const result = await simulate({
    Client: (env) => Authenticate_Client(env, credentials, hasher),
    IP: (env) => Authenticate_IP(env, registry, tokenGen),
    Service: Authenticate_Service,
  });

  const expectedSalt = "FveKfWMX";
  const expectedToken = "201767";
  const expectedHash = "rOWr8MjbsyQwrX2XOIrHiZEE2eKOtL599Tr9nqZAL4w=";

  assertEquals(result, {
    Client: {
      receives: [
        { sender: "IP", value: expectedSalt },
        { sender: "IP", value: true },
        { sender: "IP", value: expectedToken },
      ],
      return: new AuthResult_C({
        Success: true,
        Token: expectedToken,
      }),
      sends: [
        { receivers: ["IP"], value: "username" },
        {
          receivers: ["IP"],
          value: expectedHash,
        },
      ],
    },
    IP: {
      receives: [
        { sender: "Client", value: "username" },
        {
          sender: "Client",
          value: expectedHash,
        },
      ],
      return: undefined,
      sends: [
        { receivers: ["Client"], value: expectedSalt },
        { receivers: ["Client", "Service"], value: true },
        { receivers: ["Client", "Service"], value: expectedToken },
      ],
    },
    Service: {
      receives: [
        { sender: "IP", value: true },
        { sender: "IP", value: expectedToken },
      ],
      return: new AuthResult_S({ Success: true, Token: expectedToken }),
      sends: [],
    },
  });

  assertEquals(hasher.hashes, [expectedHash], "hashes");
  assertEquals(registry.checkCalls, [expectedHash], "check calls");
  assertEquals(registry.getSalts, [expectedSalt], "get salts");
  assertEquals(tokenGen.tokens, [expectedToken], "tokens");
});
