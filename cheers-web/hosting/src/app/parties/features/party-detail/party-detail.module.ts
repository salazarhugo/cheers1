import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyDetailRoutingModule} from './party-detail-routing.module';
import {PartyComponent} from "./party.component";
import {MaterialModule} from "../../../material/material.module";


@NgModule({
    declarations: [PartyComponent],
    imports: [
        CommonModule,
        PartyDetailRoutingModule,
        MaterialModule
    ],
    exports: [
        PartyComponent,
    ]
})
export class PartyDetailModule {
}
