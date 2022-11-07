import {Component, OnInit} from '@angular/core';
import {environment} from "../../../../environments/environment";
import {loadStripe} from "@stripe/stripe-js";
import {Observable, of} from "rxjs";
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {PartyService} from "../../../parties/data/party.service";
import {PaymentService} from "../../../payments/data/payment.service";
import {UntypedFormBuilder, UntypedFormControl, UntypedFormGroup, Validators} from "@angular/forms";
import {Ticket} from "../../data/models/ticket.model";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {Party} from "../../data/models/party.model";

@Component({
    selector: 'app-ticketing',
    templateUrl: './ticketing.component.html',
    styleUrls: ['./ticketing.component.sass']
})
export class TicketingComponent implements OnInit {

    stripePromise = loadStripe(environment.stripe.public_key);
    $party: Observable<Party | null> = of(null)
    hasAcceptedTerms = false
    total = 0

    tickets: Ticket[]
    ticketingForm = new UntypedFormGroup({});
    partyId: string | null

    userForm = this.fb.group({
            name: ['', [
                Validators.required,
                Validators.maxLength(50),
            ]],
            email: ['', [
                Validators.required,
                Validators.maxLength(50),
            ]],
        }
    )

    constructor(
        private route: ActivatedRoute,
        private partyService: PartyService,
        private paymentService: PaymentService,
        private router: Router,
        public afAuth: AngularFireAuth,
        private fb: UntypedFormBuilder,
    ) {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.partyId = params.get("id")
            if (this.partyId)
                this.$party = this.partyService.getParty(this.partyId)
        })
        this.afAuth.user.subscribe(user => {
            if (user != null) {
                console.log(user)
                this.userForm.patchValue(
                    {
                        email: user.email,
                        name: user.displayName
                    }
                )
            }
        })
    }

    async getPartyTickets() {
        this.tickets = await this.partyService.getPartyTickets(this.partyId!!)
        this.tickets.forEach(ticket => {
            this.ticketingForm.addControl(ticket.id, new UntypedFormControl(0))
        })
    }

    ngOnInit() {
        this.ticketingForm.valueChanges.subscribe(value => {
            let sum = 0;
            for (const key in value) {
                if (value.hasOwnProperty(key)) {
                    const pack = this.tickets.find(p => p.id == key)
                    if (pack == null) return
                    sum = Number(value[key]) * pack.price + sum;
                }
            }
            this.total = sum;
        })
        this.getPartyTickets().then()
    }

    decrement(pack: string) {
        const formControl = this.ticketingForm.get(pack)!!
        if (formControl.value == 0)
            return
        formControl.setValue(formControl.value - 1)
    }

    increment(pack: string) {
        const formControl = this.ticketingForm.get(pack)!!
        formControl.setValue(formControl.value + 1)
    }

    async checkout() {
        const order = this.ticketingForm.getRawValue()
        console.log(order)
        const paymentIntent = await this.paymentService.createPaymentIntent(order, this.partyId!!).toPromise()
        const clientSecret = paymentIntent?.clientSecret
        await this.router.navigate(['payment', clientSecret])
    }
}
