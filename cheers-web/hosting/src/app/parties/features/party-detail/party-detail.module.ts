import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyDetailRoutingModule} from './party-detail-routing.module';
import {PartyComponent} from "./party.component";
import {MaterialModule} from "../../../material/material.module";
import {MatRadioModule} from "@angular/material/radio";
import {AdminButtonComponent} from "../../../shared/ui/admin-button/admin-button.component";
import {CheersButtonComponent} from "../../../shared/ui/cheers-button/cheers-button.component";


@NgModule({
    declarations: [PartyComponent],
    imports: [
        CommonModule,
        PartyDetailRoutingModule,
        MaterialModule,
        MatRadioModule,
        AdminButtonComponent,
        CheersButtonComponent,
    ],
    exports: [
        PartyComponent,
    ]
})
export class PartyDetailModule {
}
