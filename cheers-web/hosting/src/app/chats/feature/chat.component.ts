import {Component, OnInit} from '@angular/core';
import {ChatService} from "../data/chat.service";
import {ActivatedRoute, ParamMap} from "@angular/router";
import {Chat} from "../../shared/data/models/chat.model";
import {firstValueFrom, Observable, of} from "rxjs";

@Component({
    selector: 'app-chats',
    templateUrl: './chat.component.html',
    styleUrls: ['./chat.component.sass']
})
export class ChatComponent implements OnInit {

    isLoading = true
    rooms: Chat[] | null = null
    room: Chat | undefined  = undefined

    constructor(
        public route: ActivatedRoute,
        private chatService: ChatService,
    ) {
    }

    async ngOnInit() {
        const rooms = await firstValueFrom(this.chatService.getInbox())
        this.rooms = rooms
        this.isLoading = false
        this.route.firstChild?.paramMap.subscribe(res => {
            console.log(res)
            if (res?.has("id"))
                this.setRoom(res?.get("id")!)
        })
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    setRoom(roomId: string) {
        this.room = this.rooms?.find(rooms => rooms.id == roomId)
    }
}
