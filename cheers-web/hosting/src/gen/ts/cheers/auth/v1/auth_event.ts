/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.auth.v1";

export interface AuthEvent {
  created?: CreatedAuth | undefined;
  deleted?: DeletedAuth | undefined;
}

export interface CreatedAuth {
  uid: string;
  email: string;
}

export interface DeletedAuth {
  uid: string;
  email: string;
}

function createBaseAuthEvent(): AuthEvent {
  return { created: undefined, deleted: undefined };
}

export const AuthEvent = {
  encode(message: AuthEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.created !== undefined) {
      CreatedAuth.encode(message.created, writer.uint32(10).fork()).ldelim();
    }
    if (message.deleted !== undefined) {
      DeletedAuth.encode(message.deleted, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AuthEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuthEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.created = CreatedAuth.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.deleted = DeletedAuth.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AuthEvent {
    return {
      created: isSet(object.created) ? CreatedAuth.fromJSON(object.created) : undefined,
      deleted: isSet(object.deleted) ? DeletedAuth.fromJSON(object.deleted) : undefined,
    };
  },

  toJSON(message: AuthEvent): unknown {
    const obj: any = {};
    message.created !== undefined && (obj.created = message.created ? CreatedAuth.toJSON(message.created) : undefined);
    message.deleted !== undefined && (obj.deleted = message.deleted ? DeletedAuth.toJSON(message.deleted) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<AuthEvent>, I>>(base?: I): AuthEvent {
    return AuthEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AuthEvent>, I>>(object: I): AuthEvent {
    const message = createBaseAuthEvent();
    message.created = (object.created !== undefined && object.created !== null)
      ? CreatedAuth.fromPartial(object.created)
      : undefined;
    message.deleted = (object.deleted !== undefined && object.deleted !== null)
      ? DeletedAuth.fromPartial(object.deleted)
      : undefined;
    return message;
  },
};

function createBaseCreatedAuth(): CreatedAuth {
  return { uid: "", email: "" };
}

export const CreatedAuth = {
  encode(message: CreatedAuth, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatedAuth {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatedAuth();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.uid = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.email = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatedAuth {
    return { uid: isSet(object.uid) ? String(object.uid) : "", email: isSet(object.email) ? String(object.email) : "" };
  },

  toJSON(message: CreatedAuth): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.email !== undefined && (obj.email = message.email);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatedAuth>, I>>(base?: I): CreatedAuth {
    return CreatedAuth.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatedAuth>, I>>(object: I): CreatedAuth {
    const message = createBaseCreatedAuth();
    message.uid = object.uid ?? "";
    message.email = object.email ?? "";
    return message;
  },
};

function createBaseDeletedAuth(): DeletedAuth {
  return { uid: "", email: "" };
}

export const DeletedAuth = {
  encode(message: DeletedAuth, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletedAuth {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletedAuth();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.uid = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.email = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeletedAuth {
    return { uid: isSet(object.uid) ? String(object.uid) : "", email: isSet(object.email) ? String(object.email) : "" };
  },

  toJSON(message: DeletedAuth): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.email !== undefined && (obj.email = message.email);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletedAuth>, I>>(base?: I): DeletedAuth {
    return DeletedAuth.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletedAuth>, I>>(object: I): DeletedAuth {
    const message = createBaseDeletedAuth();
    message.uid = object.uid ?? "";
    message.email = object.email ?? "";
    return message;
  },
};

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
