import {Component, ElementRef, Input, OnChanges, OnInit, SimpleChanges, ViewChild} from '@angular/core';
import {Chat} from "../../../shared/data/models/chat.model";
import {ChatMessage, toChatMessage} from "../../../shared/data/models/chat-message.model";
import {ChatService} from "../../data/chat.service";
import {Message, Message_Status, MessageItem, MessageType} from "../../../../gen/ts/cheers/chat/v1/chat_service";
import {UserService} from "../../../shared/data/services/user.service";
import {firstValueFrom} from "rxjs";
import {toUnixTimestamp} from "../../../parties/ui/party-form/party-form.component";
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
        // this.socket = new WebSocket('ws://localhost:8000/chat', [roomId]);
        this.socket = new WebSocket('wss://chat-service-r3a2dr4u4a-nw.a.run.app/chat', [roomId]);

        this.socket.addEventListener('message', (message) => {
            const msg = JSON.parse(message.data)
            msg.sender = user.id == msg.senderId
            const a = this.messages.find(m => m.text == msg.text)
            if (a != undefined) {
                const index = this.messages.indexOf(a)
                this.messages[index] = msg
            }
            else {
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
        const item: MessageItem = {
            message: {
                createTime: toUnixTimestamp(new Date()),
                likeCount: 0,
                picture: "",
                senderId: user.id,
                senderName: "Hugo",
                senderPicture: "",
                senderUsername: "",
                status: Message_Status.EMPTY,
                type: MessageType.TEXT,
                id: "",
                roomId: roomId,
                text: this.text
            },
            sender: true,
            liked: false,
        }

        const chatMessage = toChatMessage(item, "SCHEDULED")
        this.messages.push(chatMessage)
        const bytes = Message.encode(item.message!).finish()
        console.log(bytes)
        this.socket.send(bytes);
        this.text = ""
    }
}
