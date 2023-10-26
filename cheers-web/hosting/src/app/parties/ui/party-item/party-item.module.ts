import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PartyItemComponent} from "./party-item.component";
import {MaterialModule} from "../../../material/material.module";
import {RouterModule} from "@angular/router";
import {PartyDateModule} from "../../../shared/data/pipes/party-date/party-date.module";
import {CoreModule} from "../../../core/core.module";

@NgModule({
    declarations: [PartyItemComponent],
    imports: [
        CommonModule,
        CoreModule,
        MaterialModule,
        RouterModule,
        PartyDateModule,
    ],
    exports: [PartyItemComponent],
})
export class PartyItemModule {
}
