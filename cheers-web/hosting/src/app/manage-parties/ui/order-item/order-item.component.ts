import {Component, Input, OnInit} from '@angular/core';
import {Order} from "../../../../gen/ts/cheers/order/v1/order_service";

@Component({
  selector: 'app-order-item',
  templateUrl: './order-item.component.html',
  styleUrls: ['./order-item.component.sass']
})
export class OrderItemComponent implements OnInit {

  @Input() order: Order

  constructor() { }

  ngOnInit(): void {
  }

}
