import { Injectable } from '@angular/core';
import {ApiService} from "./api.service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {firstValueFrom} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class OrderService {

  constructor(
      private api: ApiService,
      private auth: AngularFireAuth,
  ) { }

  async getCurrentUserOrders() {
    const user = await firstValueFrom(this.auth.user)
    return this.api.listOrders(user?.uid!)
  }
}
