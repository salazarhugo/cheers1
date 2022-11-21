import {Component, ElementRef, Input, OnInit, ViewChild} from '@angular/core';
import {AuthService} from "../../../shared/data/services/auth.service";
import {AngularFireFunctions} from "@angular/fire/compat/functions";
import {environment} from "../../../../environments/environment";
import {loadStripe} from "@stripe/stripe-js";
import {PaymentService} from "../../data/payment.service";
import {ThemeService} from "../../../core/data/theme.service";
import {firstValueFrom} from "rxjs";

declare var Stripe: any;

@Component({
    selector: 'app-elements',
    templateUrl: './elements.component.html',
    styleUrls: ['./elements.component.sass']
})
export class ElementsComponent implements OnInit {

    stripePromise = loadStripe(environment.stripe.public_key);
    @Input() amount: number;
    @Input() clientSecret: string;
    @Input() description: string;
    @ViewChild('paymentElement') paymentElement: ElementRef;
    stripe: any; // : stripe.Stripe;
    card: any;
    cardErrors: any;
    isLoading = false;
    confirmation: any;
    elements: any

    constructor(
        private auth: AuthService,
        private paymentService: PaymentService,
        private themeService: ThemeService,
    ) {
    }

    async ngOnInit() {
        const stripe = await this.stripePromise;

        if (stripe == null)
            return

        this.stripe = Stripe(environment.stripe.public_key);

        const isDarkMode = await firstValueFrom(this.themeService.isDarkTheme)
        console.log("Fwf")
        const options = {
            clientSecret: this.clientSecret,
            appearance: {
                theme: (isDarkMode) ? "night" : "default",
            },
        };

        this.elements = this.stripe.elements(options);

        this.card = this.elements.create('payment');
        this.card.mount(this.paymentElement.nativeElement);

        this.card.addEventListener('change', (error: any) => {
            this.cardErrors = error && error.message;
        });
    }

    async handleForm(e: any) {
        e.preventDefault();

        this.isLoading = true;
        // const stripe = Stripe(environment.stripe.public_key)

        const {error} = await this.stripe.confirmPayment({
            elements: this.elements,
            confirmParams: {
                return_url: window.location.origin + "/payment/complete",
            },
        })

        if (error.type === "card_error" || error.type === "validation_error") {
            console.log(error)
        } else {
            console.log(error)
        }
        this.isLoading = false
    }
}
