/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.type";

export enum Privacy {
  FRIENDS = 0,
  PRIVATE = 1,
  PUBLIC = 2,
  GROUP = 3,
  UNRECOGNIZED = -1,
}

export function privacyFromJSON(object: any): Privacy {
  switch (object) {
    case 0:
    case "FRIENDS":
      return Privacy.FRIENDS;
    case 1:
    case "PRIVATE":
      return Privacy.PRIVATE;
    case 2:
    case "PUBLIC":
      return Privacy.PUBLIC;
    case 3:
    case "GROUP":
      return Privacy.GROUP;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Privacy.UNRECOGNIZED;
  }
}

export function privacyToJSON(object: Privacy): string {
  switch (object) {
    case Privacy.FRIENDS:
      return "FRIENDS";
    case Privacy.PRIVATE:
      return "PRIVATE";
    case Privacy.PUBLIC:
      return "PUBLIC";
    case Privacy.GROUP:
      return "GROUP";
    case Privacy.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
