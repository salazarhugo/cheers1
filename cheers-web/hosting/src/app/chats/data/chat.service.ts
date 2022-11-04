import {Injectable} from '@angular/core';
import {Observable, of} from "rxjs";
import {Chat} from "../../shared/data/models/chat.model";

@Injectable({
  providedIn: 'root'
})
export class ChatService {

  constructor(
      // private client: ChatServiceClient,
  ) {
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
    return of(this.rooms) //this.client.getRooms(this.rooms)
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
