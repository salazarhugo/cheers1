import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {map, Observable} from "rxjs";
import {PaymentIntent} from "../../shared/data/models/payment-intent.model";
import {ListOrderResponse} from "../../../gen/ts/cheers/order/v1/order_service";
import {environment} from "../../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class PaymentService {

  constructor(
      private http: HttpClient,
  ) {
  }

  createPaymentIntent(
      tickets: any,
      partyId: string,
      firstName: string,
      lastName: string,
      email: string,
  ): Observable<PaymentIntent> {
    return this.http.post<PaymentIntent>(`${environment.GATEWAY_URL}/v1/payments`, {
      party_id: partyId,
      tickets: tickets,
      email: email,
      first_name: firstName,
      last_name: lastName,
    })
  }
}
