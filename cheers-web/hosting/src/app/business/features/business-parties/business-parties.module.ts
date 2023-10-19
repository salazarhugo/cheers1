import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BusinessPartiesRoutingModule } from './business-parties-routing.module';
import { BusinessPartiesComponent } from './business-parties.component';
import {PartyItemModule} from "../../../parties/ui/party-item/party-item.module";
import {FlexModule} from "@angular/flex-layout";
import {SharedModule} from "../../../shared/shared.module";
import {MaterialModule} from "../../../material/material.module";


@NgModule({
  declarations: [
    BusinessPartiesComponent
  ],
  imports: [
    CommonModule,
    BusinessPartiesRoutingModule,
    PartyItemModule,
    FlexModule,
      SharedModule,
      MaterialModule,
  ]
})
export class BusinessPartiesModule { }
