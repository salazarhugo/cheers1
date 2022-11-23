import { Component, OnInit } from '@angular/core';
import {OrderService} from "../../../shared/data/services/order.service";

@Component({
  selector: 'app-business-orders',
  templateUrl: './business-orders.component.html',
  styleUrls: ['./business-orders.component.sass']
})
export class BusinessOrdersComponent implements OnInit {

  constructor(
      private orderService: OrderService
  ) { }

  async ngOnInit()  {
    const res = await this.orderService.getCurrentUserOrders()
    res.subscribe(res => console.log(res))
  }
}