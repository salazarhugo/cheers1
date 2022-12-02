import {Component, ElementRef, Input, OnInit, ViewChild} from '@angular/core';
import {Chat} from "../../../shared/data/models/chat.model";
import {ChatMessage} from "../../../shared/data/models/chat-message.model";
import {ChatService} from "../../data/chat.service";
import {Message, Message_Status, MessageType} from "../../../../gen/ts/cheers/chat/v1/chat_service";
import {UserService} from "../../../shared/data/services/user.service";
import {firstValueFrom} from "rxjs";
import {toUnixTimestamp} from "../../../parties/ui/party-form/party-form.component";

@Component({
    selector: 'app-chat-content',
    templateUrl: './chat-content.component.html',
    styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit {

    @ViewChild('scrollMe') private myScrollContainer: ElementRef;
    @Input() room: Chat
    socket: WebSocket
    messages: ChatMessage[] = []
    text: string = "";

    constructor(
        private chatService: ChatService,
        private userService: UserService,
    ) {
    }

    scrollToBottom(): void {
        try {
            this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight;
        } catch (err) {
        }
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

    ngAfterViewChecked() {
        this.scrollToBottom();
    }

    async sendMessage() {
        const user = await firstValueFrom(this.userService.user$)
        const roomId = this.room.id.replace(':', '')
        const message: Message = {
            createTime: toUnixTimestamp(new Date()),
            likeCount: 0,
            picture: "",
            senderId: user.id,
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
