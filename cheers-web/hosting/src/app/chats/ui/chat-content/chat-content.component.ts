import {Component, Input, OnInit} from '@angular/core';
import {Chat} from "../../../shared/data/models/chat.model";
import {ChatMessage} from "../../../shared/data/models/chat-message.model";
import {ChatService} from "../../data/chat.service";
import {Message, Message_Status, MessageType} from "../../../../gen/ts/cheers/chat/v1/chat_service";

@Component({
    selector: 'app-chat-content',
    templateUrl: './chat-content.component.html',
    styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit {

    @Input() room: Chat
    socket: WebSocket
    messages: ChatMessage[] = []
    text: string = "";

    constructor(
        private chatService: ChatService,
    ) {
    }

    ngOnInit(): void {
        const roomId = this.room.id.replace(':', '')
        this.chatService.getRoomMessages(roomId).subscribe(messages => {
            this.messages = messages
        })

        this.socket = new WebSocket('ws://localhost:8000/ws', roomId);

        this.socket.addEventListener('message', (message) => {
            const msg = JSON.parse(message.data)
            this.messages.push(msg)
            console.log(msg)
        })
    }

    sendMessage() {
        const roomId = this.room.id.replace(':', '')
        const message: Message = {
            createTime: 0,
            likeCount: 0,
            picture: "",
            senderId: "55FEvHawinQCa9jgH7ZdWESR3ri2",
            senderName: "Hugo",
            senderPicture: "",
            senderUsername: "",
            status: Message_Status.DELIVERED,
            type: MessageType.TEXT,
            id: "",
            roomId: roomId,
            text: this.text
        }
        const bytes = Message.encode(message).finish()
        this.socket.send(bytes);
        this.text = ""
    }
}
