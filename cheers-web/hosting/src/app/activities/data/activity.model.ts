import {Activity as ActivityPb} from "../../../gen/ts/cheers/activity/v1/activity_service";

export class Activity {
    picture: string | null
    media: string | null
    text: string = "Lara has started following you."
}

export function toActivity(value: ActivityPb): Activity {
    const activity = new Activity()
    Object.assign(activity, value)
    return activity
}
