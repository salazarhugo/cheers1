import {Component, OnInit} from '@angular/core';
import {loadStripe, PaymentIntent} from "@stripe/stripe-js";
import {environment} from "../../../../environments/environment";
import {ActivatedRoute} from "@angular/router";
import {first} from "rxjs/operators";
import {AnimationOptions} from "ngx-lottie";

declare var Stripe: any;

@Component({
  selector: 'app-complete',
  templateUrl: './complete.component.html',
  styleUrls: ['./complete.component.css']
})
export class CompleteComponent implements OnInit {
  stripePromise = loadStripe(environment.stripe.public_key);
  stripe: any; // : stripe.Stripe;
  status: PaymentIntent.Status | undefined = undefined
  options: AnimationOptions = {
    path: '/assets/lottie/approved.json',
    loop: false,
  };

  constructor(
      private route: ActivatedRoute,
  ) {
  }

  async ngOnInit() {
    const stripe = await this.stripePromise;

    if (stripe == null)
      return

    this.stripe = Stripe(environment.stripe.public_key);
    const clientSecret = this.route.snapshot.queryParams['payment_intent_client_secret']

    stripe.retrievePaymentIntent(clientSecret).then(({paymentIntent}) => {
      const message = document.querySelector('#message')
      this.status = paymentIntent?.status

      switch (paymentIntent?.status) {
        case 'succeeded':
          break;

        case 'processing':
          // message.innerText = "Payment processing. We'll update you when payment is received.";
          break;

        case 'requires_payment_method':
          // message.innerText = 'Payment failed. Please try another payment method.';
          // Redirect your user back to your payment page to attempt collecting
          // payment again
          break;

        default:
          // message.innerText = 'Something went wrong.';
          break;
      }
    });
  }

}
