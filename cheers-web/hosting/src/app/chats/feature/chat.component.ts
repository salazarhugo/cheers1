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

    $rooms: Observable<Chat[] | null> = of(null)
    $room: Observable<Chat | null> = of(null)

    constructor(
        private route: ActivatedRoute,
        private chatService: ChatService,
    ) {
    }

    ngOnInit(): void {
        this.$rooms = this.chatService.getRooms()

        this.route.firstChild?.paramMap.subscribe((params: ParamMap) => {
            const roomId = params.get("id")
            if (roomId) {
                this.$room = this.chatService.getRoom(roomId)
            }
        })
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }
}
