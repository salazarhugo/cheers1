import {Component, OnInit} from '@angular/core';
import {UntypedFormBuilder} from "@angular/forms";
import {Ticket} from "../../../shared/data/models/ticket.model";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {ActivatedRoute, ParamMap} from "@angular/router";
import {tick} from "@angular/core/testing";
import {filter, Observable, of} from "rxjs";


@Component({
    selector: 'app-party-tickets',
    templateUrl: './party-tickets.component.html',
    styleUrls: ['./party-tickets.component.sass']
})
export class PartyTicketsComponent implements OnInit {

    $tickets: Observable<Ticket[] | undefined>
    partyId: string

    constructor(
        private fb: UntypedFormBuilder,
        private fs: AngularFirestore,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit(): void {
        this.route.paramMap.subscribe((params: ParamMap) => {
            const partyId = params.get("id")
            if (partyId == null)
                return
            this.partyId = partyId
            this.$tickets = this.fs.collection("ticketing").doc(partyId).collection<Ticket>("tickets").valueChanges()
        })
    }

    deleteTicket(ticket: Ticket) {
        this.fs.collection("ticketing").doc(this.partyId).collection("tickets").doc(ticket.id).delete()
            .then(r => console.log("deleted"))
        // const index = this.tickets.indexOf(ticket, 0);
        // if (index > -1) {
        //     this.tickets.splice(index, 1);
        // }
    }
}
