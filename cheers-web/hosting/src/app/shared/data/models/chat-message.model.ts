import {MessageItem} from "../../../../gen/ts/cheers/chat/v1/chat_service";

export class ChatMessage {
    id: string = ""
    text: string = ""
    roomId: string = ""
    sender: boolean = false
    createTime: number = 0
}

export function toChatMessage(item: MessageItem): ChatMessage {
    let chatMessage = new ChatMessage()
    Object.assign(chatMessage, item.message)
    chatMessage.sender = item.sender
    return chatMessage
}
