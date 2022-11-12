import { Component, OnInit } from '@angular/core';
import {PartyService} from "../../../parties/data/party.service";
import {firstValueFrom, Observable} from "rxjs";
import {Order} from "../../../../gen/ts/cheers/order/v1/order_service";

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
