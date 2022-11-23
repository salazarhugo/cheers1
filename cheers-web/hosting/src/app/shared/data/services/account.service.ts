import { Injectable } from '@angular/core';
import {ApiService} from "./api.service";

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(
      private service: ApiService,
  ) { }

  getCurrentAccount() {
    return this.service.getAccount("")
  }
}
