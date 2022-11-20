import { Injectable } from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";

@Injectable({
  providedIn: 'root'
})
export class NotificationService {

  constructor(
      private service: ApiService,
  ) { }

  createToken(token: string) {
    return this.service.createToken(token)
  }
}
