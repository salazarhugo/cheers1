import {Component, Input, OnInit} from '@angular/core';
import {Chat} from "../../../shared/data/models/chat.model";

@Component({
  selector: 'app-chat-content',
  templateUrl: './chat-content.component.html',
  styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit {

  @Input() room: Chat

  constructor() {
  }

  ngOnInit(): void {
  }

}
