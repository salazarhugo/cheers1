import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyBasicinfoRoutingModule } from './manage-party-basicinfo-routing.module';
import { ManagePartyBasicinfoComponent } from './manage-party-basicinfo.component';
import {PartyFormModule} from "../../../parties/ui/party-form/party-form.module";


@NgModule({
  declarations: [
    ManagePartyBasicinfoComponent
  ],
    imports: [
        CommonModule,
        ManagePartyBasicinfoRoutingModule,
        PartyFormModule
    ],
  exports: [
    ManagePartyBasicinfoComponent
  ],
})
export class ManagePartyBasicinfoModule { }
