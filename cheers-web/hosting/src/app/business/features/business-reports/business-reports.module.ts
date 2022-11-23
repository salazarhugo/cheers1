import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BusinessReportsRoutingModule } from './business-reports-routing.module';
import { BusinessReportsComponent } from './business-reports.component';


@NgModule({
  declarations: [
    BusinessReportsComponent
  ],
  imports: [
    CommonModule,
    BusinessReportsRoutingModule
  ]
})
export class BusinessReportsModule { }
