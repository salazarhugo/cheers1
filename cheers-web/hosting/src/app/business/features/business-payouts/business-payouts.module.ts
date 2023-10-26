import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BusinessPayoutsRoutingModule } from './business-payouts-routing.module';
import { BusinessPayoutsComponent } from './business-payouts.component';
import {MatLegacyCardModule as MatCardModule} from "@angular/material/legacy-card";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
  declarations: [
    BusinessPayoutsComponent
  ],
  imports: [
    CommonModule,
    BusinessPayoutsRoutingModule,
    MatCardModule,
    FlexModule
  ],
  exports: [
    BusinessPayoutsComponent
  ],
})
export class BusinessPayoutsModule { }
