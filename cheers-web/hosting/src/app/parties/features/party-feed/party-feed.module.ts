import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyFeedRoutingModule} from './party-feed-routing.module';
import {PartiesComponent} from "./parties.component";
import {PartyItemModule} from "../../ui/party-item/party-item.module";
import {MatLegacyChipsModule as MatChipsModule} from "@angular/material/legacy-chips";
import {MaterialModule} from "../../../material/material.module";


@NgModule({
    declarations: [PartiesComponent],
    imports: [
        CommonModule,
        PartyFeedRoutingModule,
        PartyItemModule,
        MatChipsModule,
        MaterialModule,
    ],
    exports: [PartiesComponent],
})
export class PartyFeedModule {
}
