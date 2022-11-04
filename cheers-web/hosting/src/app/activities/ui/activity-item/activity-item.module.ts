import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ActivityItemRoutingModule} from './activity-item-routing.module';
import {ActivityItemComponent} from "./activity-item.component";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
    declarations: [ActivityItemComponent],
    imports: [
        CommonModule,
        ActivityItemRoutingModule,
        FlexModule
    ],
    exports: [ActivityItemComponent],
})
export class ActivityItemModule {
}
