import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyBasicinfoRoutingModule } from './manage-party-basicinfo-routing.module';
import { ManagePartyBasicinfoComponent } from './manage-party-basicinfo.component';


@NgModule({
  declarations: [
    ManagePartyBasicinfoComponent
  ],
  imports: [
    CommonModule,
    ManagePartyBasicinfoRoutingModule
  ],
  exports: [
    ManagePartyBasicinfoComponent
  ],
})
export class ManagePartyBasicinfoModule { }
