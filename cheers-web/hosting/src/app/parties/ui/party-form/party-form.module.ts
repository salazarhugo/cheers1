import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PartyFormComponent} from "./party-form.component";
import {MaterialModule} from "../../../material/material.module";
import {TicketItemModule} from "../../../tickets/ui/ticket-item/ticket-item.module";


@NgModule({
    declarations: [PartyFormComponent],
    imports: [
        CommonModule,
        MaterialModule,
        TicketItemModule,
    ],
    exports: [PartyFormComponent],
})
export class PartyFormModule {
}
