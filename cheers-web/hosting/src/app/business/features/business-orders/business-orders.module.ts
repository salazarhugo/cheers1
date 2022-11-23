import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BusinessOrdersRoutingModule } from './business-orders-routing.module';
import { BusinessOrdersComponent } from './business-orders.component';
import {OrderItemModule} from "../../../manage-parties/ui/order-item/order-item.module";


@NgModule({
  declarations: [
    BusinessOrdersComponent
  ],
    imports: [
        CommonModule,
        BusinessOrdersRoutingModule,
        OrderItemModule
    ]
})
export class BusinessOrdersModule { }
