import { Component, OnInit } from '@angular/core';
import {Observable} from "rxjs";
import {Ticket} from "../../../shared/data/models/ticket.model";
import {UntypedFormBuilder} from "@angular/forms";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {ActivatedRoute, ParamMap} from "@angular/router";

@Component({
    selector: 'app-manage-party-tickets',
    templateUrl: './manage-party-tickets.component.html',
    styleUrls: ['./manage-party-tickets.component.sass']
})
export class ManagePartyTicketsComponent implements OnInit {

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
