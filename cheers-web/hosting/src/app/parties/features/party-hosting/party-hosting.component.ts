import { Component, OnInit } from '@angular/core';
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../data/party.service";
import {firstValueFrom, Observable, of} from "rxjs";
import {Party} from "../../../shared/data/models/party.model";

@Component({
  selector: 'app-party-hosting',
  templateUrl: './party-hosting.component.html',
  styleUrls: ['./party-hosting.component.sass']
})
export class PartyHostingComponent implements OnInit {

  $parties: Observable<Party[] | null> = of(null)

  constructor(
      private userService: UserService,
      private partyService: PartyService,
  ) {
  }

  async ngOnInit() {
    const user = await firstValueFrom(this.userService.user$)
    console.log(user)
    this.$parties = this.partyService.getMyParties(user.id)
    console.log(this.$parties)
  }

}
