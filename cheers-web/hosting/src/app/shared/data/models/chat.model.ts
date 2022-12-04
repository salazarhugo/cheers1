import {Room} from "../../../../gen/ts/cheers/chat/v1/chat_service";

export class Chat {
    id: string = ""
    name: string = ""
    username: string = ""
    picture: string = ""
    latestMessageTime: number = 0
}

export function toChat(room: Room): Chat {
    let chat = new Chat()
    Object.assign(chat, room)
    return chat
}