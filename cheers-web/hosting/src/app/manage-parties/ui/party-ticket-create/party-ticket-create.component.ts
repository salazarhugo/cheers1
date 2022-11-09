import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {UntypedFormBuilder, Validators} from "@angular/forms";
import {Ticket} from "../../../shared/data/models/ticket.model";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {ActivatedRoute, ParamMap} from "@angular/router";
import {delay} from "rxjs";


@Component({
    selector: 'app-party-ticket-create',
    templateUrl: './party-ticket-create.component.html',
    styleUrls: ['./party-ticket-create.component.sass']
})
export class PartyTicketCreateComponent implements OnInit {

    @Input() tickets: Ticket[]
    @Output() onCancel = new EventEmitter<void>();
    isLoading = false

    partyId: string | null
    ticketForm = this.fb.group({
            name: ['', [
                Validators.required,
                Validators.maxLength(50),
            ]],
            quantity: [0, [
                Validators.required,
                Validators.min(1),
                Validators.max(500_000),
            ]],
            price: [0, [
                Validators.max(1_000_000)
            ]
            ],
            description: ['', [
                Validators.maxLength(2500),
            ]],
        }
    )

    get name() {
        return this.ticketForm.get('name');
    }

    get quantity() {
        return this.ticketForm.get('quantity');
    }

    constructor(
        private fb: UntypedFormBuilder,
        private fs: AngularFirestore,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit(): void {
        this.route.paramMap.subscribe((params: ParamMap) => {
            const partyId = params.get("id")
            console.log(partyId)
            this.partyId = partyId
        })
    }

    onSave() {
        if (this.partyId == null)
            return
        this.isLoading = true
        const form = this.ticketForm.getRawValue()
        form.price *= 100
        const doc = this.fs.collection("ticketing").doc(this.partyId).collection("tickets").doc().ref
        form.id = doc.id
        doc.set(form)
            .then(r => console.log(r))
            .catch(err => console.log(err))
        this.ticketForm.reset()
        this.isLoading = false
    }
}
