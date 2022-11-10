import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ManagePartyDashboardRoutingModule} from './manage-party-dashboard-routing.module';
import {ManagePartyDashboardComponent} from './manage-party-dashboard.component';
import {MatCardModule} from "@angular/material/card";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
    declarations: [
        ManagePartyDashboardComponent
    ],
    imports: [
        CommonModule,
        ManagePartyDashboardRoutingModule,
        MatCardModule,
        FlexModule
    ],
    exports: [
        ManagePartyDashboardComponent
    ],
})
export class ManagePartyDashboardModule {
}
