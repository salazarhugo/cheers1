import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {PaymentIntent} from "../../shared/data/models/payment-intent.model";

@Injectable({
  providedIn: 'root'
})
export class PaymentService {

  BASE_URL = "https://payment-service-r3a2dr4u4a-nw.a.run.app"

  constructor(
      private http: HttpClient,
  ) {
  }

  createPaymentIntent(order: any, partyId: string): Observable<PaymentIntent> {
    return this.http.post<PaymentIntent>(`${this.BASE_URL}/createPaymentIntent?partyId=${partyId}`, order)
  }
}
