import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyDetailsRoutingModule } from './manage-party-details-routing.module';
import { ManagePartyDetailsComponent } from './manage-party-details.component';


@NgModule({
  declarations: [
    ManagePartyDetailsComponent
  ],
  imports: [
    CommonModule,
    ManagePartyDetailsRoutingModule
  ]
})
export class ManagePartyDetailsModule { }
