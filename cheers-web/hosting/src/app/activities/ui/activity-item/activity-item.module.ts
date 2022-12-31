import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ActivityItemRoutingModule} from './activity-item-routing.module';
import {ActivityItemComponent} from "./activity-item.component";
import {FlexModule} from "@angular/flex-layout";
import {RelativeTimeModule} from "../../../shared/data/pipes/relative-time/relative-time.module";


@NgModule({
    declarations: [ActivityItemComponent],
    imports: [
        CommonModule,
        ActivityItemRoutingModule,
        FlexModule,
        RelativeTimeModule
    ],
    exports: [ActivityItemComponent],
})
export class ActivityItemModule {
}
