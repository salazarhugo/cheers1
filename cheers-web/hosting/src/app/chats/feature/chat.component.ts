import {Component, OnInit} from '@angular/core';
import {ChatService} from "../data/chat.service";
import {ActivatedRoute, ParamMap} from "@angular/router";
import {Chat} from "../../shared/data/models/chat.model";
import {Observable, of} from "rxjs";

@Component({
    selector: 'app-chats',
    templateUrl: './chat.component.html',
    styleUrls: ['./chat.component.sass']
})
export class ChatComponent implements OnInit {

    rooms: Chat[] | null = null
    room: Chat | undefined  = undefined

    constructor(
        private route: ActivatedRoute,
        private chatService: ChatService,
    ) {
    }

    ngOnInit(): void {
        this.chatService.getRooms().subscribe(rooms => {
            this.rooms = rooms
        })
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    setRoom(roomId: string) {
        this.room = this.rooms?.find(rooms => rooms.id == roomId)
    }
}
