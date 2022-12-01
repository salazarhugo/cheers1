import {Component, Input, OnInit} from '@angular/core';
import {ChatMessage} from "../../../shared/data/models/chat-message.model";

@Component({
  selector: 'app-chat-message-item',
  templateUrl: './chat-message-item.component.html',
  styleUrls: ['./chat-message-item.component.sass']
})
export class ChatMessageItemComponent implements OnInit {

  @Input() message: ChatMessage

  constructor() { }

  ngOnInit(): void {
  }

}
