import {Message_Status, MessageItem} from "../../../../gen/ts/cheers/chat/v1/chat_service";

export class ChatMessage {
    id: string = ""
    text: string = ""
    roomId: string = ""
    sender: boolean = false
    status: string = ""
    createTime: number = 0
}

export function toChatMessage(item: MessageItem, status: string | null = null): ChatMessage {
    let chatMessage = new ChatMessage()
    Object.assign(chatMessage, item.message)
    chatMessage.sender = item.sender
    if (status)
        chatMessage.status = status
    return chatMessage
}
