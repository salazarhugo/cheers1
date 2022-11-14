import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { OrderItemRoutingModule } from './order-item-routing.module';
import { OrderItemComponent } from './order-item.component';


@NgModule({
  declarations: [
    OrderItemComponent
  ],
  imports: [
    CommonModule,
    OrderItemRoutingModule
  ],
  exports: [
    OrderItemComponent
  ],
})
export class OrderItemModule { }
