import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyDetailRoutingModule} from './party-detail-routing.module';
import {PartyComponent} from "./party.component";
import {MaterialModule} from "../../../material/material.module";
import {AdminButtonModule} from "../../../shared/ui/admin-button/admin-button.module";
import {MatLegacyRadioModule as MatRadioModule} from "@angular/material/legacy-radio";


@NgModule({
    declarations: [PartyComponent],
    imports: [
        CommonModule,
        PartyDetailRoutingModule,
        MaterialModule,
        AdminButtonModule,
        MatRadioModule
    ],
    exports: [
        PartyComponent,
    ]
})
export class PartyDetailModule {
}
