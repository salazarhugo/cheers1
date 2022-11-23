import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BusinessOrdersRoutingModule } from './business-orders-routing.module';
import { BusinessOrdersComponent } from './business-orders.component';


@NgModule({
  declarations: [
    BusinessOrdersComponent
  ],
  imports: [
    CommonModule,
    BusinessOrdersRoutingModule
  ]
})
export class BusinessOrdersModule { }
