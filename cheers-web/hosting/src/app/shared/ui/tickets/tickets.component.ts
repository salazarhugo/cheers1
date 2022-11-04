import { Component, OnInit } from '@angular/core';
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {UserService} from "../../data/services/user.service";
import {Observable} from "rxjs";
import {Ticket} from "../../data/models/ticket.model";

@Component({
  selector: 'app-tickets',
  templateUrl: './tickets.component.html',
  styleUrls: ['./tickets.component.sass']
})
export class TicketsComponent implements OnInit {

  constructor(
      private userService: UserService,
  ) { }

    tickets$: Observable<Ticket[]>

    ngOnInit(): void {
      this.tickets$ = this.userService.getUserTickets()
  }

}
