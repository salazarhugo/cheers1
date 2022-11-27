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
  displayedColumns: string[] = ['status', 'date', 'amount', 'party','firstName', 'lastName', 'email'];
  orders: Order[]

  constructor(
      private orderService: OrderService
  ) { }

  async ngOnInit()  {
    const res = await this.orderService.getCurrentUserOrders()
    res.subscribe(data => {
      this.orders = data
    })
  }

  /** Gets the total amount of all orders. */
  getTotalCost() {
    return this.orders.map(t => t.amount).reduce((acc, value) => acc + value, 0);
  }
}