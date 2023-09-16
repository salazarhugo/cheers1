import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyFeedRoutingModule} from './party-feed-routing.module';
import {PartiesComponent} from "./parties.component";
import {PartyItemModule} from "../../ui/party-item/party-item.module";
import {MatChipsModule} from "@angular/material/chips";
import {MaterialModule} from "../../../material/material.module";
import {SHIMMER_OPTIONS, ShimmerModule} from "@sreyaj/ng-shimmer";


@NgModule({
    declarations: [PartiesComponent],
    imports: [
        CommonModule,
        PartyFeedRoutingModule,
        PartyItemModule,
        MatChipsModule,
        MaterialModule,
        ShimmerModule,
    ],
    exports: [PartiesComponent],
})
export class PartyFeedModule {
}
