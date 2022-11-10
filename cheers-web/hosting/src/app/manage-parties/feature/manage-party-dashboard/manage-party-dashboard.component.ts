import { Component, OnInit } from '@angular/core';
import {PartyService} from "../../../parties/data/party.service";
import {Order} from "../../../shared/data/models/order.model";
import {firstValueFrom, Observable} from "rxjs";

@Component({
  selector: 'app-manage-party-dashboard',
  templateUrl: './manage-party-dashboard.component.html',
  styleUrls: ['./manage-party-dashboard.component.sass']
})
export class ManagePartyDashboardComponent implements OnInit {

  orders: Observable<Order[]>

  constructor(
      private partyService: PartyService,
  ) {
  }

  ngOnInit() {
    // const party = await firstValueFrom(this.partyService.getManagedParty())
    this.partyService.getManagedParty().subscribe(party => {
      this.orders = this.partyService.getPartyOrders(party.id)
    })
  }
}
