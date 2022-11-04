import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ActivityListRoutingModule} from './activity-list-routing.module';
import {ActivitiesComponent} from "./activities.component";
import {FlexModule} from "@angular/flex-layout";
import {UserItemModule} from "../../../users/ui/user-item/user-item.module";
import {ActivityItemModule} from "../../ui/activity-item/activity-item.module";
import {MatDividerModule} from "@angular/material/divider";


@NgModule({
    declarations: [ActivitiesComponent],
    imports: [
        CommonModule,
        ActivityListRoutingModule,
        FlexModule,
        UserItemModule,
        ActivityItemModule,
        MatDividerModule,
    ],
    exports: [ActivitiesComponent],
})
export class ActivityListModule {
}
