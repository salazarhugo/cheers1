import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyTicketsRoutingModule } from './manage-party-tickets-routing.module';
import { ManagePartyTicketsComponent } from './manage-party-tickets.component';
import {MatSidenavModule} from "@angular/material/sidenav";
import {TicketItemModule} from "../../../tickets/ui/ticket-item/ticket-item.module";
import {MatDividerModule} from "@angular/material/divider";
import {ReactiveFormsModule} from "@angular/forms";
import {FlexModule} from "@angular/flex-layout";
import {PartyTicketCreateModule} from "../../ui/party-ticket-create/party-ticket-create.module";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";


@NgModule({
  declarations: [
    ManagePartyTicketsComponent
  ],
  imports: [
    CommonModule,
    ManagePartyTicketsRoutingModule,
    MatFormFieldModule,
    TicketItemModule,
    MatDividerModule,
    ReactiveFormsModule,
    FlexModule,
    MatInputModule,
    MatButtonModule,
    MatSidenavModule,
    PartyTicketCreateModule,
  ]
})
export class ManagePartyTicketsModule { }
