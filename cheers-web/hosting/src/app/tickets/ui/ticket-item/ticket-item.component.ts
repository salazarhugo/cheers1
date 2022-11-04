import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Ticket} from "../../../shared/data/models/ticket.model";

@Component({
  selector: 'app-ticket-item',
  templateUrl: './ticket-item.component.html',
  styleUrls: ['./ticket-item.component.sass']
})
export class TicketItemComponent implements OnInit {

  @Input() ticket: Ticket
  @Output() onDelete = new EventEmitter<Ticket>();

  constructor() {
  }

  ngOnInit(): void {
  }

  deleteTicket() {
    this.onDelete.emit(this.ticket)
  }

  copyTicket() {

  }

  editTicket() {

  }
}
