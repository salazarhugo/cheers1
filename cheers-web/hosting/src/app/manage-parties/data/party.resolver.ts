import { Injectable } from '@angular/core';
import {
  Router, Resolve,
  RouterStateSnapshot,
  ActivatedRouteSnapshot
} from '@angular/router';
import {firstValueFrom, Observable, of} from 'rxjs';
import {PartyService} from "../../parties/data/party.service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {ApiService} from "../../shared/data/services/api.service";
import {Party} from "../../shared/data/models/party.model";
import {UserService} from "../../shared/data/services/user.service";

@Injectable({
  providedIn: 'root'
})
export class PartyResolver implements Resolve<Party> {
  constructor(
      public userService: UserService,
      public partyService: PartyService,
      private router: Router,
      private afAuth: AngularFireAuth,
  ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<Party> {
    return new Promise(async (resolve, reject) => {

      const partyId = route.firstChild?.paramMap.get("id")
      if (!partyId)
        return reject();

      const party = await firstValueFrom(this.partyService.getParty(partyId))
      // const user = await firstValueFrom(this.userService.user$)

      // Only the host of the party can edit it or admin
      // if (!party.isHost && !user.moderator) {
      //   return reject("You are not the host of the party");
      // }

      this.partyService.setParty(party)
      console.log(party)
      resolve(party)
    })
  }
}
