import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyTicketsRoutingModule} from './party-tickets-routing.module';
import {PartyTicketsComponent} from './party-tickets.component';
import {MatFormFieldModule} from "@angular/material/form-field";
import {TicketItemModule} from "../../../tickets/ui/ticket-item/ticket-item.module";
import {MatDividerModule} from "@angular/material/divider";
import {ReactiveFormsModule} from "@angular/forms";
import {FlexModule} from "@angular/flex-layout";
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {MatSidenavModule} from "@angular/material/sidenav";
import {PartyTicketCreateModule} from "../../ui/party-ticket-create/party-ticket-create.module";


@NgModule({
    declarations: [
        PartyTicketsComponent
    ],
    imports: [
        CommonModule,
        PartyTicketsRoutingModule,
        MatFormFieldModule,
        TicketItemModule,
        MatDividerModule,
        ReactiveFormsModule,
        FlexModule,
        MatInputModule,
        MatButtonModule,
        MatSidenavModule,
        PartyTicketCreateModule
    ],
    exports: [
        PartyTicketsComponent
    ]
})
export class PartyTicketsModule {
}
