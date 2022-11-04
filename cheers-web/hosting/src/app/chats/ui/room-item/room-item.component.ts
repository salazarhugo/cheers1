import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-room-item',
  templateUrl: './room-item.component.html',
  styleUrls: ['./room-item.component.sass']
})
export class RoomItemComponent implements OnInit {

  @Input() room: any

  constructor() {
  }

  ngOnInit(): void {
  }

  onImgError(event: any) {
    event.target.src = 'assets/default_profile_picture.png';
  }
}
