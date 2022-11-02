/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../../google/protobuf/empty";
import { Story } from "../../type/story/story";
import { User } from "../../type/user/user";

export const protobufPackage = "cheers.story.v1";

export interface ListUserStoryRequest {
  userId: string;
}

export interface ListUserStoryResponse {
  user: User | undefined;
  stories: StoryResponse[];
}

export interface CreateStoryRequest {
  story: Story | undefined;
}

export interface GetStoryRequest {
  id: string;
}

export interface UpdateStoryRequest {
  story: Story | undefined;
}

export interface DeleteStoryRequest {
  id: string;
}

export interface FeedStoryRequest {
  parent: string;
  pageSize: number;
  page: number;
}

export interface FeedStoryResponse {
  items: UserWithStories[];
  nextPageToken: string;
}

export interface ViewStoryRequest {
  id: string;
}

export interface ViewStoryResponse {
  success: boolean;
}

export interface LikeStoryRequest {
  id: string;
}

export interface LikeStoryResponse {
  success: boolean;
}

export interface UnlikeStoryRequest {
  id: string;
}

export interface UnlikeStoryResponse {
  success: boolean;
}

export interface SaveStoryRequest {
  id: string;
}

export interface SaveStoryResponse {
  success: boolean;
}

export interface UnsaveStoryRequest {
  id: string;
}

export interface UnsaveStoryResponse {
  success: boolean;
}

export interface ListStoryResponse {
  user: User | undefined;
  items: UserWithStories[];
}

export interface UserWithStories {
  user: User | undefined;
  stories: StoryResponse[];
}

export interface StoryResponse {
  story: Story | undefined;
  hasLiked: boolean;
  hasViewed: boolean;
}

function createBaseListUserStoryRequest(): ListUserStoryRequest {
  return { userId: "" };
}

export const ListUserStoryRequest = {
  encode(message: ListUserStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUserStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListUserStoryRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: ListUserStoryRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListUserStoryRequest>, I>>(object: I): ListUserStoryRequest {
    const message = createBaseListUserStoryRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseListUserStoryResponse(): ListUserStoryResponse {
  return { user: undefined, stories: [] };
}

export const ListUserStoryResponse = {
  encode(message: ListUserStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.stories) {
      StoryResponse.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUserStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        case 2:
          message.stories.push(StoryResponse.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListUserStoryResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      stories: Array.isArray(object?.stories) ? object.stories.map((e: any) => StoryResponse.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListUserStoryResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    if (message.stories) {
      obj.stories = message.stories.map((e) => e ? StoryResponse.toJSON(e) : undefined);
    } else {
      obj.stories = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListUserStoryResponse>, I>>(object: I): ListUserStoryResponse {
    const message = createBaseListUserStoryResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.stories = object.stories?.map((e) => StoryResponse.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateStoryRequest(): CreateStoryRequest {
  return { story: undefined };
}

export const CreateStoryRequest = {
  encode(message: CreateStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.story !== undefined) {
      Story.encode(message.story, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.story = Story.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateStoryRequest {
    return { story: isSet(object.story) ? Story.fromJSON(object.story) : undefined };
  },

  toJSON(message: CreateStoryRequest): unknown {
    const obj: any = {};
    message.story !== undefined && (obj.story = message.story ? Story.toJSON(message.story) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateStoryRequest>, I>>(object: I): CreateStoryRequest {
    const message = createBaseCreateStoryRequest();
    message.story = (object.story !== undefined && object.story !== null) ? Story.fromPartial(object.story) : undefined;
    return message;
  },
};

function createBaseGetStoryRequest(): GetStoryRequest {
  return { id: "" };
}

export const GetStoryRequest = {
  encode(message: GetStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetStoryRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: GetStoryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetStoryRequest>, I>>(object: I): GetStoryRequest {
    const message = createBaseGetStoryRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseUpdateStoryRequest(): UpdateStoryRequest {
  return { story: undefined };
}

export const UpdateStoryRequest = {
  encode(message: UpdateStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.story !== undefined) {
      Story.encode(message.story, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.story = Story.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateStoryRequest {
    return { story: isSet(object.story) ? Story.fromJSON(object.story) : undefined };
  },

  toJSON(message: UpdateStoryRequest): unknown {
    const obj: any = {};
    message.story !== undefined && (obj.story = message.story ? Story.toJSON(message.story) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateStoryRequest>, I>>(object: I): UpdateStoryRequest {
    const message = createBaseUpdateStoryRequest();
    message.story = (object.story !== undefined && object.story !== null) ? Story.fromPartial(object.story) : undefined;
    return message;
  },
};

function createBaseDeleteStoryRequest(): DeleteStoryRequest {
  return { id: "" };
}

export const DeleteStoryRequest = {
  encode(message: DeleteStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteStoryRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: DeleteStoryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteStoryRequest>, I>>(object: I): DeleteStoryRequest {
    const message = createBaseDeleteStoryRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseFeedStoryRequest(): FeedStoryRequest {
  return { parent: "", pageSize: 0, page: 0 };
}

export const FeedStoryRequest = {
  encode(message: FeedStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.parent !== "") {
      writer.uint32(10).string(message.parent);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedStoryRequest();
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
          message.page = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeedStoryRequest {
    return {
      parent: isSet(object.parent) ? String(object.parent) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: FeedStoryRequest): unknown {
    const obj: any = {};
    message.parent !== undefined && (obj.parent = message.parent);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeedStoryRequest>, I>>(object: I): FeedStoryRequest {
    const message = createBaseFeedStoryRequest();
    message.parent = object.parent ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseFeedStoryResponse(): FeedStoryResponse {
  return { items: [], nextPageToken: "" };
}

export const FeedStoryResponse = {
  encode(message: FeedStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      UserWithStories.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(UserWithStories.decode(reader, reader.uint32()));
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

  fromJSON(object: any): FeedStoryResponse {
    return {
      items: Array.isArray(object?.items) ? object.items.map((e: any) => UserWithStories.fromJSON(e)) : [],
      nextPageToken: isSet(object.nextPageToken) ? String(object.nextPageToken) : "",
    };
  },

  toJSON(message: FeedStoryResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? UserWithStories.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    message.nextPageToken !== undefined && (obj.nextPageToken = message.nextPageToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeedStoryResponse>, I>>(object: I): FeedStoryResponse {
    const message = createBaseFeedStoryResponse();
    message.items = object.items?.map((e) => UserWithStories.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBaseViewStoryRequest(): ViewStoryRequest {
  return { id: "" };
}

export const ViewStoryRequest = {
  encode(message: ViewStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ViewStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseViewStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ViewStoryRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: ViewStoryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ViewStoryRequest>, I>>(object: I): ViewStoryRequest {
    const message = createBaseViewStoryRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseViewStoryResponse(): ViewStoryResponse {
  return { success: false };
}

export const ViewStoryResponse = {
  encode(message: ViewStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ViewStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseViewStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ViewStoryResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: ViewStoryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ViewStoryResponse>, I>>(object: I): ViewStoryResponse {
    const message = createBaseViewStoryResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseLikeStoryRequest(): LikeStoryRequest {
  return { id: "" };
}

export const LikeStoryRequest = {
  encode(message: LikeStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikeStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikeStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LikeStoryRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: LikeStoryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LikeStoryRequest>, I>>(object: I): LikeStoryRequest {
    const message = createBaseLikeStoryRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseLikeStoryResponse(): LikeStoryResponse {
  return { success: false };
}

export const LikeStoryResponse = {
  encode(message: LikeStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikeStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikeStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LikeStoryResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: LikeStoryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LikeStoryResponse>, I>>(object: I): LikeStoryResponse {
    const message = createBaseLikeStoryResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseUnlikeStoryRequest(): UnlikeStoryRequest {
  return { id: "" };
}

export const UnlikeStoryRequest = {
  encode(message: UnlikeStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnlikeStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnlikeStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnlikeStoryRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: UnlikeStoryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnlikeStoryRequest>, I>>(object: I): UnlikeStoryRequest {
    const message = createBaseUnlikeStoryRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseUnlikeStoryResponse(): UnlikeStoryResponse {
  return { success: false };
}

export const UnlikeStoryResponse = {
  encode(message: UnlikeStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnlikeStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnlikeStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnlikeStoryResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: UnlikeStoryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnlikeStoryResponse>, I>>(object: I): UnlikeStoryResponse {
    const message = createBaseUnlikeStoryResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseSaveStoryRequest(): SaveStoryRequest {
  return { id: "" };
}

export const SaveStoryRequest = {
  encode(message: SaveStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SaveStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSaveStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SaveStoryRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: SaveStoryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SaveStoryRequest>, I>>(object: I): SaveStoryRequest {
    const message = createBaseSaveStoryRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseSaveStoryResponse(): SaveStoryResponse {
  return { success: false };
}

export const SaveStoryResponse = {
  encode(message: SaveStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SaveStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSaveStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SaveStoryResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: SaveStoryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SaveStoryResponse>, I>>(object: I): SaveStoryResponse {
    const message = createBaseSaveStoryResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseUnsaveStoryRequest(): UnsaveStoryRequest {
  return { id: "" };
}

export const UnsaveStoryRequest = {
  encode(message: UnsaveStoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnsaveStoryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnsaveStoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnsaveStoryRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: UnsaveStoryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnsaveStoryRequest>, I>>(object: I): UnsaveStoryRequest {
    const message = createBaseUnsaveStoryRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseUnsaveStoryResponse(): UnsaveStoryResponse {
  return { success: false };
}

export const UnsaveStoryResponse = {
  encode(message: UnsaveStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnsaveStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnsaveStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnsaveStoryResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: UnsaveStoryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnsaveStoryResponse>, I>>(object: I): UnsaveStoryResponse {
    const message = createBaseUnsaveStoryResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseListStoryResponse(): ListStoryResponse {
  return { user: undefined, items: [] };
}

export const ListStoryResponse = {
  encode(message: ListStoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.items) {
      UserWithStories.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListStoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        case 2:
          message.items.push(UserWithStories.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListStoryResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      items: Array.isArray(object?.items) ? object.items.map((e: any) => UserWithStories.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListStoryResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    if (message.items) {
      obj.items = message.items.map((e) => e ? UserWithStories.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListStoryResponse>, I>>(object: I): ListStoryResponse {
    const message = createBaseListStoryResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.items = object.items?.map((e) => UserWithStories.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUserWithStories(): UserWithStories {
  return { user: undefined, stories: [] };
}

export const UserWithStories = {
  encode(message: UserWithStories, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.stories) {
      StoryResponse.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserWithStories {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserWithStories();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        case 2:
          message.stories.push(StoryResponse.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserWithStories {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      stories: Array.isArray(object?.stories) ? object.stories.map((e: any) => StoryResponse.fromJSON(e)) : [],
    };
  },

  toJSON(message: UserWithStories): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    if (message.stories) {
      obj.stories = message.stories.map((e) => e ? StoryResponse.toJSON(e) : undefined);
    } else {
      obj.stories = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserWithStories>, I>>(object: I): UserWithStories {
    const message = createBaseUserWithStories();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.stories = object.stories?.map((e) => StoryResponse.fromPartial(e)) || [];
    return message;
  },
};

function createBaseStoryResponse(): StoryResponse {
  return { story: undefined, hasLiked: false, hasViewed: false };
}

export const StoryResponse = {
  encode(message: StoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.story !== undefined) {
      Story.encode(message.story, writer.uint32(18).fork()).ldelim();
    }
    if (message.hasLiked === true) {
      writer.uint32(40).bool(message.hasLiked);
    }
    if (message.hasViewed === true) {
      writer.uint32(48).bool(message.hasViewed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StoryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.story = Story.decode(reader, reader.uint32());
          break;
        case 5:
          message.hasLiked = reader.bool();
          break;
        case 6:
          message.hasViewed = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoryResponse {
    return {
      story: isSet(object.story) ? Story.fromJSON(object.story) : undefined,
      hasLiked: isSet(object.hasLiked) ? Boolean(object.hasLiked) : false,
      hasViewed: isSet(object.hasViewed) ? Boolean(object.hasViewed) : false,
    };
  },

  toJSON(message: StoryResponse): unknown {
    const obj: any = {};
    message.story !== undefined && (obj.story = message.story ? Story.toJSON(message.story) : undefined);
    message.hasLiked !== undefined && (obj.hasLiked = message.hasLiked);
    message.hasViewed !== undefined && (obj.hasViewed = message.hasViewed);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StoryResponse>, I>>(object: I): StoryResponse {
    const message = createBaseStoryResponse();
    message.story = (object.story !== undefined && object.story !== null) ? Story.fromPartial(object.story) : undefined;
    message.hasLiked = object.hasLiked ?? false;
    message.hasViewed = object.hasViewed ?? false;
    return message;
  },
};

export interface StoryService {
  CreateStory(request: CreateStoryRequest): Promise<StoryResponse>;
  GetStory(request: GetStoryRequest): Promise<StoryResponse>;
  UpdateStory(request: UpdateStoryRequest): Promise<StoryResponse>;
  DeleteStory(request: DeleteStoryRequest): Promise<Empty>;
  FeedStory(request: FeedStoryRequest): Promise<FeedStoryResponse>;
  /** Return stories of a specific user */
  ListUserStory(request: ListUserStoryRequest): Promise<ListUserStoryResponse>;
  ViewStory(request: ViewStoryRequest): Promise<ViewStoryResponse>;
  LikeStory(request: LikeStoryRequest): Promise<LikeStoryResponse>;
  UnlikeStory(request: UnlikeStoryRequest): Promise<UnlikeStoryResponse>;
  SaveStory(request: SaveStoryRequest): Promise<SaveStoryResponse>;
  UnsaveStory(request: UnsaveStoryRequest): Promise<UnsaveStoryResponse>;
}

export class StoryServiceClientImpl implements StoryService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.story.v1.StoryService";
    this.rpc = rpc;
    this.CreateStory = this.CreateStory.bind(this);
    this.GetStory = this.GetStory.bind(this);
    this.UpdateStory = this.UpdateStory.bind(this);
    this.DeleteStory = this.DeleteStory.bind(this);
    this.FeedStory = this.FeedStory.bind(this);
    this.ListUserStory = this.ListUserStory.bind(this);
    this.ViewStory = this.ViewStory.bind(this);
    this.LikeStory = this.LikeStory.bind(this);
    this.UnlikeStory = this.UnlikeStory.bind(this);
    this.SaveStory = this.SaveStory.bind(this);
    this.UnsaveStory = this.UnsaveStory.bind(this);
  }
  CreateStory(request: CreateStoryRequest): Promise<StoryResponse> {
    const data = CreateStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateStory", data);
    return promise.then((data) => StoryResponse.decode(new _m0.Reader(data)));
  }

  GetStory(request: GetStoryRequest): Promise<StoryResponse> {
    const data = GetStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetStory", data);
    return promise.then((data) => StoryResponse.decode(new _m0.Reader(data)));
  }

  UpdateStory(request: UpdateStoryRequest): Promise<StoryResponse> {
    const data = UpdateStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateStory", data);
    return promise.then((data) => StoryResponse.decode(new _m0.Reader(data)));
  }

  DeleteStory(request: DeleteStoryRequest): Promise<Empty> {
    const data = DeleteStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteStory", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  FeedStory(request: FeedStoryRequest): Promise<FeedStoryResponse> {
    const data = FeedStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FeedStory", data);
    return promise.then((data) => FeedStoryResponse.decode(new _m0.Reader(data)));
  }

  ListUserStory(request: ListUserStoryRequest): Promise<ListUserStoryResponse> {
    const data = ListUserStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListUserStory", data);
    return promise.then((data) => ListUserStoryResponse.decode(new _m0.Reader(data)));
  }

  ViewStory(request: ViewStoryRequest): Promise<ViewStoryResponse> {
    const data = ViewStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ViewStory", data);
    return promise.then((data) => ViewStoryResponse.decode(new _m0.Reader(data)));
  }

  LikeStory(request: LikeStoryRequest): Promise<LikeStoryResponse> {
    const data = LikeStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "LikeStory", data);
    return promise.then((data) => LikeStoryResponse.decode(new _m0.Reader(data)));
  }

  UnlikeStory(request: UnlikeStoryRequest): Promise<UnlikeStoryResponse> {
    const data = UnlikeStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnlikeStory", data);
    return promise.then((data) => UnlikeStoryResponse.decode(new _m0.Reader(data)));
  }

  SaveStory(request: SaveStoryRequest): Promise<SaveStoryResponse> {
    const data = SaveStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SaveStory", data);
    return promise.then((data) => SaveStoryResponse.decode(new _m0.Reader(data)));
  }

  UnsaveStory(request: UnsaveStoryRequest): Promise<UnsaveStoryResponse> {
    const data = UnsaveStoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnsaveStory", data);
    return promise.then((data) => UnsaveStoryResponse.decode(new _m0.Reader(data)));
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
