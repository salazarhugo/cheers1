import {Injectable} from '@angular/core';
import {map, Observable, of} from "rxjs";
import {Chat, toChat} from "../../shared/data/models/chat.model";
import {HttpClient} from "@angular/common/http";
import {ListRoomResponse, Room} from "../../../gen/ts/cheers/chat/v1/chat_service";
import {environment} from "../../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class ChatService {

  constructor(
      private http: HttpClient,
  ) {
    // const socket = new WebSocket('ws://localhost:8081/ws');
    //
    // socket.addEventListener('open', function (event) {
    //   socket.send("Hello World!");
    // })
    //
    // socket.addEventListener('message', function (message) {
    //   console.log(message.data)
    // })
  }

  getRooms(): Observable<Chat[]> {
    return this.http.get<ListRoomResponse>(`${environment.GATEWAY_URL}/v1/chats?page=0&page_size=10`)
        .pipe(map(res => res.rooms.map(room => toChat(room))))
  }

  joinRoom(roomId: string) {
    // this.client.joinRoom(new RoomId({roomId: }))
  }
}
