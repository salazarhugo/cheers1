export class User {
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
    storyState: StoryState = StoryState.EMPTY
}

export enum StoryState {
    EMPTY,
    SEEN,
    UNSEEN,
    FRIENDS,
}

export class UserResponse {
    user: User
    postCount: number
    followersCount: number
    followingCount: number
    hasFollowed: boolean
    storyState: string
}
