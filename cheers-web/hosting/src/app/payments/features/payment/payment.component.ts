import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap} from "@angular/router";
import {loadStripe, PaymentIntent} from "@stripe/stripe-js";
import {environment} from "../../../../environments/environment";

@Component({
    selector: 'app-payment',
    templateUrl: './payment.component.html',
    styleUrls: ['./payment.component.sass']
})
export class PaymentComponent implements OnInit {
    stripePromise = loadStripe(environment.stripe.public_key);
    isLoading = false
    clientSecret: string | null = null
    paymentIntent: PaymentIntent | undefined

    constructor(
        private route: ActivatedRoute,
    ) {
        this.route.paramMap.subscribe((params: ParamMap) => {
            const clientSecret = params.get("clientSecret")
            this.clientSecret = clientSecret
        })

    }

    async ngOnInit() {
        const stripe = await this.stripePromise;

        if (stripe == null)
            return

        stripe.retrievePaymentIntent(this.clientSecret!!).then(({paymentIntent}) => {
            this.paymentIntent = paymentIntent
        });
    }

}
