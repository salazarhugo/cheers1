/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.chat.v1";

export interface ChatEvent {
  senderId: string;
  roomId: string;
}

function createBaseChatEvent(): ChatEvent {
  return { senderId: "", roomId: "" };
}

export const ChatEvent = {
  encode(message: ChatEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.senderId !== "") {
      writer.uint32(10).string(message.senderId);
    }
    if (message.roomId !== "") {
      writer.uint32(18).string(message.roomId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChatEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChatEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.senderId = reader.string();
          break;
        case 2:
          message.roomId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChatEvent {
    return {
      senderId: isSet(object.senderId) ? String(object.senderId) : "",
      roomId: isSet(object.roomId) ? String(object.roomId) : "",
    };
  },

  toJSON(message: ChatEvent): unknown {
    const obj: any = {};
    message.senderId !== undefined && (obj.senderId = message.senderId);
    message.roomId !== undefined && (obj.roomId = message.roomId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChatEvent>, I>>(object: I): ChatEvent {
    const message = createBaseChatEvent();
    message.senderId = object.senderId ?? "";
    message.roomId = object.roomId ?? "";
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
