import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TicketItemRoutingModule} from './ticket-item-routing.module';
import {TicketItemComponent} from "./ticket-item.component";
import {MatIconModule} from "@angular/material/icon";
import {FlexModule} from "@angular/flex-layout";
import {MatButtonModule} from "@angular/material/button";
import {MatMenuModule} from "@angular/material/menu";
import {MatCardModule} from "@angular/material/card";


@NgModule({
    declarations: [TicketItemComponent],
    imports: [
        CommonModule,
        TicketItemRoutingModule,
        MatButtonModule,
        MatIconModule,
        FlexModule,
        MatMenuModule,
        MatCardModule
    ],
    exports: [TicketItemComponent],
})
export class TicketItemModule {
}
