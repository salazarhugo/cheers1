import { Component, OnInit } from '@angular/core';
import {OrderService} from "../../../shared/data/services/order.service";
import {Observable} from "rxjs";
import {Order} from "../../../../gen/ts/cheers/order/v1/order_service";

@Component({
  selector: 'app-business-orders',
  templateUrl: './business-orders.component.html',
  styleUrls: ['./business-orders.component.sass']
})
export class BusinessOrdersComponent implements OnInit {

  $orders: Observable<Order[]>

  constructor(
      private orderService: OrderService
  ) { }

  async ngOnInit()  {
    const res = await this.orderService.getCurrentUserOrders()
    this.$orders = res
  }
}