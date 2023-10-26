import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { OrderItemRoutingModule } from './order-item-routing.module';
import { OrderItemComponent } from './order-item.component';
import {RelativeTimeModule} from "../../../shared/data/pipes/relative-time/relative-time.module";
import {MatLegacyCardModule as MatCardModule} from "@angular/material/legacy-card";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
  declarations: [
    OrderItemComponent
  ],
    imports: [
        CommonModule,
        OrderItemRoutingModule,
        RelativeTimeModule,
        MatCardModule,
        FlexModule
    ],
  exports: [
    OrderItemComponent
  ],
})
export class OrderItemModule { }
