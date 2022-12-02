import {Injectable} from '@angular/core';
import {map, Observable, of} from "rxjs";
import {Chat, toChat} from "../../shared/data/models/chat.model";
import {HttpClient} from "@angular/common/http";
import {
    ListRoomMessagesRequest,
    ListRoomMessagesResponse,
    ListRoomResponse,
    Room
} from "../../../gen/ts/cheers/chat/v1/chat_service";
import {environment} from "../../../environments/environment";
import {ChatMessage, toChatMessage} from "../../shared/data/models/chat-message.model";
import {sortDescendingPriority} from "@angular/flex-layout";

@Injectable({
    providedIn: 'root'
})
export class ChatService {

    socket: WebSocket

    constructor(
        private http: HttpClient,
    ) {
    }


    getRooms(): Observable<Chat[]> {
        return this.http.get<ListRoomResponse>(`${environment.GATEWAY_URL}/v1/chats?page=0&page_size=10`)
            .pipe(map(res => res.rooms.map(room => toChat(room))))
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
