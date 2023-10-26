import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PartyTicketCreateRoutingModule} from './party-ticket-create-routing.module';
import {PartyTicketCreateComponent} from './party-ticket-create.component';
import {MatLegacyFormFieldModule as MatFormFieldModule} from "@angular/material/legacy-form-field";
import {FlexModule} from "@angular/flex-layout";
import {ReactiveFormsModule} from "@angular/forms";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";
import {MatLegacyInputModule as MatInputModule} from "@angular/material/legacy-input";
import {MatDividerModule} from "@angular/material/divider";
import {MatButtonToggleModule} from "@angular/material/button-toggle";
import {MatLegacySlideToggleModule as MatSlideToggleModule} from "@angular/material/legacy-slide-toggle";
import {PriceModule} from "../../../shared/data/directives/price/price.module";


@NgModule({
    declarations: [
        PartyTicketCreateComponent
    ],
    imports: [
        CommonModule,
        PartyTicketCreateRoutingModule,
        MatFormFieldModule,
        FlexModule,
        ReactiveFormsModule,
        MatButtonModule,
        MatInputModule,
        MatDividerModule,
        MatButtonToggleModule,
        MatSlideToggleModule,
        PriceModule
    ],
    exports: [
        PartyTicketCreateComponent
    ],
})
export class PartyTicketCreateModule {
}
