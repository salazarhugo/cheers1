import { Component, OnInit } from '@angular/core';
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {UserService} from "../../data/services/user.service";
import {Observable, tap} from "rxjs";
import {Ticket} from "../../data/models/ticket.model";

@Component({
  selector: 'app-tickets',
  templateUrl: './tickets.component.html',
  styleUrls: ['./tickets.component.sass']
})
export class TicketsComponent implements OnInit {

  isLoading: boolean = true

  constructor(
      private userService: UserService,
  ) { }

    tickets: Ticket[] | null

    ngOnInit(): void {
      this.userService.getUserTickets().subscribe(value => {
        this.tickets = value
        this.isLoading = false
      })
  }

}
