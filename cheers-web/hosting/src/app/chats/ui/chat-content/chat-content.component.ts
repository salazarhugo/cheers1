import {Component, ElementRef, Input, OnChanges, OnInit, SimpleChanges, ViewChild} from '@angular/core';
import {Chat} from "../../../shared/data/models/chat.model";
import {ChatMessage, toChatMessage} from "../../../shared/data/models/chat-message.model";
import {ChatService} from "../../data/chat.service";
import {
    Message,
    Message_Status,
    MessageItem,
    MessageType,
    SendMessageRequest
} from "../../../../gen/ts/cheers/chat/v1/chat_service";
import {UserService} from "../../../shared/data/services/user.service";
import {firstValueFrom} from "rxjs";
import {AngularFireAuth} from "@angular/fire/compat/auth";

@Component({
    selector: 'app-chat-content',
    templateUrl: './chat-content.component.html',
    styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit, OnChanges {

    @ViewChild('scrollMe') private myScrollContainer: ElementRef;
    @Input() room: Chat
    socket: WebSocket
    messages: ChatMessage[] = []
    text: string = "";

    constructor(
        private chatService: ChatService,
        private authService: AngularFireAuth,
        private userService: UserService,
    ) {
    }

    ngOnInit(): void {
    }

    scrollToBottom(): void {
        try {
            this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight;
        } catch (err) {
        }
    }

    async ngOnChanges(changes: SimpleChanges) {
        this.messages = []
        const user = await firstValueFrom(this.userService.user$)
        const roomId = this.room.id

        this.chatService.getRoomMessages(roomId).subscribe(messages => {
            this.messages = messages
        })

        const result = await firstValueFrom(this.authService.idTokenResult)
        const idToken = result?.token!
        this.socket = new WebSocket('wss://chat-service-r3a2dr4u4a-nw.a.run.app/chat?token=' + idToken, [user.id]);

        this.socket.addEventListener('message', (message) => {
            const msg = JSON.parse(message.data)
            msg.sender = user.id == msg.senderId
            const a = this.messages.find(m => m.text == msg.text)
            if (a != undefined) {
                const index = this.messages.indexOf(a)
                this.messages[index] = msg
            } else {
                this.messages.push(msg)
            }
        })
    }

    ngAfterViewChecked() {
        this.scrollToBottom();
    }

    async sendMessage() {
        const user = await firstValueFrom(this.userService.user$)
        const roomId = this.room.id

        const request: SendMessageRequest = {
            clientId: new Date().getTime().toString(),
            replyTo: "",
            roomId: roomId,
            text: this.text
        }
        this.chatService.sendMessage(request).subscribe()
        this.text = ""
    }
}
