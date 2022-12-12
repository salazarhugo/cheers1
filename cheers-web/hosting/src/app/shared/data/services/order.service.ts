import { Injectable } from '@angular/core';
import {ApiService} from "./api.service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {firstValueFrom, map, Observable} from "rxjs";
import {ListUserOrdersResponse, Order} from "../../../../gen/ts/cheers/order/v1/order_service";
import {environment} from "../../../../environments/environment";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class OrderService {

  constructor(
      private http: HttpClient,
      private api: ApiService,
      private auth: AngularFireAuth,
  ) { }

  async getCurrentUserOrders() {
    const user = await firstValueFrom(this.auth.user)
    return this.api.listUserOrders(user?.uid!)
  }

  getOrganizationOrders(query: string): Observable<Order[]> {
    return this.http.get<ListUserOrdersResponse>(`${environment.GATEWAY_URL}/v1/orders/organization?query=${query}`)
        .pipe(map(r => r.orders))
  }
}
