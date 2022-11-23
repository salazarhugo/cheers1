/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.account.v1";

export interface Account {
  id: string;
  balance: number;
  ticketSold: number;
}

export interface CreateAccountRequest {
  account: Account | undefined;
}

export interface CreateAccountResponse {
  account: Account | undefined;
}

export interface GetAccountRequest {
  accountId: string;
}

export interface GetAccountResponse {
  account: Account | undefined;
}

export interface UpdateAccountRequest {
  account: Account | undefined;
}

export interface UpdateAccountResponse {
  account: Account | undefined;
}

export interface DeleteAccountRequest {
  accountId: string;
}

export interface DeleteAccountResponse {
}

export interface ListAccountRequest {
  partyId?: string | undefined;
  accountId?: string | undefined;
}

export interface ListAccountResponse {
  accounts: Account[];
}

function createBaseAccount(): Account {
  return { id: "", balance: 0, ticketSold: 0 };
}

export const Account = {
  encode(message: Account, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.balance !== 0) {
      writer.uint32(16).int32(message.balance);
    }
    if (message.ticketSold !== 0) {
      writer.uint32(24).int32(message.ticketSold);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Account {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.balance = reader.int32();
          break;
        case 3:
          message.ticketSold = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Account {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      balance: isSet(object.balance) ? Number(object.balance) : 0,
      ticketSold: isSet(object.ticketSold) ? Number(object.ticketSold) : 0,
    };
  },

  toJSON(message: Account): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.balance !== undefined && (obj.balance = Math.round(message.balance));
    message.ticketSold !== undefined && (obj.ticketSold = Math.round(message.ticketSold));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Account>, I>>(object: I): Account {
    const message = createBaseAccount();
    message.id = object.id ?? "";
    message.balance = object.balance ?? 0;
    message.ticketSold = object.ticketSold ?? 0;
    return message;
  },
};

function createBaseCreateAccountRequest(): CreateAccountRequest {
  return { account: undefined };
}

export const CreateAccountRequest = {
  encode(message: CreateAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateAccountRequest {
    return { account: isSet(object.account) ? Account.fromJSON(object.account) : undefined };
  },

  toJSON(message: CreateAccountRequest): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateAccountRequest>, I>>(object: I): CreateAccountRequest {
    const message = createBaseCreateAccountRequest();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    return message;
  },
};

function createBaseCreateAccountResponse(): CreateAccountResponse {
  return { account: undefined };
}

export const CreateAccountResponse = {
  encode(message: CreateAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateAccountResponse {
    return { account: isSet(object.account) ? Account.fromJSON(object.account) : undefined };
  },

  toJSON(message: CreateAccountResponse): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateAccountResponse>, I>>(object: I): CreateAccountResponse {
    const message = createBaseCreateAccountResponse();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    return message;
  },
};

function createBaseGetAccountRequest(): GetAccountRequest {
  return { accountId: "" };
}

export const GetAccountRequest = {
  encode(message: GetAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetAccountRequest {
    return { accountId: isSet(object.accountId) ? String(object.accountId) : "" };
  },

  toJSON(message: GetAccountRequest): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetAccountRequest>, I>>(object: I): GetAccountRequest {
    const message = createBaseGetAccountRequest();
    message.accountId = object.accountId ?? "";
    return message;
  },
};

function createBaseGetAccountResponse(): GetAccountResponse {
  return { account: undefined };
}

export const GetAccountResponse = {
  encode(message: GetAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetAccountResponse {
    return { account: isSet(object.account) ? Account.fromJSON(object.account) : undefined };
  },

  toJSON(message: GetAccountResponse): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetAccountResponse>, I>>(object: I): GetAccountResponse {
    const message = createBaseGetAccountResponse();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    return message;
  },
};

function createBaseUpdateAccountRequest(): UpdateAccountRequest {
  return { account: undefined };
}

export const UpdateAccountRequest = {
  encode(message: UpdateAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateAccountRequest {
    return { account: isSet(object.account) ? Account.fromJSON(object.account) : undefined };
  },

  toJSON(message: UpdateAccountRequest): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateAccountRequest>, I>>(object: I): UpdateAccountRequest {
    const message = createBaseUpdateAccountRequest();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    return message;
  },
};

function createBaseUpdateAccountResponse(): UpdateAccountResponse {
  return { account: undefined };
}

export const UpdateAccountResponse = {
  encode(message: UpdateAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateAccountResponse {
    return { account: isSet(object.account) ? Account.fromJSON(object.account) : undefined };
  },

  toJSON(message: UpdateAccountResponse): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateAccountResponse>, I>>(object: I): UpdateAccountResponse {
    const message = createBaseUpdateAccountResponse();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    return message;
  },
};

function createBaseDeleteAccountRequest(): DeleteAccountRequest {
  return { accountId: "" };
}

export const DeleteAccountRequest = {
  encode(message: DeleteAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteAccountRequest {
    return { accountId: isSet(object.accountId) ? String(object.accountId) : "" };
  },

  toJSON(message: DeleteAccountRequest): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteAccountRequest>, I>>(object: I): DeleteAccountRequest {
    const message = createBaseDeleteAccountRequest();
    message.accountId = object.accountId ?? "";
    return message;
  },
};

function createBaseDeleteAccountResponse(): DeleteAccountResponse {
  return {};
}

export const DeleteAccountResponse = {
  encode(_: DeleteAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): DeleteAccountResponse {
    return {};
  },

  toJSON(_: DeleteAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteAccountResponse>, I>>(_: I): DeleteAccountResponse {
    const message = createBaseDeleteAccountResponse();
    return message;
  },
};

function createBaseListAccountRequest(): ListAccountRequest {
  return { partyId: undefined, accountId: undefined };
}

export const ListAccountRequest = {
  encode(message: ListAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== undefined) {
      writer.uint32(10).string(message.partyId);
    }
    if (message.accountId !== undefined) {
      writer.uint32(18).string(message.accountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.partyId = reader.string();
          break;
        case 2:
          message.accountId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListAccountRequest {
    return {
      partyId: isSet(object.partyId) ? String(object.partyId) : undefined,
      accountId: isSet(object.accountId) ? String(object.accountId) : undefined,
    };
  },

  toJSON(message: ListAccountRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.accountId !== undefined && (obj.accountId = message.accountId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListAccountRequest>, I>>(object: I): ListAccountRequest {
    const message = createBaseListAccountRequest();
    message.partyId = object.partyId ?? undefined;
    message.accountId = object.accountId ?? undefined;
    return message;
  },
};

function createBaseListAccountResponse(): ListAccountResponse {
  return { accounts: [] };
}

export const ListAccountResponse = {
  encode(message: ListAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.accounts) {
      Account.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accounts.push(Account.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListAccountResponse {
    return { accounts: Array.isArray(object?.accounts) ? object.accounts.map((e: any) => Account.fromJSON(e)) : [] };
  },

  toJSON(message: ListAccountResponse): unknown {
    const obj: any = {};
    if (message.accounts) {
      obj.accounts = message.accounts.map((e) => e ? Account.toJSON(e) : undefined);
    } else {
      obj.accounts = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListAccountResponse>, I>>(object: I): ListAccountResponse {
    const message = createBaseListAccountResponse();
    message.accounts = object.accounts?.map((e) => Account.fromPartial(e)) || [];
    return message;
  },
};

export interface AccountService {
  CreateAccount(request: CreateAccountRequest): Promise<CreateAccountResponse>;
  GetAccount(request: GetAccountRequest): Promise<GetAccountResponse>;
  UpdateAccount(request: UpdateAccountRequest): Promise<UpdateAccountResponse>;
  DeleteAccount(request: DeleteAccountRequest): Promise<DeleteAccountResponse>;
  ListAccount(request: ListAccountRequest): Promise<ListAccountResponse>;
}

export class AccountServiceClientImpl implements AccountService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.account.v1.AccountService";
    this.rpc = rpc;
    this.CreateAccount = this.CreateAccount.bind(this);
    this.GetAccount = this.GetAccount.bind(this);
    this.UpdateAccount = this.UpdateAccount.bind(this);
    this.DeleteAccount = this.DeleteAccount.bind(this);
    this.ListAccount = this.ListAccount.bind(this);
  }
  CreateAccount(request: CreateAccountRequest): Promise<CreateAccountResponse> {
    const data = CreateAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateAccount", data);
    return promise.then((data) => CreateAccountResponse.decode(new _m0.Reader(data)));
  }

  GetAccount(request: GetAccountRequest): Promise<GetAccountResponse> {
    const data = GetAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetAccount", data);
    return promise.then((data) => GetAccountResponse.decode(new _m0.Reader(data)));
  }

  UpdateAccount(request: UpdateAccountRequest): Promise<UpdateAccountResponse> {
    const data = UpdateAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateAccount", data);
    return promise.then((data) => UpdateAccountResponse.decode(new _m0.Reader(data)));
  }

  DeleteAccount(request: DeleteAccountRequest): Promise<DeleteAccountResponse> {
    const data = DeleteAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteAccount", data);
    return promise.then((data) => DeleteAccountResponse.decode(new _m0.Reader(data)));
  }

  ListAccount(request: ListAccountRequest): Promise<ListAccountResponse> {
    const data = ListAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListAccount", data);
    return promise.then((data) => ListAccountResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
