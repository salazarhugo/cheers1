import {HttpClient} from "@angular/common/http";
import {map, Observable} from "rxjs";
import {environment} from "../../../environments/environment";
import {CreatePaymentResponse, RefundPaymentResponse} from "../../../gen/ts/cheers/payment/v1/payment_service";
import {Injectable} from "@angular/core";

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
    ): Observable<string> {
        return this.http.post<CreatePaymentResponse>(`${environment.GATEWAY_URL}/v1/payments`, {
            party_id: partyId,
            tickets: tickets,
            email: email,
            first_name: firstName,
            last_name: lastName,
        }).pipe(map(r => r.clientSecret))
    }

    refundPayment(paymentIntentId: string): Observable<RefundPaymentResponse> {
        return this.http.post<RefundPaymentResponse>(`${environment.GATEWAY_URL}/v1/payments/${paymentIntentId}/refund`, {})
    }
}
