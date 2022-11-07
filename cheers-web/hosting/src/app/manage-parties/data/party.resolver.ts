import { Injectable } from '@angular/core';
import {
  Router, Resolve,
  RouterStateSnapshot,
  ActivatedRouteSnapshot
} from '@angular/router';
import { Observable, of } from 'rxjs';
import {PartyService} from "../../parties/data/party.service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {ApiService} from "../../shared/data/services/api.service";

@Injectable({
  providedIn: 'root'
})
export class PartyResolver implements Resolve<boolean> {
  constructor(
      public partyService: PartyService,
      private router: Router,
      private afAuth: AngularFireAuth,
      private api: ApiService,
  ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> {
    const partyId = route.paramMap.get("id")
    // this.api.getParty(partyId)
    return of(true);
  }
}
