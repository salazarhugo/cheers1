import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PartyHostingRoutingModule } from './party-hosting-routing.module';
import { PartyHostingComponent } from './party-hosting.component';
import {PartyItemModule} from "../../ui/party-item/party-item.module";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
  declarations: [
    PartyHostingComponent
  ],
  imports: [
    CommonModule,
    PartyHostingRoutingModule,
    PartyItemModule,
    FlexModule
  ]
})
export class PartyHostingModule { }
