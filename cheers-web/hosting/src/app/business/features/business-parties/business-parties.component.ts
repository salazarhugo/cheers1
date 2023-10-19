import { Component, OnInit } from '@angular/core';
import {Observable, of} from "rxjs";
import {Party} from "../../../shared/data/models/party.model";
import {User} from "../../../shared/data/models/user.model";
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../../parties/data/party.service";
import {KeyValue} from "@angular/common";

@Component({
  selector: 'app-business-parties',
  templateUrl: './business-parties.component.html',
  styleUrls: ['./business-parties.component.sass']
})
export class BusinessPartiesComponent implements OnInit {

  $user: Observable<User | null> = of(null)
  $myParties: Observable<Party[] | null> = of(null)

  constructor(
      private userService: UserService,
      private partyService: PartyService,
  ) {
    this.$user = this.userService.user$
  }

  async ngOnInit() {
    await this.loadParties()
  }

  onImgError(event: any) {
    event.target.src = 'assets/default_profile_picture.png';
  }

  async loadParties() {
    this.$user.subscribe(user => {
      console.log(user)
      this.$myParties = this.partyService.getMyParties(user?.id!)
    })
  }
}
