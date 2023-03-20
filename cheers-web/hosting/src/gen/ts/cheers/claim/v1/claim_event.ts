/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.claim.v1";

export interface ClaimEvent {
  created?: CreatedClaim | undefined;
  deleted?: DeletedClaim | undefined;
}

export interface CreatedClaim {
  userId: string;
  claim: string;
}

export interface DeletedClaim {
  userId: string;
  claim: string;
}

function createBaseClaimEvent(): ClaimEvent {
  return { created: undefined, deleted: undefined };
}

export const ClaimEvent = {
  encode(message: ClaimEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.created !== undefined) {
      CreatedClaim.encode(message.created, writer.uint32(10).fork()).ldelim();
    }
    if (message.deleted !== undefined) {
      DeletedClaim.encode(message.deleted, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ClaimEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClaimEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.created = CreatedClaim.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.deleted = DeletedClaim.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ClaimEvent {
    return {
      created: isSet(object.created) ? CreatedClaim.fromJSON(object.created) : undefined,
      deleted: isSet(object.deleted) ? DeletedClaim.fromJSON(object.deleted) : undefined,
    };
  },

  toJSON(message: ClaimEvent): unknown {
    const obj: any = {};
    message.created !== undefined && (obj.created = message.created ? CreatedClaim.toJSON(message.created) : undefined);
    message.deleted !== undefined && (obj.deleted = message.deleted ? DeletedClaim.toJSON(message.deleted) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ClaimEvent>, I>>(base?: I): ClaimEvent {
    return ClaimEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ClaimEvent>, I>>(object: I): ClaimEvent {
    const message = createBaseClaimEvent();
    message.created = (object.created !== undefined && object.created !== null)
      ? CreatedClaim.fromPartial(object.created)
      : undefined;
    message.deleted = (object.deleted !== undefined && object.deleted !== null)
      ? DeletedClaim.fromPartial(object.deleted)
      : undefined;
    return message;
  },
};

function createBaseCreatedClaim(): CreatedClaim {
  return { userId: "", claim: "" };
}

export const CreatedClaim = {
  encode(message: CreatedClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.claim !== "") {
      writer.uint32(18).string(message.claim);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatedClaim {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatedClaim();
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

          message.claim = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatedClaim {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      claim: isSet(object.claim) ? String(object.claim) : "",
    };
  },

  toJSON(message: CreatedClaim): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.claim !== undefined && (obj.claim = message.claim);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatedClaim>, I>>(base?: I): CreatedClaim {
    return CreatedClaim.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatedClaim>, I>>(object: I): CreatedClaim {
    const message = createBaseCreatedClaim();
    message.userId = object.userId ?? "";
    message.claim = object.claim ?? "";
    return message;
  },
};

function createBaseDeletedClaim(): DeletedClaim {
  return { userId: "", claim: "" };
}

export const DeletedClaim = {
  encode(message: DeletedClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.claim !== "") {
      writer.uint32(18).string(message.claim);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletedClaim {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletedClaim();
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

          message.claim = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeletedClaim {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      claim: isSet(object.claim) ? String(object.claim) : "",
    };
  },

  toJSON(message: DeletedClaim): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.claim !== undefined && (obj.claim = message.claim);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletedClaim>, I>>(base?: I): DeletedClaim {
    return DeletedClaim.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletedClaim>, I>>(object: I): DeletedClaim {
    const message = createBaseDeletedClaim();
    message.userId = object.userId ?? "";
    message.claim = object.claim ?? "";
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
