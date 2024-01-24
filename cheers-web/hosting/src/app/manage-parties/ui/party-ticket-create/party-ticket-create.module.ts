import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PartyTicketCreateRoutingModule} from './party-ticket-create-routing.module';
import {PartyTicketCreateComponent} from './party-ticket-create.component';
import {MatFormFieldModule} from "@angular/material/form-field";
import {FlexModule} from "@angular/flex-layout";
import {ReactiveFormsModule} from "@angular/forms";
import {MatButtonModule} from "@angular/material/button";
import {MatInputModule} from "@angular/material/input";
import {MatDividerModule} from "@angular/material/divider";
import {MatButtonToggleModule} from "@angular/material/button-toggle";
import {MatSlideToggleModule} from "@angular/material/slide-toggle";
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
