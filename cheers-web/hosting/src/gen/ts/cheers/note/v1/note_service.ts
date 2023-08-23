/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.note.v1";

export interface ListFriendNoteRequest {
}

export interface ListFriendNoteResponse {
  items: Note[];
}

export interface Note {
  userId: string;
  text: string;
  name: string;
  username: string;
  picture: string;
  created: number;
}

export interface CreateNoteRequest {
  text: string;
}

export interface CreateNoteResponse {
  note: Note | undefined;
}

export interface DeleteNoteRequest {
  text: string;
}

export interface DeleteNoteResponse {
}

function createBaseListFriendNoteRequest(): ListFriendNoteRequest {
  return {};
}

export const ListFriendNoteRequest = {
  encode(_: ListFriendNoteRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendNoteRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendNoteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): ListFriendNoteRequest {
    return {};
  },

  toJSON(_: ListFriendNoteRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendNoteRequest>, I>>(base?: I): ListFriendNoteRequest {
    return ListFriendNoteRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendNoteRequest>, I>>(_: I): ListFriendNoteRequest {
    const message = createBaseListFriendNoteRequest();
    return message;
  },
};

function createBaseListFriendNoteResponse(): ListFriendNoteResponse {
  return { items: [] };
}

export const ListFriendNoteResponse = {
  encode(message: ListFriendNoteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      Note.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendNoteResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendNoteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.items.push(Note.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListFriendNoteResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => Note.fromJSON(e)) : [] };
  },

  toJSON(message: ListFriendNoteResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? Note.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendNoteResponse>, I>>(base?: I): ListFriendNoteResponse {
    return ListFriendNoteResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendNoteResponse>, I>>(object: I): ListFriendNoteResponse {
    const message = createBaseListFriendNoteResponse();
    message.items = object.items?.map((e) => Note.fromPartial(e)) || [];
    return message;
  },
};

function createBaseNote(): Note {
  return { userId: "", text: "", name: "", username: "", picture: "", created: 0 };
}

export const Note = {
  encode(message: Note, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.username !== "") {
      writer.uint32(34).string(message.username);
    }
    if (message.picture !== "") {
      writer.uint32(42).string(message.picture);
    }
    if (message.created !== 0) {
      writer.uint32(48).int64(message.created);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Note {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNote();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.text = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.username = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.picture = reader.string();
          continue;
        case 6:
          if (tag != 48) {
            break;
          }

          message.created = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Note {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      text: isSet(object.text) ? String(object.text) : "",
      name: isSet(object.name) ? String(object.name) : "",
      username: isSet(object.username) ? String(object.username) : "",
      picture: isSet(object.picture) ? String(object.picture) : "",
      created: isSet(object.created) ? Number(object.created) : 0,
    };
  },

  toJSON(message: Note): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.text !== undefined && (obj.text = message.text);
    message.name !== undefined && (obj.name = message.name);
    message.username !== undefined && (obj.username = message.username);
    message.picture !== undefined && (obj.picture = message.picture);
    message.created !== undefined && (obj.created = Math.round(message.created));
    return obj;
  },

  create<I extends Exact<DeepPartial<Note>, I>>(base?: I): Note {
    return Note.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Note>, I>>(object: I): Note {
    const message = createBaseNote();
    message.userId = object.userId ?? "";
    message.text = object.text ?? "";
    message.name = object.name ?? "";
    message.username = object.username ?? "";
    message.picture = object.picture ?? "";
    message.created = object.created ?? 0;
    return message;
  },
};

function createBaseCreateNoteRequest(): CreateNoteRequest {
  return { text: "" };
}

export const CreateNoteRequest = {
  encode(message: CreateNoteRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateNoteRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateNoteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.text = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateNoteRequest {
    return { text: isSet(object.text) ? String(object.text) : "" };
  },

  toJSON(message: CreateNoteRequest): unknown {
    const obj: any = {};
    message.text !== undefined && (obj.text = message.text);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateNoteRequest>, I>>(base?: I): CreateNoteRequest {
    return CreateNoteRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateNoteRequest>, I>>(object: I): CreateNoteRequest {
    const message = createBaseCreateNoteRequest();
    message.text = object.text ?? "";
    return message;
  },
};

function createBaseCreateNoteResponse(): CreateNoteResponse {
  return { note: undefined };
}

export const CreateNoteResponse = {
  encode(message: CreateNoteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.note !== undefined) {
      Note.encode(message.note, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateNoteResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateNoteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.note = Note.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateNoteResponse {
    return { note: isSet(object.note) ? Note.fromJSON(object.note) : undefined };
  },

  toJSON(message: CreateNoteResponse): unknown {
    const obj: any = {};
    message.note !== undefined && (obj.note = message.note ? Note.toJSON(message.note) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateNoteResponse>, I>>(base?: I): CreateNoteResponse {
    return CreateNoteResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateNoteResponse>, I>>(object: I): CreateNoteResponse {
    const message = createBaseCreateNoteResponse();
    message.note = (object.note !== undefined && object.note !== null) ? Note.fromPartial(object.note) : undefined;
    return message;
  },
};

function createBaseDeleteNoteRequest(): DeleteNoteRequest {
  return { text: "" };
}

export const DeleteNoteRequest = {
  encode(message: DeleteNoteRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteNoteRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteNoteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.text = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteNoteRequest {
    return { text: isSet(object.text) ? String(object.text) : "" };
  },

  toJSON(message: DeleteNoteRequest): unknown {
    const obj: any = {};
    message.text !== undefined && (obj.text = message.text);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteNoteRequest>, I>>(base?: I): DeleteNoteRequest {
    return DeleteNoteRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteNoteRequest>, I>>(object: I): DeleteNoteRequest {
    const message = createBaseDeleteNoteRequest();
    message.text = object.text ?? "";
    return message;
  },
};

function createBaseDeleteNoteResponse(): DeleteNoteResponse {
  return {};
}

export const DeleteNoteResponse = {
  encode(_: DeleteNoteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteNoteResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteNoteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): DeleteNoteResponse {
    return {};
  },

  toJSON(_: DeleteNoteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteNoteResponse>, I>>(base?: I): DeleteNoteResponse {
    return DeleteNoteResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteNoteResponse>, I>>(_: I): DeleteNoteResponse {
    const message = createBaseDeleteNoteResponse();
    return message;
  },
};

export interface NoteService {
  CreateNote(request: CreateNoteRequest): Promise<CreateNoteResponse>;
  DeleteNote(request: DeleteNoteRequest): Promise<DeleteNoteResponse>;
  ListFriendNote(request: ListFriendNoteRequest): Promise<ListFriendNoteResponse>;
}

export class NoteServiceClientImpl implements NoteService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.note.v1.NoteService";
    this.rpc = rpc;
    this.CreateNote = this.CreateNote.bind(this);
    this.DeleteNote = this.DeleteNote.bind(this);
    this.ListFriendNote = this.ListFriendNote.bind(this);
  }
  CreateNote(request: CreateNoteRequest): Promise<CreateNoteResponse> {
    const data = CreateNoteRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateNote", data);
    return promise.then((data) => CreateNoteResponse.decode(_m0.Reader.create(data)));
  }

  DeleteNote(request: DeleteNoteRequest): Promise<DeleteNoteResponse> {
    const data = DeleteNoteRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteNote", data);
    return promise.then((data) => DeleteNoteResponse.decode(_m0.Reader.create(data)));
  }

  ListFriendNote(request: ListFriendNoteRequest): Promise<ListFriendNoteResponse> {
    const data = ListFriendNoteRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFriendNote", data);
    return promise.then((data) => ListFriendNoteResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
