import {Injectable} from '@angular/core';
import {map, Observable, of} from "rxjs";
import {Chat, toChat} from "../../shared/data/models/chat.model";
import {HttpClient} from "@angular/common/http";
import {
    CreateRoomRequest, CreateRoomResponse, GetInboxResponse,
    ListRoomMessagesRequest,
    ListRoomMessagesResponse,
    ListRoomResponse,
    Room
} from "../../../gen/ts/cheers/chat/v1/chat_service";
import {environment} from "../../../environments/environment";
import {ChatMessage, toChatMessage} from "../../shared/data/models/chat-message.model";

@Injectable({
    providedIn: 'root'
})
export class ChatService {

    constructor(
        private http: HttpClient,
    ) {
    }

    createRoom(users: string[]): Observable<Chat> {
        return this.http.post<CreateRoomResponse>(`${environment.GATEWAY_URL}/v1/chats`, {
            recipient_users: users
        }).pipe(map(room => toChat(room.room!)))
    }

    getInbox(): Observable<Chat[]> {
        return this.http.get<GetInboxResponse>(`${environment.GATEWAY_URL}/v1/chats/inbox?page=0&page_size=10`)
            .pipe(map(res => res.inbox.map(inbox => toChat(inbox.room!))))
    }

    getRoomMessages(roomId: string): Observable<ChatMessage[]> {
        return this.http.get<ListRoomMessagesResponse>(`${environment.GATEWAY_URL}/v1/chats/${roomId}/messages?page=0&page_size=10`)
            .pipe(
                map(res => res.messages.map(msg => toChatMessage(msg)).sort((a, b) => a.createTime - b.createTime)),
            )
    }

    joinRoom(roomId: string) {
        // this.client.joinRoom(new RoomId({roomId: }))
    }
}
