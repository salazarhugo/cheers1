import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartySidenavRoutingModule } from './manage-party-sidenav-routing.module';
import { ManagePartySidenavComponent } from './manage-party-sidenav.component';
import {CoreModule} from "../../../core/core.module";
import {MatDividerModule} from "@angular/material/divider";


@NgModule({
  declarations: [
    ManagePartySidenavComponent
  ],
  imports: [
    CommonModule,
    ManagePartySidenavRoutingModule,
    CoreModule,
    MatDividerModule
  ],
  exports: [
    ManagePartySidenavComponent
  ],
})
export class ManagePartySidenavModule { }
