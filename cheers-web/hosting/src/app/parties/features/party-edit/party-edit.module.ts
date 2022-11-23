import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyEditRoutingModule} from './party-edit-routing.module';
import {EditPartyComponent} from "./edit-party.component";
import {MaterialModule} from "../../../material/material.module";
import {PartyFormModule} from "../../ui/party-form/party-form.module";


@NgModule({
    declarations: [EditPartyComponent],
    imports: [
        CommonModule,
        PartyEditRoutingModule,
        MaterialModule,
        PartyFormModule,
    ],
    exports: [EditPartyComponent],
})
export class PartyEditModule {
}
