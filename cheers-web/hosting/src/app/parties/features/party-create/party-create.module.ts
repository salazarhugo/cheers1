import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyCreateRoutingModule} from './party-create-routing.module';
import {CreatePartyComponent} from "./create-party.component";
import {PartyFormModule} from "../../ui/party-form/party-form.module";


@NgModule({
    declarations: [CreatePartyComponent],
    imports: [
        CommonModule,
        PartyCreateRoutingModule,
        PartyFormModule,
    ],
    exports: [CreatePartyComponent]
})
export class PartyCreateModule {
}
