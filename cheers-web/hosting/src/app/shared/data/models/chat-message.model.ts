import {Message} from "../../../../gen/ts/cheers/chat/v1/chat_service";

export class ChatMessage {
    id: string = ""
    text: string = ""
    roomId: string = ""
    isCreator: boolean = false
}

export function toChatMessage(message: Message): ChatMessage {
    let chatMessage = new ChatMessage()
    Object.assign(chatMessage, message)
    chatMessage.text = message.message
    return chatMessage
}
