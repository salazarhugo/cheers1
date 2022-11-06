import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyRoutingModule } from './manage-party-routing.module';
import { ManagePartyComponent } from './manage-party.component';
import {TopbarModule} from "../../../home/ui/topbar/topbar.module";
import {MatSidenavModule} from "@angular/material/sidenav";
import {CoreModule} from "../../../core/core.module";
import {RouterModule} from "@angular/router";
import {ManagePartySidenavModule} from "../../ui/manage-party-sidenav/manage-party-sidenav.module";


@NgModule({
  declarations: [
    ManagePartyComponent
  ],
    imports: [
        CommonModule,
        RouterModule,
        ManagePartyRoutingModule,
        TopbarModule,
        MatSidenavModule,
        CoreModule,
        ManagePartySidenavModule
    ],
    exports: [
        ManagePartyComponent
    ],
})
export class ManagePartyModule { }
