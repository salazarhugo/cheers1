/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Party } from "../../type/party/party";
import { User } from "../../type/user/user";

export const protobufPackage = "cheers.party.v1";

export enum PartyUserResponse {
  GOING = 0,
  INTERESTED = 1,
  NOT_INTERESTED = 2,
  UNRECOGNIZED = -1,
}

export function partyUserResponseFromJSON(object: any): PartyUserResponse {
  switch (object) {
    case 0:
    case "GOING":
      return PartyUserResponse.GOING;
    case 1:
    case "INTERESTED":
      return PartyUserResponse.INTERESTED;
    case 2:
    case "NOT_INTERESTED":
      return PartyUserResponse.NOT_INTERESTED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PartyUserResponse.UNRECOGNIZED;
  }
}

export function partyUserResponseToJSON(object: PartyUserResponse): string {
  switch (object) {
    case PartyUserResponse.GOING:
      return "GOING";
    case PartyUserResponse.INTERESTED:
      return "INTERESTED";
    case PartyUserResponse.NOT_INTERESTED:
      return "NOT_INTERESTED";
    case PartyUserResponse.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface CreatePartyRequest {
  party: Party | undefined;
}

export interface CreatePartyResponse {
  party: Party | undefined;
}

export interface GetPartyRequest {
  partyId: string;
}

export interface GetPartyResponse {
  party: Party | undefined;
}

export interface GetPartyItemRequest {
  partyId: string;
}

export interface GetPartyItemResponse {
  item: PartyItem | undefined;
}

export interface UpdatePartyRequest {
  party: Party | undefined;
}

export interface UpdatePartyResponse {
  party: Party | undefined;
}

export interface DeletePartyRequest {
  partyId: string;
}

export interface DeletePartyResponse {
}

export interface FeedPartyRequest {
  parent: string;
  pageSize: number;
  pageToken: string;
}

export interface FeedPartyResponse {
  items: PartyItem[];
  nextPageToken: string;
}

export interface PartyItem {
  party: Party | undefined;
  creator: User | undefined;
  goingCount: number;
  interestedCount: number;
  invitedCount: number;
  isGoing: boolean;
  isInterested: boolean;
  isCreator: boolean;
  userResponse: PartyUserResponse;
}

function createBaseCreatePartyRequest(): CreatePartyRequest {
  return { party: undefined };
}

export const CreatePartyRequest = {
  encode(message: CreatePartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePartyRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.party = Party.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreatePartyRequest {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: CreatePartyRequest): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreatePartyRequest>, I>>(object: I): CreatePartyRequest {
    const message = createBaseCreatePartyRequest();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseCreatePartyResponse(): CreatePartyResponse {
  return { party: undefined };
}

export const CreatePartyResponse = {
  encode(message: CreatePartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePartyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.party = Party.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreatePartyResponse {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: CreatePartyResponse): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreatePartyResponse>, I>>(object: I): CreatePartyResponse {
    const message = createBaseCreatePartyResponse();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseGetPartyRequest(): GetPartyRequest {
  return { partyId: "" };
}

export const GetPartyRequest = {
  encode(message: GetPartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.partyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetPartyRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: GetPartyRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyRequest>, I>>(object: I): GetPartyRequest {
    const message = createBaseGetPartyRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseGetPartyResponse(): GetPartyResponse {
  return { party: undefined };
}

export const GetPartyResponse = {
  encode(message: GetPartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.party = Party.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetPartyResponse {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: GetPartyResponse): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyResponse>, I>>(object: I): GetPartyResponse {
    const message = createBaseGetPartyResponse();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseGetPartyItemRequest(): GetPartyItemRequest {
  return { partyId: "" };
}

export const GetPartyItemRequest = {
  encode(message: GetPartyItemRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyItemRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyItemRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.partyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetPartyItemRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: GetPartyItemRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyItemRequest>, I>>(object: I): GetPartyItemRequest {
    const message = createBaseGetPartyItemRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseGetPartyItemResponse(): GetPartyItemResponse {
  return { item: undefined };
}

export const GetPartyItemResponse = {
  encode(message: GetPartyItemResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.item !== undefined) {
      PartyItem.encode(message.item, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyItemResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyItemResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.item = PartyItem.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetPartyItemResponse {
    return { item: isSet(object.item) ? PartyItem.fromJSON(object.item) : undefined };
  },

  toJSON(message: GetPartyItemResponse): unknown {
    const obj: any = {};
    message.item !== undefined && (obj.item = message.item ? PartyItem.toJSON(message.item) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyItemResponse>, I>>(object: I): GetPartyItemResponse {
    const message = createBaseGetPartyItemResponse();
    message.item = (object.item !== undefined && object.item !== null) ? PartyItem.fromPartial(object.item) : undefined;
    return message;
  },
};

function createBaseUpdatePartyRequest(): UpdatePartyRequest {
  return { party: undefined };
}

export const UpdatePartyRequest = {
  encode(message: UpdatePartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdatePartyRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdatePartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.party = Party.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdatePartyRequest {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: UpdatePartyRequest): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdatePartyRequest>, I>>(object: I): UpdatePartyRequest {
    const message = createBaseUpdatePartyRequest();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseUpdatePartyResponse(): UpdatePartyResponse {
  return { party: undefined };
}

export const UpdatePartyResponse = {
  encode(message: UpdatePartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdatePartyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdatePartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.party = Party.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdatePartyResponse {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: UpdatePartyResponse): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdatePartyResponse>, I>>(object: I): UpdatePartyResponse {
    const message = createBaseUpdatePartyResponse();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseDeletePartyRequest(): DeletePartyRequest {
  return { partyId: "" };
}

export const DeletePartyRequest = {
  encode(message: DeletePartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePartyRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.partyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeletePartyRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: DeletePartyRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletePartyRequest>, I>>(object: I): DeletePartyRequest {
    const message = createBaseDeletePartyRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseDeletePartyResponse(): DeletePartyResponse {
  return {};
}

export const DeletePartyResponse = {
  encode(_: DeletePartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePartyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePartyResponse();
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

  fromJSON(_: any): DeletePartyResponse {
    return {};
  },

  toJSON(_: DeletePartyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletePartyResponse>, I>>(_: I): DeletePartyResponse {
    const message = createBaseDeletePartyResponse();
    return message;
  },
};

function createBaseFeedPartyRequest(): FeedPartyRequest {
  return { parent: "", pageSize: 0, pageToken: "" };
}

export const FeedPartyRequest = {
  encode(message: FeedPartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.parent !== "") {
      writer.uint32(10).string(message.parent);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.pageToken !== "") {
      writer.uint32(26).string(message.pageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedPartyRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.parent = reader.string();
          break;
        case 2:
          message.pageSize = reader.int32();
          break;
        case 3:
          message.pageToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeedPartyRequest {
    return {
      parent: isSet(object.parent) ? String(object.parent) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      pageToken: isSet(object.pageToken) ? String(object.pageToken) : "",
    };
  },

  toJSON(message: FeedPartyRequest): unknown {
    const obj: any = {};
    message.parent !== undefined && (obj.parent = message.parent);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.pageToken !== undefined && (obj.pageToken = message.pageToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeedPartyRequest>, I>>(object: I): FeedPartyRequest {
    const message = createBaseFeedPartyRequest();
    message.parent = object.parent ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.pageToken = object.pageToken ?? "";
    return message;
  },
};

function createBaseFeedPartyResponse(): FeedPartyResponse {
  return { items: [], nextPageToken: "" };
}

export const FeedPartyResponse = {
  encode(message: FeedPartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      PartyItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedPartyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(PartyItem.decode(reader, reader.uint32()));
          break;
        case 2:
          message.nextPageToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeedPartyResponse {
    return {
      items: Array.isArray(object?.items) ? object.items.map((e: any) => PartyItem.fromJSON(e)) : [],
      nextPageToken: isSet(object.nextPageToken) ? String(object.nextPageToken) : "",
    };
  },

  toJSON(message: FeedPartyResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? PartyItem.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    message.nextPageToken !== undefined && (obj.nextPageToken = message.nextPageToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeedPartyResponse>, I>>(object: I): FeedPartyResponse {
    const message = createBaseFeedPartyResponse();
    message.items = object.items?.map((e) => PartyItem.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBasePartyItem(): PartyItem {
  return {
    party: undefined,
    creator: undefined,
    goingCount: 0,
    interestedCount: 0,
    invitedCount: 0,
    isGoing: false,
    isInterested: false,
    isCreator: false,
    userResponse: 0,
  };
}

export const PartyItem = {
  encode(message: PartyItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    if (message.creator !== undefined) {
      User.encode(message.creator, writer.uint32(18).fork()).ldelim();
    }
    if (message.goingCount !== 0) {
      writer.uint32(24).int64(message.goingCount);
    }
    if (message.interestedCount !== 0) {
      writer.uint32(32).int64(message.interestedCount);
    }
    if (message.invitedCount !== 0) {
      writer.uint32(40).int64(message.invitedCount);
    }
    if (message.isGoing === true) {
      writer.uint32(48).bool(message.isGoing);
    }
    if (message.isInterested === true) {
      writer.uint32(56).bool(message.isInterested);
    }
    if (message.isCreator === true) {
      writer.uint32(64).bool(message.isCreator);
    }
    if (message.userResponse !== 0) {
      writer.uint32(72).int32(message.userResponse);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PartyItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePartyItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.party = Party.decode(reader, reader.uint32());
          break;
        case 2:
          message.creator = User.decode(reader, reader.uint32());
          break;
        case 3:
          message.goingCount = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.interestedCount = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.invitedCount = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.isGoing = reader.bool();
          break;
        case 7:
          message.isInterested = reader.bool();
          break;
        case 8:
          message.isCreator = reader.bool();
          break;
        case 9:
          message.userResponse = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PartyItem {
    return {
      party: isSet(object.party) ? Party.fromJSON(object.party) : undefined,
      creator: isSet(object.creator) ? User.fromJSON(object.creator) : undefined,
      goingCount: isSet(object.goingCount) ? Number(object.goingCount) : 0,
      interestedCount: isSet(object.interestedCount) ? Number(object.interestedCount) : 0,
      invitedCount: isSet(object.invitedCount) ? Number(object.invitedCount) : 0,
      isGoing: isSet(object.isGoing) ? Boolean(object.isGoing) : false,
      isInterested: isSet(object.isInterested) ? Boolean(object.isInterested) : false,
      isCreator: isSet(object.isCreator) ? Boolean(object.isCreator) : false,
      userResponse: isSet(object.userResponse) ? partyUserResponseFromJSON(object.userResponse) : 0,
    };
  },

  toJSON(message: PartyItem): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    message.creator !== undefined && (obj.creator = message.creator ? User.toJSON(message.creator) : undefined);
    message.goingCount !== undefined && (obj.goingCount = Math.round(message.goingCount));
    message.interestedCount !== undefined && (obj.interestedCount = Math.round(message.interestedCount));
    message.invitedCount !== undefined && (obj.invitedCount = Math.round(message.invitedCount));
    message.isGoing !== undefined && (obj.isGoing = message.isGoing);
    message.isInterested !== undefined && (obj.isInterested = message.isInterested);
    message.isCreator !== undefined && (obj.isCreator = message.isCreator);
    message.userResponse !== undefined && (obj.userResponse = partyUserResponseToJSON(message.userResponse));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PartyItem>, I>>(object: I): PartyItem {
    const message = createBasePartyItem();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    message.creator = (object.creator !== undefined && object.creator !== null)
      ? User.fromPartial(object.creator)
      : undefined;
    message.goingCount = object.goingCount ?? 0;
    message.interestedCount = object.interestedCount ?? 0;
    message.invitedCount = object.invitedCount ?? 0;
    message.isGoing = object.isGoing ?? false;
    message.isInterested = object.isInterested ?? false;
    message.isCreator = object.isCreator ?? false;
    message.userResponse = object.userResponse ?? 0;
    return message;
  },
};

export interface PartyService {
  CreateParty(request: CreatePartyRequest): Promise<CreatePartyResponse>;
  GetParty(request: GetPartyRequest): Promise<GetPartyResponse>;
  UpdateParty(request: UpdatePartyRequest): Promise<UpdatePartyResponse>;
  DeleteParty(request: DeletePartyRequest): Promise<DeletePartyResponse>;
  GetPartyItem(request: GetPartyItemRequest): Promise<GetPartyItemResponse>;
  FeedParty(request: FeedPartyRequest): Promise<FeedPartyResponse>;
}

export class PartyServiceClientImpl implements PartyService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.party.v1.PartyService";
    this.rpc = rpc;
    this.CreateParty = this.CreateParty.bind(this);
    this.GetParty = this.GetParty.bind(this);
    this.UpdateParty = this.UpdateParty.bind(this);
    this.DeleteParty = this.DeleteParty.bind(this);
    this.GetPartyItem = this.GetPartyItem.bind(this);
    this.FeedParty = this.FeedParty.bind(this);
  }
  CreateParty(request: CreatePartyRequest): Promise<CreatePartyResponse> {
    const data = CreatePartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateParty", data);
    return promise.then((data) => CreatePartyResponse.decode(new _m0.Reader(data)));
  }

  GetParty(request: GetPartyRequest): Promise<GetPartyResponse> {
    const data = GetPartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetParty", data);
    return promise.then((data) => GetPartyResponse.decode(new _m0.Reader(data)));
  }

  UpdateParty(request: UpdatePartyRequest): Promise<UpdatePartyResponse> {
    const data = UpdatePartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateParty", data);
    return promise.then((data) => UpdatePartyResponse.decode(new _m0.Reader(data)));
  }

  DeleteParty(request: DeletePartyRequest): Promise<DeletePartyResponse> {
    const data = DeletePartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteParty", data);
    return promise.then((data) => DeletePartyResponse.decode(new _m0.Reader(data)));
  }

  GetPartyItem(request: GetPartyItemRequest): Promise<GetPartyItemResponse> {
    const data = GetPartyItemRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetPartyItem", data);
    return promise.then((data) => GetPartyItemResponse.decode(new _m0.Reader(data)));
  }

  FeedParty(request: FeedPartyRequest): Promise<FeedPartyResponse> {
    const data = FeedPartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FeedParty", data);
    return promise.then((data) => FeedPartyResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
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
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
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
