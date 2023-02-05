import {UserWithStories} from "../../../../gen/ts/cheers/story/v1/story_service";


export class Story {
    id: string = ""
    authorId: string = ""
    username: string = ""
    verified: boolean = false
    picture: string = ""
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


export function toStory(userWithStories: UserWithStories): Story[] {
    const stories: Story[] = []
    const user = userWithStories.user
    userWithStories.stories.forEach(storyRemote => {
        const story = new Story()
        Object.assign(story, storyRemote.story)
        story.picture = user?.picture!
        story.username = user?.username!
        stories.push(story)
    })

    return stories
}
