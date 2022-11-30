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

  rooms: Chat[] = [
    {
      id: "room01",
      name: "Dora",
      username: "dora",
      picture: "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/gettyimages-694016643-1523641375.jpg?crop=1xw:1xh;center,top&resize=480:*",
    },
    {
      id: "room02",
      name: "Marta",
      username: "dora",
      picture: "https://i.pinimg.com/236x/42/0b/d5/420bd5f2f52c1c895a4e395d59172232--freida-pinto-most-beautiful-women.jpg",
    },
    {
      id: "room03",
      name: "Adrien",
      username: "adrien94",
      picture: "https://i.pinimg.com/736x/5b/7b/ce/5b7bce45aaf4bf028ece82c28140924a.jpg",
    },
    {
      id: "room04",
      name: "Dora",
      username: "dora",
      picture: "https://i.pinimg.com/474x/6c/db/0a/6cdb0ac6cf41f7bfd71a91420d84ae3a.jpg",
    },
  ]

  getRooms(): Observable<Chat[]> {
    return this.http.get<ListRoomResponse>(`${environment.GATEWAY_URL}/v1/chats?page=0&page_size=10`)
        .pipe(map(res => res.rooms.map(room => toChat(room))))
  }

  getRoom(roomId: string): Observable<Chat | null> {
    const room = this.rooms.find(r => r.id == roomId)
    if (room == undefined)
      return of(null)
    return of(room)
  }

  joinRoom(roomId: string) {
    // this.client.joinRoom(new RoomId({roomId: }))
  }
}
