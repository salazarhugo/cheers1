import {Component, Input, OnInit} from '@angular/core';
import {Chat} from "../../../shared/data/models/chat.model";
import {ChatMessage} from "../../../shared/data/models/chat-message.model";
import {ChatService} from "../../data/chat.service";

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
        this.chatService.getRoomMessages(this.room.id).subscribe(messages => {
            this.messages = messages
        })

        const roomId = this.room.id.replace(':', '')
        this.socket = new WebSocket('ws://localhost:8000/ws', roomId);

        this.socket.addEventListener('message', (message) => {
            const msg = JSON.parse(message.data)
            this.messages.push(msg)
            console.log(msg)
        })
    }

    sendMessage() {
        const roomId = this.room.id.replace(':', '')
        const message: ChatMessage = {
            id: "",
            roomId: roomId,
            text: this.text,
            isCreator: false,
        }
        this.socket.send(JSON.stringify(message));
        this.text = ""
    }
}
