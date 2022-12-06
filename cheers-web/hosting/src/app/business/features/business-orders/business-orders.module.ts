import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BusinessOrdersRoutingModule } from './business-orders-routing.module';
import { BusinessOrdersComponent } from './business-orders.component';
import {OrderItemModule} from "../../../manage-parties/ui/order-item/order-item.module";
import {MatTableModule} from "@angular/material/table";
import {MatSortModule} from "@angular/material/sort";
import {MatChipsModule} from "@angular/material/chips";
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {MatMenuModule} from "@angular/material/menu";
import {RelativeTimeModule} from "../../../shared/data/pipes/relative-time/relative-time.module";


@NgModule({
  declarations: [
    BusinessOrdersComponent
  ],
    imports: [
        CommonModule,
        BusinessOrdersRoutingModule,
        OrderItemModule,
        MatTableModule,
        MatSortModule,
        MatChipsModule,
        MatButtonModule,
        MatIconModule,
        MatMenuModule,
        RelativeTimeModule
    ]
})
export class BusinessOrdersModule { }
