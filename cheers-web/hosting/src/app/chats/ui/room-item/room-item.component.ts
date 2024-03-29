import {Component, Input, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {Chat} from "../../../shared/data/models/chat.model";

@Component({
  selector: 'app-room-item',
  templateUrl: './room-item.component.html',
  styleUrls: ['./room-item.component.sass']
})
export class RoomItemComponent implements OnInit {

  @Input() room: Chat

  constructor(
      public router: Router,
  ) {
  }

  ngOnInit(): void {
  }

  onImgError(event: any) {
    event.target.src = 'assets/default_profile_picture.png';
  }
}
