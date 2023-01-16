import {PartyAnswer, PartyItem} from "../../../../gen/ts/cheers/party/v1/party_service";
import {Party} from "./party.model";
import {Story as StoryPb} from "../../../../gen/ts/cheers/type/story/story";


export class Story {
    id: string = ""
    authorId: string = ""
    username: string = ""
    verified: boolean = false
    profilePictureUrl: string = ""
    seen: boolean = false
    created: number = 0
    relativeTime: string = ""
    privacy: string = ""
    photo: string = ""
    videoUrl: string = ""
    latitude: number = 0
    longitude: number = 0
    altitude: number = 0
    locationName: string = ""
    tagUsersId: string[] = []
    type: string = ""
}

export function toStory(value: StoryPb): Story {
    const story = new Story()
    Object.assign(story, value)
    return story
}
