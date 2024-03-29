import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ManagePartyDashboardRoutingModule} from './manage-party-dashboard-routing.module';
import {ManagePartyDashboardComponent} from './manage-party-dashboard.component';
import {MatCardModule} from "@angular/material/card";
import {FlexModule} from "@angular/flex-layout";
import {RelativeTimeModule} from "../../../shared/data/pipes/relative-time/relative-time.module";
import {OrderItemModule} from "../../ui/order-item/order-item.module";


@NgModule({
    declarations: [
        ManagePartyDashboardComponent
    ],
    imports: [
        CommonModule,
        ManagePartyDashboardRoutingModule,
        MatCardModule,
        FlexModule,
        RelativeTimeModule,
        OrderItemModule
    ],
    exports: [
        ManagePartyDashboardComponent
    ],
})
export class ManagePartyDashboardModule {
}
