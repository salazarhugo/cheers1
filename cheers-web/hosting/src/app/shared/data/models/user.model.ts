import {UserItem} from "../../../../gen/ts/cheers/type/user/user";


export class UserModel {
    id: string = ""
    name: string = "Demo";
    email: string = "";
    website: string = "";
    username: string = "admin";
    picture: string = "";
    bio: string = "";
    postCount: number = 0
    followers: number = 0
    following: number = 0
    phoneNumber: string = ""
    verified: boolean = true
    online: boolean = false
    followBack: boolean = false
    created: number = 0
    moderator: boolean = false
    admin: boolean = false
    business: boolean = false
    storyState: StoryState = StoryState.EMPTY
}

export enum StoryState {
    EMPTY,
    SEEN,
    UNSEEN,
    FRIENDS,
}

export class UserResponse {
    user: UserModel
    postCount: number
    followersCount: number
    followingCount: number
    hasFollowed: boolean
    storyState: string
}

export function toUser(res: UserResponse): UserModel | null {
    const user = res.user
    if (user == null)
        return null
    user.postCount = res.postCount
    user.following = res.followingCount
    user.followers = res.followersCount
    user.followBack = res.hasFollowed
    const strEnum = res.storyState as unknown as StoryState;
    user.storyState = strEnum
    return user
}

export function toUserFromUserItem(value: UserItem): UserModel {
    const user = new UserModel()
    Object.assign(user, value)
    return user
}
